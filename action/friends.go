package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Friends Doc: https://dev.vk.com/ru/method/friends
type Friends struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/ru/method/friends.add
func (f *Friends) Add(actor actor.Actor) *request.FriendsAddRequest {
	return request.NewFriendsAddRequest(f.api, actor)
}

// AddList Doc: https://dev.vk.com/ru/method/friends.addList
func (f *Friends) AddList(actor actor.Actor) *request.FriendsAddListRequest {
	return request.NewFriendsAddListRequest(f.api, actor)
}

// AreFriends Doc: https://dev.vk.com/ru/method/friends.areFriends
func (f *Friends) AreFriends(actor actor.Actor) *request.FriendsAreFriendsRequest {
	return request.NewFriendsAreFriendsRequest(f.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/friends.delete
func (f *Friends) Delete(actor actor.Actor) *request.FriendsDeleteRequest {
	return request.NewFriendsDeleteRequest(f.api, actor)
}

// DeleteAllRequests Doc: https://dev.vk.com/ru/method/friends.deleteAllRequests
func (f *Friends) DeleteAllRequests(actor actor.Actor) *request.FriendsDeleteAllRequestsRequest {
	return request.NewFriendsDeleteAllRequestsRequest(f.api, actor)
}

// DeleteList Doc: https://dev.vk.com/ru/method/friends.deleteList
func (f *Friends) DeleteList(actor actor.Actor) *request.FriendsDeleteListRequest {
	return request.NewFriendsDeleteListRequest(f.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/friends.edit
func (f *Friends) Edit(actor actor.Actor) *request.FriendsEditRequest {
	return request.NewFriendsEditRequest(f.api, actor)
}

// EditList Doc: https://dev.vk.com/ru/method/friends.editList
func (f *Friends) EditList(actor actor.Actor) *request.FriendsEditListRequest {
	return request.NewFriendsEditListRequest(f.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/friends.get
func (f *Friends) Get(actor actor.Actor) *request.FriendsGetRequest {
	return request.NewFriendsGetRequest(f.api, actor)
}

// GetAppUsers Doc: https://dev.vk.com/ru/method/friends.getAppUsers
func (f *Friends) GetAppUsers(actor actor.Actor) *request.FriendsGetAppUsersRequest {
	return request.NewFriendsGetAppUsersRequest(f.api, actor)
}

// GetLists Doc: https://dev.vk.com/ru/method/friends.getLists
func (f *Friends) GetLists(actor actor.Actor) *request.FriendsGetListsRequest {
	return request.NewFriendsGetListsRequest(f.api, actor)
}

// GetMutual Doc: https://dev.vk.com/ru/method/friends.getMutual
func (f *Friends) GetMutual(actor actor.Actor) *request.FriendsGetMutualRequest {
	return request.NewFriendsGetMutualRequest(f.api, actor)
}

// GetOnline Doc: https://dev.vk.com/ru/method/friends.getOnline
func (f *Friends) GetOnline(actor actor.Actor) *request.FriendsGetOnlineRequest {
	return request.NewFriendsGetOnlineRequest(f.api, actor)
}

// GetRecent Doc: https://dev.vk.com/ru/method/friends.getRecent
func (f *Friends) GetRecent(actor actor.Actor) *request.FriendsGetRecentRequest {
	return request.NewFriendsGetRecentRequest(f.api, actor)
}

// GetRequests Doc: https://dev.vk.com/ru/method/friends.getRequests
func (f *Friends) GetRequests(actor actor.Actor) *request.FriendsGetRequestsRequest {
	return request.NewFriendsGetRequestsRequest(f.api, actor)
}

// GetRequestsExtended Doc: https://dev.vk.com/ru/method/friends.getRequests
func (f *Friends) GetRequestsExtended(actor actor.Actor) *request.FriendsGetRequestsExtendedRequest {
	return request.NewFriendsGetRequestsExtendedRequest(f.api, actor)
}

// GetSuggestions Doc: https://dev.vk.com/ru/method/friends.getSuggestions
func (f *Friends) GetSuggestions(actor actor.Actor) *request.FriendsGetSuggestionsRequest {
	return request.NewFriendsGetSuggestionsRequest(f.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/friends.search
func (f *Friends) Search(actor actor.Actor) *request.FriendsSearchRequest {
	return request.NewFriendsSearchRequest(f.api, actor)
}
