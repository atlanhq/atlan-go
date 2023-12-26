package client

import (
	"atlan-go/atlan/model"
	"encoding/json"
	"errors"
	"fmt"
)

type TypeDefClient struct {
	Client *AtlanClient
}

// NewTypeDefClient creates a new instance of TypeDefClient.
func NewTypeDefClient(client *AtlanClient) *TypeDefClient {
	return &TypeDefClient{Client: client}
}

func buildTypeDefRequest(typedef model.TypeDef) (model.TypeDefResponse, error) {
	var payload model.TypeDefResponse

	switch t := typedef.(type) {
	case *model.AtlanTagDef:
		payload = model.TypeDefResponse{
			AtlanTagDefs:       []model.AtlanTagDef{*t},
			EnumDefs:           []model.EnumDef{},
			StructDefs:         []model.StructDef{},
			EntityDefs:         []model.EntityDef{},
			RelationshipDefs:   []model.RelationshipDef{},
			CustomMetadataDefs: []model.CustomMetadataDef{},
		}
	case *model.CustomMetadataDef:
		payload = model.TypeDefResponse{
			AtlanTagDefs:       []model.AtlanTagDef{},
			EnumDefs:           []model.EnumDef{},
			StructDefs:         []model.StructDef{},
			EntityDefs:         []model.EntityDef{},
			RelationshipDefs:   []model.RelationshipDef{},
			CustomMetadataDefs: []model.CustomMetadataDef{*t},
		}
	case *model.EnumDef:
		payload = model.TypeDefResponse{
			AtlanTagDefs:       []model.AtlanTagDef{},
			EnumDefs:           []model.EnumDef{*t},
			StructDefs:         []model.StructDef{},
			EntityDefs:         []model.EntityDef{},
			RelationshipDefs:   []model.RelationshipDef{},
			CustomMetadataDefs: []model.CustomMetadataDef{},
		}
	default:
		return model.TypeDefResponse{}, errors.New("unsupported typedef category")
	}

	return payload, nil
}

func NewTypeDefResponse(rawJSON []byte) (*model.TypeDefResponse, error) {
	var response model.TypeDefResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
func RefreshCaches(typedef model.TypeDef) error {

	switch t := typedef.(type) {
	case *model.AtlanTagDef:
		atlanTagCache := NewAtlanTagCache(DefaultAtlanClient)
		return atlanTagCache.RefreshCache()
	case model.CustomMetadataDef:
		//return CustomMetadataCache.RefreshCache()
	case model.EnumDef:
		//return EnumCache.RefreshCache()
	default:
		return fmt.Errorf("unsupported typedef category: %T", t)
	}
	return nil
}

func GetAll() (*model.TypeDefResponse, error) {
	rawJSON, err := DefaultAtlanClient.CallAPI(&GET_ALL_TYPE_DEFS, nil, nil)
	if err != nil {
		return nil, err
	}
	return NewTypeDefResponse(rawJSON)
}

// Get retrieves a TypeDefResponse object that contains a list of the specified category type definitions in Atlan.
func Get(typeCategory model.AtlanTypeCategory) (*model.TypeDefResponse, error) {
	queryParams := map[string]string{"type": string(typeCategory)}
	rawJSON, err := DefaultAtlanClient.CallAPI(&GET_ALL_TYPE_DEFS, queryParams, nil)
	if err != nil {
		return nil, err
	}
	return NewTypeDefResponse(rawJSON)
}

func (c *TypeDefClient) Create(typedef model.TypeDef) (*model.TypeDefResponse, error) {
	payload, err := buildTypeDefRequest(typedef)
	rawJSON, err := c.Client.CallAPI(&CREATE_TYPE_DEFS, nil, payload)
	if err != nil {
		return nil, err
	}
	RefreshCaches(typedef)
	return NewTypeDefResponse(rawJSON)
}

func (c *TypeDefClient) Update(typedef model.TypeDef) (*model.TypeDefResponse, error) {
	payload, err := buildTypeDefRequest(typedef)
	rawJSON, err := c.Client.CallAPI(&UPDATE_TYPE_DEFS, nil, payload)
	if err != nil {
		return nil, err
	}
	RefreshCaches(typedef)
	return NewTypeDefResponse(rawJSON)
}

func (c *TypeDefClient) Purge(name string, typedefType model.TypeDef) error {
	var internalName string
	switch t := typedefType.(type) {
	case *model.CustomMetadataDef:
		//internalName = CustomMetadataCache.getIDForName(name)
	case *model.EnumDef:
		//internalName = name
	case *model.AtlanTagDef:
		//internalName := fmt.Sprintf("Name: %s\n:", GetIDForName(name))
	default:
		return fmt.Errorf("unsupported TypeDef type: %T", t)
	}

	if internalName == "" {
		return fmt.Errorf("type definition not found by name: %s", name)
	}

	c.Client.CallAPI(&DELETE_TYPE_DEF_BY_NAME, nil, nil)

	switch t := typedefType.(type) {
	case *model.CustomMetadataDef:
		//CustomMetadataCache.refreshCache()
	case *model.EnumDef:
		//EnumCache.refreshCache()
	case *model.AtlanTagDef:
		RefreshCache()
	default:
		return fmt.Errorf("unsupported TypeDef type: %T", t)
	}

	return nil
}
