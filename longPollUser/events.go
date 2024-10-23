package longPollUser

import (
	"go-vk-sdk/errors"
	"time"
)

// Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started

type EventType int

const (
	EventTypeMessageFlagsReplace        EventType = 1   // $message_id(integer), $flags(integer), extra_fields*. Replaces message flags (FLAGS := $flags).
	EventTypeMessageFlagsSet            EventType = 2   // $message_id(integer), $mask(integer), extra_fields*. Sets message flags (FLAGS = $mask).
	EventTypeMessageFlagsReset          EventType = 3   // $message_id(integer), $mask(integer), extra_fields*. Resets message flags (FLAGS &= ~$mask).
	EventTypeMessageNew                 EventType = 4   // $message_id(integer), $flags(integer), $minor_id(integer), extra_fields*. Adds a new message.
	EventTypeMessageEdit                EventType = 5   // $message_id(integer), $mask(integer), $peer_id(integer), $timestamp(integer), $new_text(string), $attachments(array). Edits a message.
	EventTypeMessagesIncomingRead       EventType = 6   // $peer_id(integer), $local_id(integer). Reads all incoming messages in $peer_id up to $local_id.
	EventTypeMessagesOutgoingRead       EventType = 7   // $peer_id(integer), $local_id(integer). Reads all outgoing messages in $peer_id up to $local_id.
	EventTypeFriendOnline               EventType = 8   // -$user_id(integer), $extra(integer), $timestamp(integer). Friend $user_id is online. $extra is non-zero if mode flag 64 is set.
	EventTypeFriendOffline              EventType = 9   // -$user_id(integer), $flags(integer), $timestamp(integer). Friend $user_id is offline. $flags is 0 if the user left the site, 1 if offline due to timeout.
	EventTypeDialogFlagsReset           EventType = 10  // $peer_id(integer), $mask(integer). Resets dialog flags $peer_id (PEER_FLAGS &= ~$flags). Only for community dialogs.
	EventTypeDialogFlagsReplace         EventType = 11  // $peer_id(integer), $flags(integer). Replaces dialog flags $peer_id (PEER_FLAGS := $flags). Only for community dialogs.
	EventTypeDialogFlagsSet             EventType = 12  // $peer_id(integer), $mask(integer). Sets dialog flags $peer_id (PEER_FLAGS = $flags). Only for community dialogs.
	EventTypeMessagesDelete             EventType = 13  // $peer_id(integer), $local_id(integer). Deletes all messages in dialog $peer_id up to $local_id.
	EventTypeMessagesRestore            EventType = 14  // $peer_id(integer), $local_id(integer). Restores recently deleted messages in dialog $peer_id up to $local_id.
	EventTypeMajorIDChange              EventType = 20  // $peer_id(integer), $major_id(integer). The $major_id in dialog $peer_id has changed.
	EventTypeMinorIDChange              EventType = 21  // $peer_id(integer), $minor_id(integer). The $minor_id in dialog $peer_id has changed.
	EventTypeChatParametersChange       EventType = 51  // $chat_id(integer), $self(integer). One of the parameters (composition, topic) of the chat $chat_id has changed. $self indicates if changes were made by the user (1/0).
	EventTypeChatInfoChange             EventType = 52  // $type_id(integer), $peer_id(integer), $info(integer). Chat info $peer_id changed with type $type_id. $info provides additional details about the change.
	EventTypeUserTypingDialog           EventType = 61  // $user_id(integer), $flags(integer). LongPoll $user_id is typing in a dialog. Event occurs every ~5 seconds while typing. $flags = 1.
	EventTypeUserTypingChat             EventType = 62  // $user_id(integer), $chat_id(integer). LongPoll $user_id is typing in a chat $chat_id.
	EventTypeUsersTypingChat            EventType = 63  // $user_ids(integer), $peer_id(integer), $total_count(integer), $ts(integer). Users $user_ids are typing in chat $peer_id. Up to five users are transmitted, total number in $total_count.
	EventTypeUsersRecordingAudioMessage EventType = 64  // $user_ids(integer), $peer_id(integer), $total_count(integer), $ts(integer). Users $user_ids are recording voice messages in chat $peer_id.
	EventTypeUserCall                   EventType = 70  // $user_id(integer), $call_id(integer). LongPoll $user_id made a call with ID $call_id.
	EventTypeMenuCounterChange          EventType = 80  // $count(integer). The counter in the left menu changed to $count.
	EventTypeNotificationSettingsChange EventType = 114 // $peer_id(integer), $sound(integer), $disabled_until(integer). Notification settings changed for $peer_id. $sound indicates on/off, $disabled_until is the mute duration (-1: forever, 0: off).
)

