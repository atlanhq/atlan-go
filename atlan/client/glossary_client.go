package client

import (
	"encoding/json"
	"errors"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/assets"
	"time"
)

const (
	MaxRetries    = 5
	RetryInterval = time.Second * 5
)

type AtlasGlossary assets.AtlasGlossary

// GlossaryClient defines the client for interacting with the model API.
type GlossaryClient struct {
	client *AtlanClient
}

// NewGlossaryClient creates a new GlossaryClient instance.
func NewGlossaryClient(ac *AtlanClient) *GlossaryClient {
	return &GlossaryClient{client: ac}
}

// Creator is used to create a new glossary asset in memory.
func (g *AtlasGlossary) Creator(name string, icon atlan.AtlanIcon) {
	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.QualifiedName = assets.StringPtr(name)
	g.AssetIcon = atlan.AtlanIconPtr(icon)
}

// Updater is used to modify a glossary asset in memory.
func (g *AtlasGlossary) Updater(name string, qualifiedName string, glossary_guid string) error {
	if name == "" || qualifiedName == "" || glossary_guid == "" {
		return errors.New("name, qualified_name, and glossary_guid are required fields")
	}

	g.TypeName = assets.StringPtr("AtlasGlossary")
	g.Name = assets.StringPtr(name)
	g.Guid = assets.StringPtr(glossary_guid)
	g.QualifiedName = assets.StringPtr(qualifiedName)

	return nil
}

// MarshalJSON filters out entities to only include those with non-empty attributes.
func (g *AtlasGlossary) MarshalJSON() ([]byte, error) {
	// Construct the custom JSON structure
	customJSON := map[string]interface{}{
		"typeName": "AtlasGlossary",
		"attributes": map[string]interface{}{
			"name": g.Name,
			// Add other attributes as necessary.
		},
		"relationshipAttributes": make(map[string]interface{}),
	}

	if g.QualifiedName != nil && *g.QualifiedName != "" {
		customJSON["attributes"].(map[string]interface{})["qualifiedName"] = *g.QualifiedName
	}

	if g.Guid != nil && *g.Guid != "" {
		customJSON["guid"] = *g.Guid
	}

	if g.DisplayName != nil && *g.DisplayName != "" {
		customJSON["attributes"].(map[string]interface{})["displayName"] = *g.DisplayName
	}

	// Marshal the custom JSON
	return json.MarshalIndent(customJSON, "", "  ")
}
