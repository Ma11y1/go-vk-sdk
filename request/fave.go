package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/fave

// FaveAddArticleRequest defines the request for fave.addArticle
//
// The method adds an article to bookmarks.
// Doc: https://dev.vk.com/method/fave.addArticle
type FaveAddArticleRequest struct {
	BaseRequest
}

// NewFaveAddArticleRequest creates a new request for fave.addArticle
func NewFaveAddArticleRequest(a *api.API, actor actor.Actor) *FaveAddArticleRequest {
	return &FaveAddArticleRequest{*NewMethodBaseRequest(a, actor, "fave.addArticle")}
}

// Exec executes the request and unmarshals the response into FaveAddArticleResponse
func (r *FaveAddArticleRequest) Exec(ctx context.Context) (response response.FaveAddArticleResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddLinkRequest defines the request for fave.addLink
//
// The method adds a link to bookmarks.
// Doc: https://dev.vk.com/method/fave.addLink
type FaveAddLinkRequest struct {
	BaseRequest
}

// NewFaveAddLinkRequest creates a new request for fave.addLink
func NewFaveAddLinkRequest(a *api.API, actor actor.Actor) *FaveAddLinkRequest {
	return &FaveAddLinkRequest{*NewMethodBaseRequest(a, actor, "fave.addLink")}
}

// Exec executes the request and unmarshals the response into FaveAddLinkResponse
func (r *FaveAddLinkRequest) Exec(ctx context.Context) (response response.FaveAddLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddPageRequest defines the request for fave.addPage
//
// The method adds a community or user to bookmarks.
// Doc: https://dev.vk.com/method/fave.addPage
type FaveAddPageRequest struct {
	BaseRequest
}

// NewFaveAddPageRequest creates a new request for fave.addPage
func NewFaveAddPageRequest(a *api.API, actor actor.Actor) *FaveAddPageRequest {
	return &FaveAddPageRequest{*NewMethodBaseRequest(a, actor, "fave.addPage")}
}

// Exec executes the request and unmarshals the response into FaveAddPageResponse
func (r *FaveAddPageRequest) Exec(ctx context.Context) (response response.FaveAddPageResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddPostRequest defines the request for fave.addPost
//
// The method adds a wall post from a user or community to bookmarks.
// Doc: https://dev.vk.com/method/fave.addPost
type FaveAddPostRequest struct {
	BaseRequest
}

// NewFaveAddPostRequest creates a new request for fave.addPost
func NewFaveAddPostRequest(a *api.API, actor actor.Actor) *FaveAddPostRequest {
	return &FaveAddPostRequest{*NewMethodBaseRequest(a, actor, "fave.addPost")}
}

// Exec executes the request and unmarshals the response into FaveAddPostResponse
func (r *FaveAddPostRequest) Exec(ctx context.Context) (response response.FaveAddPostResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddProductRequest defines the request for fave.addProduct
//
// The method adds a product to bookmarks.
// Doc: https://dev.vk.com/method/fave.addProduct
type FaveAddProductRequest struct {
	BaseRequest
}

// NewFaveAddProductRequest creates a new request for fave.addProduct
func NewFaveAddProductRequest(a *api.API, actor actor.Actor) *FaveAddProductRequest {
	return &FaveAddProductRequest{*NewMethodBaseRequest(a, actor, "fave.addProduct")}
}

// Exec executes the request and unmarshals the response into FaveAddProductResponse
func (r *FaveAddProductRequest) Exec(ctx context.Context) (response response.FaveAddProductResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddTagRequest defines the request for fave.addTag
//
// The method creates a bookmark tag.
// Doc: https://dev.vk.com/method/fave.addTag
type FaveAddTagRequest struct {
	BaseRequest
}

// NewFaveAddTagRequest creates a new request for fave.addTag
func NewFaveAddTagRequest(a *api.API, actor actor.Actor) *FaveAddTagRequest {
	return &FaveAddTagRequest{*NewMethodBaseRequest(a, actor, "fave.addTag")}
}

// Exec executes the request and unmarshals the response into FaveAddTagResponse
func (r *FaveAddTagRequest) Exec(ctx context.Context) (response response.FaveAddTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveAddVideoRequest defines the request for fave.addVideo
//
// The method adds a video to bookmarks.
// Doc: https://dev.vk.com/method/fave.addVideo
type FaveAddVideoRequest struct {
	BaseRequest
}

// NewFaveAddVideoRequest creates a new request for fave.addVideo
func NewFaveAddVideoRequest(a *api.API, actor actor.Actor) *FaveAddVideoRequest {
	return &FaveAddVideoRequest{*NewMethodBaseRequest(a, actor, "fave.addVideo")}
}

// Exec executes the request and unmarshals the response into FaveAddVideoResponse
func (r *FaveAddVideoRequest) Exec(ctx context.Context) (response response.FaveAddVideoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveEditTagRequest defines the request for fave.editTag
//
// The method edits a bookmark tag.
// Doc: https://dev.vk.com/method/fave.editTag
type FaveEditTagRequest struct {
	BaseRequest
}

// NewFaveEditTagRequest creates a new request for fave.editTag
func NewFaveEditTagRequest(a *api.API, actor actor.Actor) *FaveEditTagRequest {
	return &FaveEditTagRequest{*NewMethodBaseRequest(a, actor, "fave.editTag")}
}

// Exec executes the request and unmarshals the response into FaveEditTagResponse
func (r *FaveEditTagRequest) Exec(ctx context.Context) (response response.FaveEditTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveGetRequest defines the request for fave.get
//
// The method returns objects added to the user's bookmarks.
// Doc: https://dev.vk.com/method/fave.get
type FaveGetRequest struct {
	BaseRequest
}

// NewFaveGetRequest creates a new request for fave.get
func NewFaveGetRequest(a *api.API, actor actor.Actor) *FaveGetRequest {
	return &FaveGetRequest{*NewMethodBaseRequest(a, actor, "fave.get")}
}

// Exec executes the request and unmarshals the response into FaveGetResponse
func (r *FaveGetRequest) Exec(ctx context.Context) (response response.FaveGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveGetExtendedRequest defines the request for fave.get
//
// The method returns objects added to the user's bookmarks.
// Doc: https://dev.vk.com/method/fave.get
type FaveGetExtendedRequest struct {
	BaseRequest
}

// NewFaveGetExtendedRequest creates a new request for fave.get
func NewFaveGetExtendedRequest(a *api.API, actor actor.Actor) *FaveGetExtendedRequest {
	r := &FaveGetExtendedRequest{*NewMethodBaseRequest(a, actor, "fave.get")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into FaveGetExtendedResponse
func (r *FaveGetExtendedRequest) Exec(ctx context.Context) (response response.FaveGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveGetPagesRequest defines the request for fave.getPages
//
// The method returns pages of users and communities added to bookmarks.
// Doc: https://dev.vk.com/method/fave.getPages
type FaveGetPagesRequest struct {
	BaseRequest
}

// NewFaveGetPagesRequest creates a new request for fave.getPages
func NewFaveGetPagesRequest(a *api.API, actor actor.Actor) *FaveGetPagesRequest {
	return &FaveGetPagesRequest{*NewMethodBaseRequest(a, actor, "fave.getPages")}
}

// Exec executes the request and unmarshals the response into FaveGetPagesResponse
func (r *FaveGetPagesRequest) Exec(ctx context.Context) (response response.FaveGetPagesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveGetTagsRequest defines the request for fave.getTags
//
// The method returns a list of bookmark tags.
// Doc: https://dev.vk.com/method/fave.getTags
type FaveGetTagsRequest struct {
	BaseRequest
}

// NewFaveGetTagsRequest creates a new request for fave.getTags
func NewFaveGetTagsRequest(a *api.API, actor actor.Actor) *FaveGetTagsRequest {
	return &FaveGetTagsRequest{*NewMethodBaseRequest(a, actor, "fave.getTags")}
}

// Exec executes the request and unmarshals the response into FaveGetTagsResponse
func (r *FaveGetTagsRequest) Exec(ctx context.Context) (response response.FaveGetTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveMarkSeenRequest defines the request for fave.markSeen
//
// The method marks bookmarks as viewed.
// Doc: https://dev.vk.com/method/fave.markSeen
type FaveMarkSeenRequest struct {
	BaseRequest
}

// NewFaveMarkSeenRequest creates a new request for fave.markSeen
func NewFaveMarkSeenRequest(a *api.API, actor actor.Actor) *FaveMarkSeenRequest {
	return &FaveMarkSeenRequest{*NewMethodBaseRequest(a, actor, "fave.markSeen")}
}

// Exec executes the request and unmarshals the response into FaveMarkSeenResponse
func (r *FaveMarkSeenRequest) Exec(ctx context.Context) (response response.FaveMarkSeenResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemoveArticleRequest defines the request for fave.removeArticle
//
// The method removes an article from bookmarks.
// Doc: https://dev.vk.com/method/fave.removeArticle
type FaveRemoveArticleRequest struct {
	BaseRequest
}

// NewFaveRemoveArticleRequest creates a new request for fave.removeArticle
func NewFaveRemoveArticleRequest(a *api.API, actor actor.Actor) *FaveRemoveArticleRequest {
	return &FaveRemoveArticleRequest{*NewMethodBaseRequest(a, actor, "fave.removeArticle")}
}

// Exec executes the request and unmarshals the response into FaveRemoveArticleResponse
func (r *FaveRemoveArticleRequest) Exec(ctx context.Context) (response response.FaveRemoveArticleResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemoveLinkRequest defines the request for fave.removeLink
//
// The method removes a link from the user's bookmarks.
// Doc: https://dev.vk.com/method/fave.removeLink
type FaveRemoveLinkRequest struct {
	BaseRequest
}

// NewFaveRemoveLinkRequest creates a new request for fave.removeLink
func NewFaveRemoveLinkRequest(a *api.API, actor actor.Actor) *FaveRemoveLinkRequest {
	return &FaveRemoveLinkRequest{*NewMethodBaseRequest(a, actor, "fave.removeLink")}
}

// Exec executes the request and unmarshals the response into FaveRemoveLinkResponse
func (r *FaveRemoveLinkRequest) Exec(ctx context.Context) (response response.FaveRemoveLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemovePageRequest defines the request for fave.removePage
//
// The method removes a community or user from bookmarks.
// Doc: https://dev.vk.com/method/fave.removePage
type FaveRemovePageRequest struct {
	BaseRequest
}

// NewFaveRemovePageRequest creates a new request for fave.removePage
func NewFaveRemovePageRequest(a *api.API, actor actor.Actor) *FaveRemovePageRequest {
	return &FaveRemovePageRequest{*NewMethodBaseRequest(a, actor, "fave.removePage")}
}

// Exec executes the request and unmarshals the response into FaveRemovePageResponse
func (r *FaveRemovePageRequest) Exec(ctx context.Context) (response response.FaveRemovePageResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemovePostRequest defines the request for fave.removePost
//
// The method removes a wall post from a user or community from bookmarks.
// Doc: https://dev.vk.com/method/fave.removePost
type FaveRemovePostRequest struct {
	BaseRequest
}

// NewFaveRemovePostRequest creates a new request for fave.removePost
func NewFaveRemovePostRequest(a *api.API, actor actor.Actor) *FaveRemovePostRequest {
	return &FaveRemovePostRequest{*NewMethodBaseRequest(a, actor, "fave.removePost")}
}

// Exec executes the request and unmarshals the response into FaveRemovePostResponse
func (r *FaveRemovePostRequest) Exec(ctx context.Context) (response response.FaveRemovePostResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemoveProductRequest defines the request for fave.removeProduct
//
// The method removes a product from bookmarks.
// Doc: https://dev.vk.com/method/fave.removeProduct
type FaveRemoveProductRequest struct {
	BaseRequest
}

// NewFaveRemoveProductRequest creates a new request for fave.removeProduct
func NewFaveRemoveProductRequest(a *api.API, actor actor.Actor) *FaveRemoveProductRequest {
	return &FaveRemoveProductRequest{*NewMethodBaseRequest(a, actor, "fave.removeProduct")}
}

// Exec executes the request and unmarshals the response into FaveRemoveProductResponse
func (r *FaveRemoveProductRequest) Exec(ctx context.Context) (response response.FaveRemoveProductResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemoveTagRequest defines the request for fave.removeTag
//
// The method removes a bookmark tag.
// Doc: https://dev.vk.com/method/fave.removeTag
type FaveRemoveTagRequest struct {
	BaseRequest
}

// NewFaveRemoveTagRequest creates a new request for fave.removeTag
func NewFaveRemoveTagRequest(a *api.API, actor actor.Actor) *FaveRemoveTagRequest {
	return &FaveRemoveTagRequest{*NewMethodBaseRequest(a, actor, "fave.removeTag")}
}

// Exec executes the request and unmarshals the response into FaveRemoveTagResponse
func (r *FaveRemoveTagRequest) Exec(ctx context.Context) (response response.FaveRemoveTagResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveRemoveVideoRequest defines the request for fave.removeVideo
//
// The method removes a video from bookmarks.
// Doc: https://dev.vk.com/method/fave.removeVideo
type FaveRemoveVideoRequest struct {
	BaseRequest
}

// NewFaveRemoveVideoRequest creates a new request for fave.removeVideo
func NewFaveRemoveVideoRequest(a *api.API, actor actor.Actor) *FaveRemoveVideoRequest {
	return &FaveRemoveVideoRequest{*NewMethodBaseRequest(a, actor, "fave.removeVideo")}
}

// Exec executes the request and unmarshals the response into FaveRemoveVideoResponse
func (r *FaveRemoveVideoRequest) Exec(ctx context.Context) (response response.FaveRemoveVideoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveReorderTagsRequest defines the request for fave.reorderTags
//
// The method changes the order of bookmark tags in the list.
// Doc: https://dev.vk.com/method/fave.reorderTags
type FaveReorderTagsRequest struct {
	BaseRequest
}

// NewFaveReorderTagsRequest creates a new request for fave.reorderTags
func NewFaveReorderTagsRequest(a *api.API, actor actor.Actor) *FaveReorderTagsRequest {
	return &FaveReorderTagsRequest{*NewMethodBaseRequest(a, actor, "fave.reorderTags")}
}

// Exec executes the request and unmarshals the response into FaveReorderTagsResponse
func (r *FaveReorderTagsRequest) Exec(ctx context.Context) (response response.FaveReorderTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveSetPageTagsRequest defines the request for fave.setPageTags
//
// The method sets a tag for a user or community page in bookmarks.
// Doc: https://dev.vk.com/method/fave.setPageTags
type FaveSetPageTagsRequest struct {
	BaseRequest
}

// NewFaveSetPageTagsRequest creates a new request for fave.setPageTags
func NewFaveSetPageTagsRequest(a *api.API, actor actor.Actor) *FaveSetPageTagsRequest {
	return &FaveSetPageTagsRequest{*NewMethodBaseRequest(a, actor, "fave.setPageTags")}
}

// Exec executes the request and unmarshals the response into FaveSetPageTagsResponse
func (r *FaveSetPageTagsRequest) Exec(ctx context.Context) (response response.FaveSetPageTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveSetTagsRequest defines the request for fave.setTags
//
// The method sets a tag for an objects in the bookmarks list.
// Doc: https://dev.vk.com/method/fave.setTags
type FaveSetTagsRequest struct {
	BaseRequest
}

// NewFaveSetTagsRequest creates a new request for fave.setTags
func NewFaveSetTagsRequest(a *api.API, actor actor.Actor) *FaveSetTagsRequest {
	return &FaveSetTagsRequest{*NewMethodBaseRequest(a, actor, "fave.setTags")}
}

// Exec executes the request and unmarshals the response into FaveSetTagsResponse
func (r *FaveSetTagsRequest) Exec(ctx context.Context) (response response.FaveSetTagsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FaveTrackPageInteractionRequest defines the request for fave.trackPageInteraction
//
// The method moves a user or community page to the top of the bookmarks list.
// Doc: https://dev.vk.com/method/fave.trackPageInteraction
type FaveTrackPageInteractionRequest struct {
	BaseRequest
}

// NewFaveTrackPageInteractionRequest creates a new request for fave.trackPageInteraction
func NewFaveTrackPageInteractionRequest(a *api.API, actor actor.Actor) *FaveTrackPageInteractionRequest {
	return &FaveTrackPageInteractionRequest{*NewMethodBaseRequest(a, actor, "fave.trackPageInteraction")}
}

// Exec executes the request and unmarshals the response into FaveTrackPageInteractionResponse
func (r *FaveTrackPageInteractionRequest) Exec(ctx context.Context) (response response.FaveTrackPageInteractionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
