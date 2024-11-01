package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Wall Doc: https://dev.vk.com/ru/method/wall
type Wall struct {
	BaseAction
}

// CheckCopyrightLink Doc: https://dev.vk.com/method/wall.checkCopyrightLink
func (w *Wall) CheckCopyrightLink(actor actor.Actor) *request.WallCheckCopyrightLinkRequest {
	return request.NewWallCheckCopyrightLinkRequest(w.api, actor)
}

// CloseComments Doc: https://dev.vk.com/method/wall.closeComments
func (w *Wall) CloseComments(actor actor.Actor) *request.WallCloseCommentsRequest {
	return request.NewWallCloseCommentsRequest(w.api, actor)
}

// CreateComment Doc: https://dev.vk.com/method/wall.createComment
func (w *Wall) CreateComment(actor actor.Actor) *request.WallCreateCommentRequest {
	return request.NewWallCreateCommentRequest(w.api, actor)
}

// Delete Doc: https://dev.vk.com/method/wall.delete
func (w *Wall) Delete(actor actor.Actor) *request.WallDeleteRequest {
	return request.NewWallDeleteRequest(w.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/method/wall.deleteComment
func (w *Wall) DeleteComment(actor actor.Actor) *request.WallDeleteCommentRequest {
	return request.NewWallDeleteCommentRequest(w.api, actor)
}

// Edit Doc: https://dev.vk.com/method/wall.edit
func (w *Wall) Edit(actor actor.Actor) *request.WallEditRequest {
	return request.NewWallEditRequest(w.api, actor)
}

// EditAdsStealth Doc: https://dev.vk.com/method/wall.editAdsStealth
func (w *Wall) EditAdsStealth(actor actor.Actor) *request.WallEditAdsStealthRequest {
	return request.NewWallEditAdsStealthRequest(w.api, actor)
}

// EditComment Doc: https://dev.vk.com/method/wall.editComment
func (w *Wall) EditComment(actor actor.Actor) *request.WallEditCommentRequest {
	return request.NewWallEditCommentRequest(w.api, actor)
}

// Get Doc: https://dev.vk.com/method/wall.get
func (w *Wall) Get(actor actor.Actor) *request.WallGetRequest {
	return request.NewWallGetRequest(w.api, actor)
}

// GetExtended Doc: https://dev.vk.com/method/wall.get
func (w *Wall) GetExtended(actor actor.Actor) *request.WallGetExtendedRequest {
	return request.NewWallGetExtendedRequest(w.api, actor)
}

// GetById Doc: https://dev.vk.com/method/wall.getById
func (w *Wall) GetById(actor actor.Actor) *request.WallGetByIDRequest {
	return request.NewWallGetByIDRequest(w.api, actor)
}

// GetByIdExtended Doc: https://dev.vk.com/method/wall.getById
func (w *Wall) GetByIdExtended(actor actor.Actor) *request.WallGetByIDExtendedRequest {
	return request.NewWallGetByIDExtendedRequest(w.api, actor)
}

// GetComment Doc: https://dev.vk.com/method/wall.getComment
func (w *Wall) GetComment(actor actor.Actor) *request.WallGetCommentRequest {
	return request.NewWallGetCommentRequest(w.api, actor)
}

// GetCommentExtended Doc: https://dev.vk.com/method/wall.getComment
func (w *Wall) GetCommentExtended(actor actor.Actor) *request.WallGetCommentExtendedRequest {
	return request.NewWallGetCommentExtendedRequest(w.api, actor)
}

// GetComments Doc: https://dev.vk.com/method/wall.getComments
func (w *Wall) GetComments(actor actor.Actor) *request.WallGetCommentsRequest {
	return request.NewWallGetCommentsRequest(w.api, actor)
}

// GetCommentsExtended Doc: https://dev.vk.com/method/wall.getComments
func (w *Wall) GetCommentsExtended(actor actor.Actor) *request.WallGetCommentsExtendedRequest {
	return request.NewWallGetCommentsExtendedRequest(w.api, actor)
}

// GetReposts Doc: https://dev.vk.com/method/wall.getReposts
func (w *Wall) GetReposts(actor actor.Actor) *request.WallGetRepostsRequest {
	return request.NewWallGetRepostsRequest(w.api, actor)
}

// OpenComments Doc: https://dev.vk.com/method/wall.openComments
func (w *Wall) OpenComments(actor actor.Actor) *request.WallOpenCommentsRequest {
	return request.NewWallOpenCommentsRequest(w.api, actor)
}

// ParseAttachedLink Doc: https://dev.vk.com/method/wall.parseAttachedLink
func (w *Wall) ParseAttachedLink(actor actor.Actor) *request.WallParseAttachedLinkRequest {
	return request.NewWallParseAttachedLinkRequest(w.api, actor)
}

// Pin Doc: https://dev.vk.com/method/wall.pin
func (w *Wall) Pin(actor actor.Actor) *request.WallPinRequest {
	return request.NewWallPinRequest(w.api, actor)
}

// Post Doc: https://dev.vk.com/method/wall.post
func (w *Wall) Post(actor actor.Actor) *request.WallPostRequest {
	return request.NewWallPostRequest(w.api, actor)
}

// PostAdsStealth Doc: https://dev.vk.com/method/wall.postAdsStealth
func (w *Wall) PostAdsStealth(actor actor.Actor) *request.WallPostAdsStealthRequest {
	return request.NewWallPostAdsStealthRequest(w.api, actor)
}

// ReportComment Doc: https://dev.vk.com/method/wall.reportComment
func (w *Wall) ReportComment(actor actor.Actor) *request.WallReportCommentRequest {
	return request.NewWallReportCommentRequest(w.api, actor)
}

// ReportPost Doc: https://dev.vk.com/method/wall.reportPost
func (w *Wall) ReportPost(actor actor.Actor) *request.WallReportPostRequest {
	return request.NewWallReportPostRequest(w.api, actor)
}

// Repost Doc: https://dev.vk.com/method/wall.repost
func (w *Wall) Repost(actor actor.Actor) *request.WallRepostRequest {
	return request.NewWallRepostRequest(w.api, actor)
}

// Restore Doc: https://dev.vk.com/method/wall.restore
func (w *Wall) Restore(actor actor.Actor) *request.WallRestoreRequest {
	return request.NewWallRestoreRequest(w.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/method/wall.restoreComment
func (w *Wall) RestoreComment(actor actor.Actor) *request.WallRestoreCommentRequest {
	return request.NewWallRestoreCommentRequest(w.api, actor)
}

// Search Doc: https://dev.vk.com/method/wall.search
func (w *Wall) Search(actor actor.Actor) *request.WallSearchRequest {
	return request.NewWallSearchRequest(w.api, actor)
}

// SearchExtended Doc: https://dev.vk.com/method/wall.search
func (w *Wall) SearchExtended(actor actor.Actor) *request.WallSearchExtendedRequest {
	return request.NewWallSearchExtendedRequest(w.api, actor)
}

// Unpin Doc: https://dev.vk.com/method/wall.unpin
func (w *Wall) Unpin(actor actor.Actor) *request.WallUnpinRequest {
	return request.NewWallUnpinRequest(w.api, actor)
}
