package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/leadForms

type LeadFormsCreateResponse struct {
	BaseResponse
	Response struct {
		FormID int    `json:"form_id"`
		URL    string `json:"url"`
	} `json:"response"`
}

type LeadFormsDeleteResponse struct {
	BaseResponse
	Response struct {
		FormID int `json:"form_id"`
	} `json:"response"`
}

type LeadFormsGetResponse struct {
	BaseResponse
	Response objects.LeadFormsForm `json:"response"`
}

type LeadFormsGetLeadsResponse struct {
	BaseResponse
	Response struct {
		Leads []objects.LeadFormsLead `json:"leads"`
	} `json:"response"`
}

type LeadFormsListResponse struct {
	BaseResponse
	Response []objects.LeadFormsForm `json:"response"`
}

type LeadFormsGetUploadURLResponse struct {
	BaseResponse
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}

type LeadFormsUpdateResponse struct {
	BaseResponse
	Response struct {
		FormID int    `json:"form_id"`
		URL    string `json:"url"`
	} `json:"response"`
}

type LeadFormsUploadPhotoResponse struct {
	Photo   string `json:"photo"`
	ErrCode int    `json:"errcode"`
}
