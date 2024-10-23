package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/friends

// FriendsAddRequest defines the request for friends.add
//
// Approves or creates a friend request.
// Doc: https://dev.vk.com/method/friends.add
type FriendsAddRequest struct {
	BaseRequest
}

// NewFriendsAddRequest creates a new request for friends.add
func NewFriendsAddRequest(a *api.API, actor actor.Actor) *FriendsAddRequest {
	return &FriendsAddRequest{*NewMethodBaseRequest(a, actor, "friends.add")}
}

// Exec executes the request and unmarshals the response into FriendsAddResponse
func (r *FriendsAddRequest) Exec(ctx context.Context) (response response.FriendsAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsAddListRequest defines the request for friends.addList
//
// Creates a new friend list for the current user.
// Doc: https://dev.vk.com/method/friends.addList
type FriendsAddListRequest struct {
	BaseRequest
}

// NewFriendsAddListRequest creates a new request for friends.addList
func NewFriendsAddListRequest(a *api.API, actor actor.Actor) *FriendsAddListRequest {
	return &FriendsAddListRequest{*NewMethodBaseRequest(a, actor, "friends.addList")}
}

// Exec executes the request and unmarshals the response into FriendsAddListResponse
func (r *FriendsAddListRequest) Exec(ctx context.Context) (response response.FriendsAddListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsAreFriendsRequest defines the request for friends.areFriends
//
// Returns information about whether the current user is friends with the specified users.
// Doc: https://dev.vk.com/method/friends.areFriends
type FriendsAreFriendsRequest struct {
	BaseRequest
}

// NewFriendsAreFriendsRequest creates a new request for friends.areFriends
func NewFriendsAreFriendsRequest(a *api.API, actor actor.Actor) *FriendsAreFriendsRequest {
	return &FriendsAreFriendsRequest{*NewMethodBaseRequest(a, actor, "friends.areFriends")}
}

// Exec executes the request and unmarshals the response into FriendsAreFriendsResponse
func (r *FriendsAreFriendsRequest) Exec(ctx context.Context) (response response.FriendsAreFriendsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsDeleteRequest defines the request for friends.delete
//
// Deletes a user from the friend list or declines a friend request.
// Doc: https://dev.vk.com/method/friends.delete
type FriendsDeleteRequest struct {
	BaseRequest
}

// NewFriendsDeleteRequest creates a new request for friends.delete
func NewFriendsDeleteRequest(a *api.API, actor actor.Actor) *FriendsDeleteRequest {
	return &FriendsDeleteRequest{*NewMethodBaseRequest(a, actor, "friends.delete")}
}

// Exec executes the request and unmarshals the response into FriendsDeleteResponse
func (r *FriendsDeleteRequest) Exec(ctx context.Context) (response response.FriendsDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsDeleteAllRequestsRequest defines the request for friends.deleteAllRequests
//
// Marks all incoming friend requests as viewed.
// Doc: https://dev.vk.com/method/friends.deleteAllRequests
type FriendsDeleteAllRequestsRequest struct {
	BaseRequest
}

// NewFriendsDeleteAllRequestsRequest creates a new request for friends.deleteAllRequests
func NewFriendsDeleteAllRequestsRequest(a *api.API, actor actor.Actor) *FriendsDeleteAllRequestsRequest {
	return &FriendsDeleteAllRequestsRequest{*NewMethodBaseRequest(a, actor, "friends.deleteAllRequests")}
}

// Exec executes the request and unmarshals the response into FriendsDeleteAllRequestsResponse
func (r *FriendsDeleteAllRequestsRequest) Exec(ctx context.Context) (response response.FriendsDeleteAllRequestsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsDeleteListRequest defines the request for friends.deleteList
//
// Deletes an existing friend list for the current user.
// Doc: https://dev.vk.com/method/friends.deleteList
type FriendsDeleteListRequest struct {
	BaseRequest
}

// NewFriendsDeleteListRequest creates a new request for friends.deleteList
func NewFriendsDeleteListRequest(a *api.API, actor actor.Actor) *FriendsDeleteListRequest {
	return &FriendsDeleteListRequest{*NewMethodBaseRequest(a, actor, "friends.deleteList")}
}

// Exec executes the request and unmarshals the response into FriendsDeleteListResponse
func (r *FriendsDeleteListRequest) Exec(ctx context.Context) (response response.FriendsDeleteListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsEditRequest defines the request for friends.edit
//
// Edits friend lists for the selected friend.
// Doc: https://dev.vk.com/method/friends.edit
type FriendsEditRequest struct {
	BaseRequest
}

// NewFriendsEditRequest creates a new request for friends.edit
func NewFriendsEditRequest(a *api.API, actor actor.Actor) *FriendsEditRequest {
	return &FriendsEditRequest{*NewMethodBaseRequest(a, actor, "friends.edit")}
}

// Exec executes the request and unmarshals the response into FriendsEditResponse
func (r *FriendsEditRequest) Exec(ctx context.Context) (response response.FriendsEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsEditListRequest defines the request for friends.editList
//
// Edits an existing friend list for the current user.
// Doc: https://dev.vk.com/method/friends.editList
type FriendsEditListRequest struct {
	BaseRequest
}

// NewFriendsEditListRequest creates a new request for friends.editList
func NewFriendsEditListRequest(a *api.API, actor actor.Actor) *FriendsEditListRequest {
	return &FriendsEditListRequest{*NewMethodBaseRequest(a, actor, "friends.editList")}
}

// Exec executes the request and unmarshals the response into FriendsEditListResponse
func (r *FriendsEditListRequest) Exec(ctx context.Context) (response response.FriendsEditListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetRequest defines the request for friends.get
//
// Returns a list of user friend IDs or extended information about friends if the fields parameter is used.
// Doc: https://dev.vk.com/method/friends.get
type FriendsGetRequest struct {
	BaseRequest
}

// NewFriendsGetRequest creates a new request for friends.get
func NewFriendsGetRequest(a *api.API, actor actor.Actor) *FriendsGetRequest {
	return &FriendsGetRequest{*NewMethodBaseRequest(a, actor, "friends.get")}
}

// Exec executes the request and unmarshals the response into FriendsGetResponse
func (r *FriendsGetRequest) Exec(ctx context.Context) (response response.FriendsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetAppUsersRequest defines the request for friends.getAppUsers
//
// Returns a list of friend IDs who have installed this application.
// Doc: https://dev.vk.com/method/friends.getAppUsers
type FriendsGetAppUsersRequest struct {
	BaseRequest
}

// NewFriendsGetAppUsersRequest creates a new request for friends.getAppUsers
func NewFriendsGetAppUsersRequest(a *api.API, actor actor.Actor) *FriendsGetAppUsersRequest {
	return &FriendsGetAppUsersRequest{*NewMethodBaseRequest(a, actor, "friends.getAppUsers")}
}

// Exec executes the request and unmarshals the response into FriendsGetAppUsersResponse
func (r *FriendsGetAppUsersRequest) Exec(ctx context.Context) (response response.FriendsGetAppUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetListsRequest defines the request for friends.getLists
//
// Returns a list of friend labels.
// Doc: https://dev.vk.com/method/friends.getLists
type FriendsGetListsRequest struct {
	BaseRequest
}

// NewFriendsGetListsRequest creates a new request for friends.getLists
func NewFriendsGetListsRequest(a *api.API, actor actor.Actor) *FriendsGetListsRequest {
	return &FriendsGetListsRequest{*NewMethodBaseRequest(a, actor, "friends.getLists")}
}

// Exec executes the request and unmarshals the response into FriendsGetListsResponse
func (r *FriendsGetListsRequest) Exec(ctx context.Context) (response response.FriendsGetListsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetMutualRequest defines the request for friends.getMutual
//
// Returns a list of mutual friend IDs between a pair of users.
// Doc: https://dev.vk.com/method/friends.getMutual
type FriendsGetMutualRequest struct {
	BaseRequest
}

// NewFriendsGetMutualRequest creates a new request for friends.getMutual
func NewFriendsGetMutualRequest(a *api.API, actor actor.Actor) *FriendsGetMutualRequest {
	return &FriendsGetMutualRequest{*NewMethodBaseRequest(a, actor, "friends.getMutual")}
}

// Exec executes the request and unmarshals the response into FriendsGetMutualResponse
func (r *FriendsGetMutualRequest) Exec(ctx context.Context) (response response.FriendsGetMutualResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetOnlineRequest defines the request for friends.getOnline
//
// Returns a list of friend IDs who are currently online.
// Doc: https://dev.vk.com/method/friends.getOnline
type FriendsGetOnlineRequest struct {
	BaseRequest
}

// NewFriendsGetOnlineRequest creates a new request for friends.getOnline
func NewFriendsGetOnlineRequest(a *api.API, actor actor.Actor) *FriendsGetOnlineRequest {
	return &FriendsGetOnlineRequest{*NewMethodBaseRequest(a, actor, "friends.getOnline")}
}

// Exec executes the request and unmarshals the response into FriendsGetOnlineResponse
func (r *FriendsGetOnlineRequest) Exec(ctx context.Context) (response response.FriendsGetOnlineResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetRecentRequest defines the request for friends.getRecent
//
// Returns a list of IDs for recently added friends of the current user.
// Doc: https://dev.vk.com/method/friends.getRecent
type FriendsGetRecentRequest struct {
	BaseRequest
}

// NewFriendsGetRecentRequest creates a new request for friends.getRecent
func NewFriendsGetRecentRequest(a *api.API, actor actor.Actor) *FriendsGetRecentRequest {
	return &FriendsGetRecentRequest{*NewMethodBaseRequest(a, actor, "friends.getRecent")}
}

// Exec executes the request and unmarshals the response into FriendsGetRecentResponse
func (r *FriendsGetRecentRequest) Exec(ctx context.Context) (response response.FriendsGetRecentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetRequestsRequest defines the request for friends.getRequests
//
// Returns information about received or sent friend requests for the current user.
// Doc: https://dev.vk.com/method/friends.getRequests
type FriendsGetRequestsRequest struct {
	BaseRequest
}

// NewFriendsGetRequestsRequest creates a new request for friends.getRequests
func NewFriendsGetRequestsRequest(a *api.API, actor actor.Actor) *FriendsGetRequestsRequest {
	return &FriendsGetRequestsRequest{*NewMethodBaseRequest(a, actor, "friends.getRequests")}
}

// Exec executes the request and unmarshals the response into FriendsGetRequestsResponse
func (r *FriendsGetRequestsRequest) Exec(ctx context.Context) (response response.FriendsGetRequestsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsGetSuggestionsRequest defines the request for friends.getSuggestions
//
// Returns a list of user profiles that may be friends with the current user.
// Doc: https://dev.vk.com/method/friends.getSuggestions
type FriendsGetSuggestionsRequest struct {
	BaseRequest
}

// NewFriendsGetSuggestionsRequest creates a new request for friends.getSuggestions
func NewFriendsGetSuggestionsRequest(a *api.API, actor actor.Actor) *FriendsGetSuggestionsRequest {
	return &FriendsGetSuggestionsRequest{*NewMethodBaseRequest(a, actor, "friends.getSuggestions")}
}

// Exec executes the request and unmarshals the response into FriendsGetSuggestionsResponse
func (r *FriendsGetSuggestionsRequest) Exec(ctx context.Context) (response response.FriendsGetSuggestionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FriendsSearchRequest defines the request for friends.search
//
// Allows searching through the user's friends list.
// Doc: https://dev.vk.com/method/friends.search
type FriendsSearchRequest struct {
	BaseRequest
}

// NewFriendsSearchRequest creates a new request for friends.search
func NewFriendsSearchRequest(a *api.API, actor actor.Actor) *FriendsSearchRequest {
	return &FriendsSearchRequest{*NewMethodBaseRequest(a, actor, "friends.search")}
}

// Exec executes the request and unmarshals the response into FriendsSearchResponse
func (r *FriendsSearchRequest) Exec(ctx context.Context) (response response.FriendsSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
