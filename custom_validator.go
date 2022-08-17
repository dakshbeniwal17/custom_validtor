package custom_validator

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

type structValidator interface {
	validateStruct(any) error
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var CustomValidator structValidator = &defaultValidator{}
var validate *validator.Validate

func errorForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "is-email":
		return fmt.Sprintf("`%s` is not a valid email", fe.Value())
	case "is-phone":
		return fmt.Sprintf("`%s` is not a valid phone number", fe.Value())
	case "regex":
		return fmt.Sprintf("`%s` does not match the given regex: %s", fe.Value(), fe.Param())
	}
	return fe.Error()
}

func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("my-validator")
		v.validate.RegisterValidation("regex", ValidateRegex)
		v.validate.RegisterValidation("is-email", ValidateEmail)
		v.validate.RegisterValidation("is-phone", ValidatePhone)
	})
}

func ValidateMyStruct(x interface{}) map[string]string {
	err := CustomValidator.validateStruct(x)
	if err != nil {
		var ve validator.ValidationErrors
		errorMessages := make(map[string]string)
		if errors.As(err, &ve) {
			for _, fe := range ve {
				errorMessages[fe.Field()] = errorForTag(fe)
			}
		}
		return errorMessages
	}
	return map[string]string{}
}
