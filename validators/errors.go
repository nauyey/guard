package validators

type validationError struct {
	msg string
}

// Error implements the error interface
func (err *validationError) Error() string {
	return err.msg
}

// ValidationError implements the guard.Error interface
func (err *validationError) ValidationError() {}
