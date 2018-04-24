package guard

// Validator is the interface that defines a validator.
//
// Validate executes the current validator instance.
//
// If the validation instance is invalid then it returns a validation error.
// But if the validation execute with some other errors but not a validation error,
// it returns those errors, too.
// So the error can be validation errors or any other errors.
type Validator interface {
	Validate() error
}
