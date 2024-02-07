package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"errors"
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

func (g *AtlasGlossary) Create(name string, icon string) {
	entity := model.Glossary{
		TypeName: "AtlasGlossary",
		Attributes: model.GlossaryAttributes{
			Name:          name,
			QualifiedName: name,
			AssetIcon:     icon,
		},
	}
	if icon != "" {
		entity.Attributes.AssetIcon = icon
	}

	g.Entities = append(g.Entities, entity)
}

func (g *AtlasGlossary) CreateForModification(name string, qualifiedName string, glossary_guid string) error {
	if name == "" || qualifiedName == "" || glossary_guid == "" {
		return errors.New("name, qualified_name, and glossary_guid are required fields")
	}

	entity := model.Glossary{
		TypeName: "AtlasGlossary",
		Attributes: model.GlossaryAttributes{
			Name:          name,
			QualifiedName: qualifiedName,
		},
		Guid: glossary_guid,
	}
	g.Entities = append(g.Entities, entity)
	return nil
}

func (g *AtlasGlossary) MarshalJSON() ([]byte, error) {
	// Filter out entities to only include those with non-empty attributes
	filteredEntities := make([]model.Glossary, 0)
	for _, entity := range g.Entities {
		if entity.Attributes.Name != "" || entity.Attributes.QualifiedName != "" || entity.Attributes.AssetIcon != "" {
			filteredEntities = append(filteredEntities, entity)
		}
	}

	type Alias AtlasGlossary

	customJSON := &struct {
		Entities []model.Glossary `json:"entities"`
	}{
		Entities: filteredEntities,
	}

	return json.MarshalIndent(customJSON, "", "  ")
}

func (g *AtlasGlossary) Save() (*model.AssetMutationResponse, error) {
	glossaryJSON, err := g.MarshalJSON()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(glossaryJSON))

	var requestObj interface{}
	err = json.Unmarshal(glossaryJSON, &requestObj)
	if err != nil {
		return nil, err
	}

	api := &CREATE_ENTITIES
	resp, err := DefaultAtlanClient.CallAPI(api, nil, requestObj)

	fmt.Println("Response:", string(resp)) // Print the response for debugging

	var response model.AssetMutationResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
