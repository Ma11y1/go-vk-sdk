package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Photos Doc: https://dev.vk.com/ru/method/photos
type Photos struct {
	BaseAction
}

// ConfirmTag Doc: https://dev.vk.com/ru/method/photos.confirmTag
func (p *Photos) ConfirmTag(actor actor.Actor) *request.PhotosConfirmTagRequest {
	return request.NewPhotosConfirmTagRequest(p.api, actor)
}

// Copy Doc: https://dev.vk.com/ru/method/photos.copy
func (p *Photos) Copy(actor actor.Actor) *request.PhotosCopyRequest {
	return request.NewPhotosCopyRequest(p.api, actor)
}

// CreateAlbum Doc: https://dev.vk.com/ru/method/photos.createAlbum
func (p *Photos) CreateAlbum(actor actor.Actor) *request.PhotosCreateAlbumRequest {
	return request.NewPhotosCreateAlbumRequest(p.api, actor)
}

// CreateComment Doc: https://dev.vk.com/ru/method/photos.createComment
func (p *Photos) CreateComment(actor actor.Actor) *request.PhotosCreateCommentRequest {
	return request.NewPhotosCreateCommentRequest(p.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/photos.delete
func (p *Photos) Delete(actor actor.Actor) *request.PhotosDeleteRequest {
	return request.NewPhotosDeleteRequest(p.api, actor)
}

// DeleteAlbum Doc: https://dev.vk.com/ru/method/photos.deleteAlbum
func (p *Photos) DeleteAlbum(actor actor.Actor) *request.PhotosDeleteAlbumRequest {
	return request.NewPhotosDeleteAlbumRequest(p.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/ru/method/photos.deleteComment
func (p *Photos) DeleteComment(actor actor.Actor) *request.PhotosDeleteCommentRequest {
	return request.NewPhotosDeleteCommentRequest(p.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/photos.edit
func (p *Photos) Edit(actor actor.Actor) *request.PhotosEditRequest {
	return request.NewPhotosEditRequest(p.api, actor)
}

// EditAlbum Doc: https://dev.vk.com/ru/method/photos.editAlbum
func (p *Photos) EditAlbum(actor actor.Actor) *request.PhotosEditAlbumRequest {
	return request.NewPhotosEditAlbumRequest(p.api, actor)
}

// EditComment Doc: https://dev.vk.com/ru/method/photos.editComment
func (p *Photos) EditComment(actor actor.Actor) *request.PhotosEditCommentRequest {
	return request.NewPhotosEditCommentRequest(p.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/photos.get
func (p *Photos) Get(actor actor.Actor) *request.PhotosGetRequest {
	return request.NewPhotosGetRequest(p.api, actor)
}

// GetExtended Doc: https://dev.vk.com/ru/method/photos.get
func (p *Photos) GetExtended(actor actor.Actor) *request.PhotosGetExtendedRequest {
	return request.NewPhotosGetExtendedRequest(p.api, actor)
}

// GetAlbums Doc: https://dev.vk.com/ru/method/photos.getAlbums
func (p *Photos) GetAlbums(actor actor.Actor) *request.PhotosGetAlbumsRequest {
	return request.NewPhotosGetAlbumsRequest(p.api, actor)
}

// GetAlbumsCount Doc: https://dev.vk.com/ru/method/photos.getAlbumsCount
func (p *Photos) GetAlbumsCount(actor actor.Actor) *request.PhotosGetAlbumsCountRequest {
	return request.NewPhotosGetAlbumsCountRequest(p.api, actor)
}

// GetAll Doc: https://dev.vk.com/ru/method/photos.getAll
func (p *Photos) GetAll(actor actor.Actor) *request.PhotosGetAllRequest {
	return request.NewPhotosGetAllRequest(p.api, actor)
}

// GetAllExtended Doc: https://dev.vk.com/ru/method/photos.getAll
func (p *Photos) GetAllExtended(actor actor.Actor) *request.PhotosGetAllExtendedRequest {
	return request.NewPhotosGetAllExtendedRequest(p.api, actor)
}

// GetAllComments Doc: https://dev.vk.com/ru/method/photos.getAllComments
func (p *Photos) GetAllComments(actor actor.Actor) *request.PhotosGetAllCommentsRequest {
	return request.NewPhotosGetAllCommentsRequest(p.api, actor)
}

// GetByID Doc: https://dev.vk.com/ru/method/photos.getById
func (p *Photos) GetByID(actor actor.Actor) *request.PhotosGetByIDRequest {
	return request.NewPhotosGetByIDRequest(p.api, actor)
}

// GetByIDExtended Doc: https://dev.vk.com/ru/method/photos.getById
func (p *Photos) GetByIDExtended(actor actor.Actor) *request.PhotosGetByIDExtendedRequest {
	return request.NewPhotosGetByIDExtendedRequest(p.api, actor)
}

// GetChatUploadServer Doc: https://dev.vk.com/ru/method/photos.getChatUploadServer
func (p *Photos) GetChatUploadServer(actor actor.Actor) *request.PhotosGetChatUploadServerRequest {
	return request.NewPhotosGetChatUploadServerRequest(p.api, actor)
}

// GetComments Doc: https://dev.vk.com/ru/method/photos.getComments
func (p *Photos) GetComments(actor actor.Actor) *request.PhotosGetCommentsRequest {
	return request.NewPhotosGetCommentsRequest(p.api, actor)
}

// GetCommentsExtended Doc: https://dev.vk.com/ru/method/photos.getComments
func (p *Photos) GetCommentsExtended(actor actor.Actor) *request.PhotosGetCommentsExtendedRequest {
	return request.NewPhotosGetCommentsExtendedRequest(p.api, actor)
}

// GetMarketAlbumUploadServer Doc: https://dev.vk.com/ru/method/photos.getMarketAlbumUploadServer
func (p *Photos) GetMarketAlbumUploadServer(actor actor.Actor) *request.PhotosGetMarketAlbumUploadServerRequest {
	return request.NewPhotosGetMarketAlbumUploadServerRequest(p.api, actor)
}

// GetMessagesUploadServer Doc: https://dev.vk.com/ru/method/photos.getMessagesUploadServer
func (p *Photos) GetMessagesUploadServer(actor actor.Actor) *request.PhotosGetMessagesUploadServerRequest {
	return request.NewPhotosGetMessagesUploadServerRequest(p.api, actor)
}

// GetNewTags Doc: https://dev.vk.com/ru/method/photos.getNewTags
func (p *Photos) GetNewTags(actor actor.Actor) *request.PhotosGetNewTagsRequest {
	return request.NewPhotosGetNewTagsRequest(p.api, actor)
}

// GetOwnerCoverPhotoUploadServer Doc: https://dev.vk.com/ru/method/photos.getOwnerCoverPhotoUploadServer
func (p *Photos) GetOwnerCoverPhotoUploadServer(actor actor.Actor) *request.PhotosGetOwnerCoverPhotoUploadServerRequest {
	return request.NewPhotosGetOwnerCoverPhotoUploadServerRequest(p.api, actor)
}

// GetOwnerPhotoUploadServer Doc: https://dev.vk.com/ru/method/photos.getOwnerPhotoUploadServer
func (p *Photos) GetOwnerPhotoUploadServer(actor actor.Actor) *request.PhotosGetOwnerPhotoUploadServerRequest {
	return request.NewPhotosGetOwnerPhotoUploadServerRequest(p.api, actor)
}

// GetTags Doc: https://dev.vk.com/ru/method/photos.getTags
func (p *Photos) GetTags(actor actor.Actor) *request.PhotosGetTagsRequest {
	return request.NewPhotosGetTagsRequest(p.api, actor)
}

// GetUploadServer Doc: https://dev.vk.com/ru/method/photos.getUploadServer
func (p *Photos) GetUploadServer(actor actor.Actor) *request.PhotosGetUploadServerRequest {
	return request.NewPhotosGetUploadServerRequest(p.api, actor)
}

// GetUserPhotos Doc: https://dev.vk.com/ru/method/photos.getUserPhotos
func (p *Photos) GetUserPhotos(actor actor.Actor) *request.PhotosGetUserPhotosRequest {
	return request.NewPhotosGetUserPhotosRequest(p.api, actor)
}

// GetUserPhotosExtended Doc: https://dev.vk.com/ru/method/photos.getUserPhotos
func (p *Photos) GetUserPhotosExtended(actor actor.Actor) *request.PhotosGetUserPhotosExtendedRequest {
	return request.NewPhotosGetUserPhotosExtendedRequest(p.api, actor)
}

// GetWallUploadServer Doc: https://dev.vk.com/ru/method/photos.getWallUploadServer
func (p *Photos) GetWallUploadServer(actor actor.Actor) *request.PhotosGetWallUploadServerRequest {
	return request.NewPhotosGetWallUploadServerRequest(p.api, actor)
}

// MakeCover Doc: https://dev.vk.com/ru/method/photos.makeCover
func (p *Photos) MakeCover(actor actor.Actor) *request.PhotosMakeCoverRequest {
	return request.NewPhotosMakeCoverRequest(p.api, actor)
}

// Move Doc: https://dev.vk.com/ru/method/photos.move
func (p *Photos) Move(actor actor.Actor) *request.PhotosMoveRequest {
	return request.NewPhotosMoveRequest(p.api, actor)
}

// PutTag Doc: https://dev.vk.com/ru/method/photos.putTag
func (p *Photos) PutTag(actor actor.Actor) *request.PhotosPutTagRequest {
	return request.NewPhotosPutTagRequest(p.api, actor)
}

// RemoveTag Doc: https://dev.vk.com/ru/method/photos.removeTag
func (p *Photos) RemoveTag(actor actor.Actor) *request.PhotosRemoveTagRequest {
	return request.NewPhotosRemoveTagRequest(p.api, actor)
}

// ReorderAlbums Doc: https://dev.vk.com/ru/method/photos.reorderAlbums
func (p *Photos) ReorderAlbums(actor actor.Actor) *request.PhotosReorderAlbumsRequest {
	return request.NewPhotosReorderAlbumsRequest(p.api, actor)
}

// ReorderPhotos Doc: https://dev.vk.com/ru/method/photos.reorderPhotos
func (p *Photos) ReorderPhotos(actor actor.Actor) *request.PhotosReorderPhotosRequest {
	return request.NewPhotosReorderPhotosRequest(p.api, actor)
}

// Report Doc: https://dev.vk.com/ru/method/photos.report
func (p *Photos) Report(actor actor.Actor) *request.PhotosReportRequest {
	return request.NewPhotosReportRequest(p.api, actor)
}

// ReportComment Doc: https://dev.vk.com/ru/method/photos.reportComment
func (p *Photos) ReportComment(actor actor.Actor) *request.PhotosReportCommentRequest {
	return request.NewPhotosReportCommentRequest(p.api, actor)
}

// Restore Doc: https://dev.vk.com/ru/method/photos.restore
func (p *Photos) Restore(actor actor.Actor) *request.PhotosRestoreRequest {
	return request.NewPhotosRestoreRequest(p.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/ru/method/photos.restoreComment
func (p *Photos) RestoreComment(actor actor.Actor) *request.PhotosRestoreCommentRequest {
	return request.NewPhotosRestoreCommentRequest(p.api, actor)
}

// Save Doc: https://dev.vk.com/ru/method/photos.save
func (p *Photos) Save(actor actor.Actor) *request.PhotosSaveRequest {
	return request.NewPhotosSaveRequest(p.api, actor)
}

// SaveMarketAlbumPhoto Doc: https://dev.vk.com/ru/method/photos.saveMarketAlbumPhoto
func (p *Photos) SaveMarketAlbumPhoto(actor actor.Actor) *request.PhotosSaveMarketAlbumPhotoRequest {
	return request.NewPhotosSaveMarketAlbumPhotoRequest(p.api, actor)
}

// SaveMessagesPhoto Doc: https://dev.vk.com/ru/method/photos.saveMessagesPhoto
func (p *Photos) SaveMessagesPhoto(actor actor.Actor) *request.PhotosSaveMessagesPhotoRequest {
	return request.NewPhotosSaveMessagesPhotoRequest(p.api, actor)
}

// SaveOwnerCoverPhoto Doc: https://dev.vk.com/ru/method/photos.saveOwnerCoverPhoto
func (p *Photos) SaveOwnerCoverPhoto(actor actor.Actor) *request.PhotosSaveOwnerCoverPhotoRequest {
	return request.NewPhotosSaveOwnerCoverPhotoRequest(p.api, actor)
}

// SaveOwnerPhoto Doc: https://dev.vk.com/ru/method/photos.saveOwnerPhoto
func (p *Photos) SaveOwnerPhoto(actor actor.Actor) *request.PhotosSaveOwnerPhotoRequest {
	return request.NewPhotosSaveOwnerPhotoRequest(p.api, actor)
}

// SaveWallPhoto Doc: https://dev.vk.com/ru/method/photos.saveWallPhoto
func (p *Photos) SaveWallPhoto(actor actor.Actor) *request.PhotosSaveWallPhotoRequest {
	return request.NewPhotosSaveWallPhotoRequest(p.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/photos.search
func (p *Photos) Search(actor actor.Actor) *request.PhotosSearchRequest {
	return request.NewPhotosSearchRequest(p.api, actor)
}
