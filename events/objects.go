package events

import (
	"encoding/json"
	"go-vk-sdk/objects"
)

type EventConfirmation struct {
	GroupID int
}

func (e *EventConfirmation) EventType() EventType {
	return EventTypeConfirmation
}

type EventMessageNew struct {
	Message    objects.Message    `json:"message"`
	ClientInfo objects.ClientInfo `json:"client_info"`
}

func (e *EventMessageNew) EventType() EventType {
	return EventTypeMessageNew
}

type EventMessageReply objects.Message

func (e *EventMessageReply) EventType() EventType {
	return EventTypeMessageReply
}

type EventMessageEdit objects.Message

func (e *EventMessageEdit) EventType() EventType {
	return EventTypeMessageEdit
}

type EventMessageAllow struct {
	UserID int    `json:"user_id"`
	Key    string `json:"key"`
}

func (e *EventMessageAllow) EventType() EventType {
	return EventTypeMessageAllow
}

type EventMessageDeny struct {
	UserID int `json:"user_id"`
}

func (e *EventMessageDeny) EventType() EventType {
	return EventTypeMessageDeny
}

type EventMessageTypingState struct {
	State  string `json:"state"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}

func (e *EventMessageTypingState) EventType() EventType {
	return EventTypeMessageTypingState
}

type EventMessageEvent struct {
	UserID                int             `json:"user_id"`
	PeerID                int             `json:"peer_id"`
	EventID               string          `json:"event_id"`
	Payload               json.RawMessage `json:"payload"`
	ConversationMessageID int             `json:"conversation_message_id"`
}

func (e *EventMessageEvent) EventType() EventType {
	return EventTypeMessageEvent
}

type EventPhotoNew objects.Photo

func (e *EventPhotoNew) EventType() EventType {
	return EventTypePhotoNew
}

type EventPhotoCommentNew objects.WallWallComment

func (e *EventPhotoCommentNew) EventType() EventType {
	return EventTypePhotoCommentNew
}

type EventPhotoCommentEdit objects.WallWallComment

func (e *EventPhotoCommentEdit) EventType() EventType {
	return EventTypePhotoCommentEdit
}

type EventPhotoCommentRestore objects.WallWallComment

func (e *EventPhotoCommentRestore) EventType() EventType {
	return EventTypePhotoCommentRestore
}

type EventPhotoCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	PhotoID   int `json:"photo_id"`
}

func (e *EventPhotoCommentDelete) EventType() EventType {
	return EventTypePhotoCommentDelete
}

type EventAudioNew objects.Audio

func (e *EventAudioNew) EventType() EventType {
	return EventTypeAudioNew
}

type EventVideoNew objects.Video

func (e *EventVideoNew) EventType() EventType {
	return EventTypeVideoNew
}

type EventVideoCommentNew objects.WallWallComment

func (e *EventVideoCommentNew) EventType() EventType {
	return EventTypeVideoCommentNew
}

type EventVideoCommentEdit objects.WallWallComment

func (e *EventVideoCommentEdit) EventType() EventType {
	return EventTypeVideoCommentEdit
}

type EventVideoCommentRestore objects.WallWallComment

func (e *EventVideoCommentRestore) EventType() EventType {
	return EventTypeVideoCommentRestore
}

type EventVideoCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	VideoID   int `json:"video_id"`
}

func (e *EventVideoCommentDelete) EventType() EventType {
	return EventTypeVideoCommentDelete
}

type EventWallPostNew objects.WallWallpost

func (e *EventWallPostNew) EventType() EventType {
	return EventTypeWallPostNew
}

type EventWallRepost objects.WallWallpost

func (e *EventWallRepost) EventType() EventType {
	return EventTypeWallRepost
}

type EventWallReplyNew objects.WallWallComment

func (e *EventWallReplyNew) EventType() EventType {
	return EventTypeWallReplyNew
}

type EventWallReplyEdit objects.WallWallComment

func (e *EventWallReplyEdit) EventType() EventType {
	return EventTypeWallReplyEdit
}

type EventWallReplyRestore objects.WallWallComment

func (e *EventWallReplyRestore) EventType() EventType {
	return EventTypeWallReplyRestore
}

type EventWallReplyDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	DeleterID int `json:"deleter_id"`
	PostID    int `json:"post_id"`
}

func (e *EventWallReplyDelete) EventType() EventType {
	return EventTypeWallReplyDelete
}

type EventBoardPostNew objects.BoardTopicComment

func (e *EventBoardPostNew) EventType() EventType {
	return EventTypeBoardPostNew
}

type EventBoardPostEdit objects.BoardTopicComment

func (e *EventBoardPostEdit) EventType() EventType {
	return EventTypeBoardPostEdit
}

type EventBoardPostRestore objects.BoardTopicComment

func (e *EventBoardPostRestore) EventType() EventType {
	return EventTypeBoardPostRestore
}

type EventBoardPostDelete struct {
	TopicOwnerID int `json:"topic_owner_id"`
	TopicID      int `json:"topic_id"`
	ID           int `json:"id"`
}

func (e *EventBoardPostDelete) EventType() EventType {
	return EventTypeBoardPostDelete
}

type EventMarketCommentNew objects.WallWallComment

func (e *EventMarketCommentNew) EventType() EventType {
	return EventTypeMarketCommentNew
}

type EventMarketCommentEdit objects.WallWallComment

func (e *EventMarketCommentEdit) EventType() EventType {
	return EventTypeMarketCommentEdit
}

type EventMarketCommentRestore objects.WallWallComment

func (e *EventMarketCommentRestore) EventType() EventType {
	return EventTypeMarketCommentRestore
}

type EventMarketCommentDelete struct {
	OwnerID   int `json:"owner_id"`
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	DeleterID int `json:"deleter_id"`
	ItemID    int `json:"item_id"`
}

func (e *EventMarketCommentDelete) EventType() EventType {
	return EventTypeMarketCommentDelete
}

type EventMarketOrderNew objects.MarketOrder

func (e *EventMarketOrderNew) EventType() EventType {
	return EventTypeMarketOrderNew
}

type EventMarketOrderEdit objects.MarketOrder

func (e *EventMarketOrderEdit) EventType() EventType {
	return EventTypeMarketOrderEdit
}

type EventGroupLeave struct {
	UserID int                    `json:"user_id"`
	Self   objects.NumberFlagBool `json:"self"`
}

func (e *EventGroupLeave) EventType() EventType {
	return EventTypeGroupLeave
}

type EventGroupJoin struct {
	UserID   int    `json:"user_id"`
	JoinType string `json:"join_type"`
}

func (e *EventGroupJoin) EventType() EventType {
	return EventTypeGroupJoin
}

type EventUserBlock struct {
	AdminID     int    `json:"admin_id"`
	UserID      int    `json:"user_id"`
	UnblockDate int    `json:"unblock_date"`
	Reason      int    `json:"reason"`
	Comment     string `json:"comment"`
}

func (e *EventUserBlock) EventType() EventType {
	return EventTypeUserBlock
}

type EventUserUnblock struct {
	AdminID   int `json:"admin_id"`
	UserID    int `json:"user_id"`
	ByEndDate int `json:"by_end_date"`
}

func (e *EventUserUnblock) EventType() EventType {
	return EventTypeUserUnblock
}

type EventPollVoteNew struct {
	OwnerID  int `json:"owner_id"`
	PollID   int `json:"poll_id"`
	OptionID int `json:"option_id"`
	UserID   int `json:"user_id"`
}

func (e *EventPollVoteNew) EventType() EventType {
	return EventTypePollVoteNew
}

type EventGroupOfficersEdit struct {
	AdminID  int `json:"admin_id"`
	UserID   int `json:"user_id"`
	LevelOld int `json:"level_old"`
	LevelNew int `json:"level_new"`
}

func (e *EventGroupOfficersEdit) EventType() EventType {
	return EventTypeGroupOfficersEdit
}

type Changes struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type ChangesInt struct {
	OldValue int `json:"old_value"`
	NewValue int `json:"new_value"`
}

type EventGroupChangeSettings struct {
	UserID  int `json:"user_id"`
	Changes struct {
		Title                 Changes    `json:"title"`
		Description           Changes    `json:"description"`
		Access                ChangesInt `json:"access"`
		ScreenName            Changes    `json:"screen_name"`
		PublicCategory        ChangesInt `json:"public_category"`
		PublicSubcategory     ChangesInt `json:"public_subcategory"`
		AgeLimits             ChangesInt `json:"age_limits"`
		Website               Changes    `json:"website"`
		StatusDefault         Changes    `json:"status_default"`
		Wall                  ChangesInt `json:"wall"`
		Replies               ChangesInt `json:"replies"`
		Topics                ChangesInt `json:"topics"`
		Audio                 ChangesInt `json:"audio"`
		Photos                ChangesInt `json:"photos"`
		Video                 ChangesInt `json:"video"`
		Market                ChangesInt `json:"market"`
		Docs                  ChangesInt `json:"docs"`
		Messages              ChangesInt `json:"messages"`
		EventGroupID          ChangesInt `json:"event_group_id"`
		Links                 Changes    `json:"links"`
		Email                 Changes    `json:"email"`
		EventStartDate        ChangesInt `json:"event_start_date::"`
		EventFinishDate       ChangesInt `json:"event_finish_date:"`
		Subject               Changes    `json:"subject"`
		MarketWiki            Changes    `json:"market_wiki"`
		DisableMarketComments ChangesInt `json:"disable_market_comments"`
		Phone                 ChangesInt `json:"phone"`
		CountryID             ChangesInt `json:"country_id"`
		CityID                ChangesInt `json:"city_id"`
	} `json:"Changes"`
}

func (e *EventGroupChangeSettings) EventType() EventType {
	return EventTypeGroupChangeSettings
}

type EventGroupChangePhoto struct {
	UserID int           `json:"user_id"`
	Photo  objects.Photo `json:"photo"`
}

func (e *EventGroupChangePhoto) EventType() EventType {
	return EventTypeGroupChangePhoto
}

type EventVkpayTransaction struct {
	FromID      int    `json:"from_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}

func (e *EventVkpayTransaction) EventType() EventType {
	return EventTypeVkpayTransaction
}

type EventLeadFormsNew struct {
	LeadID   int    `json:"lead_id"`
	GroupID  int    `json:"group_id"`
	UserID   int    `json:"user_id"`
	FormID   int    `json:"form_id"`
	FormName string `json:"form_name"`
	AdID     int    `json:"ad_id"`
	Answers  []struct {
		Key      string `json:"key"`
		Question string `json:"question"`
		Answer   string `json:"answer"`
	} `json:"answers"`
}

func (e *EventLeadFormsNew) EventType() EventType {
	return EventTypeLeadFormsNew
}

type EventAppPayload struct {
	UserID  int    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Payload string `json:"payload"`
}

func (e *EventAppPayload) EventType() EventType {
	return EventTypeAppPayload
}

type EventMessageRead struct {
	FromID        int `json:"from_id"`
	PeerID        int `json:"peer_id"`
	ReadMessageID int `json:"read_message_id"`
}

func (e *EventMessageRead) EventType() EventType {
	return EventTypeMessageRead
}

type EventLikeAdd struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"`
}

