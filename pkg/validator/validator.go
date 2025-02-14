package validator

import (
	"errors"
	"github.com/jeremyalv/go-todo-api/constants"
	"reflect"
	"strings"
)

const tagName = "validate"

type Validator interface {
	Validate(any) (bool, error)
}

type DefaultValidator struct{}

func (v DefaultValidator) Validate(val any) (bool, error) {
	return true, nil
}

type TodoValidator struct{}

func (v TodoValidator) Validate(val any) (bool, error) {
	if !constants.RegexTodoId.MatchString(val.(string)) {
		return false, errors.New("invalid TodoId given")
	}
	return true, nil
}

func FindTag(tag string) Validator {
	args := strings.Split(tag, ",")
	switch args[0] {
	case constants.CtxTodoId:
		return TodoValidator{}
	default:
		return DefaultValidator{}
	}
}

func ValidateRequest(s any) *[]error {
	var errorItems []error

	v := reflect.ValueOf(s)
	for i := 0; i < reflect.Indirect(v).NumField(); i++ {
		// Returns the value of the `validate` tag from the i-th field in the passed in struct s.
		tag := reflect.Indirect(v).Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}
		validator := FindTag(tag)

		// Converts the value of the i-th field of struct s back to an interface{} / any, and pass it to our validator's Validate() method
		valid, err := validator.Validate(reflect.Indirect(v).Field(i).Interface())
		if !valid && err != nil {
			errorItems = append(errorItems, err)
		}
	}

	if len(errorItems) > 0 {
		return &errorItems
	}
	return nil
}