type Event interface {
	init([]interface{}) error
	EventType() EventType
}

type EventUpdate struct {
	Type  EventType
	Event interface{}
}

func newEventUpdate(data []interface{}, mode ExtraOptionsMode) (*EventUpdate, error) {
	event, err := newEvent(data, mode)
	if err != nil {
		return nil, err
	}

	return &EventUpdate{
		Type:  EventType(data[0].(float64)),
		Event: event,
	}, nil
}

func newEvent(data []interface{}, mode ExtraOptionsMode) (Event, error) {
	key := EventType(data[0].(float64))

	var event Event

	switch key {
	case EventTypeMessageFlagsReplace:
		event = &EventMessageFlagsReplace{}
	case EventTypeMessageFlagsSet:
		event = &EventMessageFlagsSet{}
	case EventTypeMessageFlagsReset:
		event = &EventMessageFlagsReset{}
	case EventTypeMessageNew:
		event = &EventMessageNew{}
	case EventTypeMessageEdit:
		event = &EventMessageEdit{}
	case EventTypeMessagesIncomingRead:
		event = &EventMessagesIncomingRead{}
	case EventTypeMessagesOutgoingRead:
		event = &EventMessagesOutgoingRead{}
	case EventTypeFriendOnline:
		event = &EventFriendOnline{}
	case EventTypeFriendOffline:
		event = &EventFriendOffline{}
	case EventTypeDialogFlagsReset:
		event = &EventDialogFlagsReset{}
	case EventTypeDialogFlagsReplace:
		event = &EventDialogFlagsReplace{}
	case EventTypeDialogFlagsSet:
		event = &EventDialogsFlagsSet{}
	case EventTypeMessagesDelete:
		event = &EventMessagesDelete{}
	case EventTypeMessagesRestore:
		event = &EventMessagesRestore{}
	case EventTypeMajorIDChange:
		event = &EventMajorIDChange{}
	case EventTypeMinorIDChange:
		event = &EventMinorIDChange{}
	case EventTypeChatParametersChange:
		event = &EventChatParametersChange{}
	case EventTypeChatInfoChange:
		event = &EventChatInfoChange{}
	case EventTypeUserTypingDialog:
		event = &EventUserTypingDialog{}
	case EventTypeUserTypingChat:
		event = &EventUserTypingChat{}
	case EventTypeUsersTypingChat:
		event = &EventUsersTypingChat{}
	case EventTypeUsersRecordingAudioMessage:
		event = &EventUsersRecordingAudioMessage{}
	case EventTypeUserCall:
		event = &EventUserCall{}
	case EventTypeMenuCounterChange:
		event = &EventMenuCounterChange{}
	case EventTypeNotificationSettingsChange:
		{
			e := &EventNotificationSettingsChange{}
			if mode&ExtraOptionsModeExtendedEvents != 0 {
				err := e.initMode8(data)
				return e, err
			} else {
				err := e.init(data)
				return e, err
			}
		}
	default:
		return nil, &errors.InvalidEventTypeError{Type: int(key)}
	}

	err := event.init(data)
	return event, err
}

type EventMessageFlagsReplace struct {
	MessageID int
	Flags     MessageFlag
	ExtraFieldsMessages
}

func (e *EventMessageFlagsReplace) EventType() EventType {
	return EventTypeMessageFlagsReplace
}

func (e *EventMessageFlagsReplace) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessageFlagsReplace", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = MessageFlag(int(v))
	}

	return e.ExtraFieldsMessages.init(i)
}

// EventMessageFlagsSet Code = 2
type EventMessageFlagsSet struct {
	MessageID int
	Mask      MessageFlag
	ExtraFieldsMessages
}

func (e *EventMessageFlagsSet) EventType() EventType {
	return EventTypeMessageFlagsSet
}

