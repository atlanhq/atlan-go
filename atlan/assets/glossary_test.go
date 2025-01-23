package assets

import (
	"encoding/json"
	"testing"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"

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
		Asset: structs.Asset{
			Referenceable: structs.Referenceable{
				TypeName:      structs.StringPtr("AtlasGlossary"),
				Guid:          structs.StringPtr("fc36342b-ddb5-44ba-b774-4c90cc66d5a2"),
				Status:        atlan.AtlanStatusPtr("ACTIVE"),
				QualifiedName: structs.StringPtr("test_glossary"),
			},
			Name:      structs.StringPtr("Test Glossary"),
			AssetIcon: atlan.AtlanIconPtr(atlan.AtlanIconAirplaneInFlight),
		},
		ShortDescription: structs.StringPtr("Short description"),
		LongDescription:  structs.StringPtr("Long description"),
		Language:         structs.StringPtr("English"),
		Usage:            structs.StringPtr("Usage details"),
	}

	// Marshal the AtlasGlossary object into JSON
	jsonData, err := glossary.ToJSON()

	// Define the expected JSON data
	// This expected JSON data is from the Custom Marshal function of glossary which is partially implemented therefore we only have few fields marshalled that are required for creator and updater.
	expectedJSON := []byte(`{
	  "attributes": {
		"name": "Test Glossary",
		"qualifiedName": "test_glossary"
	  },
	  "guid": "fc36342b-ddb5-44ba-b774-4c90cc66d5a2",
	  "relationshipAttributes": {},
	  "typeName": "AtlasGlossary"
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
	var term structs.AtlasGlossaryTerm
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
	term := structs.AtlasGlossaryTerm{
		Asset: structs.Asset{
			Referenceable: structs.Referenceable{
				TypeName:  structs.StringPtr("AtlasGlossaryTerm"),
				Guid:      structs.StringPtr("433b1b64-0b16-4812-9bae-14b13e9bd645"),
				Status:    atlan.AtlanStatusPtr("ACTIVE"),
				CreatedBy: structs.StringPtr("user1"),
				UpdatedBy: structs.StringPtr("user2"),
			},
		},
		ShortDescription: structs.StringPtr("Short description"),
		LongDescription:  structs.StringPtr("Long description"),
		Example:          structs.StringPtr("Example Text"),
		Abbreviation:     structs.StringPtr("Abbreviation Text"),
		Usage:            structs.StringPtr("Usage Text"),
		AdditionalAttributes: &map[string]string{
			"key": "value",
		},
		Anchor: &structs.AtlasGlossary{
			Asset: structs.Asset{
				Referenceable: structs.Referenceable{
					TypeName: structs.StringPtr("AtlasGlossaryTerm"),
					Guid:     structs.StringPtr("562067ed-c56a-470d-9306-488d9c6d6448"),
				},
			},
			Relation: structs.Relation{
				DisplayText:        structs.StringPtr("Display Text"),
				EntityStatus:       structs.StringPtr("ACTIVE"),
				RelationshipType:   structs.StringPtr("AtlasGlossaryTermAnchor"),
				RelationshipGuid:   structs.StringPtr("abe7f160-182e-4c61-bc8e-e3392404611b"),
				RelationshipStatus: atlan.AtlanStatusPtr("ACTIVE"),
			},
		},
	}

	// Marshal the AtlasGlossaryTerm object into JSON
	jsonData, err := json.Marshal(term)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	// Define the expected JSON data
	expectedJSON := []byte(`{"typeName":"AtlasGlossaryTerm","guid":"433b1b64-0b16-4812-9bae-14b13e9bd645","createdBy":"user1","updatedBy":"user2","status":"ACTIVE","shortDescription":"Short description","longDescription":"Long description","example":"Example Text","abbreviation":"Abbreviation Text","usage":"Usage Text","additionalAttributes":{"key":"value"},"anchor":{"entityStatus":"ACTIVE","relationshipType":"AtlasGlossaryTermAnchor","relationshipGuid":"abe7f160-182e-4c61-bc8e-e3392404611b","relationshipStatus":"ACTIVE","typeName":"AtlasGlossaryTerm","guid":"562067ed-c56a-470d-9306-488d9c6d6448"}}`)

	// Assert that the marshalled JSON matches the expected JSON data
	assert.JSONEq(t, string(expectedJSON), string(jsonData), "Marshalled JSON does not match the expected JSON")
}
