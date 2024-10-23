package objects

type WallAppPost struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type WallAttachedNote struct {
	Comments     int    `json:"comments"`
	Date         int    `json:"date"`
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	ReadComments int    `json:"read_comments"`
	Title        string `json:"title"`
	ViewURL      string `json:"view_url"`
}

type WallCommentAttachment struct {
	Audio             Audio             `json:"audio"`
	Doc               Document          `json:"doc"`
	Link              Link              `json:"link"`
	Market            MarketItem        `json:"market"`
	MarketMarketAlbum MarketAlbum       `json:"market_market_album"`
	Note              WallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             Photo             `json:"photo"`
	Sticker           Sticker           `json:"sticker"`
	Type              string            `json:"type"`
	Video             Video             `json:"video"`
	Graffiti          WallGraffiti      `json:"graffiti"`
}

type WallGraffiti struct {
	ID        int    `json:"id"`
	OwnerID   int    `json:"owner_id"`
	Photo200  string `json:"photo_200"`
	Photo586  string `json:"photo_586"`
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	AccessKey string `json:"access_key"`
}

type WallPostSource struct {
	Link     Link   `json:"link"`
	Data     string `json:"data"`
	Platform string `json:"platform"`
	Type     string `json:"type"`
	URL      string `json:"url"`
}

type WallPostedPhoto struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type WallViews struct {
	Count int `json:"count"`
}

type WallWallCommentThread struct {
	Count           int               `json:"count"`
	Items           []WallWallComment `json:"items"`
	CanPost         NumberFlagBool    `json:"can_post"`
	GroupsCanPost   NumberFlagBool    `json:"groups_can_post"`
	ShowReplyButton NumberFlagBool    `json:"show_reply_button"`
}