func (e *EventMessageFlagsSet) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessageFlagsSet", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Mask = MessageFlag(int(v))
	}

	return e.ExtraFieldsMessages.init(i)
}

// EventMessageFlagsReset Code = 3
type EventMessageFlagsReset struct {
	MessageID int
	Mask      MessageFlag
	ExtraFieldsMessages
}

func (e *EventMessageFlagsReset) EventType() EventType {
	return EventTypeMessageFlagsReset
}

func (e *EventMessageFlagsReset) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessageFlagsReset", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Mask = MessageFlag(int(v))
	}

	return e.ExtraFieldsMessages.init(i)
}

// EventMessageNew Code = 4
type EventMessageNew struct {
	MessageID int
	Flags     MessageFlag
	ExtraFieldsMessages
}

func (e *EventMessageNew) EventType() EventType {
	return EventTypeMessageNew
}

func (e *EventMessageNew) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessageNew", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = MessageFlag(int(v))
	}

	return e.ExtraFieldsMessages.init(i)
}

// EventMessageEdit Code = 5
type EventMessageEdit struct {
	MessageID      int
	Flags          MessageFlag
	PeerID         int
	Timestamp      time.Time
	NewText        string
	AdditionalData AdditionalData
	Attachments    Attachments
}

func (e *EventMessageEdit) EventType() EventType {
	return EventTypeMessageEdit
}

func (e *EventMessageEdit) init(i []interface{}) error {
	if len(i) < 6 {
		return &errors.TooShortEventArrayError{Action: "EventMessageEdit", Least: 6, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.MessageID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = MessageFlag(int(v))
	}

	if v, ok := i[3].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[4].(float64); ok {
		e.Timestamp = time.Unix(int64(v), 0)
	}

	if v, ok := i[5].(string); ok {
		e.NewText = v
	}

	if len(i) > 6 {
		if v, ok := i[6].(map[string]interface{}); ok {
			e.AdditionalData.init(v)
		}
	}

	if len(i) > 7 {
		if v, ok := i[7].(map[string]interface{}); ok {
			e.Attachments = v
		}
	}

	return nil
}

// EventMessagesIncomingRead Code = 6
type EventMessagesIncomingRead struct {
	PeerID  int
	LocalID int
}

func (e *EventMessagesIncomingRead) EventType() EventType {
	return EventTypeMessagesIncomingRead
}

func (e *EventMessagesIncomingRead) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessagesIncomingRead", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.LocalID = int(v)
	}

	return nil
}

// EventMessagesOutgoingRead Code = 7
type EventMessagesOutgoingRead struct {
	PeerID  int
	LocalID int
}

func (e *EventMessagesOutgoingRead) EventType() EventType {
	return EventTypeMessagesOutgoingRead
}

func (e *EventMessagesOutgoingRead) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessagesOutgoingRead", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.LocalID = int(v)
	}

	return nil
}

// EventFriendOnline Code = 8
type EventFriendOnline struct {
	UserID    int
	Extra     int
	Timestamp time.Time
	AppID     PlatformType
}

func (e *EventFriendOnline) EventType() EventType {
	return EventTypeFriendOnline
}

func (e *EventFriendOnline) init(i []interface{}) error {
	if len(i) < 4 {
		return &errors.TooShortEventArrayError{Action: "EventFriendOnline", Least: 4, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Extra = int(v)
	}

	if v, ok := i[3].(float64); ok {
		e.Timestamp = time.Unix(int64(v), 0)
	}

	if v, ok := i[4].(float64); ok {
		e.AppID = PlatformType(v)
	}

	return nil
}

// EventFriendOffline Code = 9
type EventFriendOffline struct {
	UserID    int
	Flags     int
	Timestamp time.Time
}

func (e *EventFriendOffline) EventType() EventType {
	return EventTypeFriendOffline
}

func (e *EventFriendOffline) init(i []interface{}) error {
	if len(i) < 4 {
		return &errors.TooShortEventArrayError{Action: "EventFriendOffline", Least: 4, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = int(v)
	}

	if v, ok := i[3].(float64); ok {
		e.Timestamp = time.Unix(int64(v), 0)
	}

	return nil
}

// EventDialogFlagsReset Code = 10
type EventDialogFlagsReset struct {
	PeerID int
	Mask   DialogFlag
}

func (e *EventDialogFlagsReset) EventType() EventType {
	return EventTypeDialogFlagsReset
}

func (e *EventDialogFlagsReset) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventDialogFlagsReset", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Mask = DialogFlag(int(v))
	}

	return nil
}

