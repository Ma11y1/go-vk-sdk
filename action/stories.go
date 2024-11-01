package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Stories Doc: https://dev.vk.com/ru/method/stories
type Stories struct {
	BaseAction
}

// BanOwner Doc: https://dev.vk.com/method/stories.banOwner
func (s *Stories) BanOwner(actor actor.Actor) *request.StoriesBanOwnerRequest {
	return request.NewStoriesBanOwnerRequest(s.api, actor)
}

// Delete Doc: https://dev.vk.com/method/stories.delete
func (s *Stories) Delete(actor actor.Actor) *request.StoriesDeleteRequest {
	return request.NewStoriesDeleteRequest(s.api, actor)
}

// Get Doc: https://dev.vk.com/method/stories.get
func (s *Stories) Get(actor actor.Actor) *request.StoriesGetRequest {
	return request.NewStoriesGetRequest(s.api, actor)
}

// GetExtended Doc: https://dev.vk.com/method/stories.get
func (s *Stories) GetExtended(actor actor.Actor) *request.StoriesGetExtendedRequest {
	return request.NewStoriesGetExtendedRequest(s.api, actor)
}

// GetBanned Doc: https://dev.vk.com/method/stories.getBanned
func (s *Stories) GetBanned(actor actor.Actor) *request.StoriesGetBannedRequest {
	return request.NewStoriesGetBannedRequest(s.api, actor)
}

// GetBannedExtended Doc: https://dev.vk.com/method/stories.getBanned
func (s *Stories) GetBannedExtended(actor actor.Actor) *request.StoriesGetBannedExtendedRequest {
	return request.NewStoriesGetBannedExtendedRequest(s.api, actor)
}

// GetByID Doc: https://dev.vk.com/method/stories.getById
func (s *Stories) GetByID(actor actor.Actor) *request.StoriesGetByIDRequest {
	return request.NewStoriesGetByIDRequest(s.api, actor)
}

// GetByIDExtended Doc: https://dev.vk.com/method/stories.getById
func (s *Stories) GetByIDExtended(actor actor.Actor) *request.StoriesGetByIDExtendedRequest {
	return request.NewStoriesGetByIDExtendedRequest(s.api, actor)
}

// GetDetailedStats Doc: https://dev.vk.com/method/stories.getDetailedStats
func (s *Stories) GetDetailedStats(actor actor.Actor) *request.StoriesGetDetailedStatsRequest {
	return request.NewStoriesGetDetailedStatsRequest(s.api, actor)
}

// GetPhotoUploadServer Doc: https://dev.vk.com/method/stories.getPhotoUploadServer
func (s *Stories) GetPhotoUploadServer(actor actor.Actor) *request.StoriesGetPhotoUploadServerRequest {
	return request.NewStoriesGetPhotoUploadServerRequest(s.api, actor)
}

// GetReplies Doc: https://dev.vk.com/method/stories.getReplies
func (s *Stories) GetReplies(actor actor.Actor) *request.StoriesGetRepliesRequest {
	return request.NewStoriesGetRepliesRequest(s.api, actor)
}

// GetRepliesExtended Doc: https://dev.vk.com/method/stories.getReplies
func (s *Stories) GetRepliesExtended(actor actor.Actor) *request.StoriesGetRepliesExtendedRequest {
	return request.NewStoriesGetRepliesExtendedRequest(s.api, actor)
}

// GetStats Doc: https://dev.vk.com/method/stories.getStats
func (s *Stories) GetStats(actor actor.Actor) *request.StoriesGetStatsRequest {
	return request.NewStoriesGetStatsRequest(s.api, actor)
}

// GetVideoUploadServer Doc: https://dev.vk.com/method/stories.getVideoUploadServer
func (s *Stories) GetVideoUploadServer(actor actor.Actor) *request.StoriesGetVideoUploadServerRequest {
	return request.NewStoriesGetVideoUploadServerRequest(s.api, actor)
}

// GetViewers Doc: https://dev.vk.com/method/stories.getViewers
func (s *Stories) GetViewers(actor actor.Actor) *request.StoriesGetViewersRequest {
	return request.NewStoriesGetViewersRequest(s.api, actor)
}

// HideAllReplies Doc: https://dev.vk.com/method/stories.hideAllReplies
func (s *Stories) HideAllReplies(actor actor.Actor) *request.StoriesHideAllRepliesRequest {
	return request.NewStoriesHideAllRepliesRequest(s.api, actor)
}

// HideReply Doc: https://dev.vk.com/method/stories.hideReply
func (s *Stories) HideReply(actor actor.Actor) *request.StoriesHideReplyRequest {
	return request.NewStoriesHideReplyRequest(s.api, actor)
}

// Save Doc: https://dev.vk.com/method/stories.save
func (s *Stories) Save(actor actor.Actor) *request.StoriesSaveRequest {
	return request.NewStoriesSaveRequest(s.api, actor)
}

// Search Doc: https://dev.vk.com/method/stories.search
func (s *Stories) Search(actor actor.Actor) *request.StoriesSearchRequest {
	return request.NewStoriesSearchRequest(s.api, actor)
}

// SearchExtended Doc: https://dev.vk.com/method/stories.search
func (s *Stories) SearchExtended(actor actor.Actor) *request.StoriesSearchExtendedRequest {
	return request.NewStoriesSearchExtendedRequest(s.api, actor)
}

// SendInteraction Doc: https://dev.vk.com/method/stories.sendInteraction
func (s *Stories) SendInteraction(actor actor.Actor) *request.StoriesSendInteractionRequest {
	return request.NewStoriesSendInteractionRequest(s.api, actor)
}

// UnbanOwner Doc: https://dev.vk.com/method/stories.unbanOwner
func (s *Stories) UnbanOwner(actor actor.Actor) *request.StoriesUnbanOwnerRequest {
	return request.NewStoriesUnbanOwnerRequest(s.api, actor)
}
