package custom_validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateRegex(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	regex := fl.Param()
	return matchRegex(data, regexp.MustCompile(regex))
}

func matchRegex(data string, regex *regexp.Regexp) bool {
	return regex.MatchString(data)
}