// EventDialogFlagsReplace Code = 11
type EventDialogFlagsReplace struct {
	PeerID int
	Flags  DialogFlag
}

func (e *EventDialogFlagsReplace) EventType() EventType {
	return EventTypeDialogFlagsReplace
}

func (e *EventDialogFlagsReplace) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventDialogFlagsReplace", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = DialogFlag(int(v))
	}

	return nil
}

// EventDialogsFlagsSet Code = 12
type EventDialogsFlagsSet struct {
	PeerID int
	Mask   DialogFlag
}

func (e *EventDialogsFlagsSet) EventType() EventType {
	return EventTypeDialogFlagsSet
}

func (e *EventDialogsFlagsSet) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventDialogsFlagsSet", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Mask = DialogFlag(int(v))
	}

	return nil
}

// EventMessagesDelete Code = 13
type EventMessagesDelete struct {
	PeerID  int
	LocalID int
}

func (e *EventMessagesDelete) EventType() EventType {
	return EventTypeMessagesDelete
}

func (e *EventMessagesDelete) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessagesDelete", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.LocalID = int(v)
	}

	return nil
}

// EventMessagesRestore Code = 14
type EventMessagesRestore struct {
	PeerID  int
	LocalID int
}

func (e *EventMessagesRestore) EventType() EventType {
	return EventTypeMessagesRestore
}

func (e *EventMessagesRestore) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventMessagesRestore", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.LocalID = int(v)
	}

	return nil
}

// EventMajorIDChange Code = 20
type EventMajorIDChange struct {
	PeerID  int
	MajorID int
}

func (e *EventMajorIDChange) EventType() EventType {
	return EventTypeMajorIDChange
}

func (e *EventMajorIDChange) init(i []interface{}) error {
	if len(i) < 2 {
		return &errors.TooShortEventArrayError{Action: "EventMajorIDChange", Least: 2, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.MajorID = int(v)
	}

	return nil
}

// EventMinorIDChange Code = 21
type EventMinorIDChange struct {
	PeerID  int
	MinorID int
}

func (e *EventMinorIDChange) EventType() EventType {
	return EventTypeMinorIDChange
}

func (e *EventMinorIDChange) init(i []interface{}) error {
	if len(i) < 2 {
		return &errors.TooShortEventArrayError{Action: "EventMinorIDChange", Least: 2, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.MinorID = int(v)
	}

	return nil
}

// EventChatParametersChange Code = 51
//
// One of the parameters (content, topic) of the conversation ChatID was
// changed. Self — 1 or 0 (whether the change was caused by the user).
type EventChatParametersChange struct {
	ChatID int
	Self   int
}

func (e *EventChatParametersChange) EventType() EventType {
	return EventTypeChatParametersChange
}

func (e *EventChatParametersChange) init(i []interface{}) error {
	if len(i) < 2 {
		return &errors.TooShortEventArrayError{Action: "EventChatParametersChange", Least: 2, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.ChatID = int(v)
	}

	if len(i) > 2 {
		if v, ok := i[2].(float64); ok {
			e.Self = int(v)
		}
	}

	return nil
}

// EventChatInfoChange Code = 52
type EventChatInfoChange struct {
	TypeID ExtraChatField
	PeerID int
	Info   int
}

func (e *EventChatInfoChange) EventType() EventType {
	return EventTypeChatInfoChange
}

func (e *EventChatInfoChange) init(i []interface{}) error {
	if len(i) < 4 {
		return &errors.TooShortEventArrayError{Action: "EventChatInfoChange", Least: 4, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.TypeID = ExtraChatField(int(v))
	}

	if v, ok := i[2].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		e.Info = int(v)
	}

	return nil
}

// EventUserTypingDialog Code = 61
type EventUserTypingDialog struct {
	UserID int
	Flags  int
}

func (e *EventUserTypingDialog) EventType() EventType {
	return EventTypeUserTypingDialog
}

func (e *EventUserTypingDialog) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventUserTypingDialog", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.Flags = int(v)
	}

	return nil
}

// EventUserTypingChat Code = 62
type EventUserTypingChat struct {
	UserID int
	ChatID int
}

func (e *EventUserTypingChat) EventType() EventType {
	return EventTypeUserTypingChat
}

func (e *EventUserTypingChat) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventUserTypingChat", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.ChatID = int(v)
	}

	return nil
}

