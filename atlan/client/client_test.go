package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Backup and clear environmental variables
	apiKey := os.Getenv("ATLAN_API_KEY")
	baseURL := os.Getenv("ATLAN_BASE_URL")
	defer func() {
		os.Setenv("ATLAN_API_KEY", apiKey)
		os.Setenv("ATLAN_BASE_URL", baseURL)
	}()
	os.Setenv("ATLAN_API_KEY", "mock_api_key")
	os.Setenv("ATLAN_BASE_URL", "https://example.com")

	err := Init()
	if err != nil {
		t.Errorf("Unexpected error initializing AtlanClient: %v", err)
	}
	if DefaultAtlanClient == nil {
		t.Error("DefaultAtlanClient is not initialized")
	}
}

func TestCallAPI(t *testing.T) {
	// Mock server to simulate API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	//logger := log.New(os.Stdout, "Test Logger: ", log.LstdFlags)

	// Create a new AtlanClient instance
	atlanClient := &AtlanClient{
		Session:        http.DefaultClient,
		host:           mockServer.URL,
		ApiKey:         "mock_api_key",
		loggingEnabled: false,
		//	logger:         logger,
		requestParams: make(map[string]interface{}),
	}

	// Define an API struct for testing
	api := &API{
		Method:   http.MethodGet,
		Path:     "/test",
		Status:   http.StatusOK,
		Endpoint: struct{ Atlas string }{Atlas: "/v1"},
	}

	// Make a call to the API
	response, err := atlanClient.CallAPI(api, nil, nil)
	if err != nil {
		t.Errorf("Error making API call: %v", err)
	}

	// Check if the response is correct
	expectedResponse := `{"message": "success"}`
	if string(response) != expectedResponse {
		t.Errorf("Unexpected response. Expected: %s, Got: %s", expectedResponse, string(response))
	}
}

func TestMakeRequest(t *testing.T) {
	// Create a new AtlanClient instance
	atlanClient := &AtlanClient{
		Session:        http.DefaultClient,
		host:           "http://example.com",
		ApiKey:         "mock_api_key",
		loggingEnabled: true,
		requestParams:  make(map[string]interface{}),
	}

	// Test GET request
	req, err := atlanClient.makeRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
	}
	if req.Request.Method != http.MethodGet {
		// Test POST request
		req, err = atlanClient.makeRequest(http.MethodPost, "http://example.com", map[string]interface{}{
			"data": "test data",
		})
		if err != nil {
			t.Errorf("Error creating POST request: %v", err)
		}
		if req.Request.Method != http.MethodPost {
			t.Errorf("Unexpected request method. Expected: %s, Got: %s", http.MethodPost, req.Request.Method)
		}
	}
}
