package dto

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	v10 "github.com/go-playground/validator/v10"
)

func SetupValidation() {
	if v, ok := binding.Validator.Engine().(*v10.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})
	} else {
		slog.Error("kit/dto/validate: unable to get valid validator instance")
	}
}

func Validate[T any](c *gin.Context, container *T) (ok bool, validationErrors []ValidationError, err error) {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		c.Request.Body = io.NopCloser(strings.NewReader("{}"))
	}

	err = c.ShouldBindJSON(container)
	if err != nil {
		ok = false
		re := make(map[string]ValidationError)

		var ve v10.ValidationErrors

		as := errors.As(err, &ve)
		if !as {
			return ok, validationErrors, err
		}

		err = nil

		for _, e := range ve {
			if eve, ok := re[e.Field()]; !ok {
				re[e.Field()] = ValidationError{
					Field:  e.Field(),
					Errors: []string{translate(e)},
				}
			} else {
				eve.Errors = append(eve.Errors, translate(e))
			}
		}

		for _, ve := range re {
			validationErrors = append(validationErrors, ve)
		}

		return ok, validationErrors, nil
	}

	ok = true

	return ok, validationErrors, nil
}

func translate(error v10.FieldError) string {
	switch error.Tag() {
	case "required":
		return "required"

	case "email":
		return "must be a valid email address"

	case "numeric":
		return "must be a valid number"

	case "len":
		if error.Param() == "1" {
			return fmt.Sprintf("must be %s character long", error.Param())
		} else {
			return fmt.Sprintf("must be %s characters long", error.Param())
		}

	case "gt":
		return fmt.Sprintf("must be a number greater than %s", error.Param())

	case "gte":
		return fmt.Sprintf("must be a number greater than or equal to %s", error.Param())

	case "lt":
		return fmt.Sprintf("must be a number less than %s", error.Param())

	case "lte":
		return fmt.Sprintf("must be a number less than or equal to %s", error.Param())

	case "min":
		return fmt.Sprintf("must be at least %s", error.Param())

	case "max":
		return fmt.Sprintf("must be less than %s", error.Param())

	case "oneof":
		return fmt.Sprintf("must be one of: %s", error.Param())

	case "url":
		return "must be a valid URL"

	case "e164":
		return "must be a properly formatted phone number"
	}

	return ""
}
