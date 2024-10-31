package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/status

// StatusGetRequest defines the request for status.get
//
// Retrieves the status text of a user or community.
// Doc: https://dev.vk.com/method/status.get
type StatusGetRequest struct {
	BaseRequest
}

// NewStatusGetRequest creates a new request for status.get
func NewStatusGetRequest(a *api.API, actor actor.Actor) *StatusGetRequest {
	return &StatusGetRequest{*NewMethodBaseRequest(a, actor, "status.get")}
}

// Exec executes the request and unmarshals the response into StatusGetResponse
func (r *StatusGetRequest) Exec(ctx context.Context) (response response.StatusGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StatusSetRequest defines the request for status.set
//
// Sets a new status for the current user or community.
// Doc: https://dev.vk.com/method/status.set
type StatusSetRequest struct {
	BaseRequest
}

// NewStatusSetRequest creates a new request for status.set
func NewStatusSetRequest(a *api.API, actor actor.Actor) *StatusSetRequest {
	return &StatusSetRequest{*NewMethodBaseRequest(a, actor, "status.set")}
}

// Exec executes the request and unmarshals the response into StatusSetResponse
func (r *StatusSetRequest) Exec(ctx context.Context) (response response.StatusSetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
