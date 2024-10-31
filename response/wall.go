package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/wall

type WallCheckCopyrightLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallCloseCommentsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallCreateCommentResponse struct {
	BaseResponse
	Response struct {
		CommentID    int   `json:"comment_id"`
		ParentsStack []int `json:"parents_stack"`
	} `json:"response"`
}

type WallDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallDeleteCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallEditResponse struct {
	BaseResponse
	Response struct {
		PostID int `json:"post_id"`
	} `json:"response"`
}

type WallEditAdsStealthResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallEditCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallGetResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.WallWallpost `json:"items"`
	} `json:"response"`
}

type WallGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.WallWallpost `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type WallGetByIDResponse struct {
	BaseResponse
	Response struct {
		Items []objects.WallWallpost `json:"items"`
	} `json:"response"`
}

type WallGetByIDExtendedResponse struct {
	BaseResponse
	Response struct {
		Items []objects.WallWallpost `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type WallGetCommentResponse struct {
	BaseResponse
	Response struct {
		Items             []objects.WallWallComment `json:"items"`
		CanPost           objects.BoolInt           `json:"can_post"`
		ShowReplyButton   objects.BoolInt           `json:"show_reply_button"`
		GroupsCanPost     objects.BoolInt           `json:"groups_can_post"`
		CurrentLevelCount int                       `json:"current_level_count"`
	} `json:"response"`
}

type WallGetCommentExtendedResponse struct {
	BaseResponse
	Response struct {
		Count             int                       `json:"count"`
		Items             []objects.WallWallComment `json:"items"`
		CanPost           objects.BoolInt           `json:"can_post"`
		ShowReplyButton   objects.BoolInt           `json:"show_reply_button"`
		GroupsCanPost     objects.BoolInt           `json:"groups_can_post"`
		CurrentLevelCount int                       `json:"current_level_count"`
		Profiles          []objects.UserFull        `json:"profiles"`
		Groups            []objects.GroupFull       `json:"groups"`
	} `json:"response"`
}

type WallGetCommentsResponse struct {
	BaseResponse
	Response struct {
		CanPost           objects.BoolInt           `json:"can_post"`
		ShowReplyButton   objects.BoolInt           `json:"show_reply_button"`
		GroupsCanPost     objects.BoolInt           `json:"groups_can_post"`
		CurrentLevelCount int                       `json:"current_level_count"`
		Count             int                       `json:"count"`
		Items             []objects.WallWallComment `json:"items"`
	} `json:"response"`
}

type WallGetCommentsExtendedResponse struct {
	BaseResponse
	Response struct {
		CanPost           objects.BoolInt           `json:"can_post"`
		ShowReplyButton   objects.BoolInt           `json:"show_reply_button"`
		GroupsCanPost     objects.BoolInt           `json:"groups_can_post"`
		CurrentLevelCount int                       `json:"current_level_count"`
		Count             int                       `json:"count"`
		Items             []objects.WallWallComment `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type WallGetRepostsResponse struct {
	BaseResponse
	Response struct {
		Items []objects.WallWallpost `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type WallOpenCommentsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallParseAttachedLinkResponse struct {
	BaseResponse
	Response objects.WallParseAttachedLink `json:"response"`
}

type WallPinResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallPostResponse struct {
	BaseResponse
	Response struct {
		PostID int `json:"post_id"`
	} `json:"response"`
}

type WallPostAdsStealthResponse struct {
	BaseResponse
	Response struct {
		PostID int `json:"post_id"`
	} `json:"response"`
}

type WallReportCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallReportPostResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallRepostResponse struct {
	BaseResponse
	Response struct {
		Success         int `json:"success"`
		PostID          int `json:"post_id"`
		RepostsCount    int `json:"reposts_count"`
		LikesCount      int `json:"likes_count"`
		WallRepostCount int `json:"wall_repost_count"`
		MailRepostCount int `json:"mail_repost_count"`
	} `json:"response"`
}

type WallRestoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallRestoreCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type WallSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.WallWallpost `json:"items"`
	} `json:"response"`
}

type WallSearchExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.WallWallpost `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type WallUnpinResponse struct {
	BaseResponse
	Response int `json:"response"`
}
