package assets

import (
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
	"testing"
)

var GlossaryName = atlan.MakeUnique("GLS")

func TestIntegrationGlossary(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()

	glossaryGUID, glossaryQualifiedName := testCreateGlossary(t)
	fmt.Printf("glossaryQn: %v\n", glossaryQualifiedName)
	testUpdateGlossary(t, glossaryGUID)
	testRetrieveGlossary(t, glossaryGUID)
	testRetrieveGlossarybyQualifiedName(t, glossaryQualifiedName)
	testDeleteGlossary(t, glossaryGUID)
}

func testCreateGlossary(t *testing.T) (string, string) {
	g := &AtlasGlossary{}
	// Create Glossary
	g.Creator(GlossaryName, atlan.AtlanIconAirplaneInFlight)
	response, err := Save(g)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(response.MutatedEntities.CREATE), "number of glossaries created should be 1")
	assert.Equal(t, 0, len(response.MutatedEntities.UPDATE), "number of glossaries updated should be 0")
	assert.Equal(t, 0, len(response.MutatedEntities.DELETE), "number of glossaries deleted should be 0")
	assetone := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, assetone, "glossary should not be nil")
	assert.Equal(t, GlossaryName, *assetone.Attributes.Name, "glossary name should match")
	assert.Equal(t, *g.TypeName, assetone.TypeName, "glossary type should match")

	return assetone.Guid, *assetone.Attributes.QualifiedName
}

func testUpdateGlossary(t *testing.T, glossaryGUID string) {
	g := &AtlasGlossary{}
	glossaryQualifiedName := GlossaryName + "-qual"
	DisplayName := "gsdk-test-update"
	g.Updater(GlossaryName, glossaryQualifiedName, glossaryGUID)
	g.DisplayName = &DisplayName
	updateresponse, err := Save(g)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, updateresponse, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(updateresponse.MutatedEntities.UPDATE), "number of glossaries updated should be 1")
	assert.Equal(t, *g.DisplayName, *updateresponse.MutatedEntities.UPDATE[0].Attributes.DisplayText, "glossary display name should match")
}

func testRetrieveGlossary(t *testing.T, glossaryGUID string) {
	glossary, err := GetByGuid[*AtlasGlossary](glossaryGUID)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, glossary, "fetched glossary should be nil")
	assert.Equal(t, glossaryGUID, *glossary.Guid, "glossary guid should match")
}

func testRetrieveGlossarybyQualifiedName(t *testing.T, glossaryQualifiedName string) {
	glossary, err := GetByQualifiedName[*AtlasGlossary](glossaryQualifiedName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, glossary, "fetched glossary should not be nil")
	assert.Equal(t, glossaryQualifiedName, *glossary.QualifiedName, "glossary qualified name should match")
}

func testDeleteGlossary(t *testing.T, glossaryGUID string) {
	deleteresponse, err := PurgeByGuid([]string{glossaryGUID})
	if err != nil {
		fmt.Println("Error:", err)
	}
	assert.NotNil(t, deleteresponse, "fetched glossary should not be nil")
	assert.Equal(t, 1, len(deleteresponse.MutatedEntities.DELETE), "number of glossaries deleted should be 1")
	assert.Equal(t, glossaryGUID, deleteresponse.MutatedEntities.DELETE[0].Guid, "glossary guid should match")
}
