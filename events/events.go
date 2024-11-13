package events

import (
	"encoding/json"
	"errors"
)

type EventType string

const (
	EventTypeCallback                      EventType = "callback"     // only callback
	EventTypeConfirmation                  EventType = "confirmation" // only callback
	EventTypeMessageNew                    EventType = "message_new"
	EventTypeMessageReply                  EventType = "message_reply"
	EventTypeMessageEdit                   EventType = "message_edit"
	EventTypeMessageAllow                  EventType = "message_allow"
	EventTypeMessageDeny                   EventType = "message_deny"
	EventTypeMessageTypingState            EventType = "message_typing_state"
	EventTypeMessageEvent                  EventType = "message_event"
	EventTypePhotoNew                      EventType = "photo_new"
	EventTypePhotoCommentNew               EventType = "photo_comment_new"
	EventTypePhotoCommentEdit              EventType = "photo_comment_edit"
	EventTypePhotoCommentRestore           EventType = "photo_comment_restore"
	EventTypePhotoCommentDelete            EventType = "photo_comment_delete"
	EventTypeAudioNew                      EventType = "audio_new"
	EventTypeVideoNew                      EventType = "video_new"
	EventTypeVideoCommentNew               EventType = "video_comment_new"
	EventTypeVideoCommentEdit              EventType = "video_comment_edit"
	EventTypeVideoCommentRestore           EventType = "video_comment_restore"
	EventTypeVideoCommentDelete            EventType = "video_comment_delete"
	EventTypeWallPostNew                   EventType = "wall_post_new"
	EventTypeWallRepost                    EventType = "wall_repost"
	EventTypeWallReplyNew                  EventType = "wall_reply_new"
	EventTypeWallReplyEdit                 EventType = "wall_reply_edit"
	EventTypeWallReplyRestore              EventType = "wall_reply_restore"
	EventTypeWallReplyDelete               EventType = "wall_reply_delete"
	EventTypeBoardPostNew                  EventType = "board_post_new"
	EventTypeBoardPostEdit                 EventType = "board_post_edit"
	EventTypeBoardPostRestore              EventType = "board_post_restore"
	EventTypeBoardPostDelete               EventType = "board_post_delete"
	EventTypeMarketCommentNew              EventType = "market_comment_new"
	EventTypeMarketCommentEdit             EventType = "market_comment_edit"
	EventTypeMarketCommentRestore          EventType = "market_comment_restore"
	EventTypeMarketCommentDelete           EventType = "market_comment_delete"
	EventTypeMarketOrderNew                EventType = "market_order_new"
	EventTypeMarketOrderEdit               EventType = "market_order_edit"
	EventTypeGroupLeave                    EventType = "group_leave"
	EventTypeGroupJoin                     EventType = "group_join"
	EventTypeUserBlock                     EventType = "user_block"
	EventTypeUserUnblock                   EventType = "user_unblock"
	EventTypePollVoteNew                   EventType = "poll_vote_new"
	EventTypeGroupOfficersEdit             EventType = "group_officers_edit"
	EventTypeGroupChangeSettings           EventType = "group_change_settings"
	EventTypeGroupChangePhoto              EventType = "group_change_photo"
	EventTypeVkpayTransaction              EventType = "vkpay_transaction"
	EventTypeLeadFormsNew                  EventType = "lead_forms_new"
	EventTypeAppPayload                    EventType = "app_payload"
	EventTypeMessageRead                   EventType = "message_read"
	EventTypeLikeAdd                       EventType = "like_add"
	EventTypeLikeRemove                    EventType = "like_remove"
	EventTypeDonutSubscriptionCreate       EventType = "donut_subscription_create"
	EventTypeDonutSubscriptionProlonged    EventType = "donut_subscription_prolonged"
	EventTypeDonutSubscriptionExpired      EventType = "donut_subscription_expired"
	EventTypeDonutSubscriptionCancelled    EventType = "donut_subscription_cancelled"
	EventTypeDonutSubscriptionPriceChanged EventType = "donut_subscription_price_changed"
	EventTypeDonutMoneyWithdraw            EventType = "donut_money_withdraw"
	EventTypeDonutMoneyWithdrawError       EventType = "donut_money_withdraw_error"
)

