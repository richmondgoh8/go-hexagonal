package handlers

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterValidation("urlvalid", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6 && strings.HasPrefix(fl.Field().String(), "http")
	})
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "urlvalid":
		return "Invalid url"
	}
	return fe.Error() // default error
}
