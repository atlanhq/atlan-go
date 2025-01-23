package assets

import (
	"encoding/json"
	"errors"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

type AuthPolicy structs.AuthPolicy

func (a *AuthPolicy) Updater(name string, qualifiedName string) error {
	if name == "" || qualifiedName == "" {
		return errors.New("name, qualifiedName are required fields")
	}

	a.Name = &name
	a.QualifiedName = &qualifiedName

	return nil
}

func (a *AuthPolicy) UnmarshalJSON(data []byte) error {
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
	if err := json.Unmarshal(temp.Entity.AttributesJSON, &a); err != nil {
		return err
	}

	// Set the GUID and TypeName.
	a.Guid = &temp.Entity.Guid
	a.TypeName = &temp.Entity.TypeName

	return nil
}

func (a *AuthPolicy) MarshalJSON() ([]byte, error) {
	// Marshal the AccessControl asset into a JSON object.

	// Construct the custom JSON structure
	customJSON := map[string]interface{}{
		"typeName": "AuthPolicy",
		"attributes": map[string]interface{}{
			"name":          a.Name,
			"qualifiedName": a.Name,
			// Add other attributes as necessary.
		},
	}

	attributes := customJSON["attributes"].(map[string]interface{})

	if a.QualifiedName != nil && *a.QualifiedName != "" {
		attributes["qualifiedName"] = *a.QualifiedName
	}

	//	if a.Guid != nil && *a.Guid != "" {
	//		customJSON["guid"] = *a.Guid
	//	}

	if a.DisplayName != nil && *a.DisplayName != "" {
		attributes["displayName"] = *a.DisplayName
	}

	if a.PolicyType != nil {
		attributes["policyType"] = *a.PolicyType
	}

	if a.PolicyCategory != nil && *a.PolicyCategory != "" {
		attributes["policyCategory"] = *a.PolicyCategory
	}

	if a.PolicyResources != nil {
		attributes["policyResources"] = *a.PolicyResources
	}

	if a.PolicyActions != nil {
		attributes["policyActions"] = *a.PolicyActions
	}

	if a.PolicyResourceCategory != nil && *a.PolicyResourceCategory != "" {
		attributes["policyResourceCategory"] = *a.PolicyResourceCategory
	}

	if a.PolicyServiceName != nil && *a.PolicyServiceName != "" {
		attributes["policyServiceName"] = *a.PolicyServiceName
	}

	if a.PolicySubCategory != nil && *a.PolicySubCategory != "" {
		attributes["policySubCategory"] = *a.PolicySubCategory
	}

	if a.ConnectionQualifiedName != nil && *a.ConnectionQualifiedName != "" {
		attributes["connectionQualifiedName"] = *a.ConnectionQualifiedName
	}

	if a.PolicyGroups != nil && len(*a.PolicyGroups) > 0 {
		attributes["policyGroups"] = *a.PolicyGroups
	}

	if a.PolicyUsers != nil && len(*a.PolicyUsers) > 0 {
		attributes["policyUsers"] = *a.PolicyUsers
	}

	// Handle nested AccessControl field
	if a.AccessControl != nil {
		accessControl := map[string]interface{}{}

		if a.AccessControl.Guid != nil && *a.AccessControl.Guid != "" {
			accessControl["guid"] = *a.AccessControl.Guid
		}

		if a.AccessControl.TypeName != nil && *a.AccessControl.TypeName != "" {
			accessControl["typeName"] = *a.AccessControl.TypeName
		}

		attributes["accessControl"] = accessControl
	}

	return json.MarshalIndent(customJSON, "", "  ")
}

func (a *AuthPolicy) ToJSON() ([]byte, error) {
	// Marshal the Persona object into a JSON object.
	return json.Marshal(a)
}

func (a *AuthPolicy) FromJSON(data []byte) error {
	// Unmarshal the JSON data into the Persona object.
	return json.Unmarshal(data, a)
}
