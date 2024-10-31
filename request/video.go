package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/video

// VideoAddRequest defines the request for video.add
//
// Adds a video to a user or community page.
// Doc: https://dev.vk.com/method/video.add
type VideoAddRequest struct {
	BaseRequest
}

// NewVideoAddRequest creates a new request for video.add
func NewVideoAddRequest(a *api.API, actor actor.Actor) *VideoAddRequest {
	return &VideoAddRequest{*NewMethodBaseRequest(a, actor, "video.add")}
}

// Exec executes the request and unmarshals the response into VideoAddResponse
func (r *VideoAddRequest) Exec(ctx context.Context) (response response.VideoAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoAddAlbumRequest defines the request for video.addAlbum
//
// Creates an empty album for videos.
// Doc: https://dev.vk.com/method/video.addAlbum
type VideoAddAlbumRequest struct {
	BaseRequest
}

// NewVideoAddAlbumRequest creates a new request for video.addAlbum
func NewVideoAddAlbumRequest(a *api.API, actor actor.Actor) *VideoAddAlbumRequest {
	return &VideoAddAlbumRequest{*NewMethodBaseRequest(a, actor, "video.addAlbum")}
}

// Exec executes the request and unmarshals the response into VideoAddAlbumResponse
func (r *VideoAddAlbumRequest) Exec(ctx context.Context) (response response.VideoAddAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoAddToAlbumRequest defines the request for video.addToAlbum
//
// Adds a video to an album.
// Doc: https://dev.vk.com/method/video.addToAlbum
type VideoAddToAlbumRequest struct {
	BaseRequest
}

// NewVideoAddToAlbumRequest creates a new request for video.addToAlbum
func NewVideoAddToAlbumRequest(a *api.API, actor actor.Actor) *VideoAddToAlbumRequest {
	return &VideoAddToAlbumRequest{*NewMethodBaseRequest(a, actor, "video.addToAlbum")}
}

// Exec executes the request and unmarshals the response into VideoAddToAlbumResponse
func (r *VideoAddToAlbumRequest) Exec(ctx context.Context) (response response.VideoAddToAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoCreateCommentRequest defines the request for video.createComment
//
// Adds a new comment on a video.
// Doc: https://dev.vk.com/method/video.createComment
type VideoCreateCommentRequest struct {
	BaseRequest
}

// NewVideoCreateCommentRequest creates a new request for video.createComment
func NewVideoCreateCommentRequest(a *api.API, actor actor.Actor) *VideoCreateCommentRequest {
	return &VideoCreateCommentRequest{*NewMethodBaseRequest(a, actor, "video.createComment")}
}

// Exec executes the request and unmarshals the response into VideoCreateCommentResponse
func (r *VideoCreateCommentRequest) Exec(ctx context.Context) (response response.VideoCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoDeleteRequest defines the request for video.delete
//
// Deletes a video.
// Doc: https://dev.vk.com/method/video.delete
type VideoDeleteRequest struct {
	BaseRequest
}

// NewVideoDeleteRequest creates a new request for video.delete
func NewVideoDeleteRequest(a *api.API, actor actor.Actor) *VideoDeleteRequest {
	return &VideoDeleteRequest{*NewMethodBaseRequest(a, actor, "video.delete")}
}

// Exec executes the request and unmarshals the response into VideoDeleteResponse
func (r *VideoDeleteRequest) Exec(ctx context.Context) (response response.VideoDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoDeleteAlbumRequest defines the request for video.deleteAlbum
//
// Deletes a video album.
// Doc: https://dev.vk.com/method/video.deleteAlbum
type VideoDeleteAlbumRequest struct {
	BaseRequest
}

// NewVideoDeleteAlbumRequest creates a new request for video.deleteAlbum
func NewVideoDeleteAlbumRequest(a *api.API, actor actor.Actor) *VideoDeleteAlbumRequest {
	return &VideoDeleteAlbumRequest{*NewMethodBaseRequest(a, actor, "video.deleteAlbum")}
}

// Exec executes the request and unmarshals the response into VideoDeleteAlbumResponse
func (r *VideoDeleteAlbumRequest) Exec(ctx context.Context) (response response.VideoDeleteAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoDeleteCommentRequest defines the request for video.deleteComment
//
// Deletes a comment on a video.
// Doc: https://dev.vk.com/method/video.deleteComment
type VideoDeleteCommentRequest struct {
	BaseRequest
}

// NewVideoDeleteCommentRequest creates a new request for video.deleteComment
func NewVideoDeleteCommentRequest(a *api.API, actor actor.Actor) *VideoDeleteCommentRequest {
	return &VideoDeleteCommentRequest{*NewMethodBaseRequest(a, actor, "video.deleteComment")}
}

// Exec executes the request and unmarshals the response into VideoDeleteCommentResponse
func (r *VideoDeleteCommentRequest) Exec(ctx context.Context) (response response.VideoDeleteCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoEditRequest defines the request for video.edit
//
// Edits information about a video.
// Doc: https://dev.vk.com/method/video.edit
type VideoEditRequest struct {
	BaseRequest
}

// NewVideoEditRequest creates a new request for video.edit
func NewVideoEditRequest(a *api.API, actor actor.Actor) *VideoEditRequest {
	return &VideoEditRequest{*NewMethodBaseRequest(a, actor, "video.edit")}
}

// Exec executes the request and unmarshals the response into VideoEditResponse
func (r *VideoEditRequest) Exec(ctx context.Context) (response response.VideoEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoEditAlbumRequest defines the request for video.editAlbum
//
// Edits information about a video album.
// Doc: https://dev.vk.com/method/video.editAlbum
type VideoEditAlbumRequest struct {
	BaseRequest
}

// NewVideoEditAlbumRequest creates a new request for video.editAlbum
func NewVideoEditAlbumRequest(a *api.API, actor actor.Actor) *VideoEditAlbumRequest {
	return &VideoEditAlbumRequest{*NewMethodBaseRequest(a, actor, "video.editAlbum")}
}

// Exec executes the request and unmarshals the response into VideoEditAlbumResponse
func (r *VideoEditAlbumRequest) Exec(ctx context.Context) (response response.VideoEditAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoEditCommentRequest defines the request for video.editComment
//
// Edits a comment on a video.
// Doc: https://dev.vk.com/method/video.editComment
type VideoEditCommentRequest struct {
	BaseRequest
}

// NewVideoEditCommentRequest creates a new request for video.editComment
func NewVideoEditCommentRequest(a *api.API, actor actor.Actor) *VideoEditCommentRequest {
	return &VideoEditCommentRequest{*NewMethodBaseRequest(a, actor, "video.editComment")}
}

// Exec executes the request and unmarshals the response into VideoEditCommentResponse
func (r *VideoEditCommentRequest) Exec(ctx context.Context) (response response.VideoEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetRequest defines the request for video.get
//
// Returns detailed information about videos.
// Doc: https://dev.vk.com/method/video.get
type VideoGetRequest struct {
	BaseRequest
}

// NewVideoGetRequest creates a new request for video.get
func NewVideoGetRequest(a *api.API, actor actor.Actor) *VideoGetRequest {
	return &VideoGetRequest{*NewMethodBaseRequest(a, actor, "video.get")}
}

// Exec executes the request and unmarshals the response into VideoGetResponse
func (r *VideoGetRequest) Exec(ctx context.Context) (response response.VideoGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetExtendedRequest defines the request for video.get
//
// Returns detailed information about videos.
// Doc: https://dev.vk.com/method/video.get
type VideoGetExtendedRequest struct {
	BaseRequest
}

// NewVideoGetExtendedRequest creates a new request for video.get
func NewVideoGetExtendedRequest(a *api.API, actor actor.Actor) *VideoGetExtendedRequest {
	return &VideoGetExtendedRequest{*NewMethodBaseRequest(a, actor, "video.get")}
}

// Exec executes the request and unmarshals the response into VideoGetExtendedResponse
func (r *VideoGetExtendedRequest) Exec(ctx context.Context) (response response.VideoGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetAlbumByIDRequest defines the request for video.getAlbumById
//
// Returns data about a video album by its ID.
// Doc: https://dev.vk.com/method/video.getAlbumById
type VideoGetAlbumByIDRequest struct {
	BaseRequest
}

// NewVideoGetAlbumByIDRequest creates a new request for video.getAlbumById
func NewVideoGetAlbumByIDRequest(a *api.API, actor actor.Actor) *VideoGetAlbumByIDRequest {
	return &VideoGetAlbumByIDRequest{*NewMethodBaseRequest(a, actor, "video.getAlbumById")}
}

// Exec executes the request and unmarshals the response into VideoGetAlbumByIdResponse
func (r *VideoGetAlbumByIDRequest) Exec(ctx context.Context) (response response.VideoGetAlbumByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetAlbumsRequest defines the request for video.getAlbums
//
// Returns a list of video albums owned by a user or community.
// Doc: https://dev.vk.com/method/video.getAlbums
type VideoGetAlbumsRequest struct {
	BaseRequest
}

// NewVideoGetAlbumsRequest creates a new request for video.getAlbums
func NewVideoGetAlbumsRequest(a *api.API, actor actor.Actor) *VideoGetAlbumsRequest {
	return &VideoGetAlbumsRequest{*NewMethodBaseRequest(a, actor, "video.getAlbums")}
}

// Exec executes the request and unmarshals the response into VideoGetAlbumsResponse
func (r *VideoGetAlbumsRequest) Exec(ctx context.Context) (response response.VideoGetAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetAlbumsExtendedRequest defines the request for video.getAlbums
//
// Returns a list of video albums owned by a user or community.
// Doc: https://dev.vk.com/method/video.getAlbums
type VideoGetAlbumsExtendedRequest struct {
	BaseRequest
}

// NewVideoGetAlbumsExtendedRequest creates a new request for video.getAlbums
func NewVideoGetAlbumsExtendedRequest(a *api.API, actor actor.Actor) *VideoGetAlbumsExtendedRequest {
	return &VideoGetAlbumsExtendedRequest{*NewMethodBaseRequest(a, actor, "video.getAlbums")}
}

// Exec executes the request and unmarshals the response into VideoGetAlbumsExtendedResponse
func (r *VideoGetAlbumsExtendedRequest) Exec(ctx context.Context) (response response.VideoGetAlbumsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetAlbumsByVideoRequest defines the request for video.getAlbumsByVideo
//
// Returns a list of albums that include a particular video.
// Doc: https://dev.vk.com/method/video.getAlbumsByVideo
type VideoGetAlbumsByVideoRequest struct {
	BaseRequest
}

// NewVideoGetAlbumsByVideoRequest creates a new request for video.getAlbumsByVideo
func NewVideoGetAlbumsByVideoRequest(a *api.API, actor actor.Actor) *VideoGetAlbumsByVideoRequest {
	return &VideoGetAlbumsByVideoRequest{*NewMethodBaseRequest(a, actor, "video.getAlbumsByVideo")}
}

// Exec executes the request and unmarshals the response into VideoGetAlbumsByVideoResponse
func (r *VideoGetAlbumsByVideoRequest) Exec(ctx context.Context) (response response.VideoGetAlbumsByVideoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetAlbumsByVideoExtendedRequest defines the request for video.getAlbumsByVideo
//
// Returns a list of albums that include a particular video.
// Doc: https://dev.vk.com/method/video.getAlbumsByVideo
type VideoGetAlbumsByVideoExtendedRequest struct {
	BaseRequest
}

// NewVideoGetAlbumsByVideoExtendedRequest creates a new request for video.getAlbumsByVideo
func NewVideoGetAlbumsByVideoExtendedRequest(a *api.API, actor actor.Actor) *VideoGetAlbumsByVideoExtendedRequest {
	return &VideoGetAlbumsByVideoExtendedRequest{*NewMethodBaseRequest(a, actor, "video.getAlbumsByVideo")}
}

// Exec executes the request and unmarshals the response into VideoGetAlbumsByVideoExtendedResponse
func (r *VideoGetAlbumsByVideoExtendedRequest) Exec(ctx context.Context) (response response.VideoGetAlbumsByVideoExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetCommentsRequest defines the request for video.getComments
//
// Returns a list of comments on a video.
// Doc: https://dev.vk.com/method/video.getComments
type VideoGetCommentsRequest struct {
	BaseRequest
}

// NewVideoGetCommentsRequest creates a new request for video.getComments
func NewVideoGetCommentsRequest(a *api.API, actor actor.Actor) *VideoGetCommentsRequest {
	return &VideoGetCommentsRequest{*NewMethodBaseRequest(a, actor, "video.getComments")}
}

// Exec executes the request and unmarshals the response into VideoGetCommentsResponse
func (r *VideoGetCommentsRequest) Exec(ctx context.Context) (response response.VideoGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetCommentsExtendedRequest defines the request for video.getComments
//
// Returns a list of comments on a video.
// Doc: https://dev.vk.com/method/video.getComments
type VideoGetCommentsExtendedRequest struct {
	BaseRequest
}

// NewVideoGetCommentsExtendedRequest creates a new request for video.getComments
func NewVideoGetCommentsExtendedRequest(a *api.API, actor actor.Actor) *VideoGetCommentsExtendedRequest {
	return &VideoGetCommentsExtendedRequest{*NewMethodBaseRequest(a, actor, "video.getComments")}
}

// Exec executes the request and unmarshals the response into VideoGetCommentsExtendedResponse
func (r *VideoGetCommentsExtendedRequest) Exec(ctx context.Context) (response response.VideoGetCommentsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetLongPollServerRequest defines the request for video.getLongPollServer
//
// Returns data required to connect to the video long-polling server.
// Doc: https://dev.vk.com/method/video.getLongPollServer
type VideoGetLongPollServerRequest struct {
	BaseRequest
}

// NewVideoGetLongPollServerRequest creates a new request for video.getLongPollServer
func NewVideoGetLongPollServerRequest(a *api.API, actor actor.Actor) *VideoGetLongPollServerRequest {
	return &VideoGetLongPollServerRequest{*NewMethodBaseRequest(a, actor, "video.getLongPollServer")}
}

// Exec executes the request and unmarshals the response into VideoGetLongPollServerResponse
func (r *VideoGetLongPollServerRequest) Exec(ctx context.Context) (response response.VideoGetLongPollServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoGetThumbUploadUrlRequest defines the request for video.getThumbUploadUrl TODO New method not yet documented
//
// Returns the URL for uploading a video thumbnail.
// Doc: https://dev.vk.com/method/video.getThumbUploadUrl
type VideoGetThumbUploadUrlRequest struct {
	BaseRequest
}

// NewVideoGetThumbUploadUrlRequest creates a new request for video.getThumbUploadUrl
func NewVideoGetThumbUploadUrlRequest(a *api.API, actor actor.Actor) *VideoGetThumbUploadUrlRequest {
	return &VideoGetThumbUploadUrlRequest{*NewMethodBaseRequest(a, actor, "video.getThumbUploadUrl")}
}

// Exec executes the request and unmarshals the response into VideoGetThumbUploadUrlResponse
func (r *VideoGetThumbUploadUrlRequest) Exec(ctx context.Context) (response response.VideoGetThumbUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoLiveGetCategoriesRequest defines the request for video.liveGetCategories
//
// Returns a list of categories for live broadcasts.
// Doc: https://dev.vk.com/method/video.liveGetCategories
type VideoLiveGetCategoriesRequest struct {
	BaseRequest
}

// NewVideoLiveGetCategoriesRequest creates a new request for video.liveGetCategories
func NewVideoLiveGetCategoriesRequest(a *api.API, actor actor.Actor) *VideoLiveGetCategoriesRequest {
	return &VideoLiveGetCategoriesRequest{*NewMethodBaseRequest(a, actor, "video.liveGetCategories")}
}

// Exec executes the request and unmarshals the response into VideoLiveGetCategoriesResponse
func (r *VideoLiveGetCategoriesRequest) Exec(ctx context.Context) (response response.VideoLiveGetCategoriesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoRemoveFromAlbumRequest defines the request for video.removeFromAlbum
//
// Removes a video from an album.
// Doc: https://dev.vk.com/method/video.removeFromAlbum
type VideoRemoveFromAlbumRequest struct {
	BaseRequest
}

// NewVideoRemoveFromAlbumRequest creates a new request for video.removeFromAlbum
func NewVideoRemoveFromAlbumRequest(a *api.API, actor actor.Actor) *VideoRemoveFromAlbumRequest {
	return &VideoRemoveFromAlbumRequest{*NewMethodBaseRequest(a, actor, "video.removeFromAlbum")}
}

// Exec executes the request and unmarshals the response into VideoRemoveFromAlbumResponse
func (r *VideoRemoveFromAlbumRequest) Exec(ctx context.Context) (response response.VideoRemoveFromAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoReorderAlbumsRequest defines the request for video.reorderAlbums
//
// Changes the order of video albums.
// Doc: https://dev.vk.com/method/video.reorderAlbums
type VideoReorderAlbumsRequest struct {
	BaseRequest
}

// NewVideoReorderAlbumsRequest creates a new request for video.reorderAlbums
func NewVideoReorderAlbumsRequest(a *api.API, actor actor.Actor) *VideoReorderAlbumsRequest {
	return &VideoReorderAlbumsRequest{*NewMethodBaseRequest(a, actor, "video.reorderAlbums")}
}

// Exec executes the request and unmarshals the response into VideoReorderAlbumsResponse
func (r *VideoReorderAlbumsRequest) Exec(ctx context.Context) (response response.VideoReorderAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoReorderVideosRequest defines the request for video.reorderVideos
//
// Changes the order of videos in an album.
// Doc: https://dev.vk.com/method/video.reorderVideos
type VideoReorderVideosRequest struct {
	BaseRequest
}

// NewVideoReorderVideosRequest creates a new request for video.reorderVideos
func NewVideoReorderVideosRequest(a *api.API, actor actor.Actor) *VideoReorderVideosRequest {
	return &VideoReorderVideosRequest{*NewMethodBaseRequest(a, actor, "video.reorderVideos")}
}

// Exec executes the request and unmarshals the response into VideoReorderVideosResponse
func (r *VideoReorderVideosRequest) Exec(ctx context.Context) (response response.VideoReorderVideosResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoReportRequest defines the request for video.report
//
// Reports a video for inappropriate content.
// Doc: https://dev.vk.com/method/video.report
type VideoReportRequest struct {
	BaseRequest
}

// NewVideoReportRequest creates a new request for video.report
func NewVideoReportRequest(a *api.API, actor actor.Actor) *VideoReportRequest {
	return &VideoReportRequest{*NewMethodBaseRequest(a, actor, "video.report")}
}

// Exec executes the request and unmarshals the response into VideoReportResponse
func (r *VideoReportRequest) Exec(ctx context.Context) (response response.VideoReportResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoReportCommentRequest defines the request for video.reportComment
//
// Reports a comment on a video as inappropriate.
// Doc: https://dev.vk.com/method/video.reportComment
type VideoReportCommentRequest struct {
	BaseRequest
}

// NewVideoReportCommentRequest creates a new request for video.reportComment
func NewVideoReportCommentRequest(a *api.API, actor actor.Actor) *VideoReportCommentRequest {
	return &VideoReportCommentRequest{*NewMethodBaseRequest(a, actor, "video.reportComment")}
}

// Exec executes the request and unmarshals the response into VideoReportCommentResponse
func (r *VideoReportCommentRequest) Exec(ctx context.Context) (response response.VideoReportCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoRestoreRequest defines the request for video.restore
//
// Restores a previously deleted video.
// Doc: https://dev.vk.com/method/video.restore
type VideoRestoreRequest struct {
	BaseRequest
}

// NewVideoRestoreRequest creates a new request for video.restore
func NewVideoRestoreRequest(a *api.API, actor actor.Actor) *VideoRestoreRequest {
	return &VideoRestoreRequest{*NewMethodBaseRequest(a, actor, "video.restore")}
}

// Exec executes the request and unmarshals the response into VideoRestoreResponse
func (r *VideoRestoreRequest) Exec(ctx context.Context) (response response.VideoRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoRestoreCommentRequest defines the request for video.restoreComment
//
// Restores a previously deleted comment on a video.
// Doc: https://dev.vk.com/method/video.restoreComment
type VideoRestoreCommentRequest struct {
	BaseRequest
}

// NewVideoRestoreCommentRequest creates a new request for video.restoreComment
func NewVideoRestoreCommentRequest(a *api.API, actor actor.Actor) *VideoRestoreCommentRequest {
	return &VideoRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "video.restoreComment")}
}

// Exec executes the request and unmarshals the response into VideoRestoreCommentResponse
func (r *VideoRestoreCommentRequest) Exec(ctx context.Context) (response response.VideoRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoSaveRequest defines the request for video.save
//
// Saves information about a video in the user's profile or community.
// Doc: https://dev.vk.com/method/video.save
type VideoSaveRequest struct {
	BaseRequest
}

// NewVideoSaveRequest creates a new request for video.save
func NewVideoSaveRequest(a *api.API, actor actor.Actor) *VideoSaveRequest {
	return &VideoSaveRequest{*NewMethodBaseRequest(a, actor, "video.save")}
}

// Exec executes the request and unmarshals the response into VideoSaveResponse
func (r *VideoSaveRequest) Exec(ctx context.Context) (response response.VideoSaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoSaveUploadedThumbRequest defines the request for video.saveUploadedThumb TODO New method not yet documented
//
// Saves a thumbnail uploaded for a video.
// Doc: https://dev.vk.com/method/video.saveUploadedThumb
type VideoSaveUploadedThumbRequest struct {
	BaseRequest
}

// NewVideoSaveUploadedThumbRequest creates a new request for video.saveUploadedThumb
func NewVideoSaveUploadedThumbRequest(a *api.API, actor actor.Actor) *VideoSaveUploadedThumbRequest {
	return &VideoSaveUploadedThumbRequest{*NewMethodBaseRequest(a, actor, "video.saveUploadedThumb")}
}

// Exec executes the request and unmarshals the response into VideoSaveUploadedThumbResponse
func (r *VideoSaveUploadedThumbRequest) Exec(ctx context.Context) (response response.VideoSaveUploadedThumbResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoSearchRequest defines the request for video.search
//
// Searches for videos in the VK database by a given query.
// Doc: https://dev.vk.com/method/video.search
type VideoSearchRequest struct {
	BaseRequest
}

// NewVideoSearchRequest creates a new request for video.search
func NewVideoSearchRequest(a *api.API, actor actor.Actor) *VideoSearchRequest {
	return &VideoSearchRequest{*NewMethodBaseRequest(a, actor, "video.search")}
}

// Exec executes the request and unmarshals the response into VideoSearchResponse
func (r *VideoSearchRequest) Exec(ctx context.Context) (response response.VideoSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoSearchExtendedRequest defines the request for video.search
//
// Searches for videos in the VK database by a given query.
// Doc: https://dev.vk.com/method/video.search
type VideoSearchExtendedRequest struct {
	BaseRequest
}

// NewVideoSearchExtendedRequest creates a new request for video.search
func NewVideoSearchExtendedRequest(a *api.API, actor actor.Actor) *VideoSearchExtendedRequest {
	return &VideoSearchExtendedRequest{*NewMethodBaseRequest(a, actor, "video.search")}
}

// Exec executes the request and unmarshals the response into VideoSearchExtendedResponse
func (r *VideoSearchExtendedRequest) Exec(ctx context.Context) (response response.VideoSearchExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoStartStreamingRequest defines the request for video.startStreaming
//
// Starts streaming a live video.
// Doc: https://dev.vk.com/method/video.startStreaming
type VideoStartStreamingRequest struct {
	BaseRequest
}

// NewVideoStartStreamingRequest creates a new request for video.startStreaming
func NewVideoStartStreamingRequest(a *api.API, actor actor.Actor) *VideoStartStreamingRequest {
	return &VideoStartStreamingRequest{*NewMethodBaseRequest(a, actor, "video.startStreaming")}
}

// Exec executes the request and unmarshals the response into VideoStartStreamingResponse
func (r *VideoStartStreamingRequest) Exec(ctx context.Context) (response response.VideoStartStreamingResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VideoStopStreamingRequest defines the request for video.stopStreaming
//
// Stops an active live video stream.
// Doc: https://dev.vk.com/method/video.stopStreaming
type VideoStopStreamingRequest struct {
	BaseRequest
}

// NewVideoStopStreamingRequest creates a new request for video.stopStreaming
func NewVideoStopStreamingRequest(a *api.API, actor actor.Actor) *VideoStopStreamingRequest {
	return &VideoStopStreamingRequest{*NewMethodBaseRequest(a, actor, "video.stopStreaming")}
}

// Exec executes the request and unmarshals the response into VideoStopStreamingResponse
func (r *VideoStopStreamingRequest) Exec(ctx context.Context) (response response.VideoStopStreamingResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
