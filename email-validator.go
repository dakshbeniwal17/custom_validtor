package custom_validator

import (
	"net/mail"

	"github.com/go-playground/validator/v10"
)

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func ValidateEmail(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	return isEmail(data)
}
