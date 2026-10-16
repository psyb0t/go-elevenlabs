#!/bin/bash
# Rewrites the downloaded OpenAPI document into api/openapi.patched.json, the
# file oapi-codegen reads. Two spec shapes produce Go that does not compile:
#
#   - A schema that carries both anyOf and oneOf over the same members makes
#     oapi-codegen emit the As/From/Merge union methods twice. The oneOf is
#     dropped; the anyOf lists every member, so it is kept.
#   - A component schema named like an operation's generated parameter type,
#     such as ListAssetsParams next to operation list_assets, is declared twice.
#     Such schemas get a "Schema" suffix and every $ref to them is rewritten.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR
readonly SOURCE_SPEC="${SPEC_DIR}/openapi.json"
readonly OUTPUT_SPEC="${SPEC_DIR}/openapi.patched.json"
readonly SCHEMA_REF_PREFIX="#/components/schemas/"
readonly RENAME_SUFFIX="Schema"

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

main() {
	if [[ ! -f "$SOURCE_SPEC" ]]; then
		log ERROR "source spec missing path=${SOURCE_SPEC}"
		exit 1
	fi

	local tmp
	tmp="$(mktemp "${SPEC_DIR}/.openapi.patched.json.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp now, the local is gone when EXIT fires
	trap "rm -f -- '${tmp}'" EXIT

	jq --arg prefix "$SCHEMA_REF_PREFIX" --arg suffix "$RENAME_SUFFIX" '
		def pascal: split("_") | map((.[:1] | ascii_upcase) + .[1:]) | join("");

		([.paths[][] | objects | .operationId // empty | pascal + "Params"]) as $param_types
		| ([.components.schemas | keys[] | select(. as $name | $param_types | index($name))]) as $clashes
		| reduce $clashes[] as $name (.;
			.components.schemas[$name + $suffix] = .components.schemas[$name]
			| del(.components.schemas[$name])
		)
		| walk(
			if type == "object" then
				(if has("anyOf") and has("oneOf") then del(.oneOf) else . end)
				| (if (."$ref" | type) == "string" and ((."$ref" | ltrimstr($prefix)) as $n | $clashes | index($n)) then
					."$ref" += $suffix
				else
					.
				end)
			else
				.
			end
		)
	' "$SOURCE_SPEC" >"$tmp"

	chmod 0644 "$tmp"
	mv -- "$tmp" "$OUTPUT_SPEC"
	log INFO "patched spec path=${OUTPUT_SPEC}"
}

main "$@"
