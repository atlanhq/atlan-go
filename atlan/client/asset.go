package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"fmt"
)

// SearchAssets Struct to represent assets for searching
type SearchAssets struct {
	Glossary *AtlasGlossary
	Table    *AtlasTable
	// Add other assets here
}

type AttributesFields struct {
	TYPENAME            *KeywordTextField
	GUID                *KeywordField
	CREATEDBY           *KeywordField
	UPDATEDBY           *KeywordField
	STATUS              *KeywordField
	ATLANTAGS           *KeywordTextField
	PROPOGATEDATLANTAGS *KeywordTextField
	ASSIGNEDTERMS       *KeywordTextField
	SUPERTYPENAMES      *KeywordTextField
	CREATETIME          *NumericField
	UPDATETIME          *NumericField
	QUALIFIEDNAME       *KeywordTextField
}

type AssetFields struct {
	*AttributesFields
	NAME                     *KeywordTextStemmedField
	DISPLAYNAME              *KeywordTextField
	DESCRIPTION              *KeywordTextField
	USERDESCRIPTION          *KeywordTextField
	TENETID                  *KeywordField
	CERTIFICATESTATUS        *KeywordTextField
	CERTIFICATESTATUSMESSAGE *KeywordField
	CERTIFICATEUPDATEDBY     *NumericField
	ANNOUNCEMENTTITLE        *KeywordField
	ANNOUNCEMENTMESSAGE      *KeywordTextField
	ANNOUNCEMENTTYPE         *KeywordField
	ANNOUNCEMENTUPDATEDAT    *NumericField
	ANNOUNCEMENTUPDATEDBY    *KeywordField
	OWNERUSERS               *KeywordTextField
	ADMINUSERS               *KeywordField
	VIEWERUSERS              *KeywordField
	VIEWERGROUPS             *KeywordField
	CONNECTORNAME            *KeywordTextField
	CONNECTIONNAME           *KeywordTextField
}

// AtlasGlossary represents the AtlasGlossary asset
type AtlasGlossary struct {
	AssetFields
	AtlanObject
	Entities []model.Glossary `json:"entities"`
}

type AtlasTable struct {
	AttributesFields
}

// NewSearchTable returns a new AtlasTable object for Searching
func NewSearchTable() *AtlasTable {
	return &AtlasTable{
		AttributesFields: AttributesFields{
			TYPENAME: NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
		},
	}
}

// NewSearchGlossary returns a new AtlasGlossary object for Searching
func NewSearchGlossary() *AtlasGlossary {
	return &AtlasGlossary{
		AssetFields: AssetFields{
			AttributesFields: &AttributesFields{
				TYPENAME:            NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
				GUID:                NewKeywordField("guid", "__guid"),
				CREATEDBY:           NewKeywordField("createdBy", "__createdBy"),
				UPDATEDBY:           NewKeywordField("updatedBy", "__modifiedBy"),
				STATUS:              NewKeywordField("status", "__state"),
				ATLANTAGS:           NewKeywordTextField("classificationNames", "__traitNames", "__classificationsText"),
				PROPOGATEDATLANTAGS: NewKeywordTextField("classificationNames", "__propagatedTraitNames", "__classificationsText"),
				ASSIGNEDTERMS:       NewKeywordTextField("meanings", "__meanings", "__meaningsText"),
				SUPERTYPENAMES:      NewKeywordTextField("typeName", "__superTypeNames.keyword", "__superTypeNames"),
				CREATETIME:          NewNumericField("createTime", "__timestamp"),
				UPDATETIME:          NewNumericField("updateTime", "__modificationTimestamp"),
				QUALIFIEDNAME:       NewKeywordTextField("qualifiedName", "qualifiedName", "qualifiedName.text"),
			},
			DISPLAYNAME: NewKeywordTextField("displayName", "displayName.keyword", "displayName"),
			NAME:        NewKeywordTextStemmedField("name", "name.keyword", "name", "name"),
		},
	}
}

// RetrieveMinimal retrieves an asset by its GUID, without any of its relationships.
func RetrieveMinimal(guid string) (*model.Asset, error) {
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
	var assetresponse model.Asset
	err = json.Unmarshal(response, &assetresponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling asset response: %v", err)
	}

	api.Path = originalPath // Reset the api.Path to its original value
	return &assetresponse, nil
}
