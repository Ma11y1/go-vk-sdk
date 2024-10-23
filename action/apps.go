package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Apps Doc: https://dev.vk.com/ru/method/apps
type Apps struct {
	BaseAction
}

// AddSnippet Doc: https://dev.vk.com/method/apps.addSnippet
func (a *Apps) AddSnippet(actor actor.Actor) *request.AppsAddSnippetRequest {
	return request.NewAppsAddSnippetRequest(a.api, actor)
}

// AddUsersToTestingGroup Doc: https://dev.vk.com/method/apps.addUsersToTestingGroup
func (a *Apps) AddUsersToTestingGroup(actor actor.Actor) *request.AppsAddUsersToTestingGroupRequest {
	return request.NewAppsAddUsersToTestingGroupRequest(a.api, actor)
}

// DeleteAppRequests Doc: https://dev.vk.com/method/apps.deleteAppRequests
func (a *Apps) DeleteAppRequests(actor actor.Actor) *request.AppsDeleteAppRequestsRequest {
	return request.NewAppsDeleteAppRequestsRequest(a.api, actor)
}

// DeleteSnippet Doc: https://dev.vk.com/method/apps.deleteSnippet
func (a *Apps) DeleteSnippet(actor actor.Actor) *request.AppsDeleteSnippetRequest {
	return request.NewAppsDeleteSnippetRequest(a.api, actor)
}

// Get Doc: https://dev.vk.com/method/apps.get
func (a *Apps) Get(actor actor.Actor) *request.AppsGetRequest {
	return request.NewAppsGetRequest(a.api, actor)
}

// GetCatalog Doc: https://dev.vk.com/method/apps.getCatalog
func (a *Apps) GetCatalog(actor actor.Actor) *request.AppsGetCatalogRequest {
	return request.NewAppsGetCatalogRequest(a.api, actor)
}

// GetFriendsList Doc: https://dev.vk.com/method/apps.getFriendsList
func (a *Apps) GetFriendsList(actor actor.Actor) *request.AppsGetFriendsListRequest {
	return request.NewAppsGetFriendsListRequest(a.api, actor)
}

// GetLeaderboard Doc: https://dev.vk.com/method/apps.getLeaderboard
func (a *Apps) GetLeaderboard(actor actor.Actor) *request.AppsGetLeaderboardRequest {
	return request.NewAppsGetLeaderboardRequest(a.api, actor)
}

// GetMiniAppPolicies Doc: https://dev.vk.com/method/apps.getMiniAppPolicies
func (a *Apps) GetMiniAppPolicies(actor actor.Actor) *request.AppsGetMiniAppPoliciesRequest {
	return request.NewAppsGetMiniAppPoliciesRequest(a.api, actor)
}

// GetScopes Doc: https://dev.vk.com/method/apps.getScopes
func (a *Apps) GetScopes(actor actor.Actor) *request.AppsGetScopesRequest {
	return request.NewAppsGetScopesRequest(a.api, actor)
}

// GetScore Doc: https://dev.vk.com/method/apps.getScore
func (a *Apps) GetScore(actor actor.Actor) *request.AppsGetScoreRequest {
	return request.NewAppsGetScoreRequest(a.api, actor)
}

// GetSnippets Doc: https://dev.vk.com/method/apps.getSnippets
func (a *Apps) GetSnippets(actor actor.Actor) *request.AppsGetSnippetsRequest {
	return request.NewAppsGetSnippetsRequest(a.api, actor)
}

// GetTestingGroups Doc: https://dev.vk.com/method/apps.getTestingGroups
func (a *Apps) GetTestingGroups(actor actor.Actor) *request.AppsGetTestingGroupsRequest {
	return request.NewAppsGetTestingGroupsRequest(a.api, actor)
}

// IsNotificationsAllowed Doc: https://dev.vk.com/method/apps.isNotificationsAllowed
func (a *Apps) IsNotificationsAllowed(actor actor.Actor) *request.AppsIsNotificationsAllowedRequest {
	return request.NewAppsIsNotificationsAllowedRequest(a.api, actor)
}

// PromoHasActiveGift Doc: https://dev.vk.com/method/apps.promoHasActiveGift
func (a *Apps) PromoHasActiveGift(actor actor.Actor) *request.AppsPromoHasActiveGiftRequest {
	return request.NewAppsPromoHasActiveGiftRequest(a.api, actor)
}

// PromoUseGift Doc: https://dev.vk.com/method/apps.promoUseGift
func (a *Apps) PromoUseGift(actor actor.Actor) *request.AppsPromoUseGiftRequest {
	return request.NewAppsPromoUseGiftRequest(a.api, actor)
}

// RemoveTestingGroup Doc: https://dev.vk.com/method/apps.removeTestingGroup
func (a *Apps) RemoveTestingGroup(actor actor.Actor) *request.AppsRemoveTestingGroupRequest {
	return request.NewAppsRemoveTestingGroupRequest(a.api, actor)
}

// RemoveUsersFromTestingGroups Doc: https://dev.vk.com/method/apps.removeUsersFromTestingGroups
func (a *Apps) RemoveUsersFromTestingGroups(actor actor.Actor) *request.AppsRemoveUsersFromTestingGroupsRequest {
	return request.NewAppsRemoveUsersFromTestingGroupsRequest(a.api, actor)
}

// SendRequest Doc: https://dev.vk.com/method/apps.sendRequest
func (a *Apps) SendRequest(actor actor.Actor) *request.AppsSendRequestRequest {
	return request.NewAppsSendRequestRequest(a.api, actor)
}

// UpdateMetaForTestingGroup Doc: https://dev.vk.com/method/apps.updateMetaForTestingGroup
func (a *Apps) UpdateMetaForTestingGroup(actor actor.Actor) *request.AppsUpdateMetaForTestingGroupRequest {
	return request.NewAppsUpdateMetaForTestingGroupRequest(a.api, actor)
}
