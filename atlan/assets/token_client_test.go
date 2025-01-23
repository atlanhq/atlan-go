package assets

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
	"github.com/stretchr/testify/assert"
)

var (
	TestDisplayName    = atlan.MakeUnique("test-api-token")
	TestDescription    = atlan.MakeUnique("Test API Token Description")
	MaxValiditySeconds = 409968000
)

func TestIntegrationTokenClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()

	// Test Create API Token
	createdToken := testCreateApiToken(t)

	// Test retrieval by Display Name
	testRetrieveTokenByName(t, *createdToken.Attributes.DisplayName, *createdToken.GUID)

	// Test retrieval by ID (Client ID or GUID)
	testRetrieveTokenByID(t, *createdToken.ClientID)
	testRetrieveTokenByGUID(t, *createdToken.GUID)

	// Test Update API Token
	testUpdateApiToken(t, *createdToken.GUID)

	// Test Purge API Token
	testPurgeApiToken(t, *createdToken.GUID)
}

func testCreateApiToken(t *testing.T) *structs.ApiToken {
	client := &TokenClient{}

	token, err := client.Create(&TestDisplayName, &TestDescription, nil, &MaxValiditySeconds)
	require.NoError(t, err, "error should be nil while creating an API token")
	assert.NotNil(t, token, "created token should not be nil")
	assert.Equal(t, TestDisplayName, *token.Attributes.DisplayName, "token display name should match")
	assert.Equal(t, TestDescription, *token.Attributes.Description, "token description should match")

	return token
}

func testRetrieveTokenByName(t *testing.T, displayName string, guid string) {
	client := &TokenClient{}

	token, err := client.GetByName(displayName)
	require.NoError(t, err, "error should be nil while retrieving token by display name")
	assert.NotNil(t, token, "retrieved token should not be nil")
	assert.Equal(t, displayName, *token.DisplayName, "token display name should match")
	assert.Equal(t, TestDescription, *token.Attributes.Description, "token description should match")
	assert.Equal(t, guid, *token.GUID, "token GUID should match")
}

func testRetrieveTokenByID(t *testing.T, clientID string) {
	client := &TokenClient{}

	token, err := client.GetByID(clientID)
	require.NoError(t, err, "error should be nil while retrieving token by client ID")
	assert.NotNil(t, token, "retrieved token should not be nil")
	assert.Equal(t, clientID, *token.ClientID, "token client ID should match")
}

func testRetrieveTokenByGUID(t *testing.T, guid string) {
	client := &TokenClient{}

	token, err := client.GetByGUID(guid)
	require.NoError(t, err, "error should be nil while retrieving token by GUID")
	assert.NotNil(t, token, "retrieved token should not be nil")
	assert.Equal(t, guid, *token.GUID, "token GUID should match")
}

func testUpdateApiToken(t *testing.T, guid string) {
	client := &TokenClient{}

	newDescription := atlan.MakeUnique("Updated description")
	newDisplayName := atlan.MakeUnique("Updated display name")
	token, err := client.Update(&guid, &newDisplayName, &newDescription, nil)
	require.NoError(t, err, "error should be nil while updating API token")
	assert.NotNil(t, token, "updated token should not be nil")
	assert.Equal(t, newDescription, *token.Attributes.Description, "token description should be updated")
}

func testPurgeApiToken(t *testing.T, guid string) {
	client := &TokenClient{}

	// Purge the API token
	err := client.Purge(guid)
	require.NoError(t, err, "error should be nil while purging API token")

	// Verify that the token is no longer retrievable
	token, err := client.GetByGUID(guid)
	assert.Nil(t, token, "token should be nil after purging")
}
