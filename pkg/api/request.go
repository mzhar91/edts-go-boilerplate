package api

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	
	"github.com/go-playground/validator"
	
	"sg-edts.com/edts-go-boilerplate/pkg/message"
)

// IsRequestValid check is request valid
func IsRequestValid(m interface{}) (bool, error) {
	validate := validator.New()
	_ = validate.RegisterValidation("email", validateEmail)
	
	validate.RegisterTagNameFunc(
		func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			
			return name
		},
	)
	
	err := validate.Struct(m)
	
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				return false, FromErrorCodeWithParam(message.RequiredField, []interface{}{err.Field()})
			case "email":
				return false, FromErrorCode(message.InvalidEmailFormat)
			case "len":
				return false, FromErrorCode(message.IncorrectFormat)
			case "max":
				return false, FromErrorCodeWithParam(
					message.ExceedMaxCharacter, []interface{}{err.Field(), err.Param()},
				)
			case "gte":
				return false, FromErrorCodeWithParam(message.RequiredField, []interface{}{err.Field()})
			default:
				return false, errors.New("validation error")
			}
		}
	}
	
	return true, nil
}

func validateEmail(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	field := strings.ToLower(fl.Field().String())
	isEmpty := field == ""
	
	if regex.MatchString(field) && !isEmpty {
		return true
	}
	
	if isEmpty {
		return true
	}
	
	return false
}
