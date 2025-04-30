package strrand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomID(t *testing.T) {
	ids := make(map[string]any)

	for range 1000000 {
		id := RandomID()

		_, exists := ids[id]
		assert.False(t, exists, "RandomID() = %v, already exists", id)

		ids[id] = nil
	}
}

func TestRandomTimeAlphaNumeric(t *testing.T) {
	length := 16
	result := randomTimeAlphaNumeric(alphaNumeric, length)

	assert.Equal(t, length, len(result), "Expected length %d, got %d", length, len(result))

	for _, char := range result {
		assert.Contains(t, alphaNumeric, string(char), "Unexpected character %c in result", char)
	}
}

func TestRandomAlphaNumeric(t *testing.T) {
	dictionary := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	length := 10
	result := randomAlphaNumeric(dictionary, length)

	assert.Equal(t, length, len(result), "Expected length %d, got %d", length, len(result))

	for _, char := range result {
		assert.Contains(t, dictionary, string(char), "Unexpected character %c in result", char)
	}
}
