package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/wall

// WallCheckCopyrightLinkRequest defines the request for wall.checkCopyrightLink
//
// Checks if a given link has a valid copyright for posting.
// Doc: https://dev.vk.com/method/wall.checkCopyrightLink
type WallCheckCopyrightLinkRequest struct {
	BaseRequest
}

// NewWallCheckCopyrightLinkRequest creates a new request for wall.checkCopyrightLink
func NewWallCheckCopyrightLinkRequest(a *api.API, actor actor.Actor) *WallCheckCopyrightLinkRequest {
	return &WallCheckCopyrightLinkRequest{*NewMethodBaseRequest(a, actor, "wall.checkCopyrightLink")}
}

// Exec executes the request and unmarshals the response into WallCheckCopyrightLinkResponse
func (r *WallCheckCopyrightLinkRequest) Exec(ctx context.Context) (response response.WallCheckCopyrightLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallCloseCommentsRequest defines the request for wall.closeComments
//
// Closes comments for a post on the wall.
// Doc: https://dev.vk.com/method/wall.closeComments
type WallCloseCommentsRequest struct {
	BaseRequest
}

// NewWallCloseCommentsRequest creates a new request for wall.closeComments
func NewWallCloseCommentsRequest(a *api.API, actor actor.Actor) *WallCloseCommentsRequest {
	return &WallCloseCommentsRequest{*NewMethodBaseRequest(a, actor, "wall.closeComments")}
}

// Exec executes the request and unmarshals the response into WallCloseCommentsResponse
func (r *WallCloseCommentsRequest) Exec(ctx context.Context) (response response.WallCloseCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallCreateCommentRequest defines the request for wall.createComment
//
// Adds a new comment to a post on the wall.
// Doc: https://dev.vk.com/method/wall.createComment
type WallCreateCommentRequest struct {
	BaseRequest
}

// NewWallCreateCommentRequest creates a new request for wall.createComment
func NewWallCreateCommentRequest(a *api.API, actor actor.Actor) *WallCreateCommentRequest {
	return &WallCreateCommentRequest{*NewMethodBaseRequest(a, actor, "wall.createComment")}
}

// Exec executes the request and unmarshals the response into WallCreateCommentResponse
func (r *WallCreateCommentRequest) Exec(ctx context.Context) (response response.WallCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallDeleteRequest defines the request for wall.delete
//
// Deletes a post from a user or community wall.
// Doc: https://dev.vk.com/method/wall.delete
type WallDeleteRequest struct {
	BaseRequest
}

// NewWallDeleteRequest creates a new request for wall.delete
func NewWallDeleteRequest(a *api.API, actor actor.Actor) *WallDeleteRequest {
	return &WallDeleteRequest{*NewMethodBaseRequest(a, actor, "wall.delete")}
}

// Exec executes the request and unmarshals the response into WallDeleteResponse
func (r *WallDeleteRequest) Exec(ctx context.Context) (response response.WallDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallDeleteCommentRequest defines the request for wall.deleteComment
//
// Deletes a comment on a post on the wall.
// Doc: https://dev.vk.com/method/wall.deleteComment
type WallDeleteCommentRequest struct {
	BaseRequest
}

// NewWallDeleteCommentRequest creates a new request for wall.deleteComment
func NewWallDeleteCommentRequest(a *api.API, actor actor.Actor) *WallDeleteCommentRequest {
	return &WallDeleteCommentRequest{*NewMethodBaseRequest(a, actor, "wall.deleteComment")}
}

// Exec executes the request and unmarshals the response into WallDeleteCommentResponse
func (r *WallDeleteCommentRequest) Exec(ctx context.Context) (response response.WallDeleteCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallEditRequest defines the request for wall.edit
//
// Edits a post on a user or community wall.
// Doc: https://dev.vk.com/method/wall.edit
type WallEditRequest struct {
	BaseRequest
}

// NewWallEditRequest creates a new request for wall.edit
func NewWallEditRequest(a *api.API, actor actor.Actor) *WallEditRequest {
	return &WallEditRequest{*NewMethodBaseRequest(a, actor, "wall.edit")}
}

// Exec executes the request and unmarshals the response into WallEditResponse
func (r *WallEditRequest) Exec(ctx context.Context) (response response.WallEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallEditAdsStealthRequest defines the request for wall.editAdsStealth
//
// Edits a hidden (stealth) post on a wall.
// Doc: https://dev.vk.com/method/wall.editAdsStealth
type WallEditAdsStealthRequest struct {
	BaseRequest
}

// NewWallEditAdsStealthRequest creates a new request for wall.editAdsStealth
func NewWallEditAdsStealthRequest(a *api.API, actor actor.Actor) *WallEditAdsStealthRequest {
	return &WallEditAdsStealthRequest{*NewMethodBaseRequest(a, actor, "wall.editAdsStealth")}
}

// Exec executes the request and unmarshals the response into WallEditAdsStealthResponse
func (r *WallEditAdsStealthRequest) Exec(ctx context.Context) (response response.WallEditAdsStealthResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallEditCommentRequest defines the request for wall.editComment
//
// Edits a comment on a post.
// Doc: https://dev.vk.com/method/wall.editComment
type WallEditCommentRequest struct {
	BaseRequest
}

// NewWallEditCommentRequest creates a new request for wall.editComment
func NewWallEditCommentRequest(a *api.API, actor actor.Actor) *WallEditCommentRequest {
	return &WallEditCommentRequest{*NewMethodBaseRequest(a, actor, "wall.editComment")}
}

// Exec executes the request and unmarshals the response into WallEditCommentResponse
func (r *WallEditCommentRequest) Exec(ctx context.Context) (response response.WallEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetRequest defines the request for wall.get
//
// Retrieves a list of posts from a user's or community's wall.
// Doc: https://dev.vk.com/method/wall.get
type WallGetRequest struct {
	BaseRequest
}

// NewWallGetRequest creates a new request for wall.get
func NewWallGetRequest(a *api.API, actor actor.Actor) *WallGetRequest {
	return &WallGetRequest{*NewMethodBaseRequest(a, actor, "wall.get")}
}

// Exec executes the request and unmarshals the response into WallGetResponse
func (r *WallGetRequest) Exec(ctx context.Context) (response response.WallGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetExtendedRequest defines the request for wall.get
//
// Retrieves a list of posts from a user's or community's wall.
// Doc: https://dev.vk.com/method/wall.get
type WallGetExtendedRequest struct {
	BaseRequest
}

// NewWallGetExtendedRequest creates a new request for wall.get
func NewWallGetExtendedRequest(a *api.API, actor actor.Actor) *WallGetExtendedRequest {
	return &WallGetExtendedRequest{*NewMethodBaseRequest(a, actor, "wall.get")}
}

// Exec executes the request and unmarshals the response into WallGetExtendedResponse
func (r *WallGetExtendedRequest) Exec(ctx context.Context) (response response.WallGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetByIDRequest defines the request for wall.getById
//
// Retrieves one or more posts by their IDs.
// Doc: https://dev.vk.com/method/wall.getById
type WallGetByIDRequest struct {
	BaseRequest
}

// NewWallGetByIDRequest creates a new request for wall.getById
func NewWallGetByIDRequest(a *api.API, actor actor.Actor) *WallGetByIDRequest {
	return &WallGetByIDRequest{*NewMethodBaseRequest(a, actor, "wall.getById")}
}

// Exec executes the request and unmarshals the response into WallGetByIDResponse
func (r *WallGetByIDRequest) Exec(ctx context.Context) (response response.WallGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetByIDExtendedRequest defines the request for wall.getById
//
// Retrieves one or more posts by their IDs.
// Doc: https://dev.vk.com/method/wall.getById
type WallGetByIDExtendedRequest struct {
	BaseRequest
}

// NewWallGetByIDExtendedRequest creates a new request for wall.getById
func NewWallGetByIDExtendedRequest(a *api.API, actor actor.Actor) *WallGetByIDExtendedRequest {
	return &WallGetByIDExtendedRequest{*NewMethodBaseRequest(a, actor, "wall.getById")}
}

// Exec executes the request and unmarshals the response into WallGetByIDExtendedResponse
func (r *WallGetByIDExtendedRequest) Exec(ctx context.Context) (response response.WallGetByIDExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetCommentRequest defines the request for wall.getComment
//
// Retrieves information about a comment on a post.
// Doc: https://dev.vk.com/method/wall.getComment
type WallGetCommentRequest struct {
	BaseRequest
}

// NewWallGetCommentRequest creates a new request for wall.getComment
func NewWallGetCommentRequest(a *api.API, actor actor.Actor) *WallGetCommentRequest {
	return &WallGetCommentRequest{*NewMethodBaseRequest(a, actor, "wall.getComment")}
}

// Exec executes the request and unmarshals the response into WallGetCommentResponse
func (r *WallGetCommentRequest) Exec(ctx context.Context) (response response.WallGetCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetCommentExtendedRequest defines the request for wall.getComment
//
// Retrieves information about a comment on a post.
// Doc: https://dev.vk.com/method/wall.getComment
type WallGetCommentExtendedRequest struct {
	BaseRequest
}

// NewWallGetCommentExtendedRequest creates a new request for wall.getComment
func NewWallGetCommentExtendedRequest(a *api.API, actor actor.Actor) *WallGetCommentExtendedRequest {
	return &WallGetCommentExtendedRequest{*NewMethodBaseRequest(a, actor, "wall.getComment")}
}

// Exec executes the request and unmarshals the response into WallGetCommentExtendedResponse
func (r *WallGetCommentExtendedRequest) Exec(ctx context.Context) (response response.WallGetCommentExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetCommentsRequest defines the request for wall.getComments
//
// Retrieves comments on a post.
// Doc: https://dev.vk.com/method/wall.getComments
type WallGetCommentsRequest struct {
	BaseRequest
}

// NewWallGetCommentsRequest creates a new request for wall.getComments
func NewWallGetCommentsRequest(a *api.API, actor actor.Actor) *WallGetCommentsRequest {
	return &WallGetCommentsRequest{*NewMethodBaseRequest(a, actor, "wall.getComments")}
}

// Exec executes the request and unmarshals the response into WallGetCommentsResponse
func (r *WallGetCommentsRequest) Exec(ctx context.Context) (response response.WallGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetCommentsExtendedRequest defines the request for wall.getComments
//
// Retrieves comments on a post.
// Doc: https://dev.vk.com/method/wall.getComments
type WallGetCommentsExtendedRequest struct {
	BaseRequest
}

// NewWallGetCommentsExtendedRequest creates a new request for wall.getComments
func NewWallGetCommentsExtendedRequest(a *api.API, actor actor.Actor) *WallGetCommentsExtendedRequest {
	return &WallGetCommentsExtendedRequest{*NewMethodBaseRequest(a, actor, "wall.getComments")}
}

// Exec executes the request and unmarshals the response into WallGetCommentsExtendedResponse
func (r *WallGetCommentsExtendedRequest) Exec(ctx context.Context) (response response.WallGetCommentsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallGetRepostsRequest defines the request for wall.getReposts
//
// Retrieves information about reposts of a post.
// Doc: https://dev.vk.com/method/wall.getReposts
type WallGetRepostsRequest struct {
	BaseRequest
}

// NewWallGetRepostsRequest creates a new request for wall.getReposts
func NewWallGetRepostsRequest(a *api.API, actor actor.Actor) *WallGetRepostsRequest {
	return &WallGetRepostsRequest{*NewMethodBaseRequest(a, actor, "wall.getReposts")}
}

// Exec executes the request and unmarshals the response into WallGetRepostsResponse
func (r *WallGetRepostsRequest) Exec(ctx context.Context) (response response.WallGetRepostsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallOpenCommentsRequest defines the request for wall.openComments
//
// Opens comments for a post on the wall.
// Doc: https://dev.vk.com/method/wall.openComments
type WallOpenCommentsRequest struct {
	BaseRequest
}

// NewWallOpenCommentsRequest creates a new request for wall.openComments
func NewWallOpenCommentsRequest(a *api.API, actor actor.Actor) *WallOpenCommentsRequest {
	return &WallOpenCommentsRequest{*NewMethodBaseRequest(a, actor, "wall.openComments")}
}

// Exec executes the request and unmarshals the response into WallOpenCommentsResponse
func (r *WallOpenCommentsRequest) Exec(ctx context.Context) (response response.WallOpenCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallParseAttachedLinkRequest defines the request for wall.parseAttachedLink
//
// Retrieves additional information for a link that can be used to create a snippet when posting on the wall.
// Doc: https://dev.vk.com/method/wall.parseAttachedLink
type WallParseAttachedLinkRequest struct {
	BaseRequest
}

// NewWallParseAttachedLinkRequest creates a new request for wall.parseAttachedLink
func NewWallParseAttachedLinkRequest(a *api.API, actor actor.Actor) *WallParseAttachedLinkRequest {
	return &WallParseAttachedLinkRequest{*NewMethodBaseRequest(a, actor, "wall.parseAttachedLink")}
}

// Exec executes the request and unmarshals the response into WallParseAttachedLinkResponse
func (r *WallParseAttachedLinkRequest) Exec(ctx context.Context) (response response.WallParseAttachedLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallPinRequest defines the request for wall.pin
//
// Pins a post to the top of a user or community wall.
// Doc: https://dev.vk.com/method/wall.pin
type WallPinRequest struct {
	BaseRequest
}

// NewWallPinRequest creates a new request for wall.pin
func NewWallPinRequest(a *api.API, actor actor.Actor) *WallPinRequest {
	return &WallPinRequest{*NewMethodBaseRequest(a, actor, "wall.pin")}
}

// Exec executes the request and unmarshals the response into WallPinResponse
func (r *WallPinRequest) Exec(ctx context.Context) (response response.WallPinResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallPostRequest defines the request for wall.post
//
// Adds a new post on a user or community wall.
// Doc: https://dev.vk.com/method/wall.post
type WallPostRequest struct {
	BaseRequest
}

// NewWallPostRequest creates a new request for wall.post
func NewWallPostRequest(a *api.API, actor actor.Actor) *WallPostRequest {
	return &WallPostRequest{*NewMethodBaseRequest(a, actor, "wall.post")}
}

// Exec executes the request and unmarshals the response into WallPostResponse
func (r *WallPostRequest) Exec(ctx context.Context) (response response.WallPostResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallPostAdsStealthRequest defines the request for wall.postAdsStealth
//
// Creates a hidden (stealth) post on the wall.
// Doc: https://dev.vk.com/method/wall.postAdsStealth
type WallPostAdsStealthRequest struct {
	BaseRequest
}

// NewWallPostAdsStealthRequest creates a new request for wall.postAdsStealth
func NewWallPostAdsStealthRequest(a *api.API, actor actor.Actor) *WallPostAdsStealthRequest {
	return &WallPostAdsStealthRequest{*NewMethodBaseRequest(a, actor, "wall.postAdsStealth")}
}

// Exec executes the request and unmarshals the response into WallPostAdsStealthResponse
func (r *WallPostAdsStealthRequest) Exec(ctx context.Context) (response response.WallPostAdsStealthResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallReportCommentRequest defines the request for wall.reportComment
//
// Reports (flags) a comment as inappropriate.
// Doc: https://dev.vk.com/method/wall.reportComment
type WallReportCommentRequest struct {
	BaseRequest
}

// NewWallReportCommentRequest creates a new request for wall.reportComment
func NewWallReportCommentRequest(a *api.API, actor actor.Actor) *WallReportCommentRequest {
	return &WallReportCommentRequest{*NewMethodBaseRequest(a, actor, "wall.reportComment")}
}

// Exec executes the request and unmarshals the response into WallReportCommentResponse
func (r *WallReportCommentRequest) Exec(ctx context.Context) (response response.WallReportCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallReportPostRequest defines the request for wall.reportPost
//
// Reports (flags) a post as inappropriate.
// Doc: https://dev.vk.com/method/wall.reportPost
type WallReportPostRequest struct {
	BaseRequest
}

// NewWallReportPostRequest creates a new request for wall.reportPost
func NewWallReportPostRequest(a *api.API, actor actor.Actor) *WallReportPostRequest {
	return &WallReportPostRequest{*NewMethodBaseRequest(a, actor, "wall.reportPost")}
}

// Exec executes the request and unmarshals the response into WallReportPostResponse
func (r *WallReportPostRequest) Exec(ctx context.Context) (response response.WallReportPostResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallRepostRequest defines the request for wall.repost
//
// Shares (reposts) a post to a user's wall or community wall.
// Doc: https://dev.vk.com/method/wall.repost
type WallRepostRequest struct {
	BaseRequest
}

// NewWallRepostRequest creates a new request for wall.repost
func NewWallRepostRequest(a *api.API, actor actor.Actor) *WallRepostRequest {
	return &WallRepostRequest{*NewMethodBaseRequest(a, actor, "wall.repost")}
}

// Exec executes the request and unmarshals the response into WallRepostResponse
func (r *WallRepostRequest) Exec(ctx context.Context) (response response.WallRepostResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallRestoreRequest defines the request for wall.restore
//
// Restores a deleted post on a user or community wall.
// Doc: https://dev.vk.com/method/wall.restore
type WallRestoreRequest struct {
	BaseRequest
}

// NewWallRestoreRequest creates a new request for wall.restore
func NewWallRestoreRequest(a *api.API, actor actor.Actor) *WallRestoreRequest {
	return &WallRestoreRequest{*NewMethodBaseRequest(a, actor, "wall.restore")}
}

// Exec executes the request and unmarshals the response into WallRestoreResponse
func (r *WallRestoreRequest) Exec(ctx context.Context) (response response.WallRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallRestoreCommentRequest defines the request for wall.restoreComment
//
// Restores a deleted comment on a post.
// Doc: https://dev.vk.com/method/wall.restoreComment
type WallRestoreCommentRequest struct {
	BaseRequest
}

// NewWallRestoreCommentRequest creates a new request for wall.restoreComment
func NewWallRestoreCommentRequest(a *api.API, actor actor.Actor) *WallRestoreCommentRequest {
	return &WallRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "wall.restoreComment")}
}

// Exec executes the request and unmarshals the response into WallRestoreCommentResponse
func (r *WallRestoreCommentRequest) Exec(ctx context.Context) (response response.WallRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallSearchRequest defines the request for wall.search
//
// Searches for posts on a wall.
// Doc: https://dev.vk.com/method/wall.search
type WallSearchRequest struct {
	BaseRequest
}

// NewWallSearchRequest creates a new request for wall.search
func NewWallSearchRequest(a *api.API, actor actor.Actor) *WallSearchRequest {
	return &WallSearchRequest{*NewMethodBaseRequest(a, actor, "wall.search")}
}

// Exec executes the request and unmarshals the response into WallSearchResponse
func (r *WallSearchRequest) Exec(ctx context.Context) (response response.WallSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallSearchExtendedRequest defines the request for wall.search
//
// Searches for posts on a wall.
// Doc: https://dev.vk.com/method/wall.search
type WallSearchExtendedRequest struct {
	BaseRequest
}

// NewWallSearchExtendedRequest creates a new request for wall.search
func NewWallSearchExtendedRequest(a *api.API, actor actor.Actor) *WallSearchExtendedRequest {
	return &WallSearchExtendedRequest{*NewMethodBaseRequest(a, actor, "wall.search")}
}

// Exec executes the request and unmarshals the response into WallSearchExtendedResponse
func (r *WallSearchExtendedRequest) Exec(ctx context.Context) (response response.WallSearchExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// WallUnpinRequest defines the request for wall.unpin
//
// Unpins a post from the top of a user or community wall.
// Doc: https://dev.vk.com/method/wall.unpin
type WallUnpinRequest struct {
	BaseRequest
}

// NewWallUnpinRequest creates a new request for wall.unpin
func NewWallUnpinRequest(a *api.API, actor actor.Actor) *WallUnpinRequest {
	return &WallUnpinRequest{*NewMethodBaseRequest(a, actor, "wall.unpin")}
}

// Exec executes the request and unmarshals the response into WallUnpinResponse
func (r *WallUnpinRequest) Exec(ctx context.Context) (response response.WallUnpinResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
