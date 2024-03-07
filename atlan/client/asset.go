package client

import (
	"atlan-go/atlan/model"
	Assets2 "atlan-go/atlan/model/assets"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type AtlanObject interface {
	MarshalJSON() ([]byte, error)
}

// SearchAssets Struct to represent assets for searching
type SearchAssets struct {
	Glossary *AtlasGlossary
	Table    *AtlasTable
	// Add other assets here
}

type AttributesFields struct {
	TYPENAME              *model.KeywordTextField
	GUID                  *model.KeywordField
	CREATED_BY            *model.KeywordField
	UPDATED_BY            *model.KeywordField
	STATUS                *model.KeywordField
	ATLAN_TAGS            *model.KeywordTextField
	PROPOGATED_ATLAN_TAGS *model.KeywordTextField
	ASSIGNED_TERMS        *model.KeywordTextField
	SUPERTYPE_NAMES       *model.KeywordTextField
	CREATE_TIME           *model.NumericField
	UPDATE_TIME           *model.NumericField
	QUALIFIED_NAME        *model.KeywordTextField
}

type AssetFields struct {
	*AttributesFields
	NAME                       *model.KeywordTextStemmedField
	DISPLAY_NAME               *model.KeywordTextField
	DESCRIPTION                *model.KeywordTextField
	USER_DESCRIPTION           *model.KeywordTextField
	TENET_ID                   *model.KeywordField
	CERTIFICATE_STATUS         *model.KeywordTextField
	CERTIFICATE_STATUS_MESSAGE *model.KeywordField
	CERTIFICATE_UPDATED_BY     *model.NumericField
	ANNOUNCEMENT_TITLE         *model.KeywordField
	ANNOUNCEMENT_MESSAGE       *model.KeywordTextField
	ANNOUNCEMENT_TYPE          *model.KeywordField
	ANNOUNCEMENT_UPDATED_AT    *model.NumericField
	ANNOUNCEMENT_UPDATED_BY    *model.KeywordField
	OWNER_USERS                *model.KeywordTextField
	ADMIN_USERS                *model.KeywordField
	VIEWER_USERS               *model.KeywordField
	VIEWER_GROUPS              *model.KeywordField
	CONNECTOR_NAME             *model.KeywordTextField
	CONNECTION_NAME            *model.KeywordTextField
}

// AtlasGlossary represents the AtlasGlossary asset
type AtlasGlossary struct {
	AssetFields
	AtlanObject
	Entities []Assets2.Glossary `json:"entities"`
}

type AtlasTable struct {
	AttributesFields
}

// NewSearchTable returns a new AtlasTable object for Searching
func NewSearchTable() *AtlasTable {
	return &AtlasTable{
		AttributesFields: AttributesFields{
			TYPENAME: model.NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
		},
	}
}

// NewSearchGlossary returns a new AtlasGlossary object for Searching
func NewSearchGlossary() *AtlasGlossary {
	return &AtlasGlossary{
		AssetFields: AssetFields{
			AttributesFields: &AttributesFields{
				TYPENAME:              model.NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
				GUID:                  model.NewKeywordField("guid", "__guid"),
				CREATED_BY:            model.NewKeywordField("createdBy", "__createdBy"),
				UPDATED_BY:            model.NewKeywordField("updatedBy", "__modifiedBy"),
				STATUS:                model.NewKeywordField("status", "__state"),
				ATLAN_TAGS:            model.NewKeywordTextField("classificationNames", "__traitNames", "__classificationsText"),
				PROPOGATED_ATLAN_TAGS: model.NewKeywordTextField("classificationNames", "__propagatedTraitNames", "__classificationsText"),
				ASSIGNED_TERMS:        model.NewKeywordTextField("meanings", "__meanings", "__meaningsText"),
				SUPERTYPE_NAMES:       model.NewKeywordTextField("typeName", "__superTypeNames.keyword", "__superTypeNames"),
				CREATE_TIME:           model.NewNumericField("createTime", "__timestamp"),
				UPDATE_TIME:           model.NewNumericField("updateTime", "__modificationTimestamp"),
				QUALIFIED_NAME:        model.NewKeywordTextField("qualifiedName", "qualifiedName", "qualifiedName.text"),
			},
			DISPLAY_NAME: model.NewKeywordTextField("displayName", "displayName.keyword", "displayName"),
			NAME:         model.NewKeywordTextStemmedField("name", "name.keyword", "name", "name"),
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

// WaitTillDeleted waits for an asset to be deleted.
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

// Save saves the asset in memory to the Atlas server.
func Save(asset AtlanObject) (*model.AssetMutationResponse, error) {
	assetJSON, err := asset.MarshalJSON()
	if err != nil {
		return nil, err
	}

	fmt.Println(string(assetJSON))

	var requestObj interface{}
	err = json.Unmarshal(assetJSON, &requestObj)
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
