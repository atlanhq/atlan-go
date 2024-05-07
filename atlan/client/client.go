package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// LoggerConfig defines the configuration for logging.
type LoggerConfig struct {
	Level   string
	Enabled bool
}

// AtlanClient defines the Atlan API client structure.
type AtlanClient struct {
	Session        *http.Client
	host           string
	ApiKey         string
	loggingEnabled bool
	requestParams  map[string]interface{}
	logger         *log.Logger
	SearchAssets
}

// DefaultAtlanClient represents the default AtlanClient instance.
var (
	LoggingEnabled       = true
	DefaultAtlanClient   *AtlanClient
	DefaultAtlanTagCache *AtlanTagCache
)

// Init initializes the default AtlanClient.
func Init() error {
	apiKey, baseURL := retrieveAPIConfig()

	// Configure client and logger
	client, logger := configureClient()

	// Initialize default AtlanClient
	DefaultAtlanClient = &AtlanClient{
		Session:        client,
		host:           baseURL,
		ApiKey:         apiKey,
		requestParams:  defaultRequestParams(apiKey),
		logger:         logger,
		loggingEnabled: LoggingEnabled,
		SearchAssets:   newDefaultSearchAssets(),
	}

	return nil
}

// Context creates a new AtlanClient with provided API key and base URL.
func Context(apiKey, baseURL string) (*AtlanClient, error) {
	// Configure client and logger
	client, logger := configureClient()

	atlanClient := &AtlanClient{
		Session:        client,
		host:           baseURL,
		ApiKey:         apiKey,
		requestParams:  defaultRequestParams(apiKey),
		logger:         logger,
		loggingEnabled: LoggingEnabled,
		SearchAssets:   newDefaultSearchAssets(),
	}

	// Set as default AtlanClient
	DefaultAtlanClient = atlanClient

	return atlanClient, nil
}

// NewContext initializes a new AtlanClient instance.
func NewContext() *AtlanClient {
	if err := Init(); err != nil {
		panic(fmt.Sprintf("Failed to initialize AtlanClient: %v", err))
	}

	return DefaultAtlanClient
}

// configureClient configures HTTP client and logger.
func configureClient() (*http.Client, *log.Logger) {
	client := &http.Client{}
	logger := log.New(os.Stdout, "AtlanClient: ", log.LstdFlags|log.Lshortfile)

	if !LoggingEnabled {
		logger = log.New(io.Discard, "", 0)
	}

	return client, logger
}

// defaultRequestParams returns default request parameters.
func defaultRequestParams(apiKey string) map[string]interface{} {
	VERSION := "0.0"
	headers := map[string]string{
		"x-atlan-agent":    "sdk",
		"x-atlan-agent-id": "go",
		"User-Agent":       fmt.Sprintf("Atlan-GOSDK/%s", VERSION),
	}

	headers["Authorization"] = "Bearer " + apiKey
	headers["Accept"] = "application/json"
	headers["Content-type"] = "application/json"

	return map[string]interface{}{
		"headers": headers,
	}
}

// newDefaultSearchAssets initializes default SearchAssets for AtlanClient.
func newDefaultSearchAssets() SearchAssets {
	return SearchAssets{
		Glossary:         NewSearchGlossary(),
		Table:            NewSearchTable(),
		Column:           NewSearchColumn(),
		Connection:       NewSearchConnection(),
		MaterialisedView: NewSearchMaterialisedView(),
		View:             NewSearchView(),
	}
}

// retrieveAPIConfig retrieves API configuration from environment variables.
func retrieveAPIConfig() (apiKey, baseURL string) {
	apiKey = os.Getenv("ATLAN_API_KEY")
	if apiKey == "" {
		panic("ATLAN_API_KEY not provided in environmental variables")
	}

	baseURL = os.Getenv("ATLAN_BASE_URL")
	if baseURL == "" {
		panic("ATLAN_BASE_URL not provided in environmental variables")
	}

	return apiKey, baseURL
}

