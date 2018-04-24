package validators_test

import (
	"testing"

	"github.com/nauyey/guard/validators"
)

func TestIsOdd(t *testing.T) {
	if err := (&validators.IsOdd{Value: 5}).Validate(); err != nil {
		t.Errorf("validators.IsOdd faild")
	}
	if err := (&validators.IsOdd{Value: 4}).Validate(); err == nil {
		t.Errorf("validators.IsOdd faild")
	}

	// test override error message
	err := (&validators.IsOdd{Value: 4}).Validate()
	if err == nil || err.Error() != "should be odd" {
		t.Errorf("validators.IsOdd faild")
	}
	err = (&validators.IsOdd{Value: 4}).OverrideMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.IsOdd faild")
	}
}

func TestIsEven(t *testing.T) {
	if err := (&validators.IsEven{Value: 6}).Validate(); err != nil {
		t.Errorf("validators.IsEven faild")
	}
	if err := (&validators.IsEven{Value: 7}).Validate(); err == nil {
		t.Errorf("validators.IsEven faild")
	}
}

func TestIntGreaterThan(t *testing.T) {
	if err := (&validators.IntGreaterThan{Value: 6, Target: 5}).Validate(); err != nil {
		t.Errorf("validators.IntGreaterThan faild")
	}
	if err := (&validators.IntGreaterThan{Value: 6, Target: 6}).Validate(); err == nil {
		t.Errorf("validators.IntGreaterThan faild")
	}
	if err := (&validators.IntGreaterThan{Value: 6, Target: 7}).Validate(); err == nil {
		t.Errorf("validators.IntGreaterThan faild")
	}
}

func TestIntGreaterThanOrEqualTo(t *testing.T) {
	if err := (&validators.IntGreaterThanOrEqualTo{Value: 6, Target: 5}).Validate(); err != nil {
		t.Errorf("validators.IntGreaterThanOrEqualTo faild")
	}
	if err := (&validators.IntGreaterThanOrEqualTo{Value: 6, Target: 6}).Validate(); err != nil {
		t.Errorf("validators.IntGreaterThanOrEqualTo faild")
	}
	if err := (&validators.IntGreaterThanOrEqualTo{Value: 6, Target: 7}).Validate(); err == nil {
		t.Errorf("validators.IntGreaterThanOrEqualTo faild")
	}
}

func TestIntEqualTo(t *testing.T) {
	if err := (&validators.IntEqualTo{Value: 6, Target: 5}).Validate(); err == nil {
		t.Errorf("validators.IntEqualTo faild")
	}
	if err := (&validators.IntEqualTo{Value: 6, Target: 6}).Validate(); err != nil {
		t.Errorf("validators.IntEqualTo faild")
	}
	if err := (&validators.IntEqualTo{Value: 6, Target: 7}).Validate(); err == nil {
		t.Errorf("validators.IntEqualTo faild")
	}
}

func TestIntLessThan(t *testing.T) {
	if err := (&validators.IntLessThan{Value: 6, Target: 5}).Validate(); err == nil {
		t.Errorf("validators.IntLessThan faild")
	}
	if err := (&validators.IntLessThan{Value: 6, Target: 6}).Validate(); err == nil {
		t.Errorf("validators.IntLessThan faild")
	}
	if err := (&validators.IntLessThan{Value: 6, Target: 7}).Validate(); err != nil {
		t.Errorf("validators.IntLessThan faild")
	}
}

func TestIntLessThanOrEqualTo(t *testing.T) {
	if err := (&validators.IntLessThanOrEqualTo{Value: 6, Target: 5}).Validate(); err == nil {
		t.Errorf("validators.IntLessThanOrEqualTo faild")
	}
	if err := (&validators.IntLessThanOrEqualTo{Value: 6, Target: 6}).Validate(); err != nil {
		t.Errorf("validators.IntLessThanOrEqualTo faild")
	}
	if err := (&validators.IntLessThanOrEqualTo{Value: 6, Target: 7}).Validate(); err != nil {
		t.Errorf("validators.IntLessThanOrEqualTo faild")
	}
}

func TestIntInRange(t *testing.T) {
	if err := (&validators.IntInRange{Value: 4, Left: 5, Right: 10}).Validate(); err == nil {
		t.Errorf("validators.IntInRange faild")
	}
	if err := (&validators.IntInRange{Value: 5, Left: 5, Right: 10}).Validate(); err != nil {
		t.Errorf("validators.IntInRange faild")
	}
	if err := (&validators.IntInRange{Value: 6, Left: 5, Right: 10}).Validate(); err != nil {
		t.Errorf("validators.IntInRange faild")
	}
	if err := (&validators.IntInRange{Value: 10, Left: 5, Right: 10}).Validate(); err != nil {
		t.Errorf("validators.IntInRange faild")
	}
	if err := (&validators.IntInRange{Value: 11, Left: 5, Right: 10}).Validate(); err == nil {
		t.Errorf("validators.IntInRange faild")
	}

	// test override error message
	err := (&validators.IntInRange{Value: 4, Left: 5, Right: 10}).Validate()
	if err == nil || err.Error() != "should be greater than or equal to left" {
		t.Errorf("validators.IntInRange faild")
	}
	err = (&validators.IntInRange{Value: 4, Left: 5, Right: 10}).OverrideLeftMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.IntInRange faild")
	}
	err = (&validators.IntInRange{Value: 11, Left: 5, Right: 10}).Validate()
	if err == nil || err.Error() != "should be less than or equal to right" {
		t.Errorf("validators.IntInRange faild")
	}
	err = (&validators.IntInRange{Value: 11, Left: 5, Right: 10}).OverrideRightMessage("override error message").Validate()
	if err == nil || err.Error() != "override error message" {
		t.Errorf("validators.IntInRange faild")
	}
}
