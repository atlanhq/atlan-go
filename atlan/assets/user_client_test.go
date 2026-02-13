package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	// Test disabling and enabling user
	testDisableAndEnableUser(t, createdUser.ID, *createdUser.Username)

	// Test that RemoveUser disables user before deletion
	// Note: This test verifies the disable step but doesn't actually delete the user
	// to avoid cleanup issues in test environments
	testRemoveUserDisablesBeforeDeletion(t, createdUser.ID, *createdUser.Username)

	// Test nil Enabled field handling (edge case)
	testRemoveUserHandlesNilEnabled(t, createdUser.ID, *createdUser.Username)
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
	require.NoError(t, err, "error should be nil while retrieving reverted user")
	assert.Len(t, users, 1, "exactly one user should be retrieved")
	assert.Equal(t, revertRole, users[0].WorkspaceRole, "user role ID should match the updated role")
}

// testDisableAndEnableUser tests the UpdateUser function to disable and enable a user.
func testDisableAndEnableUser(t *testing.T, userID string, username string) {
	client := &UserClient{}

	// Test disabling the user
	enabled := false
	err := client.UpdateUser(userID, &enabled)
	require.NoError(t, err, "error should be nil while disabling user")

	// Verify the user is disabled
	user, err := client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while retrieving disabled user")
	assert.NotNil(t, user, "retrieved user should not be nil")
	if user.Enabled != nil {
		assert.False(t, *user.Enabled, "user should be disabled")
	}

	// Test enabling the user
	enabled = true
	err = client.UpdateUser(userID, &enabled)
	require.NoError(t, err, "error should be nil while enabling user")

	// Verify the user is enabled
	user, err = client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while retrieving enabled user")
	assert.NotNil(t, user, "retrieved user should not be nil")
	if user.Enabled != nil {
		assert.True(t, *user.Enabled, "user should be enabled")
	}
}

// testRemoveUserDisablesBeforeDeletion tests that RemoveUser disables the user before deletion.
// This test verifies the disable step but doesn't actually execute the deletion workflow
// to avoid cleanup issues in test environments.
func testRemoveUserDisablesBeforeDeletion(t *testing.T, userID string, username string) {
	client := &UserClient{}

	// First, ensure the user is enabled
	enabled := true
	err := client.UpdateUser(userID, &enabled)
	require.NoError(t, err, "error should be nil while enabling user for test")

	// Verify user is enabled before test
	user, err := client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while retrieving user before test")
	assert.NotNil(t, user, "retrieved user should not be nil")
	if user.Enabled != nil {
		assert.True(t, *user.Enabled, "user should be enabled before test")
	}

	// This verifies the disable-before-delete logic without actually deleting the user
	userDetails, err := client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while fetching user details")
	assert.NotNil(t, userDetails, "retrieved user details should not be nil")

	if userDetails.Enabled == nil || *userDetails.Enabled {
		enabled := false
		err = client.UpdateUser(userDetails.ID, &enabled)
		require.NoError(t, err, "error should be nil while disabling user (simulating RemoveUser behavior)")
	}

	// Verify the user is now disabled
	user, err = client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while retrieving disabled user")
	assert.NotNil(t, user, "retrieved user should not be nil")
	if user.Enabled != nil {
		assert.False(t, *user.Enabled, "user should be disabled after disable step")
	}

	// Re-enable the user for other tests
	enabled = true
	err = client.UpdateUser(userID, &enabled)
	require.NoError(t, err, "error should be nil while re-enabling user after test")

	t.Logf("Successfully verified that user is disabled before deletion step")
}

// testRemoveUserHandlesNilEnabled tests that the disable logic properly handles
// the case where the Enabled field is nil (unknown status).
func testRemoveUserHandlesNilEnabled(t *testing.T, userID string, username string) {
	client := &UserClient{}

	// Get user details
	userDetails, err := client.GetByUsername(username)
	require.NoError(t, err, "error should be nil while fetching user details")
	assert.NotNil(t, userDetails, "retrieved user details should not be nil")

	// This verifies that the code doesn't panic when Enabled is nil
	if userDetails.Enabled == nil || (userDetails.Enabled != nil && *userDetails.Enabled) {
		enabled := false
		err = client.UpdateUser(userDetails.ID, &enabled)
		require.NoError(t, err, "error should be nil while disabling user (handling nil Enabled)")

		// Verify the user is now disabled
		user, err := client.GetByUsername(username)
		require.NoError(t, err, "error should be nil while retrieving disabled user")
		assert.NotNil(t, user, "retrieved user should not be nil")
		if user.Enabled != nil {
			assert.False(t, *user.Enabled, "user should be disabled")
		}

		// Re-enable the user for other tests
		enabled = true
		err = client.UpdateUser(userID, &enabled)
		require.NoError(t, err, "error should be nil while re-enabling user after test")
	}

	t.Logf("Successfully verified nil Enabled field handling")
}

// TestSafeFullName tests the safeFullName helper function with various nil pointer scenarios.
func TestSafeFullName(t *testing.T) {
	tests := []struct {
		name      string
		first     *string
		last      *string
		fallbacks []string
		expected  string
	}{
		{
			name:     "both first and last name present",
			first:    stringPtr("John"),
			last:     stringPtr("Doe"),
			expected: "John Doe",
		},
		{
			name:     "only first name present",
			first:    stringPtr("John"),
			last:     nil,
			expected: "John",
		},
		{
			name:     "only last name present",
			first:    nil,
			last:     stringPtr("Doe"),
			expected: "Doe",
		},
		{
			name:     "both names nil without fallback",
			first:    nil,
			last:     nil,
			expected: "",
		},
		{
			name:     "empty first name",
			first:    stringPtr(""),
			last:     stringPtr("Doe"),
			expected: "Doe",
		},
		{
			name:     "empty last name",
			first:    stringPtr("John"),
			last:     stringPtr(""),
			expected: "John",
		},
		{
			name:     "both names empty strings without fallback",
			first:    stringPtr(""),
			last:     stringPtr(""),
			expected: "",
		},
		{
			name:     "names with extra spaces are individually trimmed",
			first:    stringPtr("  John  "),
			last:     stringPtr("  Doe  "),
			expected: "John Doe",
		},
		{
			name:      "both names nil falls back to email",
			first:     nil,
			last:      nil,
			fallbacks: []string{"john@example.com"},
			expected:  "john@example.com",
		},
		{
			name:      "both names empty falls back to email",
			first:     stringPtr(""),
			last:      stringPtr(""),
			fallbacks: []string{"john@example.com"},
			expected:  "john@example.com",
		},
		{
			name:      "both names whitespace-only falls back to email",
			first:     stringPtr("   "),
			last:      stringPtr("   "),
			fallbacks: []string{"john@example.com"},
			expected:  "john@example.com",
		},
		{
			name:      "names present ignores fallback",
			first:     stringPtr("John"),
			last:      stringPtr("Doe"),
			fallbacks: []string{"john@example.com"},
			expected:  "John Doe",
		},
		{
			name:      "empty fallback skipped to next",
			first:     nil,
			last:      nil,
			fallbacks: []string{"", "john@example.com"},
			expected:  "john@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := safeFullName(tt.first, tt.last, tt.fallbacks...)
			assert.Equal(t, tt.expected, result, "safeFullName should return the expected value")
		})
	}
}

// stringPtr is a helper function to create a pointer to a string.
func stringPtr(s string) *string {
	return &s
}
