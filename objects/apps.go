package objects

import "go-vk-sdk/constants"

type AppsCatalogBanner struct {
	Description      string `json:"description"`
	TitleColor       string `json:"title_color"`
	BackgroundColor  string `json:"background_color"`
	DescriptionColor string `json:"description_color"`
}

type AppsApp struct {
	ID                    int                             `json:"id"`
	AuthorOwnerID         int                             `json:"author_owner_id"`
	AuthorURL             string                          `json:"author_url"`
	Banner1120            string                          `json:"banner_1120"`
	Banner560             string                          `json:"banner_560"`
	CatalogPosition       int                             `json:"catalog_position"`
	Description           string                          `json:"description"`
	Friends               []int                           `json:"friends"`
	Genre                 string                          `json:"genre"`
	GenreID               int                             `json:"genre_id"`
	Icon139               string                          `json:"icon_139"`
	Icon150               string                          `json:"icon_150"`
	Icon278               string                          `json:"icon_278"`
	Icon75                string                          `json:"icon_75"`
	International         NumberFlagBool                  `json:"international"`
	IsInCatalog           NumberFlagBool                  `json:"is_in_catalog"`
	Installed             NumberFlagBool                  `json:"installed"`
	PushEnabled           NumberFlagBool                  `json:"push_enabled"`
	HideTabbar            NumberFlagBool                  `json:"hide_tabbar"`
	IsNew                 NumberFlagBool                  `json:"is_new"`
	New                   NumberFlagBool                  `json:"new"`
	IsInstalled           NumberFlagBool                  `json:"is_installed"`
	HasVkConnect          NumberFlagBool                  `json:"has_vk_connect"`
	LeaderboardType       int                             `json:"leaderboard_type"`
	MembersCount          int                             `json:"members_count"`
	PlatformID            int                             `json:"platform_id"`
	PublishedDate         int                             `json:"published_date"`
	ScreenName            string                          `json:"screen_name"`
	Screenshots           []Photo                         `json:"screenshots"`
	Section               string                          `json:"section"`
	Title                 string                          `json:"title"`
	Type                  string                          `json:"type"`
	Icon16                string                          `json:"icon_16"`
	Icon576               string                          `json:"icon_576"`
	ScreenOrientation     constants.AppsScreenOrientation `json:"screen_orientation"`
	CatalogBanner         AppsCatalogBanner               `json:"catalog_banner"`
	MobileControlsType    int                             `json:"mobile_controls_type"`
	MobileViewSupportType int                             `json:"mobile_view_support_type"`
}

type AppsLeaderboard struct {
	Level  int `json:"level"`
	Points int `json:"points"`
	Score  int `json:"score"`
	UserID int `json:"user_id"`
}

type AppsScope struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type AppsTestingGroup struct {
	GroupID   int      `json:"group_id"`
	UserIDs   []int    `json:"user_ids"`
	Name      string   `json:"name"`
	Webview   string   `json:"webview"`
	Platforms []string `json:"platforms"`
}
