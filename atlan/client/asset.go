package client

import (
	"atlan-go/atlan/model"
	Assets2 "atlan-go/atlan/model/assets"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
	"time"
)

type AtlanObject interface {
	MarshalJSON() ([]byte, error)
}

// SearchAssets Struct to represent assets for searching
type SearchAssets struct {
	Glossary *AtlasGlossaryFields
	Table    *AtlasTableFields
	// Add other assets here
}

type AttributesFields struct {
	TYPENAME              *KeywordTextField
	GUID                  *KeywordField
	CREATED_BY            *KeywordField
	UPDATED_BY            *KeywordField
	STATUS                *KeywordField
	ATLAN_TAGS            *KeywordTextField
	PROPOGATED_ATLAN_TAGS *KeywordTextField
	ASSIGNED_TERMS        *KeywordTextField
	SUPERTYPE_NAMES       *KeywordTextField
	CREATE_TIME           *NumericField
	UPDATE_TIME           *NumericField
	QUALIFIED_NAME        *KeywordTextField
}

type AssetFields struct {
	*AttributesFields
	NAME                       *KeywordTextStemmedField
	DISPLAY_NAME               *KeywordTextField
	DESCRIPTION                *KeywordTextField
	USER_DESCRIPTION           *KeywordTextField
	TENET_ID                   *KeywordField
	CERTIFICATE_STATUS         *KeywordTextField
	CERTIFICATE_STATUS_MESSAGE *KeywordField
	CERTIFICATE_UPDATED_BY     *NumericField
	ANNOUNCEMENT_TITLE         *KeywordField
	ANNOUNCEMENT_MESSAGE       *KeywordTextField
	ANNOUNCEMENT_TYPE          *KeywordField
	ANNOUNCEMENT_UPDATED_AT    *NumericField
	ANNOUNCEMENT_UPDATED_BY    *KeywordField
	OWNER_USERS                *KeywordTextField
	ADMIN_USERS                *KeywordField
	VIEWER_USERS               *KeywordField
	VIEWER_GROUPS              *KeywordField
	CONNECTOR_NAME             *KeywordTextField
	CONNECTION_NAME            *KeywordTextField
}

// AtlasGlossary represents the AtlasGlossary asset
type AtlasGlossaryFields struct {
	AssetFields
	AtlanObject
}

type AtlasTableFields struct {
	AttributesFields
}

// NewSearchTable returns a new AtlasTable object for Searching
func NewSearchTable() *AtlasTableFields {
	return &AtlasTableFields{
		AttributesFields: AttributesFields{
			TYPENAME: NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
		},
	}
}

// NewSearchGlossary returns a new AtlasGlossary object for Searching
func NewSearchGlossary() *AtlasGlossaryFields {
	return &AtlasGlossaryFields{
		AssetFields: AssetFields{
			AttributesFields: &AttributesFields{
				TYPENAME:              NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
				GUID:                  NewKeywordField("guid", "__guid"),
				CREATED_BY:            NewKeywordField("createdBy", "__createdBy"),
				UPDATED_BY:            NewKeywordField("updatedBy", "__modifiedBy"),
				STATUS:                NewKeywordField("status", "__state"),
				ATLAN_TAGS:            NewKeywordTextField("classificationNames", "__traitNames", "__classificationsText"),
				PROPOGATED_ATLAN_TAGS: NewKeywordTextField("classificationNames", "__propagatedTraitNames", "__classificationsText"),
				ASSIGNED_TERMS:        NewKeywordTextField("meanings", "__meanings", "__meaningsText"),
				SUPERTYPE_NAMES:       NewKeywordTextField("typeName", "__superTypeNames.keyword", "__superTypeNames"),
				CREATE_TIME:           NewNumericField("createTime", "__timestamp"),
				UPDATE_TIME:           NewNumericField("updateTime", "__modificationTimestamp"),
				QUALIFIED_NAME:        NewKeywordTextField("qualifiedName", "qualifiedName", "qualifiedName.text"),
			},
			DISPLAY_NAME: NewKeywordTextField("displayName", "displayName.keyword", "displayName"),
			NAME:         NewKeywordTextStemmedField("name", "name.keyword", "name", "name"),
		},
	}
}

// Methods on assets

// RetrieveMinimal retrieves an asset by its GUID, without any of its relationships.
func RetrieveMinimal(guid string) (*Assets2.Asset, error) {
	if DefaultAtlanClient == nil {
		return nil, fmt.Errorf("default AtlanClient not initialized")
	}

	api := &GET_ENTITY_BY_GUID
	originalPath := api.Path

	api.Path += guid

	// Add query parameters to ignore relationships
	queryParams := make(map[string]string)
	queryParams["min_ext_info"] = "true"
	queryParams["ignore_relationships"] = "true"

	response, err := DefaultAtlanClient.CallAPI(api, queryParams, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into asset json structure
	var assetresponse Assets2.Asset
	err = json.Unmarshal(response, &assetresponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling asset response: %v", err)
	}

	api.Path = originalPath // Reset the api.Path to its original value
	return &assetresponse, nil
}

// PurgeByGuid HARD deletes assets by their GUIDs.
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

// DeleteByGuid SOFT deletes assets by their GUIDs.
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
		if *asset.TypeName == "AtlasGlossaryCategory" {
			return nil, fmt.Errorf("asset %s of type %s cannot be archived", guid, *asset.TypeName)
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

// WaitTillDeleted waits for an asset to be deleted.
func WaitTillDeleted(guid string) error {
	for i := 0; i < MaxRetries; i++ {
		asset, err := RetrieveMinimal(guid)
		if err != nil {
			return fmt.Errorf("error retrieving asset: %v", err)
		}

		if *asset.Status == "DELETED" {
			return nil
		}

		// If the asset is not deleted, wait for a while before retrying
		time.Sleep(RetryInterval)
	}

	// If the asset is still not deleted after all retries, return an error
	return errors.New("retry limit overrun waiting for asset to be deleted")
}

type SaveRequest struct {
	Entities []AtlanObject `json:"entities"`
}

// Save saves the assets in memory to the Atlas server.
func Save(assets ...AtlanObject) (*model.AssetMutationResponse, error) {
	request := SaveRequest{
		Entities: assets,
	}

	api := &CREATE_ENTITIES
	resp, err := DefaultAtlanClient.CallAPI(api, nil, request)
	if err != nil {
		return nil, err
	}

	var response model.AssetMutationResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func generateCacheKey(baseURL, apiKey string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(fmt.Sprintf("%s/%s", baseURL, apiKey)))
	return fmt.Sprintf("%d", h.Sum32())
}
