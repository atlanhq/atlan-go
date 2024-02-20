package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"fmt"
)

// Struct to represent assets for searching
type SearchAssets struct {
	Glossary *AtlasGlossary
	Table    *AtlasTable
	// Add other assets here
}

type Attributes struct {
	TypeName            *KeywordTextField
	GUID                *KeywordField
	CreatedBy           *KeywordField
	UpdatedBy           *KeywordField
	Status              *KeywordField
	AtlanTags           *KeywordTextField
	PropagatedAtlanTags *KeywordTextField
	AssignedTerms       *KeywordTextField
	SuperTypeNames      *KeywordTextField
	CreateTime          *NumericField
	UpdateTime          *NumericField
	QualifiedName       *KeywordTextField
}

type Asset struct {
	*Attributes
	Name                     *KeywordTextStemmedField
	DislayName               *KeywordTextField
	Description              *KeywordTextField
	UserDescription          *KeywordTextField
	TenetID                  *KeywordField
	CetificateStatus         *KeywordTextField
	CertificateStatusMessage *KeywordField
	CertificateUpdatedBy     *NumericField
	AnnouncementTitle        *KeywordField
	AnnouncementMessage      *KeywordTextField
	AnnouncementType         *KeywordField
	AnnouncementUpdatedAt    *NumericField
	AnnouncementUpdatedBy    *KeywordField
	OwnerUsers               *KeywordTextField
	AdminUsers               *KeywordField
	ViewerUsers              *KeywordField
	ViewerGroups             *KeywordField
	ConnectorName            *KeywordTextField
	ConnectionName           *KeywordTextField
}

// AtlasGlossary represents the AtlasGlossary asset
type AtlasGlossary struct {
	Asset
	AtlanObject
	Entities []model.Glossary `json:"entities"`
}

type AtlasTable struct {
	Attributes
}

// NewSearchTable returns a new AtlasTable object for Searching
func NewSearchTable() *AtlasTable {
	return &AtlasTable{
		Attributes: Attributes{
			TypeName: NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
		},
	}
}

// NewSearchGlossary returns a new AtlasGlossary object for Searching
func NewSearchGlossary() *AtlasGlossary {
	return &AtlasGlossary{
		Asset: Asset{
			Attributes: &Attributes{
				TypeName:            NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
				GUID:                NewKeywordField("guid", "__guid"),
				CreatedBy:           NewKeywordField("createdBy", "__createdBy"),
				UpdatedBy:           NewKeywordField("updatedBy", "__modifiedBy"),
				Status:              NewKeywordField("status", "__state"),
				AtlanTags:           NewKeywordTextField("classificationNames", "__traitNames", "__classificationsText"),
				PropagatedAtlanTags: NewKeywordTextField("classificationNames", "__propagatedTraitNames", "__classificationsText"),
				AssignedTerms:       NewKeywordTextField("meanings", "__meanings", "__meaningsText"),
				SuperTypeNames:      NewKeywordTextField("typeName", "__superTypeNames.keyword", "__superTypeNames"),
				CreateTime:          NewNumericField("createTime", "__timestamp"),
				UpdateTime:          NewNumericField("updateTime", "__modificationTimestamp"),
				QualifiedName:       NewKeywordTextField("qualifiedName", "qualifiedName", "qualifiedName.text"),
			},
			DislayName: NewKeywordTextField("displayName", "displayName.keyword", "displayName"),
			Name:       NewKeywordTextStemmedField("name", "name.keyword", "name", "name"),
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
