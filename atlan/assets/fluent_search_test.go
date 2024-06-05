package assets

import (
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationFluentSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := NewContext()

	// Create a glossary
	g := &AtlasGlossary{}
	g.Creator(GlossaryName, atlan.AtlanIconAirplaneInFlight)
	response, err := Save(g)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched glossary should not be nil")

	time.Sleep(5 * time.Second)
	// Search for glossary with Active Status and Name as GlossaryName
	searchResult, err := NewFluentSearch().
		PageSizes(10).
		ActiveAssets().
		Where(ctx.Glossary.NAME.Eq(GlossaryName)).
		//IncludeOnResults("guid").
		Execute()

	if err != nil {
		fmt.Printf("Error executing search: %v\n", err)
		return
	}

	assert.NotNil(t, searchResult, "search result should not be nil")
	assert.Equal(t, 1, len(searchResult), "number of glossaries should be 1")
	assert.Equal(t, GlossaryName, *searchResult[0].Entities[0].DisplayName, "glossary name should match")

	// Search for glossaries starts with letter G and sort them in ascending order by name
	searchResult, err = NewFluentSearch().
		PageSizes(10).
		ActiveAssets().
		Where(ctx.Glossary.NAME.StartsWith("gsdk", nil)).
		Sort(NAME, atlan.SortOrderAscending).
		Execute()

	assert.Equal(t, 1, len(searchResult), "number of glossaries should be 1")
	assert.Equal(t, "g", string((*searchResult[0].Entities[0].DisplayName)[0]), "glossary name should start with G")

	// Delete already created glossary
	deleteresponse, _ := PurgeByGuid([]string{response.MutatedEntities.CREATE[0].Guid})
	assert.NotNil(t, deleteresponse, "fetched glossary should not be nil")
}
