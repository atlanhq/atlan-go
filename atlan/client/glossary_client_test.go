package client

import (
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ModuleName = atlan.MakeUnique("GLS")

func TestIntegrationGlossary(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	NewContext()
	g := &AtlasGlossary{}
	// Create Glossary
	g.Creator(ModuleName, atlan.AtlanIconAirplaneInFlight)
	response, err := Save(g)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, response, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(response.MutatedEntities.CREATE), "number of glossaries created should be 1")
	assert.Equal(t, 0, len(response.MutatedEntities.UPDATE), "number of glossaries updated should be 0")
	assert.Equal(t, 0, len(response.MutatedEntities.DELETE), "number of glossaries deleted should be 0")
	assetone := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, assetone, "glossary should not be nil")
	assert.Equal(t, ModuleName, *assetone.Attributes.Name, "glossary name should match")
	assert.Equal(t, *g.TypeName, assetone.TypeName, "glossary type should match")

	// Update Glossary
	glossaryQualifiedName := *assetone.Attributes.QualifiedName
	glossaryGUID := assetone.Guid
	DisplayName := "gsdk-test-update"
	g.Updater(ModuleName, glossaryQualifiedName, glossaryGUID)
	g.DisplayName = &DisplayName
	updateresponse, err := Save(g)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, updateresponse, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(updateresponse.MutatedEntities.UPDATE), "number of glossaries updated should be 1")
	assert.Equal(t, *g.DisplayName, *updateresponse.MutatedEntities.UPDATE[0].Attributes.DisplayText, "glossary display name should match")

	// Delete Glossary (Hard)
	deleteresponse, _ := PurgeByGuid([]string{glossaryGUID})
	assert.NotNil(t, deleteresponse, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(deleteresponse.MutatedEntities.DELETE), "number of glossaries deleted should be 1")
	assert.Equal(t, glossaryGUID, deleteresponse.MutatedEntities.DELETE[0].Guid, "glossary guid should match")

	// Retrieve Glossary by Guid
	glossary, err := GetGlossaryByGuid(glossaryGUID)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, glossary, "fetched glossary should be nil")
	assert.Equal(t, glossaryGUID, glossary.Guid, "glossary guid should match")

}
