package validators

// numeric validation error messages
const (
	isOddMsg                   = "should be odd"
	isEvenMsg                  = "should be even"
	intGreaterThanMsg          = "should be great than"
	intGreaterThanOrEqualToMsg = "should be great than or equal to"
	intEqualToMsg              = "should equal to"
	intLessThanMsg             = "should be less than"
	intLessThanOrEqualToMsg    = "should be less than or equal to"
	intInRangeLeftMsg          = "should be greater than or equal to left"
	intInRangeRightMsg         = "should be less than or equal to right"
)

// IsOdd is a validator which will check whether the field Value is odd.
type IsOdd struct {
	Value int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IsOdd) Validate() error {
	if v.Value%2 == 0 {
		return &validationError{returnDefaultStringIfNil(v.message, isOddMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IsOdd) OverrideMessage(msg string) *IsOdd {
	v.message = &msg
	return v
}

// IsEven is a validator which will check whether the field Value is even.
type IsEven struct {
	Value int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IsEven) Validate() error {
	if v.Value%2 != 0 {
		return &validationError{returnDefaultStringIfNil(v.message, isEvenMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IsEven) OverrideMessage(msg string) *IsEven {
	v.message = &msg
	return v
}

// IntGreaterThan is a validator which will check whether the field Value is greater than field Target.
type IntGreaterThan struct {
	Value  int
	Target int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IntGreaterThan) Validate() error {
	if v.Value <= v.Target {
		return &validationError{returnDefaultStringIfNil(v.message, intGreaterThanMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IntGreaterThan) OverrideMessage(msg string) *IntGreaterThan {
	v.message = &msg
	return v
}

// IntGreaterThanOrEqualTo is a validator which will check whether the field Value is greater than or equal to field Target.
type IntGreaterThanOrEqualTo struct {
	Value  int
	Target int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IntGreaterThanOrEqualTo) Validate() error {
	if v.Value < v.Target {
		return &validationError{returnDefaultStringIfNil(v.message, intGreaterThanOrEqualToMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IntGreaterThanOrEqualTo) OverrideMessage(msg string) *IntGreaterThanOrEqualTo {
	v.message = &msg
	return v
}

// IntEqualTo is a validator which will check whether the field Value equals to field Target.
type IntEqualTo struct {
	Value  int
	Target int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IntEqualTo) Validate() error {
	if v.Value != v.Target {
		return &validationError{returnDefaultStringIfNil(v.message, intEqualToMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IntEqualTo) OverrideMessage(msg string) *IntEqualTo {
	v.message = &msg
	return v
}

// IntLessThan is a validator which will check whether the field Value is less than field Target.
type IntLessThan struct {
	Value  int
	Target int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IntLessThan) Validate() error {
	if v.Value >= v.Target {
		return &validationError{returnDefaultStringIfNil(v.message, intLessThanMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IntLessThan) OverrideMessage(msg string) *IntLessThan {
	v.message = &msg
	return v
}

// IntLessThanOrEqualTo is a validator which will check whether the field Value is less than or equal to field Target.
type IntLessThanOrEqualTo struct {
	Value  int
	Target int

	message *string
}

// Validate implements the guard.Validator interface
func (v *IntLessThanOrEqualTo) Validate() error {
	if v.Value > v.Target {
		return &validationError{returnDefaultStringIfNil(v.message, intLessThanOrEqualToMsg)}
	}
	return nil
}

// OverrideMessage overrides the validation error message of current validator
func (v *IntLessThanOrEqualTo) OverrideMessage(msg string) *IntLessThanOrEqualTo {
	v.message = &msg
	return v
}

// IntInRange is a validator which will check whether the field Value is in range of field Left and Right.
type IntInRange struct {
	Value int
	Left  int
	Right int

	leftMessage  *string
	rightMessage *string
}

// Validate implements the guard.Validator interface
func (v *IntInRange) Validate() error {
	if v.Value < v.Left {
		return &validationError{returnDefaultStringIfNil(v.leftMessage, intInRangeLeftMsg)}
	}
	if v.Value > v.Right {
		return &validationError{returnDefaultStringIfNil(v.rightMessage, intInRangeRightMsg)}
	}
	return nil
}

// OverrideLeftMessage overrides the error message of the out of left range validation
func (v *IntInRange) OverrideLeftMessage(msg string) *IntInRange {
	v.leftMessage = &msg
	return v
}

// OverrideRightMessage overrides the error message of the out of right range validation
func (v *IntInRange) OverrideRightMessage(msg string) *IntInRange {
	v.rightMessage = &msg
	return v
}
