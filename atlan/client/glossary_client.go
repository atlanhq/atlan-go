package client

import (
	"atlan-go/atlan"
	"atlan-go/atlan/model/assets"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	MaxRetries    = 5
	RetryInterval = time.Second * 5
)

type AtlasGlossary assets.AtlasGlossary

// GlossaryClient defines the client for interacting with the model API.
type GlossaryClient struct {
	client *AtlanClient
}

// NewGlossaryClient creates a new GlossaryClient instance.
func NewGlossaryClient(ac *AtlanClient) *GlossaryClient {
	return &GlossaryClient{client: ac}
}

// GetGlossaryByGuid retrieves a glossary by its GUID.
func GetGlossaryByGuid(glossaryGuid string) (*assets.AtlasGlossary, error) {
	if DefaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	api := &GET_ENTITY_BY_GUID
	api.Path += glossaryGuid

	response, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return nil, err
	}

	g, err := assets.FromJSON(response)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// GetGlossaryTermByGuid retrieves a glossary term by its GUID.
func GetGlossaryTermByGuid(glossaryGuid string) (*assets.AtlasGlossaryTerm, error) {
	if DefaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	api := &GET_ENTITY_BY_GUID
	api.Path += glossaryGuid

	response, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return nil, err
	}

	gt, err := assets.FromJSONTerm(response)
	if err != nil {
		return nil, err
	}

	return gt, nil
}

// Creator is used to create a new glossary asset in memory.
func (g *AtlasGlossary) Creator(name string, icon atlan.AtlanIcon) {
	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.QualifiedName = assets.StringPtr("CBtveYe0Avp5iwU8q3M7Y12")
	g.AssetIcon = atlan.AtlanIconPtr(icon)
}

// Updater is used to modify a glossary asset in memory.
func (g *AtlasGlossary) Updater(name string, qualifiedName string, glossary_guid string) error {
	if name == "" || qualifiedName == "" || glossary_guid == "" {
		return errors.New("name, qualified_name, and glossary_guid are required fields")
	}

	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.Guid = assets.StringPtr(glossary_guid)
	g.QualifiedName = assets.StringPtr(qualifiedName)

	return nil
}

// MarshalJSON filters out entities to only include those with non-empty attributes.
func (g *AtlasGlossary) MarshalJSON() ([]byte, error) {
	// Construct the custom JSON structure
	customJSON := map[string]interface{}{
		"typeName": "AtlasGlossary",
		"attributes": map[string]interface{}{
			"name": g.Name,
			// Add other attributes as necessary.
		},
		"relationshipAttributes": make(map[string]interface{}),
	}

	if g.QualifiedName != nil && *g.QualifiedName != "" {
		customJSON["attributes"].(map[string]interface{})["qualifiedName"] = *g.QualifiedName
	}

	if g.Guid != nil && *g.Guid != "" {
		customJSON["guid"] = *g.Guid
	}

	// Marshal the custom JSON
	return json.MarshalIndent(customJSON, "", "  ")
}
