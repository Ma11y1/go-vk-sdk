package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/secure

type SecureAddAppEventResponse struct {
	BaseResponse
	Response objects.NumberFlagBool `json:"response"`
}

type SecureCheckTokenResponse struct {
	BaseResponse
	Response objects.SecureTokenChecked `json:"response"`
}

type SecureGetAppBalanceResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type SecureGetSMSHistoryResponse struct {
	BaseResponse
	Response []objects.SecureSmsNotification `json:"response"`
}

type SecureGetTransactionsHistoryResponse struct {
	BaseResponse
	Response []objects.SecureTransaction `json:"response"`
}

type SecureGetUserLevelResponse struct {
	BaseResponse
	Response []objects.SecureLevel `json:"response"`
}

type SecureGiveEventStickerResponse struct {
	BaseResponse
	Response []struct {
		UserID int    `json:"user_id"`
		Status string `json:"status"`
	} `json:"response"`
}

type SecureSendNotificationResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type SecureSendSMSNotificationResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type SecureSetCounterResponse struct {
	BaseResponse
	Response int `json:"response"`
}
