package model

type AtlanTypeCategory string
type AtlanTagColor string
type IndexType string
type Cardinality string
type AtlanIcon string

type TypeDef interface {
	GetCategory() AtlanTypeCategory
}

type EnumDef struct {
	TypeDef
}

type StructDef struct {
	TypeDef
}

type EntityDef struct {
	TypeDef
}

type RelationshipDef struct {
	TypeDef
}

func (a *AtlanTagDef) GetCategory() AtlanTypeCategory {
	return a.Category
}

func (a *CustomMetadataDef) GetCategory() AtlanTypeCategory {
	return *a.Category
}

type TypeDefResponse struct {
	EnumDefs           []EnumDef           `json:"enumDefs"`
	StructDefs         []StructDef         `json:"structDefs"`
	AtlanTagDefs       []AtlanTagDef       `json:"classificationDefs"`
	EntityDefs         []EntityDef         `json:"entityDefs"`
	RelationshipDefs   []RelationshipDef   `json:"relationshipDefs"`
	CustomMetadataDefs []CustomMetadataDef `json:"businessMetadataDefs"`
}

type TypeDefBase struct {
	Category    AtlanTypeCategory `json:"category"`
	CreateTime  int64             `json:"createTime,omitempty"`
	CreatedBy   string            `json:"createdBy,omitempty"`
	Description string            `json:"description,omitempty"`
	GUID        string            `json:"guid,omitempty"`
	Name        string            `json:"name"`
	TypeVersion string            `json:"typeVersion,omitempty"`
	UpdateTime  int64             `json:"updateTime,omitempty"`
	UpdatedBy   string            `json:"updatedBy,omitempty"`
	Version     int               `json:"version,omitempty"`
}

// AtlanTagDef represents the AtlanTagDef(Classifications) structure.
type AtlanTagDef struct {
	TypeDefBase
	TypeDef
	Options       map[string]interface{} `json:"options"`
	AttributeDefs []AttributesDefsTags   `json:"attributeDefs"`
	DisplayName   string                 `json:"displayName"`
	EntityTypes   []string               `json:"entityTypes"`
	SubTypes      []string               `json:"subTypes"`
	SuperTypes    []string               `json:"superTypes"`
}

type AttributesDefsTags struct {
	name                  string `json:"name"`
	typename              string `json:"typeName"`
	isOptional            bool   `json:"isOptional"`
	cardinality           string `json:"cardinality"`
	valuesmincount        int    `json:"valuesMinCount"`
	valuesmaxcount        int    `json:"valuesMaxCount"`
	isUnique              bool   `json:"isUnique"`
	isIndexable           bool   `json:"isIndexable"`
	includeInNotification bool   `json:"includeInNotification"`
	skipScrubbing         bool   `json:"skipScrubbing"`
	searchWeight          int    `json:"searchWeight"`
	displayName           string `json:"displayName"`
	isDefaultValueNull    bool   `json:"isDefaultValueNull"`
}

// AttributeOptions represents options for customizing an attribute.
type AttributeOptions struct {
	CustomMetadataVersion       *string `json:"customMetadataVersion,omitempty"`
	Description                 *string `json:"description,omitempty"`
	ApplicableEntityTypes       *string `json:"applicableEntityTypes,omitempty"`
	CustomApplicableEntityTypes *string `json:"customApplicableEntityTypes,omitempty"`
	AllowSearch                 *string `json:"allowSearch,omitempty"`
	MaxStrLength                *string `json:"maxStrLength,omitempty"`
	AllowFiltering              *string `json:"allowFiltering,omitempty"`
	MultiValueSelect            *string `json:"multiValueSelect,omitempty"`
	ShowInOverview              *string `json:"showInOverview,omitempty"`
	IsDeprecated                *string `json:"isDeprecated,omitempty"`
	IsEnum                      *string `json:"isEnum,omitempty"`
	EnumType                    *string `json:"enumType,omitempty"`
	CustomType                  *string `json:"customType,omitempty"`
	HasTimePrecision            *bool   `json:"hasTimePrecision,omitempty"`
	IsArchived                  bool    `json:"isArchived,omitempty"`
	ArchivedAt                  *int64  `json:"archivedAt,omitempty"` // Using int64 for timestamp
	ArchivedBy                  *string `json:"archivedBy,omitempty"`
	IsSoftReference             *string `json:"isSoftReference,omitempty"`
	IsAppendOnPartialUpdate     *string `json:"isAppendOnPartialUpdate,omitempty"`
	PrimitiveType               *string `json:"primitiveType,omitempty"`
	ApplicableConnections       *string `json:"applicableConnections,omitempty"`
	ApplicableGlossaries        *string `json:"applicableGlossaries,omitempty"`
	ApplicableAssetTypes        *string `json:"assetTypesList,omitempty"`
	ApplicableGlossaryTypes     *string `json:"glossaryTypeList,omitempty"`
	ApplicableOtherAssetTypes   *string `json:"otherAssetTypeList,omitempty"`
}

// AttributeDef represents the definition of an attribute.
type AttributeDef struct {
	IsNew                 *bool                         `json:"isNew,omitempty"`
	Cardinality           *Cardinality                  `json:"cardinality,omitempty"`
	Constraints           *[]map[string]interface{}     `json:"constraints,omitempty"`
	EnumValues            *[]string                     `json:"enumValues,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	DefaultValue          *string                       `json:"defaultValue,omitempty"`
	DisplayName           *string                       `json:"displayName,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	IncludeInNotification *bool                         `json:"includeInNotification,omitempty"`
	IndexType             *IndexType                    `json:"indexType,omitempty"`
	IsIndexable           *bool                         `json:"isIndexable,omitempty"`
	IsOptional            *bool                         `json:"isOptional,omitempty"`
	IsUnique              *bool                         `json:"isUnique,omitempty"`
	Options               *AttributeOptions             `json:"options,omitempty"`
	SearchWeight          *float64                      `json:"searchWeight,omitempty"`
	SkipScrubbing         *bool                         `json:"skipScrubbing,omitempty"`
	TypeName              *string                       `json:"typeName,omitempty"`
	ValuesMinCount        *float64                      `json:"valuesMinCount,omitempty"`
	ValuesMaxCount        *float64                      `json:"valuesMaxCount,omitempty"`
	IndexTypeESConfig     *map[string]string            `json:"indexTypeESConfig,omitempty"`
	IndexTypeESFields     *map[string]map[string]string `json:"indexTypeESFields,omitempty"`
	IsDefaultValueNull    *bool                         `json:"isDefaultValueNull,omitempty"`
}

// CustomMetadataDefOptions represents options for customizing metadata definitions.
type CustomMetadataDefOptions struct {
	Emoji     *string        `json:"emoji,omitempty"`
	ImageID   *string        `json:"imageId,omitempty"`
	IsLocked  *bool          `json:"isLocked,omitempty"`
	LogoType  *string        `json:"logoType,omitempty"`
	LogoURL   *string        `json:"logoUrl,omitempty"`
	IconColor *AtlanTagColor `json:"iconColor,omitempty"`
	IconName  *AtlanIcon     `json:"iconName,omitempty"`
}

// CustomMetadataDef represents the definition of custom metadata.
type CustomMetadataDef struct {
	TypeDefBase
	TypeDef
	AttributeDefs []AttributeDef            `json:"attributeDefs"`
	Category      *AtlanTypeCategory        `json:"category,omitempty"`
	DisplayName   *string                   `json:"displayName,omitempty"`
	Options       *CustomMetadataDefOptions `json:"options,omitempty"`
}
