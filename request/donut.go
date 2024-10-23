package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/donut
// https://dev.vk.com/ru/api/donut/getting-started

// DonutGetFriendsRequest defines the request for donut.getFriends
//
// The method returns a list of VK Donuts who are subscribed to specific communities from the user's friends.
// Doc: https://dev.vk.com/method/donut.getFriends
type DonutGetFriendsRequest struct {
	BaseRequest
}

// NewDonutGetFriendsRequest creates a new request for donut.getFriends
func NewDonutGetFriendsRequest(a *api.API, actor actor.Actor) *DonutGetFriendsRequest {
	return &DonutGetFriendsRequest{*NewMethodBaseRequest(a, actor, "donut.getFriends")}
}

// Exec executes the request and unmarshals the response into DonutGetFriendsResponse
func (r *DonutGetFriendsRequest) Exec(ctx context.Context) (response response.DonutGetFriendsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DonutGetSubscriptionRequest defines the request for donut.getSubscription
//
// The method returns information about a VK Donut subscription.
// Doc: https://dev.vk.com/method/donut.getSubscription
type DonutGetSubscriptionRequest struct {
	BaseRequest
}

// NewDonutGetSubscriptionRequest creates a new request for donut.getSubscription
func NewDonutGetSubscriptionRequest(a *api.API, actor actor.Actor) *DonutGetSubscriptionRequest {
	return &DonutGetSubscriptionRequest{*NewMethodBaseRequest(a, actor, "donut.getSubscription")}
}

// Exec executes the request and unmarshals the response into DonutGetSubscriptionResponse
func (r *DonutGetSubscriptionRequest) Exec(ctx context.Context) (response response.DonutGetSubscriptionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DonutGetSubscriptionsRequest defines the request for donut.getSubscriptions
//
// The method returns information about the user's VK Donut subscriptions.
// Doc: https://dev.vk.com/method/donut.getSubscriptions
type DonutGetSubscriptionsRequest struct {
	BaseRequest
}

// NewDonutGetSubscriptionsRequest creates a new request for donut.getSubscriptions
func NewDonutGetSubscriptionsRequest(a *api.API, actor actor.Actor) *DonutGetSubscriptionsRequest {
	return &DonutGetSubscriptionsRequest{*NewMethodBaseRequest(a, actor, "donut.getSubscriptions")}
}

// Exec executes the request and unmarshals the response into DonutGetSubscriptionsResponse
func (r *DonutGetSubscriptionsRequest) Exec(ctx context.Context) (response response.DonutGetSubscriptionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DonutIsDonRequest defines the request for donut.isDon
//
// The method returns information about whether a user is subscribed to paid content (is a VK Donut).
// Doc: https://dev.vk.com/method/donut.isDon
type DonutIsDonRequest struct {
	BaseRequest
}

// NewDonutIsDonRequest creates a new request for donut.isDon
func NewDonutIsDonRequest(a *api.API, actor actor.Actor) *DonutIsDonRequest {
	return &DonutIsDonRequest{*NewMethodBaseRequest(a, actor, "donut.isDon")}
}

// Exec executes the request and unmarshals the response into DonutIsDonResponse
func (r *DonutIsDonRequest) Exec(ctx context.Context) (response response.DonutIsDonResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
