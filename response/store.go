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

type StoreGetProductsResponse struct {
	BaseResponse
	Response []objects.Product `json:"response"`
	//Response struct {
	//	Count int               `json:"count"`
	//	Items []objects.Product `json:"items"`
	//} `json:"response"`
}

type StoreGetStickersKeywordsResponse struct {
	BaseResponse
	Response objects.StickersKeywords `json:"response"`
}

type StoreRemoveStickersFromFavoriteResponse struct {
	BaseResponse
	Response int `json:"response"`
}
