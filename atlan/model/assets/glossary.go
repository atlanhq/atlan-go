package assets

import (
	"encoding/json"
	"github.com/atlanhq/atlan-go/atlan"
)

type AtlasGlossaryAttributes AtlasGlossary

type AtlasGlossary struct {
	Relation
	Asset
	ShortDescription     *string                  `json:"shortDescription,omitempty"`
	LongDescription      *string                  `json:"longDescription,omitempty"`
	Language             *string                  `json:"language,omitempty"`
	Usage                *string                  `json:"usage,omitempty"`
	AdditionalAttributes *map[string]string       `json:"additionalAttributes,omitempty"`
	Terms                *[]AtlasGlossaryTerm     `json:"terms,omitempty"`
	Categories           *[]AtlasGlossaryCategory `json:"categories,omitempty"`
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
				SchemaRegistrySubjects []SchemaRegistrySubject `json:"schemaRegistrySubjects"`
				McMonitors             []MCMonitor             `json:"mcMonitors"`
				Terms                  []AtlasGlossaryTerm     `json:"terms"`
				OutputPortDataProducts []string                `json:"outputPortDataProducts"`
				Files                  []File                  `json:"files"`
				McIncidents            []MCIncident            `json:"mcIncidents"`
				Links                  []Link                  `json:"links"`
				Categories             []AtlasGlossaryCategory `json:"categories"`
				Metrics                []Metric                `json:"metrics"`
				Readme                 []Readme                `json:"readme"`
				Meanings               []Meaning               `json:"meanings"`
				SodaChecks             []SodaCheck             `json:"sodaChecks"`
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

	var asset AtlasGlossaryAttributes
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

func (ag *AtlasGlossary) ToJSON() ([]byte, error) {
	return json.MarshalIndent(ag, "", "  ")
}

func (ag *AtlasGlossary) FromJSON(data []byte) error {
	return json.Unmarshal(data, ag)
}
