package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/likes

type LikesAddResponse struct {
	BaseResponse
	Response struct {
		Likes int `json:"likes"`
	} `json:"response"`
}

type LikesDeleteResponse struct {
	BaseResponse
	Response struct {
		Likes int `json:"likes"`
	} `json:"response"`
}

type LikesGetListResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type LikesGetListExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type LikesIsLikedResponse struct {
	BaseResponse
	Response struct {
		Liked  objects.BoolInt `json:"liked"`
		Copied objects.BoolInt `json:"copied"`
	} `json:"response"`
}
