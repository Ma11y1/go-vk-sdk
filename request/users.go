package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// UsersGetRequest defines the request for users.get
//
// Retrieves information about users.
// Doc: https://dev.vk.com/method/users.get
type UsersGetRequest struct {
	BaseRequest
}

// NewUsersGetRequest creates a new request for users.get
func NewUsersGetRequest(a *api.API, actor actor.Actor) *UsersGetRequest {
	return &UsersGetRequest{*NewMethodBaseRequest(a, actor, "users.get")}
}

// Exec executes the request and unmarshals the response into UsersGetResponse
func (r *UsersGetRequest) Exec(ctx context.Context) (response response.UsersGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UsersGetFollowersRequest defines the request for users.getFollowers
//
// Retrieves a list of followers of a user.
// Doc: https://dev.vk.com/method/users.getFollowers
type UsersGetFollowersRequest struct {
	BaseRequest
}

// NewUsersGetFollowersRequest creates a new request for users.getFollowers
func NewUsersGetFollowersRequest(a *api.API, actor actor.Actor) *UsersGetFollowersRequest {
	return &UsersGetFollowersRequest{*NewMethodBaseRequest(a, actor, "users.getFollowers")}
}

// Exec executes the request and unmarshals the response into UsersGetFollowersResponse
func (r *UsersGetFollowersRequest) Exec(ctx context.Context) (response response.UsersGetFollowersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UsersGetSubscriptionsRequest defines the request for users.getSubscriptions
//
// Retrieves a list of subscriptions of a user.
// Doc: https://dev.vk.com/method/users.getSubscriptions
type UsersGetSubscriptionsRequest struct {
	BaseRequest
}

// NewUsersGetSubscriptionsRequest creates a new request for users.getSubscriptions
func NewUsersGetSubscriptionsRequest(a *api.API, actor actor.Actor) *UsersGetSubscriptionsRequest {
	return &UsersGetSubscriptionsRequest{*NewMethodBaseRequest(a, actor, "users.getSubscriptions")}
}

// Exec executes the request and unmarshals the response into UsersGetSubscriptionsResponse
func (r *UsersGetSubscriptionsRequest) Exec(ctx context.Context) (response response.UsersGetSubscriptionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UsersReportRequest defines the request for users.report
//
// Reports a user for inappropriate behavior or spam.
// Doc: https://dev.vk.com/method/users.report
type UsersReportRequest struct {
	BaseRequest
}

// NewUsersReportRequest creates a new request for users.report
func NewUsersReportRequest(a *api.API, actor actor.Actor) *UsersReportRequest {
	return &UsersReportRequest{*NewMethodBaseRequest(a, actor, "users.report")}
}

// Exec executes the request and unmarshals the response into UsersReportResponse
func (r *UsersReportRequest) Exec(ctx context.Context) (response response.UsersReportResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UsersSearchRequest defines the request for users.search
//
// Searches for users by specified criteria.
// Doc: https://dev.vk.com/method/users.search
type UsersSearchRequest struct {
	BaseRequest
}

// NewUsersSearchRequest creates a new request for users.search
func NewUsersSearchRequest(a *api.API, actor actor.Actor) *UsersSearchRequest {
	return &UsersSearchRequest{*NewMethodBaseRequest(a, actor, "users.search")}
}

// Exec executes the request and unmarshals the response into UsersSearchResponse
func (r *UsersSearchRequest) Exec(ctx context.Context) (response response.UsersSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
