package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/friends

type FriendsAddResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FriendsAddListResponse struct {
	BaseResponse
	Response struct {
		ListID int `json:"list_id"`
	} `json:"response"`
}

type FriendsAreFriendsResponse struct {
	BaseResponse
	Response []objects.FriendsStatus `json:"response"`
}

type FriendsDeleteResponse struct {
	BaseResponse
	Response struct {
		Success           objects.NumberFlagBool `json:"success"`
		FriendDeleted     objects.NumberFlagBool `json:"friend_deleted"`
		OutRequestDeleted objects.NumberFlagBool `json:"out_request_deleted"`
		InRequestDeleted  objects.NumberFlagBool `json:"in_request_deleted"`
		SuggestionDeleted objects.NumberFlagBool `json:"suggestion_deleted"`
	} `json:"response"`
}

type FriendsDeleteAllRequestsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FriendsDeleteListResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FriendsEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FriendsEditListResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type FriendsGetResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type FriendsGetFieldsResponse struct {
	BaseResponse
	Response struct {
		Count int                           `json:"count"`
		Items []objects.FriendsUserXtrLists `json:"items"`
	} `json:"response"`
}

type FriendsGetAppUsersResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type FriendsGetByPhonesResponse struct {
	BaseResponse
	Response []objects.FriendsUserXtrPhone `json:"response"`
}

type FriendsGetListsResponse struct {
	BaseResponse
	Response struct {
		Count int                   `json:"count"`
		Items []objects.FriendsList `json:"items"`
	} `json:"response"`
}

type FriendsGetMutualResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type FriendsGetOnlineResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type FriendsGetOnlineOnlineMobileResponse struct {
	BaseResponse
	Response struct {
		Online       []int `json:"online"`
		OnlineMobile []int `json:"online_mobile"`
	} `json:"response"`
}

type FriendsGetRecentResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type FriendsGetRequestsResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"` // Total requests number
		Items []int `json:"items"`
	} `json:"response"`
}

type FriendsGetRequestsNeedMutualResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"` // Total requests number
		Items []objects.FriendsRequests `json:"items"`
	} `json:"response"`
}

type FriendsGetRequestsExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                                 `json:"count"`
		Items []objects.FriendsRequestsXtrMessage `json:"items"`
	} `json:"response"`
}

type FriendsGetSuggestionsResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type FriendsSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}
