package objects

import (
	"fmt"
)

type Video struct {
	// Video access key.
	AccessKey string `json:"access_key"`

	// Date when the video has been added in Unixtime.
	AddingDate int `json:"adding_date"`

	// Date when the video has been released in Unixtime.
	ReleaseDate int `json:"release_date"`

	// Information whether current user can add the video.
	CanAdd BoolInt `json:"can_add"`

	// Information whether current user can add the video to faves.
	CanAddToFaves BoolInt `json:"can_add_to_faves"`

	// Information whether current user can comment the video.
	CanComment BoolInt `json:"can_comment"`

	// Information whether current user can edit the video.
	CanEdit BoolInt `json:"can_edit"`

	// Information whether current user can like the video.
	CanLike BoolInt `json:"can_like"`

	// Information whether current user can download the video.
	CanDownload int `json:"can_download"`

	// Information whether current user can repost this video.
	CanRepost         BoolInt      `json:"can_repost"`
	CanSubscribe      BoolInt      `json:"can_subscribe"`
	CanAttachLink     BoolInt      `json:"can_attach_link"`
	IsFavorite        BoolInt      `json:"is_favorite"`
	IsPrivate         BoolInt      `json:"is_private"`
	IsExplicit        BoolInt      `json:"is_explicit"`
	IsSubscribed      BoolInt      `json:"is_subscribed"`
	Added             BoolInt      `json:"added"`
	Repeat            BoolInt      `json:"repeat"` // Information whether the video is repeated
	ContentRestricted int          `json:"content_restricted"`
	Live              BoolInt      `json:"live"` // Returns if the video is a live stream
	Upcoming          BoolInt      `json:"upcoming"`
	Comments          int          `json:"comments"`    // Number of comments
	Date              int          `json:"date"`        // Date when video has been uploaded in Unixtime
	Description       string       `json:"description"` // Video description
	Duration          int          `json:"duration"`    // Video duration in seconds
	Files             VideoFiles   `json:"files"`
	Trailer           VideoFiles   `json:"trailer,omitempty"`
	FirstFrame        []VideoImage `json:"first_frame"`
	Image             []VideoImage `json:"image"`
	Height            int          `json:"height"`   // Video height
	ID                int          `json:"id"`       // Video ID
	OwnerID           int          `json:"owner_id"` // Video owner ID
	UserID            int          `json:"user_id"`
	Photo130          string       `json:"photo_130"`  // url of the preview image with 130 px in width
	Photo320          string       `json:"photo_320"`  // url of the preview image with 320 px in width
	Photo640          string       `json:"photo_640"`  // url of the preview image with 640 px in width
	Photo800          string       `json:"photo_800"`  // url of the preview image with 800 px in width
	Photo1280         string       `json:"photo_1280"` // url of the preview image with 1280 px in width

	// url of the page with a player that can be used to play the video in the browser.
	Player                   string            `json:"player"`
	Processing               int               `json:"processing"` // Returns if the video is processing
	Title                    string            `json:"title"`      // Video title
	Subtitle                 string            `json:"subtitle"`   // Video subtitle
	Type                     string            `json:"type"`
	Views                    int               `json:"views"` // Number of views
	Width                    int               `json:"width"` // Video width
	Platform                 string            `json:"platform"`
	LocalViews               int               `json:"local_views"`
	Likes                    LikesInfo         `json:"likes"`   // Count of likes
	Reposts                  RepostsInfo       `json:"reposts"` // Count of views
	TrackCode                string            `json:"track_code"`
	PrivacyView              Privacy           `json:"privacy_view"`
	PrivacyComment           Privacy           `json:"privacy_comment"`
	ActionButton             VideoActionButton `json:"action_button"`
	Restriction              VideoRestriction  `json:"restriction"`
	ContentRestrictedMessage string            `json:"content_restricted_message"`
	MainArtists              []AudioArtist     `json:"main_artists"`
	FeaturedArtists          []AudioArtist     `json:"featured_artists"`
	Genres                   []IDName          `json:"genres"`
	OvID                     string            `json:"ov_id,omitempty"`
}

func (video *Video) ToAttachment() string {
	return fmt.Sprintf("video%d_%d", video.OwnerID, video.ID)
}

