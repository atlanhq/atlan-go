package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

type IndexType string

// CustomBool is a custom type that implements the json.Unmarshaler interface
type CustomBool bool

type TypeDef interface {
	GetCategory() atlan.AtlanTypeCategory
}

type TypeDefResponse struct {
	EnumDefs           []EnumDef           `json:"enumDefs"`
	StructDefs         []StructDef         `json:"structDefs"`
	AtlanTagDefs       []AtlanTagDef       `json:"classificationDefs"`
	EntityDefs         []EntityDef         `json:"entityDefs"`
	RelationshipDefs   []RelationshipDef   `json:"relationshipDefs"`
	CustomMetadataDefs []CustomMetadataDef `json:"businessMetadataDefs"`
	//	reservedEntityDefs   []EntityDef
	//	customEntityDefs     []EntityDef
	//	customEntityDefNames map[string]struct{}
}

type TypeDefBase struct {
	Category    atlan.AtlanTypeCategory `json:"category"`
	CreateTime  int64                   `json:"createTime,omitempty"`
	CreatedBy   string                  `json:"createdBy,omitempty"`
	Description string                  `json:"description,omitempty"`
	GUID        string                  `json:"guid,omitempty"`
	Name        string                  `json:"name"`
	TypeVersion string                  `json:"typeVersion,omitempty"`
	UpdateTime  int64                   `json:"updateTime,omitempty"`
	UpdatedBy   string                  `json:"updatedBy,omitempty"`
	Version     int                     `json:"version,omitempty"`
}

// AttributeDef represents the definition of an attribute.
type AttributeDef struct {
	IsNew                 CustomBool                    `json:"isNew,omitempty"`
	Cardinality           *atlan.Cardinality            `json:"cardinality,omitempty"`
	Constraints           *[]map[string]interface{}     `json:"constraints,omitempty"`
	EnumValues            *[]string                     `json:"enumValues,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	DefaultValue          *string                       `json:"defaultValue,omitempty"`
	DisplayName           *string                       `json:"displayName,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	IncludeInNotification CustomBool                    `json:"includeInNotification,omitempty"`
	IndexType             *IndexType                    `json:"indexType,omitempty"`
	IsIndexable           CustomBool                    `json:"isIndexable,omitempty"`
	IsOptional            CustomBool                    `json:"isOptional,omitempty"`
	IsUnique              CustomBool                    `json:"isUnique,omitempty"`
	Options               *AttributeOptions             `json:"options,omitempty"`
	SearchWeight          *float64                      `json:"searchWeight,omitempty"`
	SkipScrubbing         CustomBool                    `json:"skipScrubbing,omitempty"`
	TypeName              *string                       `json:"typeName,omitempty"`
	ValuesMinCount        *float64                      `json:"valuesMinCount,omitempty"`
	ValuesMaxCount        *float64                      `json:"valuesMaxCount,omitempty"`
	IndexTypeESConfig     *map[string]string            `json:"indexTypeESConfig,omitempty"`
	IndexTypeESFields     *map[string]map[string]string `json:"indexTypeESFields,omitempty"`
	IsDefaultValueNull    CustomBool                    `json:"isDefaultValueNull,omitempty"`
}

