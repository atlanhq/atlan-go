package client

import (
	"encoding/json"
	"errors"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/assets"
	"time"
)

const (
	MaxRetries    = 5
	RetryInterval = time.Second * 5
)

type AtlasGlossary assets.AtlasGlossary

// GlossaryClient defines the client for interacting with the model API.
type GlossaryClient struct {
	client *AtlanClient
}

// NewGlossaryClient creates a new GlossaryClient instance.
func NewGlossaryClient(ac *AtlanClient) *GlossaryClient {
	return &GlossaryClient{client: ac}
}

// Creator is used to create a new glossary asset in memory.
func (g *AtlasGlossary) Creator(name string, icon atlan.AtlanIcon) {
	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.QualifiedName = assets.StringPtr(name)
	g.AssetIcon = atlan.AtlanIconPtr(icon)
}

// Updater is used to modify a glossary asset in memory.
func (g *AtlasGlossary) Updater(name string, qualifiedName string, glossary_guid string) error {
	if name == "" || qualifiedName == "" || glossary_guid == "" {
		return errors.New("name, qualified_name, and glossary_guid are required fields")
	}

	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.Guid = assets.StringPtr(glossary_guid)
	g.QualifiedName = assets.StringPtr(qualifiedName)

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
				SchemaRegistrySubjects []assets.SchemaRegistrySubject `json:"schemaRegistrySubjects"`
				McMonitors             []assets.MCMonitor             `json:"mcMonitors"`
				Terms                  []assets.AtlasGlossaryTerm     `json:"terms"`
				OutputPortDataProducts []string                       `json:"outputPortDataProducts"`
				Files                  []assets.File                  `json:"files"`
				McIncidents            []assets.MCIncident            `json:"mcIncidents"`
				Links                  []assets.Link                  `json:"links"`
				Categories             []assets.AtlasGlossaryCategory `json:"categories"`
				Metrics                []assets.Metric                `json:"metrics"`
				Readme                 []assets.Readme                `json:"readme"`
				Meanings               []assets.Meaning               `json:"meanings"`
				SodaChecks             []assets.SodaCheck             `json:"sodaChecks"`
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

	var asset assets.AtlasGlossaryAttributes
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

	// Marshal the custom JSON
	return json.MarshalIndent(customJSON, "", "  ")
}

func (ag *AtlasGlossary) ToJSON() ([]byte, error) {
	return json.MarshalIndent(ag, "", "  ")
}

func (ag *AtlasGlossary) FromJSON(data []byte) error {
	return json.Unmarshal(data, ag)
}