type VideoRestriction struct {
	Title          string  `json:"title"`
	Text           string  `json:"text"`
	AlwaysShown    BoolInt `json:"always_shown"`
	Blur           BoolInt `json:"blur"`
	CanPlay        BoolInt `json:"can_play"`
	CanPreview     BoolInt `json:"can_preview"`
	CardIcon       []Image `json:"card_icon"`
	ListIcon       []Image `json:"list_icon"`
	DisclaimerType int     `json:"disclaimer_type"`
}

type VideoActionButton struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	URL     string       `json:"url"`
	Snippet VideoSnippet `json:"snippet"`
}

type VideoSnippet struct {
	Description string  `json:"description"`
	OpenTitle   string  `json:"open_title"`
	Title       string  `json:"title"`
	TypeName    string  `json:"type_name"`
	Date        int     `json:"date"`
	Image       []Image `json:"image"`
}

type VideoFiles struct {
	External     string `json:"external,omitempty"` // url of the external player
	Mp4_1080     string `json:"mp4_1080,omitempty"` // url of the mpeg4 file with 1080p quality
	Mp4_1440     string `json:"mp4_1440,omitempty"` // url of the mpeg4 file with 2k quality
	Mp4_2160     string `json:"mp4_2160,omitempty"` // url of the mpeg4 file with 4k quality
	Mp4_240      string `json:"mp4_240,omitempty"`  // url of the mpeg4 file with 240p quality
	Mp4_360      string `json:"mp4_360,omitempty"`  // url of the mpeg4 file with 360p quality
	Mp4_480      string `json:"mp4_480,omitempty"`  // url of the mpeg4 file with 480p quality
	Mp4_720      string `json:"mp4_720,omitempty"`  // url of the mpeg4 file with 720p quality
	Live         string `json:"live,omitempty"`
	HLS          string `json:"hls,omitempty"`
	DashUni      string `json:"dash_uni,omitempty"`
	DashSep      string `json:"dash_sep,omitempty"`
	DashWebm     string `json:"dash_webm,omitempty"`
	FailoverHost string `json:"failover_host,omitempty"`
}

type VideoCatBlock struct {
	CanHide BoolInt           `json:"can_hide"`
	ID      int               `json:"id"`
	Items   []VideoCatElement `json:"items"`
	Name    string            `json:"name"`
	Next    string            `json:"next"`
	Type    string            `json:"type"`
	View    string            `json:"view"`
}

type VideoCatElement struct {
	CanAdd      BoolInt `json:"can_add"`
	CanEdit     BoolInt `json:"can_edit"`
	IsPrivate   BoolInt `json:"is_private"`
	Comments    int     `json:"comments"`
	Count       int     `json:"count"`
	Date        int     `json:"date"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	ID          int     `json:"id"`
	OwnerID     int     `json:"owner_id"`
	Photo130    string  `json:"photo_130"`
	Photo160    string  `json:"photo_160"`
	Photo320    string  `json:"photo_320"`
	Photo640    string  `json:"photo_640"`
	Photo800    string  `json:"photo_800"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	UpdatedTime int     `json:"updated_time"`
	Views       int     `json:"views"`
}

type VideoSaveResult struct {
	Description string `json:"description"` // Video description
	OwnerID     int    `json:"owner_id"`    // Video owner ID
	Title       string `json:"title"`       // Video title
	UploadURL   string `json:"upload_url"`  // url for the video uploading
	VideoID     int    `json:"video_id"`    // Video ID
	AccessKey   string `json:"access_key"`  // Video access key
}

type VideoUploadResponse struct {
	Size    int `json:"size"`
	VideoID int `json:"video_id"`
}

type VideoAlbum struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
}

type VideoAlbumFull struct {
	Count       int          `json:"count"`        // Total number of videos in album
	ID          int          `json:"id"`           // Album ID
	Image       []VideoImage `json:"image"`        // Album cover image in different sizes
	IsSystem    BoolInt      `json:"is_system"`    // Information whether album is system
	OwnerID     int          `json:"owner_id"`     // Album owner's ID
	Photo160    string       `json:"photo_160"`    // url of the preview image with 160px in width
	Photo320    string       `json:"photo_320"`    // url of the preview image with 320px in width
	Title       string       `json:"title"`        // Album title
	UpdatedTime int          `json:"updated_time"` // Date when the album has been updated last time in Unixtime
	ImageBlur   int          `json:"image_blur"`
	Privacy     Privacy      `json:"privacy"`
	CanDelete   BoolInt      `json:"can_delete"`
	CanEdit     BoolInt      `json:"can_edit"`
	CanUpload   BoolInt      `json:"can_upload"`
}

