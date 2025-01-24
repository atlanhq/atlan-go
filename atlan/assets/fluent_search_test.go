package assets

import (
	"fmt"
	"testing"
	"time"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
	"github.com/stretchr/testify/assert"
)

const GlossaryDescription = "Automated testing of GO SDK."

var AnnouncementType = atlan.AnnouncementTypeWARNING

const (
	AnnouncementTitle   = "GO SDK testing."
	AnnouncementMessage = "Automated testing of the GO SDK."
)

func TestIntegrationFluentSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	ctx := NewContext()

	// Create a glossary
	g := &AtlasGlossary{}
	g.Creator(GlossaryName, atlan.AtlanIconAirplaneInFlight)
	g.Description = structs.StringPtr(GlossaryDescription)
	g.AnnouncementType = &AnnouncementType
	g.AnnouncementTitle = structs.StringPtr(AnnouncementTitle)
	g.AnnouncementMessage = structs.StringPtr(AnnouncementMessage)

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
		IncludeOnResults("description", "announcementType", "announcementTitle", "announcementMessage").
		Execute()
	if err != nil {
		fmt.Printf("Error executing search: %v\n", err)
		return
	}

	assert.NotNil(t, searchResult, "search result should not be nil")
	assert.Len(t, searchResult, 1, "number of glossaries should be 1")
	assert.Equal(t, GlossaryName, *searchResult[0].Entities[0].DisplayName, "glossary name should match")
	assert.Equal(t, GlossaryDescription, *searchResult[0].Entities[0].Description, "glossary description should exist")
	assert.Equal(t, AnnouncementType, *searchResult[0].Entities[0].AnnouncementType, "announcement type should exist")
	assert.Equal(t, AnnouncementTitle, *searchResult[0].Entities[0].AnnouncementTitle, "announcement title should exist")
	assert.Equal(t, AnnouncementMessage, *searchResult[0].Entities[0].AnnouncementMessage, "announcement message should exist")

	// Search for glossaries starts with letter G and sort them in ascending order by name
	searchResult, err = NewFluentSearch().
		PageSizes(10).
		ActiveAssets().
		Where(ctx.Glossary.NAME.StartsWith("gsdk", nil)).
		Sort(NAME, atlan.SortOrderAscending).
		Execute()
	if err != nil {
		fmt.Printf("Error executing search: %v\n", err)
		return
	}

	assert.Len(t, searchResult, 1, "number of glossaries should be 1")
	assert.Equal(t, "g", string((*searchResult[0].Entities[0].DisplayName)[0]), "glossary name should start with G")

	// Delete already created glossary
	deleteresponse, _ := PurgeByGuid([]string{response.MutatedEntities.CREATE[0].Guid})
	assert.NotNil(t, deleteresponse, "fetched glossary should not be nil")
}
