package errors

import (
	"errors"
	"fmt"
	"go-vk-sdk/logger"
)

const errorMessagePrefix = "GOVk-api-error: "

func Error(from, message string) error {
	return errors.New(fmt.Sprintf("[%s] %s: %s\n", errorMessagePrefix, from, message))
}

type InternalError struct {
	From    string
	Message string
}

func (e *InternalError) Debug() {
	logger.Log(e.From, e.Message)
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("[%s] %s: %s\n", errorMessagePrefix, e.From, e.Message)
}
