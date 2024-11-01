package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Marusia Doc: https://dev.vk.com/ru/method/marusia
type Marusia struct {
	BaseAction
}

// GetPictureUploadLink Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) GetPictureUploadLink(actor actor.Actor) *request.MarusiaGetPictureUploadLinkRequest {
	return request.NewMarusiaGetPictureUploadLinkRequest(a.api, actor)
}

// SavePicture Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) SavePicture(actor actor.Actor) *request.MarusiaSavePictureRequest {
	return request.NewMarusiaSavePictureRequest(a.api, actor)
}

// GetPictures Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) GetPictures(actor actor.Actor) *request.MarusiaGetPicturesRequest {
	return request.NewMarusiaGetPicturesRequest(a.api, actor)
}

// DeletePicture Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) DeletePicture(actor actor.Actor) *request.MarusiaDeletePictureRequest {
	return request.NewMarusiaDeletePictureRequest(a.api, actor)
}

// GetAudioUploadLink Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) GetAudioUploadLink(actor actor.Actor) *request.MarusiaGetAudioUploadLinkRequest {
	return request.NewMarusiaGetAudioUploadLinkRequest(a.api, actor)
}

// CreateAudio Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) CreateAudio(actor actor.Actor) *request.MarusiaCreateAudioRequest {
	return request.NewMarusiaCreateAudioRequest(a.api, actor)
}

// GetAudios Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) GetAudios(actor actor.Actor) *request.MarusiaGetAudiosRequest {
	return request.NewMarusiaGetAudiosRequest(a.api, actor)
}

// DeleteAudio Doc: https://dev.vk.com/ru/method/marusia
func (a *Marusia) DeleteAudio(actor actor.Actor) *request.MarusiaDeleteAudioRequest {
	return request.NewMarusiaDeleteAudioRequest(a.api, actor)
}
