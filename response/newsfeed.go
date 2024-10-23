package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/newsfeed

type NewsfeedAddBanResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NewsfeedDeleteBanResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NewsfeedDeleteListResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NewsfeedGetResponse struct {
	BaseResponse
	Response struct {
		Items []objects.NewsfeedNewsfeedItem `json:"items"`
		objects.UsersAndGroups
		NextFrom string `json:"next_from"`
	} `json:"response"`
}

type NewsfeedGetBannedResponse struct {
	BaseResponse
	Response struct {
		Members []int `json:"members"`
		Groups  []int `json:"groups"`
	} `json:"response"`
}

type NewsfeedGetBannedExtendedResponse struct {
	BaseResponse
	Response struct {
		objects.UsersAndGroups
	} `json:"response"`
}

type NewsfeedGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Items []objects.NewsfeedNewsfeedItem `json:"items"`
		objects.UsersAndGroups
		NextFrom string `json:"next_from"`
	} `json:"response"`
}

type NewsfeedGetListsResponse struct {
	BaseResponse
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			NoReposts int    `json:"no_reposts"`
			SourceIDs []int  `json:"source_ids"`
		} `json:"items"`
	} `json:"response"`
}

type NewsfeedGetMentionsResponse struct {
	BaseResponse
	Response struct {
		Count int                        `json:"count"`
		Items []objects.WallWallpostToID `json:"items"`
	} `json:"response"`
}

type NewsfeedGetRecommendedResponse struct {
	BaseResponse
	Response struct {
		Items      []objects.NewsfeedNewsfeedItem `json:"items"`
		Profiles   []objects.UserFull             `json:"profiles"`
		Groups     []objects.GroupFull            `json:"groups"`
		NextOffset string                         `json:"next_offset"`
		NextFrom   string                         `json:"next_from"`
	} `json:"response"`
}

type NewsfeedGetSuggestedSourcesResponse struct {
	BaseResponse
	Response struct {
		Count int `json:"count"`
		Items []struct {
			objects.UserFull
			objects.GroupFull
		} `json:"items"`
	} `json:"response"`
}

type NewsfeedIgnoreItemResponse struct {
	BaseResponse
	Response struct {
		Status  bool    `json:"status"`
		Message *string `json:"message"`
	} `json:"response"`
}

type NewsfeedSaveListResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NewsfeedSearchResponse struct {
	BaseResponse
	Response struct {
		Items      []objects.WallWallpost `json:"items"`
		Count      int                    `json:"count"`
		TotalCount int                    `json:"total_count"`
		NextFrom   string                 `json:"next_from"`
	} `json:"response"`
}

type NewsfeedSearchExtendedResponse struct {
	BaseResponse
	Response struct {
		Items      []objects.WallWallpost `json:"items"`
		Count      int                    `json:"count"`
		TotalCount int                    `json:"total_count"`
		Profiles   []objects.UserFull     `json:"profiles"`
		Groups     []objects.GroupFull    `json:"groups"`
		NextFrom   string                 `json:"next_from"`
	} `json:"response"`
}

type NewsfeedUnignoreItemResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NewsfeedUnsubscribeResponse struct {
	BaseResponse
	Response int `json:"response"`
}
