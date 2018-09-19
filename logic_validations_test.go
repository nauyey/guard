package guard_test

import (
	"github.com/nauyey/guard"
	"github.com/nauyey/guard/validators"
	"testing"
)

func TestOrValidator(t *testing.T) {
	err := guard.Validate(
		guard.Or(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
		),
	)

	if err == nil {
		t.Errorf("guard.Or failed to revolse none having passed correctly")
	}

	err = guard.Validate(
		guard.Or(
			&validators.StringNotBlank{Value: "NOT BLANK"},
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
		),
	)

	if err != nil {
		t.Errorf("guard.Or failed to resolve one having passed correctly")
	}

	err = guard.Validate(
		guard.Or(
			&validators.StringNotBlank{Value: "NOT BLANK"},
			&validators.StringNotBlank{Value: "NOT BLANK"},
			&validators.StringNotBlank{Value: "NOT BLANK"},
		),
	)

	if err != nil {
		t.Errorf("guard.Or failed to resolve all having passed correctly")
	}
}

func TestXorValidator(t *testing.T) {
	err := guard.Validate(
		guard.Xor(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
		),
	)

	if err == nil {
		t.Errorf("guard.Xor failed to resolve to few having passed correctly")
	}

	err = guard.Validate(
		guard.Xor(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
		),
	)

	if err == nil {
		t.Errorf("guard.Xor failed to resolve to many having passed correctly")
	}

	err = guard.Validate(
		guard.Xor(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
			&validators.StringNotBlank{Value: ""},
		),
	)

	if err != nil {
		t.Errorf("guard.Xor failed to resolve only one having passed correctly")
	}
}

func TestNandValidator(t *testing.T) {
	err := guard.Validate(
		guard.Nand(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: ""},
		),
	)

	if err != nil {
		t.Errorf("guard.Xor failed to resolve none having passed correctly")
	}

	err = guard.Validate(
		guard.Nand(
			&validators.StringNotBlank{Value: ""},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
		),
	)

	if err != nil {
		t.Errorf("guard.Xor failed to resolve not all but one having passed correctly")
	}

	err = guard.Validate(
		guard.Nand(
			&validators.StringNotBlank{Value: "ONLY_ONE"},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
			&validators.StringNotBlank{Value: "ONLY_ONE"},
		),
	)

	if err == nil {
		t.Errorf("guard.Xor failed to resolve all having passed correctly")
	}
}
