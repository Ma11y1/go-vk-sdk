package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/donut

type DonutGetFriendsResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type DonutGetSubscriptionResponse struct {
	BaseResponse
	Response objects.DonutDonatorSubscriptionInfo `json:"response"`
}

type DonutGetSubscriptionsResponse struct {
	BaseResponse
	Response struct {
		Subscriptions []objects.DonutDonatorSubscriptionInfo `json:"subscriptions"`
		Count         int                                    `json:"count"`
		Profiles      []objects.UserFull                     `json:"profiles"`
		Groups        []objects.GroupFull                    `json:"groups"`
	} `json:"response"`
}

type DonutIsDonResponse struct {
	BaseResponse
	Response int `json:"response"`
}
