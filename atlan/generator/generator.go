package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/assets"
	"github.com/atlanhq/atlan-go/atlan/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"
	"time"
)

var typeDefFile = "atlan/generator/typedefs.json"

// Define a template for the header (only needs to be executed once)
const headerTemplate = `// # **************************************
// # CODE BELOW IS GENERATED DO NOT MODIFY  **
// # **************************************

package generator
`

// Cache struct to hold different typedefs
type Cache struct {
	EnumDefCache         map[string]model.EnumDef
	StructDefCache       map[string]model.StructDef
	EntityDefCache       map[string]model.EntityDef
	RelationshipDefCache map[string]model.RelationshipDef
	mu                   sync.Mutex
}

func getTypeDefs() (*model.TypeDefResponse, error) {
	fileInfo, err := os.Stat(typeDefFile)
	if err != nil || time.Since(fileInfo.ModTime()) > 24*time.Hour {
		return nil, &assets.GenerationError{
			AtlanError: assets.AtlanError{
				ErrorCode: assets.ErrorInfo{},
				//				Args:          []interface{}{err.Error()},
				OriginalError: "File containing typedefs does not exist or is not current. Please run create_typedefs.go to create typedefs.json.",
			},
		}
	}

	fileContent, err := os.ReadFile(typeDefFile)
	if err != nil {
		return nil, err
	}
	var typeDefs model.TypeDefResponse
	if err := json.Unmarshal(fileContent, &typeDefs); err != nil {
		return nil, err
	}

	return &typeDefs, nil
}

// Create a new cache and populate it with typedef data from the JSON file
func NewCache() (*Cache, error) {
	cache := &Cache{
		EnumDefCache:         make(map[string]model.EnumDef),
		StructDefCache:       make(map[string]model.StructDef),
		EntityDefCache:       make(map[string]model.EntityDef),
		RelationshipDefCache: make(map[string]model.RelationshipDef),
	}

	err := cache.populateCache()
	if err != nil {
		return nil, err
	}
	return cache, nil
}

// Populate the cache with typedefs from the typedefs.json file
func (c *Cache) populateCache() error {
	typeDefs, err := getTypeDefs()
	if err != nil {
		return err
	}

	// Populate Enum Cache
	for _, enumDef := range typeDefs.EnumDefs {
		c.EnumDefCache[enumDef.Name] = enumDef
	}

	// Populate Struct Cache
	for _, structDef := range typeDefs.StructDefs {
		c.StructDefCache[structDef.Name] = structDef
	}

	// Populate Entity Cache
	for _, entityDef := range typeDefs.EntityDefs {
		c.EntityDefCache[entityDef.Name] = entityDef
	}

	// Populate Relationship Cache
	for _, relationshipDef := range typeDefs.RelationshipDefs {
		c.RelationshipDefCache[relationshipDef.Name] = relationshipDef
	}

	return nil
}

var enumTemplateFile = "atlan/generator/templates/enums.tmpl"
var structTemplateFile = "atlan/generator/templates/structs.tmpl"

