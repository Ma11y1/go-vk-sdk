package objects

import (
	"encoding/json"
	"fmt"
	"go-vk-sdk/constants"
)

type MessagesAudioMessage struct {
	AccessKey string `json:"access_key"` // Access key for the document
	ID        int    `json:"id"`         // Document ID
	OwnerID   int    `json:"owner_id"`   // Document owner ID
	Duration  int    `json:"duration"`   // Audio message duration in seconds
	LinkMp3   string `json:"link_mp3"`   // MP3 file url
	LinkOgg   string `json:"link_ogg"`   // OGG file url
	Waveform  []int  `json:"waveform"`   // Sound visualisation
}

func (doc MessagesAudioMessage) ToAttachment() string {
	return fmt.Sprintf("doc%d_%d", doc.OwnerID, doc.ID)
}

type MessagesGraffiti struct {
	AccessKey string `json:"access_key"` // Access key for the document
	ID        int    `json:"id"`         // Document ID
	OwnerID   int    `json:"owner_id"`   // Document owner ID
	URL       string `json:"url"`        // Graffiti URL
	Width     int    `json:"width"`      // Graffiti width
	Height    int    `json:"height"`     // Graffiti height
}

func (doc MessagesGraffiti) ToAttachment() string {
	return fmt.Sprintf("doc%d_%d", doc.OwnerID, doc.ID)
}

type Message struct {
	// Only for messages from community. Contains user ID of community admin,
	// who sent this message.
	AdminAuthorID int                         `json:"admin_author_id"`
	Action        MessagesMessageAction       `json:"action"`
	Attachments   []MessagesMessageAttachment `json:"attachments"`

	// Unique auto-incremented number for all messages with this peer.
	ConversationMessageID int `json:"conversation_message_id"`

	// Date when the message has been sent in Unixtime.
	Date int `json:"date"`

	// EventType author's ID.
	FromID int `json:"from_id"`

	// Forwarded messages.
	FwdMessages  []Message        `json:"fwd_Messages"`
	ReplyMessage *Message         `json:"reply_message"`
	Geo          MessageGeo       `json:"geo"`
	PinnedAt     int              `json:"pinned_at,omitempty"`
	ID           int              `json:"id"`        // EventType ID
	Deleted      BoolInt          `json:"deleted"`   // Is it an deleted message
	Important    BoolInt          `json:"important"` // Is it an important message
	IsHidden     BoolInt          `json:"is_hidden"`
	IsCropped    BoolInt          `json:"is_cropped"`
	IsSilent     BoolInt          `json:"is_silent"`
	Out          BoolInt          `json:"out"` // Information whether the message is outcoming
	WasListened  BoolInt          `json:"was_listened,omitempty"`
	Keyboard     MessagesKeyboard `json:"keyboard"`
	Template     MessagesTemplate `json:"template"`
	Payload      string           `json:"payload"`
	PeerID       int              `json:"peer_id"` // Peer ID

	// ID used for sending messages. It returned only for outgoing messages.
	RandomID     int    `json:"random_id"`
	Ref          string `json:"ref"`
	RefSource    string `json:"ref_source"`
	Text         string `json:"text"`          // EventType text
	UpdateTime   int    `json:"update_time"`   // Date when the message has been updated in Unixtime
	MembersCount int    `json:"members_count"` // Members number
	ExpireTTL    int    `json:"expire_ttl"`
	MessageTag   string `json:"message_tag"` // for https://notify.mail.ru/
}

type MessagesBasePayload struct {
	ButtonType string `json:"button_type,omitempty"`
	Command    string `json:"command,omitempty"`
	Payload    string `json:"payload,omitempty"`
}

type MessagesKeyboard struct {
	AuthorID int                        `json:"author_id,omitempty"` // Community or bot, which set this keyboard
	Buttons  [][]MessagesKeyboardButton `json:"buttons"`
	OneTime  BoolInt                    `json:"one_time,omitempty"` // Should this keyboard disappear on first use
	Inline   BoolInt                    `json:"inline,omitempty"`
}

