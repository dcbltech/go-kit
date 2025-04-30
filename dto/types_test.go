package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError(t *testing.T) {
	err := ValidationError{
		Field:  "name",
		Errors: []string{"required"},
	}

	assert.Equal(t, "name", err.Field)
	assert.Contains(t, err.Errors, "required")
}

func TestValidationFailedResponse(t *testing.T) {
	resp := ValidationFailedResponse{
		Errors: []ValidationError{
			{
				Field:  "name",
				Errors: []string{"required"},
			},
		},
	}

	assert.Len(t, resp.Errors, 1)
	assert.Equal(t, "name", resp.Errors[0].Field)
}

func TestErrorResponse(t *testing.T) {
	err := ErrorResponse{
		Message: "unable to find users using id",
	}

	assert.Equal(t, "unable to find users using id", err.Message)
}
