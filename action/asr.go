package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Asr Doc: https://dev.vk.com/ru/method/asr
type Asr struct {
	BaseAction
}

// CheckStatus Doc: https://dev.vk.com/ru/method/asr.checkStatus
func (a *Asr) CheckStatus(actor actor.Actor) *request.AsrCheckStatusRequest {
	return request.NewAsrCheckStatusRequest(a.api, actor)
}

// GetUploadURL Doc: https://dev.vk.com/ru/method/asr.getUploadUrl
func (a *Asr) GetUploadURL(actor actor.Actor) *request.AsrGetUploadURLRequest {
	return request.NewAsrGetUploadURLRequest(a.api, actor)
}

// Process Doc: https://dev.vk.com/ru/method/asr.process
func (a *Asr) Process(actor actor.Actor) *request.AsrProcessRequest {
	return request.NewAsrProcessRequest(a.api, actor)
}
