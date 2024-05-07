package client

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestEnvConfig(t *testing.T) {
	// Set up environment variables
	apiKey := "your_api_key"
	baseURL := "your_base_url"

	os.Getenv("ATLAN_API_KEY")
	os.Getenv("ATLAN_BASE_URL")
	
	// Initialize client
	err := Init()
	assert.NoError(t, err)

	// Assert API key and base URL are correctly set
	assert.Equal(t, apiKey, DefaultAtlanClient.ApiKey)
	assert.Equal(t, baseURL, DefaultAtlanClient.host)
}

func TestEnvConfigUsingContext(t *testing.T) {
	// Set up environment variables
	apiKey := "your_api_key"
	baseURL := "your_base_url"

	// Initialize client
	ctx, err := Context(apiKey, baseURL)

	assert.NoError(t, err)

	// Assert API key and base URL are correctly set
	assert.Equal(t, apiKey, ctx.ApiKey)
	assert.Equal(t, baseURL, ctx.host)
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
	ctx, _ := Context("api_key", ts.URL)

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
	assert.NoError(t, err)

	// Define expected response as JSON object
	expectedResponse := map[string]string{"message": "success"}

	// Unmarshal actual response into JSON object
	var actualResponse map[string]string
	err = json.Unmarshal(response, &actualResponse)
	assert.NoError(t, err)

	// Compare expected and actual responses
	assert.Equal(t, expectedResponse, actualResponse)
}
