package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/likes

// LikesAddRequest defines the request for likes.add
//
// Adds the specified object to the current user's list of Likes.
// Doc: https://dev.vk.com/method/likes.add
type LikesAddRequest struct {
	BaseRequest
}

// NewLikesAddRequest creates a new request for likes.add
func NewLikesAddRequest(a *api.API, actor actor.Actor) *LikesAddRequest {
	return &LikesAddRequest{*NewMethodBaseRequest(a, actor, "likes.add")}
}

// Exec executes the request and unmarshals the response into LikesAddResponse
func (r *LikesAddRequest) Exec(ctx context.Context) (response response.LikesAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LikesDeleteRequest defines the request for likes.delete
//
// Removes the specified object from the current user's list of Likes.
// Doc: https://dev.vk.com/method/likes.delete
type LikesDeleteRequest struct {
	BaseRequest
}

// NewLikesDeleteRequest creates a new request for likes.delete
func NewLikesDeleteRequest(a *api.API, actor actor.Actor) *LikesDeleteRequest {
	return &LikesDeleteRequest{*NewMethodBaseRequest(a, actor, "likes.delete")}
}

// Exec executes the request and unmarshals the response into LikesDeleteResponse
func (r *LikesDeleteRequest) Exec(ctx context.Context) (response response.LikesDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LikesGetListRequest defines the request for likes.getList
//
// Retrieves a list of user IDs who liked the specified object.
// Doc: https://dev.vk.com/method/likes.getList
type LikesGetListRequest struct {
	BaseRequest
}

// NewLikesGetListRequest creates a new request for likes.getList
func NewLikesGetListRequest(a *api.API, actor actor.Actor) *LikesGetListRequest {
	return &LikesGetListRequest{*NewMethodBaseRequest(a, actor, "likes.getList")}
}

// Exec executes the request and unmarshals the response into LikesGetListResponse
func (r *LikesGetListRequest) Exec(ctx context.Context) (response response.LikesGetListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LikesGetListExtendedRequest defines the request for likes.getList
//
// Retrieves a list of user IDs who liked the specified object.
// Doc: https://dev.vk.com/method/likes.getList
type LikesGetListExtendedRequest struct {
	BaseRequest
}

// NewLikesGetListExtendedRequest creates a new request for likes.getList
func NewLikesGetListExtendedRequest(a *api.API, actor actor.Actor) *LikesGetListExtendedRequest {
	r := &LikesGetListExtendedRequest{*NewMethodBaseRequest(a, actor, "likes.getList")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into LikesGetListExtendedResponse
func (r *LikesGetListExtendedRequest) Exec(ctx context.Context) (response response.LikesGetListExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LikesIsLikedRequest defines the request for likes.isLiked
//
// Checks if the specified object is in the list of Likes of the specified user.
// Doc: https://dev.vk.com/method/likes.isLiked
type LikesIsLikedRequest struct {
	BaseRequest
}

// NewLikesIsLikedRequest creates a new request for likes.isLiked
func NewLikesIsLikedRequest(a *api.API, actor actor.Actor) *LikesIsLikedRequest {
	return &LikesIsLikedRequest{*NewMethodBaseRequest(a, actor, "likes.isLiked")}
}

// Exec executes the request and unmarshals the response into LikesIsLikedResponse
func (r *LikesIsLikedRequest) Exec(ctx context.Context) (response response.LikesIsLikedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
