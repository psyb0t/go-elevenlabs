# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `elevenlabs_test`, and use only the exported API.
- **TST2** Tests never call ElevenLabs. They point the client at an `httptest` server and answer with a canned status and body.
- **TST3** Every test calls `t.Parallel()`. Assert the request method, path and `xi-api-key` header, then the decoded `JSONnnn` field.
- **TST4** `example_test.go` holds the README usage example so `make test` compiles it. Keep the two in sync.
- **TST5** `make test` runs `go test -race` and must pass.
