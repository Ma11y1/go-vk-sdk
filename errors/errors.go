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

type NumberAPIError int

func (e *NumberAPIError) Error() string {
	return fmt.Sprintf("%s error number: %d", errorMessagePrefix, e)
}

type APIError struct {
	Code             int    `json:"error_code"`
	Subcode          int    `json:"error_subcode,omitempty"`
	Message          string `json:"error_msg"`
	Text             string `json:"error_text"`
	CaptchaSID       string `json:"captcha_sid,omitempty"`
	CaptchaImg       string `json:"captcha_img,omitempty"`
	ConfirmationText string `json:"confirmation_text,omitempty"`
	RedirectURI      string `json:"redirect_uri,omitempty"`
	RequestParams    []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s error code: %d -> %s", errorMessagePrefix, e.Code, e.Message)
}

type ExecuteAPIErrors []ExecuteAPIError

func (e *ExecuteAPIErrors) Error() string {
	return fmt.Sprintf("%s execute errors: %w", errorMessagePrefix, e)
}

type ExecuteAPIError struct {
	Code    int    `json:"error_code"`
	Method  string `json:"method"`
	Message string `json:"error_msg"`
}

func (e *ExecuteAPIError) Error() string {
	return fmt.Sprintf("%s execute error code: %d -> %s", errorMessagePrefix, e.Code, e.Message)
}

type UploadAPIError struct {
	Err         string `json:"errors"`
	Code        int    `json:"error_code"`
	Description string `json:"error_descr"`
	IsLogged    bool   `json:"error_is_logged"`
}

func (e *UploadAPIError) Error() string {
	if e.Err != "" {
		return fmt.Sprintf("%s upload error code:%d -> %s", errorMessagePrefix, e.Code, e.Err)
	}

	return fmt.Sprintf("%s upload error code:%d -> %s", errorMessagePrefix, e.Code, e.Description)
}

type AdsAPIError struct {
	Code        string `json:"error_code"`
	Description string `json:"error_desc"`
}

func (e *AdsAPIError) Error() string {
	return fmt.Sprintf("%s Ads error: %s\n %s", errorMessagePrefix, e.Code, e.Description)
}
