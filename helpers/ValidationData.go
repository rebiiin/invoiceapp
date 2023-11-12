package helpers

import (
	"github.com/go-playground/validator/v10"
)

type IError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	//Value string `json:"value"`
}

func InputValidation(input interface{}) []*IError {
	v := validator.New()
	errorValidation := v.Struct(input)

	var errors []*IError

	if errorValidation != nil {
		for _, err := range errorValidation.(validator.ValidationErrors) {
			el := &IError{
				Field: err.Field(),
				Tag:   err.Tag(),
				//Value: err.Param(),
			}
			errors = append(errors, el)
		}
	}

	return errors
}
