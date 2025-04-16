package requests

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type SendMessageRequest struct {
	To      string `json:"to" validate:"required,e164,min=11,max=15"`
	Message string `json:"message" validate:"required"`
}

var validate = validator.New()

func (r *SendMessageRequest) Validate() map[string]string {
	err := validate.Struct(r)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		switch err.Tag() {
		case "required":
			errors[field] = fmt.Sprintf("%s is required", field)
		case "e164":
			errors[field] = fmt.Sprintf("%s must be a valid international phone number (E.164 format)", field)
		case "min":
			errors[field] = fmt.Sprintf("%s must be at least %s characters", field, err.Param())
		case "max":
			errors[field] = fmt.Sprintf("%s must be at most %s characters", field, err.Param())
		default:
			errors[field] = fmt.Sprintf("Invalid value for %s", field)
		}
	}

	return errors
}
