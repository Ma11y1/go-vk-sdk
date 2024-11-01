package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Newsfeed Doc: https://dev.vk.com/ru/method/newsfeed
type Newsfeed struct {
	BaseAction
}

// AddBan Doc: https://dev.vk.com/ru/method/newsfeed.addBan
func (n *Newsfeed) AddBan(actor actor.Actor) *request.NewsfeedAddBanRequest {
	return request.NewNewsfeedAddBanRequest(n.api, actor)
}

// DeleteBan Doc: https://dev.vk.com/ru/method/newsfeed.deleteBan
func (n *Newsfeed) DeleteBan(actor actor.Actor) *request.NewsfeedDeleteBanRequest {
	return request.NewNewsfeedDeleteBanRequest(n.api, actor)
}

// DeleteList Doc: https://dev.vk.com/ru/method/newsfeed.deleteList
func (n *Newsfeed) DeleteList(actor actor.Actor) *request.NewsfeedDeleteListRequest {
	return request.NewNewsfeedDeleteListRequest(n.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/newsfeed.get
func (n *Newsfeed) Get(actor actor.Actor) *request.NewsfeedGetRequest {
	return request.NewNewsfeedGetRequest(n.api, actor)
}

// GetBanned Doc: https://dev.vk.com/ru/method/newsfeed.getBanned
func (n *Newsfeed) GetBanned(actor actor.Actor) *request.NewsfeedGetBannedRequest {
	return request.NewNewsfeedGetBannedRequest(n.api, actor)
}

// GetBannedExtended Doc: https://dev.vk.com/ru/method/newsfeed.getBanned
func (n *Newsfeed) GetBannedExtended(actor actor.Actor) *request.NewsfeedGetBannedExtendedRequest {
	return request.NewNewsfeedGetBannedExtendedRequest(n.api, actor)
}

// GetComments Doc: https://dev.vk.com/ru/method/newsfeed.getComments
func (n *Newsfeed) GetComments(actor actor.Actor) *request.NewsfeedGetCommentsRequest {
	return request.NewNewsfeedGetCommentsRequest(n.api, actor)
}

// GetLists Doc: https://dev.vk.com/ru/method/newsfeed.getLists
func (n *Newsfeed) GetLists(actor actor.Actor) *request.NewsfeedGetListsRequest {
	return request.NewNewsfeedGetListsRequest(n.api, actor)
}

// GetMentions Doc: https://dev.vk.com/ru/method/newsfeed.getMentions
func (n *Newsfeed) GetMentions(actor actor.Actor) *request.NewsfeedGetMentionsRequest {
	return request.NewNewsfeedGetMentionsRequest(n.api, actor)
}

// GetRecommended Doc: https://dev.vk.com/ru/method/newsfeed.getRecommended
func (n *Newsfeed) GetRecommended(actor actor.Actor) *request.NewsfeedGetRecommendedRequest {
	return request.NewNewsfeedGetRecommendedRequest(n.api, actor)
}

// GetSuggestedSources Doc: https://dev.vk.com/ru/method/newsfeed.getSuggestedSources
func (n *Newsfeed) GetSuggestedSources(actor actor.Actor) *request.NewsfeedGetSuggestedSourcesRequest {
	return request.NewNewsfeedGetSuggestedSourcesRequest(n.api, actor)
}

// IgnoreItem Doc: https://dev.vk.com/ru/method/newsfeed.ignoreItem
func (n *Newsfeed) IgnoreItem(actor actor.Actor) *request.NewsfeedIgnoreItemRequest {
	return request.NewNewsfeedIgnoreItemRequest(n.api, actor)
}

// SaveList Doc: https://dev.vk.com/ru/method/newsfeed.saveList
func (n *Newsfeed) SaveList(actor actor.Actor) *request.NewsfeedSaveListRequest {
	return request.NewNewsfeedSaveListRequest(n.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/newsfeed.search
func (n *Newsfeed) Search(actor actor.Actor) *request.NewsfeedSearchRequest {
	return request.NewNewsfeedSearchRequest(n.api, actor)
}

// SearchExtended Doc: https://dev.vk.com/ru/method/newsfeed.search
func (n *Newsfeed) SearchExtended(actor actor.Actor) *request.NewsfeedSearchExtendedRequest {
	return request.NewNewsfeedSearchExtendedRequest(n.api, actor)
}

// UnignoreItem Doc: https://dev.vk.com/ru/method/newsfeed.unignoreItem
func (n *Newsfeed) UnignoreItem(actor actor.Actor) *request.NewsfeedUnignoreItemRequest {
	return request.NewNewsfeedUnignoreItemRequest(n.api, actor)
}

// Unsubscribe Doc: https://dev.vk.com/ru/method/newsfeed.unsubscribe
func (n *Newsfeed) Unsubscribe(actor actor.Actor) *request.NewsfeedUnsubscribeRequest {
	return request.NewNewsfeedUnsubscribeRequest(n.api, actor)
}
