package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	LenError       = errors.New("LenError")
	MinError       = errors.New("MinError")
	MaxError       = errors.New("MaxError")
	InError        = errors.New("InError")
	NotStructError = errors.New("NotStructError")
	WrongRuleError = errors.New("WrongRuleError")
	RegexpError    = errors.New("RegexpError")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	buff := strings.Builder{}
	for _, e := range v {
		buff.WriteString(fmt.Sprintf("Field: %s, Error: %v\n", e.Field, e.Err))
	}
	return buff.String()
}

func Validate(v interface{}) error {

	var vErrors = ValidationErrors{}
	e := reflect.ValueOf(v)
	if e.Kind() != reflect.Struct {
		return NotStructError
	}

	for i := 0; i < e.NumField(); i++ {

		field := e.Type().Field(i)
		varName := field.Name
		varTag := field.Tag
		if varTag == "" {
			continue
		}

		varValue := e.Field(i).Interface()

		if e.Field(i).Kind() == reflect.Slice {
			for _, sliceVal := range varValue.([]string) {
				err := validateValue(varTag, sliceVal)
				if err != nil {
					if !isValidationError(&err) {
						return err
					}
					vErrors = append(vErrors, ValidationError{Field: varName, Err: err})
				}
			}

		} else {
			err := validateValue(varTag, varValue)
			if err != nil {
				if !isValidationError(&err) {
					return err
				}
				vErrors = append(vErrors, ValidationError{Field: varName, Err: err})
			}
		}

	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}

func isValidationError(err *error) bool {
	if errors.Is(*err, LenError) ||
		errors.Is(*err, MinError) ||
		errors.Is(*err, MaxError) ||
		errors.Is(*err, InError) ||
		errors.Is(*err, NotStructError) ||
		errors.Is(*err, WrongRuleError) ||
		errors.Is(*err, RegexpError) {
		return true
	}
	return false
}

func validateValue(varTag reflect.StructTag, varValue interface{}) error {

	if varTagVal := varTag.Get("validate"); varTagVal != "" {
		valRules := strings.Split(varTagVal, "|")

		for _, rawRule := range valRules {
			valRule := strings.Split(rawRule, ":")

			if len(valRule) != 2 {
				return WrongRuleError
				//return ValidationError{Field: varName, Err: WrongRuleError}, nil
			}

			rule := valRule[0]
			val := valRule[1]

			switch rule {
			case "len":

				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}

				if intVal != len(varValue.(string)) {
					return LenError
				}

			case "min":
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}

				if intVal > varValue.(int) {
					return MinError
				}

			case "max":
				intVal, err := strconv.Atoi(val)
				if err != nil {
					return err
				}

				if intVal < varValue.(int) {
					return MaxError
				}

			case "in":
				inStr := strings.Split(val, ",")

				if false == contains(inStr, fmt.Sprintf("%v", varValue)) {
					return InError
				}
			case "regexp":
				matched, err := regexp.MatchString(val, fmt.Sprintf("%v", varValue))
				if err != nil {
					return err
				}

				if !matched {
					return RegexpError
				}
			default:
			}

		}

	}
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
