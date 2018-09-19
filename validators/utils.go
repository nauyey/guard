package validators

func contains(sl []string, target string) bool {
	for _, s := range sl {
		if s == target {
			return true
		}
	}

	return false
}

func returnDefaultStringIfNil(s *string, d string) string {
	if s != nil {
		return *s
	}
	return d
}

// ReturnGenericError provides errors for validations outside the validators packages
func ReturnGenericError(s string) error {
	return &validationError{s}
}
