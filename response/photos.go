package response

import (
	"go-vk-sdk/errors"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/method/photos

type PhotosConfirmTagResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosCopyResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosCreateAlbumResponse struct {
	BaseResponse
	Response objects.PhotosPhotoAlbumFull `json:"response"`
}

type PhotosCreateCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosDeleteAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosDeleteCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosEditAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosEditCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosGetResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"` // Total number
		Items []objects.Photo `json:"items"`
	} `json:"response"`
}

type PhotosGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"` // Total number
		Items []objects.PhotosPhotoFull `json:"items"`
	} `json:"response"`
}

type PhotosGetAlbumsResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"` // Total number
		Items []objects.PhotosPhotoAlbumFull `json:"items"`
	} `json:"response"`
}

type PhotosGetAlbumsCountResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosGetAllResponse struct {
	BaseResponse
	Response struct {
		Count int                                `json:"count"` // Total number
		Items []objects.PhotosPhotoXtrRealOffset `json:"items"`
		More  objects.BoolInt                    `json:"more"` // Information whether next page is presented
	} `json:"response"`
}

type PhotosGetAllExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                                    `json:"count"` // Total number
		Items []objects.PhotosPhotoFullXtrRealOffset `json:"items"`
		More  objects.BoolInt                        `json:"more"` // Information whether next page is presented
	} `json:"response"`
}

type PhotosGetAllCommentsResponse struct {
	BaseResponse
	Response struct {
		Count int                           `json:"count"` // Total number
		Items []objects.PhotosCommentXtrPid `json:"items"`
	} `json:"response"`
}

type PhotosGetByIDResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosGetByIDExtendedResponse struct {
	BaseResponse
	Response []objects.PhotosPhotoFull `json:"response"`
}

type PhotosGetChatUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PhotosGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Count      int                       `json:"count"`       // Total number
		RealOffset int                       `json:"real_offset"` // Real offset of the comments
		Items      []objects.WallWallComment `json:"items"`
	} `json:"response"`
}

type PhotosGetCommentsExtendedResponse struct {
	BaseResponse
	Response struct {
		Count      int                       `json:"count"`       // Total number
		RealOffset int                       `json:"real_offset"` // Real offset of the comments
		Items      []objects.WallWallComment `json:"items"`
		Profiles   []objects.UserFull        `json:"profiles"`
		Groups     []objects.GroupFull       `json:"groups"`
	} `json:"response"`
}

type PhotosGetMarketAlbumUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PhotosGetMarketUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PhotosGetMessagesUploadServerResponse struct {
	BaseResponse
	Response struct {
		AlbumID   int    `json:"album_id"`
		UploadURL string `json:"upload_url"`
		UserID    int    `json:"user_id,omitempty"`
		GroupID   int    `json:"group_id,omitempty"`
	} `json:"response"`
}

type PhotosGetNewTagsResponse struct {
	BaseResponse
	Response struct {
		Count int                             `json:"count"` // Total number
		Items []objects.PhotosPhotoXtrTagInfo `json:"items"`
	} `json:"response"`
}

type PhotosGetOwnerCoverPhotoUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PhotosGetOwnerPhotoUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PhotosGetTagsResponse struct {
	BaseResponse
	Response []objects.PhotosPhotoTag `json:"response"`
}

type PhotosGetUploadServerResponse struct {
	BaseResponse
	Response objects.PhotosPhotoUpload `json:"response"`
}

type PhotosGetUserPhotosResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"` // Total number
		Items []objects.Photo `json:"items"`
	} `json:"response"`
}

type PhotosGetUserPhotosExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"` // Total number
		Items []objects.PhotosPhotoFull `json:"items"`
	} `json:"response"`
}

type PhotosGetWallUploadServerResponse struct {
	BaseResponse
	Response objects.PhotosPhotoUpload `json:"response"`
}

type PhotosMakeCoverResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosMoveResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosPutTagResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosRemoveTagResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosReorderAlbumsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosReorderPhotosResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosReportResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosReportCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosRestoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosRestoreCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PhotosSaveResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosSaveMarketAlbumPhotoResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosSaveMarketPhotoResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosSaveMessagesPhotoResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosSaveOwnerCoverPhotoResponse struct {
	BaseResponse
	Response struct {
		Images []objects.PhotosImage `json:"images"`
	} `json:"response"`
}

type PhotosSaveOwnerPhotoResponse struct {
	BaseResponse
	Response struct {
		PhotoHash     string `json:"photo_hash"`
		PhotoSrc      string `json:"photo_src"`
		PhotoSrcBig   string `json:"photo_src_big"`
		PhotoSrcSmall string `json:"photo_src_small"`
		Saved         int    `json:"saved"`
		PostID        int    `json:"post_id"`
	} `json:"response"`
}

type PhotosSaveWallPhotoResponse struct {
	BaseResponse
	Response []objects.Photo `json:"response"`
}

type PhotosSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"` // Total number
		Items []objects.PhotosPhotoFull `json:"items"`
	} `json:"response"`
}

type PhotosUploadPhotoAlbumResponse struct {
	AID        int    `json:"aid"`         // Album ID
	Hash       string `json:"hash"`        // Uploading hash
	PhotosList string `json:"photos_list"` // Uploaded photos data
	Server     int    `json:"server"`      // Upload server number
}

type PhotosUploadPhotoWallResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

type PhotosUploadPhotoOwnerResponse struct {
	MID         int    `json:"mid"`
	Hash        string `json:"hash"`   // Uploading hash
	Photo       string `json:"photo"`  // Uploaded photo data
	Server      int    `json:"server"` // Upload server number
	MessageCode int    `json:"message_code"`
	ProfileAID  int    `json:"profile_aid"`
}

type PhotosUploadPhotoMessageResponse struct {
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

type PhotosUploadPhotoChatResponse struct {
	Response string `json:"response"` // Uploaded photo data
}

type PhotosUploadPhotoMarketResponse struct {
	CropData string `json:"crop_data"` // Crop data
	CropHash string `json:"crop_hash"` // Crop hash
	GroupID  int    `json:"group_id"`  // Community ID
	Hash     string `json:"hash"`      // Uploading hash
	Photo    string `json:"photo"`     // Uploaded photo data
	Server   int    `json:"server"`    // Upload server number
}

type PhotosUploadPhotoMarketAlbumResponse struct {
	GID    int    `json:"gid"`    // Community ID
	Hash   string `json:"hash"`   // Uploading hash
	Photo  string `json:"photo"`  // Uploaded photo data
	Server int    `json:"server"` // Upload server number
}

type PhotosUploadVideoResponse errors.UploadAPIError

type PhotosUploadDocumentResponse errors.UploadAPIError

type PhotosUploadPhotoStoriesResponse struct {
	Response struct {
		UploadResult string `json:"upload_result"`
		Sig          string `json:"_sig"`
	} `json:"response"`
	Error struct {
		ErrorCode int    `json:"error_code"`
		Type      string `json:"type"`
	} `json:"error"`
}
