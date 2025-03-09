package utils

type MultiError struct {
	Errors []error
}

func NewMultiError(errs []error) *MultiError {
	return &MultiError{Errors: errs}
}

func NewEmptyMultiError() *MultiError {
	return &MultiError{Errors: []error{}}
}

func (m *MultiError) Add(err error) {
	m.Errors = append(m.Errors, err)
}

func (m *MultiError) HasErrors() bool {
	return len(m.Errors) > 0
}

func (m *MultiError) SelfOrNil() *MultiError {
	if m.HasErrors() {
		return m
	}

	return nil
}

func (m *MultiError) Error() string {
	result := ""
	for _, err := range m.Errors {
		result += err.Error() + "\n"
	}

	return "multiple errors occurred:\n" + result
}
