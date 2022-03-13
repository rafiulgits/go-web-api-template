package models

import (
	"fmt"
	"net/mail"
	"reflect"
	"strings"

	lib "github.com/go-playground/validator/v10"
)

var validator *lib.Validate

func NewValidator() *lib.Validate {
	v := lib.New()
	v.RegisterValidation("optional_email", optionalEmailValidation)
	v.RegisterTagNameFunc(configureFieldNameConvention)
	validator = v
	return v
}

func GetValidator() *lib.Validate {
	return validator
}

func configureFieldNameConvention(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

// @param err must be `ValidationErrors` type
func ParseModelErrors(err error) []map[string]string {
	validationErr, ok := err.(lib.ValidationErrors)
	if !ok {
		return nil
	}
	fieldErrs := make([]map[string]string, 0)
	for _, e := range validationErr {
		var msg string
		switch e.Tag() {
		case requireFailed:
			msg = "This field is required"
		case maxLengthFailed:
			msg = fmt.Sprintf("Maximum %s characters allowed", e.Param())
		case minLengthFailed:
			msg = fmt.Sprintf("Minimum %s characters required", e.Param())
		case emailFailed, optionalEmailFailed:
			msg = "Invalid email address"
		case lengthFailed:
			msg = fmt.Sprintf("Length should be %s", e.Param())
		default:
			msg = e.Error()
		}
		fieldErrs = append(fieldErrs, map[string]string{e.Field(): msg})
	}
	return fieldErrs
}

const requireFailed = "required"
const maxLengthFailed = "max"
const minLengthFailed = "min"
const lengthFailed = "len"
const emailFailed = "email"
const optionalEmailFailed = "optional_email"

func optionalEmailValidation(field lib.FieldLevel) bool {
	if field.Field().Kind() != reflect.String {
		return false
	}
	if field.Field().Len() == 0 {
		return true
	}
	_, err := mail.ParseAddress(field.Field().String())
	return err == nil
}
