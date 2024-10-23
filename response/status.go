package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/status

type StatusGetResponse struct {
	BaseResponse
	Response struct {
		Audio objects.Audio `json:"audio"`
		Text  string        `json:"text"`
	} `json:"response"`
}

type StatusSetResponse struct {
	BaseResponse
	Response int `json:"response"`
}