func NewMessagesKeyboard(oneTime BoolInt) *MessagesKeyboard {
	return &MessagesKeyboard{
		Buttons: [][]MessagesKeyboardButton{},
		OneTime: oneTime,
	}
}

func NewMessagesKeyboardInline() *MessagesKeyboard {
	return &MessagesKeyboard{
		Buttons: [][]MessagesKeyboardButton{},
		Inline:  true,
	}
}

func (keyboard *MessagesKeyboard) AddRow() *MessagesKeyboard {
	if len(keyboard.Buttons) == 0 {
		keyboard.Buttons = make([][]MessagesKeyboardButton, 1)
	} else {
		row := make([]MessagesKeyboardButton, 0)
		keyboard.Buttons = append(keyboard.Buttons, row)
	}

	return keyboard
}

func (keyboard *MessagesKeyboard) AddTextButton(label string, payload interface{}, color string) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonText,
			Label:   label,
			Payload: string(b),
		},
		Color: color,
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard *MessagesKeyboard) AddOpenLinkButton(link, label string, payload interface{}) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonOpenLink,
			Payload: string(b),
			Label:   label,
			Link:    link,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard *MessagesKeyboard) AddLocationButton(payload interface{}) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonLocation,
			Payload: string(b),
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard *MessagesKeyboard) AddVKPayButton(payload interface{}, hash string) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonVKPay,
			Payload: string(b),
			Hash:    hash,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard *MessagesKeyboard) AddVKAppsButton(
	appID, ownerID int,
	payload interface{},
	label, hash string,
) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonVKApp,
			AppID:   appID,
			OwnerID: ownerID,
			Payload: string(b),
			Label:   label,
			Hash:    hash,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard *MessagesKeyboard) AddCallbackButton(label string, payload interface{}, color string) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    constants.ButtonCallback,
			Label:   label,
			Payload: string(b),
		},
		Color: color,
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)

	return keyboard
}

func (keyboard MessagesKeyboard) ToJSON() string {
	b, _ := json.Marshal(keyboard)
	return string(b)
}

type MessagesKeyboardButton struct {
	Action MessagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color,omitempty"` // Button color
}

type MessagesKeyboardButtonAction struct {
	AppID   int    `json:"app_id,omitempty"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash,omitempty"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label,omitempty"`    // Label for button
	OwnerID int    `json:"owner_id,omitempty"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload,omitempty"`  // Additional data sent along with message for developer convenience
	Type    string `json:"type"`               // Button type
	Link    string `json:"link,omitempty"`     // Link url
}

type MessagesEventDataShowSnackbar struct {
	Text string `json:"text,omitempty"`
}

type MessagesEventDataOpenLink struct {
	Link string `json:"link,omitempty"`
}

