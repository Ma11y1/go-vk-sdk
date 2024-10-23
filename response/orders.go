package response

import (
	"go-vk-sdk/constants"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/method/orders

type OrdersCancelSubscriptionResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type OrdersChangeStateResponse struct {
	BaseResponse
	Response constants.OrderState `json:"response"`
}

type OrdersGetResponse struct {
	BaseResponse
	Response []objects.OrdersOrder `json:"response"`
}

type OrdersGetAmountResponse struct {
	BaseResponse
	Response []objects.OrdersAmount `json:"response"`
}

type OrdersGetByIDResponse struct {
	BaseResponse
	Response []objects.OrdersOrder `json:"response"`
}

type OrdersGetUserSubscriptionByIDResponse struct {
	BaseResponse
	Response objects.OrdersSubscription `json:"response"`
}

type OrdersGetUserSubscriptionsResponse struct {
	BaseResponse
	Response struct {
		Count int                          `json:"count"` // Total number
		Items []objects.OrdersSubscription `json:"items"`
	} `json:"response"`
}

type OrdersUpdateSubscriptionResponse struct {
	BaseResponse
	Response int `json:"response"`
}
