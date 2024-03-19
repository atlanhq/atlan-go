package client

import (
	"atlan-go/atlan"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGlossaryByGuid(t *testing.T) {
	// Define test cases with input GUID and expected attributes
	testCases := []struct {
		GUID             string
		ExpectedName     string
		ExpectedQualName string
		ExpectedIcon     atlan.AtlanIcon
	}{
		{
			GUID:             "fc36342b-ddb5-44ba-b774-4c90cc66d5a2",
			ExpectedName:     "Test Glossary",
			ExpectedQualName: "test_glossary",
			ExpectedIcon:     atlan.AtlanIconAirplaneInFlight,
		},
		// Add more test cases here if needed
	}

	// Mock server to simulate API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL to determine the response
		if r.URL.Path == "/api/meta/entity/guid/fc36342b-ddb5-44ba-b774-4c90cc66d5a2" {
			// Respond with a mock glossary JSON
			response := `{
				"referredEntities":{},
				"entity":{
					"typeName":"AtlasGlossary",
					"attributes":{
						"name":"Test Glossary",
						"qualifiedName":"test_glossary",
						"assetIcon":"PhAirplaneInFlight"
					}
				}
			}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(response))
		} else {
			// Respond with a generic success message for other requests
			response := `{"message": "fail"}`
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(response))
		}
	}))
	defer mockServer.Close()
	//logger := log.New(os.Stdout, "Test Logger: ", log.LstdFlags)

	LoggingEnabled = false
	atlanClient, _ := Context("mock_api_key", mockServer.URL)

	// Set the default AtlanClient to the mock client
	DefaultAtlanClient = atlanClient

	// Iterate over test cases
	for _, tc := range testCases {
		// Call the function to get the glossary by GUID
		glossary, err := GetGlossaryByGuid(tc.GUID)
		assert.NoError(t, err, "Expected no error calling GetGlossaryByGuid")

		// Check if the glossary object is not nil
		assert.NotNil(t, glossary, "Expected a glossary object, got nil")

		// Check if the glossary attributes match the expected values
		assert.Equal(t, tc.ExpectedName, glossary.Name, "Expected glossary name to be '%s', got '%s'", tc.ExpectedName, glossary.Name)
		assert.Equal(t, tc.ExpectedQualName, glossary.QualifiedName, "Expected glossary qualified name to be '%s', got '%s'", tc.ExpectedQualName, glossary.QualifiedName)
		assert.Equal(t, tc.ExpectedIcon, glossary.AssetIcon, "Expected glossary icon to be '%s', got '%s'", tc.ExpectedIcon, glossary.Asset.AssetIcon)
	}
}