type MessagesEventDataOpenApp struct {
	AppID   int    `json:"app_id,omitempty"`
	OwnerID int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type MessagesEventData struct {
	Type string `json:"type"`
	MessagesEventDataShowSnackbar
	MessagesEventDataOpenLink
	MessagesEventDataOpenApp
}

// NewMessagesEventDataShowSnackbar show disappearing message.
//
// Contains the field text - the text you want to print
// (maximum 90 characters). Snackbar is shown for 10 seconds and automatically
// hides, while the user has the ability to flick it off the screen.
func NewMessagesEventDataShowSnackbar(text string) *MessagesEventData {
	return &MessagesEventData{
		Type: "show_snackbar",
		MessagesEventDataShowSnackbar: MessagesEventDataShowSnackbar{
			Text: text,
		},
	}
}

// NewMessagesEventDataOpenLink open the link. Click on the specified address.
func NewMessagesEventDataOpenLink(link string) *MessagesEventData {
	return &MessagesEventData{
		Type: "open_link",
		MessagesEventDataOpenLink: MessagesEventDataOpenLink{
			Link: link,
		},
	}
}

// NewMessagesEventDataOpenApp open the link. Click on the specified address.
func NewMessagesEventDataOpenApp(appID, ownerID int, hash string) *MessagesEventData {
	return &MessagesEventData{
		Type: "open_app",
		MessagesEventDataOpenApp: MessagesEventDataOpenApp{
			AppID:   appID,
			OwnerID: ownerID,
			Hash:    hash,
		},
	}
}

func (eventData MessagesEventData) ToJSON() string {
	b, _ := json.Marshal(eventData)
	return string(b)
}

// MessagesTemplate struct.
//
// https://dev.vk.com/ru/api/bots/development/messages
type MessagesTemplate struct {
	Type     string                    `json:"type"`
	Elements []MessagesTemplateElement `json:"elements"`
}

func (template MessagesTemplate) ToJSON() string {
	b, _ := json.Marshal(template)
	return string(b)
}

type MessagesTemplateElement struct {
	MessagesTemplateElementCarousel
}

type MessagesTemplateElementCarousel struct {
	Title       string                                `json:"title,omitempty"`
	Action      MessagesTemplateElementCarouselAction `json:"action,omitempty"`
	Description string                                `json:"description,omitempty"`
	Photo       *Photo                                `json:"photo,omitempty"`    // Only read
	PhotoID     string                                `json:"photo_id,omitempty"` // Only for send
	Buttons     []MessagesKeyboardButton              `json:"buttons,omitempty"`
}

type MessagesTemplateElementCarouselAction struct {
	Type string `json:"type"`
	Link string `json:"link,omitempty"`
}

type MessageContentSourceMessage struct {
	OwnerID               int `json:"owner_id,omitempty"`
	PeerID                int `json:"peer_id,omitempty"`
	ConversationMessageID int `json:"conversation_message_id,omitempty"`
}

type MessageContentSourceURL struct {
	URL string `json:"url,omitempty"`
}

// MessageContentSource struct.
//
// https://dev.vk.com/ru/api/bots/development/messages
type MessageContentSource struct {
	Type                        string `json:"type"`
	MessageContentSourceMessage        // type message
	MessageContentSourceURL            // type url
}

func NewMessageContentSourceMessage(ownerID, peerID, conversationMessageID int) *MessageContentSource {
	return &MessageContentSource{
		Type: "message",
		MessageContentSourceMessage: MessageContentSourceMessage{
			OwnerID:               ownerID,
			PeerID:                peerID,
			ConversationMessageID: conversationMessageID,
		},
	}
}

func NewMessageContentSourceURL(u string) *MessageContentSource {
	return &MessageContentSource{
		Type: "url",
		MessageContentSourceURL: MessageContentSourceURL{
			URL: u,
		},
	}
}

func (contentSource MessageContentSource) ToJSON() string {
	b, _ := json.Marshal(contentSource)
	return string(b)
}

type MessagesChat struct {
	AdminID        int     `json:"admin_id"` // Chat creator ID
	ID             int     `json:"id"`       // Chat ID
	IsDefaultPhoto BoolInt `json:"is_default_photo"`
	IsGroupChannel BoolInt `json:"is_group_channel"`
	Photo100       string  `json:"photo_100"` // url of the preview image with 100 px in width
	Photo200       string  `json:"photo_200"` // url of the preview image with 200 px in width
	Photo50        string  `json:"photo_50"`  // url of the preview image with 50 px in width
	Title          string  `json:"title"`     // Chat title
	Type           string  `json:"type"`      // Chat type
	Users          []int   `json:"users"`
	MembersCount   int     `json:"members_count"`
}

type MessagesChatPreview struct {
	AdminID      int                              `json:"admin_id"`
	MembersCount int                              `json:"members_count"`
	Members      []int                            `json:"members"`
	Title        string                           `json:"title"`
	Photo        MessagesChatSettingsPhoto        `json:"photo"`
	LocalID      int                              `json:"local_id"`
	Joined       bool                             `json:"joined"`
	ChatSettings MessagesConversationChatSettings `json:"chat_settings"`
	IsMember     BoolInt                          `json:"is_member"`
}

type MessagesChatPushSettings struct {
	DisabledUntil int     `json:"disabled_until"` // Time until that notifications are disabled
	Sound         BoolInt `json:"sound"`          // Information whether the sound is on
}

type MessagesChatSettingsPhoto struct {
	Photo100           string  `json:"photo_100"`
	Photo200           string  `json:"photo_200"`
	Photo50            string  `json:"photo_50"`
	IsDefaultPhoto     BoolInt `json:"is_default_photo"`
	IsDefaultCallPhoto bool    `json:"is_default_call_photo"`
}

type MessagesConversation struct {
	CanWrite                  MessagesConversationCanWrite     `json:"can_write"`
	ChatSettings              MessagesConversationChatSettings `json:"chat_settings"`
	InRead                    int                              `json:"in_read"`         // Last message user have read
	LastMessageID             int                              `json:"last_message_id"` // ID of the last message in conversation
	Mentions                  []int                            `json:"mentions"`        // IDs of messages with mentions
	MessageRequest            string                           `json:"message_request"`
	LastConversationMessageID int                              `json:"last_conversation_message_id"`
	InReadCMID                int                              `json:"in_read_cmid"`
	OutReadCMID               int                              `json:"out_read_cmid"`
	OutRead                   int                              `json:"out_read"`
	Peer                      MessagesConversationPeer         `json:"peer"`
	PushSettings              MessagesConversationPushSettings `json:"push_settings"`
	Important                 BoolInt                          `json:"important"`
	Unanswered                BoolInt                          `json:"unanswered"`
	IsMarkedUnread            BoolInt                          `json:"is_marked_unread"`
	CanSendMoney              BoolInt                          `json:"can_send_money"`
	CanReceiveMoney           BoolInt                          `json:"can_receive_money"`
	IsNew                     BoolInt                          `json:"is_new"`
	IsArchived                BoolInt                          `json:"is_archived"`
	UnreadCount               int                              `json:"unread_count"` // Unread messages number
	CurrentKeyboard           MessagesKeyboard                 `json:"current_keyboard"`
	SortID                    struct {
		MajorID int `json:"major_id"`
		MinorID int `json:"minor_id"`
	} `json:"sort_id"`
}

type MessagesConversationCanWrite struct {
	Allowed BoolInt `json:"allowed"`
	Reason  int     `json:"reason"`
}

type MessagesConversationChatSettings struct {
	MembersCount  int                       `json:"members_count"`
	FriendsCount  int                       `json:"friends_count"`
	Photo         MessagesChatSettingsPhoto `json:"photo"`
	PinnedMessage MessagesPinnedMessage     `json:"pinned_message"`
	State         string                    `json:"state"`
	Title         string                    `json:"title"`
	ActiveIDs     []int                     `json:"active_ids"`
	ACL           struct {
		CanInvite            BoolInt `json:"can_invite"`
		CanChangeInfo        BoolInt `json:"can_change_info"`
		CanChangePin         BoolInt `json:"can_change_pin"`
		CanPromoteUsers      BoolInt `json:"can_promote_users"`
		CanSeeInviteLink     BoolInt `json:"can_see_invite_link"`
		CanChangeInviteLink  BoolInt `json:"can_change_invite_link"`
		CanCopyChat          BoolInt `json:"can_copy_chat"`
		CanModerate          BoolInt `json:"can_moderate"`
		CanCall              BoolInt `json:"can_call"`
		CanUseMassMentions   BoolInt `json:"can_use_mass_mentions"`
		CanChangeServiceType BoolInt `json:"can_change_service_type"`
		CanChangeStyle       BoolInt `json:"can_change_style"`
	} `json:"acl"`
	IsGroupChannel   BoolInt                 `json:"is_group_channel"`
	IsDisappearing   BoolInt                 `json:"is_disappearing"`
	IsService        BoolInt                 `json:"is_service"`
	IsCreatedForCall BoolInt                 `json:"is_created_for_call"`
	OwnerID          int                     `json:"owner_id"`
	AdminIDs         []int                   `json:"admin_ids"`
	Permissions      MessagesChatPermissions `json:"permissions"`
}

type MessagesConversationMember struct {
	MemberID  int     `json:"member_id"`
	JoinDate  int     `json:"join_date"`
	InvitedBy int     `json:"invited_by"`
	IsOwner   BoolInt `json:"is_owner,omitempty"`
	IsAdmin   BoolInt `json:"is_admin,omitempty"`
	CanKick   BoolInt `json:"can_kick,omitempty"`
}

type MessagesChatPermissions struct {
	Invite          constants.MessagesChatPermission `json:"invite"`
	ChangeInfo      constants.MessagesChatPermission `json:"change_info"`
	ChangePin       constants.MessagesChatPermission `json:"change_pin"`
	UseMassMentions constants.MessagesChatPermission `json:"use_mass_mentions"`
	SeeInviteLink   constants.MessagesChatPermission `json:"see_invite_link"`
	Call            constants.MessagesChatPermission `json:"call"`
	ChangeAdmins    constants.MessagesChatPermission `json:"change_admins"`
	ChangeStyle     constants.MessagesChatPermission `json:"change_style"`
}

type MessagesConversationPeer struct {
	ID      int    `json:"id"`
	LocalID int    `json:"local_id"`
	Type    string `json:"type"`
}

type MessagesConversationPushSettings struct {
	DisabledUntil        int     `json:"disabled_until"`
	DisabledForever      BoolInt `json:"disabled_forever"`
	NoSound              BoolInt `json:"no_sound"`
	DisabledMentions     BoolInt `json:"disabled_mentions"`
	DisabledMassMentions BoolInt `json:"disabled_mass_mentions"`
}

type MessagesConversationWithMessage struct {
	Conversation MessagesConversation `json:"conversation"`
	LastMessage  Message              `json:"last_message"`
}

type MessagesDialog struct {
	Important  int     `json:"important"`
	InRead     int     `json:"in_read"`
	Message    Message `json:"message"`
	OutRead    int     `json:"out_read"`
	Unanswered int     `json:"unanswered"`
	Unread     int     `json:"unread"`
}

type MessagesHistoryAttachment struct {
	Attachment MessagesHistoryMessageAttachment `json:"attachment"`
	MessageID  int                              `json:"message_id"` // EventType ID
	FromID     int                              `json:"from_id"`
}

type MessagesHistoryMessageAttachment struct {
	Audio  Audio    `json:"audio"`
	Doc    Document `json:"doc"`
	Link   Link     `json:"link"`
	Market Link     `json:"market"`
	Photo  Photo    `json:"photo"`
	Share  Link     `json:"share"`
	Type   string   `json:"type"`
	Video  Video    `json:"video"`
	Wall   Link     `json:"wall"`
}

type MessagesLastActivity struct {
	Online BoolInt `json:"online"` // Information whether user is online
	Time   int     `json:"time"`   // Time when user was online in Unixtime
}

type MessagesLongPollParams struct {
	Key    string `json:"key"`    // Key
	Pts    int    `json:"pts"`    // Persistent timestamp
	Server string `json:"server"` // Server url
	Ts     int    `json:"ts"`     // Timestamp
}

type MessagesMessageAction struct {
	ConversationMessageID int `json:"conversation_message_id"` // Message ID

	// Email address for chat_invite_user or chat_kick_user actions.
	Email    string                     `json:"email"`
	MemberID int                        `json:"member_id"` // OAuthUser or email peer ID
	Message  string                     `json:"message"`   // Message body of related message
	Photo    MessagesMessageActionPhoto `json:"photo"`

	// New chat title for chat_create and chat_title_update actions.
	Text string `json:"text"`
	Type string `json:"type"`
}

type MessagesMessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // url of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // url of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // url of the preview image with 50px in width
}

