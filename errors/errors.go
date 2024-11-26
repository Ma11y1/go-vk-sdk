package errors

import (
	"errors"
	"fmt"
	"go-vk-sdk/logger"
)

const MessagePrefix = "GOVk-api-error: "

func Error(from, message string) error {
	return errors.New(fmt.Sprintf("[%s] %s: %s\n", MessagePrefix, from, message))
}

func ErrorLog(from, message string) error {
	str := fmt.Sprintf("[%s] %s: %s\n", MessagePrefix, from, message)
	logger.LogMessage(str)
	return errors.New(str)
}

type InternalError struct {
	From    string
	Message string
}

func (e *InternalError) Log() {
	logger.Log(e.From, e.Message)
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("[%s] %s: %s\n", MessagePrefix, e.From, e.Message)
}
