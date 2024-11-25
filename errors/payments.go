package errors

import (
	"strconv"
)

// PaymentsError
//
// When critical errors occur, the order cancels and the app event onOrderFail is sent.
// If the error is temporary, a notification will be resent after some time and the user will wait for the process is completed.
//
// Doc: https://dev.vk.com/ru/api/payments/notifications/overview
type PaymentsError struct {
	Code PaymentsErrorCode `json:"error_code"`
	Msg  string            `json:"error_msg,omitempty"` // Msg description in easy-to-read format (required for error_code >= 100).

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

func (e *PaymentsError) Error() string {
	if e.Msg != "" {
		return errorMessagePrefix + " Payments: " + e.Msg
	}

	return errorMessagePrefix + " Payments: error with code " + strconv.Itoa(int(e.Code))
}
