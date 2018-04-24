package guard_test

import (
	"errors"
	"testing"

	"github.com/nauyey/guard"
)

type validationError struct {
	msg string
}

func (err *validationError) Error() string {
	return err.msg
}

func (err *validationError) ValidationError() {}

type validationErrors struct {
	errs []error
}

func (err *validationErrors) Error() string {
	return "validation errors"
}

func (err *validationErrors) ValidationErrors() []error {
	return err.errs
}

type testValidator struct {
	err error
}

func (v *testValidator) Validate() error {
	return v.err
}

func TestValidate(t *testing.T) {
	// test without validation errors
	err := guard.Validate(
		&testValidator{},
		guard.Strict(&testValidator{}),
		&testValidator{},
	)
	if err != nil {
		t.Errorf("guard.Validate failed with err=%v", err)
	}

	// test with validation errors
	err = guard.Validate(
		&testValidator{
			err: &validationError{},
		},
		&testValidator{},
		guard.Strict(&testValidator{}),
		&testValidator{
			err: &validationError{},
		},
		&testValidator{
			err: &validationError{},
		},
	)
	if err == nil {
		t.Fatalf("guard.Validate failed with err=nil")
	}
	errs, ok := err.(guard.Errors)
	if !ok {
		t.Fatalf("guard.Validate failed to return err(type guard.Errors)")
	}
	vErrs := errs.ValidationErrors()
	if len(vErrs) != 3 {
		t.Errorf("guard.Validate failed with len(vErrs)=%d, want len(vErrs)=3", len(vErrs))
	}

	// test with validation errors and guard.Errors
	err = guard.Validate(
		&testValidator{
			err: &validationError{},
		},
		&testValidator{},
		guard.Strict(&testValidator{}),
		&testValidator{
			err: &validationError{},
		},
		&testValidator{
			err: &validationErrors{
				errs: []error{
					&validationError{},
					&validationError{},
				},
			},
		},
		&testValidator{
			err: &validationError{},
		},
	)
	if err == nil {
		t.Fatalf("guard.Validate failed with err=nil")
	}
	errs, ok = err.(guard.Errors)
	if !ok {
		t.Fatalf("guard.Validate failed to return err(type guard.Errors)")
	}
	vErrs = errs.ValidationErrors()
	if len(vErrs) != 5 {
		t.Errorf("guard.Validate failed with len(vErrs)=%d, want len(vErrs)=5", len(vErrs))
	}

	// test with validation errors and strict validation errors
	err = guard.Validate(
		&testValidator{
			err: &validationError{},
		},
		&testValidator{},
		guard.Strict(&testValidator{
			err: &validationError{},
		}),
		&testValidator{
			err: errors.New("non-validation error"),
		},
		&testValidator{
			err: &validationError{},
		},
		&testValidator{
			err: &validationError{},
		},
	)
	if err == nil {
		t.Fatalf("guard.Validate failed with err=nil")
	}
	errs, ok = err.(guard.Errors)
	if !ok {
		t.Fatalf("guard.Validate failed to return err(type guard.Errors)")
	}
	vErrs = errs.ValidationErrors()
	if len(vErrs) != 2 {
		t.Errorf("guard.Validate failed with len(vErrs)=%d, want len(vErrs)=2", len(vErrs))
	}

	// test with non-validation errors
	err = guard.Validate(
		&testValidator{
			err: &validationError{},
		},
		&testValidator{},
		&testValidator{
			err: errors.New("non-validation error"),
		},
		guard.Strict(&testValidator{
			err: &validationError{},
		}),
		&testValidator{
			err: &validationError{},
		},
	)
	if err == nil {
		t.Fatalf("guard.Validate failed with err=nil")
	}
	errs, ok = err.(guard.Errors)
	if ok {
		t.Fatalf("guard.Validate failed to return non-validation error")
	}
}

func TestValidateStrictly(t *testing.T) {
	// test without validation errors
	err := guard.Validate(
		&testValidator{},
		guard.Strict(
			&testValidator{},
			guard.Strict(&testValidator{}),
			&testValidator{},
		),
	)
	if err != nil {
		t.Errorf("guard.Validate Strictly failed with err=%v", err)
	}

	// test with validation errors
	err = guard.Validate(
		guard.Strict(
			&testValidator{
				err: &validationError{},
			},
			&testValidator{},
			&testValidator{
				err: &validationError{},
			},
		),
	)
	if err == nil {
		t.Fatalf("guard.Validate Strictly failed with err=nil")
	}
	errs, ok := err.(guard.Errors)
	if !ok {
		t.Fatalf("guard.Validate Strictly failed to return err(type guard.Errors)")
	}
	vErrs := errs.ValidationErrors()
	if len(vErrs) != 1 {
		t.Errorf("guard.Validate Strictly failed with len(vErrs)=%d, want len(vErrs)=1", len(vErrs))
	}

	// test with guard.Errors
	err = guard.Validate(
		guard.Strict(
			&testValidator{},
			guard.Strict(&testValidator{}),
			&testValidator{
				err: &validationErrors{
					errs: []error{
						&validationError{},
						&validationError{},
					},
				},
			},
			&testValidator{
				err: &validationError{},
			},
		),
	)
	if err == nil {
		t.Fatalf("guard.ValidateStrictly failed with err=nil")
	}
	errs, ok = err.(guard.Errors)
	if !ok {
		t.Fatalf("guard.ValidateStrictly failed to return err(type guard.Errors)")
	}
	vErrs = errs.ValidationErrors()
	if len(vErrs) != 2 {
		t.Errorf("guard.ValidateStrictly failed with len(vErrs)=%d, want len(vErrs)=2", len(vErrs))
	}
}
