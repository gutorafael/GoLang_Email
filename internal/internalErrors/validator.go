package internalerrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct (obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)
		validationError := validationErr[0]
		switch validationError.Tag() {
			case "required":
				return errors.New(validationError.StructField() + " cannot be empty")
			case "min":
				return errors.New(validationError.StructField() + " cannot be less than " + validationError.Param() + " characters")
			case "max":
				return errors.New(validationError.StructField() + " cannot be greater than " + validationError.Param() + " characters")
			case "email":
				return errors.New(validationError.StructField() + " must be a valid email")
			case "gte":
				return errors.New(validationError.StructField() + " must have at least one contact")
		}
	}
	return nil
}