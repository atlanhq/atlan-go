package assets

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

const UserEmail = "gsdk-test-user@atlan.com"

// UserEmail     = fmt.Sprintf("%s@atlan.com", strings.ToLower(atlan.MakeUnique("test_user")))
var WorkspaceRole = "$guest"

func TestIntegrationUserClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	NewContext()

	// ctx.EnableLogging("debug")

	// Test user creation
	createdUser := getOrCreateTestUser(t)

	// Test retrieval by email
	testRetrieveUserByEmail(t, createdUser.Email)

	// Test retrieval by username
	testRetrieveUserByUsername(t, *createdUser.Username)

	// Test updating user's role
	testChangeUserRole(t, createdUser.ID)
}

func getOrCreateTestUser(t *testing.T) *AtlanUser {
	client := &UserClient{}

	// Check if user already exists
	existingUser, err := client.GetByEmail(UserEmail, 1, 0)
	if err == nil && len(existingUser) > 0 {
		t.Logf("User already exists: %s", UserEmail)
		return &existingUser[0]
	}

	t.Logf("User does not exist, creating new user: %s", UserEmail)
	users := []AtlanUser{
		{
			Email:         UserEmail,
			WorkspaceRole: WorkspaceRole,
		},
	}

	createdUsers, err := client.CreateUsers(users, true)

	require.NoError(t, err, "error should be nil while creating a user")
	assert.NotNil(t, createdUsers, "created users should not be nil")
	assert.Len(t, createdUsers, 1, "exactly one user should be created")

	user := createdUsers[0]
	assert.Equal(t, UserEmail, user.Email, "user email should match")
	assert.Equal(t, WorkspaceRole, user.WorkspaceRole, "user role should match")

	return &user
}

func testRetrieveUserByEmail(t *testing.T, email string) {
	client := &UserClient{}

	users, err := client.GetByEmail(email, 1, 0)
	require.NoError(t, err, "error should be nil while retrieving user by email")
	assert.NotNil(t, users, "retrieved users should not be nil")
	assert.Len(t, users, 1, "exactly one user should be retrieved")

	user := users[0]
	assert.Equal(t, email, user.Email, "user email should match")
}

func testRetrieveUserByUsername(t *testing.T, username string) {
	client := &UserClient{}

	user, err := client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while retrieving user by username")
	assert.NotNil(t, user, "retrieved user should not be nil")
	assert.Equal(t, username, *user.Username, "user username should match")
}

func testChangeUserRole(t *testing.T, userID string) {
	client := &UserClient{}

	role := "$member"
	newRoleID, _ := GetRoleIDForRoleName(role)

	err := client.ChangeUserRole(userID, newRoleID)
	require.NoError(t, err, "error should be nil while updating user's role")

	// Verify the role change
	users, err := client.GetByEmails([]string{UserEmail}, 1, 0)
	require.NoError(t, err, "error should be nil while retrieving updated user")
	assert.Len(t, users, 1, "exactly one user should be retrieved")
	assert.Equal(t, role, users[0].WorkspaceRole, "user role ID should match the updated role")

	// Revert to original role
	revertRole := "$guest"
	revertRoleId, _ := GetRoleIDForRoleName(revertRole)
	err = client.ChangeUserRole(userID, revertRoleId)
	require.NoError(t, err, "error should be nil while updating user's role")

	users, err = client.GetByEmails([]string{UserEmail}, 1, 0)
	assert.Len(t, users, 1, "exactly one user should be retrieved")
	assert.Equal(t, revertRole, users[0].WorkspaceRole, "user role ID should match the updated role")
}
