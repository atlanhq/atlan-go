// glossary_test.go
package model

import (
	"testing"
)

func TestGlossaryJSONConversion(t *testing.T) {
	glossaryInstance := Glossary{
		Entity: struct {
			TypeName   string `json:"typeName"`
			Attributes struct {
				Meanings []Meaning `json:"meanings"`
			} `json:"attributes"`
			Guid       string `json:"guid"`
			IsComplete bool   `json:"isIncomplete"`
			Status     string `json:"status"`
			CreatedBy  string `json:"createdBy"`
			CreateTime int    `json:"createTime"`
			UpdateTime int    `json:"updateTime"`
		}{
			TypeName:   "AtlasGlossaryTerm",
			Guid:       "c8bc0b75-bc8f-4cc3-bc2c-68703b7078c9",
			CreatedBy:  "karanjot.singh",
			IsComplete: false,
			CreateTime: 123456789,
			UpdateTime: 123456789,
			Attributes: struct {
				Meanings []Meaning `json:"meanings"`
			}{
				Meanings: []Meaning{
					{
						TermGuid:     "termGuid1",
						RelationGuid: "relationGuid1",
						DisplayText:  "displayText1",
						Confidence:   100,
					},
					{
						TermGuid:     "termGuid2",
						RelationGuid: "relationGuid2",
						DisplayText:  "displayText2",
						Confidence:   100,
					},
				},
			},
		},
	}

	// Convert the Glossary instance to JSON
	jsonData, err := glossaryInstance.ToJSON()
	if err != nil {
		t.Fatalf("Error converting to JSON: %v", err)
	}

	// Convert JSON back to a new Glossary instance
	newGlossary, err := FromJSON(jsonData)
	if err != nil {
		t.Fatalf("Error converting from JSON: %v", err)
	}

	// Compare the original and new Glossary instances
	if !AreEqual(&glossaryInstance, newGlossary) {
		t.Errorf("Original and new Glossary instances are different.\nOriginal: %+v\nNew: %+v", glossaryInstance, *newGlossary)
	}
}
