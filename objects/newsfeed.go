package objects

type NewsfeedEventActivity struct {
	Address      string `json:"address"`       // address of event
	ButtonText   string `json:"button_text"`   // text of attach
	Friends      []int  `json:"friends"`       // array of friends ids
	MemberStatus int    `json:"member_status"` // Current user's member status
	Text         string `json:"text"`          // text of attach
	Time         int    `json:"time"`          // event start time
}

type NewsfeedItemAudio struct {
	Audio NewsfeedItemAudioAudio `json:"audio"`
}

type NewsfeedItemAudioAudio struct {
	Count int     `json:"count"` // Audios number
	Items []Audio `json:"items"`
}

type NewsfeedItemDigest struct {
	ButtonText  string         `json:"button_text"`
	FeedID      string         `json:"feed_id"` // id of feed in digest
	Items       []WallWallpost `json:"items"`
	MainPostIDs []string       `json:"main_post_ids"`
	Template    string         `json:"template"` // type of digest
	Title       string         `json:"title"`
	TrackCode   string         `json:"track_code"`
}

type NewsfeedItemFriend struct {
	Friends NewsfeedItemFriendFriends `json:"friends"`
}

type NewsfeedItemFriendFriends struct {
	Count int      `json:"count"` // Number of friends has been added
	Items []UserID `json:"items"`
}

type NewsfeedItemNote struct {
	Notes NewsfeedItemNoteNotes `json:"notes"`
}

type NewsfeedItemNoteNotes struct {
	Count int                    `json:"count"` // Notes number
	Items []NewsfeedNewsfeedNote `json:"items"`
}

type NewsfeedItemPhoto struct {
	Photos NewsfeedItemPhotoPhotos `json:"photos"`
}

type NewsfeedItemPhotoPhotos struct {
	Count int               `json:"count"` // Photos number
	Items []PhotosPhotoFull `json:"items"`
}

type NewsfeedItemPhotoTag struct {
	PhotoTags NewsfeedItemPhotoTagPhotoTags `json:"photo_tags"`
}

type NewsfeedItemPhotoTagPhotoTags struct {
	Count int               `json:"count"` // Tags number
	Items []PhotosPhotoFull `json:"items"`
}

type NewsfeedItemStoriesBlock struct {
	BlockType string         `json:"block_type"`
	Stories   []StoriesStory `json:"stories"`
}

type NewsfeedItemTopic struct{}

type NewsfeedItemVideo struct {
	Video NewsfeedItemVideoVideo `json:"video"`
}

type NewsfeedItemVideoVideo struct {
	Count int     `json:"count"` // Tags number
	Items []Video `json:"items"`
}

type NewsfeedItemWallpost struct {
	Activity       NewsfeedEventActivity    `json:"activity"`
	Attachments    []WallWallpostAttachment `json:"attachments"`
	Comments       CommentsInfo             `json:"comments"`
	FromID         int                      `json:"from_id"`
	CopyHistory    []WallWallpost           `json:"copy_history"`
	Geo            Geo                      `json:"geo"`
	Likes          LikesInfo                `json:"likes"`
	PostSource     WallPostSource           `json:"post_source"`
	PostType       string                   `json:"post_type"`
	Reposts        RepostsInfo              `json:"reposts"`
	MarkedAsAds    int                      `json:"marked_as_ads,omitempty"`
	Views          interface{}              `json:"views,omitempty"` // Views int or wallViews
	IsFavorite     BoolInt                  `json:"is_favorite,omitempty"`
	CanDelete      BoolInt                  `json:"can_delete"`
	CanArchive     BoolInt                  `json:"can_archive"`
	IsArchived     BoolInt                  `json:"is_archived"`
	SignerID       int                      `json:"signer_id,omitempty"`
	Text           string                   `json:"text"` // Post text
	Copyright      WallPostCopyright        `json:"copyright"`
	CategoryAction NewsfeedCategoryAction   `json:"category_action"`
}

type NewsfeedCategoryAction struct {
	Action struct {
		Target string `json:"target"`
		Type   string `json:"type"`
		URL    string `json:"url"`
	} `json:"action"`
	Name string `json:"name"`
}

type NewsfeedList struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// NewsfeedItemMarket struct.
type NewsfeedItemMarket struct {
	MarketItem
}

// NewsfeedNewsfeedItem struct.
type NewsfeedNewsfeedItem struct {
	Type     string `json:"type"`
	SourceID int    `json:"source_id"`
	Date     int    `json:"date"`
	TopicID  int    `json:"topic_id"`

	PostID int `json:"post_id,omitempty"`

	NewsfeedItemWallpost
	NewsfeedItemPhoto
	NewsfeedItemPhotoTag
	NewsfeedItemFriend
	NewsfeedItemNote
	NewsfeedItemAudio
	NewsfeedItemTopic
	NewsfeedItemVideo
	NewsfeedItemDigest
	NewsfeedItemStoriesBlock
	NewsfeedItemMarket

	CreatedBy        int     `json:"created_by,omitempty"`
	CanEdit          BoolInt `json:"can_edit,omitempty"`
	CanDelete        BoolInt `json:"can_delete,omitempty"`
	CanDoubtCategory BoolInt `json:"can_doubt_category"`
	CanSetCategory   BoolInt `json:"can_set_category"`
}

type NewsfeedNewsfeedNote struct {
	Comments int    `json:"comments"`
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Title    string `json:"title"`
}
