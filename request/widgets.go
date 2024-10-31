package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/widgets

// WidgetsGetCommentsRequest defines the request for widgets.getComments
//
// Retrieves a list of comments added via a widget on external websites.
// Doc: https://dev.vk.com/method/widgets.getComments
type WidgetsGetCommentsRequest struct {
	BaseRequest
}

// NewWidgetsGetCommentsRequest creates a new request for widgets.getComments
func NewWidgetsGetCommentsRequest(a *api.API, actor actor.Actor) *WidgetsGetCommentsRequest {
	return &WidgetsGetCommentsRequest{*NewMethodBaseRequest(a, actor, "widgets.getComments")}
}

// Exec executes the request and unmarshals the response into WidgetsGetCommentsResponse
func (r *WidgetsGetCommentsRequest) Exec(ctx context.Context) (response response.WidgetsGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WidgetsGetPagesRequest defines the request for widgets.getPages
//
// Retrieves a list of pages on external websites that have installed widgets.
// Doc: https://dev.vk.com/method/widgets.getPages
type WidgetsGetPagesRequest struct {
	BaseRequest
}

// NewWidgetsGetPagesRequest creates a new request for widgets.getPages
func NewWidgetsGetPagesRequest(a *api.API, actor actor.Actor) *WidgetsGetPagesRequest {
	return &WidgetsGetPagesRequest{*NewMethodBaseRequest(a, actor, "widgets.getPages")}
}

// Exec executes the request and unmarshals the response into WidgetsGetPagesResponse
func (r *WidgetsGetPagesRequest) Exec(ctx context.Context) (response response.WidgetsGetPagesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
