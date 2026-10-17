# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.4, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.
- The README says most, not every, `...Params` struct has an `XiApiKey` field.

## v0.1.3, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.2, 2026-10-17

- The client adds `ListEnvironments`, `CreateEnvironment` (and `CreateEnvironmentWithBody`) and `DeleteEnvironment` to `Client`, `ClientInterface`, `ClientWithResponses` and `ClientWithResponsesInterface`, with the matching `WithResponse` methods.
- These endpoints add the `CreateEnvironmentRequest`, `EnvironmentResponse`, `EnvironmentsListResponse`, `ListEnvironmentsParams`, `CreateEnvironmentParams`, `DeleteEnvironmentParams`, `ListEnvironmentsResponse`, `CreateEnvironmentResponse` and `DeleteEnvironmentResponse` types.
- The doc comment on `UpdateEnvironmentVariableRequest.Values` and `UpdateEnvironmentVariable` now says that null removes the value for an environment, not the environment itself, and points to the new `DeleteEnvironment` method to remove an environment.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

## v0.1.1, 2026-10-16

- Added the `go-elevenlabs` agent skill. It installs with `claude plugin install go-elevenlabs@psyb0t` in Claude Code, `codex plugin add go-elevenlabs@psyb0t` in Codex and `openclaw skills install @psyb0t/go-elevenlabs` from ClawHub.

## v0.1.0, 2026-10-16

- Added `elevenlabs`, a REST client generated with oapi-codegen from the ElevenLabs OpenAPI 3.1 document (spec version 1.0, 318 paths).
- `make generate` downloads the current spec from `api.elevenlabs.io`, patches the two shapes that break the generated code (duplicate `anyOf`/`oneOf` unions, a schema named like an operation's params type), and regenerates the client. The spec is not committed.
