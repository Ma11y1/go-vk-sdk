package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/orders

// OrdersCancelSubscriptionRequest defines the request for orders.cancelSubscription
//
// Cancels a subscription.
// Doc: https://dev.vk.com/method/orders.cancelSubscription
type OrdersCancelSubscriptionRequest struct {
	BaseRequest
}

// NewOrdersCancelSubscriptionRequest creates a new request for orders.cancelSubscription
func NewOrdersCancelSubscriptionRequest(a *api.API, actor actor.Actor) *OrdersCancelSubscriptionRequest {
	return &OrdersCancelSubscriptionRequest{*NewMethodBaseRequest(a, actor, "orders.cancelSubscription")}
}

// Exec executes the request and unmarshals the response into OrdersCancelSubscriptionResponse
func (r *OrdersCancelSubscriptionRequest) Exec(ctx context.Context) (response response.OrdersCancelSubscriptionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersChangeStateRequest defines the request for orders.changeState
//
// Changes the state of an order.
// Doc: https://dev.vk.com/method/orders.changeState
type OrdersChangeStateRequest struct {
	BaseRequest
}

// NewOrdersChangeStateRequest creates a new request for orders.changeState
func NewOrdersChangeStateRequest(a *api.API, actor actor.Actor) *OrdersChangeStateRequest {
	return &OrdersChangeStateRequest{*NewMethodBaseRequest(a, actor, "orders.changeState")}
}

// Exec executes the request and unmarshals the response into OrdersChangeStateResponse
func (r *OrdersChangeStateRequest) Exec(ctx context.Context) (response response.OrdersChangeStateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersGetRequest defines the request for orders.get
//
// Returns a list of orders.
// Doc: https://dev.vk.com/method/orders.get
type OrdersGetRequest struct {
	BaseRequest
}

// NewOrdersGetRequest creates a new request for orders.get
func NewOrdersGetRequest(a *api.API, actor actor.Actor) *OrdersGetRequest {
	return &OrdersGetRequest{*NewMethodBaseRequest(a, actor, "orders.get")}
}

// Exec executes the request and unmarshals the response into OrdersGetResponse
func (r *OrdersGetRequest) Exec(ctx context.Context) (response response.OrdersGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersGetAmountRequest defines the request for orders.getAmount
//
// Returns the cost of votes in the user's currency.
// Doc: https://dev.vk.com/method/orders.getAmount
type OrdersGetAmountRequest struct {
	BaseRequest
}

// NewOrdersGetAmountRequest creates a new request for orders.getAmount
func NewOrdersGetAmountRequest(a *api.API, actor actor.Actor) *OrdersGetAmountRequest {
	return &OrdersGetAmountRequest{*NewMethodBaseRequest(a, actor, "orders.getAmount")}
}

// Exec executes the request and unmarshals the response into OrdersGetAmountResponse
func (r *OrdersGetAmountRequest) Exec(ctx context.Context) (response response.OrdersGetAmountResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersGetByIDRequest defines the request for orders.getById
//
// Returns information about a specific order.
// Doc: https://dev.vk.com/method/orders.getById
type OrdersGetByIDRequest struct {
	BaseRequest
}

// NewOrdersGetByIDRequest creates a new request for orders.getById
func NewOrdersGetByIDRequest(a *api.API, actor actor.Actor) *OrdersGetByIDRequest {
	return &OrdersGetByIDRequest{*NewMethodBaseRequest(a, actor, "orders.getById")}
}

// Exec executes the request and unmarshals the response into OrdersGetByIdResponse
func (r *OrdersGetByIDRequest) Exec(ctx context.Context) (response response.OrdersGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersGetUserSubscriptionByIDRequest defines the request for orders.getUserSubscriptionById
//
// Gets subscription information by its ID.
// Doc: https://dev.vk.com/method/orders.getUserSubscriptionById
type OrdersGetUserSubscriptionByIDRequest struct {
	BaseRequest
}

// NewOrdersGetUserSubscriptionByIDRequest creates a new request for orders.getUserSubscriptionById
func NewOrdersGetUserSubscriptionByIDRequest(a *api.API, actor actor.Actor) *OrdersGetUserSubscriptionByIDRequest {
	return &OrdersGetUserSubscriptionByIDRequest{*NewMethodBaseRequest(a, actor, "orders.getUserSubscriptionById")}
}

// Exec executes the request and unmarshals the response into OrdersGetUserSubscriptionByIDResponse
func (r *OrdersGetUserSubscriptionByIDRequest) Exec(ctx context.Context) (response response.OrdersGetUserSubscriptionByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OrdersGetUserSubscriptionsRequest defines the request for orders.getUserSubscriptions
//
// Gets a list of active subscriptions of the user.
// Doc: https://dev.vk.com/method/orders.getUserSubscriptions
type OrdersGetUserSubscriptionsRequest struct {
	BaseRequest
}

// NewOrdersGetUserSubscriptionsRequest creates a new request for orders.getUserSubscriptions
func NewOrdersGetUserSubscriptionsRequest(a *api.API, actor actor.Actor) *OrdersGetUserSubscriptionsRequest {
	return &OrdersGetUserSubscriptionsRequest{*NewMethodBaseRequest(a, actor, "orders.getUserSubscriptions")}
}

// Exec executes the request and unmarshals the response into OrdersGetUserSubscriptionsResponse
func (r *OrdersGetUserSubscriptionsRequest) Exec(ctx context.Context) (response response.OrdersGetUserSubscriptionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
