package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/users

type UsersGetResponse struct {
	BaseResponse
	Response []objects.UserFull `json:"response"`
}

type UsersGetFollowersResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type UsersGetFollowersFieldsResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type UsersGetSubscriptionsResponse struct {
	BaseResponse
	Response struct {
		Users struct {
			Count int   `json:"count"`
			Items []int `json:"items"`
		} `json:"users"`
		Groups struct {
			Count int   `json:"count"`
			Items []int `json:"items"`
		} `json:"groups"`
	} `json:"response"`
}

type UsersReportResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type UsersSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}
