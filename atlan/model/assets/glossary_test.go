package assets

import (
	"atlan-go/atlan"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stringPtr(s string) *string {
	return &s
}

func TestUnmarshalling(t *testing.T) {
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

	fmt.Println("Unmarshalled glossary:", glossary)

	// Assert that there is no error during unmarshalling
	assert.NoError(t, err, "Error unmarshalling JSON")

	// Assert that the unmarshalled glossary matches the expected glossary
	assert.Equal(t, "Test Glossary", *glossary.Name, "Unexpected glossary name")
	assert.Equal(t, "test_glossary", *glossary.QualifiedName, "Unexpected glossary qualified name")
	assert.Equal(t, atlan.AtlanIconAirplaneInFlight, glossary.AssetIcon, "Unexpected glossary asset icon")
	assert.Equal(t, "Short description", *glossary.ShortDescription, "Unexpected glossary short description")
	assert.Equal(t, "Long description", *glossary.LongDescription, "Unexpected glossary long description")
	assert.Equal(t, "English", *glossary.Language, "Unexpected glossary language")
	assert.Equal(t, "Usage details", *glossary.Usage, "Unexpected glossary usage")
}

func TestMarshalling(t *testing.T) {
	// Define a sample AtlasGlossary object
	glossary := AtlasGlossary{
		Asset: Asset{
			Referenceable: Referenceable{
				TypeName: stringPtr("AtlasGlossary"),
				Guid:     stringPtr("fc36342b-ddb5-44ba-b774-4c90cc66d5a2"),
				Status:   atlan.AtlanStatusPtr("ACTIVE"),
			},
			Name:          stringPtr("Test Glossary"),
			QualifiedName: stringPtr("test_glossary"),
			AssetIcon:     atlan.AtlanIconAirplaneInFlight,
		},
		ShortDescription: stringPtr("Short description"),
		LongDescription:  stringPtr("Long description"),
		Language:         stringPtr("English"),
		Usage:            stringPtr("Usage details"),
	}

	// Marshal the AtlasGlossary object into JSON
	jsonData, err := glossary.ToJSON()

	fmt.Println("Marshalled JSON:", string(jsonData))

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

func TestMarshallingAndUnmarshalling(t *testing.T) {
	// Define a sample AtlasGlossary object
	glossary := AtlasGlossary{
		Asset: Asset{
			Referenceable: Referenceable{TypeName: stringPtr("AtlasGlossary")},
			Name:          stringPtr("Test Glossary"),
			QualifiedName: stringPtr("test_glossary"),
			AssetIcon:     atlan.AtlanIconAirplaneInFlight,
		},
		ShortDescription: stringPtr("Short description"),
		LongDescription:  stringPtr("Long description"),
		Language:         stringPtr("English"),
		Usage:            stringPtr("Usage details"),
	}

	// Marshal the AtlasGlossary object to JSON
	jsonData, err := glossary.ToJSON()
	assert.NoError(t, err, "Error marshalling AtlasGlossary to JSON")
	fmt.Println("Marshalled JSON:", string(jsonData))

	// Unmarshal the JSON data back to AtlasGlossary
	decodedGlossary, err := FromJSON(jsonData)
	assert.NoError(t, err, "Error unmarshalling JSON to AtlasGlossary")
	fmt.Println("Unmarshalled glossary:", decodedGlossary)

	// Assert that the unmarshalled glossary matches the original glossary
	assert.Equal(t, glossary, decodedGlossary, "Unmarshalled glossary does not match original glossary")
}
