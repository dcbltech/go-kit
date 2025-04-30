package dto

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidateOrRespond validates the incoming request data against the defined
// contract. In the case of an error, or validation failing, it also fulfils
// the request to the client.
func ValidateOrRespond[T any](c *gin.Context, payload *T) (ok bool) {
	ok, validationErrors, err := Validate(c, payload)
	if err != nil {
		slog.Error("kit/dto/respond: something went wrong", "error", err)

		InternalServerError().Respond(c)

		return
	}

	if !ok {
		ValidationFailed(validationErrors).Respond(c)

		return
	}

	return
}

func Created() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusCreated,
		Data:   nil,
	}
}

func OK() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusOK,
		Data:   nil,
	}
}

func OKWithData(data any) *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusOK,
		Data:   data,
	}
}

func Redirect(url string) *Response {
	return &Response{
		Type:   responseTypeRedirect,
		Status: http.StatusFound,
		Data:   url,
	}
}

func BadRequest() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusBadRequest,
		Data:   nil,
	}
}

func BadRequestWithData(data any) *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusBadRequest,
		Data:   data,
	}
}

func PaymentRequired() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusPaymentRequired,
		Data:   nil,
	}
}

func PaymentRequiredWithData(data any) *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusPaymentRequired,
		Data:   data,
	}
}

func ValidationFailed(errors []ValidationError) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Data: &ValidationFailedResponse{
			Errors: errors,
		},
	}
}

func ValidationFailedWith(field, error string) *Response {
	return ValidationFailed([]ValidationError{
		{
			Field:  field,
			Errors: []string{error},
		},
	})
}

func Unauthorized() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusUnauthorized,
		Data:   nil,
	}
}

func Forbidden() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusForbidden,
		Data:   nil,
	}
}

func NotFound() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusNotFound,
		Data:   nil,
	}
}

func UnprocessableEntityWithError(error string) *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusUnprocessableEntity,
		Data:   ErrorResponse{Message: error},
	}
}

func Locked() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusLocked,
		Data:   nil,
	}
}

func InternalServerError() *Response {
	return &Response{
		Type:   responseTypeRespond,
		Status: http.StatusInternalServerError,
		Data:   nil,
	}
}
