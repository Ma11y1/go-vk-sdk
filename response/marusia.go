package response

import (
	"encoding/json"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/marusia/getting-started

type MarusiaGetPictureUploadLinkResponse struct {
	BaseResponse
	Response struct {
		PictureUploadLink string `json:"picture_upload_link"` // Link
	} `json:"response"`
}

type MarusiaSavePictureResponse struct {
	BaseResponse
	Response struct {
		AppID   int `json:"app_id"`
		PhotoID int `json:"photo_id"`
	} `json:"response"`
}

type MarusiaGetPicturesResponse struct {
	BaseResponse
	Response struct {
		Count int                      `json:"count"`
		Items []objects.MarusiaPicture `json:"items"`
	} `json:"response"`
}

type MarusiaDeletePictureResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarusiaGetAudioUploadLinkResponse struct {
	BaseResponse
	Response struct {
		AudioUploadLink string `json:"audio_upload_link"` // Link
	} `json:"response"`
}

type MarusiaCreateAudioResponse struct {
	BaseResponse
	Response struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"response"`
}

type MarusiaGetAudiosResponse struct {
	BaseResponse
	Response struct {
		Count  int                    `json:"count"`
		Audios []objects.MarusiaAudio `json:"audios"`
	} `json:"response"`
}

type MarusiaDeleteAudioResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarusiaUploadPictureResponse struct {
	Hash        string          `json:"hash"`
	Photo       json.RawMessage `json:"photo"`
	Server      int             `json:"server"`
	AID         int             `json:"aid"`
	MessageCode int             `json:"message_code"`
}

type MarusiaUploadAudioResponse struct {
	Sha       string                   `json:"sha"`
	Secret    string                   `json:"secret"`
	Meta      objects.MarusiaAudioMeta `json:"meta"`
	Hash      string                   `json:"hash"`
	Server    string                   `json:"server"`
	UserID    int                      `json:"user_id"`
	RequestID string                   `json:"request_id"`
}
