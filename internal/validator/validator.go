package validator

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validator interface {
	Validate(interface{}) (string, error)
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (v *CustomValidator) Validate(i interface{}) (string, error) {

	var errors []string
	err := v.validator.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is required")
		}
	}

	return strings.Join(errors, ", "), err
}
