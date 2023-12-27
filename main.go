package main

import (
	"email/internal/domain/campaign"
	"github.com/go-playground/validator/v10"
)

func main () {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)
		for _, v:= range validationErr {
			
			switch v.Tag() {
				case "required":
					println(v.StructField() + " cannot be empty")
				case "min":
					println(v.StructField() + " cannot be less than " + v.Param() + " characters")
				case "max":
					println(v.StructField() + " cannot be greater than " + v.Param() + " characters")
				case "email":
					println(v.StructField() + " must be a valid email")
			}
		}
	}
}