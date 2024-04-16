package assets

import (
	"encoding/json"
	"github.com/atlanhq/atlan-go/atlan"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtlasGlossaryUnmarshalling(t *testing.T) {
	// Define a sample JSON data representing an AtlasGlossary
	jsonData := []byte(`{
		"referredEntities":{},
		"entity":{
			"typeName":"AtlasGlossary",
			"attributes":{
				"name":"Test Glossary",
				"qualifiedName":"test_glossary",
				"assetIcon":"PhAirplaneInFlight",
				"shortDescription":"Short description",
				"longDescription":"Long description",
				"language":"English",
				"usage":"Usage details"
			}
		}
	}`)

	// Unmarshal the JSON data into an AtlasGlossary object
	var glossary AtlasGlossary
	err := glossary.UnmarshalJSON(jsonData)

	// Assert that there is no error during unmarshalling
	assert.NoError(t, err, "Error unmarshalling JSON")

	// Assert that the unmarshalled glossary matches the expected glossary
	assert.Equal(t, "Test Glossary", *glossary.Name, "Unexpected glossary name")
	assert.Equal(t, "test_glossary", *glossary.QualifiedName, "Unexpected glossary qualified name")
	assert.Equal(t, atlan.AtlanIconAirplaneInFlight, *glossary.AssetIcon, "Unexpected glossary asset icon")
	assert.Equal(t, "Short description", *glossary.ShortDescription, "Unexpected glossary short description")
	assert.Equal(t, "Long description", *glossary.LongDescription, "Unexpected glossary long description")
	assert.Equal(t, "English", *glossary.Language, "Unexpected glossary language")
	assert.Equal(t, "Usage details", *glossary.Usage, "Unexpected glossary usage")
}

func TestAtlasGlossaryMarshalling(t *testing.T) {
	// Define a sample AtlasGlossary object
	glossary := AtlasGlossary{
		Asset: Asset{
			Referenceable: Referenceable{
				TypeName:      StringPtr("AtlasGlossary"),
				Guid:          StringPtr("fc36342b-ddb5-44ba-b774-4c90cc66d5a2"),
				Status:        atlan.AtlanStatusPtr("ACTIVE"),
				QualifiedName: StringPtr("test_glossary"),
			},
			Name:      StringPtr("Test Glossary"),
			AssetIcon: atlan.AtlanIconPtr(atlan.AtlanIconAirplaneInFlight),
		},
		ShortDescription: StringPtr("Short description"),
		LongDescription:  StringPtr("Long description"),
		Language:         StringPtr("English"),
		Usage:            StringPtr("Usage details"),
	}

	// Marshal the AtlasGlossary object into JSON
	jsonData, err := glossary.ToJSON()

	// Define the expected JSON data
	expectedJSON := []byte(`{
	  "typeName": "AtlasGlossary",
	  "guid": "fc36342b-ddb5-44ba-b774-4c90cc66d5a2",
	  "status": "ACTIVE",
	  "assetIcon": "PhAirplaneInFlight",
	  "name": "Test Glossary",
	  "qualifiedName": "test_glossary",
	  "shortDescription": "Short description",
	  "longDescription": "Long description",
	  "language": "English",
	  "usage": "Usage details"
	}`)

	// Assert that there is no error during marshalling
	assert.NoError(t, err, "Error marshalling AtlasGlossary to JSON")

	// Assert that the marshalled JSON data matches the expected JSON data
	assert.JSONEq(t, string(expectedJSON), string(jsonData), "Marshalled JSON does not match expected JSON")
}

func TestAtlasGlossaryTermUnmarshalling(t *testing.T) {
	// Define a sample JSON data representing an AtlasGlossaryTerm
	jsonData := []byte(`{
		"entity": {
			"typeName": "AtlasGlossaryTerm",
			"attributes": {
				"shortDescription": "Short description",
				"longDescription": "Long description",
				"example": "Example",
				"abbreviation": "Abbreviation",
				"usage": "Usage",
				"anchor": {
					"guid": "some_guid",
					"typeName": "AtlasGlossary",
					"entityStatus": "ACTIVE",
					"displayText": "Display Text",
					"relationshipType": "RELATIONSHIP_TYPE",
					"relationshipGuid": "RELATIONSHIP_GUID",
					"relationshipStatus": "ACTIVE"
				}
			}
		}
	}`)

	// Unmarshal the JSON data into an AtlasGlossaryTerm object
	var term AtlasGlossaryTerm
	err := term.UnmarshalJSON(jsonData)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Perform assertions on the unmarshalled object's fields
	assert.Equal(t, "AtlasGlossaryTerm", *term.TypeName, "Unexpected term type name")
	assert.Equal(t, "Short description", *term.ShortDescription, "Unexpected term short description")
	assert.Equal(t, "Long description", *term.LongDescription, "Unexpected term long description")
	assert.Equal(t, "Example", *term.Example, "Unexpected term example")
	assert.Equal(t, "Abbreviation", *term.Abbreviation, "Unexpected term abbreviation")
	assert.Equal(t, "Usage", *term.Usage, "Unexpected term usage")

	// Assert anchor details
	assert.NotNil(t, term.Anchor, "Anchor should not be nil")
	assert.Equal(t, "AtlasGlossary", *term.Anchor.TypeName, "Unexpected anchor type name")
	assert.Equal(t, "some_guid", *term.Anchor.Guid, "Unexpected anchor guid")

}

func TestAtlasGlossaryTermMarshalling(t *testing.T) {
	// Define a sample AtlasGlossaryTerm object
	term := AtlasGlossaryTerm{
		Asset: Asset{
			Referenceable: Referenceable{
				TypeName:  StringPtr("AtlasGlossaryTerm"),
				Guid:      StringPtr("433b1b64-0b16-4812-9bae-14b13e9bd645"),
				Status:    atlan.AtlanStatusPtr("ACTIVE"),
				CreatedBy: StringPtr("user1"),
				UpdatedBy: StringPtr("user2"),
			},
		},
		ShortDescription: StringPtr("Short description"),
		LongDescription:  StringPtr("Long description"),
		Example:          StringPtr("Example Text"),
		Abbreviation:     StringPtr("Abbreviation Text"),
		Usage:            StringPtr("Usage Text"),
		AdditionalAttributes: &map[string]string{
			"key": "value",
		},
		Anchor: &AtlasGlossary{
			Asset: Asset{Referenceable: Referenceable{
				TypeName: StringPtr("AtlasGlossaryTerm"),
				Guid:     StringPtr("562067ed-c56a-470d-9306-488d9c6d6448"),
			},
			},
			Relation: Relation{
				displayText:        StringPtr("Display Text"),
				entityStatus:       StringPtr("ACTIVE"),
				relationshipType:   StringPtr("AtlasGlossaryTermAnchor"),
				relationshipGuid:   StringPtr("abe7f160-182e-4c61-bc8e-e3392404611b"),
				relationshipStatus: atlan.AtlanStatusPtr("ACTIVE"),
			},
		},
	}

	// Marshal the AtlasGlossaryTerm object into JSON
	jsonData, err := json.Marshal(term)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	// Define the expected JSON data
	expectedJSON := []byte(`{"typeName":"AtlasGlossaryTerm","guid":"433b1b64-0b16-4812-9bae-14b13e9bd645","createdBy":"user1","updatedBy":"user2","status":"ACTIVE","shortDescription":"Short description","longDescription":"Long description","example":"Example Text","abbreviation":"Abbreviation Text","usage":"Usage Text","additionalAttributes":{"key":"value"},"anchor":{"typeName":"AtlasGlossaryTerm","guid":"562067ed-c56a-470d-9306-488d9c6d6448"}}`)

	// Assert that the marshalled JSON matches the expected JSON data
	assert.JSONEq(t, string(expectedJSON), string(jsonData), "Marshalled JSON does not match the expected JSON")
}
