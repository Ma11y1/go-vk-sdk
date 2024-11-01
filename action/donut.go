package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Donut Doc: https://dev.vk.com/ru/method/donut
type Donut struct {
	BaseAction
}

// GetFriends Doc: https://dev.vk.com/ru/method/donut.getFriends
func (d *Donut) GetFriends(actor actor.Actor) *request.DonutGetFriendsRequest {
	return request.NewDonutGetFriendsRequest(d.api, actor)
}

// GetSubscription Doc: https://dev.vk.com/ru/method/donut.getSubscription
func (d *Donut) GetSubscription(actor actor.Actor) *request.DonutGetSubscriptionRequest {
	return request.NewDonutGetSubscriptionRequest(d.api, actor)
}

// GetSubscriptions Doc: https://dev.vk.com/ru/method/donut.getSubscriptions
func (d *Donut) GetSubscriptions(actor actor.Actor) *request.DonutGetSubscriptionsRequest {
	return request.NewDonutGetSubscriptionsRequest(d.api, actor)
}

// IsDon Doc: https://dev.vk.com/ru/method/donut.isDon
func (d *Donut) IsDon(actor actor.Actor) *request.DonutIsDonRequest {
	return request.NewDonutIsDonRequest(d.api, actor)
}
