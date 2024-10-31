package objects

import (
	"fmt"
)

type Photo struct {
	AccessKey          string             `json:"access_key"` // Access key for the photo
	AlbumID            int                `json:"album_id"`   // Album ID
	Date               int                `json:"date"`       // Date when uploaded
	Height             int                `json:"height"`     // Original photo height
	ID                 int                `json:"id"`         // Photo ID
	Images             []PhotosImage      `json:"images"`
	Lat                float64            `json:"lat"`      // Latitude
	Long               float64            `json:"long"`     // Longitude
	OwnerID            int                `json:"owner_id"` // Photo owner's ID
	PostID             int                `json:"post_id"`  // Post ID
	Text               string             `json:"text"`     // Photo caption
	UserID             int                `json:"user_id"`  // ID of the user who have uploaded the photo
	Width              int                `json:"width"`    // Original photo width
	CanUpload          BoolInt            `json:"can_upload"`
	CommentsDisabled   BoolInt            `json:"comments_disabled"`
	ThumbIsLast        BoolInt            `json:"thumb_is_last"`
	UploadByAdminsOnly BoolInt            `json:"upload_by_admins_only"`
	HasTags            BoolInt            `json:"has_tags"`
	Created            int                `json:"created"`
	Description        string             `json:"description"`
	PrivacyComment     []string           `json:"privacy_comment"`
	PrivacyView        []string           `json:"privacy_view"`
	Size               int                `json:"size"`
	Sizes              []PhotosPhotoSizes `json:"sizes"`
	ThumbID            int                `json:"thumb_id"`
	ThumbSrc           string             `json:"thumb_src"`
	Title              string             `json:"title"`
	Updated            int                `json:"updated"`
	Color              string             `json:"color"`
}

func (photo *Photo) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

func (photo *Photo) MaxSize() (maxPhotoSize PhotosPhotoSizes) {
	var maxSize float64

	for _, photoSize := range photo.Sizes {
		size := photoSize.Height * photoSize.Width
		if size > maxSize {
			maxSize = size
			maxPhotoSize = photoSize
		}
	}

	return
}

func (photo *Photo) MinSize() (minPhotoSize PhotosPhotoSizes) {
	var minSize float64

	for _, photoSize := range photo.Sizes {
		size := photoSize.Height * photoSize.Width
		if size < minSize || minSize == 0 {
			minSize = size
			minPhotoSize = photoSize
		}
	}

	return
}

type PhotosCommentXtrPid struct {
	Attachments    []WallCommentAttachment `json:"attachments"`
	Date           int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromID         int                     `json:"from_id"` // Author ID
	ID             int                     `json:"id"`      // Comment ID
	Likes          LikesInfo               `json:"likes"`
	ParentsStack   []int                   `json:"parents_stack"`
	Pid            int                     `json:"pid"`              // Photo ID
	ReplyToComment int                     `json:"reply_to_comment"` // Replied comment ID
	ReplyToUser    int                     `json:"reply_to_user"`    // Replied user ID
	Text           string                  `json:"text"`             // Comment text
	Thread         WallWallCommentThread   `json:"thread"`
}

type PhotosImage struct {
	Image
	Type string `json:"type"`
}

type PhotosPhotoAlbum struct {
	Created     int    `json:"created"`     // Date when the album has been created in Unixtime
	Description string `json:"description"` // Photo album description
	ID          int    `json:"id"`          // Photo album ID
	OwnerID     int    `json:"owner_id"`    // Album owner's ID
	Size        int    `json:"size"`        // Photos number
	Thumb       Photo  `json:"thumb"`
	Title       string `json:"title"`   // Photo album title
	Updated     int    `json:"updated"` // Date when the album has been updated last time in Unixtime
}

func (album *PhotosPhotoAlbum) ToAttachment() string {
	return fmt.Sprintf("album%d_%d", album.OwnerID, album.ID)
}

type PhotosPhotoAlbumFull struct {
	// Information whether current user can upload photo to the album.
	CanUpload        BoolInt            `json:"can_upload"`
	CommentsDisabled BoolInt            `json:"comments_disabled"` // Information whether album comments are disabled
	Created          int                `json:"created"`           // Date when the album has been created in Unixtime
	Description      string             `json:"description"`       // Photo album description
	ID               int                `json:"id"`                // Photo album ID
	OwnerID          int                `json:"owner_id"`          // Album owner's ID
	Size             int                `json:"size"`              // Photos number
	PrivacyComment   Privacy            `json:"privacy_comment"`
	PrivacyView      Privacy            `json:"privacy_view"`
	Sizes            []PhotosPhotoSizes `json:"sizes"`
	ThumbID          int                `json:"thumb_id"` // Thumb photo ID

	// Information whether the album thumb is last photo.
	ThumbIsLast int    `json:"thumb_is_last"`
	ThumbSrc    string `json:"thumb_src"` // url of the thumb image
	Title       string `json:"title"`     // Photo album title

	// Date when the album has been updated last time in Unixtime.
	Updated int `json:"updated"`

	// Information whether only community administrators can upload photos.
	UploadByAdminsOnly int `json:"upload_by_admins_only"`
}

func (album *PhotosPhotoAlbumFull) ToAttachment() string {
	return fmt.Sprintf("album%d_%d", album.OwnerID, album.ID)
}

