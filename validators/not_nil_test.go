package validators_test

import (
	"testing"

	"github.com/nauyey/guard/validators"
)

func TestNotNil(t *testing.T) {
	type test struct{}

	var (
		nilTest  *test
		nilTest2 *test = nil
		nilTest3 int   = 0
	)

	if err := (&validators.NotNil{Value: "abc"}).Validate(); err != nil {
		t.Errorf("validators.NotNil faild")
	}
	if err := (&validators.NotNil{Value: nil}).Validate(); err == nil {
		t.Errorf("validators.NotNil faild")
	}
	if err := (&validators.NotNil{Value: nilTest}).Validate(); err == nil {
		t.Errorf("validators.NotNil faild")
	}
	if err := (&validators.NotNil{Value: nilTest2}).Validate(); err == nil {
		t.Errorf("validators.NotNil faild")
	}
	if err := (&validators.NotNil{Value: nilTest3}).Validate(); err != nil {
		t.Errorf("validators.NotNil faild")
	}

	// test override error message
	err := (&validators.NotNil{Value: nilTest}).Validate()
	if err == nil || err.Error() != "shouldn't be nil" {
		t.Errorf("validators.NotNil faild")
	}
	err = (&validators.NotNil{Value: nilTest}).OverrideMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.NotNil faild")
	}
}
