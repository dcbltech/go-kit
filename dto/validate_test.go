package dto

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	SetupValidation()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	type Payload struct {
		Name string `json:"name" binding:"required"`
	}

	c.Request = httptest.NewRequest(http.MethodPost, "http://test.com", nil)

	var payload Payload

	ok, validationErrors, err := Validate(c, &payload)

	assert.False(t, ok)
	assert.NotNil(t, validationErrors)
	assert.Nil(t, err)
}
