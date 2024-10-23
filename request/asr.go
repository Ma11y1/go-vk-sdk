package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/asr

// AsrCheckStatusRequest defines the request for asr.checkStatus
//
//	The method checks the status of the audio recording processing task and returns a text transcript of the audio recording.
//	Doc: https://dev.vk.com/method/asr.checkStatus
type AsrCheckStatusRequest struct {
	BaseRequest
}

// NewAsrCheckStatusRequest creates a new request for asr.checkStatus
func NewAsrCheckStatusRequest(a *api.API, actor actor.Actor) *AsrCheckStatusRequest {
	return &AsrCheckStatusRequest{*NewMethodBaseRequest(a, actor, "asr.checkStatus")}
}

// Exec executes the request and unmarshals the response into AsrCheckStatusResponse
func (r *AsrCheckStatusRequest) Exec(ctx context.Context) (response response.AsrCheckStatusResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// TaskID required method define parameter task_id.
//
//	Identifier of the created task for processing audio recordings in UUID format.
//	https://dev.vk.com/ru/method/asr.process
func (r *AsrCheckStatusRequest) TaskID(id string) *AsrCheckStatusRequest {
	r.parameters.Set(constants.ParameterNameTaskID, id)
	return r
}

// AsrGetUploadURLRequest defines the request for asr.getUploadUrl
//
//	The method returns a link to the server address for downloading the audio recording.
//	The link is available for 24 hours.
//	Doc: https://dev.vk.com/method/asr.getUploadUrl
type AsrGetUploadURLRequest struct {
	BaseRequest
}

// NewAsrGetUploadURLRequest creates a new request for asr.getUploadUrl
func NewAsrGetUploadURLRequest(a *api.API, actor actor.Actor) *AsrGetUploadURLRequest {
	return &AsrGetUploadURLRequest{*NewMethodBaseRequest(a, actor, "asr.getUploadUrl")}
}

// Exec executes the request and unmarshals the response into AsrGetUploadURLResponse
func (r *AsrGetUploadURLRequest) Exec(ctx context.Context) (response response.AsrGetUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AsrProcessRequest defines the request for asr.process
//
//	The method performs speech recognition from the loaded audio file.
//	Total maximum duration of audio recording files per day: 100 minutes.
//	Doc: https://dev.vk.com/method/asr.process
type AsrProcessRequest struct {
	BaseRequest
}

// NewAsrProcessRequest creates a new request for asr.process
func NewAsrProcessRequest(a *api.API, actor actor.Actor) *AsrProcessRequest {
	return &AsrProcessRequest{*NewMethodBaseRequest(a, actor, "asr.process")}
}

// Exec executes the request and unmarshals the response into AsrProcessResponse
func (r *AsrProcessRequest) Exec(ctx context.Context) (response response.AsrProcessResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Audio required method define parameter audio.
//
//	JSON response from a request to send an audio file to the download server address.
func (r *AsrProcessRequest) Audio(audio string) *AsrProcessRequest {
	r.parameters.Set(constants.ParameterNameAudio, audio)
	return r
}

// Model required method define parameter model.
// The speech recognition model to use. Possible values:
//
//	"neutral" - recognizes intelligible speech, like in an interview or TV show.
//	"spontaneous" - speech recognition with slang and profanity.
func (r *AsrProcessRequest) Model(model constants.AsrProcessModel) *AsrProcessRequest {
	r.parameters.Set(constants.ParameterNameModel, string(model))
	return r
}
