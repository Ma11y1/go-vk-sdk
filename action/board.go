package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Board Doc: https://dev.vk.com/ru/method/board
type Board struct {
	BaseAction
}

// AddTopic Doc: https://dev.vk.com/ru/method/board.addTopic
func (b *Board) AddTopic(actor actor.Actor) *request.BoardAddTopicRequest {
	return request.NewBoardAddTopicRequest(b.api, actor)
}

// CloseTopic Doc: https://dev.vk.com/ru/method/board.closeTopic
func (b *Board) CloseTopic(actor actor.Actor) *request.BoardCloseTopicRequest {
	return request.NewBoardCloseTopicRequest(b.api, actor)
}

// CreateComment Doc: https://dev.vk.com/ru/method/board.createComment
func (b *Board) CreateComment(actor actor.Actor) *request.BoardCreateCommentRequest {
	return request.NewBoardCreateCommentRequest(b.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/ru/method/board.deleteComment
func (b *Board) DeleteComment(actor actor.Actor) *request.BoardDeleteCommentRequest {
	return request.NewBoardDeleteCommentRequest(b.api, actor)
}

// DeleteTopic Doc: https://dev.vk.com/ru/method/board.deleteTopic
func (b *Board) DeleteTopic(actor actor.Actor) *request.BoardDeleteTopicRequest {
	return request.NewBoardDeleteTopicRequest(b.api, actor)
}

// EditComment Doc: https://dev.vk.com/ru/method/board.editComment
func (b *Board) EditComment(actor actor.Actor) *request.BoardEditCommentRequest {
	return request.NewBoardEditCommentRequest(b.api, actor)
}

// EditTopic Doc: https://dev.vk.com/ru/method/board.editTopic
func (b *Board) EditTopic(actor actor.Actor) *request.BoardEditTopicRequest {
	return request.NewBoardEditTopicRequest(b.api, actor)
}

// FixTopic Doc: https://dev.vk.com/ru/method/board.fixTopic
func (b *Board) FixTopic(actor actor.Actor) *request.BoardFixTopicRequest {
	return request.NewBoardFixTopicRequest(b.api, actor)
}

// GetComments Doc: https://dev.vk.com/ru/method/board.getComments
func (b *Board) GetComments(actor actor.Actor) *request.BoardGetCommentsRequest {
	return request.NewBoardGetCommentsRequest(b.api, actor)
}

// GetCommentsExtended Doc: https://dev.vk.com/ru/method/board.getComments
func (b *Board) GetCommentsExtended(actor actor.Actor) *request.BoardGetCommentsExtendedRequest {
	return request.NewBoardGetCommentsExtendedRequest(b.api, actor)
}

// GetTopics Doc: https://dev.vk.com/ru/method/board.getTopics
func (b *Board) GetTopics(actor actor.Actor) *request.BoardGetTopicsRequest {
	return request.NewBoardGetTopicsRequest(b.api, actor)
}

// GetTopicsExtended Doc: https://dev.vk.com/ru/method/board.getTopics
func (b *Board) GetTopicsExtended(actor actor.Actor) *request.BoardGetTopicsExtendedRequest {
	return request.NewBoardGetTopicsExtendedRequest(b.api, actor)
}

// OpenTopic Doc: https://dev.vk.com/ru/method/board.openTopic
func (b *Board) OpenTopic(actor actor.Actor) *request.BoardOpenTopicRequest {
	return request.NewBoardOpenTopicRequest(b.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/ru/method/board.restoreComment
func (b *Board) RestoreComment(actor actor.Actor) *request.BoardRestoreCommentRequest {
	return request.NewBoardRestoreCommentRequest(b.api, actor)
}

// UnfixTopic Doc: https://dev.vk.com/ru/method/board.unfixTopic
func (b *Board) UnfixTopic(actor actor.Actor) *request.BoardUnfixTopicRequest {
	return request.NewBoardUnfixTopicRequest(b.api, actor)
}
