package validators_test

import (
	"testing"

	"github.com/nauyey/guard/validators"
)

func TestStringNotBlank(t *testing.T) {
	if err := (&validators.StringNotBlank{Value: "abc"}).Validate(); err != nil {
		t.Errorf("validators.StringNotBlank faild")
	}
	if err := (&validators.StringNotBlank{Value: ""}).Validate(); err == nil {
		t.Errorf("validators.StringNotBlank faild")
	}
	if err := (&validators.StringNotBlank{Value: " "}).Validate(); err == nil {
		t.Errorf("validators.StringNotBlank faild")
	}
	if err := (&validators.StringNotBlank{Value: " 	\t\n"}).Validate(); err == nil {
		t.Errorf("validators.StringNotBlank faild")
	}

	// test override error message
	err := (&validators.StringNotBlank{Value: ""}).Validate()
	if err == nil || err.Error() != "shouldn't be blank" {
		t.Errorf("validators.StringNotBlank faild")
	}
	err = (&validators.StringNotBlank{Value: ""}).OverrideMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.StringNotBlank faild")
	}
}

func TestStringInclusion(t *testing.T) {
	if err := (&validators.StringInclusion{}).Validate(); err == nil {
		t.Errorf("validators.StringInclusion faild")
	}
	if err := (&validators.StringInclusion{Value: "abc"}).Validate(); err == nil {
		t.Errorf("validators.StringInclusion faild")
	}
	if err := (&validators.StringInclusion{Value: "abc", In: []string{"bcd", "cdf"}}).Validate(); err == nil {
		t.Errorf("validators.StringInclusion faild")
	}
	if err := (&validators.StringInclusion{Value: "abc", In: []string{"abc", "bcd", "cdf"}}).Validate(); err != nil {
		t.Errorf("validators.StringInclusion faild")
	}
}

func TestStringExclusion(t *testing.T) {
	if err := (&validators.StringExclusion{}).Validate(); err != nil {
		t.Errorf("validators.StringExclusion faild")
	}
	if err := (&validators.StringExclusion{Value: "abc"}).Validate(); err != nil {
		t.Errorf("validators.StringExclusion faild")
	}
	if err := (&validators.StringExclusion{Value: "abc", In: []string{"bcd", "cdf"}}).Validate(); err != nil {
		t.Errorf("validators.StringExclusion faild")
	}
	if err := (&validators.StringExclusion{Value: "abc", In: []string{"abc", "bcd", "cdf"}}).Validate(); err == nil {
		t.Errorf("validators.StringExclusion faild")
	}
}

func TestStringLength(t *testing.T) {
	if err := (&validators.StringLength{}).Validate(); err != nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "ab"}).Validate(); err == nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "ab", Min: 3, Max: 5}).Validate(); err == nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "abc", Min: 3, Max: 5}).Validate(); err != nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "abcd", Min: 3, Max: 5}).Validate(); err != nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "abcdf", Min: 3, Max: 5}).Validate(); err != nil {
		t.Errorf("validators.StringLength faild")
	}
	if err := (&validators.StringLength{Value: "abcdfg", Min: 3, Max: 5}).Validate(); err == nil {
		t.Errorf("validators.StringLength faild")
	}

	// test override error message
	err := (&validators.StringLength{Value: "ab", Min: 3, Max: 5}).Validate()
	if err == nil || err.Error() != "too short" {
		t.Errorf("validators.StringLength faild")
	}
	err = (&validators.StringLength{Value: "ab", Min: 3, Max: 5}).OverrideTooShortMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.StringLength faild")
	}
	err = (&validators.StringLength{Value: "abcdfg", Min: 3, Max: 5}).Validate()
	if err == nil || err.Error() != "too long" {
		t.Errorf("validators.StringLength faild")
	}
	err = (&validators.StringLength{Value: "abcdfg", Min: 3, Max: 5}).OverrideTooLongMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.StringLength faild")
	}
}
