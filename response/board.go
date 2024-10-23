package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/board

type BoardAddTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardCloseTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardCreateCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardDeleteCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardDeleteTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardEditCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardEditTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardFixTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardGetCommentsResponse struct {
	Count      int                         `json:"count"`
	Items      []objects.BoardTopicComment `json:"items"`
	Poll       objects.BoardTopicPoll      `json:"poll"`
	RealOffset int                         `json:"real_offset"`
}

type BoardGetCommentsExtendedResponse struct {
	Count      int                         `json:"count"`
	Items      []objects.BoardTopicComment `json:"items"`
	Poll       objects.BoardTopicPoll      `json:"poll"`
	RealOffset int                         `json:"real_offset"`
	Profiles   []objects.UserFull          `json:"profiles"`
	Groups     []objects.GroupFull         `json:"groups"`
}

type BoardGetTopicsResponse struct {
	Count        int                    `json:"count"`
	Items        []objects.BoardTopic   `json:"items"`
	DefaultOrder int                    `json:"default_order"`
	CanAddTopics objects.NumberFlagBool `json:"can_add_topics"`
}

type BoardGetTopicsExtendedResponse struct {
	Count        int                    `json:"count"`
	Items        []objects.BoardTopic   `json:"items"`
	DefaultOrder int                    `json:"default_order"`
	CanAddTopics objects.NumberFlagBool `json:"can_add_topics"`
	Profiles     []objects.UserFull     `json:"profiles"`
	Groups       []objects.GroupFull    `json:"groups"`
}

type BoardOpenTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardRestoreCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type BoardUnfixTopicResponse struct {
	BaseResponse
	Response int `json:"response"`
}
