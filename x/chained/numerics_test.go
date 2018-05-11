package chained_test

import (
	"testing"

	"github.com/nauyey/guard"
	. "github.com/nauyey/guard/validators"
	"github.com/nauyey/guard/x/chained"
)

func TestValidateInt(t *testing.T) {
	age := 7
	err := guard.Validate(
		chained.ValidateInt(age).
			With((&IsEven{}).OverrideMessage("override message")).
			With(&IntInRange{Left: 18, Right: 65}),
	)

	if err == nil {
		t.Fatalf("ValidateInt failed")
	}

	errs, ok := err.(guard.Errors)
	if !ok {
		t.Fatalf("ValidateInt failed without return guard.Errors")
	}

	if len(errs.ValidationErrors()) != 1 {
		t.Errorf("ValidateInt failed by validating data not strictly")
	}
}
