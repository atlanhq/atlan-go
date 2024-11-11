package generator

import (
	"encoding/json"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/assets"
	"log"
	"os"
	"path/filepath"
)

// GenerateTypedefsFile generates the typedefs.json file in the generator package.
func GenerateTypedefsFile() error {
	// Fetch typedefs using the existing GetAll() function from the assets package
	typedefs, err := assets.GetAll()
	if err != nil {
		return fmt.Errorf("failed to fetch typedefs: %v", err)
	}

	// Define the file path for typedefs.json in the generator folder
	typedefsFilePath := filepath.Join("atlan", "generator", "typedefs.json")

	// Create the typedefs.json file
	file, err := os.Create(typedefsFilePath)
	if err != nil {
		return fmt.Errorf("failed to create typedefs.json: %v", err)
	}
	defer file.Close()

	// Write the typedefs to the file in JSON format
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Add indentation for better readability
	if err := encoder.Encode(typedefs); err != nil {
		return fmt.Errorf("failed to write typedefs to file: %v", err)
	}

	fmt.Println("typedefs.json successfully generated at", typedefsFilePath)
	return nil
}

// Main function to trigger the typedefs generation
func main() {
	if err := GenerateTypedefsFile(); err != nil {
		log.Fatalf("Error generating typedefs file: %v", err)
	}
}
