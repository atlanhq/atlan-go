package client

import (
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTypeDefinitionsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	NewContext()

	// Example categories to test
	categories := []atlan.AtlanTypeCategory{
		atlan.AtlanTypeCategoryClassification,
	}

	// Test getting type definitions
	response, err := Get(categories)
	assert.NoError(t, err, "Expected no error from Get")
	assert.NotNil(t, response, "Expected a valid response from Get")

	// Validate the response structure
	assert.NotNil(t, response.StructDefs, "Expected StructDefs to be non-nil")
	assert.NotEmpty(t, response.StructDefs, "Expected StructDefs to be non-empty")
	assert.NotNil(t, response.AtlanTagDefs, "Expected AtlanTagDefs to be non-nil")
	assert.NotEmpty(t, response.AtlanTagDefs, "Expected AtlanTagDefs to be non-empty")

}
