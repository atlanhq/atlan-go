package assets

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

func TestEnvConfig(t *testing.T) {
	// Save the current environment variables
	originalAPIKey := os.Getenv("ATLAN_API_KEY")
	originalBaseURL := os.Getenv("ATLAN_BASE_URL")

	// Set up environment variables for the duration of the test
	os.Setenv("ATLAN_API_KEY", "your_api_key")
	os.Setenv("ATLAN_BASE_URL", "your_base_url")

	// Clean up environment variables after the test
	defer func() {
		os.Setenv("ATLAN_API_KEY", originalAPIKey)
		os.Setenv("ATLAN_BASE_URL", originalBaseURL)
	}()

	// Initialize client
	err := Init()
	require.NoError(t, err)

	// Assert API key and base URL are correctly set
	assert.Equal(t, "your_api_key", DefaultAtlanClient.ApiKey)
	assert.Equal(t, "https://your_base_url", DefaultAtlanClient.host)
}

func TestEnvConfigUsingContext(t *testing.T) {
	// Set up environment variables
	apiKey := "your_api_key"
	baseURL := "https://your_base_url"

	// Initialize client
	ctx, err := Context(baseURL, apiKey)

	require.NoError(t, err)

	// Assert API key and base URL are correctly set
	assert.Equal(t, apiKey, ctx.ApiKey)
	assert.Equal(t, baseURL, ctx.host)
}

func TestContextWithNormalizedURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"devx9.atlan.com", "https://devx9.atlan.com"},
		{"https://devx9.atlan.com", "https://devx9.atlan.com"},
		{"devx9.atlan.com/some/path", "https://devx9.atlan.com"},
		{"devx9.atlan.com?query=param", "https://devx9.atlan.com"},
		{"https://devx9.atlan.com/some/path", "https://devx9.atlan.com"},
	}

	apiKey := "your_api_key"

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// Initialize client context with the test input
			ctx, err := Context(test.input, apiKey)

			require.NoError(t, err)

			// Ensure the base URL is normalized correctly
			assert.Equal(t, test.expected, ctx.host)
		})
	}
}

func TestLoggerConfig(t *testing.T) {
	ctx, _ := Context("api_key", "baseurl")

	// Initialize logger with enabled logging
	ctx.SetLogger(true, "info")

	// Verify that logger is created and logging is enabled
	assert.NotNil(t, ctx.logger.Log)

	// Verify that logs are produced when logging is enabled
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := slog.New(handler)
	ctx.logger.Log = logger
	ctx.logger.Infof("Test log message")
	assert.Contains(t, buf.String(), "Test log message")

	// Reset buffer
	buf.Reset()

	// Initialize logger with disabled logging
	ctx.SetLogger(false, "")
	ctx.DisableLogging()

	// Verify that logger is created and logging is disabled
	assert.NotNil(t, ctx.logger.Log)

	// Verify that no logs are produced when logging is disabled
	ctx.logger.Infof("Test log message")
	assert.Empty(t, buf.String())
}

func TestCallAPI(t *testing.T) {
	// Create a new test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a sample JSON response
		response := map[string]string{"message": "success"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Initialize AtlanClient with the test server URL
	ctx, _ := Context(ts.URL, "api_key")

	// Define a sample API object
	api := &API{
		Method:   http.MethodGet,
		Endpoint: Endpoint{Atlas: "/test"},
		Path:     "/endpoint",
		Status:   http.StatusOK,
	}

	// Call the API
	response, err := ctx.CallAPI(api, nil, nil)

	// Check if there's no error
	require.NoError(t, err)

	// Define expected response as JSON object
	expectedResponse := map[string]string{"message": "success"}

	// Unmarshal actual response into JSON object
	var actualResponse map[string]string
	err = json.Unmarshal(response, &actualResponse)
	require.NoError(t, err)

	// Compare expected and actual responses
	assert.Equal(t, expectedResponse, actualResponse)
}
