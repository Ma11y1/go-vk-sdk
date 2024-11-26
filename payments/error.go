package payments

import (
	internalErrors "go-vk-sdk/errors"
	"strconv"
)

// ErrorCode
//
//	Errors 100-999 are specified by the app. Such errors always include an error description.
//	See https://dev.vk.com/ru/api/payments/notifications/overview
type ErrorCode int

const (
	CommonError            ErrorCode = 1  // Critical: true/false.
	TemporaryDatabaseError ErrorCode = 2  // Critical: false.
	BadSignatures          ErrorCode = 10 // Critical: true.
	// BadRequest Request parameters do not comply with the specification.
	// No required fields in the request.
	// Other request integrity errors.
	BadRequest        ErrorCode = 11 // Critical: true.
	ProductNotExist   ErrorCode = 20 // Critical: true.
	ProductOutOfStock ErrorCode = 21 // Critical: true.
	UserNotExist      ErrorCode = 22 // Critical: true.
)

// Error
//
// When critical errors occur, the order cancels and the app event onOrderFail is sent.
// If the error is temporary, a notification will be resent after some time and the user will wait for the process is completed.
//
// Doc: https://dev.vk.com/ru/api/payments/notifications/overview
type Error struct {
	Code ErrorCode `json:"error_code"`
	Msg  string    `json:"error_msg,omitempty"` // Msg description in easy-to-read format (required for error_code >= 100).

	// IsCritical severity. Possible values are:
	//
	// true — if a notification with identical parameters is passed
	// repeatedly, the same error will occur. For example, the indicated
	// product does not exist. The notification will not be resent as the user
	// will receive an error.
	//
	// false — if the error is temporary, a notification may be processed
	// later. For example, a temporary error in posting to the database.
	// The notification will be sent after some time and the user will wait
	// for the response.
	IsCritical bool `json:"critical"`
}

func (e *Error) Error() string {
	if e.Msg != "" {
		return internalErrors.MessagePrefix + " Payments: " + e.Msg
	}

	return internalErrors.MessagePrefix + " Payments: error with code " + strconv.Itoa(int(e.Code))
}
