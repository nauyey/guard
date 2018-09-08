package chained

import "github.com/nauyey/guard"

type valueValidator struct {
	validators []guard.Validator
}

func (v *valueValidator) Validate() error {
	return guard.Validate(
		v.validators...,
	)
}
