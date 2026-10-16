package elevenlabs_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-elevenlabs"
)

// The README usage example. It compiles as part of the test build; it has no
// Output comment, so it is not run against the real API.
//
//nolint:testableexamples // needs a real API key, compile check only
func Example() {
	auth := elevenlabs.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("xi-api-key", os.Getenv("ELEVENLABS_API_KEY"))

			return nil
		},
	)

	client, err := elevenlabs.NewClientWithResponses(
		"https://api.elevenlabs.io",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.GetVoicesWithResponse(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
