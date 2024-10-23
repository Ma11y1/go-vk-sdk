package errors

import (
	"errors"
	"fmt"
)

type FailedError struct {
	Code int
}

func (e *FailedError) Error() string {
	return fmt.Sprintf("%s Long poll: failed code %d", errorMessagePrefix, e.Code)
}

var NotValidVersionError = errors.New(errorMessagePrefix + " Long poll user: not valid version")

type TooShortEventArrayError struct {
	Action string
	Least  int
	Got    int
}

func (e *TooShortEventArrayError) Error() string {
	return fmt.Sprintf("%s Long poll user: cannot init array into struct %s: expected at least %d element(s), got %d", errorMessagePrefix, e.Action, e.Least, e.Got)
}

type ExpectedSliceError struct {
	V interface{}
}

func (e *ExpectedSliceError) Error() string {
	return fmt.Sprintf("%s Long poll user: expected a slice, got %T", errorMessagePrefix, e.V)
}

type FailedCastError struct {
	V interface{}
}

func (e *FailedCastError) Error() string {
	return fmt.Sprintf("%s Long poll user: cast failed, value type: %T", errorMessagePrefix, e.V)
}

type InvalidEventTypeError struct {
	Type int
}

func (e *InvalidEventTypeError) Error() string {
	return fmt.Sprintf("%s Long poll user: invalid event type: %d", errorMessagePrefix, e.Type)
}
