package longPollUser

import (
	"go-vk-sdk/errors"
	"time"
)

// Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started

type ExtraFieldsMessages struct {
	PeerID         int
	Timestamp      time.Time
	Text           string
	Attachments    Attachments // attachments, if mode = 2 was chosen ExtraOptionsMode = ExtraOptionsModeReceiveAttachments
	AdditionalData AdditionalData
}

func (e *ExtraFieldsMessages) init(i []interface{}) error {
	length := len(i)

	if length > 3 {
		if v, ok := i[3].(float64); ok {
			e.PeerID = int(v)
		}
	}

	if length > 4 {
		if v, ok := i[4].(float64); ok {
			e.Timestamp = time.Unix(int64(v), 0)
		}
	}

	if length > 5 {
		if v, ok := i[5].(string); ok {
			e.Text = v
		}
	}

	if length > 6 {
		if v, ok := i[6].(map[string]interface{}); ok {
			e.AdditionalData.init(v)
		}
	}

	if length > 7 {
		if v, ok := i[7].(map[string]interface{}); ok {
			e.Attachments = v
		}
	}

	return nil
}

type Attachments map[string]interface{}

// AdditionalData struct. If mode contains the flag 2 along with text and the
// message topic, a JSON-object may be passed. This object contains media
// attachments or other additional information. Descriptions of the object
// fields are listed below.
type AdditionalData struct {
	Title     string // Message's subject.
	From      string // user ID of who sent the message if the message is from a chat
	RefSource string
	ExpireTTL string
	IsExpired string
	// FromAdmin ID of the administrator who sent the message. It is returned for
	// messages sent from a community (only for community administrators).
	FromAdmin string
	Emoji     string // Message contains emoji.
	Action
}

func (a *AdditionalData) init(v map[string]interface{}) {
	a.Title = getString(v, "title")
	a.From = getString(v, "from")
	a.RefSource = getString(v, "ref_source")
	a.FromAdmin = getString(v, "from_admin")
	a.ExpireTTL = getString(v, "expire_ttl")
	a.IsExpired = getString(v, "is_expired")
	a.Emoji = getString(v, "emoji")
	a.Action.init(v)
}

type Action struct {
	SourceAct           string // Service action name with multiple dialogs
	SourceMid           string // user ID to whom the service action concerns
	SourceText          string
	SourceOldText       string
	SourceMessage       string
	SourceChatLocalID   string
	SourceIsUpdatePhoto bool
}

func (a *Action) init(v map[string]interface{}) {
	a.SourceAct = getString(v, "source_act")

	switch a.SourceAct {
	case "chat_create", "chat_title_update":
		a.SourceText = getString(v, "source_text")
		if a.SourceAct == "chat_title_update" {
			a.SourceOldText = getString(v, "source_old_text")
		}
	case "chat_invite_user", "chat_kick_user":
		a.SourceMid = getString(v, "source_mid")
	case "chat_pin_message", "chat_unpin_message":
		a.SourceMid = getString(v, "source_mid")
		a.SourceMessage = getString(v, "source_message")
		a.SourceChatLocalID = getString(v, "source_chat_local_id")
	case "chat_photo_update":
		a.SourceIsUpdatePhoto = true
	}
}

func getString(m map[string]interface{}, key string) string {
	if val, ok := m[key].(string); ok {
		return val
	}
	return ""
}

func interfaceToStringIntMap(m interface{}) (map[string]int, error) {
	mapInterface, ok := m.(map[string]interface{})
	if !ok {
		return nil, &errors.ExpectedSliceError{V: m}
	}

	result := make(map[string]int, len(mapInterface)) // Предварительно задаем размер

	for key, value := range mapInterface {
		val, ok := value.(float64)
		if !ok {
			return nil, &errors.FailedCastError{V: value}
		}
		result[key] = int(val)
	}

	return result, nil
}

func interfaceToIDSlice(slice interface{}) ([]int, error) {
	sliceInterface, ok := slice.([]interface{})
	if !ok {
		return nil, &errors.ExpectedSliceError{V: slice}
	}

	result := make([]int, len(sliceInterface)) // Предварительно задаем размер

	for i, value := range sliceInterface {
		val, ok := value.(float64)
		if !ok {
			return nil, &errors.FailedCastError{V: value}
		}
		result[i] = int(val)
	}

	return result, nil
}
