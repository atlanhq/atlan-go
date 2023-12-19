package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallAPI(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request method and path
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, but got %s", r.Method)
		}
		if r.URL.Path != "/test" {
			t.Errorf("Expected path '/test', but got %s", r.URL.Path)
		}

		// Write a sample response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	// Create an AtlanClient for testing
	client := NewAtlanClient("test-api-key", mockServer.URL)

	// Test GET request
	response, err := client.CallAPI("/test", http.MethodGet, nil, nil)
	if err != nil {
		t.Fatalf("Error calling GET API: %v", err)
	}
	expectedResponse := []byte(`{"message": "success"}`)
	if string(response) != string(expectedResponse) {
		t.Errorf("Expected response %s, but got %s", expectedResponse, response)
	}
}