func (album *PhotosPhotoAlbumFull) MaxSize() (maxPhotoSize PhotosPhotoSizes) {
	var maxSize float64

	for _, photoSize := range album.Sizes {
		size := photoSize.Height * photoSize.Width
		if size > maxSize {
			maxSize = size
			maxPhotoSize = photoSize
		}
	}

	return
}

func (album *PhotosPhotoAlbumFull) MinSize() (minPhotoSize PhotosPhotoSizes) {
	var minSize float64

	for _, photoSize := range album.Sizes {
		size := photoSize.Height * photoSize.Width
		if size < minSize || minSize == 0 {
			minSize = size
			minPhotoSize = photoSize
		}
	}

	return
}

type PhotosPhotoFull struct {
	AccessKey  string             `json:"access_key"`  // Access key for the photo
	AlbumID    int                `json:"album_id"`    // Album ID
	CanComment BoolInt            `json:"can_comment"` // Information whether current user can comment the photo
	CanRepost  BoolInt            `json:"can_repost"`  // Information whether current user can repost the photo
	HasTags    BoolInt            `json:"has_tags"`
	Comments   Count              `json:"comments"`
	Date       int                `json:"date"`   // Date when uploaded
	Height     int                `json:"height"` // Original photo height
	ID         int                `json:"id"`     // Photo ID
	Images     []PhotosImage      `json:"images"`
	Lat        float64            `json:"lat"` // Latitude
	Likes      Likes              `json:"likes"`
	Long       float64            `json:"long"`     // Longitude
	OwnerID    int                `json:"owner_id"` // Photo owner's ID
	PostID     int                `json:"post_id"`  // Post ID
	Reposts    RepostsInfo        `json:"reposts"`
	Tags       Count              `json:"tags"`
	Text       string             `json:"text"`       // Photo caption
	UserID     int                `json:"user_id"`    // ID of the user who have uploaded the photo
	Width      int                `json:"width"`      // Original photo width
	Hidden     int                `json:"hidden"`     // Returns if the photo is hidden above the wall
	Photo75    string             `json:"photo_75"`   // url of image with 75 px width
	Photo130   string             `json:"photo_130"`  // url of image with 130 px width
	Photo604   string             `json:"photo_604"`  // url of image with 604 px width
	Photo807   string             `json:"photo_807"`  // url of image with 807 px width
	Photo1280  string             `json:"photo_1280"` // url of image with 1280 px width
	Photo2560  string             `json:"photo_2560"` // url of image with 2560 px width
	Sizes      []PhotosPhotoSizes `json:"sizes"`
	OrigPhoto  PhotosPhotoSizes   `json:"orig_photo"`
}

func (photo *PhotosPhotoFull) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

func (photo *PhotosPhotoFull) MaxSize() (maxPhotoSize PhotosPhotoSizes) {
	var maxSize float64

	for _, photoSize := range photo.Sizes {
		size := photoSize.Height * photoSize.Width
		if size > maxSize {
			maxSize = size
			maxPhotoSize = photoSize
		}
	}

	return
}

func (photo *PhotosPhotoFull) MinSize() (minPhotoSize PhotosPhotoSizes) {
	var minSize float64

	for _, photoSize := range photo.Sizes {
		size := photoSize.Height * photoSize.Width
		if size < minSize || minSize == 0 {
			minSize = size
			minPhotoSize = photoSize
		}
	}

	return
}

type PhotosPhotoFullXtrRealOffset struct {
	PhotosPhotoFull
	RealOffset int `json:"real_offset"` // Real position of the photo
}

// PhotosPhotoSizes
// BUG(VK): json: cannot unmarshal number 180.000000 into Go struct field PhotosPhotoSizes.height of type int
type PhotosPhotoSizes struct {
	Image
}

type PhotosPhotoTag struct {
	Date        int     `json:"date"`        // Date when tag has been added in Unixtime
	ID          int     `json:"id"`          // Tag ID
	PlacerID    int     `json:"placer_id"`   // ID of the tag creator
	TaggedName  string  `json:"tagged_name"` // Tag description
	Description string  `json:"description"` // Tagged description.
	UserID      int     `json:"user_id"`     // Tagged user ID
	Viewed      BoolInt `json:"viewed"`      // Information whether the tag is reviewed
	X           float64 `json:"x"`           // Coordinate X of the left upper corner
	X2          float64 `json:"x2"`          // Coordinate X of the right lower corner
	Y           float64 `json:"y"`           // Coordinate Y of the left upper corner
	Y2          float64 `json:"y2"`          // Coordinate Y of the right lower corner
}

type PhotosPhotoUpload struct {
	AlbumID   int    `json:"album_id"`   // Album ID
	UploadURL string `json:"upload_url"` // url to upload photo
	UserID    int    `json:"user_id"`    // OAuthUser ID
}

type PhotosPhotoUploadResponse struct {
	AID        int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

type PhotosPhotoXtrRealOffset struct {
	Photo
	RealOffset int `json:"real_offset"` // Real position of the photo
}

type PhotosPhotoXtrTagInfo struct {
	Photo
	TagCreated int `json:"tag_created"` // Date when tag has been added in Unixtime
	TagID      int `json:"tag_id"`      // Tag ID
}
