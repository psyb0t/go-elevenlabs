// Package elevenlabs is a Go client for the ElevenLabs REST API, generated
// with oapi-codegen from the OpenAPI document ElevenLabs publishes.
package elevenlabs

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/patch-spec.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.patched.json
