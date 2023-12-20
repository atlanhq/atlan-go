package client

import (
	"atlan-go/atlan/model"
	"fmt"
	"net/http"
)

// GlossaryClient defines the client for interacting with the model API.
type GlossaryClient struct {
	client *AtlanClient
}

// NewGlossaryClient creates a new GlossaryClient instance.
func NewGlossaryClient(ac *AtlanClient) *GlossaryClient {
	return &GlossaryClient{client: ac}
}

// GetGlossaryByGuid retrieves a model by its GUID using the default AtlanClient.
func GetGlossaryByGuid(glossaryGuid string) (*model.Glossary, error) {
	if defaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	apiPath := fmt.Sprintf("/api/meta/entity/guid/%s", glossaryGuid)
	response, err := defaultAtlanClient.CallAPI(apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	g, err := model.FromJSON(response)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// GetGlossaryTermByGuid retrieves a model term by its GUID using the default AtlanClient.

func GetGlossaryTermByGuid(glossaryGuid string) (*model.GlossaryTerm, error) {
	if defaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	apiPath := fmt.Sprintf("/api/meta/entity/guid/%s", glossaryGuid)
	response, err := defaultAtlanClient.CallAPI(apiPath, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	gt, err := model.FromJSONTerm(response)
	if err != nil {
		return nil, err
	}

	return gt, nil
}
