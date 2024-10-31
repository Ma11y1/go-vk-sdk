package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/notes

// NotesAddRequest defines the request for notes.add
//
// Creates a new note for the current user.
// Doc: https://dev.vk.com/method/notes.add
type NotesAddRequest struct {
	BaseRequest
}

// NewNotesAddRequest creates a new request for notes.add
func NewNotesAddRequest(a *api.API, actor actor.Actor) *NotesAddRequest {
	return &NotesAddRequest{*NewMethodBaseRequest(a, actor, "notes.add")}
}

// Exec executes the request and unmarshals the response into NotesAddResponse
func (r *NotesAddRequest) Exec(ctx context.Context) (response response.NotesAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesCreateCommentRequest defines the request for notes.createComment
//
// Adds a new comment to a note.
// Doc: https://dev.vk.com/method/notes.createComment
type NotesCreateCommentRequest struct {
	BaseRequest
}

// NewNotesCreateCommentRequest creates a new request for notes.createComment
func NewNotesCreateCommentRequest(a *api.API, actor actor.Actor) *NotesCreateCommentRequest {
	return &NotesCreateCommentRequest{*NewMethodBaseRequest(a, actor, "notes.createComment")}
}

// Exec executes the request and unmarshals the response into NotesCreateCommentResponse
func (r *NotesCreateCommentRequest) Exec(ctx context.Context) (response response.NotesCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesDeleteRequest defines the request for notes.delete
//
// Deletes a note of the current user.
// Doc: https://dev.vk.com/method/notes.delete
type NotesDeleteRequest struct {
	BaseRequest
}

// NewNotesDeleteRequest creates a new request for notes.delete
func NewNotesDeleteRequest(a *api.API, actor actor.Actor) *NotesDeleteRequest {
	return &NotesDeleteRequest{*NewMethodBaseRequest(a, actor, "notes.delete")}
}

// Exec executes the request and unmarshals the response into NotesDeleteResponse
func (r *NotesDeleteRequest) Exec(ctx context.Context) (response response.NotesDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesEditRequest defines the request for notes.edit
//
// Edits a note of the current user.
// Doc: https://dev.vk.com/method/notes.edit
type NotesEditRequest struct {
	BaseRequest
}

// NewNotesEditRequest creates a new request for notes.edit
func NewNotesEditRequest(a *api.API, actor actor.Actor) *NotesEditRequest {
	return &NotesEditRequest{*NewMethodBaseRequest(a, actor, "notes.edit")}
}

// Exec executes the request and unmarshals the response into NotesEditResponse
func (r *NotesEditRequest) Exec(ctx context.Context) (response response.NotesEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesEditCommentRequest defines the request for notes.editComment
//
// Edits a specified comment on a note.
// Doc: https://dev.vk.com/method/notes.editComment
type NotesEditCommentRequest struct {
	BaseRequest
}

// NewNotesEditCommentRequest creates a new request for notes.editComment
func NewNotesEditCommentRequest(a *api.API, actor actor.Actor) *NotesEditCommentRequest {
	return &NotesEditCommentRequest{*NewMethodBaseRequest(a, actor, "notes.editComment")}
}

// Exec executes the request and unmarshals the response into NotesEditCommentResponse
func (r *NotesEditCommentRequest) Exec(ctx context.Context) (response response.NotesEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesGetRequest defines the request for notes.get
//
// Returns a list of notes created by the user.
// Doc: https://dev.vk.com/method/notes.get
type NotesGetRequest struct {
	BaseRequest
}

// NewNotesGetRequest creates a new request for notes.get
func NewNotesGetRequest(a *api.API, actor actor.Actor) *NotesGetRequest {
	return &NotesGetRequest{*NewMethodBaseRequest(a, actor, "notes.get")}
}

// Exec executes the request and unmarshals the response into NotesGetResponse
func (r *NotesGetRequest) Exec(ctx context.Context) (response response.NotesGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesGetByIdRequest defines the request for notes.getById
//
// Returns a note by its ID.
// Doc: https://dev.vk.com/method/notes.getById
type NotesGetByIdRequest struct {
	BaseRequest
}

// NewNotesGetByIdRequest creates a new request for notes.getById
func NewNotesGetByIdRequest(a *api.API, actor actor.Actor) *NotesGetByIdRequest {
	return &NotesGetByIdRequest{*NewMethodBaseRequest(a, actor, "notes.getById")}
}

// Exec executes the request and unmarshals the response into NotesGetByIdResponse
func (r *NotesGetByIdRequest) Exec(ctx context.Context) (response response.NotesGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesGetCommentsRequest defines the request for notes.getComments
//
// Returns a list of comments on a note.
// Doc: https://dev.vk.com/method/notes.getComments
type NotesGetCommentsRequest struct {
	BaseRequest
}

// NewNotesGetCommentsRequest creates a new request for notes.getComments
func NewNotesGetCommentsRequest(a *api.API, actor actor.Actor) *NotesGetCommentsRequest {
	return &NotesGetCommentsRequest{*NewMethodBaseRequest(a, actor, "notes.getComments")}
}

// Exec executes the request and unmarshals the response into NotesGetCommentsResponse
func (r *NotesGetCommentsRequest) Exec(ctx context.Context) (response response.NotesGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NotesRestoreCommentRequest defines the request for notes.restoreComment
//
// Restores a deleted comment.
// Doc: https://dev.vk.com/method/notes.restoreComment
type NotesRestoreCommentRequest struct {
	BaseRequest
}

// NewNotesRestoreCommentRequest creates a new request for notes.restoreComment
func NewNotesRestoreCommentRequest(a *api.API, actor actor.Actor) *NotesRestoreCommentRequest {
	return &NotesRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "notes.restoreComment")}
}

// Exec executes the request and unmarshals the response into NotesRestoreCommentResponse
func (r *NotesRestoreCommentRequest) Exec(ctx context.Context) (response response.NotesRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
