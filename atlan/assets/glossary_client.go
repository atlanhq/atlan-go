package assets

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

const (
	MaxRetries    = 5
	RetryInterval = time.Second * 5
)

type AtlasGlossary structs.AtlasGlossary

// Creator is used to create a new glossary asset in memory.
func (g *AtlasGlossary) Creator(name string, icon atlan.AtlanIcon) {
	g.TypeName = structs.StringPtr("AtlasGlossary")
	g.Name = structs.StringPtr(name)
	g.QualifiedName = structs.StringPtr(name)
	g.AssetIcon = atlan.AtlanIconPtr(icon)
}

// Updater is used to modify a glossary asset in memory.
func (g *AtlasGlossary) Updater(name string, qualifiedName string, glossary_guid string) error {
	if name == "" || qualifiedName == "" || glossary_guid == "" {
		return errors.New("name, qualified_name, and glossary_guid are required fields")
	}

	g.TypeName = structs.StringPtr("AtlasGlossary")
	g.Name = structs.StringPtr(name)
	g.Guid = structs.StringPtr(glossary_guid)
	g.QualifiedName = structs.StringPtr(qualifiedName)

	return nil
}

func (ag *AtlasGlossary) UnmarshalJSON(data []byte) error {
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
				Files                  []structs.File                  `json:"files"`
				McIncidents            []structs.MCIncident            `json:"mcIncidents"`
				Links                  []structs.Link                  `json:"links"`
				Categories             []structs.AtlasGlossaryCategory `json:"categories"`
				Metrics                []structs.Metric                `json:"metrics"`
				Readme                 []structs.Readme                `json:"readme"`
				Meanings               []structs.Meaning               `json:"meanings"`
				SodaChecks             []structs.SodaCheck             `json:"sodaChecks"`
			} `json:"relationshipAttributes"`
			Labels []interface{} `json:"labels"`
		} `json:"entity"`
	}

	// Unmarshal the JSON into the temporary structure
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Map the fields from the temporary structure to your AtlasGlossary struct
	ag.TypeName = &temp.Entity.TypeName
	ag.Guid = &temp.Entity.Guid
	ag.IsIncomplete = &temp.Entity.IsIncomplete
	ag.Status = &temp.Entity.Status
	ag.CreatedBy = &temp.Entity.CreatedBy
	ag.UpdatedBy = &temp.Entity.UpdatedBy
	ag.CreateTime = &temp.Entity.CreateTime
	ag.UpdateTime = &temp.Entity.UpdateTime

	var asset structs.AtlasGlossaryAttributes
	if err := json.Unmarshal(temp.Entity.AttributesJSON, &asset); err != nil {
		return err
	}

	// Map Asset fields
	ag.Name = asset.Name
	ag.AssetIcon = asset.AssetIcon
	ag.QualifiedName = asset.QualifiedName
	ag.ShortDescription = asset.ShortDescription
	ag.LongDescription = asset.LongDescription
	ag.Language = asset.Language
	ag.Usage = asset.Usage
	ag.AssetIcon = asset.AssetIcon

	return nil
}

// MarshalJSON filters out entities to only include those with non-empty attributes.
func (g *AtlasGlossary) MarshalJSON() ([]byte, error) {
	// Construct the custom JSON structure
	customJSON := map[string]interface{}{
		"typeName": "AtlasGlossary",
		"attributes": map[string]interface{}{
			"name": g.Name,
			// Add other attributes as necessary.
		},
		"relationshipAttributes": make(map[string]interface{}),
	}

	if g.QualifiedName != nil && *g.QualifiedName != "" {
		customJSON["attributes"].(map[string]interface{})["qualifiedName"] = *g.QualifiedName
	}

	if g.Guid != nil && *g.Guid != "" {
		customJSON["guid"] = *g.Guid
	}

	if g.DisplayName != nil && *g.DisplayName != "" {
		customJSON["attributes"].(map[string]interface{})["displayName"] = *g.DisplayName
	}
	if g.Description != nil && *g.Description != "" {
		customJSON["attributes"].(map[string]interface{})["description"] = *g.Description
	}

	// Marshal the custom JSON
	return json.MarshalIndent(customJSON, "", "  ")
}

func (ag *AtlasGlossary) ToJSON() ([]byte, error) {
	return json.MarshalIndent(ag, "", "  ")
}

func (ag *AtlasGlossary) FromJSON(data []byte) error {
	return json.Unmarshal(data, ag)
}
