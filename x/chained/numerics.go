package chained

import "github.com/nauyey/guard"

func ValidateInt(value int) *intValidator {
	return &intValidator{
		value:          value,
		valueValidator: &valueValidator{},
	}
}

type intValueValidator interface {
	SetValue(value int)
	guard.Validator
}

type intValidator struct {
	value int
	*valueValidator
}

func (v *intValidator) With(validator intValueValidator) *intValidator {
	validator.SetValue(v.value)
	v.validators = append(v.validators, guard.Strict(validator))
	return v
}
