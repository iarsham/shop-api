package common

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const IrPhoneRegex string = `^(\+98|0)?9\d{9}$`

func IrPhoneValidate(phone string) bool {
	ok, err := regexp.MatchString(IrPhoneRegex, phone)
	if err != nil {
		return false
	}
	return ok
}

func IrPhoneValidator(v validator.FieldLevel) bool {
	value, ok := v.Field().Interface().(string)
	if !ok {
		return false
	}
	return IrPhoneValidate(value)
}
