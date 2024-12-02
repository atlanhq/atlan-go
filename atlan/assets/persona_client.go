package assets

import (
	"encoding/json"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

type Persona structs.Persona

// Creator is used to create a new persona asset in memory.
func (p *Persona) Creator(name string) {
	p.TypeName = structs.StringPtr("Persona")
	p.Name = structs.StringPtr(name)
}

/*
	Example Usage :
	persona := &Persona{}
	policy, err := persona.CreateMetadataPolicy(
		"MyPolicy",
		"persona-guid-1234",
		atlan.AuthPolicyTypeMetadata,
		[]atlan.PersonaMetadataAction{
			atlan.PersonaMetadataActionRead,
			atlan.PersonaMetadataActionWrite,
		},
		"connection-guid-5678",
		[]string{"resource1", "resource2"},

	)
*/

// CreateMetadataPolicy creates a new metadata policy for a persona.
func (p *Persona) CreateMetadataPolicy(
	name string,
	personaID string,
	policyType atlan.AuthPolicyType,
	actions []atlan.PersonaMetadataAction,
	connectionQualifiedName string,
	resources []string) (*AuthPolicy, error) {

	// Convert actions to their string values
	var policyActions []string
	for _, action := range actions {
		policyActions = append(policyActions, action.String())
	}

	// Create the policy object
	policy := &AuthPolicy{
		Asset: structs.Asset{
			Name:                    structs.StringPtr(name),
			ConnectionQualifiedName: &connectionQualifiedName,
			Referenceable: structs.Referenceable{
				Guid: &personaID,
			},
		},
		PolicyType:             &policyType,
		PolicyCategory:         structs.StringPtr(atlan.AuthPolicyCategoryPersona.String()),
		PolicyResources:        &resources,
		PolicyActions:          &policyActions,
		PolicyResourceCategory: structs.StringPtr(atlan.AuthPolicyResourceCategoryCustom.String()),
		PolicyServiceName:      structs.StringPtr("atlas"),
		PolicySubCategory:      structs.StringPtr("metadata"),
		AccessControl: &structs.AccessControl{
			Asset: structs.Asset{
				Referenceable: structs.Referenceable{
					Guid:     &personaID,
					TypeName: structs.StringPtr("Persona"),
				},
			},
		},
	}
	return policy, nil
}

// CreateDataPolicy creates a new data policy for a persona.
func (p *Persona) CreateDataPolicy(
	name string,
	personaID string,
	policyType atlan.AuthPolicyType,
	connectionQualifiedName string,
	resources []string) (*AuthPolicy, error) {

	// Add "entity-type:*" to resources
	resources = append(resources, "entity-type:*")

	// Define default policy actions
	policyActions := []string{atlan.DataActionSelect.String()}

	// Create the policy object
	policy := &AuthPolicy{
		Asset: structs.Asset{
			ConnectionQualifiedName: &connectionQualifiedName,
			Name:                    &name,
			Referenceable: structs.Referenceable{
				Guid: &personaID,
			},
		},
		PolicyType:             &policyType,
		PolicyCategory:         structs.StringPtr(atlan.AuthPolicyCategoryPersona.String()),
		PolicyActions:          &policyActions,
		PolicyResources:        &resources,
		PolicyResourceCategory: structs.StringPtr(atlan.AuthPolicyResourceCategoryEntity.String()),
		PolicyServiceName:      structs.StringPtr("heka"),
		PolicySubCategory:      structs.StringPtr("data"),
		AccessControl: &structs.AccessControl{
			Asset: structs.Asset{
				Referenceable: structs.Referenceable{
					Guid: &personaID,
				},
			},
		},
	}

	return policy, nil
}

// CreateGlossaryPolicy creates a new glossary policy for a persona.
func (p *Persona) CreateGlossaryPolicy(
	name string,
	personaID string,
	policyType atlan.AuthPolicyType,
	actions []atlan.PersonaGlossaryAction,
	resources []string) (*AuthPolicy, error) {

	// Convert actions to their string values
	policyActions := make([]string, len(actions))
	for i, action := range actions {
		policyActions[i] = action.String()
	}

	// Create the policy object
	policy := &AuthPolicy{
		Asset: structs.Asset{
			Name: &name,
			Referenceable: structs.Referenceable{
				Guid: &personaID,
			},
		},
		PolicyType:             &policyType,
		PolicyCategory:         structs.StringPtr(atlan.AuthPolicyCategoryPersona.String()),
		PolicyActions:          &policyActions,
		PolicyResources:        &resources,
		PolicyResourceCategory: structs.StringPtr(atlan.AuthPolicyResourceCategoryCustom.String()),
		PolicyServiceName:      structs.StringPtr("atlas"),
		PolicySubCategory:      structs.StringPtr("glossary"),
		AccessControl: &structs.AccessControl{
			Asset: structs.Asset{
				Referenceable: structs.Referenceable{
					Guid: &personaID,
				},
			},
		},
	}

	return policy, nil
}

// CreateDomainPolicy creates a new domain policy for a persona.
func (p *Persona) CreateDomainPolicy(
	name string,
	personaID string,
	actions []atlan.PersonaDomainAction,
	resources []string,
) (*AuthPolicy, error) {

	// Convert actions to their string values
	policyActions := make([]string, len(actions))
	for i, action := range actions {
		policyActions[i] = action.String()
	}

	// Create the policy object
	policy := &AuthPolicy{
		Asset: structs.Asset{
			Name: &name,
			Referenceable: structs.Referenceable{
				Guid: &personaID,
			},
		},
		PolicyType:             &atlan.AuthPolicyTypeAllow,
		PolicyCategory:         structs.StringPtr(atlan.AuthPolicyCategoryPersona.String()),
		PolicyActions:          &policyActions,
		PolicyResources:        &resources,
		PolicyResourceCategory: structs.StringPtr(atlan.AuthPolicyResourceCategoryCustom.String()),
		PolicyServiceName:      structs.StringPtr("atlas"),
		PolicySubCategory:      structs.StringPtr("domain"),
		AccessControl: &structs.AccessControl{
			Asset: structs.Asset{
				Referenceable: structs.Referenceable{
					Guid: &personaID,
				},
			},
		},
	}

	return policy, nil
}

func (p *Persona) Updater(qualifiedName, name string, isEnabled bool) error {
	// Validate required fields
	if qualifiedName == "" || name == "" {
		return fmt.Errorf("missing required fields: qualifiedName and name")
	}
	p.QualifiedName = &qualifiedName
	p.Name = &name
	p.IsAccessControlEnabled = &isEnabled
	return nil
}

func (p *Persona) UnmarshalJSON(data []byte) error {
	// Define a temporary structure with the expected JSON structure.
	var temp struct {
		ReferredEntities map[string]interface{} `json:"referredEntities"`
		Entity           struct {
			TypeName               string            `json:"typeName"`
			AttributesJSON         json.RawMessage   `json:"attributes"`
			Guid                   string            `json:"guid"`
			IsIncomplete           bool              `json:"isIncomplete"`
			Status                 atlan.AtlanStatus `json:"status"`
			CreatedBy              string            `json:"createdBy"`
			UpdatedBy              string            `json:"updatedBy"`
			CreateTime             int64             `json:"createTime"`
			UpdateTime             int64             `json:"updateTime"`
			Version                int               `json:"version"`
			RelationshipAttributes struct {
				SchemaRegistrySubjects []structs.SchemaRegistrySubject `json:"schemaRegistrySubjects"`
				McMonitors             []structs.MCMonitor             `json:"mcMonitors"`
				Terms                  []structs.AtlasGlossaryTerm     `json:"terms"`
				OutputPortDataProducts []string                        `json:"outputPortDataProducts"`
				AtlasGlossary          []structs.AtlasGlossary         `json:"AtlasGlossary"`
				AccessControl          []structs.AccessControl         `json:"AccessControl"`
				Policies               []structs.AuthPolicy            `json:"policies"`
			} `json:"relationshipAttributes"`
		}
	}

	// Unmarshal the JSON data into the temporary structure.
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Unmarshal the attributes JSON into the entity.
	if err := json.Unmarshal(temp.Entity.AttributesJSON, &p); err != nil {
		return err
	}

	// Set the GUID and TypeName.
	p.Guid = &temp.Entity.Guid
	p.TypeName = &temp.Entity.TypeName

	return nil
}

func (p *Persona) MarshalJSON() ([]byte, error) {
	// Marshal the AccessControl asset into a JSON object.

	// Construct the custom JSON structure
	customJSON := map[string]interface{}{
		"typeName": "Persona",
		"attributes": map[string]interface{}{
			"name":          p.Name,
			"qualifiedName": p.Name,
			//		"displayName":            p.Name,
			"isAccessControlEnabled": true,
			// Add other attributes as necessary.
		},
	}

	if p.IsAccessControlEnabled != nil {
		customJSON["attributes"].(map[string]interface{})["isAccessControlEnabled"] = *p.IsAccessControlEnabled
	}

	if p.QualifiedName != nil && *p.QualifiedName != "" {
		customJSON["attributes"].(map[string]interface{})["qualifiedName"] = *p.QualifiedName
	}

	if p.Guid != nil && *p.Guid != "" {
		customJSON["guid"] = *p.Guid
	}

	if p.DisplayName != nil && *p.DisplayName != "" {
		customJSON["attributes"].(map[string]interface{})["displayName"] = *p.DisplayName
	}

	if p.PersonaUsers != nil {
		customJSON["attributes"].(map[string]interface{})["personaUsers"] = p.PersonaUsers
	}

	if p.PersonaGroups != nil {
		customJSON["attributes"].(map[string]interface{})["personaGroups"] = p.PersonaGroups
	}

	return json.MarshalIndent(customJSON, "", "  ")
}

func (p *Persona) ToJSON() ([]byte, error) {
	// Marshal the Persona object into a JSON object.
	return json.Marshal(p)
}

func (p *Persona) FromJSON(data []byte) error {
	// Unmarshal the JSON data into the Persona object.
	return json.Unmarshal(data, p)
}