type VideoFull struct {
	AccessKey     string     `json:"access_key"`  // Video access key
	AddingDate    int        `json:"adding_date"` // Date when the video has been added in Unixtime
	IsFavorite    BoolInt    `json:"is_favorite"`
	CanAdd        BoolInt    `json:"can_add"`     // Information whether current user can add the video
	CanComment    BoolInt    `json:"can_comment"` // Information whether current user can comment the video
	CanEdit       BoolInt    `json:"can_edit"`    // Information whether current user can edit the video
	CanRepost     BoolInt    `json:"can_repost"`  // Information whether current user can comment the video
	CanLike       BoolInt    `json:"can_like"`
	CanAddToFaves BoolInt    `json:"can_add_to_faves"`
	Repeat        BoolInt    `json:"repeat"`      // Information whether the video is repeated
	Comments      int        `json:"comments"`    // Number of comments
	Date          int        `json:"date"`        // Date when video has been uploaded in Unixtime
	Description   string     `json:"description"` // Video description
	Duration      int        `json:"duration"`    // Video duration in seconds
	Files         VideoFiles `json:"files"`
	Trailer       VideoFiles `json:"trailer"`
	ID            int        `json:"id"` // Video ID
	Likes         Likes      `json:"likes"`
	Live          int        `json:"live"`     // Returns if the video is live translation
	OwnerID       int        `json:"owner_id"` // Video owner ID

	// url of the page with a player that can be used to play the video in the browser.
	Player     string       `json:"player"`
	Processing int          `json:"processing"` // Returns if the video is processing
	Title      string       `json:"title"`      // Video title
	Views      int          `json:"views"`      // Number of views
	Width      int          `json:"width"`
	Height     int          `json:"height"`
	Image      []VideoImage `json:"image"`
	FirstFrame []VideoImage `json:"first_frame"`
	Added      int          `json:"added"`
	Type       string       `json:"type"`
	Reposts    RepostsInfo  `json:"reposts"`
}

func (video *VideoFull) ToAttachment() string {
	return fmt.Sprintf("video%d_%d", video.OwnerID, video.ID)
}

type VideoTag struct {
	Date       int     `json:"date"`
	ID         int     `json:"id"`
	PlacerID   int     `json:"placer_id"`
	TaggedName string  `json:"tagged_name"`
	UserID     int     `json:"user_id"`
	Viewed     BoolInt `json:"viewed"`
}

type VideoTagInfo struct {
	AccessKey   string     `json:"access_key"`
	AddingDate  int        `json:"adding_date"`
	CanAdd      BoolInt    `json:"can_add"`
	CanEdit     BoolInt    `json:"can_edit"`
	Comments    int        `json:"comments"`
	Date        int        `json:"date"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	Files       VideoFiles `json:"files"`
	ID          int        `json:"id"`
	Live        int        `json:"live"`
	OwnerID     int        `json:"owner_id"`
	Photo130    string     `json:"photo_130"`
	Photo320    string     `json:"photo_320"`
	Photo800    string     `json:"photo_800"`
	PlacerID    int        `json:"placer_id"`
	Player      string     `json:"player"`
	Processing  int        `json:"processing"`
	TagCreated  int        `json:"tag_created"`
	TagID       int        `json:"tag_id"`
	Title       string     `json:"title"`
	Views       int        `json:"views"`
}

type VideoImage struct {
	Image
	WithPadding BoolInt `json:"with_padding"`
}

type VideoLive struct {
	OwnerID     int             `json:"owner_id"`
	VideoID     int             `json:"video_id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	AccessKey   string          `json:"access_key"`
	Stream      VideoLiveStream `json:"stream"`
}

type VideoLiveStream struct {
	URL     string `json:"url"`
	Key     string `json:"key"`
	OKMPURL string `json:"okmp_url"`
}

type VideoLiveCategory struct {
	ID      int                 `json:"id"`
	Label   string              `json:"label"`
	Sublist []VideoLiveCategory `json:"sublist,omitempty"`
}