type MessagesMessageAttachment struct {
	Audio             Audio           `json:"audio"`
	Doc               Document        `json:"doc"`
	Gift              GiftsLayout     `json:"gift"`
	Link              Link            `json:"link"`
	Market            MarketItem      `json:"market"`
	MarketMarketAlbum MarketAlbum     `json:"market_market_album"`
	Photo             Photo           `json:"photo"`
	Sticker           Sticker         `json:"sticker"`
	Type              string          `json:"type"`
	Video             Video           `json:"video"`
	Wall              WallWallpost    `json:"wall"`
	WallReply         WallWallComment `json:"wall_reply"`
	AudioMessage      Document        `json:"audio_message"`
	Graffiti          Document        `json:"graffiti"`
	Poll              PollsPoll       `json:"poll"`
	Call              MessageCall     `json:"call"`
	Story             StoriesStory    `json:"story"`
	Podcast           PodcastsEpisode `json:"podcast"`
}

type MessageCall struct {
	InitiatorID int     `json:"initiator_id"`
	ReceiverID  int     `json:"receiver_id"`
	State       string  `json:"state"`
	Time        int     `json:"time"`
	Duration    int     `json:"duration"`
	Video       BoolInt `json:"video"`
}

type MessagesPinnedMessage struct {
	Attachments []MessagesMessageAttachment `json:"attachments"`

	// Unique auto-incremented number for all Messages with this peer.
	ConversationMessageID int `json:"conversation_message_id"`

	// Date when the message has been sent in Unixtime.
	Date int `json:"date"`

	// EventType author's ID.
	FromID       int        `json:"from_id"`
	FwdMessages  []*Message `json:"fwd_Messages"`
	Geo          MessageGeo `json:"geo"`
	ID           int        `json:"id"`      // EventType ID
	PeerID       int        `json:"peer_id"` // Peer ID
	ReplyMessage *Message   `json:"reply_message"`
	Text         string     `json:"text"` // EventType text
}

