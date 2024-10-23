package errors

import "fmt"

type AuthDirectError struct {
	Type           string `json:"error"`
	Reason         string `json:"error_reason,omitempty"`
	Description    string `json:"error_description,omitempty"`
	CaptchaSID     string `json:"captcha_sid,omitempty"`
	CaptchaImg     string `json:"captcha_img,omitempty"`
	RedirectURI    string `json:"redirect_uri,omitempty"`
	ValidationType string `json:"validation_type,omitempty"`
	PhoneMask      string `json:"phone_mask,omitempty"`
}

func (e *AuthDirectError) Error() string {
	return fmt.Sprintf("%s Auth Direct error: %s\n %s", errorMessagePrefix, e.Type, e.Description)
}

type AuthCodeFlowError struct {
	Type        string `json:"errors"`
	Reason      string `json:"error_reason,omitempty"`
	Description string `json:"error_description,omitempty"`
}

func (e *AuthCodeFlowError) Error() string {
	return fmt.Sprintf("%s Auth Code flow error: %s\n %s", errorMessagePrefix, e.Type, e.Description)
}
