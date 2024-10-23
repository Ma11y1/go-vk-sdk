package objects

import (
	"fmt"
)

type Audio struct {
	ID                  int            `json:"id"`
	OwnerID             int            `json:"owner_id"`
	URL                 string         `json:"url"`
	Title               string         `json:"title"`
	Subtitle            string         `json:"subtitle"`
	Artist              string         `json:"artist"`
	Date                int            `json:"date"`
	Duration            int            `json:"duration"`
	AccessKey           string         `json:"access_key"`
	IsHq                NumberFlagBool `json:"is_hq"`
	IsExplicit          NumberFlagBool `json:"is_explicit"`
	StoriesAllowed      NumberFlagBool `json:"stories_allowed"`
	ShortVideosAllowed  NumberFlagBool `json:"short_videos_allowed"`
	IsFocusTrack        NumberFlagBool `json:"is_focus_track"`
	IsLicensed          NumberFlagBool `json:"is_licensed"`
	StoriesCoverAllowed NumberFlagBool `json:"stories_cover_allowed"`
	LyricsID            int            `json:"lyrics_id"`
	AlbumID             int            `json:"album_id"`
	GenreID             int            `json:"genre_id"`
	TrackCode           string         `json:"track_code"`
	NoSearch            int            `json:"no_search"`
	MainArtists         []AudioArtist  `json:"main_artists"`
	Ads                 AudioAds       `json:"ads"`
}

func (a *Audio) ToAttachment() string {
	return fmt.Sprintf("a%d_%d", a.OwnerID, a.ID)
}

type AudioAds struct {
	ContentID      string `json:"content_id"`
	PUID1          string `json:"puid1"`
	PUID22         string `json:"puid22"`
	Duration       string `json:"duration"`
	AccountAgeType string `json:"account_age_type"`
}

type AudioArtist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type AudioLyrics struct {
	LyricsID int    `json:"lyrics_id"`
	Text     string `json:"text"`
}
