package validator

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

// SetUpValidation sets up the custom validations for the validator
func SetUpValidation() {
	customValidator := validator.New()
	Validate = customValidator
}
