package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"strconv"
)

// Doc: https://dev.vk.com/ru/method/messages

// Methods for working with personal messages.
//  To instantly receive incoming messages, use the LongPoll server. https://dev.vk.com/ru/api/user-long-poll/getting-started
//  Please note: access to working with section methods with a user key is limited.
//  Information about the Messages api limitation is in the Roadmap. https://dev.vk.com/ru/reference/roadmap#%D0%9E%D0%B3%D1%80%D0%B0%D0%BD%D0%B8%D1%87%D0%B5%D0%BD%D0%B8%D0%B5%20Messages%20API
//  Please note: methods for working with calls have been moved to the new calls section. Old calling methods from the messages section have been marked deprecated and may be removed in future versions of the api.

// MessagesAddChatUserRequest defines the request for messages.addChatUser
//
// Adds a new user to a group chat.
// Doc: https://dev.vk.com/method/messages.addChatUser
type MessagesAddChatUserRequest struct {
	BaseRequest
}

// NewMessagesAddChatUserRequest creates a new request for messages.addChatUser
func NewMessagesAddChatUserRequest(a *api.API, actor actor.Actor) *MessagesAddChatUserRequest {
	return &MessagesAddChatUserRequest{*NewMethodBaseRequest(a, actor, "messages.addChatUser")}
}

