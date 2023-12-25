package client

import (
	"atlan-go/atlan/model"
	"fmt"
)

// GlossaryClient defines the client for interacting with the model API.
type GlossaryClient struct {
	client *AtlanClient
}

// NewGlossaryClient creates a new GlossaryClient instance.
func NewGlossaryClient(ac *AtlanClient) *GlossaryClient {
	return &GlossaryClient{client: ac}
}

func GetGlossaryByGuid(glossaryGuid string) (*model.Glossary, error) {
	if DefaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	api := &GET_ENTITY_BY_GUID
	api.Path += glossaryGuid

	response, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return nil, err
	}

	g, err := model.FromJSON(response)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func GetGlossaryTermByGuid(glossaryGuid string) (*model.GlossaryTerm, error) {
	if DefaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	api := &GET_ENTITY_BY_GUID
	api.Path += glossaryGuid

	response, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return nil, err
	}

	gt, err := model.FromJSONTerm(response)
	if err != nil {
		return nil, err
	}

	return gt, nil
}