func (e *EventLikeAdd) EventType() EventType {
	return EventTypeLikeAdd
}

type EventLikeRemove struct {
	LikerID       int    `json:"liker_id"`
	ObjectType    string `json:"object_type"`
	ObjectOwnerID int    `json:"object_owner_id"`
	ObjectID      int    `json:"object_id"`
	ThreadReplyID int    `json:"thread_reply_id"`
	PostID        int    `json:"post_id"`
}

func (e *EventLikeRemove) EventType() EventType {
	return EventTypeLikeRemove
}

type EventDonutSubscriptionCreate struct {
	Amount           float64 `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
	UserID           float64 `json:"user_id"`
}

func (e *EventDonutSubscriptionCreate) EventType() EventType {
	return EventTypeDonutSubscriptionCreate
}

type EventDonutSubscriptionProlonged struct {
	Amount           float64 `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
	UserID           float64 `json:"user_id"`
}

func (e *EventDonutSubscriptionProlonged) EventType() EventType {
	return EventTypeDonutSubscriptionProlonged
}

type EventDonutSubscriptionExpired struct {
	UserID float64 `json:"user_id"`
}

func (e *EventDonutSubscriptionExpired) EventType() EventType {
	return EventTypeDonutSubscriptionExpired
}

type EventDonutSubscriptionCancelled struct {
	UserID float64 `json:"user_id"`
}

func (e *EventDonutSubscriptionCancelled) EventType() EventType {
	return EventTypeDonutSubscriptionCancelled
}

type EventDonutSubscriptionPriceChanged struct {
	AmountOld            float64 `json:"amount_old"`
	AmountNew            float64 `json:"amount_new"`
	AmountDiff           float64 `json:"amount_diff"`
	AmountDiffWithoutFee float64 `json:"amount_diff_without_fee"`
	UserID               float64 `json:"user_id"`
}

func (e *EventDonutSubscriptionPriceChanged) EventType() EventType {
	return EventTypeDonutSubscriptionPriceChanged
}

type EventDonutMoneyWithdraw struct {
	Amount           float64 `json:"amount"`
	AmountWithoutFee float64 `json:"amount_without_fee"`
}

func (e *EventDonutMoneyWithdraw) EventType() EventType {
	return EventTypeDonutMoneyWithdraw
}

type EventDonutMoneyWithdrawError struct {
	Reason string `json:"reason"`
}

func (e *EventDonutMoneyWithdrawError) EventType() EventType {
	return EventTypeDonutMoneyWithdrawError
}
