package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/gifts

type GiftsGetResponse struct {
	BaseResponse
	Response struct {
		Count int            `json:"count"`
		Items []objects.Gift `json:"items"`
	} `json:"response"`
}
