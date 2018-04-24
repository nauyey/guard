package validators

import (
	"reflect"
)

const (
	notNilMsg = "shouldn't be nil"
)

// NotNil is a validator which will check whether the field Value is not nil.
type NotNil struct {
	Value interface{}

	message *string
}

// Validate implements the guard.Validator interface
func (v *NotNil) Validate() error {
	vpv := reflect.ValueOf(v.Value)
	if vpv.Kind() == reflect.Invalid || (vpv.Kind() == reflect.Ptr) && vpv.IsNil() {
		return &validationError{returnDefaultStringIfNil(v.message, notNilMsg)}
	}

	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *NotNil) OverrideMessage(msg string) *NotNil {
	v.message = &msg
	return v
}
