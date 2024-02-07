package client

import "atlan-go/atlan/model"

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

type AtlasGlossary struct {
	Asset
	Entities []model.Glossary `json:"entities"`
}

type AtlasTable struct {
	Attributes
}

func NewTable() *AtlasTable {
	return &AtlasTable{
		Attributes: Attributes{
			TypeName: NewKeywordTextField("typeName", "__typeName.keyword", "__typeName"),
		},
	}
}
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

/*
	glossary := NewGlossary()
    query := glossary.TypeName.StartsWith("H", nil)
	searchResult, err := client.NewFluentSearch().
		PageSizes(10).
where(query).
*/
