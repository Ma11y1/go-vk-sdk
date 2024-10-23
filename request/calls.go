package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/calls

// CallsForceFinishRequest defines the request for calls.forceFinish
//
// Force call to end
// Doc: https://dev.vk.com/ru/method/calls.forceFinish
type CallsForceFinishRequest struct {
	BaseRequest
}

// NewCallsForceFinishRequest creates a new request for calls.forceFinish
func NewCallsForceFinishRequest(a *api.API, actor actor.Actor) *CallsForceFinishRequest {
	return &CallsForceFinishRequest{*NewMethodBaseRequest(a, actor, "calls.forceFinish")}
}

// Exec executes the request and unmarshals the response into CallsForceFinishRequest
func (r *CallsForceFinishRequest) Exec(ctx context.Context) (response response.CallsForceFinish, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// CallsStartRequest defines the request for calls.start
//
// Create a new call on behalf of a user or community
// Doc: https://dev.vk.com/ru/method/calls.start
type CallsStartRequest struct {
	BaseRequest
}

// NewCallsStartRequest creates a new request for calls.start
func NewCallsStartRequest(a *api.API, actor actor.Actor) *CallsStartRequest {
	return &CallsStartRequest{*NewMethodBaseRequest(a, actor, "calls.start")}
}

// Exec executes the request and unmarshals the response into CloseTopicResponse
func (r *CallsStartRequest) Exec(ctx context.Context) (response response.CallsStartResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
