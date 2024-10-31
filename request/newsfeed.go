package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/newsfeed

// NewsfeedAddBanRequest defines the request for newsfeed.addBan
//
// Blocks specified users and groups from appearing in the newsfeed of the current user.
// Doc: https://dev.vk.com/method/newsfeed.addBan
type NewsfeedAddBanRequest struct {
	BaseRequest
}

// NewNewsfeedAddBanRequest creates a new request for newsfeed.addBan
func NewNewsfeedAddBanRequest(a *api.API, actor actor.Actor) *NewsfeedAddBanRequest {
	return &NewsfeedAddBanRequest{*NewMethodBaseRequest(a, actor, "newsfeed.addBan")}
}

// Exec executes the request and unmarshals the response into NewsfeedAddBanResponse
func (r *NewsfeedAddBanRequest) Exec(ctx context.Context) (response response.NewsfeedAddBanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedDeleteBanRequest defines the request for newsfeed.deleteBan
//
// Allows specified users and groups to reappear in the newsfeed of the current user.
// Doc: https://dev.vk.com/method/newsfeed.deleteBan
type NewsfeedDeleteBanRequest struct {
	BaseRequest
}

// NewNewsfeedDeleteBanRequest creates a new request for newsfeed.deleteBan
func NewNewsfeedDeleteBanRequest(a *api.API, actor actor.Actor) *NewsfeedDeleteBanRequest {
	return &NewsfeedDeleteBanRequest{*NewMethodBaseRequest(a, actor, "newsfeed.deleteBan")}
}

// Exec executes the request and unmarshals the response into NewsfeedDeleteBanResponse
func (r *NewsfeedDeleteBanRequest) Exec(ctx context.Context) (response response.NewsfeedDeleteBanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedDeleteListRequest defines the request for newsfeed.deleteList
//
// Deletes a custom newsfeed list.
// Doc: https://dev.vk.com/method/newsfeed.deleteList
type NewsfeedDeleteListRequest struct {
	BaseRequest
}

// NewNewsfeedDeleteListRequest creates a new request for newsfeed.deleteList
func NewNewsfeedDeleteListRequest(a *api.API, actor actor.Actor) *NewsfeedDeleteListRequest {
	return &NewsfeedDeleteListRequest{*NewMethodBaseRequest(a, actor, "newsfeed.deleteList")}
}

// Exec executes the request and unmarshals the response into NewsfeedDeleteListResponse
func (r *NewsfeedDeleteListRequest) Exec(ctx context.Context) (response response.NewsfeedDeleteListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetRequest defines the request for newsfeed.get
//
// Returns data necessary to display the current user's newsfeed.
// Doc: https://dev.vk.com/method/newsfeed.get
type NewsfeedGetRequest struct {
	BaseRequest
}

// NewNewsfeedGetRequest creates a new request for newsfeed.get
func NewNewsfeedGetRequest(a *api.API, actor actor.Actor) *NewsfeedGetRequest {
	return &NewsfeedGetRequest{*NewMethodBaseRequest(a, actor, "newsfeed.get")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetResponse
func (r *NewsfeedGetRequest) Exec(ctx context.Context) (response response.NewsfeedGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetBannedRequest defines the request for newsfeed.getBanned
//
// Returns a list of users and groups hidden from the current user's newsfeed.
// Doc: https://dev.vk.com/method/newsfeed.getBanned
type NewsfeedGetBannedRequest struct {
	BaseRequest
}

// NewNewsfeedGetBannedRequest creates a new request for newsfeed.getBanned
func NewNewsfeedGetBannedRequest(a *api.API, actor actor.Actor) *NewsfeedGetBannedRequest {
	return &NewsfeedGetBannedRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getBanned")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetBannedResponse
func (r *NewsfeedGetBannedRequest) Exec(ctx context.Context) (response response.NewsfeedGetBannedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetCommentsRequest defines the request for newsfeed.getComments
//
// Returns data necessary to display the comments section of the user's newsfeed.
// Doc: https://dev.vk.com/method/newsfeed.getComments
type NewsfeedGetCommentsRequest struct {
	BaseRequest
}

// NewNewsfeedGetCommentsRequest creates a new request for newsfeed.getComments
func NewNewsfeedGetCommentsRequest(a *api.API, actor actor.Actor) *NewsfeedGetCommentsRequest {
	return &NewsfeedGetCommentsRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getComments")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetCommentsResponse
func (r *NewsfeedGetCommentsRequest) Exec(ctx context.Context) (response response.NewsfeedGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetListsRequest defines the request for newsfeed.getLists
//
// Returns custom newsfeed lists.
// Doc: https://dev.vk.com/method/newsfeed.getLists
type NewsfeedGetListsRequest struct {
	BaseRequest
}

// NewNewsfeedGetListsRequest creates a new request for newsfeed.getLists
func NewNewsfeedGetListsRequest(a *api.API, actor actor.Actor) *NewsfeedGetListsRequest {
	return &NewsfeedGetListsRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getLists")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetListsResponse
func (r *NewsfeedGetListsRequest) Exec(ctx context.Context) (response response.NewsfeedGetListsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetMentionsRequest defines the request for newsfeed.getMentions
//
// Returns a list of posts on user walls where the specified user is mentioned.
// Doc: https://dev.vk.com/method/newsfeed.getMentions
type NewsfeedGetMentionsRequest struct {
	BaseRequest
}

// NewNewsfeedGetMentionsRequest creates a new request for newsfeed.getMentions
func NewNewsfeedGetMentionsRequest(a *api.API, actor actor.Actor) *NewsfeedGetMentionsRequest {
	return &NewsfeedGetMentionsRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getMentions")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetMentionsResponse
func (r *NewsfeedGetMentionsRequest) Exec(ctx context.Context) (response response.NewsfeedGetMentionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetRecommendedRequest defines the request for newsfeed.getRecommended
//
// Retrieves recommended newsfeed items for the user.
// Doc: https://dev.vk.com/method/newsfeed.getRecommended
type NewsfeedGetRecommendedRequest struct {
	BaseRequest
}

// NewNewsfeedGetRecommendedRequest creates a new request for newsfeed.getRecommended
func NewNewsfeedGetRecommendedRequest(a *api.API, actor actor.Actor) *NewsfeedGetRecommendedRequest {
	return &NewsfeedGetRecommendedRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getRecommended")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetRecommendedResponse
func (r *NewsfeedGetRecommendedRequest) Exec(ctx context.Context) (response response.NewsfeedGetRecommendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedGetSuggestedSourcesRequest defines the request for newsfeed.getSuggestedSources
//
// Returns communities and users suggested for the user to follow.
// Doc: https://dev.vk.com/method/newsfeed.getSuggestedSources
type NewsfeedGetSuggestedSourcesRequest struct {
	BaseRequest
}

// NewNewsfeedGetSuggestedSourcesRequest creates a new request for newsfeed.getSuggestedSources
func NewNewsfeedGetSuggestedSourcesRequest(a *api.API, actor actor.Actor) *NewsfeedGetSuggestedSourcesRequest {
	return &NewsfeedGetSuggestedSourcesRequest{*NewMethodBaseRequest(a, actor, "newsfeed.getSuggestedSources")}
}

// Exec executes the request and unmarshals the response into NewsfeedGetSuggestedSourcesResponse
func (r *NewsfeedGetSuggestedSourcesRequest) Exec(ctx context.Context) (response response.NewsfeedGetSuggestedSourcesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedIgnoreItemRequest defines the request for newsfeed.ignoreItem
//
// Allows hiding an item from the newsfeed.
// Doc: https://dev.vk.com/method/newsfeed.ignoreItem
type NewsfeedIgnoreItemRequest struct {
	BaseRequest
}

// NewNewsfeedIgnoreItemRequest creates a new request for newsfeed.ignoreItem
func NewNewsfeedIgnoreItemRequest(a *api.API, actor actor.Actor) *NewsfeedIgnoreItemRequest {
	return &NewsfeedIgnoreItemRequest{*NewMethodBaseRequest(a, actor, "newsfeed.ignoreItem")}
}

// Exec executes the request and unmarshals the response into NewsfeedIgnoreItemResponse
func (r *NewsfeedIgnoreItemRequest) Exec(ctx context.Context) (response response.NewsfeedIgnoreItemResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedSaveListRequest defines the request for newsfeed.saveList
//
// Creates or edits custom newsfeed lists for viewing.
// Doc: https://dev.vk.com/method/newsfeed.saveList
type NewsfeedSaveListRequest struct {
	BaseRequest
}

// NewNewsfeedSaveListRequest creates a new request for newsfeed.saveList
func NewNewsfeedSaveListRequest(a *api.API, actor actor.Actor) *NewsfeedSaveListRequest {
	return &NewsfeedSaveListRequest{*NewMethodBaseRequest(a, actor, "newsfeed.saveList")}
}

// Exec executes the request and unmarshals the response into NewsfeedSaveListResponse
func (r *NewsfeedSaveListRequest) Exec(ctx context.Context) (response response.NewsfeedSaveListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedSearchRequest defines the request for newsfeed.search
//
// Returns search results by statuses, in descending order from newest to oldest.
// Doc: https://dev.vk.com/method/newsfeed.search
type NewsfeedSearchRequest struct {
	BaseRequest
}

// NewNewsfeedSearchRequest creates a new request for newsfeed.search
func NewNewsfeedSearchRequest(a *api.API, actor actor.Actor) *NewsfeedSearchRequest {
	return &NewsfeedSearchRequest{*NewMethodBaseRequest(a, actor, "newsfeed.search")}
}

// Exec executes the request and unmarshals the response into NewsfeedSearchResponse
func (r *NewsfeedSearchRequest) Exec(ctx context.Context) (response response.NewsfeedSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedUnignoreItemRequest defines the request for newsfeed.unignoreItem
//
// Allows previously hidden item to reappear in the newsfeed.
// Doc: https://dev.vk.com/method/newsfeed.unignoreItem
type NewsfeedUnignoreItemRequest struct {
	BaseRequest
}

// NewNewsfeedUnignoreItemRequest creates a new request for newsfeed.unignoreItem
func NewNewsfeedUnignoreItemRequest(a *api.API, actor actor.Actor) *NewsfeedUnignoreItemRequest {
	return &NewsfeedUnignoreItemRequest{*NewMethodBaseRequest(a, actor, "newsfeed.unignoreItem")}
}

// Exec executes the request and unmarshals the response into NewsfeedUnignoreItemResponse
func (r *NewsfeedUnignoreItemRequest) Exec(ctx context.Context) (response response.NewsfeedUnignoreItemResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NewsfeedUnsubscribeRequest defines the request for newsfeed.unsubscribe
//
// Unsubscribes the current user from comments on a specified item.
// Doc: https://dev.vk.com/method/newsfeed.unsubscribe
type NewsfeedUnsubscribeRequest struct {
	BaseRequest
}

// NewNewsfeedUnsubscribeRequest creates a new request for newsfeed.unsubscribe
func NewNewsfeedUnsubscribeRequest(a *api.API, actor actor.Actor) *NewsfeedUnsubscribeRequest {
	return &NewsfeedUnsubscribeRequest{*NewMethodBaseRequest(a, actor, "newsfeed.unsubscribe")}
}

// Exec executes the request and unmarshals the response into NewsfeedUnsubscribeResponse
func (r *NewsfeedUnsubscribeRequest) Exec(ctx context.Context) (response response.NewsfeedUnsubscribeResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
