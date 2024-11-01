package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Video Doc: https://dev.vk.com/ru/method/video
type Video struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/method/video.add
func (v *Video) Add(actor actor.Actor) *request.VideoAddRequest {
	return request.NewVideoAddRequest(v.api, actor)
}

// AddAlbum Doc: https://dev.vk.com/method/video.addAlbum
func (v *Video) AddAlbum(actor actor.Actor) *request.VideoAddAlbumRequest {
	return request.NewVideoAddAlbumRequest(v.api, actor)
}

// AddToAlbum Doc: https://dev.vk.com/method/video.addToAlbum
func (v *Video) AddToAlbum(actor actor.Actor) *request.VideoAddToAlbumRequest {
	return request.NewVideoAddToAlbumRequest(v.api, actor)
}

// CreateComment Doc: https://dev.vk.com/method/video.createComment
func (v *Video) CreateComment(actor actor.Actor) *request.VideoCreateCommentRequest {
	return request.NewVideoCreateCommentRequest(v.api, actor)
}

// Delete Doc: https://dev.vk.com/method/video.delete
func (v *Video) Delete(actor actor.Actor) *request.VideoDeleteRequest {
	return request.NewVideoDeleteRequest(v.api, actor)
}

// DeleteAlbum Doc: https://dev.vk.com/method/video.deleteAlbum
func (v *Video) DeleteAlbum(actor actor.Actor) *request.VideoDeleteAlbumRequest {
	return request.NewVideoDeleteAlbumRequest(v.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/method/video.deleteComment
func (v *Video) DeleteComment(actor actor.Actor) *request.VideoDeleteCommentRequest {
	return request.NewVideoDeleteCommentRequest(v.api, actor)
}

// Edit Doc: https://dev.vk.com/method/video.edit
func (v *Video) Edit(actor actor.Actor) *request.VideoEditRequest {
	return request.NewVideoEditRequest(v.api, actor)
}

// EditAlbum Doc: https://dev.vk.com/method/video.editAlbum
func (v *Video) EditAlbum(actor actor.Actor) *request.VideoEditAlbumRequest {
	return request.NewVideoEditAlbumRequest(v.api, actor)
}

// EditComment Doc: https://dev.vk.com/method/video.editComment
func (v *Video) EditComment(actor actor.Actor) *request.VideoEditCommentRequest {
	return request.NewVideoEditCommentRequest(v.api, actor)
}

// Get Doc: https://dev.vk.com/method/video.get
func (v *Video) Get(actor actor.Actor) *request.VideoGetRequest {
	return request.NewVideoGetRequest(v.api, actor)
}

// GetExtended Doc: https://dev.vk.com/method/video.get
func (v *Video) GetExtended(actor actor.Actor) *request.VideoGetExtendedRequest {
	return request.NewVideoGetExtendedRequest(v.api, actor)
}

// GetAlbumById Doc: https://dev.vk.com/method/video.getAlbumById
func (v *Video) GetAlbumById(actor actor.Actor) *request.VideoGetAlbumByIDRequest {
	return request.NewVideoGetAlbumByIDRequest(v.api, actor)
}

// GetAlbums Doc: https://dev.vk.com/method/video.getAlbums
func (v *Video) GetAlbums(actor actor.Actor) *request.VideoGetAlbumsRequest {
	return request.NewVideoGetAlbumsRequest(v.api, actor)
}

// GetAlbumsExtended Doc: https://dev.vk.com/method/video.getAlbums
func (v *Video) GetAlbumsExtended(actor actor.Actor) *request.VideoGetAlbumsExtendedRequest {
	return request.NewVideoGetAlbumsExtendedRequest(v.api, actor)
}

// GetAlbumsByVideo Doc: https://dev.vk.com/method/video.getAlbumsByVideo
func (v *Video) GetAlbumsByVideo(actor actor.Actor) *request.VideoGetAlbumsByVideoRequest {
	return request.NewVideoGetAlbumsByVideoRequest(v.api, actor)
}

// GetAlbumsByVideoExtended Doc: https://dev.vk.com/method/video.getAlbumsByVideo
func (v *Video) GetAlbumsByVideoExtended(actor actor.Actor) *request.VideoGetAlbumsByVideoExtendedRequest {
	return request.NewVideoGetAlbumsByVideoExtendedRequest(v.api, actor)
}

// GetComments Doc: https://dev.vk.com/method/video.getComments
func (v *Video) GetComments(actor actor.Actor) *request.VideoGetCommentsRequest {
	return request.NewVideoGetCommentsRequest(v.api, actor)
}

// GetCommentsExtended Doc: https://dev.vk.com/method/video.getComments
func (v *Video) GetCommentsExtended(actor actor.Actor) *request.VideoGetCommentsExtendedRequest {
	return request.NewVideoGetCommentsExtendedRequest(v.api, actor)
}

// GetLongPollServer Doc: https://dev.vk.com/method/video.getLongPollServer
func (v *Video) GetLongPollServer(actor actor.Actor) *request.VideoGetLongPollServerRequest {
	return request.NewVideoGetLongPollServerRequest(v.api, actor)
}

// GetThumbUploadUrl Doc: https://dev.vk.com/method/video.getThumbUploadUrl
func (v *Video) GetThumbUploadUrl(actor actor.Actor) *request.VideoGetThumbUploadUrlRequest {
	return request.NewVideoGetThumbUploadUrlRequest(v.api, actor)
}

// LiveGetCategories Doc: https://dev.vk.com/method/video.liveGetCategories
func (v *Video) LiveGetCategories(actor actor.Actor) *request.VideoLiveGetCategoriesRequest {
	return request.NewVideoLiveGetCategoriesRequest(v.api, actor)
}

// RemoveFromAlbum Doc: https://dev.vk.com/method/video.removeFromAlbum
func (v *Video) RemoveFromAlbum(actor actor.Actor) *request.VideoRemoveFromAlbumRequest {
	return request.NewVideoRemoveFromAlbumRequest(v.api, actor)
}

// ReorderAlbums Doc: https://dev.vk.com/method/video.reorderAlbums
func (v *Video) ReorderAlbums(actor actor.Actor) *request.VideoReorderAlbumsRequest {
	return request.NewVideoReorderAlbumsRequest(v.api, actor)
}

// ReorderVideos Doc: https://dev.vk.com/method/video.reorderVideos
func (v *Video) ReorderVideos(actor actor.Actor) *request.VideoReorderVideosRequest {
	return request.NewVideoReorderVideosRequest(v.api, actor)
}

// Report Doc: https://dev.vk.com/method/video.report
func (v *Video) Report(actor actor.Actor) *request.VideoReportRequest {
	return request.NewVideoReportRequest(v.api, actor)
}

// ReportComment Doc: https://dev.vk.com/method/video.reportComment
func (v *Video) ReportComment(actor actor.Actor) *request.VideoReportCommentRequest {
	return request.NewVideoReportCommentRequest(v.api, actor)
}

// Restore Doc: https://dev.vk.com/method/video.restore
func (v *Video) Restore(actor actor.Actor) *request.VideoRestoreRequest {
	return request.NewVideoRestoreRequest(v.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/method/video.restoreComment
func (v *Video) RestoreComment(actor actor.Actor) *request.VideoRestoreCommentRequest {
	return request.NewVideoRestoreCommentRequest(v.api, actor)
}

// Save Doc: https://dev.vk.com/method/video.save
func (v *Video) Save(actor actor.Actor) *request.VideoSaveRequest {
	return request.NewVideoSaveRequest(v.api, actor)
}

// SaveUploadedThumb Doc: https://dev.vk.com/method/video.saveUploadedThumb
func (v *Video) SaveUploadedThumb(actor actor.Actor) *request.VideoSaveUploadedThumbRequest {
	return request.NewVideoSaveUploadedThumbRequest(v.api, actor)
}

// Search Doc: https://dev.vk.com/method/video.search
func (v *Video) Search(actor actor.Actor) *request.VideoSearchRequest {
	return request.NewVideoSearchRequest(v.api, actor)
}

// SearchExtended Doc: https://dev.vk.com/method/video.search
func (v *Video) SearchExtended(actor actor.Actor) *request.VideoSearchExtendedRequest {
	return request.NewVideoSearchExtendedRequest(v.api, actor)
}

// StartStreaming Doc: https://dev.vk.com/method/video.startStreaming
func (v *Video) StartStreaming(actor actor.Actor) *request.VideoStartStreamingRequest {
	return request.NewVideoStartStreamingRequest(v.api, actor)
}

// StopStreaming Doc: https://dev.vk.com/method/video.stopStreaming
func (v *Video) StopStreaming(actor actor.Actor) *request.VideoStopStreamingRequest {
	return request.NewVideoStopStreamingRequest(v.api, actor)
}
