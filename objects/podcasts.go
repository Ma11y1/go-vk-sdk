package objects

type PodcastsSearchExternalData struct {
	OwnerName string  `json:"owner_name"`
	OwnerURL  string  `json:"owner_url"`
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Cover     []Image `json:"cover"`
}

type PodcastsItem struct {
	OwnerID int `json:"owner_id"`
}

type PodcastsCategory struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Cover []Image `json:"cover"`
}

type PodcastsEpisode struct {
	ID                  int                 `json:"id"`
	OwnerID             int                 `json:"owner_id"`
	Artist              string              `json:"artist"`
	Title               string              `json:"title"`
	Duration            int                 `json:"duration"`
	Date                int                 `json:"date"`
	URL                 string              `json:"url"`
	LyricsID            int                 `json:"lyrics_id"`
	NoSearch            int                 `json:"no_search"`
	TrackCode           string              `json:"track_code"`
	IsHq                NumberFlagBool      `json:"is_hq"`
	IsFocusTrack        NumberFlagBool      `json:"is_focus_track"`
	IsExplicit          NumberFlagBool      `json:"is_explicit"`
	ShortVideosAllowed  NumberFlagBool      `json:"short_videos_allowed"`
	StoriesAllowed      NumberFlagBool      `json:"stories_allowed"`
	StoriesCoverAllowed NumberFlagBool      `json:"stories_cover_allowed"`
	PodcastInfo         PodcastsPodcastInfo `json:"podcast_info"`
}

type PodcastsPodcastInfo struct {
	Cover struct {
		Sizes []Image `json:"cover"`
	}
	Plays       int            `json:"plays"`
	IsFavorite  NumberFlagBool `json:"is_favorite"`
	Description string         `json:"description"`
	Position    int            `json:"position"`
}
