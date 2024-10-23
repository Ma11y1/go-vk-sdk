package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/apps

type AppsAddSnippetResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsAddUsersToTestingGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsDeleteAppRequestsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsDeleteSnippetResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsGetResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.AppsApp `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type AppsGetCatalogResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.AppsApp `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type AppsGetFriendsListResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type AppsGetFriendsListExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type AppsGetLeaderboardResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.AppsLeaderboard `json:"items"`
	} `json:"response"`
}

type AppsGetLeaderboardExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Score  int `json:"score"`
			UserID int `json:"user_id"`
		} `json:"items"`
		Profiles []objects.UserFull `json:"profiles"`
	} `json:"response"`
}

type AppsGetMiniAppPoliciesResponse struct {
	BaseResponse
	Response struct {
		PrivacyPolicy string `json:"privacy_policy"`
		Term          string `json:"term"`
	} `json:"response"`
}

type AppsGetScopesResponse struct {
	BaseResponse
	Response struct {
		Count int                 `json:"count"`
		Items []objects.AppsScope `json:"items"`
	} `json:"response"`
}

type AppsGetScoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsGetSnippetsResponse struct {
	BaseResponse
	Response struct {
		Items []struct {
			VkRef         []string `json:"vk_ref"`
			GroupID       []int    `json:"group_id"`
			Hash          []string `json:"hash"`
			SnippetID     int      `json:"snippet_id"`
			Title         string   `json:"title,omitempty"`
			Description   string   `json:"description,omitempty"`
			ImageURL      string   `json:"image_url,omitempty"`
			SmallImageURL string   `json:"small_image_url,omitempty"`
			Button        string   `json:"button,omitempty"`
		} `json:"items"`
	} `json:"response"`
}

type AppsGetTestingGroupsResponse struct {
	BaseResponse
	Response []objects.AppsTestingGroup `json:"response"`
}

type AppsIsNotificationsAllowedResponse struct {
	BaseResponse
	Response struct {
		IsAllowed bool `json:"is_allowed"`
	} `json:"response"`
}

type AppsPromoHasActiveGiftResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsPromoUseGiftResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsRemoveTestingGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsRemoveUsersFromTestingGroupsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsSendRequestResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppsUpdateMetaForTestingGroupResponse struct {
	BaseResponse
	Response struct {
		GroupID int `json:"group_id"`
	} `json:"response"`
}

type AppUploadWidgetsImageResponse struct {
	Image string `json:"image"`
	Hash  string `json:"hash"`
}

type AppUploadWidgetsGroupImageResponse struct {
	Image string `json:"image"`
	Hash  string `json:"hash"`
}
