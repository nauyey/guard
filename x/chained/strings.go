package chained

import "github.com/nauyey/guard"

func ValidateString(value string) *stringValidator {
	return &stringValidator{
		value:          value,
		valueValidator: &valueValidator{},
	}
}

type stringValueValidator interface {
	SetValue(value string)
	guard.Validator
}

type stringValidator struct {
	value string
	*valueValidator
}

func (v *stringValidator) With(validator stringValueValidator) *stringValidator {
	validator.SetValue(v.value)
	v.validators = append(v.validators, guard.Strict(validator))
	return v
}
