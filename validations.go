package guard

import "reflect"

// Validate executes the validators one by one.
//
// If a validator is invalid, it's validation error will be collected into a Errors instance.
//
// If a validator is invalid and strict, it's validation error will be collected into the Errors instance.
// Then Validate will stop executing validators and immediately return the Errors instance as an error.
//
// If a validator faild by internal error, Validate returns the error immediately.
//
// If the Errors instance contains any error, Validate returns the Errors instance.
//
// Otherwise, there aren't any validation errors. And nil will be returned.
//
func Validate(validators ...Validator) error {
	errs := []error{}

	for _, v := range validators {
		if err := v.Validate(); err != nil {
			switch vErr := err.(type) {
			default:
				return err
			case Errors:
				errs = append(errs, vErr.ValidationErrors()...)
				if _, ok := v.(*strictValidators); ok {
					return &errors{errs: errs}
				}
			case Error:
				errs = append(errs, err)
				if _, ok := v.(*strictValidators); ok {
					return &errors{errs: errs}
				}
			}
		}
	}

	if len(errs) != 0 {
		return &errors{errs: errs}
	}
	return nil
}

// Strict transfers a validator to be a strict validator.
// If a strict valdator faild, guard.Validate will stop executing the next validator but return validation errors.
func Strict(validators ...Validator) Validator {
	return &strictValidators{
		validators: validators,
	}
}

// AllowNil wraps a validator that if the validator is nil
// then this validator will be always valid. Otherwise, the validator
// will be the way as it before.
func AllowNil(v Validator) Validator {
	return &allowNilValidator{
		validator: v,
	}
}

type strictValidators struct {
	validators []Validator
}

// Validate implements the Validator interface
func (v *strictValidators) Validate() error {
	for _, v := range v.validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (v *strictValidators) strict() {}

type allowNilValidator struct {
	validator Validator
}

// Validate implements the Validator interface
func (v *allowNilValidator) Validate() error {
	vpv := reflect.ValueOf(v.validator)
	if vpv.Kind() == reflect.Invalid || (vpv.Kind() == reflect.Ptr) && vpv.IsNil() {
		return nil
	}
	return v.validator.Validate()
}

func (v *allowNilValidator) strict() {}

type errors struct {
	errs []error
}

// Error implements the error interface
func (err *errors) Error() string {
	return "validation errors"
}

// ValidationErrors implements the Errors interface
func (err *errors) ValidationErrors() []error {
	return err.errs
}
