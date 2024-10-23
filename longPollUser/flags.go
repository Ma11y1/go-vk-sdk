package longPollUser

// Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started

type FailedType int

const (
	FailedTypeOutdatedStory    FailedType = 1
	FailedTypeExpiredKey       FailedType = 2
	FailedTypeOutdatedUserInfo FailedType = 3
	FailedTypeInvalidVersion   FailedType = 4
)

type ExtraOptionsMode int

const (
	ExtraOptionsModeReceiveAttachments ExtraOptionsMode = 1 << 1 // receive attachments
	ExtraOptionsModeExtendedEvents     ExtraOptionsMode = 1 << 3 // receive more events
	ExtraOptionsModeReturnPts          ExtraOptionsMode = 1 << 5 // receive pts (used in messages.getLongPollHistory)
	ExtraOptionsModeCode8ExtraFields   ExtraOptionsMode = 1 << 6 // extra fields in event type 8(friend become online)
	ExtraOptionsModeReturnRandomID     ExtraOptionsMode = 1 << 7 // return random_id field
)

type ExtraChatField int

const (
	ExtraChatFieldNameChange     ExtraChatField = iota + 1 // The name of the conversation has changed
	ExtraChatFieldCoverChange                              // The cover of the conversation has changed
	ExtraChatFieldAdminAssigned                            // New administrator has been appointed.
	ExtraChatFieldPinMessage                               // Message pin
	ExtraChatFieldUserCome                                 // LongPoll has joined the conversation
	ExtraChatFieldUserLeave                                // LongPoll left the conversation.
	ExtraChatFieldUserKicked                               // LongPoll was excluded from the conversation.
	ExtraChatFieldAdminDismissed                           // Administrator rights have been removed from the user
)

type MessageFlag int

func (b MessageFlag) Has(flag MessageFlag) bool { return b&flag != 0 }

// Each message has a flag, which is a value received by summing up any of the following parameters.
const (
	MessageFlagUnread       MessageFlag = 1 << iota // Message is unread
	MessageFlagOutbox                               // Message is outgoing
	MessageFlagReplied                              // Message was answered
	MessageFlagImportant                            // Message is marked as important
	MessageFlagChat                                 // Message sent via chat
	MessageFlagFriends                              // Message sent by a friend
	MessageFlagSpam                                 // Message marked as "Spam"
	MessageFlagDeleted                              // Message was deleted
	MessageFlagFixed                                // Message was user-checked for spam
	MessageFlagMedia                                // Message has media content
	MessageFlagHidden       MessageFlag = 1 << 16   // Greeting message from a community
	MessageFlagDeleteForAll MessageFlag = 1 << 17   // Message was deleted for all
	MessageFlagNotDelivered MessageFlag = 1 << 18   // Incoming message not delivered
)

type DialogFlag int

func (b DialogFlag) Has(flag DialogFlag) bool { return b&flag != 0 }

// Each dialog has flags, which are values received by summing up any of the
// following parameters.
const (
	DialogFlagImportant  DialogFlag = 1 << iota // Important dialog
	DialogFlagUnanswered                        // Dialog without a community reply
)

type PlatformType int

func (b PlatformType) Has(flag PlatformType) bool { return b&flag != 0 }

const (
	PlatformTypeMobile PlatformType = iota + 1
	PlatformTypeIPhone
	PlatformTypeIPad
	PlatformTypeAndroid
	PlatformTypeWPhone
	PlatformTypeWindows
	PlatformTypeWeb
)

type ChatAction string

const (
	ChatActionCreate       ChatAction = "chat_create"       // create chat
	ChatActionTitleUpdate  ChatAction = "chat_title_update" // change chat name
	ChatActionPhotoUpdate  ChatAction = "chat_photo_update" // change chat photo
	ChatActionInviteUser   ChatAction = "chat_invite_user"  // invite user to chat
	ChatActionKickUser     ChatAction = "chat_kick_user"    // kick out user from chat
	ChatActionPinMessage   ChatAction = "chat_pin_message"
	ChatActionUnpinMessage ChatAction = "chat_unpin_message"
)
