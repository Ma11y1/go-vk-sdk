package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/notifications

// NotificationsGetRequest defines the request for notifications.get
//
// Returns a list of alerts about other users' replies to the current user's posts.
// Doc: https://dev.vk.com/method/notifications.get
type NotificationsGetRequest struct {
	BaseRequest
}

// NewNotificationsGetRequest creates a new request for notifications.get
func NewNotificationsGetRequest(a *api.API, actor actor.Actor) *NotificationsGetRequest {
	return &NotificationsGetRequest{*NewMethodBaseRequest(a, actor, "notifications.get")}
}

// Exec executes the request and unmarshals the response into NotificationsGetResponse
func (r *NotificationsGetRequest) Exec(ctx context.Context) (response response.NotificationsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotificationsMarkAsViewedRequest defines the request for notifications.markAsViewed
//
// Resets the counter of unviewed notifications about other users' replies to the current user's posts.
// Doc: https://dev.vk.com/method/notifications.markAsViewed
type NotificationsMarkAsViewedRequest struct {
	BaseRequest
}

// NewNotificationsMarkAsViewedRequest creates a new request for notifications.markAsViewed
func NewNotificationsMarkAsViewedRequest(a *api.API, actor actor.Actor) *NotificationsMarkAsViewedRequest {
	return &NotificationsMarkAsViewedRequest{*NewMethodBaseRequest(a, actor, "notifications.markAsViewed")}
}

// Exec executes the request and unmarshals the response into NotificationsMarkAsViewedResponse
func (r *NotificationsMarkAsViewedRequest) Exec(ctx context.Context) (response response.NotificationsMarkAsViewedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotificationsSendMessageRequest defines the request for notifications.sendMessage
//
//	The method sends a notification to the widget user.
//	Doc: https://dev.vk.com/method/notifications.sendMessage
type NotificationsSendMessageRequest struct {
	BaseRequest
}

// NewNotificationsSendMessageRequest creates a new request for notifications.sendMessage
func NewNotificationsSendMessageRequest(a *api.API, actor actor.Actor) *NotificationsSendMessageRequest {
	return &NotificationsSendMessageRequest{*NewMethodBaseRequest(a, actor, "notifications.sendMessage")}
}

// Exec executes the request and unmarshals the response into NotificationsMarkAsViewedResponse
func (r *NotificationsSendMessageRequest) Exec(ctx context.Context) (response response.NotificationsSendMessageResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
