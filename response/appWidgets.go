package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/appWidgets
type AppWidgetsGetAppImageUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type AppWidgetsGetAppImagesResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.AppWidgetsImage `json:"items"`
	} `json:"response"`
}

type AppWidgetsGetGroupImageUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type AppWidgetsGetGroupImagesResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.AppWidgetsImage `json:"items"`
	} `json:"response"`
}

type AppWidgetsGetImagesByIDResponse struct {
	BaseResponse
	Response objects.AppWidgetsImage `json:"response"`
}

type AppWidgetsSaveAppImageResponse struct {
	BaseResponse
	Response objects.AppWidgetsImage `json:"response"`
}

type AppWidgetsSaveGroupImageResponse struct {
	BaseResponse
	Response objects.AppWidgetsImage `json:"response"`
}

type AppWidgetsUpdateResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AppWidgetsAppImageUploadResponse struct {
	BaseResponse
	Response struct {
		Image string `json:"image"`
		Hash  string `json:"hash"`
	} `json:"response"`
}

type AppWidgetsGroupImageUploadResponse struct {
	BaseResponse
	Response struct {
		Image string `json:"image"`
		Hash  string `json:"hash"`
	} `json:"response"`
}
