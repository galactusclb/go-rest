package validation

import (
	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

func NewValidator() *XValidator {
	return &XValidator{
		validator: validator.New(),
	}
}

func (v *XValidator) Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	err := v.validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var elm ErrorResponse

			elm.Error = true
			elm.FailedField = err.Field()
			elm.Tag = err.Tag()
			elm.Value = err.Value()

			validationErrors = append(validationErrors, elm)
		}
	}

	return validationErrors
}
