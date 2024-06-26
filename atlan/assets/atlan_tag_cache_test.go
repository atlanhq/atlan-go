package assets

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.NoError(t, err)
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

	//_ = cache.RefreshCache()

	fmt.Println("Tag Name: ", tagName)
	id, err := cache.GetIDForName(tagName)

	// Verify
	assert.NoError(t, err)
	assert.NotEmpty(t, id) // ID should be non-empty if the tag exists

	// Test not found scenario
	_, err = cache.GetIDForName("NonExistentTag")
	assert.Nil(t, nil) // Expect error(nil) since tag does not exist
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

	// Assuming "BBDjIBZUNHtKPExR1Z3a5I" is a valid GUID
	name, err := cache.GetNameForID(id)

	// Verify
	assert.NoError(t, err)
	assert.NotEmpty(t, name) // Name should be non-empty if the ID is valid

	// Test not found scenario
	_, err = cache.GetNameForID("123456")
	assert.Nil(t, nil) // Expect error(nil) since ID does not exist
}
