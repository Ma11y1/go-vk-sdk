package objects

import (
	"encoding/json"
	"go-vk-sdk/constants"
)

type StoriesViewer struct {
	IsLiked bool `json:"is_liked"`
	UserID  int  `json:"user_id"`

	// For extended
	User struct {
		Type            string `json:"type"`
		ID              int    `json:"id"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		IsClosed        bool   `json:"is_closed"`
		CanAccessClosed bool   `json:"can_access_closed"`
	} `json:"user,omitempty"`
}

type StoriesNarrativeInfo struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Views  int    `json:"views"`
}

type StoriesPromoBlock struct {
	Name        string         `json:"name"`
	Photo50     string         `json:"photo_50"`
	Photo100    string         `json:"photo_100"`
	NotAnimated NumberFlagBool `json:"not_animated"`
}

type StoriesStoryLink struct {
	Text string `json:"text"` // Link text
	URL  string `json:"url"`  // Link URL
}

type StoriesReplies struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

type StoriesQuestions struct {
	Count int `json:"count"` // Replies number.
	New   int `json:"new"`   // New replies number.
}

type StoriesStoryStats struct {
	Answer      StoriesStoryStatsStat `json:"answer"`
	Bans        StoriesStoryStatsStat `json:"bans"`
	OpenLink    StoriesStoryStatsStat `json:"open_link"`
	Replies     StoriesStoryStatsStat `json:"replies"`
	Shares      StoriesStoryStatsStat `json:"shares"`
	Subscribers StoriesStoryStatsStat `json:"subscribers"`
	Views       StoriesStoryStatsStat `json:"views"`
	Likes       StoriesStoryStatsStat `json:"likes"`
}

type StoriesStoryStatsStat struct {
	Count int    `json:"count"` // Stat value
	State string `json:"state"`
}

type StoriesStory struct {
	AccessKey string         `json:"access_key"` // Access key for private objects.
	ExpiresAt int            `json:"expires_at"` // Story expiration time. Unixtime.
	CanHide   NumberFlagBool `json:"can_hide"`
	// Information whether story has question sticker and current user can send question to the author
	CanAsk NumberFlagBool `json:"can_ask"`
	// Information whether story has question sticker and current user can send anonymous question to the author
	CanAskAnonymous NumberFlagBool `json:"can_ask_anonymous"`

	// Information whether current user can comment the story (0 - no, 1 - yes).
	CanComment NumberFlagBool `json:"can_comment"`

	// Information whether current user can reply to the story
	// (0 - no, 1 - yes).
	CanReply NumberFlagBool `json:"can_reply"`

	// Information whether current user can see the story (0 - no, 1 - yes).
	CanSee NumberFlagBool `json:"can_see"`

	// Information whether current user can share the story (0 - no, 1 - yes).
	CanShare NumberFlagBool `json:"can_share"`

	// Information whether the story is deleted (false - no, true - yes).
	IsDeleted NumberFlagBool `json:"is_deleted"`

	// Information whether the story is expired (false - no, true - yes).
	IsExpired NumberFlagBool `json:"is_expired"`

	// Is video without sound
	NoSound NumberFlagBool `json:"no_sound"`

	// Does author have stories privacy restrictions
	IsRestricted NumberFlagBool `json:"is_restricted"`

	CanUseInNarrative NumberFlagBool `json:"can_use_in_narrative"`

	// Information whether current user has seen the story or not
	// (0 - no, 1 - yes).
	Seen                 NumberFlagBool           `json:"seen"`
	IsOwnerPinned        NumberFlagBool           `json:"is_owner_pinned"`
	IsOneTime            NumberFlagBool           `json:"is_one_time"`
	IsAdvice             NumberFlagBool           `json:"is_advice,omitempty"`
	NeedMute             NumberFlagBool           `json:"need_mute"`
	MuteReply            NumberFlagBool           `json:"mute_reply"`
	CanLike              NumberFlagBool           `json:"can_like"`
	Date                 int                      `json:"date"` // Date when story has been added in Unixtime.
	ID                   int                      `json:"id"`   // Story ID.
	Link                 StoriesStoryLink         `json:"link"`
	OwnerID              int                      `json:"owner_id"` // Story owner's ID.
	ParentStory          *StoriesStory            `json:"parent_story"`
	ParentStoryAccessKey string                   `json:"parent_story_access_key"` // Access key for private objects.
	ParentStoryID        int                      `json:"parent_story_id"`         // Parent story ID.
	ParentStoryOwnerID   int                      `json:"parent_story_owner_id"`   // Parent story owner's ID.
	Photo                Photo                    `json:"photo"`
	Replies              StoriesReplies           `json:"replies"` // Replies to current story.
	Type                 string                   `json:"type"`
	Video                Video                    `json:"video"`
	Views                int                      `json:"views"` // Views number.
	ClickableStickers    StoriesClickableStickers `json:"clickable_stickers"`
	TrackCode            string                   `json:"track_code"`
	LikesCount           int                      `json:"likes_count"`
	NarrativeID          int                      `json:"narrative_id"`
	NarrativeOwnerID     int                      `json:"narrative_owner_id"`
	NarrativeInfo        StoriesNarrativeInfo     `json:"narrative_info"`
	NarrativesCount      int                      `json:"narratives_count"`
	FirstNarrativeTitle  string                   `json:"first_narrative_title"`
	Questions            StoriesQuestions         `json:"questions"`
	ReactionSetID        string                   `json:"reaction_set_id"`
}

type StoriesFeedItem struct {
	Type           constants.StoriesFeedItemType `json:"type"`
	ID             string                        `json:"id"`
	Stories        []StoriesStory                `json:"stories"`
	Grouped        constants.StoriesFeedItemType `json:"grouped"`
	App            AppsApp                       `json:"app"`
	BirthdayUserID int                           `json:"birthday_user_id"`
	TrackCode      string                        `json:"track_code"`
	HasUnseen      NumberFlagBool                `json:"has_unseen"`
	Name           string                        `json:"name"`
	PromoData      StoriesPromoBlock             `json:"promo_data"`
}

// StoriesClickableStickers struct.
//
// The field clickable_stickers is available in the history objects.
// The sticker objects is pasted by the developer on the client himself, only
// coordinates are transmitted to the server.
//
// https://dev.vk.com/ru/reference/objects/clickable-sticker
type StoriesClickableStickers struct {
	OriginalWidth     int                       `json:"original_width"`
	OriginalHeight    int                       `json:"original_height"`
	ClickableStickers []StoriesClickableSticker `json:"clickable_stickers"`
}

// NewClickableStickers return new StoriesClickableStickers.
//
// Requires the width and height of the original photo or video.
func NewClickableStickers(width, height int) *StoriesClickableStickers {
	return &StoriesClickableStickers{
		OriginalWidth:     width,
		OriginalHeight:    height,
		ClickableStickers: []StoriesClickableSticker{},
	}
}

// AddMention add mention sticker.
//
// Mention should be in the format of a VK mentioning, for example: [id1|name] or [club1|name].
func (cs *StoriesClickableStickers) AddMention(mention string, area []StoriesClickablePoint) *StoriesClickableStickers {
	cs.ClickableStickers = append(cs.ClickableStickers, StoriesClickableSticker{
		Type:          constants.ClickableStickerMention,
		ClickableArea: area,
		Mention:       mention,
	})

	return cs
}

// AddHashtag add hashtag sticker.
//
// Hashtag must necessarily begin with the symbol #.
func (cs *StoriesClickableStickers) AddHashtag(hashtag string, area []StoriesClickablePoint) *StoriesClickableStickers {
	cs.ClickableStickers = append(cs.ClickableStickers, StoriesClickableSticker{
		Type:          constants.ClickableStickerHashtag,
		ClickableArea: area,
		Hashtag:       hashtag,
	})

	return cs
}

func (cs StoriesClickableStickers) ToJSON() string {
	b, _ := json.Marshal(cs)
	return string(b)
}

type StoriesClickableSticker struct { //nolint: maligned
	ID            int                     `json:"id"`
	Type          string                  `json:"type"`
	ClickableArea []StoriesClickablePoint `json:"clickable_area"`
	Style         string                  `json:"style,omitempty"`

	// type=post
	PostOwnerID int `json:"post_owner_id,omitempty"`
	PostID      int `json:"post_id,omitempty"`

	// type=sticker
	StickerID     int `json:"sticker_id,omitempty"`
	StickerPackID int `json:"sticker_pack_id,omitempty"`

	// type=place or geo
	PlaceID int `json:"place_id,omitempty"`
	// Title
	CategoryID int `json:"category_id,omitempty"`

	// type=question
	Question               string         `json:"question,omitempty"`
	QuestionButton         string         `json:"question_button,omitempty"`
	QuestionDefaultPrivate NumberFlagBool `json:"question_default_private,omitempty"`
	Color                  string         `json:"color,omitempty"`

	// type=mention
	Mention string `json:"mention,omitempty"`

	// type=hashtag
	Hashtag string `json:"hashtag,omitempty"`

	// type=link
	LinkObject     Link   `json:"link_object,omitempty"`
	TooltipText    string `json:"tooltip_text,omitempty"`
	TooltipTextKey string `json:"tooltip_text_key,omitempty"`

	// type=time
	TimestampMs int64  `json:"timestamp_ms,omitempty"`
	Date        string `json:"date,omitempty"`
	Title       string `json:"title,omitempty"`

	// type=market_item
	Subtype string `json:"subtype,omitempty"`
	// LinkObject Link         `json:"link_object,omitempty"` // subtype=aliexpress_product
	MarketItem MarketItem `json:"market_item,omitempty"` // subtype=market_item

	// type=story_reply
	OwnerID int `json:"owner_id,omitempty"`
	StoryID int `json:"story_id,omitempty"`

	// type=owner
	// OwnerID int `json:"owner_id,omitempty"`

	// type=poll
	Poll PollsPoll `json:"poll,omitempty"`

	// type=music
	Audio          Audio `json:"audio,omitempty"`
	AudioStartTime int   `json:"audio_start_time,omitempty"`

	// type=app
	App                      AppsApp        `json:"app,omitempty"`
	AppContext               string         `json:"app_context,omitempty"`
	HasNewInteractions       NumberFlagBool `json:"has_new_interactions,omitempty"`
	IsBroadcastNotifyAllowed NumberFlagBool `json:"is_broadcast_notify_allowed,omitempty"`

	// type=emoji
	Emoji string `json:"emoji,omitempty"`

	// type=text
	Text            string `json:"text,omitempty"`
	BackgroundStyle string `json:"background_style,omitempty"`
	Alignment       string `json:"alignment,omitempty"`
	SelectionColor  string `json:"selection_color,omitempty"`
}

// StoriesClickablePoint struct.
type StoriesClickablePoint struct {
	X int `json:"x"`
	Y int `json:"y"`
}
