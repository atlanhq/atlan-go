package model

type AtlanTypeCategory string
type AtlanTagColor string

// Constants representing tag colors
const (
	AtlanTagColorGreen  AtlanTagColor = "Green"
	AtlanTagColorYellow AtlanTagColor = "Yellow"
	AtlanTagColorRed    AtlanTagColor = "Red"
	AtlanTagColorGray   AtlanTagColor = "Gray"
)

// Constants representing type categories
const (
	AtlanTypeCategoryEnum           AtlanTypeCategory = "ENUM"
	AtlanTypeCategoryStruct         AtlanTypeCategory = "STRUCT"
	AtlanTypeCategoryClassification AtlanTypeCategory = "CLASSIFICATION"
	AtlanTypeCategoryEntity         AtlanTypeCategory = "ENTITY"
	AtlanTypeCategoryRelationship   AtlanTypeCategory = "RELATIONSHIP"
	AtlanTypeCategoryCustomMetadata AtlanTypeCategory = "BUSINESS_METADATA"
)

type TypeDef interface {
	GetCategory() AtlanTypeCategory
}

func (a *AtlanTagDef) GetCategory() AtlanTypeCategory {
	return a.Category
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
	Options       map[string]interface{}   `json:"options"`
	AttributeDefs []map[string]interface{} `json:"attributeDefs"`
	DisplayName   string                   `json:"displayName"`
	EntityTypes   []string                 `json:"entityTypes"`
	SubTypes      []string                 `json:"subTypes"`
	SuperTypes    []string                 `json:"superTypes"`
}

type TypeDefResponse struct {
	EnumDefs           []EnumDef           `json:"enumDefs"`
	StructDefs         []StructDef         `json:"structDefs"`
	AtlanTagDefs       []AtlanTagDef       `json:"classificationDefs"`
	EntityDefs         []EntityDef         `json:"entityDefs"`
	RelationshipDefs   []RelationshipDef   `json:"relationshipDefs"`
	CustomMetadataDefs []CustomMetadataDef `json:"businessMetadataDefs"`
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

type CustomMetadataDef struct {
	TypeDef
}
