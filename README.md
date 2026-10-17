# go-elevenlabs

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-elevenlabs.svg)](https://pkg.go.dev/github.com/psyb0t/go-elevenlabs)
[![CI](https://github.com/psyb0t/go-elevenlabs/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-elevenlabs/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-elevenlabs/badges/version.svg)](https://github.com/psyb0t/go-elevenlabs/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-elevenlabs/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-elevenlabs/badges/importers.svg)](https://github.com/psyb0t/go-elevenlabs/blob/badges/importers.md)

Unofficial Go client for [ElevenLabs](https://elevenlabs.io/docs/api-reference), generated from the OpenAPI document ElevenLabs publishes.

| Directory | Package | API | Generator |
|---|---|---|---|
| repo root | `elevenlabs` | REST, `https://api.elevenlabs.io` | [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) |

`client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-elevenlabs
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. ElevenLabs authenticates with one header, `xi-api-key`, taken from the API keys page of your account settings. Set it once with a `RequestEditorFn`:

```go
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-elevenlabs"
)

func main() {
	auth := elevenlabs.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("xi-api-key", os.Getenv("ELEVENLABS_API_KEY"))

		return nil
	})

	client, err := elevenlabs.NewClientWithResponses("https://api.elevenlabs.io", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.GetVoicesWithResponse(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The spec also lists `xi-api-key` as an optional header parameter on most operations, so most `...Params` structs have an `XiApiKey` field. Leave it nil and let the request editor set the header. Request editors run after the params are applied, so a client-wide editor overrides a per-call `XiApiKey`. The same code lives in `example_test.go`, so it is compiled by `make test`.

The full API, with the spec's descriptions as doc comments, is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-elevenlabs) and in `go doc github.com/psyb0t/go-elevenlabs`.

## Agent integrations

The [skill](.agents/skills/go-elevenlabs) teaches a coding agent how to use go-elevenlabs when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-elevenlabs@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-elevenlabs@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-elevenlabs
```

## Regenerating

The ElevenLabs spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

1. `scripts/fetch-spec.sh` downloads `https://api.elevenlabs.io/openapi.json` to `api/openapi.json`.
2. `scripts/patch-spec.sh` writes `api/openapi.patched.json`, fixing two spec shapes that otherwise produce Go that does not compile:
   - Schemas that carry both `anyOf` and `oneOf` over the same members (the speech-to-text responses) make oapi-codegen emit each union method twice. The `oneOf` is dropped and the `anyOf` kept.
   - A component schema named like an operation's generated parameter type (`ListAssetsParams` next to operation `list_assets`) is declared twice. Such schemas get a `Schema` suffix and every `$ref` to them is rewritten.
3. oapi-codegen runs with `config.yaml` on the patched document.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from the API specification ElevenLabs publishes. The specification itself is not redistributed here. I am not affiliated with ElevenLabs.
