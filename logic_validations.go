package guard

import (
	"github.com/nauyey/guard/validators"
)

const (
	atLeastOneMsg = "at leat one" // OR
	onlyOneMsg    = "only one"    // XOR
	notAllMsg     = "not all"     // NAND
)

type (
	orValidation struct {
		validators []Validator
	}
	xorValidation struct {
		validators []Validator
	}
	nandValidation struct {
		validators []Validator
	}
)

// Or wraps validators in a group in which
// *at least one* of the validators is allowed to pass.
func Or(validators ...Validator) Validator {
	return &orValidation{
		validators: validators,
	}
}

// Validate implements the Validator interface
func (v *orValidation) Validate() error {
	// evaluate validators inside or-group
	errs := Validate(v.validators...)

	if errs != nil && len(v.validators)-len(errs.(Errors).ValidationErrors()) == 0 {
		return validators.ReturnGenericError(atLeastOneMsg)
	}

	return nil
}

// Xor wraps validators in a group in which
// *only one* of the validators is allowed to pass.
func Xor(validators ...Validator) Validator {
	return &xorValidation{
		validators: validators,
	}
}

// Validate implements the Validator interface
func (v *xorValidation) Validate() error {
	// evaluate validators inside xor-group
	errs := Validate(v.validators...)

	if errs == nil || len(v.validators)-len(errs.(Errors).ValidationErrors()) != 1 {
		return validators.ReturnGenericError(onlyOneMsg)
	}

	return nil
}

// Nand wraps validators in a group in which
// *not all* of the validators are allowed pass.
func Nand(validators ...Validator) Validator {
	return &nandValidation{
		validators: validators,
	}
}

// Validate implements the Validator interface
func (v *nandValidation) Validate() error {
	// evaluate validators inside nand-group
	errs := Validate(v.validators...)

	// no error -> all have passed -> error
	if errs == nil {
		return validators.ReturnGenericError(notAllMsg)
	}

	return nil
}
