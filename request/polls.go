package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/polls

// PollsAddVoteRequest defines the request for polls.addVote
//
// Allows the current user to vote for a selected answer in the specified poll.
// Doc: https://dev.vk.com/method/polls.addVote
type PollsAddVoteRequest struct {
	BaseRequest
}

// NewPollsAddVoteRequest creates a new request for polls.addVote
func NewPollsAddVoteRequest(a *api.API, actor actor.Actor) *PollsAddVoteRequest {
	return &PollsAddVoteRequest{*NewMethodBaseRequest(a, actor, "polls.addVote")}
}

// Exec executes the request and unmarshals the response into PollsAddVoteResponse
func (r *PollsAddVoteRequest) Exec(ctx context.Context) (response response.PollsAddVoteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsCreateRequest defines the request for polls.create
//
// Allows creating polls that can be attached to posts on a user or community page.
// Doc: https://dev.vk.com/method/polls.create
type PollsCreateRequest struct {
	BaseRequest
}

// NewPollsCreateRequest creates a new request for polls.create
func NewPollsCreateRequest(a *api.API, actor actor.Actor) *PollsCreateRequest {
	return &PollsCreateRequest{*NewMethodBaseRequest(a, actor, "polls.create")}
}

// Exec executes the request and unmarshals the response into PollsCreateResponse
func (r *PollsCreateRequest) Exec(ctx context.Context) (response response.PollsCreateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsDeleteVoteRequest defines the request for polls.deleteVote
//
// Removes the current user's vote from the selected answer in the specified poll.
// Doc: https://dev.vk.com/method/polls.deleteVote
type PollsDeleteVoteRequest struct {
	BaseRequest
}

// NewPollsDeleteVoteRequest creates a new request for polls.deleteVote
func NewPollsDeleteVoteRequest(a *api.API, actor actor.Actor) *PollsDeleteVoteRequest {
	return &PollsDeleteVoteRequest{*NewMethodBaseRequest(a, actor, "polls.deleteVote")}
}

// Exec executes the request and unmarshals the response into PollsDeleteVoteResponse
func (r *PollsDeleteVoteRequest) Exec(ctx context.Context) (response response.PollsDeleteVoteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsEditRequest defines the request for polls.edit
//
// Allows editing created polls.
// Doc: https://dev.vk.com/method/polls.edit
type PollsEditRequest struct {
	BaseRequest
}

// NewPollsEditRequest creates a new request for polls.edit
func NewPollsEditRequest(a *api.API, actor actor.Actor) *PollsEditRequest {
	return &PollsEditRequest{*NewMethodBaseRequest(a, actor, "polls.edit")}
}

// Exec executes the request and unmarshals the response into PollsEditResponse
func (r *PollsEditRequest) Exec(ctx context.Context) (response response.PollsEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsGetBackgroundsRequest defines the request for polls.getBackgrounds
//
// Returns options for background images for polls.
// Doc: https://dev.vk.com/method/polls.getBackgrounds
type PollsGetBackgroundsRequest struct {
	BaseRequest
}

// NewPollsGetBackgroundsRequest creates a new request for polls.getBackgrounds
func NewPollsGetBackgroundsRequest(a *api.API, actor actor.Actor) *PollsGetBackgroundsRequest {
	return &PollsGetBackgroundsRequest{*NewMethodBaseRequest(a, actor, "polls.getBackgrounds")}
}

// Exec executes the request and unmarshals the response into PollsGetBackgroundsResponse
func (r *PollsGetBackgroundsRequest) Exec(ctx context.Context) (response response.PollsGetBackgroundsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsGetByIDRequest defines the request for polls.getById
//
// Returns detailed information about a poll by its identifier.
// Doc: https://dev.vk.com/method/polls.getById
type PollsGetByIDRequest struct {
	BaseRequest
}

// NewPollsGetByIDRequest creates a new request for polls.getById
func NewPollsGetByIDRequest(a *api.API, actor actor.Actor) *PollsGetByIDRequest {
	return &PollsGetByIDRequest{*NewMethodBaseRequest(a, actor, "polls.getById")}
}

// Exec executes the request and unmarshals the response into PollsGetByIDResponse
func (r *PollsGetByIDRequest) Exec(ctx context.Context) (response response.PollsGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsGetPhotoUploadServerRequest defines the request for polls.getPhotoUploadServer
//
// Returns the server address for uploading background photos to the poll.
// Doc: https://dev.vk.com/method/polls.getPhotoUploadServer
type PollsGetPhotoUploadServerRequest struct {
	BaseRequest
}

// NewPollsGetPhotoUploadServerRequest creates a new request for polls.getPhotoUploadServer
func NewPollsGetPhotoUploadServerRequest(a *api.API, actor actor.Actor) *PollsGetPhotoUploadServerRequest {
	return &PollsGetPhotoUploadServerRequest{*NewMethodBaseRequest(a, actor, "polls.getPhotoUploadServer")}
}

// Exec executes the request and unmarshals the response into PollsGetPhotoUploadServerResponse
func (r *PollsGetPhotoUploadServerRequest) Exec(ctx context.Context) (response response.PollsGetPhotoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsGetVotersRequest defines the request for polls.getVoters
//
// Gets a list of user identifiers who chose specific answer options in the poll.
// Doc: https://dev.vk.com/method/polls.getVoters
type PollsGetVotersRequest struct {
	BaseRequest
}

// NewPollsGetVotersRequest creates a new request for polls.getVoters
func NewPollsGetVotersRequest(a *api.API, actor actor.Actor) *PollsGetVotersRequest {
	return &PollsGetVotersRequest{*NewMethodBaseRequest(a, actor, "polls.getVoters")}
}

// Exec executes the request and unmarshals the response into PollsGetVotersResponse
func (r *PollsGetVotersRequest) Exec(ctx context.Context) (response response.PollsGetVotersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PollsSavePhotoRequest defines the request for polls.savePhoto
//
// Saves a photo uploaded to the poll.
// Doc: https://dev.vk.com/method/polls.savePhoto
type PollsSavePhotoRequest struct {
	BaseRequest
}

// NewPollsSavePhotoRequest creates a new request for polls.savePhoto
func NewPollsSavePhotoRequest(a *api.API, actor actor.Actor) *PollsSavePhotoRequest {
	return &PollsSavePhotoRequest{*NewMethodBaseRequest(a, actor, "polls.savePhoto")}
}

// Exec executes the request and unmarshals the response into PollsSavePhotoResponse
func (r *PollsSavePhotoRequest) Exec(ctx context.Context) (response response.PollsSavePhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