// Exec executes the request and unmarshals the response into MessagesAddChatUserResponse
func (r *MessagesAddChatUserRequest) Exec(ctx context.Context) (response response.MessagesAddChatUserResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesAllowMessagesFromGroupRequest defines the request for messages.allowMessagesFromGroup
//
// Allows a community to send messages to the current user.
// Doc: https://dev.vk.com/method/messages.allowMessagesFromGroup
type MessagesAllowMessagesFromGroupRequest struct {
	BaseRequest
}

// NewMessagesAllowMessagesFromGroupRequest creates a new request for messages.allowMessagesFromGroup
func NewMessagesAllowMessagesFromGroupRequest(a *api.API, actor actor.Actor) *MessagesAllowMessagesFromGroupRequest {
	return &MessagesAllowMessagesFromGroupRequest{*NewMethodBaseRequest(a, actor, "messages.allowMessagesFromGroup")}
}

// Exec executes the request and unmarshals the response into MessagesAllowMessagesFromGroupResponse
func (r *MessagesAllowMessagesFromGroupRequest) Exec(ctx context.Context) (response response.MessagesAllowMessagesFromGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesCreateChatRequest defines the request for messages.createChat
//
// Creates a chat with multiple participants.
// Doc: https://dev.vk.com/method/messages.createChat
type MessagesCreateChatRequest struct {
	BaseRequest
}

// NewMessagesCreateChatRequest creates a new request for messages.createChat
func NewMessagesCreateChatRequest(a *api.API, actor actor.Actor) *MessagesCreateChatRequest {
	return &MessagesCreateChatRequest{*NewMethodBaseRequest(a, actor, "messages.createChat")}
}

// Exec executes the request and unmarshals the response into MessagesCreateChatResponse
func (r *MessagesCreateChatRequest) Exec(ctx context.Context) (response response.MessagesCreateChatResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesDeleteRequest defines the request for messages.delete
//
// Deletes a message.
// Doc: https://dev.vk.com/method/messages.delete
type MessagesDeleteRequest struct {
	BaseRequest
}

// NewMessagesDeleteRequest creates a new request for messages.delete
func NewMessagesDeleteRequest(a *api.API, actor actor.Actor) *MessagesDeleteRequest {
	return &MessagesDeleteRequest{*NewMethodBaseRequest(a, actor, "messages.delete")}
}

// Exec executes the request and unmarshals the response into MessagesDeleteResponse
func (r *MessagesDeleteRequest) Exec(ctx context.Context) (response response.MessagesDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesDeleteChatPhotoRequest defines the request for messages.deleteChatPhoto
//
// Deletes the chat photo.
// Doc: https://dev.vk.com/method/messages.deleteChatPhoto
type MessagesDeleteChatPhotoRequest struct {
	BaseRequest
}

// NewMessagesDeleteChatPhotoRequest creates a new request for messages.deleteChatPhoto
func NewMessagesDeleteChatPhotoRequest(a *api.API, actor actor.Actor) *MessagesDeleteChatPhotoRequest {
	return &MessagesDeleteChatPhotoRequest{*NewMethodBaseRequest(a, actor, "messages.deleteChatPhoto")}
}

// Exec executes the request and unmarshals the response into MessagesDeleteChatPhotoResponse
func (r *MessagesDeleteChatPhotoRequest) Exec(ctx context.Context) (response response.MessagesDeleteChatPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesDeleteConversationRequest defines the request for messages.deleteConversation
//
// Deletes a conversation.
// Doc: https://dev.vk.com/method/messages.deleteConversation
type MessagesDeleteConversationRequest struct {
	BaseRequest
}

// NewMessagesDeleteConversationRequest creates a new request for messages.deleteConversation
func NewMessagesDeleteConversationRequest(a *api.API, actor actor.Actor) *MessagesDeleteConversationRequest {
	return &MessagesDeleteConversationRequest{*NewMethodBaseRequest(a, actor, "messages.deleteConversation")}
}

// Exec executes the request and unmarshals the response into MessagesDeleteConversationResponse
func (r *MessagesDeleteConversationRequest) Exec(ctx context.Context) (response response.MessagesDeleteConversationResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesDeleteReactionRequest defines the request for messages.deleteReaction
//
// Deletes a previously set reaction.
// Doc: https://dev.vk.com/method/messages.deleteReaction
type MessagesDeleteReactionRequest struct {
	BaseRequest
}

// NewMessagesDeleteReactionRequest creates a new request for messages.deleteReaction
func NewMessagesDeleteReactionRequest(a *api.API, actor actor.Actor) *MessagesDeleteReactionRequest {
	return &MessagesDeleteReactionRequest{*NewMethodBaseRequest(a, actor, "messages.deleteReaction")}
}

// Exec executes the request and unmarshals the response into MessagesDeleteReactionResponse
func (r *MessagesDeleteReactionRequest) Exec(ctx context.Context) (response response.MessagesDeleteReactionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesDenyMessagesFromGroupRequest defines the request for messages.denyMessagesFromGroup
//
// Blocks a community from sending messages to the current user.
// Doc: https://dev.vk.com/method/messages.denyMessagesFromGroup
type MessagesDenyMessagesFromGroupRequest struct {
	BaseRequest
}

// NewMessagesDenyMessagesFromGroupRequest creates a new request for messages.denyMessagesFromGroup
func NewMessagesDenyMessagesFromGroupRequest(a *api.API, actor actor.Actor) *MessagesDenyMessagesFromGroupRequest {
	return &MessagesDenyMessagesFromGroupRequest{*NewMethodBaseRequest(a, actor, "messages.denyMessagesFromGroup")}
}

// Exec executes the request and unmarshals the response into MessagesDenyMessagesFromGroupResponse
func (r *MessagesDenyMessagesFromGroupRequest) Exec(ctx context.Context) (response response.MessagesDenyMessagesFromGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesEditRequest defines the request for messages.edit
//
// Edits a message.
// Doc: https://dev.vk.com/method/messages.edit
type MessagesEditRequest struct {
	BaseRequest
}

// NewMessagesEditRequest creates a new request for messages.edit
func NewMessagesEditRequest(a *api.API, actor actor.Actor) *MessagesEditRequest {
	return &MessagesEditRequest{*NewMethodBaseRequest(a, actor, "messages.edit")}
}

// Exec executes the request and unmarshals the response into MessagesEditResponse
func (r *MessagesEditRequest) Exec(ctx context.Context) (response response.MessagesEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesEditChatRequest defines the request for messages.editChat
//
// Changes the chat name.
// Doc: https://dev.vk.com/method/messages.editChat
type MessagesEditChatRequest struct {
	BaseRequest
}

// NewMessagesEditChatRequest creates a new request for messages.editChat
func NewMessagesEditChatRequest(a *api.API, actor actor.Actor) *MessagesEditChatRequest {
	return &MessagesEditChatRequest{*NewMethodBaseRequest(a, actor, "messages.editChat")}
}

// Exec executes the request and unmarshals the response into MessagesEditChatResponse
func (r *MessagesEditChatRequest) Exec(ctx context.Context) (response response.MessagesEditChatResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetByConversationMessageIdRequest defines the request for messages.getByConversationMessageId
//
// Returns messages by conversation_message_id.
// Doc: https://dev.vk.com/method/messages.getByConversationMessageId
type MessagesGetByConversationMessageIdRequest struct {
	BaseRequest
}

// NewMessagesGetByConversationMessageIdRequest creates a new request for messages.getByConversationMessageId
func NewMessagesGetByConversationMessageIdRequest(a *api.API, actor actor.Actor) *MessagesGetByConversationMessageIdRequest {
	return &MessagesGetByConversationMessageIdRequest{*NewMethodBaseRequest(a, actor, "messages.getByConversationMessageId")}
}

// Exec executes the request and unmarshals the response into MessagesGetByConversationMessageIdResponse
func (r *MessagesGetByConversationMessageIdRequest) Exec(ctx context.Context) (response response.MessagesGetByConversationMessageIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetByIdRequest defines the request for messages.getById
//
// Returns messages by their IDs.
// Doc: https://dev.vk.com/method/messages.getById
type MessagesGetByIdRequest struct {
	BaseRequest
}

// NewMessagesGetByIdRequest creates a new request for messages.getById
func NewMessagesGetByIdRequest(a *api.API, actor actor.Actor) *MessagesGetByIdRequest {
	return &MessagesGetByIdRequest{*NewMethodBaseRequest(a, actor, "messages.getById")}
}

// Exec executes the request and unmarshals the response into MessagesGetByIdResponse
func (r *MessagesGetByIdRequest) Exec(ctx context.Context) (response response.MessagesGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetChatRequest defines the request for messages.getChat
//
// Returns information about a chat.
// Doc: https://dev.vk.com/method/messages.getChat
type MessagesGetChatRequest struct {
	BaseRequest
}

// NewMessagesGetChatRequest creates a new request for messages.getChat
func NewMessagesGetChatRequest(a *api.API, actor actor.Actor) *MessagesGetChatRequest {
	return &MessagesGetChatRequest{*NewMethodBaseRequest(a, actor, "messages.getChat")}
}

// Exec executes the request and unmarshals the response into MessagesGetChatResponse
func (r *MessagesGetChatRequest) Exec(ctx context.Context) (response response.MessagesGetChatResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetChatPreviewRequest defines the request for messages.getChatPreview
//
// Gets data for a chat preview via invitation link.
// Doc: https://dev.vk.com/method/messages.getChatPreview
type MessagesGetChatPreviewRequest struct {
	BaseRequest
}

// NewMessagesGetChatPreviewRequest creates a new request for messages.getChatPreview
func NewMessagesGetChatPreviewRequest(a *api.API, actor actor.Actor) *MessagesGetChatPreviewRequest {
	return &MessagesGetChatPreviewRequest{*NewMethodBaseRequest(a, actor, "messages.getChatPreview")}
}

// Exec executes the request and unmarshals the response into MessagesGetChatPreviewResponse
func (r *MessagesGetChatPreviewRequest) Exec(ctx context.Context) (response response.MessagesGetChatPreviewResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetConversationMembersRequest defines the request for messages.getConversationMembers
//
// Returns the list of conversation participants.
// Doc: https://dev.vk.com/method/messages.getConversationMembers
type MessagesGetConversationMembersRequest struct {
	BaseRequest
}

// NewMessagesGetConversationMembersRequest creates a new request for messages.getConversationMembers
func NewMessagesGetConversationMembersRequest(a *api.API, actor actor.Actor) *MessagesGetConversationMembersRequest {
	return &MessagesGetConversationMembersRequest{*NewMethodBaseRequest(a, actor, "messages.getConversationMembers")}
}

// Exec executes the request and unmarshals the response into MessagesGetConversationMembersResponse
func (r *MessagesGetConversationMembersRequest) Exec(ctx context.Context) (response response.MessagesGetConversationMembersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetConversationsRequest defines the request for messages.getConversations
//
// Returns the user's conversation list.
// Doc: https://dev.vk.com/method/messages.getConversations
type MessagesGetConversationsRequest struct {
	BaseRequest
}

// NewMessagesGetConversationsRequest creates a new request for messages.getConversations
func NewMessagesGetConversationsRequest(a *api.API, actor actor.Actor) *MessagesGetConversationsRequest {
	return &MessagesGetConversationsRequest{*NewMethodBaseRequest(a, actor, "messages.getConversations")}
}

// Exec executes the request and unmarshals the response into MessagesGetConversationsResponse
func (r *MessagesGetConversationsRequest) Exec(ctx context.Context) (response response.MessagesGetConversationsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetConversationsByIdRequest defines the request for messages.getConversationsById
//
// Returns a conversation by its ID.
// Doc: https://dev.vk.com/method/messages.getConversationsById
type MessagesGetConversationsByIdRequest struct {
	BaseRequest
}

// NewMessagesGetConversationsByIdRequest creates a new request for messages.getConversationsById
func NewMessagesGetConversationsByIdRequest(a *api.API, actor actor.Actor) *MessagesGetConversationsByIdRequest {
	return &MessagesGetConversationsByIdRequest{*NewMethodBaseRequest(a, actor, "messages.getConversationsById")}
}

// Exec executes the request and unmarshals the response into MessagesGetConversationsByIdResponse
func (r *MessagesGetConversationsByIdRequest) Exec(ctx context.Context) (response response.MessagesGetConversationsByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetHistoryRequest defines the request for messages.getHistory
//
// Returns the message history for a specified dialogue.
// Doc: https://dev.vk.com/method/messages.getHistory
type MessagesGetHistoryRequest struct {
	BaseRequest
}

// NewMessagesGetHistoryRequest creates a new request for messages.getHistory
func NewMessagesGetHistoryRequest(a *api.API, actor actor.Actor) *MessagesGetHistoryRequest {
	return &MessagesGetHistoryRequest{*NewMethodBaseRequest(a, actor, "messages.getHistory")}
}

// Exec executes the request and unmarshals the response into MessagesGetHistoryResponse
func (r *MessagesGetHistoryRequest) Exec(ctx context.Context) (response response.MessagesGetHistoryResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetHistoryAttachmentsRequest defines the request for messages.getHistoryAttachments
//
// Returns attachments from a dialogue or conversation.
// Doc: https://dev.vk.com/method/messages.getHistoryAttachments
type MessagesGetHistoryAttachmentsRequest struct {
	BaseRequest
}

// NewMessagesGetHistoryAttachmentsRequest creates a new request for messages.getHistoryAttachments
func NewMessagesGetHistoryAttachmentsRequest(a *api.API, actor actor.Actor) *MessagesGetHistoryAttachmentsRequest {
	return &MessagesGetHistoryAttachmentsRequest{*NewMethodBaseRequest(a, actor, "messages.getHistoryAttachments")}
}

// Exec executes the request and unmarshals the response into MessagesGetHistoryAttachmentsResponse
func (r *MessagesGetHistoryAttachmentsRequest) Exec(ctx context.Context) (response response.MessagesGetHistoryAttachmentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetImportantMessagesRequest defines the request for messages.getImportantMessages
//
// Returns a list of important messages for the user.
// Doc: https://dev.vk.com/method/messages.getImportantMessages
type MessagesGetImportantMessagesRequest struct {
	BaseRequest
}

// NewMessagesGetImportantMessagesRequest creates a new request for messages.getImportantMessages
func NewMessagesGetImportantMessagesRequest(a *api.API, actor actor.Actor) *MessagesGetImportantMessagesRequest {
	return &MessagesGetImportantMessagesRequest{*NewMethodBaseRequest(a, actor, "messages.getImportantMessages")}
}

// Exec executes the request and unmarshals the response into MessagesGetImportantMessagesResponse
func (r *MessagesGetImportantMessagesRequest) Exec(ctx context.Context) (response response.MessagesGetImportantMessagesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetIntentUsersRequest defines the request for messages.getIntentUsers
//
// Returns users who subscribed to specific intents.
// Doc: https://dev.vk.com/method/messages.getIntentUsers
type MessagesGetIntentUsersRequest struct {
	BaseRequest
}

// NewMessagesGetIntentUsersRequest creates a new request for messages.getIntentUsers
func NewMessagesGetIntentUsersRequest(a *api.API, actor actor.Actor) *MessagesGetIntentUsersRequest {
	return &MessagesGetIntentUsersRequest{*NewMethodBaseRequest(a, actor, "messages.getIntentUsers")}
}

// Exec executes the request and unmarshals the response into MessagesGetIntentUsersResponse
func (r *MessagesGetIntentUsersRequest) Exec(ctx context.Context) (response response.MessagesGetIntentUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetInviteLinkRequest defines the request for messages.getInviteLink
//
// Gets an invite link for the conversation.
// Doc: https://dev.vk.com/method/messages.getInviteLink
type MessagesGetInviteLinkRequest struct {
	BaseRequest
}

// NewMessagesGetInviteLinkRequest creates a new request for messages.getInviteLink
func NewMessagesGetInviteLinkRequest(a *api.API, actor actor.Actor) *MessagesGetInviteLinkRequest {
	return &MessagesGetInviteLinkRequest{*NewMethodBaseRequest(a, actor, "messages.getInviteLink")}
}

// Exec executes the request and unmarshals the response into MessagesGetInviteLinkResponse
func (r *MessagesGetInviteLinkRequest) Exec(ctx context.Context) (response response.MessagesGetInviteLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetLastActivityRequest defines the request for messages.getLastActivity
//
// Gets the current status and the date of the user's last activity.
// Doc: https://dev.vk.com/method/messages.getLastActivity
type MessagesGetLastActivityRequest struct {
	BaseRequest
}

// NewMessagesGetLastActivityRequest creates a new request for messages.getLastActivity
func NewMessagesGetLastActivityRequest(a *api.API, actor actor.Actor) *MessagesGetLastActivityRequest {
	return &MessagesGetLastActivityRequest{*NewMethodBaseRequest(a, actor, "messages.getLastActivity")}
}

// Exec executes the request and unmarshals the response into MessagesGetLastActivityResponse
func (r *MessagesGetLastActivityRequest) Exec(ctx context.Context) (response response.MessagesGetLastActivityResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetLongPollHistoryRequest defines the request for messages.getLongPollHistory
//
// Returns updates in the user's private messages.
// Doc: https://dev.vk.com/method/messages.getLongPollHistory
type MessagesGetLongPollHistoryRequest struct {
	BaseRequest
}

// NewMessagesGetLongPollHistoryRequest creates a new request for messages.getLongPollHistory
func NewMessagesGetLongPollHistoryRequest(a *api.API, actor actor.Actor) *MessagesGetLongPollHistoryRequest {
	return &MessagesGetLongPollHistoryRequest{*NewMethodBaseRequest(a, actor, "messages.getLongPollHistory")}
}

// Exec executes the request and unmarshals the response into MessagesGetLongPollHistoryResponse
func (r *MessagesGetLongPollHistoryRequest) Exec(ctx context.Context) (response response.MessagesGetLongPollHistoryResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetLongPollServerRequest defines the request for messages.getLongPollServer
//
// Returns data required for connecting to the Long Poll server.
// Doc: https://dev.vk.com/method/messages.getLongPollServer
type MessagesGetLongPollServerRequest struct {
	BaseRequest
}

// NewMessagesGetLongPollServerRequest creates a new request for messages.getLongPollServer
func NewMessagesGetLongPollServerRequest(a *api.API, actor actor.Actor) *MessagesGetLongPollServerRequest {
	return &MessagesGetLongPollServerRequest{*NewMethodBaseRequest(a, actor, "messages.getLongPollServer")}
}

// Exec executes the request and unmarshals the response into MessagesGetLongPollServerResponse
func (r *MessagesGetLongPollServerRequest) Exec(ctx context.Context) (response response.MessagesGetLongPollServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// NeedPts Return the pts field required for the messages.getLongPollHistory method to work
func (r *MessagesGetLongPollServerRequest) NeedPts(v bool) *MessagesGetLongPollServerRequest {
	if v {
		r.parameters.Set(constants.ParameterNameNeedPts, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameNeedPts)
	}
	return r
}

// GroupID Community ID (for community posts with a user access key)
func (r *MessagesGetLongPollServerRequest) GroupID(id int) *MessagesGetLongPollServerRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// LpVersion version for connection to Long Poll. Current version: 3
func (r *MessagesGetLongPollServerRequest) LpVersion(v int) *MessagesGetLongPollServerRequest {
	r.parameters.Set(constants.ParameterNameLPVersion, strconv.Itoa(v))
	return r
}

// MessagesGetMessagesReactionsRequest defines the request for messages.getMessagesReactions
//
// Gets the current reaction counters on messages.
// Doc: https://dev.vk.com/method/messages.getMessagesReactions
type MessagesGetMessagesReactionsRequest struct {
	BaseRequest
}

// NewMessagesGetMessagesReactionsRequest creates a new request for messages.getMessagesReactions
func NewMessagesGetMessagesReactionsRequest(a *api.API, actor actor.Actor) *MessagesGetMessagesReactionsRequest {
	return &MessagesGetMessagesReactionsRequest{*NewMethodBaseRequest(a, actor, "messages.getMessagesReactions")}
}

// Exec executes the request and unmarshals the response into MessagesGetMessagesReactionsResponse
func (r *MessagesGetMessagesReactionsRequest) Exec(ctx context.Context) (response response.MessagesGetMessagesReactionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetReactedPeersRequest defines the request for messages.getReactedPeers
//
// Gets the list of users and communities that reacted to a message.
// Doc: https://dev.vk.com/method/messages.getReactedPeers
type MessagesGetReactedPeersRequest struct {
	BaseRequest
}

// NewMessagesGetReactedPeersRequest creates a new request for messages.getReactedPeers
func NewMessagesGetReactedPeersRequest(a *api.API, actor actor.Actor) *MessagesGetReactedPeersRequest {
	return &MessagesGetReactedPeersRequest{*NewMethodBaseRequest(a, actor, "messages.getReactedPeers")}
}

// Exec executes the request and unmarshals the response into MessagesGetReactedPeersResponse
func (r *MessagesGetReactedPeersRequest) Exec(ctx context.Context) (response response.MessagesGetReactedPeersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesGetReactionsAssetsRequest defines the request for messages.getReactionsAssets
//
// Gets reaction assets.
// Doc: https://dev.vk.com/method/messages.getReactionsAssets
type MessagesGetReactionsAssetsRequest struct {
	BaseRequest
}

// NewMessagesGetReactionsAssetsRequest creates a new request for messages.getReactionsAssets
func NewMessagesGetReactionsAssetsRequest(a *api.API, actor actor.Actor) *MessagesGetReactionsAssetsRequest {
	return &MessagesGetReactionsAssetsRequest{*NewMethodBaseRequest(a, actor, "messages.getReactionsAssets")}
}

// Exec executes the request and unmarshals the response into MessagesGetReactionsAssetsResponse
func (r *MessagesGetReactionsAssetsRequest) Exec(ctx context.Context) (response response.MessagesGetReactionsAssetsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesIsMessagesFromGroupAllowedRequest defines the request for messages.isMessagesFromGroupAllowed
//
// Checks if sending messages from a group to the user is allowed.
// Doc: https://dev.vk.com/method/messages.isMessagesFromGroupAllowed
type MessagesIsMessagesFromGroupAllowedRequest struct {
	BaseRequest
}

// NewMessagesIsMessagesFromGroupAllowedRequest creates a new request for messages.isMessagesFromGroupAllowed
func NewMessagesIsMessagesFromGroupAllowedRequest(a *api.API, actor actor.Actor) *MessagesIsMessagesFromGroupAllowedRequest {
	return &MessagesIsMessagesFromGroupAllowedRequest{*NewMethodBaseRequest(a, actor, "messages.isMessagesFromGroupAllowed")}
}

// Exec executes the request and unmarshals the response into MessagesIsMessagesFromGroupAllowedResponse
func (r *MessagesIsMessagesFromGroupAllowedRequest) Exec(ctx context.Context) (response response.MessagesIsMessagesFromGroupAllowedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesJoinChatByInviteLinkRequest defines the request for messages.joinChatByInviteLink
//
// Allows joining a chat by invitation link.
// Doc: https://dev.vk.com/method/messages.joinChatByInviteLink
type MessagesJoinChatByInviteLinkRequest struct {
	BaseRequest
}

// NewMessagesJoinChatByInviteLinkRequest creates a new request for messages.joinChatByInviteLink
func NewMessagesJoinChatByInviteLinkRequest(a *api.API, actor actor.Actor) *MessagesJoinChatByInviteLinkRequest {
	return &MessagesJoinChatByInviteLinkRequest{*NewMethodBaseRequest(a, actor, "messages.joinChatByInviteLink")}
}

// Exec executes the request and unmarshals the response into MessagesJoinChatByInviteLinkResponse
func (r *MessagesJoinChatByInviteLinkRequest) Exec(ctx context.Context) (response response.MessagesJoinChatByInviteLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesMarkAsAnsweredConversationRequest defines the request for messages.markAsAnsweredConversation
//
// Marks a conversation as answered or removes the mark.
// Doc: https://dev.vk.com/method/messages.markAsAnsweredConversation
type MessagesMarkAsAnsweredConversationRequest struct {
	BaseRequest
}

// NewMessagesMarkAsAnsweredConversationRequest creates a new request for messages.markAsAnsweredConversation
func NewMessagesMarkAsAnsweredConversationRequest(a *api.API, actor actor.Actor) *MessagesMarkAsAnsweredConversationRequest {
	return &MessagesMarkAsAnsweredConversationRequest{*NewMethodBaseRequest(a, actor, "messages.markAsAnsweredConversation")}
}

// Exec executes the request and unmarshals the response into MessagesMarkAsAnsweredConversationResponse
func (r *MessagesMarkAsAnsweredConversationRequest) Exec(ctx context.Context) (response response.MessagesMarkAsAnsweredConversationResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesMarkAsImportantRequest defines the request for messages.markAsImportant
//
// Marks messages as important or removes the mark.
// Doc: https://dev.vk.com/method/messages.markAsImportant
type MessagesMarkAsImportantRequest struct {
	BaseRequest
}

// NewMessagesMarkAsImportantRequest creates a new request for messages.markAsImportant
func NewMessagesMarkAsImportantRequest(a *api.API, actor actor.Actor) *MessagesMarkAsImportantRequest {
	return &MessagesMarkAsImportantRequest{*NewMethodBaseRequest(a, actor, "messages.markAsImportant")}
}

// Exec executes the request and unmarshals the response into MessagesMarkAsImportantResponse
func (r *MessagesMarkAsImportantRequest) Exec(ctx context.Context) (response response.MessagesMarkAsImportantResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesMarkAsImportantConversationRequest defines the request for messages.markAsImportantConversation
//
// Marks a conversation as important or removes the mark.
// Doc: https://dev.vk.com/method/messages.markAsImportantConversation
type MessagesMarkAsImportantConversationRequest struct {
	BaseRequest
}

// NewMessagesMarkAsImportantConversationRequest creates a new request for messages.markAsImportantConversation
func NewMessagesMarkAsImportantConversationRequest(a *api.API, actor actor.Actor) *MessagesMarkAsImportantConversationRequest {
	return &MessagesMarkAsImportantConversationRequest{*NewMethodBaseRequest(a, actor, "messages.markAsImportantConversation")}
}

// Exec executes the request and unmarshals the response into MessagesMarkAsImportantConversationResponse
func (r *MessagesMarkAsImportantConversationRequest) Exec(ctx context.Context) (response response.MessagesMarkAsImportantConversationResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesMarkAsReadRequest defines the request for messages.markAsRead
//
// Marks messages as read.
// Doc: https://dev.vk.com/method/messages.markAsRead
type MessagesMarkAsReadRequest struct {
	BaseRequest
}

// NewMessagesMarkAsReadRequest creates a new request for messages.markAsRead
func NewMessagesMarkAsReadRequest(a *api.API, actor actor.Actor) *MessagesMarkAsReadRequest {
	return &MessagesMarkAsReadRequest{*NewMethodBaseRequest(a, actor, "messages.markAsRead")}
}

// Exec executes the request and unmarshals the response into MessagesMarkAsReadResponse
func (r *MessagesMarkAsReadRequest) Exec(ctx context.Context) (response response.MessagesMarkAsReadResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesMarkReactionsAsReadRequest defines the request for messages.markReactionsAsRead
//
// Marks all reactions on messages as read with the specified cmids.
// Doc: https://dev.vk.com/method/messages.markReactionsAsRead
type MessagesMarkReactionsAsReadRequest struct {
	BaseRequest
}

// NewMessagesMarkReactionsAsReadRequest creates a new request for messages.markReactionsAsRead
func NewMessagesMarkReactionsAsReadRequest(a *api.API, actor actor.Actor) *MessagesMarkReactionsAsReadRequest {
	return &MessagesMarkReactionsAsReadRequest{*NewMethodBaseRequest(a, actor, "messages.markReactionsAsRead")}
}

// Exec executes the request and unmarshals the response into MessagesMarkReactionsAsReadResponse
func (r *MessagesMarkReactionsAsReadRequest) Exec(ctx context.Context) (response response.MessagesMarkReactionsAsReadResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesPinRequest defines the request for messages.pin
//
// Pins a message.
// Doc: https://dev.vk.com/method/messages.pin
type MessagesPinRequest struct {
	BaseRequest
}

// NewMessagesPinRequest creates a new request for messages.pin
func NewMessagesPinRequest(a *api.API, actor actor.Actor) *MessagesPinRequest {
	return &MessagesPinRequest{*NewMethodBaseRequest(a, actor, "messages.pin")}
}

// Exec executes the request and unmarshals the response into MessagesPinResponse
func (r *MessagesPinRequest) Exec(ctx context.Context) (response response.MessagesPinResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesRemoveChatUserRequest defines the request for messages.removeChatUser
//
// Excludes a user from a chat if the current user or community is the admin or the current user invited the excluded user.
// Doc: https://dev.vk.com/method/messages.removeChatUser
type MessagesRemoveChatUserRequest struct {
	BaseRequest
}

// NewMessagesRemoveChatUserRequest creates a new request for messages.removeChatUser
func NewMessagesRemoveChatUserRequest(a *api.API, actor actor.Actor) *MessagesRemoveChatUserRequest {
	return &MessagesRemoveChatUserRequest{*NewMethodBaseRequest(a, actor, "messages.removeChatUser")}
}

// Exec executes the request and unmarshals the response into MessagesRemoveChatUserResponse
func (r *MessagesRemoveChatUserRequest) Exec(ctx context.Context) (response response.MessagesRemoveChatUserResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesRestoreRequest defines the request for messages.restore
//
// Restores a deleted message.
// Doc: https://dev.vk.com/method/messages.restore
type MessagesRestoreRequest struct {
	BaseRequest
}

// NewMessagesRestoreRequest creates a new request for messages.restore
func NewMessagesRestoreRequest(a *api.API, actor actor.Actor) *MessagesRestoreRequest {
	return &MessagesRestoreRequest{*NewMethodBaseRequest(a, actor, "messages.restore")}
}

// Exec executes the request and unmarshals the response into MessagesRestoreResponse
func (r *MessagesRestoreRequest) Exec(ctx context.Context) (response response.MessagesRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSearchRequest defines the request for messages.search
//
// Returns a list of found messages based on the search query.
// Doc: https://dev.vk.com/method/messages.search
type MessagesSearchRequest struct {
	BaseRequest
}

// NewMessagesSearchRequest creates a new request for messages.search
func NewMessagesSearchRequest(a *api.API, actor actor.Actor) *MessagesSearchRequest {
	return &MessagesSearchRequest{*NewMethodBaseRequest(a, actor, "messages.search")}
}

// Exec executes the request and unmarshals the response into MessagesSearchResponse
func (r *MessagesSearchRequest) Exec(ctx context.Context) (response response.MessagesSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSearchConversationsRequest defines the request for messages.searchConversations
//
// Allows searching for conversations.
// Doc: https://dev.vk.com/method/messages.searchConversations
type MessagesSearchConversationsRequest struct {
	BaseRequest
}

// NewMessagesSearchConversationsRequest creates a new request for messages.searchConversations
func NewMessagesSearchConversationsRequest(a *api.API, actor actor.Actor) *MessagesSearchConversationsRequest {
	return &MessagesSearchConversationsRequest{*NewMethodBaseRequest(a, actor, "messages.searchConversations")}
}

// Exec executes the request and unmarshals the response into MessagesSearchConversationsResponse
func (r *MessagesSearchConversationsRequest) Exec(ctx context.Context) (response response.MessagesSearchConversationsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSendRequest defines the request for messages.send
//
// Sends a message.
// Doc: https://dev.vk.com/method/messages.send
type MessagesSendRequest struct {
	BaseRequest
}

// NewMessagesSendRequest creates a new request for messages.send
func NewMessagesSendRequest(a *api.API, actor actor.Actor) *MessagesSendRequest {
	return &MessagesSendRequest{*NewMethodBaseRequest(a, actor, "messages.send")}
}

// Exec executes the request and unmarshals the response into MessagesSendResponse
func (r *MessagesSendRequest) Exec(ctx context.Context) (response response.MessagesSendResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSendMessageEventAnswerRequest defines the request for messages.sendMessageEventAnswer
//
// Sends an event with an action that occurs when a callback button is pressed.
// Doc: https://dev.vk.com/method/messages.sendMessageEventAnswer
type MessagesSendMessageEventAnswerRequest struct {
	BaseRequest
}

// NewMessagesSendMessageEventAnswerRequest creates a new request for messages.sendMessageEventAnswer
func NewMessagesSendMessageEventAnswerRequest(a *api.API, actor actor.Actor) *MessagesSendMessageEventAnswerRequest {
	return &MessagesSendMessageEventAnswerRequest{*NewMethodBaseRequest(a, actor, "messages.sendMessageEventAnswer")}
}

// Exec executes the request and unmarshals the response into MessagesSendMessageEventAnswerResponse
func (r *MessagesSendMessageEventAnswerRequest) Exec(ctx context.Context) (response response.MessagesSendMessageEventAnswerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSendReactionRequest defines the request for messages.sendReaction
//
// Sets a reaction to a message.
// Doc: https://dev.vk.com/method/messages.sendReaction
type MessagesSendReactionRequest struct {
	BaseRequest
}

// NewMessagesSendReactionRequest creates a new request for messages.sendReaction
func NewMessagesSendReactionRequest(a *api.API, actor actor.Actor) *MessagesSendReactionRequest {
	return &MessagesSendReactionRequest{*NewMethodBaseRequest(a, actor, "messages.sendReaction")}
}

// Exec executes the request and unmarshals the response into MessagesSendReactionResponse
func (r *MessagesSendReactionRequest) Exec(ctx context.Context) (response response.MessagesSendReactionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSetActivityRequest defines the request for messages.setActivity
//
// Changes the typing status of the user in a dialogue.
// Doc: https://dev.vk.com/method/messages.setActivity
type MessagesSetActivityRequest struct {
	BaseRequest
}

// NewMessagesSetActivityRequest creates a new request for messages.setActivity
func NewMessagesSetActivityRequest(a *api.API, actor actor.Actor) *MessagesSetActivityRequest {
	return &MessagesSetActivityRequest{*NewMethodBaseRequest(a, actor, "messages.setActivity")}
}

// Exec executes the request and unmarshals the response into MessagesSetActivityResponse
func (r *MessagesSetActivityRequest) Exec(ctx context.Context) (response response.MessagesSetActivityResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesSetChatPhotoRequest defines the request for messages.setChatPhoto
//
// The method saves the cover of the conversation after it has been successfully uploaded to the server.
// Doc: https://dev.vk.com/method/messages.setChatPhoto
type MessagesSetChatPhotoRequest struct {
	BaseRequest
}

// NewMessagesSetChatPhotoRequest creates a new request for messages.setChatPhoto
func NewMessagesSetChatPhotoRequest(a *api.API, actor actor.Actor) *MessagesSetChatPhotoRequest {
	return &MessagesSetChatPhotoRequest{*NewMethodBaseRequest(a, actor, "messages.setChatPhoto")}
}

// Exec executes the request and unmarshals the response into MessagesSetChatPhotoResponse
func (r *MessagesSetChatPhotoRequest) Exec(ctx context.Context) (response response.MessagesSetChatPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MessagesUnpinRequest defines the request for messages.unpin
//
// Unpins a message.
// Doc: https://dev.vk.com/method/messages.unpin
type MessagesUnpinRequest struct {
	BaseRequest
}

// NewMessagesUnpinRequest creates a new request for messages.unpin
func NewMessagesUnpinRequest(a *api.API, actor actor.Actor) *MessagesUnpinRequest {
	return &MessagesUnpinRequest{*NewMethodBaseRequest(a, actor, "messages.unpin")}
}

// Exec executes the request and unmarshals the response into MessagesUnpinResponse
func (r *MessagesUnpinRequest) Exec(ctx context.Context) (response response.MessagesUnpinResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
