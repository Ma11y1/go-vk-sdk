package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"strconv"
)

// Doc: https://dev.vk.com/method/photos

// PhotosConfirmTagRequest defines the request for photos.confirmTag
//
// Confirms a tag on a photo.
// Doc: https://dev.vk.com/method/photos.confirmTag
type PhotosConfirmTagRequest struct {
	BaseRequest
}

// NewPhotosConfirmTagRequest creates a new request for photos.confirmTag
func NewPhotosConfirmTagRequest(a *api.API, actor actor.Actor) *PhotosConfirmTagRequest {
	return &PhotosConfirmTagRequest{*NewMethodBaseRequest(a, actor, "photos.confirmTag")}
}

// Exec executes the request and unmarshals the response into PhotosConfirmTagResponse
func (r *PhotosConfirmTagRequest) Exec(ctx context.Context) (response response.PhotosConfirmTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosCopyRequest defines the request for photos.copy
//
// Copies a photo to the "Saved Photos" album.
// Doc: https://dev.vk.com/method/photos.copy
type PhotosCopyRequest struct {
	BaseRequest
}

// NewPhotosCopyRequest creates a new request for photos.copy
func NewPhotosCopyRequest(a *api.API, actor actor.Actor) *PhotosCopyRequest {
	return &PhotosCopyRequest{*NewMethodBaseRequest(a, actor, "photos.copy")}
}

// Exec executes the request and unmarshals the response into PhotosCopyResponse
func (r *PhotosCopyRequest) Exec(ctx context.Context) (response response.PhotosCopyResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosCreateAlbumRequest defines the request for photos.createAlbum
//
// Creates an empty photo album.
// Doc: https://dev.vk.com/method/photos.createAlbum
type PhotosCreateAlbumRequest struct {
	BaseRequest
}

// NewPhotosCreateAlbumRequest creates a new request for photos.createAlbum
func NewPhotosCreateAlbumRequest(a *api.API, actor actor.Actor) *PhotosCreateAlbumRequest {
	return &PhotosCreateAlbumRequest{*NewMethodBaseRequest(a, actor, "photos.createAlbum")}
}

// Exec executes the request and unmarshals the response into PhotosCreateAlbumResponse
func (r *PhotosCreateAlbumRequest) Exec(ctx context.Context) (response response.PhotosCreateAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosCreateCommentRequest defines the request for photos.createComment
//
// Creates a new comment for a photo.
// Doc: https://dev.vk.com/method/photos.createComment
type PhotosCreateCommentRequest struct {
	BaseRequest
}

// NewPhotosCreateCommentRequest creates a new request for photos.createComment
func NewPhotosCreateCommentRequest(a *api.API, actor actor.Actor) *PhotosCreateCommentRequest {
	return &PhotosCreateCommentRequest{*NewMethodBaseRequest(a, actor, "photos.createComment")}
}

// Exec executes the request and unmarshals the response into PhotosCreateCommentResponse
func (r *PhotosCreateCommentRequest) Exec(ctx context.Context) (response response.PhotosCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosDeleteRequest defines the request for photos.delete
//
// Deletes a photo.
// Doc: https://dev.vk.com/method/photos.delete
type PhotosDeleteRequest struct {
	BaseRequest
}

// NewPhotosDeleteRequest creates a new request for photos.delete
func NewPhotosDeleteRequest(a *api.API, actor actor.Actor) *PhotosDeleteRequest {
	return &PhotosDeleteRequest{*NewMethodBaseRequest(a, actor, "photos.delete")}
}

// Exec executes the request and unmarshals the response into PhotosDeleteResponse
func (r *PhotosDeleteRequest) Exec(ctx context.Context) (response response.PhotosDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosDeleteAlbumRequest defines the request for photos.deleteAlbum
//
// Deletes a specified photo album.
// Doc: https://dev.vk.com/method/photos.deleteAlbum
type PhotosDeleteAlbumRequest struct {
	BaseRequest
}

// NewPhotosDeleteAlbumRequest creates a new request for photos.deleteAlbum
func NewPhotosDeleteAlbumRequest(a *api.API, actor actor.Actor) *PhotosDeleteAlbumRequest {
	return &PhotosDeleteAlbumRequest{*NewMethodBaseRequest(a, actor, "photos.deleteAlbum")}
}

// Exec executes the request and unmarshals the response into PhotosDeleteAlbumResponse
func (r *PhotosDeleteAlbumRequest) Exec(ctx context.Context) (response response.PhotosDeleteAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosDeleteCommentRequest defines the request for photos.deleteComment
//
// Deletes a comment on a photo.
// Doc: https://dev.vk.com/method/photos.deleteComment
type PhotosDeleteCommentRequest struct {
	BaseRequest
}

// NewPhotosDeleteCommentRequest creates a new request for photos.deleteComment
func NewPhotosDeleteCommentRequest(a *api.API, actor actor.Actor) *PhotosDeleteCommentRequest {
	return &PhotosDeleteCommentRequest{*NewMethodBaseRequest(a, actor, "photos.deleteComment")}
}

// Exec executes the request and unmarshals the response into PhotosDeleteCommentResponse
func (r *PhotosDeleteCommentRequest) Exec(ctx context.Context) (response response.PhotosDeleteCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosEditRequest defines the request for photos.edit
//
// Edits the description or geotag of a photo.
// Doc: https://dev.vk.com/method/photos.edit
type PhotosEditRequest struct {
	BaseRequest
}

// NewPhotosEditRequest creates a new request for photos.edit
func NewPhotosEditRequest(a *api.API, actor actor.Actor) *PhotosEditRequest {
	return &PhotosEditRequest{*NewMethodBaseRequest(a, actor, "photos.edit")}
}

// Exec executes the request and unmarshals the response into PhotosEditResponse
func (r *PhotosEditRequest) Exec(ctx context.Context) (response response.PhotosEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosEditAlbumRequest defines the request for photos.editAlbum
//
// Edits album details.
// Doc: https://dev.vk.com/method/photos.editAlbum
type PhotosEditAlbumRequest struct {
	BaseRequest
}

// NewPhotosEditAlbumRequest creates a new request for photos.editAlbum
func NewPhotosEditAlbumRequest(a *api.API, actor actor.Actor) *PhotosEditAlbumRequest {
	return &PhotosEditAlbumRequest{*NewMethodBaseRequest(a, actor, "photos.editAlbum")}
}

// Exec executes the request and unmarshals the response into PhotosEditAlbumResponse
func (r *PhotosEditAlbumRequest) Exec(ctx context.Context) (response response.PhotosEditAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosEditCommentRequest defines the request for photos.editComment
//
// Edits the text of a photo comment.
// Doc: https://dev.vk.com/method/photos.editComment
type PhotosEditCommentRequest struct {
	BaseRequest
}

// NewPhotosEditCommentRequest creates a new request for photos.editComment
func NewPhotosEditCommentRequest(a *api.API, actor actor.Actor) *PhotosEditCommentRequest {
	return &PhotosEditCommentRequest{*NewMethodBaseRequest(a, actor, "photos.editComment")}
}

// Exec executes the request and unmarshals the response into PhotosEditCommentResponse
func (r *PhotosEditCommentRequest) Exec(ctx context.Context) (response response.PhotosEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetRequest defines the request for photos.get
//
// Returns a list of photos in an album.
// Doc: https://dev.vk.com/method/photos.get
type PhotosGetRequest struct {
	BaseRequest
}

// NewPhotosGetRequest creates a new request for photos.get
func NewPhotosGetRequest(a *api.API, actor actor.Actor) *PhotosGetRequest {
	return &PhotosGetRequest{*NewMethodBaseRequest(a, actor, "photos.get")}
}

// Exec executes the request and unmarshals the response into PhotosGetResponse
func (r *PhotosGetRequest) Exec(ctx context.Context) (response response.PhotosGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetExtendedRequest defines the request for photos.get
//
// Returns a list of photos in an album.
// Doc: https://dev.vk.com/method/photos.get
type PhotosGetExtendedRequest struct {
	BaseRequest
}

// NewPhotosGetExtendedRequest creates a new request for photos.get
func NewPhotosGetExtendedRequest(a *api.API, actor actor.Actor) *PhotosGetExtendedRequest {
	r := &PhotosGetExtendedRequest{*NewMethodBaseRequest(a, actor, "photos.get")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into PhotosGetExtendedResponse
func (r *PhotosGetExtendedRequest) Exec(ctx context.Context) (response response.PhotosGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetAlbumsRequest defines the request for photos.getAlbums
//
// Returns a list of photo albums.
// Doc: https://dev.vk.com/method/photos.getAlbums
type PhotosGetAlbumsRequest struct {
	BaseRequest
}

// NewPhotosGetAlbumsRequest creates a new request for photos.getAlbums
func NewPhotosGetAlbumsRequest(a *api.API, actor actor.Actor) *PhotosGetAlbumsRequest {
	return &PhotosGetAlbumsRequest{*NewMethodBaseRequest(a, actor, "photos.getAlbums")}
}

// Exec executes the request and unmarshals the response into PhotosGetAlbumsResponse
func (r *PhotosGetAlbumsRequest) Exec(ctx context.Context) (response response.PhotosGetAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetAlbumsCountRequest defines the request for photos.getAlbumsCount
//
// Returns the number of available albums.
// Doc: https://dev.vk.com/method/photos.getAlbumsCount
type PhotosGetAlbumsCountRequest struct {
	BaseRequest
}

// NewPhotosGetAlbumsCountRequest creates a new request for photos.getAlbumsCount
func NewPhotosGetAlbumsCountRequest(a *api.API, actor actor.Actor) *PhotosGetAlbumsCountRequest {
	return &PhotosGetAlbumsCountRequest{*NewMethodBaseRequest(a, actor, "photos.getAlbumsCount")}
}

// Exec executes the request and unmarshals the response into PhotosGetAlbumsCountResponse
func (r *PhotosGetAlbumsCountRequest) Exec(ctx context.Context) (response response.PhotosGetAlbumsCountResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetAllRequest defines the request for photos.getAll
//
// Returns all photos in reverse chronological order.
// Doc: https://dev.vk.com/method/photos.getAll
type PhotosGetAllRequest struct {
	BaseRequest
}

// NewPhotosGetAllRequest creates a new request for photos.getAll
func NewPhotosGetAllRequest(a *api.API, actor actor.Actor) *PhotosGetAllRequest {
	return &PhotosGetAllRequest{*NewMethodBaseRequest(a, actor, "photos.getAll")}
}

// Exec executes the request and unmarshals the response into PhotosGetAllResponse
func (r *PhotosGetAllRequest) Exec(ctx context.Context) (response response.PhotosGetAllResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetAllExtendedRequest defines the request for photos.getAll
//
// Returns all photos in reverse chronological order.
// Doc: https://dev.vk.com/method/photos.getAll
type PhotosGetAllExtendedRequest struct {
	BaseRequest
}

// NewPhotosGetAllExtendedRequest creates a new request for photos.getAll
func NewPhotosGetAllExtendedRequest(a *api.API, actor actor.Actor) *PhotosGetAllExtendedRequest {
	r := &PhotosGetAllExtendedRequest{*NewMethodBaseRequest(a, actor, "photos.getAll")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into PhotosGetAllExtendedResponse
func (r *PhotosGetAllExtendedRequest) Exec(ctx context.Context) (response response.PhotosGetAllExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetAllCommentsRequest defines the request for photos.getAllComments
//
// Returns all comments in reverse chronological order.
// Doc: https://dev.vk.com/method/photos.getAllComments
type PhotosGetAllCommentsRequest struct {
	BaseRequest
}

// NewPhotosGetAllCommentsRequest creates a new request for photos.getAllComments
func NewPhotosGetAllCommentsRequest(a *api.API, actor actor.Actor) *PhotosGetAllCommentsRequest {
	return &PhotosGetAllCommentsRequest{*NewMethodBaseRequest(a, actor, "photos.getAllComments")}
}

// Exec executes the request and unmarshals the response into PhotosGetAllCommentsResponse
func (r *PhotosGetAllCommentsRequest) Exec(ctx context.Context) (response response.PhotosGetAllCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetByIDRequest defines the request for photos.getById
//
// Returns information about photos by their IDs.
// Doc: https://dev.vk.com/method/photos.getById
type PhotosGetByIDRequest struct {
	BaseRequest
}

// NewPhotosGetByIDRequest creates a new request for photos.getById
func NewPhotosGetByIDRequest(a *api.API, actor actor.Actor) *PhotosGetByIDRequest {
	return &PhotosGetByIDRequest{*NewMethodBaseRequest(a, actor, "photos.getById")}
}

// Exec executes the request and unmarshals the response into PhotosGetByIDResponse
func (r *PhotosGetByIDRequest) Exec(ctx context.Context) (response response.PhotosGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetByIDExtendedRequest defines the request for photos.getById
//
// Returns information about photos by their IDs.
// Doc: https://dev.vk.com/method/photos.getById
type PhotosGetByIDExtendedRequest struct {
	BaseRequest
}

// NewPhotosGetByIDExtendedRequest creates a new request for photos.getById
func NewPhotosGetByIDExtendedRequest(a *api.API, actor actor.Actor) *PhotosGetByIDExtendedRequest {
	r := &PhotosGetByIDExtendedRequest{*NewMethodBaseRequest(a, actor, "photos.getById")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into PhotosGetByIDExtendedResponse
func (r *PhotosGetByIDExtendedRequest) Exec(ctx context.Context) (response response.PhotosGetByIDExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetChatUploadServerRequest defines the request for photos.getChatUploadServer
//
// Gets the url for uploading a chat cover.
// Doc: https://dev.vk.com/method/photos.getChatUploadServer
type PhotosGetChatUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetChatUploadServerRequest creates a new request for photos.getChatUploadServer
func NewPhotosGetChatUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetChatUploadServerRequest {
	return &PhotosGetChatUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getChatUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetChatUploadServerResponse
func (r *PhotosGetChatUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetChatUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetCommentsRequest defines the request for photos.getComments
//
// Returns a list of comments on a photo.
// Doc: https://dev.vk.com/method/photos.getComments
type PhotosGetCommentsRequest struct {
	BaseRequest
}

// NewPhotosGetCommentsRequest creates a new request for photos.getComments
func NewPhotosGetCommentsRequest(a *api.API, actor actor.Actor) *PhotosGetCommentsRequest {
	return &PhotosGetCommentsRequest{*NewMethodBaseRequest(a, actor, "photos.getComments")}
}

// Exec executes the request and unmarshals the response into PhotosGetCommentsResponse
func (r *PhotosGetCommentsRequest) Exec(ctx context.Context) (response response.PhotosGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetCommentsExtendedRequest defines the request for photos.getComments
//
// Returns a list of comments on a photo.
// Doc: https://dev.vk.com/method/photos.getComments
type PhotosGetCommentsExtendedRequest struct {
	BaseRequest
}

// NewPhotosGetCommentsExtendedRequest creates a new request for photos.getComments
func NewPhotosGetCommentsExtendedRequest(a *api.API, actor actor.Actor) *PhotosGetCommentsExtendedRequest {
	r := &PhotosGetCommentsExtendedRequest{*NewMethodBaseRequest(a, actor, "photos.getComments")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into PhotosGetCommentsExtendedResponse
func (r *PhotosGetCommentsExtendedRequest) Exec(ctx context.Context) (response response.PhotosGetCommentsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetMarketAlbumUploadServerRequest defines the request for photos.getMarketAlbumUploadServer
//
// Gets the url for uploading a market album photo.
// Doc: https://dev.vk.com/method/photos.getMarketAlbumUploadServer
type PhotosGetMarketAlbumUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetMarketAlbumUploadServerRequest creates a new request for photos.getMarketAlbumUploadServer
func NewPhotosGetMarketAlbumUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetMarketAlbumUploadServerRequest {
	return &PhotosGetMarketAlbumUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getMarketAlbumUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetMarketAlbumUploadServerResponse
func (r *PhotosGetMarketAlbumUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetMarketAlbumUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetMessagesUploadServerRequest defines the request for photos.getMessagesUploadServer
//
// Gets the url for uploading a photo in a private message.
// Doc: https://dev.vk.com/method/photos.getMessagesUploadServer
type PhotosGetMessagesUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetMessagesUploadServerRequest creates a new request for photos.getMessagesUploadServer
func NewPhotosGetMessagesUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetMessagesUploadServerRequest {
	return &PhotosGetMessagesUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getMessagesUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetMessagesUploadServerResponse
func (r *PhotosGetMessagesUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetMessagesUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetNewTagsRequest defines the request for photos.getNewTags
//
// Returns a list of photos with unviewed tags.
// Doc: https://dev.vk.com/method/photos.getNewTags
type PhotosGetNewTagsRequest struct {
	BaseRequest
}

// NewPhotosGetNewTagsRequest creates a new request for photos.getNewTags
func NewPhotosGetNewTagsRequest(a *api.API, actor actor.Actor) *PhotosGetNewTagsRequest {
	return &PhotosGetNewTagsRequest{*NewMethodBaseRequest(a, actor, "photos.getNewTags")}
}

// Exec executes the request and unmarshals the response into PhotosGetNewTagsResponse
func (r *PhotosGetNewTagsRequest) Exec(ctx context.Context) (response response.PhotosGetNewTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetOwnerCoverPhotoUploadServerRequest defines the request for photos.getOwnerCoverPhotoUploadServer
//
// Gets the url for uploading a community cover photo.
// Doc: https://dev.vk.com/method/photos.getOwnerCoverPhotoUploadServer
type PhotosGetOwnerCoverPhotoUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetOwnerCoverPhotoUploadServerRequest creates a new request for photos.getOwnerCoverPhotoUploadServer
func NewPhotosGetOwnerCoverPhotoUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetOwnerCoverPhotoUploadServerRequest {
	return &PhotosGetOwnerCoverPhotoUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getOwnerCoverPhotoUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetOwnerCoverPhotoUploadServerResponse
func (r *PhotosGetOwnerCoverPhotoUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetOwnerCoverPhotoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetOwnerPhotoUploadServerRequest defines the request for photos.getOwnerPhotoUploadServer
//
// Gets the url for uploading the main photo on a user or community page.
// Doc: https://dev.vk.com/method/photos.getOwnerPhotoUploadServer
type PhotosGetOwnerPhotoUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetOwnerPhotoUploadServerRequest creates a new request for photos.getOwnerPhotoUploadServer
func NewPhotosGetOwnerPhotoUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetOwnerPhotoUploadServerRequest {
	return &PhotosGetOwnerPhotoUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getOwnerPhotoUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetOwnerPhotoUploadServerResponse
func (r *PhotosGetOwnerPhotoUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetOwnerPhotoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetTagsRequest defines the request for photos.getTags
//
// Returns a list of tags on a photo.
// Doc: https://dev.vk.com/method/photos.getTags
type PhotosGetTagsRequest struct {
	BaseRequest
}

// NewPhotosGetTagsRequest creates a new request for photos.getTags
func NewPhotosGetTagsRequest(a *api.API, actor actor.Actor) *PhotosGetTagsRequest {
	return &PhotosGetTagsRequest{*NewMethodBaseRequest(a, actor, "photos.getTags")}
}

// Exec executes the request and unmarshals the response into PhotosGetTagsResponse
func (r *PhotosGetTagsRequest) Exec(ctx context.Context) (response response.PhotosGetTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetUploadServerRequest defines the request for photos.getUploadServer
//
//	Gets the url for uploading photos to an album.
//	To upload photos to a community, enter the community ID in the group_id parameter.
//	Doc: https://dev.vk.com/method/photos.getUploadServer
type PhotosGetUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetUploadServerRequest creates a new request for photos.getUploadServer
func NewPhotosGetUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetUploadServerRequest {
	return &PhotosGetUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetUploadServerResponse
func (r *PhotosGetUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetUserPhotosRequest defines the request for photos.getUserPhotos
//
// Returns a list of photos where the user is tagged.
// Doc: https://dev.vk.com/method/photos.getUserPhotos
type PhotosGetUserPhotosRequest struct {
	BaseRequest
}

// NewPhotosGetUserPhotosRequest creates a new request for photos.getUserPhotos
func NewPhotosGetUserPhotosRequest(a *api.API, actor actor.Actor) *PhotosGetUserPhotosRequest {
	return &PhotosGetUserPhotosRequest{*NewMethodBaseRequest(a, actor, "photos.getUserPhotos")}
}

// Exec executes the request and unmarshals the response into PhotosGetUserPhotosResponse
func (r *PhotosGetUserPhotosRequest) Exec(ctx context.Context) (response response.PhotosGetUserPhotosResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetUserPhotosExtendedRequest defines the request for photos.getUserPhotos
//
// Returns a list of photos where the user is tagged.
// Doc: https://dev.vk.com/method/photos.getUserPhotos
type PhotosGetUserPhotosExtendedRequest struct {
	BaseRequest
}

// NewPhotosGetUserPhotosExtendedRequest creates a new request for photos.getUserPhotos
func NewPhotosGetUserPhotosExtendedRequest(a *api.API, actor actor.Actor) *PhotosGetUserPhotosExtendedRequest {
	r := &PhotosGetUserPhotosExtendedRequest{*NewMethodBaseRequest(a, actor, "photos.getUserPhotos")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into PhotosGetUserPhotosExtendedResponse
func (r *PhotosGetUserPhotosExtendedRequest) Exec(ctx context.Context) (response response.PhotosGetUserPhotosExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosGetWallUploadServerRequest defines the request for photos.getWallUploadServer
//
// Gets the url for uploading photos to a user's or community's wall.
// Doc: https://dev.vk.com/method/photos.getWallUploadServer
type PhotosGetWallUploadServerRequest struct {
	BaseRequest
}

// NewPhotosGetWallUploadServerRequest creates a new request for photos.getWallUploadServer
func NewPhotosGetWallUploadServerRequest(a *api.API, actor actor.Actor) *PhotosGetWallUploadServerRequest {
	return &PhotosGetWallUploadServerRequest{*NewMethodBaseRequest(a, actor, "photos.getWallUploadServer")}
}

// Exec executes the request and unmarshals the response into PhotosGetWallUploadServerResponse
func (r *PhotosGetWallUploadServerRequest) Exec(ctx context.Context) (response response.PhotosGetWallUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosMakeCoverRequest defines the request for photos.makeCover
//
// Makes a photo the album cover.
// Doc: https://dev.vk.com/method/photos.makeCover
type PhotosMakeCoverRequest struct {
	BaseRequest
}

// NewPhotosMakeCoverRequest creates a new request for photos.makeCover
func NewPhotosMakeCoverRequest(a *api.API, actor actor.Actor) *PhotosMakeCoverRequest {
	return &PhotosMakeCoverRequest{*NewMethodBaseRequest(a, actor, "photos.makeCover")}
}

// Exec executes the request and unmarshals the response into PhotosMakeCoverResponse
func (r *PhotosMakeCoverRequest) Exec(ctx context.Context) (response response.PhotosMakeCoverResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosMoveRequest defines the request for photos.move
//
// Moves a photo from one album to another.
// Doc: https://dev.vk.com/method/photos.move
type PhotosMoveRequest struct {
	BaseRequest
}

// NewPhotosMoveRequest creates a new request for photos.move
func NewPhotosMoveRequest(a *api.API, actor actor.Actor) *PhotosMoveRequest {
	return &PhotosMoveRequest{*NewMethodBaseRequest(a, actor, "photos.move")}
}

// Exec executes the request and unmarshals the response into PhotosMoveResponse
func (r *PhotosMoveRequest) Exec(ctx context.Context) (response response.PhotosMoveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosPutTagRequest defines the request for photos.putTag
//
// Adds a tag to a photo.
// Doc: https://dev.vk.com/method/photos.putTag
type PhotosPutTagRequest struct {
	BaseRequest
}

// NewPhotosPutTagRequest creates a new request for photos.putTag
func NewPhotosPutTagRequest(a *api.API, actor actor.Actor) *PhotosPutTagRequest {
	return &PhotosPutTagRequest{*NewMethodBaseRequest(a, actor, "photos.putTag")}
}

// Exec executes the request and unmarshals the response into PhotosPutTagResponse
func (r *PhotosPutTagRequest) Exec(ctx context.Context) (response response.PhotosPutTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosRemoveTagRequest defines the request for photos.removeTag
//
// Removes a tag from a photo.
// Doc: https://dev.vk.com/method/photos.removeTag
type PhotosRemoveTagRequest struct {
	BaseRequest
}

// NewPhotosRemoveTagRequest creates a new request for photos.removeTag
func NewPhotosRemoveTagRequest(a *api.API, actor actor.Actor) *PhotosRemoveTagRequest {
	return &PhotosRemoveTagRequest{*NewMethodBaseRequest(a, actor, "photos.removeTag")}
}

// Exec executes the request and unmarshals the response into PhotosRemoveTagResponse
func (r *PhotosRemoveTagRequest) Exec(ctx context.Context) (response response.PhotosRemoveTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosReorderAlbumsRequest defines the request for photos.reorderAlbums
//
// Changes the order of albums in the user's album list.
// Doc: https://dev.vk.com/method/photos.reorderAlbums
type PhotosReorderAlbumsRequest struct {
	BaseRequest
}

// NewPhotosReorderAlbumsRequest creates a new request for photos.reorderAlbums
func NewPhotosReorderAlbumsRequest(a *api.API, actor actor.Actor) *PhotosReorderAlbumsRequest {
	return &PhotosReorderAlbumsRequest{*NewMethodBaseRequest(a, actor, "photos.reorderAlbums")}
}

// Exec executes the request and unmarshals the response into PhotosReorderAlbumsResponse
func (r *PhotosReorderAlbumsRequest) Exec(ctx context.Context) (response response.PhotosReorderAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosReorderPhotosRequest defines the request for photos.reorderPhotos
//
// Changes the order of photos in an album.
// Doc: https://dev.vk.com/method/photos.reorderPhotos
type PhotosReorderPhotosRequest struct {
	BaseRequest
}

// NewPhotosReorderPhotosRequest creates a new request for photos.reorderPhotos
func NewPhotosReorderPhotosRequest(a *api.API, actor actor.Actor) *PhotosReorderPhotosRequest {
	return &PhotosReorderPhotosRequest{*NewMethodBaseRequest(a, actor, "photos.reorderPhotos")}
}

// Exec executes the request and unmarshals the response into PhotosReorderPhotosResponse
func (r *PhotosReorderPhotosRequest) Exec(ctx context.Context) (response response.PhotosReorderPhotosResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosReportRequest defines the request for photos.report
//
// Reports a photo.
// Doc: https://dev.vk.com/method/photos.report
type PhotosReportRequest struct {
	BaseRequest
}

// NewPhotosReportRequest creates a new request for photos.report
func NewPhotosReportRequest(a *api.API, actor actor.Actor) *PhotosReportRequest {
	return &PhotosReportRequest{*NewMethodBaseRequest(a, actor, "photos.report")}
}

// Exec executes the request and unmarshals the response into PhotosReportResponse
func (r *PhotosReportRequest) Exec(ctx context.Context) (response response.PhotosReportResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosReportCommentRequest defines the request for photos.reportComment
//
// Reports a comment on a photo.
// Doc: https://dev.vk.com/method/photos.reportComment
type PhotosReportCommentRequest struct {
	BaseRequest
}

// NewPhotosReportCommentRequest creates a new request for photos.reportComment
func NewPhotosReportCommentRequest(a *api.API, actor actor.Actor) *PhotosReportCommentRequest {
	return &PhotosReportCommentRequest{*NewMethodBaseRequest(a, actor, "photos.reportComment")}
}

// Exec executes the request and unmarshals the response into PhotosReportCommentResponse
func (r *PhotosReportCommentRequest) Exec(ctx context.Context) (response response.PhotosReportCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosRestoreRequest defines the request for photos.restore
//
// Restores a deleted photo.
// Doc: https://dev.vk.com/method/photos.restore
type PhotosRestoreRequest struct {
	BaseRequest
}

// NewPhotosRestoreRequest creates a new request for photos.restore
func NewPhotosRestoreRequest(a *api.API, actor actor.Actor) *PhotosRestoreRequest {
	return &PhotosRestoreRequest{*NewMethodBaseRequest(a, actor, "photos.restore")}
}

// Exec executes the request and unmarshals the response into PhotosRestoreResponse
func (r *PhotosRestoreRequest) Exec(ctx context.Context) (response response.PhotosRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosRestoreCommentRequest defines the request for photos.restoreComment
//
// Restores a deleted comment on a photo.
// Doc: https://dev.vk.com/method/photos.restoreComment
type PhotosRestoreCommentRequest struct {
	BaseRequest
}

// NewPhotosRestoreCommentRequest creates a new request for photos.restoreComment
func NewPhotosRestoreCommentRequest(a *api.API, actor actor.Actor) *PhotosRestoreCommentRequest {
	return &PhotosRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "photos.restoreComment")}
}

// Exec executes the request and unmarshals the response into PhotosRestoreCommentResponse
func (r *PhotosRestoreCommentRequest) Exec(ctx context.Context) (response response.PhotosRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosSaveRequest defines the request for photos.save
//
// Saves photos to an album after they have been successfully uploaded to the server.
// Doc: https://dev.vk.com/method/photos.save
type PhotosSaveRequest struct {
	BaseRequest
}

// NewPhotosSaveRequest creates a new request for photos.save
func NewPhotosSaveRequest(a *api.API, actor actor.Actor) *PhotosSaveRequest {
	return &PhotosSaveRequest{*NewMethodBaseRequest(a, actor, "photos.save")}
}

// Exec executes the request and unmarshals the response into PhotosSaveResponse
func (r *PhotosSaveRequest) Exec(ctx context.Context) (response response.PhotosSaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *PhotosSaveRequest) AlbumID(id int) *PhotosSaveRequest {
	r.parameters.Set(constants.ParameterNameAlbumID, strconv.Itoa(id))
	return r
}

func (r *PhotosSaveRequest) GroupID(id int) *PhotosSaveRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// PhotosSaveMarketAlbumPhotoRequest defines the request for photos.saveMarketAlbumPhoto
//
// Saves a market album photo after it has been successfully uploaded.
// Doc: https://dev.vk.com/method/photos.saveMarketAlbumPhoto
type PhotosSaveMarketAlbumPhotoRequest struct {
	BaseRequest
}

// NewPhotosSaveMarketAlbumPhotoRequest creates a new request for photos.saveMarketAlbumPhoto
func NewPhotosSaveMarketAlbumPhotoRequest(a *api.API, actor actor.Actor) *PhotosSaveMarketAlbumPhotoRequest {
	return &PhotosSaveMarketAlbumPhotoRequest{*NewMethodBaseRequest(a, actor, "photos.saveMarketAlbumPhoto")}
}

// Exec executes the request and unmarshals the response into PhotosSaveMarketAlbumPhotoResponse
func (r *PhotosSaveMarketAlbumPhotoRequest) Exec(ctx context.Context) (response response.PhotosSaveMarketAlbumPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosSaveMessagesPhotoRequest defines the request for photos.saveMessagesPhoto
//
// Saves a photo in a private message after it has been successfully uploaded.
// Doc: https://dev.vk.com/method/photos.saveMessagesPhoto
type PhotosSaveMessagesPhotoRequest struct {
	BaseRequest
}

// NewPhotosSaveMessagesPhotoRequest creates a new request for photos.saveMessagesPhoto
func NewPhotosSaveMessagesPhotoRequest(a *api.API, actor actor.Actor) *PhotosSaveMessagesPhotoRequest {
	return &PhotosSaveMessagesPhotoRequest{*NewMethodBaseRequest(a, actor, "photos.saveMessagesPhoto")}
}

// Exec executes the request and unmarshals the response into PhotosSaveMessagesPhotoResponse
func (r *PhotosSaveMessagesPhotoRequest) Exec(ctx context.Context) (response response.PhotosSaveMessagesPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosSaveOwnerCoverPhotoRequest defines the request for photos.saveOwnerCoverPhoto
//
// Saves a community or profile cover after it has been successfully uploaded.
// Doc: https://dev.vk.com/method/photos.saveOwnerCoverPhoto
type PhotosSaveOwnerCoverPhotoRequest struct {
	BaseRequest
}

// NewPhotosSaveOwnerCoverPhotoRequest creates a new request for photos.saveOwnerCoverPhoto
func NewPhotosSaveOwnerCoverPhotoRequest(a *api.API, actor actor.Actor) *PhotosSaveOwnerCoverPhotoRequest {
	return &PhotosSaveOwnerCoverPhotoRequest{*NewMethodBaseRequest(a, actor, "photos.saveOwnerCoverPhoto")}
}

// Exec executes the request and unmarshals the response into PhotosSaveOwnerCoverPhotoResponse
func (r *PhotosSaveOwnerCoverPhotoRequest) Exec(ctx context.Context) (response response.PhotosSaveOwnerCoverPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosSaveOwnerPhotoRequest defines the request for photos.saveOwnerPhoto
//
// Saves the main photo after it has been successfully uploaded.
// Doc: https://dev.vk.com/method/photos.saveOwnerPhoto
type PhotosSaveOwnerPhotoRequest struct {
	BaseRequest
}

// NewPhotosSaveOwnerPhotoRequest creates a new request for photos.saveOwnerPhoto
func NewPhotosSaveOwnerPhotoRequest(a *api.API, actor actor.Actor) *PhotosSaveOwnerPhotoRequest {
	return &PhotosSaveOwnerPhotoRequest{*NewMethodBaseRequest(a, actor, "photos.saveOwnerPhoto")}
}

// Exec executes the request and unmarshals the response into PhotosSaveOwnerPhotoResponse
func (r *PhotosSaveOwnerPhotoRequest) Exec(ctx context.Context) (response response.PhotosSaveOwnerPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PhotosSaveWallPhotoRequest defines the request for photos.saveWallPhoto
//
// Saves photos to the wall after they have been successfully uploaded.
// Doc: https://dev.vk.com/method/photos.saveWallPhoto
type PhotosSaveWallPhotoRequest struct {
	BaseRequest
}

// NewPhotosSaveWallPhotoRequest creates a new request for photos.saveWallPhoto
func NewPhotosSaveWallPhotoRequest(a *api.API, actor actor.Actor) *PhotosSaveWallPhotoRequest {
	return &PhotosSaveWallPhotoRequest{*NewMethodBaseRequest(a, actor, "photos.saveWallPhoto")}
}

// Exec executes the request and unmarshals the response into PhotosSaveWallPhotoResponse
func (r *PhotosSaveWallPhotoRequest) Exec(ctx context.Context) (response response.PhotosSaveWallPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *PhotosSaveWallPhotoRequest) GroupID(id int) *PhotosSaveWallPhotoRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// PhotosSearchRequest defines the request for photos.search
//
// Searches for photos by location or description.
// Doc: https://dev.vk.com/method/photos.search
type PhotosSearchRequest struct {
	BaseRequest
}

// NewPhotosSearchRequest creates a new request for photos.search
func NewPhotosSearchRequest(a *api.API, actor actor.Actor) *PhotosSearchRequest {
	return &PhotosSearchRequest{*NewMethodBaseRequest(a, actor, "photos.search")}
}

// Exec executes the request and unmarshals the response into PhotosSearchResponse
func (r *PhotosSearchRequest) Exec(ctx context.Context) (response response.PhotosSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