// AttributeOptions represents options for customizing an attribute.
type AttributeOptions struct {
	CustomMetadataVersion       *string    `json:"customMetadataVersion,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	ApplicableEntityTypes       *string    `json:"applicableEntityTypes,omitempty"`
	CustomApplicableEntityTypes *string    `json:"customApplicableEntityTypes,omitempty"`
	AllowSearch                 *string    `json:"allowSearch,omitempty"`
	MaxStrLength                *string    `json:"maxStrLength,omitempty"`
	AllowFiltering              CustomBool `json:"allowFiltering,omitempty"`
	MultiValueSelect            CustomBool `json:"multiValueSelect,omitempty"`
	ShowInOverview              CustomBool `json:"showInOverview,omitempty"`
	IsDeprecated                CustomBool `json:"isDeprecated,omitempty"`
	IsEnum                      CustomBool `json:"isEnum,omitempty"`
	EnumType                    *string    `json:"enumType,omitempty"`
	CustomType                  *string    `json:"customType,omitempty"`
	HasTimePrecision            CustomBool `json:"hasTimePrecision,omitempty"`
	IsArchived                  CustomBool `json:"isArchived,omitempty"`
	ArchivedAt                  *string    `json:"archivedAt,omitempty"` // Using int64 for timestamp
	ArchivedBy                  *string    `json:"archivedBy,omitempty"`
	IsSoftReference             *string    `json:"isSoftReference,omitempty"`
	IsAppendOnPartialUpdate     *string    `json:"isAppendOnPartialUpdate,omitempty"`
	PrimitiveType               *string    `json:"primitiveType,omitempty"`
	ApplicableConnections       *string    `json:"applicableConnections,omitempty"`
	ApplicableGlossaries        *string    `json:"applicableGlossaries,omitempty"`
	ApplicableAssetTypes        *string    `json:"assetTypesList,omitempty"`
	ApplicableGlossaryTypes     *string    `json:"glossaryTypeList,omitempty"`
	ApplicableOtherAssetTypes   *string    `json:"otherAssetTypeList,omitempty"`
}

func CreateOptions(attributeType atlan.AtlanCustomAttributePrimitiveType, optionsName *string) (*AttributeOptions, error) {
	if attributeType.Name == "" {
		return nil, errors.New("attribute type is required")
	}

	options := &AttributeOptions{
		//CustomMetadataVersion: "v2",
		PrimitiveType:         structs.StringPtr(attributeType.Name),
		ApplicableEntityTypes: structs.StringPtr(`["Asset"]`),
		AllowSearch:           structs.StringPtr("false"),
		MaxStrLength:          structs.StringPtr("100000000"),
		AllowFiltering:        true,
		MultiValueSelect:      false,
		ShowInOverview:        false,
		IsEnum:                false,
	}

	// Handle special cases for certain attribute types
	switch attributeType.Name {
	case atlan.AtlanCustomAttributeTypeUsers.Name, atlan.AtlanCustomAttributeTypeGroups.Name, atlan.AtlanCustomAttributeTypeURL.Name, atlan.AtlanCustomAttributeTypeSQL.Name:
		options.CustomType = structs.StringPtr(attributeType.Name)
	case atlan.AtlanCustomAttributeTypeOptions.Name:
		options.IsEnum = true
		if optionsName != nil {
			options.EnumType = optionsName
		} else {
			return nil, errors.New("enum type name must be provided for options")
		}
	}

	return options, nil
}

// EnumDef struct
type EnumDef struct {
	TypeDefBase
	TypeDef
	Category    atlan.AtlanTypeCategory `json:"category"`
	ElementDefs []ElementDef            `json:"elementDefs"`
	Options     map[string]interface{}  `json:"options,omitempty"`
	ServiceType *string                 `json:"service_type,omitempty"`
	Name        string                  `json:"name"`
}

// ElementDef struct
type ElementDef struct {
	Name        string  `json:"Name"`
	Value       string  `json:"value"`
	Description *string `json:"description,omitempty"`
	Ordinal     *int    `json:"ordinal,omitempty"`
}

// Function to create an ElementDef with ordinal and value
func NewElementDef(ordinal int, value string) (*ElementDef, error) {
	if value == "" {
		return nil, errors.New("value cannot be empty")
	}
	return &ElementDef{
		Value:   value,
		Ordinal: &ordinal,
	}, nil
}

// Function to generate a list of ElementDefs from a list of values
func ListFrom(values []string) ([]ElementDef, error) {
	if len(values) == 0 {
		return nil, errors.New("values cannot be empty")
	}
	var elements []ElementDef
	for i, value := range values {
		element, err := NewElementDef(i, value)
		if err != nil {
			return nil, err
		}
		elements = append(elements, *element)
	}
	return elements, nil
}

// Function to extend the elements without duplications
func ExtendElements(current []string, new []string) []string {
	uniqueElements := make(map[string]bool)
	for _, element := range current {
		uniqueElements[element] = true
	}
	for _, element := range new {
		if !uniqueElements[element] {
			current = append(current, element)
			uniqueElements[element] = true
		}
	}
	return current
}

// Function to create a new EnumDef
func NewEnumDef(name string, values []string) (*EnumDef, error) {
	if name == "" || len(values) == 0 {
		return nil, errors.New("name and values cannot be empty")
	}

	elementDefs, err := ListFrom(values)
	if err != nil {
		return nil, err
	}

	return &EnumDef{
		Category:    atlan.AtlanTypeCategoryEnum,
		ElementDefs: elementDefs,
		Name:        name,
	}, nil
}

// Function to update EnumDef
func UpdateEnumDef(name string, values []string, replaceExisting bool, currentValues []string) (*EnumDef, error) {
	if name == "" || len(values) == 0 {
		return nil, errors.New("name and values cannot be empty")
	}

	var updateValues []string
	if replaceExisting {
		updateValues = values
	} else {
		updateValues = ExtendElements(currentValues, values)
	}

	elementDefs, err := ListFrom(updateValues)
	if err != nil {
		return nil, err
	}

	return &EnumDef{
		Name:        name,
		Category:    atlan.AtlanTypeCategoryEnum,
		ElementDefs: elementDefs,
	}, nil
}

// Function to get valid values from EnumDef
func (e *EnumDef) GetValidValues() []string {
	var values []string
	for _, element := range e.ElementDefs {
		values = append(values, element.Value)
	}
	return values
}

type StructDef struct {
	TypeDef
	TypeDefBase
	AttributeDefs []*AttributeDef `json:"attributeDefs"`
	ServiceType   *string         `json:"serviceType,omitempty"` // Optional service type
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
	name                  string     `json:"name"`
	typename              string     `json:"typeName"`
	isOptional            CustomBool `json:"isOptional"`
	cardinality           string     `json:"cardinality"`
	valuesmincount        int        `json:"valuesMinCount"`
	valuesmaxcount        int        `json:"valuesMaxCount"`
	isUnique              CustomBool `json:"isUnique"`
	isIndexable           CustomBool `json:"isIndexable"`
	includeInNotification CustomBool `json:"includeInNotification"`
	skipScrubbing         CustomBool `json:"skipScrubbing"`
	searchWeight          int        `json:"searchWeight"`
	displayName           string     `json:"displayName"`
	isDefaultValueNull    CustomBool `json:"isDefaultValueNull"`
}

func (a *AtlanTagDef) GetCategory() atlan.AtlanTypeCategory {
	return a.Category
}

type EntityDef struct {
	TypeDef
	TypeDefBase
	AttributeDefs             []map[string]interface{} `json:"attributeDefs"`
	BusinessAttributeDefs     map[string][]map[string]interface{}
	RelationshipAttributeDefs []map[string]interface{}
	ServiceType               *string   `json:"serviceType"`
	SubTypes                  []*string `json:"subTypes"`
	SuperTypes                []*string `json:"superTypes"`
}

type RelationshipDef struct {
	TypeDef
	TypeDefBase
	AttributeDefs        []map[string]interface{} `json:"attributeDefs"`
	EndDef1              map[string]interface{}   `json:"endDef1"`
	EndDef2              map[string]interface{}   `json:"endDef2"`
	PropagateTags        *string                  `json:"propagateTags"`
	RelationshipCategory *string                  `json:"relationshipCategory"`
	RelationshipLabel    *string                  `json:"relationshipLabel"`
	ServiceType          *string                  `json:"serviceType"`
}

func (a *CustomMetadataDef) GetCategory() atlan.AtlanTypeCategory {
	return *a.Category
}

// CustomMetadataDefOptions represents options for customizing metadata definitions.
type CustomMetadataDefOptions struct {
	Emoji     *string              `json:"emoji,omitempty"`
	ImageID   *string              `json:"imageId,omitempty"`
	IsLocked  string               `json:"isLocked,omitempty"`
	LogoType  *string              `json:"logoType,omitempty"`
	LogoURL   *string              `json:"logoUrl,omitempty"`
	IconColor *atlan.AtlanTagColor `json:"iconColor,omitempty"`
	IconName  *atlan.AtlanIcon     `json:"iconName,omitempty"`
}

// CustomMetadataDef represents the definition of custom metadata.
type CustomMetadataDef struct {
	TypeDefBase
	TypeDef
	AttributeDefs []AttributeDef            `json:"attributeDefs"`
	Category      *atlan.AtlanTypeCategory  `json:"category,omitempty"`
	DisplayName   *string                   `json:"displayName,omitempty"`
	Options       *CustomMetadataDefOptions `json:"options,omitempty"`
}

// Methods

// UnmarshalJSON implements the json.Unmarshaler interface
func (cb *CustomBool) UnmarshalJSON(data []byte) error {
	var b bool
	if err := json.Unmarshal(data, &b); err == nil {
		*cb = CustomBool(b)
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "true":
		*cb = CustomBool(true)
	case "false":
		*cb = CustomBool(false)
	default:
		return fmt.Errorf("invalid boolean value: %s", s)
	}

	return nil
}