type MessagesUserXtrInvitedBy struct{}

type MessagesForward struct {
	// EventType owner. It is worth passing it on if you want to forward messages
	// from the community to a dialog.
	OwnerID int `json:"owner_id,omitempty"`

	// Identifier of the place from which the messages are to be sent.
	PeerID int `json:"peer_id,omitempty"`

	// Messages can be passed to conversation_message_ids array:
	//
	// - that are in a personal dialog with the bot;
	//
	// - which are outbound messages from the bot;
	//
	// - written after the bot has entered the conversation.
	ConversationMessageIDs []int `json:"conversation_message_ids,omitempty"`
	MessageIDs             []int `json:"message_ids,omitempty"`

	// Reply to messages. It is worth passing if you want to reply to messages
	// in the chat room where the messages are. In this case there should be
	// only one element in the conversation_message_ids/message_ids.
	IsReply bool `json:"is_reply,omitempty"`
}

func (forward MessagesForward) ToJSON() string {
	b, _ := json.Marshal(forward)
	return string(b)
}

type MessagesReactionCountersResponseItem struct {
	CMID     int                                   `json:"cmid"`
	Counters []MessagesReactionCounterResponseItem `json:"counters"`
}

type MessagesReactionCounterResponseItem struct {
	Count      int   `json:"count"`
	ReactionID int   `json:"reaction_id"`
	UserIDs    []int `json:"user_ids"`
}

type MessagesReactionResponseItem struct {
	ReactionID int `json:"reaction_id"`
	UserID     int `json:"user_id"`
}

type MessagesReactionAssetItem struct {
	ReactionID int                    `json:"reaction_id"`
	Links      ReactionAssetItemLinks `json:"links"`
}

type ReactionAssetItemLinks struct {
	BigAnimation   string `json:"big_animation"`
	SmallAnimation string `json:"small_animation"`
	Static         string `json:"static"`
}
