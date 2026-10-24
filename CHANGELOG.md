# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.5.1, 2026-10-24

- Adds the `AgentDefinitionSourceArchitect` value. `AgentDefinitionSource.Valid` now accepts `"architect"`.
- Adds `MobileFullscreen` to `WidgetConfigInput`, `WidgetConfigOutput` and `WidgetConfigResponseModel`.
- Marks the dubbing operations `ListDubs`, `CreateDubbing`, `DeleteDubbing`, `GetDubbedMetadata`, `GetDubbedFile` and `GetDubbingTranscripts` as deprecated upstream, along with their `WithResponse` variants and `CreateDubbingMultipartRequestBody`. The methods still work.
- Updates the doc comment on the agent duplicate methods: the new agent now gets a copy of the live traffic split, with any share held by an uncopied branch moved to main.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

## v0.5.0, 2026-10-23

- Breaking: `OwnerView` is removed from `ListProductsParams` and `ListServicesParams`. Callers that set it must change their code. The owner listings now use the new `OwnerListProductsParams` and `OwnerListServicesParams` tool types.
- Breaking: `LLMInfoModel` is no longer comparable, because the new `AvailableServiceTiers` field is a slice. Code that compares `LLMInfoModel` values with `==` or uses them as map keys must change.
- `LLMInfoModel.AvailableServiceTiers` is a required field with no `omitempty`.
- Adds the `LLMServiceTier` type with `LLMServiceTierStandard` and `LLMServiceTierFast`, and a `ServiceTier` field on `LLMUsageCalculatorRequestModel`, `LLMUsageCalculatorPublicRequestModel`, `PromptAgentAPIModelInput`, `PromptAgentAPIModelOutput`, `PromptAgentAPIModelWorkflowOverrideInput` and `PromptAgentAPIModelWorkflowOverrideOutput`.
- Adds the owner SMB tool types `OwnerGetClientAppointmentsParams`, `OwnerListAssetsParams`, `OwnerListLocationsParams`, `OwnerListProductsParams`, `OwnerListServicesParams` and `OwnerListStaffParams`, each with a `SmbToolType` enum. `SMBToolConfig_Params` gains matching `As...`, `From...` and `Merge...` methods for each.
- Adds `StartProcedureIds` to `ConversationInitiationClientDataInternal`, `ConversationInitiationClientDataRequestInput` and `ConversationInitiationClientDataRequestOutput`. The doc comment on `EnableProcedureIdsFromClient` now covers `start_procedure_ids`.
- Adds `PerResourceTimezones` to `CreateAssetParams` and `UpdateAssetParams`.
- Adds the `AgentDeploymentSourceDuplicate` value.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

## v0.4.0, 2026-10-22

- Breaking: `RebaseBranchOntoMain` now takes a JSON request body. `Client.RebaseBranchOntoMain`, `ClientInterface.RebaseBranchOntoMain`, `ClientWithResponses.RebaseBranchOntoMainWithResponse` and `NewRebaseBranchOntoMainRequest` gain a `RebaseBranchOntoMainJSONRequestBody` argument after the params, so callers must change their code.
- Adds `RebaseBranchOntoMainWithBody`, `RebaseBranchOntoMainWithBodyWithResponse` and `NewRebaseBranchOntoMainRequestWithBody` for sending a raw body. The body type is `BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost`.
- Adds `UpdateTestInvocationRoute` and `UpdateTestInvocationRouteWithBody` to `Client` and `ClientInterface`, with `UpdateTestInvocationRouteWithResponse` and `UpdateTestInvocationRouteWithBodyWithResponse` on `ClientWithResponses` and `ClientWithResponsesInterface`. The call is PATCH `/v1/convai/test-invocations/{test_invocation_id}`.
- Adds `UpdateTestInvocationRouteParams`, `UpdateTestInvocationRouteJSONRequestBody`, `UpdateTestInvocationRouteResponse`, `UpdateTestSuiteInvocationRequestModel`, `NewUpdateTestInvocationRouteRequest`, `NewUpdateTestInvocationRouteRequestWithBody` and `ParseUpdateTestInvocationRouteResponse`.
- Adds `Name` to `RunAgentTestsRequestModel`, `GetTestSuiteInvocationResponseModel` and `TestInvocationSummaryResponseModel`.
- The minimum Go version in `go.mod` is now 1.26.10, which picks up the fix for GO-2026-6809 in `crypto/x509`. The development image and the setup notes use the same version.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

