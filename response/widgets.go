package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/widgets

type WidgetsGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Posts []objects.WidgetsWidgetComment `json:"posts"`
	} `json:"response"`
}

type WidgetsGetPagesResponse struct {
	BaseResponse
	Response struct {
		Count int                         `json:"count"`
		Pages []objects.WidgetsWidgetPage `json:"pages"`
	} `json:"response"`
}
