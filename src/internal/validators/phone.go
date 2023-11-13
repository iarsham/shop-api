package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/iarsham/shop-api/internal/common"
	"regexp"
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

func RegisterValidators(logs *common.Logger) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("phone", IrPhoneValidator, true)
		common.LogError(logs, err)
	}
}
