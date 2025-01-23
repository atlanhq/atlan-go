package assets

import (
	"testing"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
)

var PersonaName = atlan.MakeUnique("Persona")

func TestIntegrationPersona(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()

	personaID, personaQualifiedName := testCreatePersona(t)
	testRetrievePersona(t, personaID)
	testUpdatePersona(t, personaQualifiedName)
	testDeletePersona(t, personaID)
}

func testCreatePersona(t *testing.T) (string, string) {
	p := &Persona{}
	// Create Persona
	p.Creator(PersonaName)
	response, err := Save(p)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched persona should not be nil")
	assert.Len(t, response.MutatedEntities.CREATE, 1, "number of personas created should be 1")
	assert.Empty(t, response.MutatedEntities.UPDATE, "number of personas updated should be 0")
	assert.Empty(t, response.MutatedEntities.DELETE, "number of personas deleted should be 0")
	CreatedPersona := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, CreatedPersona, "persona should not be nil")
	assert.Equal(t, PersonaName, *CreatedPersona.Attributes.Name, "persona name should match")
	assert.Equal(t, *p.TypeName, CreatedPersona.TypeName, "persona type should match")

	return CreatedPersona.Guid, *CreatedPersona.Attributes.QualifiedName
}

func testRetrievePersona(t *testing.T, personaID string) {
	persona, err := GetByGuid[*Persona](personaID)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, persona, "fetched persona should not be nil")
	assert.Equal(t, PersonaName, *persona.Name, "persona name should match")
}

func testUpdatePersona(t *testing.T, personaQualifiedName string) {
	p := &Persona{}
	Name := "gsdk-test-update"
	p.Updater(personaQualifiedName, PersonaName, true)
	p.Name = &Name
	updateresponse, err := Save(p)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, updateresponse, "fetched persona should not be nil")
	assert.Len(t, updateresponse.MutatedEntities.UPDATE, 1, "number of personas updated should be 1")
	assert.Equal(t, *p.Name, *updateresponse.MutatedEntities.UPDATE[0].Attributes.Name, "persona display name should match")
}

func testDeletePersona(t *testing.T, personaID string) {
	deleteresponse, err := PurgeByGuid([]string{personaID})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, deleteresponse, "fetched persona should not be nil")
	assert.Len(t, deleteresponse.MutatedEntities.DELETE, 1, "number of personas deleted should be 1")
	assert.Equal(t, personaID, deleteresponse.MutatedEntities.DELETE[0].Guid, "persona guid should match")
}

// Add tests related to creating policies using Persona when the Managing connections would be implemented by the sdk