type WallWallComment struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`
	Deleted        NumberFlagBool          `json:"deleted"`
	FromID         int                     `json:"from_id"`
	ID             int                     `json:"id"`
	Likes          LikesInfo               `json:"likes"`
	RealOffset     int                     `json:"real_offset"`
	ReplyToComment int                     `json:"reply_to_comment"`
	ReplyToUser    int                     `json:"reply_to_user"`
	Text           string                  `json:"text"`
	PostID         int                     `json:"post_id"`
	PostOwnerID    int                     `json:"post_owner_id"`
	PhotoID        int                     `json:"photo_id"`
	PhotoOwnerID   int                     `json:"photo_owner_id"`
	VideoID        int                     `json:"video_id"`
	VideoOwnerID   int                     `json:"video_owner_id"`
	ItemID         int                     `json:"item_id"`
	MarketOwnerID  int                     `json:"market_owner_id"`
	ParentsStack   []int                   `json:"parents_stack"`
	OwnerID        int                     `json:"owner_id"`
	Thread         WallWallCommentThread   `json:"thread"`
	Donut          WallWallCommentDonut    `json:"donut"`
}

type WallWallCommentDonut struct {
	IsDonut     NumberFlagBool `json:"is_donut"`
	Placeholder string         `json:"placeholder"`
}

type WallWallpost struct {
	AccessKey      string                   `json:"access_key"`
	ID             int                      `json:"id"`
	OwnerID        int                      `json:"owner_id"`
	FromID         int                      `json:"from_id"`
	CreatedBy      int                      `json:"created_by"`
	Date           int                      `json:"date"`
	Text           string                   `json:"text"`
	ReplyOwnerID   int                      `json:"reply_owner_id"`
	ReplyPostID    int                      `json:"reply_post_id"`
	FriendsOnly    int                      `json:"friends_only"`
	Comments       CommentsInfo             `json:"comments"`
	Likes          LikesInfo                `json:"likes"`
	Reposts        RepostsInfo              `json:"reposts"`
	Views          WallViews                `json:"views"`
	PostType       string                   `json:"post_type"`
	PostSource     WallPostSource           `json:"post_source"`
	Attachments    []WallWallpostAttachment `json:"attachments"`
	Geo            Geo                      `json:"geo"`
	SignerID       int                      `json:"signer_id"`
	CopyHistory    []WallWallpost           `json:"copy_history"`
	CanPin         NumberFlagBool           `json:"can_pin"`
	CanDelete      NumberFlagBool           `json:"can_delete"`
	CanEdit        NumberFlagBool           `json:"can_edit"`
	IsPinned       NumberFlagBool           `json:"is_pinned"`
	IsFavorite     NumberFlagBool           `json:"is_favorite"`
	IsArchived     NumberFlagBool           `json:"is_archived"`
	IsDeleted      NumberFlagBool           `json:"is_deleted"`
	MarkedAsAds    NumberFlagBool           `json:"marked_as_ads"`
	Edited         int                      `json:"edited"`
	Copyright      WallPostCopyright        `json:"copyright"`
	PostID         int                      `json:"post_id"`
	PostponedID    int                      `json:"postponed_id"`
	ParentsStack   []int                    `json:"parents_stack"`
	Donut          WallWallpostDonut        `json:"donut"`
	ShortTextRate  float64                  `json:"short_text_rate"`
	CarouselOffset int                      `json:"carousel_offset"`
	Header         WallWallpostHeader       `json:"header"`
	Hash           string                   `json:"hash"`
}

type WallWallpostAttachment struct {
	AccessKey         string            `json:"access_key"`
	Album             PhotosPhotoAlbum  `json:"album"`
	App               WallAppPost       `json:"app"`
	Audio             Audio             `json:"audio"`
	Doc               Document          `json:"doc"`
	Event             EventsEventAttach `json:"event"`
	Graffiti          WallGraffiti      `json:"graffiti"`
	Link              Link              `json:"link"`
	Market            MarketItem        `json:"market"`
	MarketMarketAlbum MarketAlbum       `json:"market_market_album"`
	Note              WallAttachedNote  `json:"note"`
	Page              PagesWikipageFull `json:"page"`
	Photo             Photo             `json:"photo"`
	PhotosList        []string          `json:"photos_list"`
	Poll              PollsPoll         `json:"poll"`
	PostedPhoto       WallPostedPhoto   `json:"posted_photo"`
	Type              string            `json:"type"`
	Video             Video             `json:"video"`
	Podcast           PodcastsEpisode   `json:"podcast"`
}

type WallWallpostToID struct {
	Attachments   []WallWallpostAttachment `json:"attachments"`
	Comments      CommentsInfo             `json:"comments"`
	CopyOwnerID   int                      `json:"copy_owner_id"`
	CopyPostID    int                      `json:"copy_post_id"`
	Date          int                      `json:"date"`
	FromID        int                      `json:"from_id"`
	Geo           Geo                      `json:"geo"`
	ID            int                      `json:"id"`
	Likes         LikesInfo                `json:"likes"`
	PostID        int                      `json:"post_id"`
	PostSource    WallPostSource           `json:"post_source"`
	PostType      string                   `json:"post_type"`
	Reposts       RepostsInfo              `json:"reposts"`
	SignerID      int                      `json:"signer_id"`
	Text          string                   `json:"text"`
	ToID          int                      `json:"to_id"`
	IsFavorite    NumberFlagBool           `json:"is_favorite"`
	MarkedAsAds   NumberFlagBool           `json:"marked_as_ads"`
	ParentsStack  []int                    `json:"parents_stack"`
	Donut         WallWallpostDonut        `json:"donut"`
	ShortTextRate float64                  `json:"short_text_rate"`
	Views         WallViews                `json:"views"`
	Header        WallWallpostHeader       `json:"header"`
}

type WallWallpostDonut struct {
	IsDonut            NumberFlagBool `json:"is_donut"`
	CanPublishFreeCopy NumberFlagBool `json:"can_publish_free_copy"`
	PaidDuration       int            `json:"paid_duration"`
	EditMode           string         `json:"edit_mode"`
	Durations          []IDName       `json:"durations"`
}

type WallPostCopyright struct {
	ID   int    `json:"id,omitempty"`
	Link string `json:"link"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type WallWallpostHeader struct {
	Type              string                              `json:"type"`
	CustomDescription WallWallpostHeaderCustomDescription `json:"custom_description"`
}

type WallWallpostHeaderCustomDescription struct {
	SourceID int `json:"source_id"`
	Date     int `json:"date"`
}
