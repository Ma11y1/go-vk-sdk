package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/fave

type FaveAddArticleResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveAddLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveAddPageResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveAddPostResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveAddProductResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveAddTagResponse struct {
	BaseResponse
	Response objects.FaveTag `json:"response"`
}

type FaveAddVideoResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveEditTagResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveGetResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.FaveItem `json:"items"`
	} `json:"response"`
}

type FaveGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.FaveItem `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type FaveGetPagesResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.FavePage `json:"items"`
	} `json:"response"`
}

type FaveGetTagsResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.FaveTag `json:"items"`
	} `json:"response"`
}

type FaveMarkSeenResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemoveArticleResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemoveLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemovePageResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemovePostResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemoveProductResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemoveTagResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveRemoveVideoResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveReorderTagsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveSetPageTagsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveSetTagsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FaveTrackPageInteractionResponse struct {
	BaseResponse
	Response int `json:"response"`
}
