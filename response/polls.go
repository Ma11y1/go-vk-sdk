package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/polls

type PollsAddVoteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PollsCreateResponse struct {
	BaseResponse
	Response objects.PollsPoll `json:"response"`
}

type PollsDeleteVoteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PollsEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PollsGetBackgroundsResponse struct {
	BaseResponse
	Response []objects.PollsBackground `json:"response"`
}

type PollsGetByIDResponse struct {
	BaseResponse
	Response objects.PollsPoll `json:"response"`
}

type PollsGetPhotoUploadServerResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"`
	} `json:"response"`
}

type PollsGetVotersResponse struct {
	BaseResponse
	Response []objects.PollsVoters `json:"response"`
}

type PollsGetVotersFieldsResponse struct {
	BaseResponse
	Response []objects.PollsVotersFields `json:"response"`
}

type PollsSavePhotoResponse struct {
	BaseResponse
	Response objects.PollsPhoto `json:"response"`
}

type PollsUploadPhotoResponse struct {
	Photo string `json:"photo"` // Uploaded photo data
	Hash  string `json:"hash"`  // Uploaded hash
}
