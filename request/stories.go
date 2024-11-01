package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/stories

// StoriesBanOwnerRequest defines the request for stories.banOwner
//
// Hides stories from selected sources in the news feed.
// Doc: https://dev.vk.com/method/stories.banOwner
type StoriesBanOwnerRequest struct {
	BaseRequest
}

// NewStoriesBanOwnerRequest creates a new request for stories.banOwner
func NewStoriesBanOwnerRequest(a *api.API, actor actor.Actor) *StoriesBanOwnerRequest {
	return &StoriesBanOwnerRequest{*NewMethodBaseRequest(a, actor, "stories.banOwner")}
}

// Exec executes the request and unmarshals the response into StoriesBanOwnerResponse
func (r *StoriesBanOwnerRequest) Exec(ctx context.Context) (response response.StoriesBanOwnerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesDeleteRequest defines the request for stories.delete
//
// Deletes a story.
// Doc: https://dev.vk.com/method/stories.delete
type StoriesDeleteRequest struct {
	BaseRequest
}

// NewStoriesDeleteRequest creates a new request for stories.delete
func NewStoriesDeleteRequest(a *api.API, actor actor.Actor) *StoriesDeleteRequest {
	return &StoriesDeleteRequest{*NewMethodBaseRequest(a, actor, "stories.delete")}
}

// Exec executes the request and unmarshals the response into StoriesDeleteResponse
func (r *StoriesDeleteRequest) Exec(ctx context.Context) (response response.StoriesDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetRequest defines the request for stories.get
//
// Returns stories available for the current user.
// Doc: https://dev.vk.com/method/stories.get
type StoriesGetRequest struct {
	BaseRequest
}

// NewStoriesGetRequest creates a new request for stories.get
func NewStoriesGetRequest(a *api.API, actor actor.Actor) *StoriesGetRequest {
	return &StoriesGetRequest{*NewMethodBaseRequest(a, actor, "stories.get")}
}

// Exec executes the request and unmarshals the response into StoriesGetResponse
func (r *StoriesGetRequest) Exec(ctx context.Context) (response response.StoriesGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetExtendedRequest defines the request for stories.get
//
// Returns stories available for the current user.
// Doc: https://dev.vk.com/method/stories.get
type StoriesGetExtendedRequest struct {
	BaseRequest
}

// NewStoriesGetExtendedRequest creates a new request for stories.get
func NewStoriesGetExtendedRequest(a *api.API, actor actor.Actor) *StoriesGetExtendedRequest {
	r := &StoriesGetExtendedRequest{*NewMethodBaseRequest(a, actor, "stories.get")}
	_ = r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into StoriesGetExtendedResponse
func (r *StoriesGetExtendedRequest) Exec(ctx context.Context) (response response.StoriesGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetBannedRequest defines the request for stories.getBanned
//
// Returns a list of story sources hidden from the current user's feed.
// Doc: https://dev.vk.com/method/stories.getBanned
type StoriesGetBannedRequest struct {
	BaseRequest
}

// NewStoriesGetBannedRequest creates a new request for stories.getBanned
func NewStoriesGetBannedRequest(a *api.API, actor actor.Actor) *StoriesGetBannedRequest {
	return &StoriesGetBannedRequest{*NewMethodBaseRequest(a, actor, "stories.getBanned")}
}

// Exec executes the request and unmarshals the response into StoriesGetBannedResponse
func (r *StoriesGetBannedRequest) Exec(ctx context.Context) (response response.StoriesGetBannedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetBannedExtendedRequest defines the request for stories.getBanned
//
// Returns a list of story sources hidden from the current user's feed.
// Doc: https://dev.vk.com/method/stories.getBanned
type StoriesGetBannedExtendedRequest struct {
	BaseRequest
}

// NewStoriesGetBannedExtendedRequest creates a new request for stories.getBanned
func NewStoriesGetBannedExtendedRequest(a *api.API, actor actor.Actor) *StoriesGetBannedExtendedRequest {
	r := &StoriesGetBannedExtendedRequest{*NewMethodBaseRequest(a, actor, "stories.getBanned")}
	_ = r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into StoriesGetBannedExtendedResponse
func (r *StoriesGetBannedExtendedRequest) Exec(ctx context.Context) (response response.StoriesGetBannedExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetByIDRequest defines the request for stories.getById
//
// Returns information about a story by its ID.
// Doc: https://dev.vk.com/method/stories.getById
type StoriesGetByIDRequest struct {
	BaseRequest
}

// NewStoriesGetByIDRequest creates a new request for stories.getById
func NewStoriesGetByIDRequest(a *api.API, actor actor.Actor) *StoriesGetByIDRequest {
	return &StoriesGetByIDRequest{*NewMethodBaseRequest(a, actor, "stories.getById")}
}

// Exec executes the request and unmarshals the response into StoriesGetByIDResponse
func (r *StoriesGetByIDRequest) Exec(ctx context.Context) (response response.StoriesGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetByIDExtendedRequest defines the request for stories.getById
//
// Returns information about a story by its ID.
// Doc: https://dev.vk.com/method/stories.getById
type StoriesGetByIDExtendedRequest struct {
	BaseRequest
}

// NewStoriesGetByIDExtendedRequest creates a new request for stories.getById
func NewStoriesGetByIDExtendedRequest(a *api.API, actor actor.Actor) *StoriesGetByIDExtendedRequest {
	r := &StoriesGetByIDExtendedRequest{*NewMethodBaseRequest(a, actor, "stories.getById")}
	_ = r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into StoriesGetByIDExtendedResponse
func (r *StoriesGetByIDExtendedRequest) Exec(ctx context.Context) (response response.StoriesGetByIDExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetDetailedStatsRequest defines the request for stories.getDetailedStats
//
// Retrieves detailed statistics for a story.
// Doc: https://dev.vk.com/method/stories.getDetailedStats
type StoriesGetDetailedStatsRequest struct {
	BaseRequest
}

// NewStoriesGetDetailedStatsRequest creates a new request for stories.getDetailedStats
func NewStoriesGetDetailedStatsRequest(a *api.API, actor actor.Actor) *StoriesGetDetailedStatsRequest {
	return &StoriesGetDetailedStatsRequest{*NewMethodBaseRequest(a, actor, "stories.getDetailedStats")}
}

// Exec executes the request and unmarshals the response into StoriesGetDetailedStatsResponse
func (r *StoriesGetDetailedStatsRequest) Exec(ctx context.Context) (response response.StoriesGetDetailedStatsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetPhotoUploadServerRequest defines the request for stories.getPhotoUploadServer
//
// Gets the URL for uploading a photo to a story.
// Doc: https://dev.vk.com/method/stories.getPhotoUploadServer
type StoriesGetPhotoUploadServerRequest struct {
	BaseRequest
}

// NewStoriesGetPhotoUploadServerRequest creates a new request for stories.getPhotoUploadServer
func NewStoriesGetPhotoUploadServerRequest(a *api.API, actor actor.Actor) *StoriesGetPhotoUploadServerRequest {
	return &StoriesGetPhotoUploadServerRequest{*NewMethodBaseRequest(a, actor, "stories.getPhotoUploadServer")}
}

// Exec executes the request and unmarshals the response into StoriesGetPhotoUploadServerResponse
func (r *StoriesGetPhotoUploadServerRequest) Exec(ctx context.Context) (response response.StoriesGetPhotoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetRepliesRequest defines the request for stories.getReplies
//
// Retrieves replies to a story.
// Doc: https://dev.vk.com/method/stories.getReplies
type StoriesGetRepliesRequest struct {
	BaseRequest
}

// NewStoriesGetRepliesRequest creates a new request for stories.getReplies
func NewStoriesGetRepliesRequest(a *api.API, actor actor.Actor) *StoriesGetRepliesRequest {
	return &StoriesGetRepliesRequest{*NewMethodBaseRequest(a, actor, "stories.getReplies")}
}

// Exec executes the request and unmarshals the response into StoriesGetRepliesResponse
func (r *StoriesGetRepliesRequest) Exec(ctx context.Context) (response response.StoriesGetRepliesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetRepliesExtendedRequest defines the request for stories.getReplies
//
// Retrieves replies to a story.
// Doc: https://dev.vk.com/method/stories.getReplies
type StoriesGetRepliesExtendedRequest struct {
	BaseRequest
}

// NewStoriesGetRepliesExtendedRequest creates a new request for stories.getReplies
func NewStoriesGetRepliesExtendedRequest(a *api.API, actor actor.Actor) *StoriesGetRepliesExtendedRequest {
	r := &StoriesGetRepliesExtendedRequest{*NewMethodBaseRequest(a, actor, "stories.getReplies")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into StoriesGetRepliesExtendedResponse
func (r *StoriesGetRepliesExtendedRequest) Exec(ctx context.Context) (response response.StoriesGetRepliesExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetStatsRequest defines the request for stories.getStats
//
// Returns story statistics.
// Doc: https://dev.vk.com/method/stories.getStats
type StoriesGetStatsRequest struct {
	BaseRequest
}

// NewStoriesGetStatsRequest creates a new request for stories.getStats
func NewStoriesGetStatsRequest(a *api.API, actor actor.Actor) *StoriesGetStatsRequest {
	return &StoriesGetStatsRequest{*NewMethodBaseRequest(a, actor, "stories.getStats")}
}

// Exec executes the request and unmarshals the response into StoriesGetStatsResponse
func (r *StoriesGetStatsRequest) Exec(ctx context.Context) (response response.StoriesGetStatsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetVideoUploadServerRequest defines the request for stories.getVideoUploadServer
//
// Gets the URL for uploading a video to a story.
// Doc: https://dev.vk.com/method/stories.getVideoUploadServer
type StoriesGetVideoUploadServerRequest struct {
	BaseRequest
}

// NewStoriesGetVideoUploadServerRequest creates a new request for stories.getVideoUploadServer
func NewStoriesGetVideoUploadServerRequest(a *api.API, actor actor.Actor) *StoriesGetVideoUploadServerRequest {
	return &StoriesGetVideoUploadServerRequest{*NewMethodBaseRequest(a, actor, "stories.getVideoUploadServer")}
}

// Exec executes the request and unmarshals the response into StoriesGetVideoUploadServerResponse
func (r *StoriesGetVideoUploadServerRequest) Exec(ctx context.Context) (response response.StoriesGetVideoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesGetViewersRequest defines the request for stories.getViewers
//
// Returns a list of users who viewed the story.
// Doc: https://dev.vk.com/method/stories.getViewers
type StoriesGetViewersRequest struct {
	BaseRequest
}

// NewStoriesGetViewersRequest creates a new request for stories.getViewers
func NewStoriesGetViewersRequest(a *api.API, actor actor.Actor) *StoriesGetViewersRequest {
	return &StoriesGetViewersRequest{*NewMethodBaseRequest(a, actor, "stories.getViewers")}
}

// Exec executes the request and unmarshals the response into StoriesGetViewersResponse
func (r *StoriesGetViewersRequest) Exec(ctx context.Context) (response response.StoriesGetViewersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesHideAllRepliesRequest defines the request for stories.hideAllReplies
//
// Hides all replies from the author for the last 24 hours.
// Doc: https://dev.vk.com/method/stories.hideAllReplies
type StoriesHideAllRepliesRequest struct {
	BaseRequest
}

// NewStoriesHideAllRepliesRequest creates a new request for stories.hideAllReplies
func NewStoriesHideAllRepliesRequest(a *api.API, actor actor.Actor) *StoriesHideAllRepliesRequest {
	return &StoriesHideAllRepliesRequest{*NewMethodBaseRequest(a, actor, "stories.hideAllReplies")}
}

// Exec executes the request and unmarshals the response into StoriesHideAllRepliesResponse
func (r *StoriesHideAllRepliesRequest) Exec(ctx context.Context) (response response.StoriesHideAllRepliesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesHideReplyRequest defines the request for stories.hideReply
//
// Hides a reply to the story.
// Doc: https://dev.vk.com/method/stories.hideReply
type StoriesHideReplyRequest struct {
	BaseRequest
}

// NewStoriesHideReplyRequest creates a new request for stories.hideReply
func NewStoriesHideReplyRequest(a *api.API, actor actor.Actor) *StoriesHideReplyRequest {
	return &StoriesHideReplyRequest{*NewMethodBaseRequest(a, actor, "stories.hideReply")}
}

// Exec executes the request and unmarshals the response into StoriesHideReplyResponse
func (r *StoriesHideReplyRequest) Exec(ctx context.Context) (response response.StoriesHideReplyResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesSaveRequest defines the request for stories.save
//
// Saves a story in the profile after it has been uploaded successfully.
// Doc: https://dev.vk.com/method/stories.save
type StoriesSaveRequest struct {
	BaseRequest
}

// NewStoriesSaveRequest creates a new request for stories.save
func NewStoriesSaveRequest(a *api.API, actor actor.Actor) *StoriesSaveRequest {
	return &StoriesSaveRequest{*NewMethodBaseRequest(a, actor, "stories.save")}
}

// Exec executes the request and unmarshals the response into StoriesSaveResponse
func (r *StoriesSaveRequest) Exec(ctx context.Context) (response response.StoriesSaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesSearchRequest defines the request for stories.search
//
// Returns search results for stories.
// Doc: https://dev.vk.com/method/stories.search
type StoriesSearchRequest struct {
	BaseRequest
}

// NewStoriesSearchRequest creates a new request for stories.search
func NewStoriesSearchRequest(a *api.API, actor actor.Actor) *StoriesSearchRequest {
	return &StoriesSearchRequest{*NewMethodBaseRequest(a, actor, "stories.search")}
}

// Exec executes the request and unmarshals the response into StoriesSearchResponse
func (r *StoriesSearchRequest) Exec(ctx context.Context) (response response.StoriesSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesSearchExtendedRequest defines the request for stories.search
//
// Returns search results for stories.
// Doc: https://dev.vk.com/method/stories.search
type StoriesSearchExtendedRequest struct {
	BaseRequest
}

// NewStoriesSearchExtendedRequest creates a new request for stories.search
func NewStoriesSearchExtendedRequest(a *api.API, actor actor.Actor) *StoriesSearchExtendedRequest {
	r := &StoriesSearchExtendedRequest{*NewMethodBaseRequest(a, actor, "stories.search")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into StoriesSearchExtendedResponse
func (r *StoriesSearchExtendedRequest) Exec(ctx context.Context) (response response.StoriesSearchExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesSendInteractionRequest defines the request for stories.sendInteraction
//
// Sends feedback for a story.
// Doc: https://dev.vk.com/method/stories.sendInteraction
type StoriesSendInteractionRequest struct {
	BaseRequest
}

// NewStoriesSendInteractionRequest creates a new request for stories.sendInteraction
func NewStoriesSendInteractionRequest(a *api.API, actor actor.Actor) *StoriesSendInteractionRequest {
	return &StoriesSendInteractionRequest{*NewMethodBaseRequest(a, actor, "stories.sendInteraction")}
}

// Exec executes the request and unmarshals the response into StoriesSendInteractionResponse
func (r *StoriesSendInteractionRequest) Exec(ctx context.Context) (response response.StoriesSendInteractionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoriesUnbanOwnerRequest defines the request for stories.unbanOwner
//
// Unbans a user or community from appearing in the stories feed.
// Doc: https://dev.vk.com/method/stories.unbanOwner
type StoriesUnbanOwnerRequest struct {
	BaseRequest
}

// NewStoriesUnbanOwnerRequest creates a new request for stories.unbanOwner
func NewStoriesUnbanOwnerRequest(a *api.API, actor actor.Actor) *StoriesUnbanOwnerRequest {
	return &StoriesUnbanOwnerRequest{*NewMethodBaseRequest(a, actor, "stories.unbanOwner")}
}

// Exec executes the request and unmarshals the response into StoriesUnbanOwnerResponse
func (r *StoriesUnbanOwnerRequest) Exec(ctx context.Context) (response response.StoriesUnbanOwnerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