// CallAPI makes a generic API call.
func (ac *AtlanClient) CallAPI(api *API, queryParams map[string]string, requestObj interface{}) ([]byte, error) {
	params := deepCopy(ac.requestParams)
	path := ac.host + api.Endpoint.Atlas + api.Path

	if queryParams != nil {
		params["params"] = queryParams
	}

	if requestObj != nil {
		//fmt.Println("Request Object:", requestObj)
		requestJSON, err := json.Marshal(requestObj)
		//fmt.Println("Request JSON:", string(requestJSON))
		if err != nil {
			return nil, fmt.Errorf("error marshaling request object: %v", err)
		}
		params["data"] = bytes.NewBuffer(requestJSON)
	}

	ac.logAPICall(api.Method, path)

	response, err := ac.makeRequest(api.Method, path, params)
	if err != nil {
		fmt.Println(err)
		errorMessage, _ := ioutil.ReadAll(response.Body)
		return nil, handleApiError(response, string(errorMessage))
	}

	ac.logHTTPStatus(response)

	responseJSON, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != api.Status {
		return nil, fmt.Errorf("unexpected HTTP status: %s", response.Status)
	}

	ac.logResponse(responseJSON)

	return responseJSON, nil
}

// makeRequest makes an HTTP request.
func (ac *AtlanClient) makeRequest(method, path string, params map[string]interface{}) (*http.Response, error) {
	var req *http.Request
	var err error
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, path, nil)
	case http.MethodPost, http.MethodPut:
		body, ok := params["data"].(io.Reader)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'data' parameter for POST/PUT/DELETE request")
		}
		req, err = http.NewRequest(method, path, body)
		req.Header.Set("Content-Type", "application/json")
	case http.MethodDelete:
		// DELETE requests may not always have a body.
		var body io.Reader
		if data, ok := params["data"]; ok {
			body, ok = data.(io.Reader)
			if !ok {
				return nil, fmt.Errorf("invalid 'data' parameter for DELETE request")
			}
		}
		req, err = http.NewRequest(method, path, body)

		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	// Set request headers
	headers, ok := params["headers"].(map[string]string)
	if ok {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// Set query parameters
	queryParams, ok := params["params"].(map[string]string)
	if ok {
		// This implementation can be improved, not doing so since requires significant changes to codebase
		var query string
		for key, value := range queryParams {
			// Check if the key is "guid" and value contains commas
			if key == "guid" && strings.Contains(value, ",") {
				// Split the value by commas
				guids := strings.Split(value, ",")
				for _, guid := range guids {
					// Append each guid to the query string with &guid=
					query += "&guid=" + guid
				}
			} else {
				// For other keys, add them normally
				query += "&" + key + "=" + value
			}
		}
		// Remove the leading "&" from the query string
		if len(query) > 0 {
			query = query[1:]
		}
		req.URL.RawQuery = query
	}

	return ac.Session.Do(req)
}

func (ac *AtlanClient) logAPICall(method, path string) {
	if ac.loggingEnabled {
		ac.logger.Println("------------------------------------------------------")
		ac.logger.Printf("Call         : %s %s\n", method, path)
		ac.logger.Printf("Content-type : application/json\n")
		ac.logger.Printf("Accept       : application/json\n")
	}
}

func (ac *AtlanClient) logHTTPStatus(response *http.Response) {
	if ac.loggingEnabled {
		ac.logger.Printf("HTTP Status: %s\n", response.Status)
		if response.StatusCode < 200 || response.StatusCode >= 300 {
			// Read the response body for the error message
			errorMessage, err := ioutil.ReadAll(response.Body)
			if err != nil {
				ac.logger.Printf("Error reading response body: %v\n", err)
			}
			ac.logger.Printf("Error: %s\n", handleApiError(response, string(errorMessage)))
		}
	}
}

func (ac *AtlanClient) logResponse(responseJSON []byte) {
	if ac.loggingEnabled {
		ac.logger.Println("<== __call_api", string(responseJSON))
	}
}

func deepCopy(original map[string]interface{}) map[string]interface{} {
	dcopy := make(map[string]interface{})
	for key, value := range original {
		dcopy[key] = value
	}
	return dcopy
}
