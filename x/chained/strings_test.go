package chained_test

import (
	"testing"

	"github.com/nauyey/guard"
	. "github.com/nauyey/guard/validators"
	"github.com/nauyey/guard/x/chained"
)

func TestValidateString(t *testing.T) {
	name := "nauyey"
	err := chained.ValidateString(name).
		With((&StringNotBlank{}).OverrideMessage("override message")).
		With(&StringLength{Min: 2, Max: 255}).
		With(&StringInclusion{In: []string{"Kent", "Bob", "Fowler", "DHH"}}).
		With(&StringExclusion{In: []string{"Jack", "Pony", "Robin"}}).
		Validate()

	if err == nil {
		t.Fatalf("ValidateString failed")
	}

	errs, ok := err.(guard.Errors)
	if !ok {
		t.Fatalf("ValidateString failed without return guard.Errors")
	}

	vErrs := errs.ValidationErrors()
	if len(vErrs) != 1 {
		t.Errorf("ValidateString failed by validating data not strictly")
	}

	if vErrs[0].Error() != "should be in" {
		t.Errorf("ValidateString failed")
	}
}
