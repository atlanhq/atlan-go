package atlan

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid"
)

// TestID represents a test identifier
type TestID struct {
	sessionID string
}

// NewTestID generates a new TestID
func NewTestID() *TestID {
	sessionID := generateNanoid(5)
	return &TestID{sessionID: sessionID}
}

// GenerateNanoid generates a random string of given length
func generateNanoid(size int) string {
	const alphabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, err := gonanoid.Generate(alphabet, size)
	if err != nil {
		panic(err)
	}
	return id
}

// MakeUnique creates a unique identifier using the input string and session ID
func (t *TestID) MakeUnique(input string) string {
	return fmt.Sprintf("gsdk_%s_%s", input, t.sessionID)
}