## v0.3.0, 2026-10-21

- Breaking: the speech to speech multipart body types are renamed. `BodySpeechToSpeechV1SpeechToSpeechVoiceIdPost` becomes `BodyVoiceChangerV1VoiceChangerVoiceIdPost`, and `BodySpeechToSpeechStreamingV1SpeechToSpeechVoiceIdStreamPost` becomes `BodyVoiceChangerStreamingV1VoiceChangerVoiceIdStreamPost`. Callers that use the old names must change their code.
- Breaking: the matching `FileFormat` types are renamed the same way, including the `Other` and `PcmS16le16` constants, for example `BodyVoiceChangerV1VoiceChangerVoiceIdPostFileFormat` and `BodyVoiceChangerStreamingV1VoiceChangerVoiceIdStreamPostFileFormatPcmS16le16`.
- Adds `ExcludePatterns` to `BodyCreateCrawlJobV1ConvaiKnowledgeBaseCrawlPost` and `GetCrawlJobResponseModel`.
- Adds `OwnerView` to `ListProductsParams` and `ListServicesParams`.
- Adds `CloneStrength` to `VoiceSettingsResponseModel` and `ToDialogueSettingsResponseModel`.
- Adds `VoiceSurchargeCharge` to `ConversationChargingCommonModel`.
- Adds the `SFXModelIdElevenTextToSoundV3` value.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

## v0.2.0, 2026-10-20

- Breaking: `SoundGenerationParams.OutputFormat` changes from `*AllowedOutputFormats` to an anonymous union struct, so callers that set it must change their code.
- Breaking: the `ToolTypeApiIntegrationMcp` constant is removed from `ToolType`, and `ToolType.Valid` no longer accepts `"api_integration_mcp"`.
- Adds `CreateRealtimeSession` and `CreateRealtimeSessionWithBody` to `Client` and `ClientInterface`, with `CreateRealtimeSessionWithResponse` and `CreateRealtimeSessionWithBodyWithResponse` on `ClientWithResponses` and `ClientWithResponsesInterface`. New types include `CreateRealtimeSessionParams`, `CreateRealtimeSessionJSONBody`, `CreateRealtimeSessionResponse`, `CreateClientSecretRequest`, `ClientSecretResponse`, `ExpiresAfter` and `RealtimeModel`.
- Adds the OpenAI realtime session types: `OpenAISessionConfigInput`, `OpenAISessionConfigOutput`, the audio config types, `OpenAIFunctionTool`, `OpenAIMCPTool`, `OpenAITurnDetection`, `OpenAITracingConfiguration`, `AudioFormat`, `AudioInputPacing`, `Eagerness` and `TurnDetectionType`.
- Adds `DisableServiceAccountApiKey` and `DisableServiceAccountApiKeyWithResponse` with `DisableServiceAccountApiKeyParams` and `DisableServiceAccountApiKeyResponse`.
- Adds the `AgentSortByUpdatedAt` value, `UpdatedAtUnixSecs` on `AgentSummaryResponseModel` and `SpeechEngineSummaryResponse`, `Description` on `BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch`, `Force` on `DeleteAgentRouteParams`, and `SourceBranchId` on `RebaseBranchOntoMainParams` and `RebasePreviewRouteParams`.
- Adds `VoiceSharingResponseModelReviewStatusDeclinedResubmittable`.
- Regenerated from the ElevenLabs OpenAPI spec, still version 1.0.

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
