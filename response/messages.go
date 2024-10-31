package response

import (
	"go-vk-sdk/errors"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/method/messages

// Methods for working with personal messages.
//  To instantly receive incoming messages, use the LongPoll server. https://dev.vk.com/ru/api/user-long-poll/getting-started
//  Please note: access to working with section methods with a user key is limited.
//  Information about the Messages api limitation is in the Roadmap. https://dev.vk.com/ru/reference/roadmap#%D0%9E%D0%B3%D1%80%D0%B0%D0%BD%D0%B8%D1%87%D0%B5%D0%BD%D0%B8%D0%B5%20Messages%20API
//  Please note: methods for working with calls have been moved to the new calls section. Old calling methods from the messages section have been marked deprecated and may be removed in future versions of the api.

type MessagesAddChatUserResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesAllowMessagesFromGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesCreateChatResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesDeleteResponse struct {
	BaseResponse
	Response map[string]int `json:"response"`
}

type MessagesDeleteChatPhotoResponse struct {
	BaseResponse
	Response struct {
		MessageID int                  `json:"message_id"`
		Chat      objects.MessagesChat `json:"chat"`
	} `json:"response"`
}

type MessagesDeleteConversationResponse struct {
	BaseResponse
	Response struct {
		LastDeletedID int `json:"last_deleted_id"` // Id of the last message, that was deleted
	} `json:"response"`
}

type MessagesDeleteReactionResponse struct {
	BaseResponse
	Response objects.BoolInt `json:"response"`
}

type MessagesDenyMessagesFromGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesEditChatResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesForceCallFinishResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesGetByConversationMessageIDResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Message `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetByIDResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Message `json:"items"`
	} `json:"response"`
}

type MessagesGetByIDUsersAndGroups struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Message `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetChatResponse struct {
	BaseResponse
	Response objects.MessagesChat `json:"response"`
}

type MessagesGetChatChatIDsResponse struct {
	BaseResponse
	Response []objects.MessagesChat `json:"response"`
}

type MessagesGetChatPreviewResponse struct {
	BaseResponse
	Response struct {
		Preview objects.MessagesChatPreview `json:"preview"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetConversationMembersResponse struct {
	BaseResponse
	Response struct {
		Items            []objects.MessagesConversationMember `json:"items"`
		Count            int                                  `json:"count"`
		ChatRestrictions struct {
			OnlyAdminsInvite   objects.BoolInt `json:"only_admins_invite"`
			OnlyAdminsEditPin  objects.BoolInt `json:"only_admins_edit_pin"`
			OnlyAdminsEditInfo objects.BoolInt `json:"only_admins_edit_info"`
			AdminsPromoteUsers objects.BoolInt `json:"admins_promote_users"`
		} `json:"chat_restrictions"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetConversationsResponse struct {
	BaseResponse
	Response struct {
		Count       int                                       `json:"count"`
		Items       []objects.MessagesConversationWithMessage `json:"items"`
		UnreadCount int                                       `json:"unread_count"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetConversationsByIDResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Items []objects.MessagesConversation `json:"items"`
	} `json:"response"`
}

type MessagesGetConversationsByIDUsersAndGroups struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Items []objects.MessagesConversation `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetHistoryResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Message `json:"items"`
		objects.UsersAndGroups
		Conversations []objects.MessagesConversation `json:"conversations,omitempty"`
	} `json:"response"`
}

type MessagesGetHistoryAttachmentsResponse struct {
	BaseResponse
	Response struct {
		Items    []objects.MessagesHistoryAttachment `json:"items"`
		NextFrom string                              `json:"next_from"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetImportantMessagesResponse struct {
	BaseResponse
	Response struct {
		Messages struct {
			Count int               `json:"count"`
			Items []objects.Message `json:"items"`
		} `json:"messages"`
		Conversations []objects.MessagesConversation `json:"conversations"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesGetIntentUsersResponse struct {
	BaseResponse
	Response struct {
		Count    int               `json:"count"`
		Items    []int             `json:"items"`
		Profiles []objects.Message `json:"profiles,omitempty"`
	} `json:"response"`
}

type MessagesGetInviteLinkResponse struct {
	BaseResponse
	Response struct {
		Link string `json:"link"`
	} `json:"response"`
}

type MessagesGetLastActivityResponse struct {
	BaseResponse
	Response objects.MessagesLastActivity `json:"response"`
}

type MessagesGetLongPollHistoryResponse struct {
	BaseResponse
	Response struct {
		History  [][]int             `json:"history"`
		Groups   []objects.GroupFull `json:"groups"`
		Messages struct {
			Count int               `json:"count"`
			Items []objects.Message `json:"items"`
		} `json:"messages"`
		Profiles []objects.UserFull `json:"profiles"`
		// Chats struct {} `json:"chats"`
		NewPTS        int                            `json:"new_pts"`
		FromPTS       int                            `json:"from_pts"`
		More          objects.BoolInt                `json:"chats"`
		Conversations []objects.MessagesConversation `json:"conversations"`
	} `json:"response"`
}

type MessagesGetLongPollServerResponse struct {
	BaseResponse
	Response objects.MessagesLongPollParams `json:"response"`
}

type MessagesGetMessagesReactionsResponse struct {
	BaseResponse
	Response struct {
		Groups   []objects.GroupFull                            `json:"groups"`
		Profiles []objects.UserFull                             `json:"profiles"`
		Items    []objects.MessagesReactionCountersResponseItem `json:"items"`
	} `json:"response"`
}

type MessagesGetReactedPeersResponse struct {
	BaseResponse
	Response struct {
		Count     int                                           `json:"count"`
		Counters  []objects.MessagesReactionCounterResponseItem `json:"counters"`
		Groups    []objects.GroupFull                           `json:"groups"`
		Profiles  []objects.UserFull                            `json:"profiles"`
		Reactions []objects.MessagesReactionResponseItem        `json:"reactions"`
	} `json:"response"`
}

type MessagesGetReactionsAssetsResponse struct {
	BaseResponse
	Response struct {
		Version        int                                 `json:"version"`
		Assets         []objects.MessagesReactionAssetItem `json:"assets"`
		OverrideAssets []objects.MessagesReactionAssetItem `json:"override_assets"`
	} `json:"response"`
}

type MessagesMarkReactionsAsReadResponse struct {
	BaseResponse
	Response objects.BoolInt `json:"response"`
}

type MessagesSendReactionResponse struct {
	BaseResponse
	Response objects.BoolInt `json:"response"`
}

type MessagesIsMessagesFromGroupAllowedResponse struct {
	BaseResponse
	Response struct {
		IsAllowed objects.BoolInt `json:"is_allowed"`
	} `json:"response"`
}

type MessagesJoinChatByInviteLinkResponse struct {
	BaseResponse
	Response struct {
		ChatID int `json:"chat_id"`
	} `json:"response"`
}

type MessagesMarkAsAnsweredConversationResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesMarkAsImportantResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type MessagesMarkAsImportantConversationResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesMarkAsReadResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesPinResponse struct {
	BaseResponse
	Response objects.Message `json:"response"`
}

type MessagesRemoveChatUserResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesRestoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesSearchResponse struct {
	BaseResponse
	Response struct {
		Count int               `json:"count"`
		Items []objects.Message `json:"items"`
		objects.UsersAndGroups
		Conversations []objects.MessagesConversation `json:"conversations,omitempty"`
	} `json:"response"`
}

type MessagesSearchConversationsResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Items []objects.MessagesConversation `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MessagesSendResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesSendPeerIDsResponse struct {
	BaseResponse
	Response []struct {
		PeerID                int             `json:"peer_id"`
		MessageID             int             `json:"message_id"`
		ConversationMessageID int             `json:"conversation_message_id"`
		Error                 errors.APIError `json:"error"`
	} `json:"response"`
}

type MessagesSendMessageEventAnswerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesSendStickerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesSetActivityResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MessagesSetChatPhotoResponse struct {
	BaseResponse
	Response struct {
		MessageID int                  `json:"message_id"`
		Chat      objects.MessagesChat `json:"chat"`
	} `json:"response"`
}

type MessagesStartCallResponse struct {
	BaseResponse
	Response struct {
		JoinLink string `json:"join_link"`
		CallID   string `json:"call_id"`
	} `json:"response"`
}

type MessagesUnpinResponse struct {
	BaseResponse
	Response int `json:"response"`
}