func NewEvent(eventUpdate *EventUpdate) (Event, error) {
	var event Event
	var err error

	switch eventUpdate.Type {
	case EventTypeMessageNew:
		event = &EventMessageNew{}
	case EventTypeMessageReply:
		event = &EventMessageReply{}
	case EventTypeMessageEdit:
		event = &EventMessageEdit{}
	case EventTypeMessageAllow:
		event = &EventMessageAllow{}
	case EventTypeMessageDeny:
		event = &EventMessageDeny{}
	case EventTypeMessageTypingState:
		event = &EventMessageTypingState{}
	case EventTypeMessageEvent:
		event = &EventMessageEvent{}
	case EventTypePhotoNew:
		event = &EventPhotoNew{}
	case EventTypePhotoCommentNew:
		event = &EventPhotoCommentNew{}
	case EventTypePhotoCommentEdit:
		event = &EventPhotoCommentEdit{}
	case EventTypePhotoCommentRestore:
		event = &EventPhotoCommentRestore{}
	case EventTypePhotoCommentDelete:
		event = &EventPhotoCommentDelete{}
	case EventTypeAudioNew:
		event = &EventAudioNew{}
	case EventTypeVideoNew:
		event = &EventVideoNew{}
	case EventTypeVideoCommentNew:
		event = &EventVideoCommentNew{}
	case EventTypeVideoCommentEdit:
		event = &EventVideoCommentEdit{}
	case EventTypeVideoCommentRestore:
		event = &EventVideoCommentRestore{}
	case EventTypeVideoCommentDelete:
		event = &EventVideoCommentDelete{}
	case EventTypeWallPostNew:
		event = &EventWallPostNew{}
	case EventTypeWallRepost:
		event = &EventWallRepost{}
	case EventTypeWallReplyNew:
		event = &EventWallReplyNew{}
	case EventTypeWallReplyEdit:
		event = &EventWallReplyEdit{}
	case EventTypeWallReplyRestore:
		event = &EventWallReplyRestore{}
	case EventTypeWallReplyDelete:
		event = &EventWallReplyDelete{}
	case EventTypeBoardPostNew:
		event = &EventBoardPostNew{}
	case EventTypeBoardPostEdit:
		event = &EventBoardPostEdit{}
	case EventTypeBoardPostRestore:
		event = &EventBoardPostRestore{}
	case EventTypeBoardPostDelete:
		event = &EventBoardPostDelete{}
	case EventTypeMarketCommentNew:
		event = &EventMarketCommentNew{}
	case EventTypeMarketCommentEdit:
		event = &EventMarketCommentEdit{}
	case EventTypeMarketCommentRestore:
		event = &EventMarketCommentRestore{}
	case EventTypeMarketCommentDelete:
		event = &EventMarketCommentDelete{}
	case EventTypeMarketOrderNew:
		event = &EventMarketOrderNew{}
	case EventTypeMarketOrderEdit:
		event = &EventMarketOrderEdit{}
	case EventTypeGroupLeave:
		event = &EventGroupLeave{}
	case EventTypeGroupJoin:
		event = &EventGroupJoin{}
	case EventTypeUserBlock:
		event = &EventUserBlock{}
	case EventTypeUserUnblock:
		event = &EventUserUnblock{}
	case EventTypePollVoteNew:
		event = &EventPollVoteNew{}
	case EventTypeGroupOfficersEdit:
		event = &EventGroupOfficersEdit{}
	case EventTypeGroupChangeSettings:
		event = &EventGroupChangeSettings{}
	case EventTypeGroupChangePhoto:
		event = &EventGroupChangePhoto{}
	case EventTypeVkpayTransaction:
		event = &EventVkpayTransaction{}
	case EventTypeLeadFormsNew:
		event = &EventLeadFormsNew{}
	case EventTypeAppPayload:
		event = &EventAppPayload{}
	case EventTypeMessageRead:
		event = &EventMessageRead{}
	case EventTypeLikeAdd:
		event = &EventLikeAdd{}
	case EventTypeLikeRemove:
		event = &EventLikeRemove{}
	case EventTypeDonutSubscriptionCreate:
		event = &EventDonutSubscriptionCreate{}
	case EventTypeDonutSubscriptionProlonged:
		event = &EventDonutSubscriptionProlonged{}
	case EventTypeDonutSubscriptionExpired:
		event = &EventDonutSubscriptionExpired{}
	case EventTypeDonutSubscriptionCancelled:
		event = &EventDonutSubscriptionCancelled{}
	case EventTypeDonutSubscriptionPriceChanged:
		event = &EventDonutSubscriptionPriceChanged{}
	case EventTypeDonutMoneyWithdraw:
		event = &EventDonutMoneyWithdraw{}
	case EventTypeDonutMoneyWithdrawError:
		event = &EventDonutMoneyWithdrawError{}
	default:
		err = errors.New("events.NewEvent(): unknown event type")
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(eventUpdate.Object, event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
