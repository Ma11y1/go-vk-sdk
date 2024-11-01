package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Doc: https://dev.vk.com/ru/reference/objects

// Basic structs
type (
	JSONObject interface {
		ToJSON() string
	}

	IDTitle struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	KeyValue struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	IDName struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	Count struct {
		Count int `json:"count"`
	}

	UserID struct {
		UserID int `json:"user_id"`
	}

	RequestMutual struct {
		Count int   `json:"count"`
		Users []int `json:"users"`
	}

	// BoolInt type. Used to convert JSON string values to boolean type when decoding JSON
	BoolInt bool
)

// UnmarshalJSON func called when JSON is decoded.
func (b *BoolInt) UnmarshalJSON(data []byte) (err error) {
	switch {
	case bytes.Equal(data, []byte("1")), bytes.Equal(data, []byte("true")):
		*b = true
	case bytes.Equal(data, []byte("0")), bytes.Equal(data, []byte("false")):
		*b = false
	default:
		// return json errors
		err = &json.UnmarshalTypeError{
			Value: string(data),
			Type:  reflect.TypeOf((*BoolInt)(nil)),
		}
	}

	return
}

func (b *BoolInt) MarshalJSON() (data []byte, err error) {
	if *b {
		return []byte("1"), nil
	} else {
		return []byte("0"), nil
	}
}

// Base api structs
type Attachment interface {
	ToAttachment() string
}

type City IDTitle
type Country IDTitle

type CommentsInfo struct {
	Count         int     `json:"count"`
	CanPost       BoolInt `json:"can_post"`
	GroupsCanPost BoolInt `json:"groups_can_post"`
	CanClose      BoolInt `json:"can_close"`
	CanOpen       BoolInt `json:"can_open"`
}

// Geo related structs
type (
	Geo struct {
		Coordinates string `json:"coordinates"`
		Place       Place  `json:"place"`
		Showmap     int    `json:"showmap"`
		Type        string `json:"type"`
	}

	MessageGeo struct {
		Coordinates GeoCoordinates `json:"coordinates"`
		Place       Place          `json:"place"`
		Showmap     int            `json:"showmap"`
		Type        string         `json:"type"`
	}

	GeoCoordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)

type Image struct {
	Type   string  `json:"type"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	ID     int     `json:"id"`
	Theme  string  `json:"theme"` // light, dark
}

// renamedBaseImage Used during decoding to convert two properties url and src to url
type renamedBaseImage struct {
	Type   string  `json:"type"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Src    string  `json:"src"`
}

// UnmarshalJSON is required to support images with `src` field.
func (obj *Image) UnmarshalJSON(data []byte) (err error) {
	var renamedObj renamedBaseImage

	err = json.Unmarshal(data, &renamedObj)
	if err != nil {
		return fmt.Errorf("objects.Image(): %w", err)
	}

	obj.Height = renamedObj.Height
	obj.Width = renamedObj.Width
	obj.Type = renamedObj.Type

	if renamedObj.Src == "" {
		obj.URL = renamedObj.URL
	} else {
		obj.URL = renamedObj.Src
	}

	return nil
}

func maxSizeImage(images *[]Image) (maxImage *Image) {
	var maxSize float64
	for _, image := range *images {
		if size := image.Height * image.Width; size > maxSize {
			maxSize = size
			maxImage = &image
		}
	}
	return
}

func minSizeImage(images *[]Image) (minImage *Image) {
	var minSize float64
	for _, image := range *images {
		if size := image.Height * image.Width; size < minSize || minSize == 0 {
			minSize = size
			minImage = &image
		}
	}
	return
}

// LikesInfo structs
type (
	Likes struct {
		UserLikes BoolInt `json:"user_likes"` // Information whether current user likes
		Count     int     `json:"count"`      // Likes number
	}

	LikesInfo struct {
		CanLike    BoolInt `json:"can_like"`    // Information whether current user can like the post
		CanPublish BoolInt `json:"can_publish"` // Information whether current user can repost
		UserLikes  BoolInt `json:"user_likes"`  // Information whether current uer has liked the post
		Count      int     `json:"count"`       // Likes number
	}
)

// Link structs
type (
	Link struct {
		Application  LinkApplication `json:"application"`
		Button       LinkButton      `json:"button"`
		ButtonText   string          `json:"button_text"`
		ButtonAction string          `json:"button_action"`
		Caption      string          `json:"caption"`
		Description  string          `json:"description"`
		Photo        Photo           `json:"photo"`
		Video        Video           `json:"video"`
		PreviewPage  string          `json:"preview_page"`
		PreviewURL   string          `json:"preview_url"`
		Product      LinkProduct     `json:"product"`
		Rating       LinkRating      `json:"rating"`
		Title        string          `json:"title"`
		Target       string          `json:"target"`
		Text         string          `json:"text"`
		URL          string          `json:"url"`
		IsFavorite   BoolInt         `json:"is_favorite"`
	}

	LinkApplication struct {
		AppID float64              `json:"app_id"`
		Store LinkApplicationStore `json:"store"`
	}

	LinkApplicationStore struct {
		ID   float64 `json:"id"`
		Name string  `json:"name"`
	}

	LinkButton struct {
		Action LinkButtonAction `json:"action"`
		Title  string           `json:"title"`
	}

	LinkButtonAction struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	}

	LinkProduct struct {
		Price       MarketPrice `json:"price"`
		Merchant    string      `json:"merchant"`
		OrdersCount int         `json:"orders_count"`
	}

	LinkRating struct {
		ReviewsCount json.Number `json:"reviews_count"`
		Stars        float64     `json:"stars"`
	}
)

