package elevenlabs_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-elevenlabs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey       = "test-key"
	headerAPIKey     = "xi-api-key"
	headerContent    = "Content-Type"
	contentTypeJSON  = "application/json"
	pathUser         = "/v1/user"
	pathVoices       = "/v1/voices"
	userResponseBody = `{"user_id":"u-1","subscription":{},` +
		`"is_new_user":false,"xi_api_key":"k",` +
		`"can_use_delayed_payment_methods":false,` +
		`"is_onboarding_completed":true,` +
		`"is_onboarding_checklist_completed":true,` +
		`"created_at":0,"first_name":"Ada"}`
	voicesResponseBody = `{"voices":[{"voice_id":"v-1","name":"Rachel"}]}`
)

func newTestClient(
	t *testing.T,
	wantPath, body string,
) *elevenlabs.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, wantPath, r.URL.Path)
			assert.Equal(t, testAPIKey, r.Header.Get(headerAPIKey))

			w.Header().Set(headerContent, contentTypeJSON)
			_, err := w.Write([]byte(body))
			assert.NoError(t, err)
		},
	))
	t.Cleanup(srv.Close)

	auth := elevenlabs.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(headerAPIKey, testAPIKey)

			return nil
		},
	)

	client, err := elevenlabs.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestGetUserInfo(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, pathUser, userResponseBody)

	res, err := client.GetUserInfoWithResponse(t.Context(), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, "u-1", res.JSON200.UserId)
	require.NotNil(t, res.JSON200.FirstName)
	assert.Equal(t, "Ada", *res.JSON200.FirstName)
}

func TestGetVoices(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, pathVoices, voicesResponseBody)

	res, err := client.GetVoicesWithResponse(t.Context(), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Voices, 1)

	assert.Equal(t, "v-1", res.JSON200.Voices[0].VoiceId)
	assert.Equal(t, "Rachel", res.JSON200.Voices[0].Name)
}
