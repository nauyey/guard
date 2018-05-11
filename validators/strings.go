package validators

import "regexp"

const blank = `\A[[:space:]]*\z`

var blankReg *regexp.Regexp

func init() {
	var err error
	if blankReg, err = regexp.Compile(blank); err != nil {
		panic(err)
	}
}

// string validation error messages
const (
	stringNotBlankMsg  = "shouldn't be blank"
	stringInclusionMsg = "should be in"
	stringExclusionMsg = "shouldn't be in"
	tooShortMsg        = "too short"
	tooLongMsg         = "too long"
)

// StringNotBlank is a validator which will check whether the field Value is not blank.
//
// A string is blank if it's empty or contains whitespaces only:
// ""       -> blank
// "  "     -> blank
// "	"   -> blank
// "\t\n\r" -> blank
// " abc "  -> not blank
//
type StringNotBlank struct {
	Value string

	message *string
}

// Validate implements the guard.Validator interface
func (v *StringNotBlank) Validate() error {
	if len(v.Value) == 0 || blankReg.Match([]byte(v.Value)) {
		return &validationError{returnDefaultStringIfNil(v.message, stringNotBlankMsg)}
	}

	return nil
}

func (v *StringNotBlank) SetValue(value string) {
	v.Value = value
}

// OverrideMessage overrides the validation error message of current validator
func (v *StringNotBlank) OverrideMessage(msg string) *StringNotBlank {
	v.message = &msg
	return v
}

// StringInclusion is a validator which will check whether the field Value is included by field In.
type StringInclusion struct {
	Value string
	In    []string

	message *string
}

// Validate implements the guard.Validator interface
func (v *StringInclusion) Validate() error {
	if !contains(v.In, v.Value) {
		return &validationError{returnDefaultStringIfNil(v.message, stringInclusionMsg)}
	}

	return nil
}

func (v *StringInclusion) SetValue(value string) {
	v.Value = value
}

// OverrideMessage overrides the validation error message of current validator
func (v *StringInclusion) OverrideMessage(msg string) *StringInclusion {
	v.message = &msg
	return v
}

// StringExclusion is a validator which will check whether the field Value is excluded by field In.
type StringExclusion struct {
	Value string
	In    []string

	message *string
}

// Validate implements the guard.Validator interface
func (v *StringExclusion) Validate() error {
	if contains(v.In, v.Value) {
		return &validationError{returnDefaultStringIfNil(v.message, stringExclusionMsg)}
	}

	return nil
}

func (v *StringExclusion) SetValue(value string) {
	v.Value = value
}

// OverrideMessage overrides the validation error message of current validator
func (v *StringExclusion) OverrideMessage(msg string) *StringExclusion {
	v.message = &msg
	return v
}

// StringLength is a validator which will check whether the length of field Value is in range of filed Min and Max.
type StringLength struct {
	Value string
	Min   int
	Max   int

	tooShortMessage *string
	tooLongMessage  *string
}

// Validate implements the guard.Validator interface
func (v *StringLength) Validate() error {
	if len(v.Value) < v.Min {
		return &validationError{returnDefaultStringIfNil(v.tooShortMessage, tooShortMsg)}
	}

	if len(v.Value) > v.Max {
		return &validationError{returnDefaultStringIfNil(v.tooLongMessage, tooLongMsg)}
	}

	return nil
}

func (v *StringLength) SetValue(value string) {
	v.Value = value
}

// OverrideTooShortMessage overrides the error message of the too short string validation
func (v *StringLength) OverrideTooShortMessage(msg string) *StringLength {
	v.tooShortMessage = &msg
	return v
}

// OverrideTooLongMessage overrides the error message of the too long string validation
func (v *StringLength) OverrideTooLongMessage(msg string) *StringLength {
	v.tooLongMessage = &msg
	return v
}
