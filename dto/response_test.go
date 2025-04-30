package dto

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponse_Respond(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	resp := &Response{
		Type:   responseTypeRespond,
		Status: http.StatusOK,
		Data:   gin.H{"message": "success"},
	}

	resp.Respond(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "success")
}

func TestResponse_Redirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodGet, "http://original.com", nil)

	resp := &Response{
		Type:   responseTypeRedirect,
		Status: http.StatusFound,
		Data:   "http://redirect.com",
	}

	resp.Respond(c)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "http://redirect.com", w.Header().Get("Location"))
}
