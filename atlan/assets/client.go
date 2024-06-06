package assets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/logger"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// AtlanClient defines the Atlan API client structure.
type AtlanClient struct {
	Session       *http.Client
	host          string
	ApiKey        string
	requestParams map[string]interface{}
	logger        logger.Logger
	SearchAssets
}

// DefaultAtlanClient represents the default AtlanClient instance.
var (
	DefaultAtlanClient   *AtlanClient
	DefaultAtlanTagCache *AtlanTagCache
)

// Init initializes the default AtlanClient.
func Init() error {
	apiKey, baseURL := retrieveAPIConfig()

	// Normalize the baseURL
	baseURL = normalizeURL(baseURL)

	// Configure client and logger
	client, logger := configureClient()

	// Initialize default AtlanClient
	DefaultAtlanClient = &AtlanClient{
		Session:       client,
		host:          baseURL,
		ApiKey:        apiKey,
		requestParams: defaultRequestParams(apiKey),
		logger:        *logger,
		SearchAssets:  newDefaultSearchAssets(),
	}

	return nil
}

// Context creates a new AtlanClient with provided API key and base URL.
func Context(baseURL, apiKey string) (*AtlanClient, error) {
	// Normalize the baseURL
	baseURL = normalizeURL(baseURL)

	// Configure client and logger
	client, logger := configureClient()

	atlanClient := &AtlanClient{
		Session:       client,
		host:          baseURL,
		ApiKey:        apiKey,
		requestParams: defaultRequestParams(apiKey),
		logger:        *logger,
		SearchAssets:  newDefaultSearchAssets(),
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
func configureClient() (*http.Client, *logger.Logger) {
	client := &http.Client{}

	var loggerInstance *logger.Logger

	// Check if the logger is already set by the user
	if DefaultAtlanClient != nil && DefaultAtlanClient.logger.Log != nil {
		loggerInstance = &DefaultAtlanClient.logger
	} else {
		// Configure logger with default values
		loggerCfg := &logger.LoggerConfig{Level: "info", Enabled: true}
		newLogger := logger.NewLogger(loggerCfg)
		loggerInstance = &newLogger
	}

	return client, loggerInstance
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
		logger.Log.Error("ATLAN_API_KEY not provided in environmental variables")
		panic("ATLAN_API_KEY not provided in environmental variables")
	}

	baseURL = os.Getenv("ATLAN_BASE_URL")
	if baseURL == "" {
		logger.Log.Error("ATLAN_BASE_URL not provided in environmental variables")
		panic("ATLAN_BASE_URL not provided in environmental variables")
	}

	return apiKey, baseURL
}

// normalizeURL ensures the URL starts with "https://" and truncates after the domain.
func normalizeURL(rawURL string) string {
	// Ensure URL starts with "http://" or "https://"
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(fmt.Sprintf("Invalid URL: %v", err))
	}

	// Truncate the URL after the host
	return parsedURL.Scheme + "://" + parsedURL.Host
}

// SetLogger enables or disables logging and sets the log level.
func (ac *AtlanClient) SetLogger(enabled bool, level string) {
	// Create a new logger configuration
	loggerCfg := &logger.LoggerConfig{
		Level:   level,
		Enabled: enabled,
	}
	// Create a new logger instance
	ac.logger = logger.NewLogger(loggerCfg)
}

// EnableLogging enables logging with the specified log level.
func (ac *AtlanClient) EnableLogging(level string) {
	ac.SetLogger(true, level)
}

// DisableLogging disables logging.
func (ac *AtlanClient) DisableLogging() {
	ac.SetLogger(false, "")
}

// CallAPI makes a generic API call.
func (ac *AtlanClient) CallAPI(api *API, queryParams interface{}, requestObj interface{}) ([]byte, error) {
	params := deepCopy(ac.requestParams)
	path := ac.host + api.Endpoint.Atlas + api.Path

	query := url.Values{}
	switch v := queryParams.(type) {
	case map[string]string:
		for key, value := range v {
			query.Add(key, value)
		}
	case map[string][]string:
		for key, values := range v {
			for _, value := range values {
				query.Add(key, value)
			}
		}
	default:
		params["params"] = queryParams
	}

	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	if requestObj != nil {
		//fmt.Println("Request Object:", requestObj)
		requestJSON, err := json.Marshal(requestObj)
		logger.Log.Debugf("Request JSON: %s", string(requestJSON))
		if err != nil {
			ac.logger.Errorf("error marshaling request object: %v", err)
			return nil, fmt.Errorf("error marshaling request object: %v", err)
		}
		params["data"] = bytes.NewBuffer(requestJSON)
	}

	ac.logAPICall(api.Method, path)

	//logger.Log.Debugf("Params: %v", params)
	response, err := ac.makeRequest(api.Method, path, params)
	if err != nil {
		return nil, handleApiError(response, err)
	}

	ac.logHTTPStatus(response)

	responseJSON, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != api.Status {
		return nil, handleApiError(response, err)
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
		if err != nil {
			return nil, ThrowAtlanError(err, CONNECTION_ERROR, nil)
		}
	case http.MethodPost, http.MethodPut:
		body, ok := params["data"].(io.Reader)
		if !ok {
			return nil, fmt.Errorf("missing or invalid 'data' parameter for POST/PUT/DELETE request")
		}
		req, err = http.NewRequest(method, path, body)
		if err != nil {
			return nil, ThrowAtlanError(err, CONNECTION_ERROR, nil)
		}
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
		if err != nil {
			return nil, ThrowAtlanError(err, CONNECTION_ERROR, nil)
		}
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
	ac.logger.Debugf("------------------------------------------------------")
	ac.logger.Debugf("Call         : %s %s", method, path)
	ac.logger.Debugf("Content-type : application/json")
	ac.logger.Debugf("Accept       : application/json")
}

func (ac *AtlanClient) logHTTPStatus(response *http.Response) {
	ac.logger.Debugf("HTTP Status: %s", response.Status)
}

func (ac *AtlanClient) logResponse(responseJSON []byte) {
	ac.logger.Debugf("<== __call_api %s", string(responseJSON))
}

func deepCopy(original map[string]interface{}) map[string]interface{} {
	dcopy := make(map[string]interface{})
	for key, value := range original {
		dcopy[key] = value
	}
	return dcopy
}
