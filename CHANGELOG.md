# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `elevenlabs`, a REST client generated with oapi-codegen from the ElevenLabs OpenAPI 3.1 document (spec version 1.0, 318 paths).
- `make generate` downloads the current spec from `api.elevenlabs.io`, patches the two shapes that break the generated code (duplicate `anyOf`/`oneOf` unions, a schema named like an operation's params type), and regenerates the client. The spec is not committed.
