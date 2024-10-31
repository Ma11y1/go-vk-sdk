package objects

type WidgetsCommentMedia struct {
	ItemID   int    `json:"item_id"`
	OwnerID  int    `json:"owner_id"`
	ThumbSrc string `json:"thumb_src"`
	Type     string `json:"type"`
}

type WidgetsCommentReplies struct {
	CanPost       BoolInt                     `json:"can_post"`
	GroupsCanPost BoolInt                     `json:"groups_can_post"`
	Count         int                         `json:"count"`
	Replies       []WidgetsCommentRepliesItem `json:"replies"`
}

type WidgetsCommentRepliesItem struct {
	Cid   int                `json:"cid"`
	Date  int                `json:"date"`
	Likes WidgetsWidgetLikes `json:"likes"`
	Text  string             `json:"text"`
	UID   int                `json:"uid"`
	User  UserFull           `json:"user"`
}

type WidgetsWidgetComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	CanDelete   BoolInt                 `json:"can_delete"`
	IsFavorite  BoolInt                 `json:"is_favorite"`
	Comments    WidgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Likes       LikesInfo               `json:"likes"`
	Media       WidgetsCommentMedia     `json:"media"`
	PostType    string                  `json:"post_type"`
	Reposts     RepostsInfo             `json:"reposts"`
	Text        string                  `json:"text"`
	ToID        int                     `json:"to_id"`
	PostSource  WallPostSource          `json:"post_source"`
	Views       struct {
		Count int `json:"count"`
	} `json:"views"`
	Donut         WallWallpostDonut  `json:"donut"`
	ShortTextRate float64            `json:"short_text_rate"`
	Header        WallWallpostHeader `json:"header"`
}

type WidgetsWidgetLikes struct {
	Count int `json:"count"`
}

type WidgetsWidgetPage struct {
	Comments    WidgetsWidgetLikes `json:"comments,omitempty"`
	Date        int                `json:"date,omitempty"`
	Description string             `json:"description,omitempty"`
	ID          int                `json:"id,omitempty"`
	Likes       WidgetsWidgetLikes `json:"likes,omitempty"`
	PageID      string             `json:"page_id,omitempty"`
	Photo       string             `json:"photo,omitempty"`
	Title       string             `json:"title,omitempty"`
	URL         string             `json:"url,omitempty"`
}
