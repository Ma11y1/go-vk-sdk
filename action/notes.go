package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Notes Doc: https://dev.vk.com/ru/method/notes
type Notes struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/ru/method/notes.add
func (n *Notes) Add(actor actor.Actor) *request.NotesAddRequest {
	return request.NewNotesAddRequest(n.api, actor)
}

// CreateComment Doc: https://dev.vk.com/ru/method/notes.createComment
func (n *Notes) CreateComment(actor actor.Actor) *request.NotesCreateCommentRequest {
	return request.NewNotesCreateCommentRequest(n.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/notes.delete
func (n *Notes) Delete(actor actor.Actor) *request.NotesDeleteRequest {
	return request.NewNotesDeleteRequest(n.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/ru/method/notes.deleteComment
func (n *Notes) DeleteComment(actor actor.Actor) *request.NotesDeleteCommentRequest {
	return request.NewNotesDeleteCommentRequest(n.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/notes.edit
func (n *Notes) Edit(actor actor.Actor) *request.NotesEditRequest {
	return request.NewNotesEditRequest(n.api, actor)
}

// EditComment Doc: https://dev.vk.com/ru/method/notes.editComment
func (n *Notes) EditComment(actor actor.Actor) *request.NotesEditCommentRequest {
	return request.NewNotesEditCommentRequest(n.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/notes.get
func (n *Notes) Get(actor actor.Actor) *request.NotesGetRequest {
	return request.NewNotesGetRequest(n.api, actor)
}

// GetById Doc: https://dev.vk.com/ru/method/notes.getById
func (n *Notes) GetById(actor actor.Actor) *request.NotesGetByIDRequest {
	return request.NewNotesGetByIDRequest(n.api, actor)
}

// GetComments Doc: https://dev.vk.com/ru/method/notes.getComments
func (n *Notes) GetComments(actor actor.Actor) *request.NotesGetCommentsRequest {
	return request.NewNotesGetCommentsRequest(n.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/ru/method/notes.restoreComment
func (n *Notes) RestoreComment(actor actor.Actor) *request.NotesRestoreCommentRequest {
	return request.NewNotesRestoreCommentRequest(n.api, actor)
}
