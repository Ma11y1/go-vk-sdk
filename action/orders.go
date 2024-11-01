package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Orders Doc: https://dev.vk.com/ru/method/orders
type Orders struct {
	BaseAction
}

// CancelSubscription Doc: https://dev.vk.com/ru/method/orders.cancelSubscription
func (o *Orders) CancelSubscription(actor actor.Actor) *request.OrdersCancelSubscriptionRequest {
	return request.NewOrdersCancelSubscriptionRequest(o.api, actor)
}

// ChangeState Doc: https://dev.vk.com/ru/method/orders.changeState
func (o *Orders) ChangeState(actor actor.Actor) *request.OrdersChangeStateRequest {
	return request.NewOrdersChangeStateRequest(o.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/orders.get
func (o *Orders) Get(actor actor.Actor) *request.OrdersGetRequest {
	return request.NewOrdersGetRequest(o.api, actor)
}

// GetAmount Doc: https://dev.vk.com/ru/method/orders.getAmount
func (o *Orders) GetAmount(actor actor.Actor) *request.OrdersGetAmountRequest {
	return request.NewOrdersGetAmountRequest(o.api, actor)
}

// GetById Doc: https://dev.vk.com/ru/method/orders.getById
func (o *Orders) GetById(actor actor.Actor) *request.OrdersGetByIDRequest {
	return request.NewOrdersGetByIDRequest(o.api, actor)
}

// GetUserSubscriptionById Doc: https://dev.vk.com/ru/method/orders.getUserSubscriptionById
func (o *Orders) GetUserSubscriptionById(actor actor.Actor) *request.OrdersGetUserSubscriptionByIDRequest {
	return request.NewOrdersGetUserSubscriptionByIDRequest(o.api, actor)
}

// GetUserSubscriptions Doc: https://dev.vk.com/ru/method/orders.getUserSubscriptions
func (o *Orders) GetUserSubscriptions(actor actor.Actor) *request.OrdersGetUserSubscriptionsRequest {
	return request.NewOrdersGetUserSubscriptionsRequest(o.api, actor)
}
