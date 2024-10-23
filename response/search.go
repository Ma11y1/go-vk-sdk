package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/search

type SearchGetHintsResponse struct {
	BaseResponse
	Response struct {
		Count int                  `json:"count"`
		Items []objects.SearchHint `json:"items"`
	} `json:"response"`
}