// GenerateEnums will use typedefs and a template to generate the enums.go file
func GenerateEnums(enumDefs []model.EnumDef) error {
	// Define helper functions to be used in the template
	funcMap := template.FuncMap{
		"enumDef": func(def model.EnumDef) string {
			return def.Name // Access the name of the EnumDef
		},
		"elementDef": func(elem model.ElementDef) string {
			return elem.Value // Access the value of the ElementDef
		},
		"replace": func(input, old, new string) string {
			return strings.ReplaceAll(input, old, new)
		},
		"title": func(s string) string {
			return cases.Title(language.Und).String(s)
		},
		"contains": func(s, substr string) bool {
			return strings.Contains(s, substr) // Check if string contains substring
		},
	}

	// Read the template file and associate helper functions
	tmpl, err := template.New("enums.tmpl").Funcs(funcMap).ParseFiles(enumTemplateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create a buffer to store the output
	var out bytes.Buffer

	// Execute the template with the enumDefs slice as the data
	err = tmpl.Execute(&out, enumDefs)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Write the output to the generated file (enums.go)
	err = os.WriteFile("atlan/generator/enums.go", out.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write enums.go file: %w", err)
	}

	fmt.Println("Successfully generated enums.go")
	return nil
}

func GenerateStructs(structDefs []model.StructDef) error {
	// Define helper functions for the template
	funcMap := template.FuncMap{
		"title": func(t string) string { return cases.Title(language.Und).String(t) },
		"lower": strings.ToLower,
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"typeAsString": func(t interface{}) string {
			// Check if t is a pointer and dereference it
			if ptr, ok := t.(*string); ok && ptr != nil {
				typeName := *ptr // Dereference the pointer to get the type name as a string
				switch typeName {
				case "string":
					return "*string"
				case "int":
					return "*int"
				case "long":
					return "*int64"
				case "float":
					return "*float64"
				case "date": // Assuming date maps to time.Time
					return "*time.Time"
				case "SourceCostUnitType": // Custom type
					return "*SourceCostUnitType"
				case "boolean":
					return "*bool"
				case "array<string>":
					return "[]*string"
				case "array<float>":
					return "[]*float64"
				case "map<string,string>":
					return "map[string]*string"
				case "array<SourceTagAttachmentValue>":
					return "[]*SourceTagAttachmentValue"
				// Add more cases as necessary for other types
				default:
					fmt.Printf("Debug - Found unknown type: %s\n", typeName) // Debugging statement
					return fmt.Sprintf("*%s", typeName)                      // Return a pointer type for unknown types
				}
			}
			// Handle non-pointer types or nil cases
			fmt.Printf("Debug - Found non-pointer or nil type: %T\n", t) // Debugging statement
			return fmt.Sprintf("%v", t)
		},
	}

	var out bytes.Buffer

	// Debugging statement to check the struct definitions passed in
	// fmt.Println("Debug - Struct definitions passed in:", structDefs)

	// Execute the header template once
	tmplHeader, err := template.New("header").Parse(headerTemplate)
	if err != nil {
		fmt.Println("Error parsing header template:", err)

	}

	err = tmplHeader.Execute(&out, nil)
	if err != nil {
		fmt.Println("Error executing header template:", err)

	}

	// Parse the template file
	tmpl, err := template.New("structs.tmpl").Funcs(funcMap).ParseFiles(structTemplateFile)
	if err != nil {
		return fmt.Errorf("failed to parse structs template: %w", err)
	}

	// Execute the template for each struct
	for _, structDef := range structDefs {
		err = tmpl.Execute(&out, structDef)
		if err != nil {
			return fmt.Errorf("failed to execute structs template: %w", err)
		}
	}

	// Write the generated content to structs.go file
	err = os.WriteFile("atlan/generator/structs.go", out.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write structs.go file: %w", err)
	}

	fmt.Println("Successfully generated structs.go")
	return nil
}

func RunGenerator() {
	//if err := GenerateTypedefsFile(); err != nil {
	//	log.Fatalf("Error generating typedefs file: %v", err)
	//}

	// Initialize and populate cache from typedefs.json
	cache, err := NewCache()
	if err != nil {
		log.Fatalf("Failed to load typedefs: %v", err)
	}

	// Convert cache.EnumDefCache (map[string]model.EnumDef) to a slice ([]model.EnumDef)
	enumDefsSlice := make([]model.EnumDef, 0, len(cache.EnumDefCache))
	for _, enumDef := range cache.EnumDefCache {
		enumDefsSlice = append(enumDefsSlice, enumDef)
	}
	//fmt.Printf("Loaded EnumDefs: %+v\n", cache.EnumDefCache)

	// Generate enums using the typedefs
	err = GenerateEnums(enumDefsSlice)
	if err != nil {
		log.Fatalf("Failed to generate enums: %v", err)
	}

	structDefsSlice := make([]model.StructDef, 0, len(cache.StructDefCache))
	for _, structDef := range cache.StructDefCache {
		structDefsSlice = append(structDefsSlice, structDef)
	}

	err = GenerateStructs(structDefsSlice)
	if err != nil {
		log.Fatalf("Failed to generate structs: %v", err)
	}
	/*
		// Initialize generator and create modules
		generator, err := NewGenerator()
		if err != nil {
			fmt.Println(err)
			return
		}

		for name, asset := range generator.templates.Assets {
			fmt.Printf("Asset: %s, Order: %d, IsCore: %t\n", name, asset.Order, asset.IsCoreAsset)
		}

	*/
	//UpdateSubTypeNamesToIgnore()
	//fmt.Println(typeDefs)

}
