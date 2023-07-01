package ex

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

type AlreadyExistError struct {
	msg   string
	Field string
}

type ErrAlreadyExists struct {
	Field      string
	Constraint string
}

func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("Duplicate value for field '%s' violated constraint '%s'", e.Field, e.Constraint)
}

type ErrValidation struct {
	Field  string
	Reason string
}

func (e *ErrValidation) Error() string {
	return fmt.Sprintf("Field '%s' %s", e.Field, e.Reason)
}
