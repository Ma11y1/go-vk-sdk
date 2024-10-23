package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/notifications

type NotificationsGetResponse struct {
	BaseResponse
	Response struct {
		Count      int                                 `json:"count"`
		Items      []objects.NotificationsNotification `json:"items"`
		Profiles   []objects.UserFull                  `json:"profiles"`
		Groups     []objects.GroupFull                 `json:"groups"`
		Photos     []objects.Photo                     `json:"photos"`
		Videos     []objects.Video                     `json:"videos"`
		Apps       []objects.AppsApp                   `json:"apps"`
		LastViewed int                                 `json:"last_viewed"`
		NextFrom   string                              `json:"next_from"`
		TTL        int                                 `json:"ttl"`
	} `json:"response"`
}

type NotificationsMarkAsViewedResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type NotificationsSendMessageResponse struct {
	BaseResponse
	Response []struct {
		UserID int                    `json:"user_id"`
		Status objects.NumberFlagBool `json:"status"`
		Error  struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"error"`
	} `json:"response"`
}
