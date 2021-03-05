package utils

import "github.com/go-playground/validator/v10"

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = "field " + err.StructField() + " is not valid"
	}

	return fields
}
