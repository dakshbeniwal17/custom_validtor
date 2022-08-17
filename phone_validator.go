package custom_validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const phoneRegex = `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`

func ValidatePhone(fl validator.FieldLevel) bool {
	phoneRegex := regexp.MustCompile(phoneRegex)
	return phoneRegex.MatchString(fl.Field().String())
}
