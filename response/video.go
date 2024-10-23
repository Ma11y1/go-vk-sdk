package response

import (
	"go-vk-sdk/errors"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/method/video

type VideoAddResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoAddAlbumResponse struct {
	BaseResponse
	Response struct {
		AlbumID int `json:"album_id"`
	} `json:"response"`
}

type VideoAddToAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoCreateCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoDeleteAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoDeleteCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoEditResponse struct {
	BaseResponse
	Response struct {
		Success   objects.NumberFlagBool `json:"success"`
		AccessKey string                 `json:"access_key"`
	} `json:"response"`
}

type VideoEditAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoEditCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoGetResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"`
		Items []objects.Video `json:"items"`
	} `json:"response"`
}

type VideoGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"`
		Items []objects.Video `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type VideoGetAlbumByIDResponse struct {
	BaseResponse
	Response objects.VideoAlbumFull `json:"response"`
}

type VideoGetAlbumsResponse struct {
	BaseResponse
	Response struct {
		Count int                  `json:"count"`
		Items []objects.VideoAlbum `json:"items"`
	} `json:"response"`
}

type VideoGetAlbumsExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                      `json:"count"`
		Items []objects.VideoAlbumFull `json:"items"`
	} `json:"response"`
}

type VideoGetAlbumsByVideoResponse struct {
	BaseResponse
	Response []int `json:"response"`
}

type VideoGetAlbumsByVideoExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                      `json:"count"`
		Items []objects.VideoAlbumFull `json:"items"`
	} `json:"response"`
}

type VideoGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.WallWallComment `json:"items"`
	} `json:"response"`
}

type VideoGetCommentsExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.WallWallComment `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type VideoLiveGetCategoriesResponse struct {
	BaseResponse
	Response []objects.VideoLiveCategory `json:"response"`
}

type VideoRemoveFromAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoReorderAlbumsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoReorderVideosResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoReportResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoReportCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoRestoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoRestoreCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type VideoSaveResponse struct {
	BaseResponse
	Response objects.VideoSaveResult `json:"response"`
}

type VideoSearchResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"`
		Items []objects.Video `json:"items"`
	} `json:"response"`
}

type VideoSearchExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int             `json:"count"`
		Items []objects.Video `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type VideoStartStreamingResponse struct {
	BaseResponse
	Response objects.VideoLive `json:"response"`
}

type VideoStopStreamingResponse struct {
	BaseResponse
	Response struct {
		UniqueViewers int `json:"unique_viewers"`
	} `json:"response"`
}

type VideosUploadVideoStoriesResponse struct {
	Response struct {
		UploadResult string `json:"upload_result"`
	} `json:"response"`
	errors.UploadAPIError
}
