package atlan

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid"
	"time"
)

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
func MakeUnique(input string) string {
	return fmt.Sprintf("gsdk_%s_%s", input, generateNanoid(5))
}

var sNextID = int64(time.Now().UnixNano()/int64(time.Millisecond)) + 1

func NextID() string {
	sNextID++
	return fmt.Sprintf("-%d", sNextID)
}