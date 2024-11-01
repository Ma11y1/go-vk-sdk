package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Notifications Doc: https://dev.vk.com/ru/method/notifications
type Notifications struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/ru/method/notifications.get
func (n *Notifications) Get(actor actor.Actor) *request.NotificationsGetRequest {
	return request.NewNotificationsGetRequest(n.api, actor)
}

// MarkAsViewed Doc: https://dev.vk.com/ru/method/notifications.markAsViewed
func (n *Notifications) MarkAsViewed(actor actor.Actor) *request.NotificationsMarkAsViewedRequest {
	return request.NewNotificationsMarkAsViewedRequest(n.api, actor)
}

// SendMessage Doc: https://dev.vk.com/ru/method/notifications.sendMessage
func (n *Notifications) SendMessage(actor actor.Actor) *request.NotificationsSendMessageRequest {
	return request.NewNotificationsSendMessageRequest(n.api, actor)
}
