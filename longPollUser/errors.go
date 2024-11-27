package longPollUser

import (
	"errors"
	"fmt"
	internalErrors "go-vk-sdk/errors"
)

type FailedError struct {
	Code int
}

func (e *FailedError) Error() string {
	return fmt.Sprintf("%s Long poll: failed code %d", internalErrors.MessagePrefix, e.Code)
}

var NotValidVersionError = errors.New(internalErrors.MessagePrefix + " Long poll user: not valid version")

type TooShortEventArrayError struct {
	Action string
	Least  int
	Got    int
}

func (e *TooShortEventArrayError) Error() string {
	return fmt.Sprintf("%s Long poll user: cannot init array into struct %s: expected at least %d element(s), got %d", internalErrors.MessagePrefix, e.Action, e.Least, e.Got)
}

type ExpectedSliceError struct {
	V interface{}
}

func (e *ExpectedSliceError) Error() string {
	return fmt.Sprintf("%s Long poll user: expected a slice, got %T", internalErrors.MessagePrefix, e.V)
}

type FailedCastError struct {
	V interface{}
}

func (e *FailedCastError) Error() string {
	return fmt.Sprintf("%s Long poll user: cast failed, value type: %T", internalErrors.MessagePrefix, e.V)
}

type InvalidEventTypeError struct {
	Type int
}

func (e *InvalidEventTypeError) Error() string {
	return fmt.Sprintf("%s Long poll user: invalid event type: %d", internalErrors.MessagePrefix, e.Type)
}
