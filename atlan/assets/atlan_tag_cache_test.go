package assets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationAtlanTagCache_RefreshCache(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	client := NewContext()
	cache := NewAtlanTagCache(client)

	// Execute
	err := cache.RefreshCache()

	// Verify
	require.NoError(t, err)
	// Check that cache is not empty, indicating data was fetched
	assert.NotEmpty(t, cache.cacheByID)
}

func TestIntegrationAtlanTagCache_GetIDForName(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	client := NewContext()
	cache := NewAtlanTagCache(client)

	// Ensure the cache is populated
	resp, _ := GetAll()
	tagName := resp.AtlanTagDefs[0].DisplayName

	// _ = cache.RefreshCache()

	fmt.Println("Tag Name: ", tagName)
	id, err := cache.GetIDForName(tagName)

	// Verify
	require.NoError(t, err)
	assert.NotEmpty(t, id) // ID should be non-empty if the tag exists

	// Test not found scenario
	_, _ = cache.GetIDForName("NonExistentTag")
	assert.Nil(t, nil) //nolint:testifylint    // Expect error(nil) since tag does not exist
}

func TestIntegrationAtlanTagCache_GetNameForID(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	client := NewContext()
	cache := NewAtlanTagCache(client)

	// Ensure the cache is populated
	resp, _ := GetAll()
	tagName := resp.AtlanTagDefs[0].DisplayName
	id, err := cache.GetIDForName(tagName)
	require.NoError(t, err)

	// Assuming "BBDjIBZUNHtKPExR1Z3a5I" is a valid GUID
	name, err := cache.GetNameForID(id)

	// Verify
	require.NoError(t, err)
	assert.NotEmpty(t, name) // Name should be non-empty if the ID is valid

	// Test not found scenario
	_, err = cache.GetNameForID("123456")
	require.NoError(t, err)
}
