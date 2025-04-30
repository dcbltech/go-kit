package dto

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateOrRespond(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	type Payload struct {
		Name string `json:"name" binding:"required"`
	}

	c.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	c.Request.Header.Set("Content-Type", "application/json")

	var payload Payload

	ok := ValidateOrRespond(c, &payload)

	assert.False(t, ok)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "required")
}

func TestResponseHelpers(t *testing.T) {
	assert.Equal(t, http.StatusCreated, Created().Status)
	assert.Equal(t, http.StatusOK, OK().Status)
	assert.Equal(t, http.StatusOK, OKWithData("data").Status)
	assert.Equal(t, http.StatusFound, Redirect("/url").Status)
	assert.Equal(t, http.StatusBadRequest, BadRequest().Status)
	assert.Equal(t, http.StatusBadRequest, BadRequestWithData("data").Status)
	assert.Equal(t, http.StatusPaymentRequired, PaymentRequired().Status)
	assert.Equal(t, http.StatusPaymentRequired, PaymentRequiredWithData("data").Status)
	assert.Equal(t, http.StatusUnauthorized, Unauthorized().Status)
	assert.Equal(t, http.StatusForbidden, Forbidden().Status)
	assert.Equal(t, http.StatusNotFound, NotFound().Status)
	assert.Equal(t, http.StatusUnprocessableEntity, UnprocessableEntityWithError("error").Status)
	assert.Equal(t, http.StatusLocked, Locked().Status)
	assert.Equal(t, http.StatusInternalServerError, InternalServerError().Status)
}
