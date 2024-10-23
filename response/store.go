package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/store

type StoreAddStickersToFavoriteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoreGetFavoriteStickersResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Sticker `json:"items"`
	} `json:"response"`
}

type StoreRemoveStickersFromFavoriteResponse struct {
	BaseResponse
	Response int `json:"response"`
}