// EventUsersTypingChat Code = 63
type EventUsersTypingChat struct {
	UserIDs    []int
	PeerID     int
	TotalCount int
	Ts         time.Time
}

func (e *EventUsersTypingChat) EventType() EventType {
	return EventTypeUsersTypingChat
}

func (e *EventUsersTypingChat) init(i []interface{}) error {
	if len(i) < 5 {
		return &errors.TooShortEventArrayError{Action: "EventUsersTypingChat", Least: 5, Got: len(i)}
	}

	userIDs, err := interfaceToIDSlice(i[1])
	if err != nil {
		return err
	}

	e.UserIDs = userIDs

	if v, ok := i[2].(float64); ok {
		e.PeerID = int(v)
	}

	if v, ok := i[3].(float64); ok {
		e.TotalCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		e.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// EventUsersRecordingAudioMessage Code = 64
type EventUsersRecordingAudioMessage struct {
	PeerID     int
	UserIDs    []int
	TotalCount int
	Ts         time.Time
}

func (e *EventUsersRecordingAudioMessage) EventType() EventType {
	return EventTypeUsersRecordingAudioMessage
}

func (e *EventUsersRecordingAudioMessage) init(i []interface{}) error {
	if len(i) < 5 {
		return &errors.TooShortEventArrayError{Action: "EventUsersRecordingAudioMessage", Least: 5, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.PeerID = int(v)
	}

	v, err := interfaceToIDSlice(i[2])
	if err != nil {
		return err
	}

	e.UserIDs = v

	if v, ok := i[3].(float64); ok {
		e.TotalCount = int(v)
	}

	if v, ok := i[4].(float64); ok {
		e.Ts = time.Unix(int64(v), 0)
	}

	return nil
}

// EventUserCall Code = 70
type EventUserCall struct {
	UserID int
	CallID int
}

func (e *EventUserCall) EventType() EventType {
	return EventTypeUserCall
}

func (e *EventUserCall) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventUserCall", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.UserID = int(v)
	}

	if v, ok := i[2].(float64); ok {
		e.CallID = int(v)
	}

	return nil
}

// EventMenuCounterChange Code = 80
type EventMenuCounterChange struct {
	Count int
}

func (e *EventMenuCounterChange) EventType() EventType {
	return EventTypeMenuCounterChange
}

func (e *EventMenuCounterChange) init(i []interface{}) error {
	if len(i) < 2 {
		return &errors.TooShortEventArrayError{Action: "EventMenuCounterChange", Least: 2, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.Count = int(v)
	}

	return nil
}

// EventNotificationSettingsChange Code = 114
type EventNotificationSettingsChange struct {
	PeerID        int
	Sound         bool
	DisabledUntil int
}

func (e *EventNotificationSettingsChange) EventType() EventType {
	return EventTypeNotificationSettingsChange
}

// initMode8 should be called if ExtendedEvents flag set.
func (e *EventNotificationSettingsChange) initMode8(i []interface{}) error {
	if len(i) < 2 {
		return &errors.TooShortEventArrayError{Action: "EventNotificationSettingsChange", Least: 2, Got: len(i)}
	}

	v, err := interfaceToStringIntMap(i[1])
	if err != nil {
		return err
	}

	e.PeerID = v["peer_id"]
	e.Sound = v["sound"] > 0
	e.DisabledUntil = v["disabled_until"]

	return nil
}

func (e *EventNotificationSettingsChange) init(i []interface{}) error {
	if len(i) < 3 {
		return &errors.TooShortEventArrayError{Action: "EventNotificationSettingsChange", Least: 3, Got: len(i)}
	}

	if v, ok := i[1].(float64); ok {
		e.Sound = int(v) > 0
	}

	if v, ok := i[2].(float64); ok {
		e.DisabledUntil = int(v)
	}

	return nil
}
