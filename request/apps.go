package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

//TODO доделать методы ввода параметров method.Apps

// Doc: https://dev.vk.com/ru/method/apps

// AppsAddSnippetRequest defines the request for apps.addSnippet
// Adds a new snippet to the mini-app or game collection.
// Doc: https://dev.vk.com/method/apps.addSnippet
type AppsAddSnippetRequest struct {
	BaseRequest
}

// NewAppsAddSnippetRequest creates a new request for apps.addSnippet
func NewAppsAddSnippetRequest(a *api.API, actor actor.Actor) *AppsAddSnippetRequest {
	return &AppsAddSnippetRequest{*NewMethodBaseRequest(a, actor, "apps.addSnippet")}
}

// Exec executes the request and unmarshals the response into AppsAddSnippetResponse
func (r *AppsAddSnippetRequest) Exec(ctx context.Context) (response response.AppsAddSnippetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsAddUsersToTestingGroupRequest defines the request for apps.addUsersToTestingGroup
// Adds users to a testing group of the mini-app.
// Doc: https://dev.vk.com/method/apps.addUsersToTestingGroup
type AppsAddUsersToTestingGroupRequest struct {
	BaseRequest
}

// NewAppsAddUsersToTestingGroupRequest creates a new request for apps.addUsersToTestingGroup
func NewAppsAddUsersToTestingGroupRequest(a *api.API, actor actor.Actor) *AppsAddUsersToTestingGroupRequest {
	return &AppsAddUsersToTestingGroupRequest{*NewMethodBaseRequest(a, actor, "apps.addUsersToTestingGroup")}
}

// Exec executes the request and unmarshals the response into AppsAddUsersToTestingGroupResponse
func (r *AppsAddUsersToTestingGroupRequest) Exec(ctx context.Context) (response response.AppsAddUsersToTestingGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsDeleteAppRequestsRequest defines the request for apps.deleteAppRequests
// Deletes all notifications of requests sent from the current app.
// Doc: https://dev.vk.com/method/apps.deleteAppRequests
type AppsDeleteAppRequestsRequest struct {
	BaseRequest
}

// NewAppsDeleteAppRequestsRequest creates a new request for apps.deleteAppRequests
func NewAppsDeleteAppRequestsRequest(a *api.API, actor actor.Actor) *AppsDeleteAppRequestsRequest {
	return &AppsDeleteAppRequestsRequest{*NewMethodBaseRequest(a, actor, "apps.deleteAppRequests")}
}

// Exec executes the request and unmarshals the response into AppsDeleteAppRequestsResponse
func (r *AppsDeleteAppRequestsRequest) Exec(ctx context.Context) (response response.AppsDeleteAppRequestsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsDeleteSnippetRequest defines the request for apps.deleteSnippet
// Deletes a snippet from the mini-app or game.
// Doc: https://dev.vk.com/method/apps.deleteSnippet
type AppsDeleteSnippetRequest struct {
	BaseRequest
}

// NewAppsDeleteSnippetRequest creates a new request for apps.deleteSnippet
func NewAppsDeleteSnippetRequest(a *api.API, actor actor.Actor) *AppsDeleteSnippetRequest {
	return &AppsDeleteSnippetRequest{*NewMethodBaseRequest(a, actor, "apps.deleteSnippet")}
}

// Exec executes the request and unmarshals the response into AppsDeleteSnippetResponse
func (r *AppsDeleteSnippetRequest) Exec(ctx context.Context) (response response.AppsDeleteSnippetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetRequest defines the request for apps.get
// Retrieves data about apps.
// Doc: https://dev.vk.com/method/apps.get
type AppsGetRequest struct {
	BaseRequest
}

// NewAppsGetRequest creates a new request for apps.get
func NewAppsGetRequest(a *api.API, actor actor.Actor) *AppsGetRequest {
	return &AppsGetRequest{*NewMethodBaseRequest(a, actor, "apps.get")}
}

// Exec executes the request and unmarshals the response into AppsGetResponse
func (r *AppsGetRequest) Exec(ctx context.Context) (response response.AppsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetCatalogRequest defines the request for apps.getCatalog
// Returns a list of apps available to users through the app catalog.
// Doc: https://dev.vk.com/method/apps.getCatalog
type AppsGetCatalogRequest struct {
	BaseRequest
}

// NewAppsGetCatalogRequest creates a new request for apps.getCatalog
func NewAppsGetCatalogRequest(a *api.API, actor actor.Actor) *AppsGetCatalogRequest {
	return &AppsGetCatalogRequest{*NewMethodBaseRequest(a, actor, "apps.getCatalog")}
}

// Exec executes the request and unmarshals the response into AppsGetCatalogResponse
func (r *AppsGetCatalogRequest) Exec(ctx context.Context) (response response.AppsGetCatalogResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetFriendsListRequest defines the request for apps.getFriendsList
// Creates a friends list for sending invites and game requests.
// Doc: https://dev.vk.com/method/apps.getFriendsList
type AppsGetFriendsListRequest struct {
	BaseRequest
}

// NewAppsGetFriendsListRequest creates a new request for apps.getFriendsList
func NewAppsGetFriendsListRequest(a *api.API, actor actor.Actor) *AppsGetFriendsListRequest {
	return &AppsGetFriendsListRequest{*NewMethodBaseRequest(a, actor, "apps.getFriendsList")}
}

// Exec executes the request and unmarshals the response into AppsGetFriendsListResponse
func (r *AppsGetFriendsListRequest) Exec(ctx context.Context) (response response.AppsGetFriendsListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetFriendsListExtendedRequest defines the request for apps.getFriendsList
// Creates a friends list for sending invites and game requests with extended data.
// Doc: https://dev.vk.com/method/apps.getFriendsList
type AppsGetFriendsListExtendedRequest struct {
	AppsGetFriendsListRequest
}

// NewAppsGetFriendsListExtendedRequest creates a new request for apps.getFriendsList
func NewAppsGetFriendsListExtendedRequest(a *api.API, actor actor.Actor) *AppsGetFriendsListExtendedRequest {
	r := &AppsGetFriendsListExtendedRequest{*NewAppsGetFriendsListRequest(a, actor)}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into AppsGetFriendsListExtendedResponse
func (r *AppsGetFriendsListExtendedRequest) Exec(ctx context.Context) (response response.AppsGetFriendsListExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AppsGetFriendsListExtendedRequest) Fields(fields string) *AppsGetFriendsListExtendedRequest {
	r.parameters.Set(constants.ParameterNameFields, fields)
	return r
}

// AppsGetLeaderboardRequest defines the request for apps.getLeaderboard
// Retrieves user rankings in the game.
// Doc: https://dev.vk.com/method/apps.getLeaderboard
type AppsGetLeaderboardRequest struct {
	BaseRequest
}

// NewAppsGetLeaderboardRequest creates a new request for apps.getLeaderboard
func NewAppsGetLeaderboardRequest(a *api.API, actor actor.Actor) *AppsGetLeaderboardRequest {
	return &AppsGetLeaderboardRequest{*NewMethodBaseRequest(a, actor, "apps.getLeaderboard")}
}

// Exec executes the request and unmarshals the response into AppsGetLeaderboardResponse
func (r *AppsGetLeaderboardRequest) Exec(ctx context.Context) (response response.AppsGetLeaderboardResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetLeaderboardExtendedRequest defines the request for apps.getLeaderboard
// Retrieves user rankings in the game with extended data.
// Doc: https://dev.vk.com/method/apps.getLeaderboard
type AppsGetLeaderboardExtendedRequest struct {
	AppsGetLeaderboardRequest
}

// NewAppsGetLeaderboardExtendedRequest creates a new request for apps.getLeaderboard
func NewAppsGetLeaderboardExtendedRequest(a *api.API, actor actor.Actor) *AppsGetLeaderboardExtendedRequest {
	r := &AppsGetLeaderboardExtendedRequest{*NewAppsGetLeaderboardRequest(a, actor)}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into AppsGetLeaderboardResponse
func (r *AppsGetLeaderboardExtendedRequest) Exec(ctx context.Context) (response response.AppsGetLeaderboardExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetMiniAppPoliciesRequest defines the request for apps.getMiniAppPolicies
// Retrieves links specified in the terms of service and privacy policy of the mini-app.
// Doc: https://dev.vk.com/method/apps.getMiniAppPolicies
type AppsGetMiniAppPoliciesRequest struct {
	BaseRequest
}

// NewAppsGetMiniAppPoliciesRequest creates a new request for apps.getMiniAppPolicies
func NewAppsGetMiniAppPoliciesRequest(a *api.API, actor actor.Actor) *AppsGetMiniAppPoliciesRequest {
	return &AppsGetMiniAppPoliciesRequest{*NewMethodBaseRequest(a, actor, "apps.getMiniAppPolicies")}
}

// Exec executes the request and unmarshals the response into AppsGetMiniAppPoliciesResponse
func (r *AppsGetMiniAppPoliciesRequest) Exec(ctx context.Context) (response response.AppsGetMiniAppPoliciesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetScopesRequest defines the request for apps.getScopes
// Retrieves the access rights (scopes) for the app.
// Doc: https://dev.vk.com/method/apps.getScopes
type AppsGetScopesRequest struct {
	BaseRequest
}

// NewAppsGetScopesRequest creates a new request for apps.getScopes
func NewAppsGetScopesRequest(a *api.API, actor actor.Actor) *AppsGetScopesRequest {
	return &AppsGetScopesRequest{*NewMethodBaseRequest(a, actor, "apps.getScopes")}
}

// Exec executes the request and unmarshals the response into AppsGetScopesResponse
func (r *AppsGetScopesRequest) Exec(ctx context.Context) (response response.AppsGetScopesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetScoreRequest defines the request for apps.getScore
// Retrieves the user's score in the game.
// Doc: https://dev.vk.com/method/apps.getScore
type AppsGetScoreRequest struct {
	BaseRequest
}

// NewAppsGetScoreRequest creates a new request for apps.getScore
func NewAppsGetScoreRequest(a *api.API, actor actor.Actor) *AppsGetScoreRequest {
	return &AppsGetScoreRequest{*NewMethodBaseRequest(a, actor, "apps.getScore")}
}

// Exec executes the request and unmarshals the response into AppsGetScoreResponse
func (r *AppsGetScoreRequest) Exec(ctx context.Context) (response response.AppsGetScoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetSnippetsRequest defines the request for apps.getSnippets
// Retrieves information about the snippets created using apps.addSnippet.
// Doc: https://dev.vk.com/method/apps.getSnippets
type AppsGetSnippetsRequest struct {
	BaseRequest
}

// NewAppsGetSnippetsRequest creates a new request for apps.getSnippets
func NewAppsGetSnippetsRequest(a *api.API, actor actor.Actor) *AppsGetSnippetsRequest {
	return &AppsGetSnippetsRequest{*NewMethodBaseRequest(a, actor, "apps.getSnippets")}
}

// Exec executes the request and unmarshals the response into AppsGetSnippetsResponse
func (r *AppsGetSnippetsRequest) Exec(ctx context.Context) (response response.AppsGetSnippetsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsGetTestingGroupsRequest defines the request for apps.getTestingGroups
// Retrieves the testing groups of the mini-app.
// Doc: https://dev.vk.com/method/apps.getTestingGroups
type AppsGetTestingGroupsRequest struct {
	BaseRequest
}

// NewAppsGetTestingGroupsRequest creates a new request for apps.getTestingGroups
func NewAppsGetTestingGroupsRequest(a *api.API, actor actor.Actor) *AppsGetTestingGroupsRequest {
	return &AppsGetTestingGroupsRequest{*NewMethodBaseRequest(a, actor, "apps.getTestingGroups")}
}

// Exec executes the request and unmarshals the response into AppsGetTestingGroupsResponse
func (r *AppsGetTestingGroupsRequest) Exec(ctx context.Context) (response response.AppsGetTestingGroupsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsIsNotificationsAllowedRequest defines the request for apps.isNotificationsAllowed
// Checks whether the user has allowed notifications in the mini-app.
// Doc: https://dev.vk.com/method/apps.isNotificationsAllowed
type AppsIsNotificationsAllowedRequest struct {
	BaseRequest
}

// NewAppsIsNotificationsAllowedRequest creates a new request for apps.isNotificationsAllowed
func NewAppsIsNotificationsAllowedRequest(a *api.API, actor actor.Actor) *AppsIsNotificationsAllowedRequest {
	return &AppsIsNotificationsAllowedRequest{*NewMethodBaseRequest(a, actor, "apps.isNotificationsAllowed")}
}

// Exec executes the request and unmarshals the response into AppsIsNotificationsAllowedResponse
func (r *AppsIsNotificationsAllowedRequest) Exec(ctx context.Context) (response response.AppsIsNotificationsAllowedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsPromoHasActiveGiftRequest defines the request for apps.promoHasActiveGift
// Checks whether the user has an active gift in the game.
// Doc: https://dev.vk.com/method/apps.promoHasActiveGift
type AppsPromoHasActiveGiftRequest struct {
	BaseRequest
}

// NewAppsPromoHasActiveGiftRequest creates a new request for apps.promoHasActiveGift
func NewAppsPromoHasActiveGiftRequest(a *api.API, actor actor.Actor) *AppsPromoHasActiveGiftRequest {
	return &AppsPromoHasActiveGiftRequest{*NewMethodBaseRequest(a, actor, "apps.promoHasActiveGift")}
}

// Exec executes the request and unmarshals the response into AppsPromoHasActiveGiftResponse
func (r *AppsPromoHasActiveGiftRequest) Exec(ctx context.Context) (response response.AppsPromoHasActiveGiftResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsPromoUseGiftRequest defines the request for apps.promoUseGift
// Marks the user's gift as used in a promo campaign.
// Doc: https://dev.vk.com/method/apps.promoUseGift
type AppsPromoUseGiftRequest struct {
	BaseRequest
}

// NewAppsPromoUseGiftRequest creates a new request for apps.promoUseGift
func NewAppsPromoUseGiftRequest(a *api.API, actor actor.Actor) *AppsPromoUseGiftRequest {
	return &AppsPromoUseGiftRequest{*NewMethodBaseRequest(a, actor, "apps.promoUseGift")}
}

// Exec executes the request and unmarshals the response into AppsPromoUseGiftResponse
func (r *AppsPromoUseGiftRequest) Exec(ctx context.Context) (response response.AppsPromoUseGiftResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsRemoveTestingGroupRequest defines the request for apps.removeTestingGroup
// Removes the specified testing group from the mini-app.
// Doc: https://dev.vk.com/method/apps.removeTestingGroup
type AppsRemoveTestingGroupRequest struct {
	BaseRequest
}

// NewAppsRemoveTestingGroupRequest creates a new request for apps.removeTestingGroup
func NewAppsRemoveTestingGroupRequest(a *api.API, actor actor.Actor) *AppsRemoveTestingGroupRequest {
	return &AppsRemoveTestingGroupRequest{*NewMethodBaseRequest(a, actor, "apps.removeTestingGroup")}
}

// Exec executes the request and unmarshals the response into AppsRemoveTestingGroupResponse
func (r *AppsRemoveTestingGroupRequest) Exec(ctx context.Context) (response response.AppsRemoveTestingGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsRemoveUsersFromTestingGroupsRequest defines the request for apps.removeUsersFromTestingGroups
// Removes specified users from the testing groups of the mini-app.
// Doc: https://dev.vk.com/method/apps.removeUsersFromTestingGroups
type AppsRemoveUsersFromTestingGroupsRequest struct {
	BaseRequest
}

// NewAppsRemoveUsersFromTestingGroupsRequest creates a new request for apps.removeUsersFromTestingGroups
func NewAppsRemoveUsersFromTestingGroupsRequest(a *api.API, actor actor.Actor) *AppsRemoveUsersFromTestingGroupsRequest {
	return &AppsRemoveUsersFromTestingGroupsRequest{*NewMethodBaseRequest(a, actor, "apps.removeUsersFromTestingGroups")}
}

// Exec executes the request and unmarshals the response into AppsRemoveUsersFromTestingGroupsResponse
func (r *AppsRemoveUsersFromTestingGroupsRequest) Exec(ctx context.Context) (response response.AppsRemoveUsersFromTestingGroupsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsSendRequestRequest defines the request for apps.sendRequest
// Sends a request to another user using VKontakte authorization.
// Doc: https://dev.vk.com/method/apps.sendRequest
type AppsSendRequestRequest struct {
	BaseRequest
}

// NewAppsSendRequestRequest creates a new request for apps.sendRequest
func NewAppsSendRequestRequest(a *api.API, actor actor.Actor) *AppsSendRequestRequest {
	return &AppsSendRequestRequest{*NewMethodBaseRequest(a, actor, "apps.sendRequest")}
}

// Exec executes the request and unmarshals the response into AppsSendRequestResponse
func (r *AppsSendRequestRequest) Exec(ctx context.Context) (response response.AppsSendRequestResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AppsUpdateMetaForTestingGroupRequest defines the request for apps.updateMetaForTestingGroup
// Creates or updates a testing group in the mini-app.
// Doc: https://dev.vk.com/method/apps.updateMetaForTestingGroup
type AppsUpdateMetaForTestingGroupRequest struct {
	BaseRequest
}

// NewAppsUpdateMetaForTestingGroupRequest creates a new request for apps.updateMetaForTestingGroup
func NewAppsUpdateMetaForTestingGroupRequest(a *api.API, actor actor.Actor) *AppsUpdateMetaForTestingGroupRequest {
	return &AppsUpdateMetaForTestingGroupRequest{*NewMethodBaseRequest(a, actor, "apps.updateMetaForTestingGroup")}
}

// Exec executes the request and unmarshals the response into AppsUpdateMetaForTestingGroupResponse
func (r *AppsUpdateMetaForTestingGroupRequest) Exec(ctx context.Context) (response response.AppsUpdateMetaForTestingGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
