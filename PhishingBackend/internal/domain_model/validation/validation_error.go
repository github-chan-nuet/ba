package validation

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Errors []FieldError
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		Errors: make([]FieldError, 0, 5),
	}
}

func (v *ValidationError) Len() int {
	return len(v.Errors)
}

func (v *ValidationError) HasErr() bool {
	return len(v.Errors) == 0
}

func (v *ValidationError) Add(field string, reason Reason) {
	v.Errors = append(v.Errors, FieldError{field, reason})
}

func (v *ValidationError) Error() string {
	errors := make([]string, 0, len(v.Errors))
	for _, err := range v.Errors {
		errors = append(errors, err.Error())
	}
	fields := strings.Join(errors, ", ")
	return "validation error occurred. Fields: " + fields
}

type FieldError struct {
	Field  string
	Reason Reason
}

func (f *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", f.Field, f.Reason)
}

type Reason string

const (
	Mandatory  Reason = "must be set"
	NoFieldSet Reason = "no field at all was set but at least one must be"
)
