package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/board

// BoardAddTopicRequest defines the request for board.addTopic
//
// The method creates a new topic in the group discussion list.
// Doc: https://dev.vk.com/method/board.addTopic
type BoardAddTopicRequest struct {
	BaseRequest
}

// NewBoardAddTopicRequest creates a new request for board.addTopic
func NewBoardAddTopicRequest(a *api.API, actor actor.Actor) *BoardAddTopicRequest {
	return &BoardAddTopicRequest{*NewMethodBaseRequest(a, actor, "board.addTopic")}
}

// Exec executes the request and unmarshals the response into BoardAddTopicResponse
func (r *BoardAddTopicRequest) Exec(ctx context.Context) (response response.BoardAddTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardCloseTopicRequest defines the request for board.closeTopic
//
// The method closes a discussion in the group.
// Doc: https://dev.vk.com/method/board.closeTopic
type BoardCloseTopicRequest struct {
	BaseRequest
}

// NewBoardCloseTopicRequest creates a new request for board.closeTopic
func NewBoardCloseTopicRequest(a *api.API, actor actor.Actor) *BoardCloseTopicRequest {
	return &BoardCloseTopicRequest{*NewMethodBaseRequest(a, actor, "board.closeTopic")}
}

// Exec executes the request and unmarshals the response into BoardCloseTopicResponse
func (r *BoardCloseTopicRequest) Exec(ctx context.Context) (response response.BoardCloseTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardCreateCommentRequest defines the request for board.createComment
//
// The method adds a new comment to a discussion in the group.
// Doc: https://dev.vk.com/method/board.createComment
type BoardCreateCommentRequest struct {
	BaseRequest
}

// NewBoardCreateCommentRequest creates a new request for board.createComment
func NewBoardCreateCommentRequest(a *api.API, actor actor.Actor) *BoardCreateCommentRequest {
	return &BoardCreateCommentRequest{*NewMethodBaseRequest(a, actor, "board.createComment")}
}

// Exec executes the request and unmarshals the response into BoardCreateCommentResponse
func (r *BoardCreateCommentRequest) Exec(ctx context.Context) (response response.BoardCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardDeleteCommentRequest defines the request for board.deleteComment
//
// The method deletes a comment from a discussion in the group.
// Doc: https://dev.vk.com/method/board.deleteComment
type BoardDeleteCommentRequest struct {
	BaseRequest
}

// NewBoardDeleteCommentRequest creates a new request for board.deleteComment
func NewBoardDeleteCommentRequest(a *api.API, actor actor.Actor) *BoardDeleteCommentRequest {
	return &BoardDeleteCommentRequest{*NewMethodBaseRequest(a, actor, "board.deleteComment")}
}

// Exec executes the request and unmarshals the response into BoardDeleteCommentResponse
func (r *BoardDeleteCommentRequest) Exec(ctx context.Context) (response response.BoardDeleteCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardDeleteTopicRequest defines the request for board.deleteTopic
//
// The method deletes a discussion in the group.
// Doc: https://dev.vk.com/method/board.deleteTopic
type BoardDeleteTopicRequest struct {
	BaseRequest
}

// NewBoardDeleteTopicRequest creates a new request for board.deleteTopic
func NewBoardDeleteTopicRequest(a *api.API, actor actor.Actor) *BoardDeleteTopicRequest {
	return &BoardDeleteTopicRequest{*NewMethodBaseRequest(a, actor, "board.deleteTopic")}
}

// Exec executes the request and unmarshals the response into BoardDeleteTopicResponse
func (r *BoardDeleteTopicRequest) Exec(ctx context.Context) (response response.BoardDeleteTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardEditCommentRequest defines the request for board.editComment
//
// The method edits a comment in a group discussion.
// Doc: https://dev.vk.com/method/board.editComment
type BoardEditCommentRequest struct {
	BaseRequest
}

// NewBoardEditCommentRequest creates a new request for board.editComment
func NewBoardEditCommentRequest(a *api.API, actor actor.Actor) *BoardEditCommentRequest {
	return &BoardEditCommentRequest{*NewMethodBaseRequest(a, actor, "board.editComment")}
}

// Exec executes the request and unmarshals the response into BoardEditCommentResponse
func (r *BoardEditCommentRequest) Exec(ctx context.Context) (response response.BoardEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardEditTopicRequest defines the request for board.editTopic
//
// The method edits the title of a discussion in the group.
// Doc: https://dev.vk.com/method/board.editTopic
type BoardEditTopicRequest struct {
	BaseRequest
}

// NewBoardEditTopicRequest creates a new request for board.editTopic
func NewBoardEditTopicRequest(a *api.API, actor actor.Actor) *BoardEditTopicRequest {
	return &BoardEditTopicRequest{*NewMethodBaseRequest(a, actor, "board.editTopic")}
}

// Exec executes the request and unmarshals the response into BoardEditTopicResponse
func (r *BoardEditTopicRequest) Exec(ctx context.Context) (response response.BoardEditTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardFixTopicRequest defines the request for board.fixTopic
//
// The method pins a discussion in the group.
// Doc: https://dev.vk.com/method/board.fixTopic
type BoardFixTopicRequest struct {
	BaseRequest
}

// NewBoardFixTopicRequest creates a new request for board.fixTopic
func NewBoardFixTopicRequest(a *api.API, actor actor.Actor) *BoardFixTopicRequest {
	return &BoardFixTopicRequest{*NewMethodBaseRequest(a, actor, "board.fixTopic")}
}

// Exec executes the request and unmarshals the response into BoardFixTopicResponse
func (r *BoardFixTopicRequest) Exec(ctx context.Context) (response response.BoardFixTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardGetCommentsRequest defines the request for board.getComments
//
// The method returns a list of comments from a specified topic.
// Doc: https://dev.vk.com/method/board.getComments
type BoardGetCommentsRequest struct {
	BaseRequest
}

// NewBoardGetCommentsRequest creates a new request for board.getComments
func NewBoardGetCommentsRequest(a *api.API, actor actor.Actor) *BoardGetCommentsRequest {
	return &BoardGetCommentsRequest{*NewMethodBaseRequest(a, actor, "board.getComments")}
}

// Exec executes the request and unmarshals the response into BoardGetCommentsResponse
func (r *BoardGetCommentsRequest) Exec(ctx context.Context) (response response.BoardGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardGetTopicsRequest defines the request for board.getTopics
//
// The method returns a list of topics in the group discussions.
// Doc: https://dev.vk.com/method/board.getTopics
type BoardGetTopicsRequest struct {
	BaseRequest
}

// NewBoardGetTopicsRequest creates a new request for board.getTopics
func NewBoardGetTopicsRequest(a *api.API, actor actor.Actor) *BoardGetTopicsRequest {
	return &BoardGetTopicsRequest{*NewMethodBaseRequest(a, actor, "board.getTopics")}
}

// Exec executes the request and unmarshals the response into BoardGetTopicsResponse
func (r *BoardGetTopicsRequest) Exec(ctx context.Context) (response response.BoardGetTopicsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardOpenTopicRequest defines the request for board.openTopic
//
// The method reopens a previously closed discussion.
// Doc: https://dev.vk.com/method/board.openTopic
type BoardOpenTopicRequest struct {
	BaseRequest
}

// NewBoardOpenTopicRequest creates a new request for board.openTopic
func NewBoardOpenTopicRequest(a *api.API, actor actor.Actor) *BoardOpenTopicRequest {
	return &BoardOpenTopicRequest{*NewMethodBaseRequest(a, actor, "board.openTopic")}
}

// Exec executes the request and unmarshals the response into BoardOpenTopicResponse
func (r *BoardOpenTopicRequest) Exec(ctx context.Context) (response response.BoardOpenTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardRestoreCommentRequest defines the request for board.restoreComment
//
// The method restores a deleted comment in a group discussion.
// Doc: https://dev.vk.com/method/board.restoreComment
type BoardRestoreCommentRequest struct {
	BaseRequest
}

// NewBoardRestoreCommentRequest creates a new request for board.restoreComment
func NewBoardRestoreCommentRequest(a *api.API, actor actor.Actor) *BoardRestoreCommentRequest {
	return &BoardRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "board.restoreComment")}
}

// Exec executes the request and unmarshals the response into BoardRestoreCommentResponse
func (r *BoardRestoreCommentRequest) Exec(ctx context.Context) (response response.BoardRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// BoardUnfixTopicRequest defines the request for board.unfixTopic
//
// The method unpins a discussion from the top of the group discussions list.
// Doc: https://dev.vk.com/method/board.unfixTopic
type BoardUnfixTopicRequest struct {
	BaseRequest
}

// NewBoardUnfixTopicRequest creates a new request for board.unfixTopic
func NewBoardUnfixTopicRequest(a *api.API, actor actor.Actor) *BoardUnfixTopicRequest {
	return &BoardUnfixTopicRequest{*NewMethodBaseRequest(a, actor, "board.unfixTopic")}
}

// Exec executes the request and unmarshals the response into BoardUnfixTopicResponse
func (r *BoardUnfixTopicRequest) Exec(ctx context.Context) (response response.BoardUnfixTopicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
