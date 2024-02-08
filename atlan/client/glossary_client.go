package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	MaxRetries    = 5
	RetryInterval = time.Second * 5
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

func PurgeByGuid(guids []string) (*model.AssetMutationResponse, error) {
	if len(guids) == 0 {
		return nil, fmt.Errorf("no GUIDs provided for deletion")
	}

	api := &DELETE_ENTITIES_BY_GUIDS

	// Construct the query parameters
	queryParams := make(map[string]string)
	queryParams["deleteType"] = "HARD"

	// Convert the GUIDs slice to a comma-separated string
	guidString := strings.Join(guids, ",")

	// Add the comma-separated string of GUIDs to the query parameters
	queryParams["guid"] = guidString

	// Call the API
	resp, err := DefaultAtlanClient.CallAPI(api, queryParams, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into the AssetMutationResponse struct
	var response model.AssetMutationResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &response, nil
}

func DeleteByGuid(guids []string) (*model.AssetMutationResponse, error) {
	if len(guids) == 0 {
		return nil, fmt.Errorf("no GUIDs provided for deletion")
	}

	for _, guid := range guids {
		asset, err := RetrieveMinimal(guid)
		if err != nil {
			return nil, fmt.Errorf("error retrieving asset: %v", err)
		}

		// Assuming the asset has a CanBeArchived field that indicates if it can be archived
		if asset.TypeName == "AtlasGlossaryCategory" {
			return nil, fmt.Errorf("asset %s of type %s cannot be archived", guid, asset.TypeName)
		}
	}

	api := &DELETE_ENTITIES_BY_GUIDS

	// Construct the query parameters
	queryParams := make(map[string]string)
	queryParams["deleteType"] = "SOFT"

	// Convert the GUIDs slice to a comma-separated string
	guidString := strings.Join(guids, ",")

	// Add the comma-separated string of GUIDs to the query parameters
	queryParams["guid"] = guidString

	fmt.Println("Query Params:", queryParams)
	// Call the API
	resp, err := DefaultAtlanClient.CallAPI(api, queryParams, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into the AssetMutationResponse struct
	var response model.AssetMutationResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	// Wait until each asset is deleted
	for _, guid := range guids {
		err = WaitTillDeleted(guid)
		if err != nil {
			return nil, err
		}
	}

	return &response, nil
}

func WaitTillDeleted(guid string) error {
	for i := 0; i < MaxRetries; i++ {
		asset, err := RetrieveMinimal(guid)
		if err != nil {
			return fmt.Errorf("error retrieving asset: %v", err)
		}

		if asset.Status == "DELETED" {
			return nil
		}

		// If the asset is not deleted, wait for a while before retrying
		time.Sleep(RetryInterval)
	}

	// If the asset is still not deleted after all retries, return an error
	return errors.New("retry limit overrun waiting for asset to be deleted")
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

	var response model.AssetMutationResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
