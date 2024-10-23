package objects

type FaveTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FavePage struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	User        UserFull  `json:"user"`
	Group       GroupFull `json:"group"`
	Tags        []FaveTag `json:"tags"`
	UpdatedDate int       `json:"updated_date"`
}

type FaveFavesLink struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Photo       Photo          `json:"photo"`
	Caption     string         `json:"caption"`
	IsFavorite  NumberFlagBool `json:"is_favorite"`
}

type FaveItem struct {
	Type      string         `json:"type"`
	Seen      NumberFlagBool `json:"seen"`
	AddedDate int            `json:"added_date"`
	Tags      []FaveTag      `json:"tags"`
	Link      FaveFavesLink  `json:"link,omitempty"`
	Post      WallWallpost   `json:"post,omitempty"`
	Video     Video          `json:"video,omitempty"`
	Product   MarketItem     `json:"product,omitempty"`
	Article   Article        `json:"article,omitempty"`
}
