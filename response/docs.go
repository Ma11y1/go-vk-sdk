package response

import "go-vk-sdk/objects"

// https://dev.vk.com/ru/method/docs

type DocsAddResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type DocsDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type DocsEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type DocsGetResponse struct {
	BaseResponse
	Response struct {
		BaseResponse
		Response int `json:"response"`
	} `json:"response"`
}

type DocsGetByIDResponse struct {
	BaseResponse
	Response []objects.Document `json:"response"`
}

type DocsGetMessagesUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type DocsGetTypesResponse struct {
	BaseResponse
	Response struct {
		Count int                     `json:"count"`
		Items []objects.DocumentTypes `json:"items"`
	} `json:"response"`
}

type DocsGetUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type DocsGetWallUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type DocsSaveResponse struct {
	BaseResponse
	Response struct {
		Type         string                       `json:"type"`
		AudioMessage objects.MessagesAudioMessage `json:"audio_message"`
		Doc          objects.Document             `json:"doc"`
		Graffiti     objects.MessagesGraffiti     `json:"graffiti"`
	} `json:"response"`
}

type DocsSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.Document `json:"items"`
	} `json:"response"`
}
