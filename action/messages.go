package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Messages Doc: https://dev.vk.com/ru/method/messages
type Messages struct {
	BaseAction
}

// AddChatUser Doc: https://dev.vk.com/ru/method/messages.addChatUser
func (m *Messages) AddChatUser(actor actor.Actor) *request.MessagesAddChatUserRequest {
	return request.NewMessagesAddChatUserRequest(m.api, actor)
}

// AllowMessagesFromGroup Doc: https://dev.vk.com/ru/method/messages.allowMessagesFromGroup
func (m *Messages) AllowMessagesFromGroup(actor actor.Actor) *request.MessagesAllowMessagesFromGroupRequest {
	return request.NewMessagesAllowMessagesFromGroupRequest(m.api, actor)
}

// CreateChat Doc: https://dev.vk.com/ru/method/messages.createChat
func (m *Messages) CreateChat(actor actor.Actor) *request.MessagesCreateChatRequest {
	return request.NewMessagesCreateChatRequest(m.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/messages.delete
func (m *Messages) Delete(actor actor.Actor) *request.MessagesDeleteRequest {
	return request.NewMessagesDeleteRequest(m.api, actor)
}

// DeleteChatPhoto Doc: https://dev.vk.com/ru/method/messages.deleteChatPhoto
func (m *Messages) DeleteChatPhoto(actor actor.Actor) *request.MessagesDeleteChatPhotoRequest {
	return request.NewMessagesDeleteChatPhotoRequest(m.api, actor)
}

// DeleteConversation Doc: https://dev.vk.com/ru/method/messages.deleteConversation
func (m *Messages) DeleteConversation(actor actor.Actor) *request.MessagesDeleteConversationRequest {
	return request.NewMessagesDeleteConversationRequest(m.api, actor)
}

// DeleteReaction Doc: https://dev.vk.com/ru/method/messages.deleteReaction
func (m *Messages) DeleteReaction(actor actor.Actor) *request.MessagesDeleteReactionRequest {
	return request.NewMessagesDeleteReactionRequest(m.api, actor)
}

// DenyMessagesFromGroup Doc: https://dev.vk.com/ru/method/messages.denyMessagesFromGroup
func (m *Messages) DenyMessagesFromGroup(actor actor.Actor) *request.MessagesDenyMessagesFromGroupRequest {
	return request.NewMessagesDenyMessagesFromGroupRequest(m.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/messages.edit
func (m *Messages) Edit(actor actor.Actor) *request.MessagesEditRequest {
	return request.NewMessagesEditRequest(m.api, actor)
}

// EditChat Doc: https://dev.vk.com/ru/method/messages.editChat
func (m *Messages) EditChat(actor actor.Actor) *request.MessagesEditChatRequest {
	return request.NewMessagesEditChatRequest(m.api, actor)
}

// GetByConversationMessageId Doc: https://dev.vk.com/ru/method/messages.getByConversationMessageId
func (m *Messages) GetByConversationMessageId(actor actor.Actor) *request.MessagesGetByConversationMessageIdRequest {
	return request.NewMessagesGetByConversationMessageIdRequest(m.api, actor)
}

// GetById Doc: https://dev.vk.com/ru/method/messages.getById
func (m *Messages) GetById(actor actor.Actor) *request.MessagesGetByIdRequest {
	return request.NewMessagesGetByIdRequest(m.api, actor)
}

// GetChat Doc: https://dev.vk.com/ru/method/messages.getChat
func (m *Messages) GetChat(actor actor.Actor) *request.MessagesGetChatRequest {
	return request.NewMessagesGetChatRequest(m.api, actor)
}

// GetChatPreview Doc: https://dev.vk.com/ru/method/messages.getChatPreview
func (m *Messages) GetChatPreview(actor actor.Actor) *request.MessagesGetChatPreviewRequest {
	return request.NewMessagesGetChatPreviewRequest(m.api, actor)
}

// GetConversationMembers Doc: https://dev.vk.com/ru/method/messages.getConversationMembers
func (m *Messages) GetConversationMembers(actor actor.Actor) *request.MessagesGetConversationMembersRequest {
	return request.NewMessagesGetConversationMembersRequest(m.api, actor)
}

// GetConversations Doc: https://dev.vk.com/ru/method/messages.getConversations
func (m *Messages) GetConversations(actor actor.Actor) *request.MessagesGetConversationsRequest {
	return request.NewMessagesGetConversationsRequest(m.api, actor)
}

// GetConversationsById Doc: https://dev.vk.com/ru/method/messages.getConversationsById
func (m *Messages) GetConversationsById(actor actor.Actor) *request.MessagesGetConversationsByIdRequest {
	return request.NewMessagesGetConversationsByIdRequest(m.api, actor)
}

// GetHistory Doc: https://dev.vk.com/ru/method/messages.getHistory
func (m *Messages) GetHistory(actor actor.Actor) *request.MessagesGetHistoryRequest {
	return request.NewMessagesGetHistoryRequest(m.api, actor)
}

// GetHistoryAttachments Doc: https://dev.vk.com/ru/method/messages.getHistoryAttachments
func (m *Messages) GetHistoryAttachments(actor actor.Actor) *request.MessagesGetHistoryAttachmentsRequest {
	return request.NewMessagesGetHistoryAttachmentsRequest(m.api, actor)
}

// GetImportantMessages Doc: https://dev.vk.com/ru/method/messages.getImportantMessages
func (m *Messages) GetImportantMessages(actor actor.Actor) *request.MessagesGetImportantMessagesRequest {
	return request.NewMessagesGetImportantMessagesRequest(m.api, actor)
}

// GetIntentUsers Doc: https://dev.vk.com/ru/method/messages.getIntentUsers
func (m *Messages) GetIntentUsers(actor actor.Actor) *request.MessagesGetIntentUsersRequest {
	return request.NewMessagesGetIntentUsersRequest(m.api, actor)
}

// GetInviteLink Doc: https://dev.vk.com/ru/method/messages.getInviteLink
func (m *Messages) GetInviteLink(actor actor.Actor) *request.MessagesGetInviteLinkRequest {
	return request.NewMessagesGetInviteLinkRequest(m.api, actor)
}

// GetLastActivity Doc: https://dev.vk.com/ru/method/messages.getLastActivity
func (m *Messages) GetLastActivity(actor actor.Actor) *request.MessagesGetLastActivityRequest {
	return request.NewMessagesGetLastActivityRequest(m.api, actor)
}

// GetLongPollHistory Doc: https://dev.vk.com/ru/method/messages.getLongPollHistory
func (m *Messages) GetLongPollHistory(actor actor.Actor) *request.MessagesGetLongPollHistoryRequest {
	return request.NewMessagesGetLongPollHistoryRequest(m.api, actor)
}

// GetLongPollServer Doc: https://dev.vk.com/ru/method/messages.getLongPollServer
func (m *Messages) GetLongPollServer(actor actor.Actor) *request.MessagesGetLongPollServerRequest {
	return request.NewMessagesGetLongPollServerRequest(m.api, actor)
}

// GetMessagesReactions Doc: https://dev.vk.com/ru/method/messages.getMessagesReactions
func (m *Messages) GetMessagesReactions(actor actor.Actor) *request.MessagesGetMessagesReactionsRequest {
	return request.NewMessagesGetMessagesReactionsRequest(m.api, actor)
}

// GetReactedPeers Doc: https://dev.vk.com/ru/method/messages.getReactedPeers
func (m *Messages) GetReactedPeers(actor actor.Actor) *request.MessagesGetReactedPeersRequest {
	return request.NewMessagesGetReactedPeersRequest(m.api, actor)
}

// GetReactionsAssets Doc: https://dev.vk.com/ru/method/messages.getReactionsAssets
func (m *Messages) GetReactionsAssets(actor actor.Actor) *request.MessagesGetReactionsAssetsRequest {
	return request.NewMessagesGetReactionsAssetsRequest(m.api, actor)
}

// IsMessagesFromGroupAllowed Doc: https://dev.vk.com/ru/method/messages.isMessagesFromGroupAllowed
func (m *Messages) IsMessagesFromGroupAllowed(actor actor.Actor) *request.MessagesIsMessagesFromGroupAllowedRequest {
	return request.NewMessagesIsMessagesFromGroupAllowedRequest(m.api, actor)
}

// JoinChatByInviteLink Doc: https://dev.vk.com/ru/method/messages.joinChatByInviteLink
func (m *Messages) JoinChatByInviteLink(actor actor.Actor) *request.MessagesJoinChatByInviteLinkRequest {
	return request.NewMessagesJoinChatByInviteLinkRequest(m.api, actor)
}

// MarkAsAnsweredConversation Doc: https://dev.vk.com/ru/method/messages.markAsAnsweredConversation
func (m *Messages) MarkAsAnsweredConversation(actor actor.Actor) *request.MessagesMarkAsAnsweredConversationRequest {
	return request.NewMessagesMarkAsAnsweredConversationRequest(m.api, actor)
}

// MarkAsImportant Doc: https://dev.vk.com/ru/method/messages.markAsImportant
func (m *Messages) MarkAsImportant(actor actor.Actor) *request.MessagesMarkAsImportantRequest {
	return request.NewMessagesMarkAsImportantRequest(m.api, actor)
}

// MarkAsImportantConversation Doc: https://dev.vk.com/ru/method/messages.markAsImportantConversation
func (m *Messages) MarkAsImportantConversation(actor actor.Actor) *request.MessagesMarkAsImportantConversationRequest {
	return request.NewMessagesMarkAsImportantConversationRequest(m.api, actor)
}

// MarkAsRead Doc: https://dev.vk.com/ru/method/messages.markAsRead
func (m *Messages) MarkAsRead(actor actor.Actor) *request.MessagesMarkAsReadRequest {
	return request.NewMessagesMarkAsReadRequest(m.api, actor)
}

// MarkReactionsAsRead Doc: https://dev.vk.com/ru/method/messages.markReactionsAsRead
func (m *Messages) MarkReactionsAsRead(actor actor.Actor) *request.MessagesMarkReactionsAsReadRequest {
	return request.NewMessagesMarkReactionsAsReadRequest(m.api, actor)
}

// Pin Doc: https://dev.vk.com/ru/method/messages.pin
func (m *Messages) Pin(actor actor.Actor) *request.MessagesPinRequest {
	return request.NewMessagesPinRequest(m.api, actor)
}

// RemoveChatUser Doc: https://dev.vk.com/ru/method/messages.removeChatUser
func (m *Messages) RemoveChatUser(actor actor.Actor) *request.MessagesRemoveChatUserRequest {
	return request.NewMessagesRemoveChatUserRequest(m.api, actor)
}

// Restore Doc: https://dev.vk.com/ru/method/messages.restore
func (m *Messages) Restore(actor actor.Actor) *request.MessagesRestoreRequest {
	return request.NewMessagesRestoreRequest(m.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/messages.search
func (m *Messages) Search(actor actor.Actor) *request.MessagesSearchRequest {
	return request.NewMessagesSearchRequest(m.api, actor)
}

// SearchConversations Doc: https://dev.vk.com/ru/method/messages.searchConversations
func (m *Messages) SearchConversations(actor actor.Actor) *request.MessagesSearchConversationsRequest {
	return request.NewMessagesSearchConversationsRequest(m.api, actor)
}

// Send Doc: https://dev.vk.com/ru/method/messages.send
func (m *Messages) Send(actor actor.Actor) *request.MessagesSendRequest {
	return request.NewMessagesSendRequest(m.api, actor)
}

// SendMessageEventAnswer Doc: https://dev.vk.com/ru/method/messages.sendMessageEventAnswer
func (m *Messages) SendMessageEventAnswer(actor actor.Actor) *request.MessagesSendMessageEventAnswerRequest {
	return request.NewMessagesSendMessageEventAnswerRequest(m.api, actor)
}

// SendReaction Doc: https://dev.vk.com/ru/method/messages.sendReaction
func (m *Messages) SendReaction(actor actor.Actor) *request.MessagesSendReactionRequest {
	return request.NewMessagesSendReactionRequest(m.api, actor)
}

// SetActivity Doc: https://dev.vk.com/ru/method/messages.setActivity
func (m *Messages) SetActivity(actor actor.Actor) *request.MessagesSetActivityRequest {
	return request.NewMessagesSetActivityRequest(m.api, actor)
}

// SetChatPhoto Doc: https://dev.vk.com/ru/method/messages.setChatPhoto
func (m *Messages) SetChatPhoto(actor actor.Actor) *request.MessagesSetChatPhotoRequest {
	return request.NewMessagesSetChatPhotoRequest(m.api, actor)
}

// Unpin Doc: https://dev.vk.com/ru/method/messages.unpin
func (m *Messages) Unpin(actor actor.Actor) *request.MessagesUnpinRequest {
	return request.NewMessagesUnpinRequest(m.api, actor)
}
