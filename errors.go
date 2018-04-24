package guard

// Error is the interface that defines a validation error.
// Any type implements this interface will be treated as a validation error.
type Error interface {
	ValidationError()
}

// Errors is the interface that defines a multiple validation errors type.
//
// ValidationErrors returns a error slice.
type Errors interface {
	ValidationErrors() []error
}