type Place struct {
	Address        string         `json:"address"`
	Checkins       int            `json:"checkins"`
	City           interface{}    `json:"city"`
	Country        interface{}    `json:"country"`
	Created        int            `json:"created"`
	ID             int            `json:"id"`
	Icon           string         `json:"icon"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	Title          string         `json:"title"`
	Type           string         `json:"type"`
	IsDeleted      BoolInt        `json:"is_deleted"`
	TotalCheckins  int            `json:"total_checkins"`
	Updated        int            `json:"updated"`
	CategoryObject CategoryObject `json:"category_object"`
}

type CategoryObject struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Icons []Image `json:"icons"`
}

type RepostsInfo struct {
	Count        int `json:"count"`
	WallCount    int `json:"wall_count"`
	MailCount    int `json:"mail_count"`
	UserReposted int `json:"user_reposted"`
}

type Sticker struct {
	Images               []Image            `json:"images"`
	ImagesWithBackground []Image            `json:"images_with_background"`
	ProductID            int                `json:"product_id"`
	StickerID            int                `json:"sticker_id"`
	IsAllowed            bool               `json:"is_allowed"`
	AnimationURL         string             `json:"animation_url"`
	Animations           []StickerAnimation `json:"animations"`
	InnerType            string             `json:"inner_type"`
}

type StickerAnimation struct {
	Type string `json:"type"` // light, dark
	URL  string `json:"url"`
}

// MaxSize return the largest Sticker.
func (s *Sticker) MaxSize() (maxImageSize *Image) {
	return maxSizeImage(&s.Images)
}

// MinSize return the smallest Sticker.
func (s *Sticker) MinSize() (minImageSize *Image) {
	return minSizeImage(&s.Images)
}

// MaxSizeBackground return the largest Sticker with background.
func (s *Sticker) MaxSizeBackground() (maxImageSize *Image) {
	return maxSizeImage(&s.ImagesWithBackground)
}

// MinSizeBackground return the smallest Sticker with background.
func (s *Sticker) MinSizeBackground() (minImageSize *Image) {
	return minSizeImage(&s.ImagesWithBackground)
}

// Privacy It allows the user or developer to determine which users or owners can see or interact with existing content
type Privacy struct {
	Category string `json:"category,omitempty"`
	Lists    struct {
		Allowed  []int `json:"allowed"`
		Excluded []int `json:"excluded"`
	} `json:"lists,omitempty"`
	Owners struct {
		Allowed  []int `json:"allowed"`
		Excluded []int `json:"excluded"`
	} `json:"owners,omitempty"`
}

type EventsEventAttach struct {
	Address      string  `json:"address,omitempty"`
	ButtonText   string  `json:"button_text"`
	Friends      []int   `json:"friends"`
	ID           int     `json:"id"`
	IsFavorite   BoolInt `json:"is_favorite"`
	MemberStatus int     `json:"member_status,omitempty"`
	Text         string  `json:"text"`
	Time         int     `json:"time,omitempty"`
}

// Article Represents an article published on VKontakte and is used to work with the article api on the platform
type Article struct {
	ID            int     `json:"id"`
	OwnerID       int     `json:"owner_id"`
	OwnerName     string  `json:"owner_name"`
	OwnerPhoto    string  `json:"owner_photo"`
	State         string  `json:"state"`
	CanReport     BoolInt `json:"can_report"`
	IsFavorite    BoolInt `json:"is_favorite"`
	NoFooter      BoolInt `json:"no_footer"`
	Title         string  `json:"title"`
	Subtitle      string  `json:"subtitle"`
	Views         int     `json:"views"`
	Shares        int     `json:"shares"`
	URL           string  `json:"url"`
	ViewURL       string  `json:"view_url"`
	AccessKey     string  `json:"access_key"`
	PublishedDate int     `json:"published_date"`
	Photo         Photo   `json:"photo"`
}

type UsersAndGroups struct {
	Profiles []UserFull  `json:"profiles,omitempty"`
	Groups   []GroupFull `json:"groups,omitempty"`
}

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       BoolInt  `json:"keyboard"`
	InlineKeyboard BoolInt  `json:"inline_keyboard"`
	Carousel       BoolInt  `json:"carousel"`
	LangID         int      `json:"lang_id"`
}
