package helper

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

var TtValidation validator.Func = func(fl validator.FieldLevel) bool {
	//date, ok := fl.Field().Interface().(time.Time)
	//if ok {
	//}
	return true
}
