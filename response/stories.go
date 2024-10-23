package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/stories

type StoriesBanOwnerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoriesDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoriesGetResponse struct {
	BaseResponse
	Response struct {
		Count            int                       `json:"count"`
		Items            []objects.StoriesFeedItem `json:"items"`
		PromoData        objects.StoriesPromoBlock `json:"promo_data"`
		NeedUploadScreen objects.NumberFlagBool    `json:"need_upload_screen"`
	} `json:"response"`
}

type StoriesGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count            int                       `json:"count"`
		Items            []objects.StoriesFeedItem `json:"items"`
		PromoData        objects.StoriesPromoBlock `json:"promo_data"`
		NeedUploadScreen objects.NumberFlagBool    `json:"need_upload_screen"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesGetBannedResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type StoriesGetBannedExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesGetByIDResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.StoriesStory `json:"items"`
	} `json:"response"`
}

type StoriesGetByIDExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.StoriesStory `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesGetPhotoUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
		PeerIDs   []int  `json:"peer_ids"`
		UserIDs   []int  `json:"user_ids"`
	} `json:"response"`
}

type StoriesGetRepliesResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.StoriesFeedItem `json:"items"`
	} `json:"response"`
}

type StoriesGetRepliesExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.StoriesFeedItem `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesGetStatsResponse struct {
	BaseResponse
	Response objects.StoriesStoryStats `json:"response"`
}

type StoriesGetVideoUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
		PeerIDs   []int  `json:"peer_ids"`
		UserIDs   []int  `json:"user_ids"`
	} `json:"response"`
}

type StoriesGetViewersResponse struct {
	BaseResponse
	Response struct {
		Count int                     `json:"count"`
		Items []objects.StoriesViewer `json:"items"`
	} `json:"response"`
}

type StoriesHideAllRepliesResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoriesHideReplyResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoriesSaveResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.StoriesStory `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.StoriesFeedItem `json:"items"`
	} `json:"response"`
}

type StoriesSearchExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.StoriesFeedItem `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type StoriesSendInteractionResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type StoriesUnbanOwnerResponse struct {
	BaseResponse
	Response int `json:"response"`
}
