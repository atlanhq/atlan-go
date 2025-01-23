package assets

import (
	"testing"
	"time"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
)

var PurposeName = atlan.MakeUnique("Purpose")

func TestIntegrationPurpose(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()
	// ctx.EnableLogging("debug")

	purposeID, purposeQualifiedName := testCreatePurpose(t)
	testRetrievePurpose(t, purposeID)
	testPurposeCreateMetadataPolicy(t, purposeID)
	testPurposeCreateDataPolicy(t, purposeID)
	testFindPurposesByName(t)
	testUpdatePurpose(t, purposeQualifiedName)
	testDeletePurpose(t, purposeID)
}

func testCreatePurpose(t *testing.T) (string, string) {
	p := &Purpose{}
	// Create Purpose
	atlanTags := []string{"Issue", "Confidential"}
	err := p.Creator(PurposeName, atlanTags)
	assert.NoError(t, err, "creator should not return an error")

	response, err := Save(p)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched purpose should not be nil")
	assert.Equal(t, 1, len(response.MutatedEntities.CREATE), "number of purposes created should be 1")
	assert.Equal(t, 0, len(response.MutatedEntities.UPDATE), "number of purposes updated should be 0")
	assert.Equal(t, 0, len(response.MutatedEntities.DELETE), "number of purposes deleted should be 0")
	CreatedPurpose := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, CreatedPurpose, "purpose should not be nil")
	assert.Equal(t, PurposeName, *CreatedPurpose.Attributes.Name, "purpose name should match")
	assert.Equal(t, *p.TypeName, CreatedPurpose.TypeName, "purpose type should match")

	return CreatedPurpose.Guid, *CreatedPurpose.Attributes.QualifiedName
}

func testRetrievePurpose(t *testing.T, purposeID string) {
	purpose, err := GetByGuid[*Purpose](purposeID)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, purpose, "fetched purpose should not be nil")
	assert.Equal(t, PurposeName, *purpose.Name, "purpose name should match")
}

func testFindPurposesByName(t *testing.T) {
	time.Sleep(3 * time.Second)
	purposes, err := FindPurposesByName(PurposeName)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, purposes.Entities, "fetched purposes should not be nil")
	assert.Equal(t, int64(1), purposes.ApproximateCount, "number of purposes fetched should be 1")
	assert.Equal(t, PurposeName, *purposes.Entities[0].Name, "purpose name should match")
}

func testPurposeCreateMetadataPolicy(t *testing.T, purposeID string) {
	p := &Purpose{}
	policy, err := p.CreateMetadataPolicy(
		PurposeName,
		purposeID,
		atlan.AuthPolicyTypeAllow,
		[]atlan.PurposeMetadataAction{
			atlan.PurposeMetadataActionRead,
		},
		nil,
		nil,
		true,
	)
	response, err := Save(policy)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched policy should not be nil")
	assert.Equal(t, 1, len(response.MutatedEntities.CREATE), "number of policies added should be 1")
	CreatedPolicy := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, CreatedPolicy, "policy should not be nil")
}

func testPurposeCreateDataPolicy(t *testing.T, purposeID string) {
	p := &Purpose{}
	policy, err := p.CreateDataPolicy(
		PurposeName,
		purposeID,
		atlan.AuthPolicyTypeAllow,
		nil,
		nil,
		true,
	)
	response, err := Save(policy)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, response, "fetched policy should not be nil")
	assert.Equal(t, 1, len(response.MutatedEntities.CREATE), "number of policies added should be 1")
	CreatedPolicy := response.MutatedEntities.CREATE[0]
	assert.NotNil(t, CreatedPolicy, "policy should not be nil")
}

func testUpdatePurpose(t *testing.T, purposeQualifiedName string) {
	p := &Purpose{}
	NewName := atlan.MakeUnique("test-update-purpose")
	Description := atlan.MakeUnique("test-update-purpose-description")
	err := p.Updater(purposeQualifiedName, PurposeName, true)
	assert.NoError(t, err, "updater should not return an error")

	p.Name = &NewName
	p.Description = &Description
	UpdaterResponse, err := Save(p)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.NotNil(t, UpdaterResponse, "fetched purpose should not be nil")
	assert.Equal(t, 1, len(UpdaterResponse.MutatedEntities.UPDATE), "number of purposes updated should be 1")
	assert.Equal(t, *p.Name, *UpdaterResponse.MutatedEntities.UPDATE[0].Attributes.Name, "purpose name should match")
	assert.Equal(t, *p.Description, *UpdaterResponse.MutatedEntities.UPDATE[0].Attributes.Description, "purpose description should match")
}

func testDeletePurpose(t *testing.T, purposeID string) {
	DeleteResponse, err := PurgeByGuid([]string{purposeID})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// for _, deleted := range DeleteResponse.MutatedEntities.DELETE {
	//	t.Logf("Deleted: %v", deleted)
	//}
	assert.NotNil(t, DeleteResponse, "fetched purpose should not be nil")
	assert.Equal(t, 3, len(DeleteResponse.MutatedEntities.DELETE), "number of purposes deleted should be 3") // 3 because of the metadata and data policies

	// Collect GUIDs from the server response
	serverGuids := make([]string, len(DeleteResponse.MutatedEntities.DELETE))
	for i, deleted := range DeleteResponse.MutatedEntities.DELETE {
		serverGuids[i] = deleted.Guid
	}

	// Ensure the expected purposeID is in the list of server-provided GUIDs
	assert.Contains(t, serverGuids, purposeID, "purpose guid should match one of the server-provided GUIDs")
}
