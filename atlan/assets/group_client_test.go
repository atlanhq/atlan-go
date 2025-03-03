package assets

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
)

var TestGroupAlias = strings.ToLower(atlan.MakeUnique("test_group"))

func TestIntegrationGroupClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()

	// ctx.EnableLogging("debug")

	// Test creating a group
	group := testCreateGroup(t)

	// Test Retrieve all groups
	testRetrieveAllGroups(t)

	// Test retrieving group by name
	testRetrieveGroupByName(t)

	// Test Add Users to group
	testAddUsersToGroup(t, *group.ID)

	// Test updating the group
	testUpdateGroup(t, *group.ID, *group.Path)

	// Test Retrieve Members
	memberID := testRetrieveMembers(t, *group.ID)

	// Test removing users from the group
	testRemoveUsersFromGroup(t, *group.ID, memberID)

	// Test purging the group
	testPurgeGroup(t, *group.ID)
}

func testCreateGroup(t *testing.T) *AtlanGroup {
	client := &GroupClient{}

	// Create a Group
	group := AtlanGroup{}
	toBeCreated, err := group.Create(TestGroupAlias)
	require.NoError(t, err, "error should be nil while creating group object")
	// users := []string{}
	response, err := client.Create(toBeCreated, nil)

	require.NoError(t, err, "error should be nil while creating group")
	assert.NotNil(t, response, "response should not be nil")
	assert.NotNil(t, response.Group, "created group should not be nil")

	// Retrieve created group
	groups, err := client.GetByName(TestGroupAlias, 10, 0)
	require.NoError(t, err, "error should be nil while retrieving group by name")
	assert.NotNil(t, groups, "retrieved groups should not be nil")
	assert.Equal(t, TestGroupAlias, *groups[0].Alias, "group alias should match")

	return groups[0]
}

func testRetrieveAllGroups(t *testing.T) {
	client := &GroupClient{}

	groups, err := client.GetAll(10, 0, "")
	require.NoError(t, err, "error should be nil while retrieving all groups")
	assert.NotNil(t, groups, "retrieved groups should not be nil")
	assert.GreaterOrEqual(t, len(groups), 1, "at least one group should exist")
	// for _, group := range groups {
	//	log.Printf("Group name: %s", *group.Name)
	//}
}

func testRetrieveGroupByName(t *testing.T) {
	client := &GroupClient{}

	groups, err := client.GetByName(TestGroupAlias, 10, 0)
	require.NoError(t, err, "error should be nil while retrieving group by name")
	assert.NotNil(t, groups, "retrieved groups should not be nil")
	assert.GreaterOrEqual(t, len(groups), 1, "at least one group should be retrieved")
	assert.Equal(t, TestGroupAlias, *groups[0].Alias, "group alias should match")
	// for _, group := range groups {
	//	log.Printf("Retrieved group name: %s", *group.Name)
	//}
}

func testAddUsersToGroup(t *testing.T, groupID string) {
	client := &GroupClient{}

	// This is a dependency issue on user-client test (which creates the user). This test runs before the user-client test
	// which creates an issue if the user doesn't exist. So, we add a check here.
	// Checks if the test user exists or not and creates it if it doesn't
	testUser := getOrCreateTestUser(t)
	// Directly uses the Email from the created / existing user
	user, err := client.UserClient.GetByEmail(testUser.Email, 1, 0)
	require.NoError(t, err, "error should be nil while getting user by email")
	err = client.UserClient.AddUserToGroups(user[0].ID, []string{groupID})
	require.NoError(t, err, "error should be nil while adding user to group")

	// Verify user was added (Also tests GetMembers Endpoint)
	members, err := client.GetMembers(groupID, nil)
	require.NoError(t, err, "error should be nil while retrieving group members after adding user")
	found := false
	for _, member := range members {
		if member.ID == user[0].ID {
			found = true
			break
		}
	}
	assert.True(t, found, "added user should be present in the group")
}

func testUpdateGroup(t *testing.T, groupID string, path string) {
	client := &GroupClient{}
	group := AtlanGroup{}

	toBeUpdated, err := group.Updater(groupID, path)
	require.NoError(t, err, "error should be nil while creating updater object")

	updatedName := []string{atlan.MakeUnique("updated_alias")}
	updatedDescription := []string{"This is the updated description"}
	toBeUpdated.Attributes.Description = updatedDescription
	toBeUpdated.Attributes.Alias = updatedName

	err = client.Update(toBeUpdated)
	require.NoError(t, err, "error should be nil while updating group")

	// Verify the update
	updatedGroups, err := client.GetByName(updatedName[0], 1, 0)
	require.NoError(t, err, "error should be nil while retrieving updated group")
	assert.Len(t, updatedGroups, 1, "exactly one group should be retrieved")
	assert.Equal(t, updatedDescription, updatedGroups[0].Attributes.Description, "group description should match the updated value")
}

func testRetrieveMembers(t *testing.T, guid string) string {
	client := &GroupClient{}

	members, err := client.GetMembers(guid, nil)
	require.NoError(t, err, "error should be nil while retrieving group members")
	assert.NotEmpty(t, members, "group should have at least one member")

	return members[0].ID
}

func testRemoveUsersFromGroup(t *testing.T, guid string, memberID string) {
	client := &GroupClient{}

	userIDs := []string{memberID}

	err := client.RemoveUsers(guid, userIDs)
	require.NoError(t, err, "error should be nil while removing users from group")

	// Verify users are removed
	members, err := client.GetMembers(guid, nil)
	require.NoError(t, err, "error should be nil while retrieving group members after removal")
	// Ensure none of the removed user IDs are in the member list
	for _, member := range members {
		for _, userID := range userIDs {
			assert.NotEqual(t, userID, member.ID, "removed user ID should not be in the group")
		}
	}
}

func testPurgeGroup(t *testing.T, groupID string) {
	client := &GroupClient{}
	err := client.Purge(groupID)
	require.NoError(t, err, "error should be nil while purging the group")
}
