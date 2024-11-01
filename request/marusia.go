package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/marusia/media-api

// MarusiaGetPictureUploadLinkRequest defines the request for marusia.getPictureUploadLink
//
//	To get the URL: this is where you will post the images.
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaGetPictureUploadLinkRequest struct {
	BaseRequest
}

// NewMarusiaGetPictureUploadLinkRequest creates a new request for marusia.getPictureUploadLink
func NewMarusiaGetPictureUploadLinkRequest(a *api.API, actor actor.Actor) *MarusiaGetPictureUploadLinkRequest {
	return &MarusiaGetPictureUploadLinkRequest{*NewMethodBaseRequest(a, actor, "marusia.getPictureUploadLink")}
}

// Exec executes the request and unmarshals the response into MarusiaGetPictureUploadLinkResponse
func (r *MarusiaGetPictureUploadLinkRequest) Exec(ctx context.Context) (response response.MarusiaGetPictureUploadLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaSavePictureRequest defines the request for marusia.savePicture
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaSavePictureRequest struct {
	BaseRequest
}

// NewMarusiaSavePictureRequest creates a new request for marusia.savePicture
func NewMarusiaSavePictureRequest(a *api.API, actor actor.Actor) *MarusiaSavePictureRequest {
	return &MarusiaSavePictureRequest{*NewMethodBaseRequest(a, actor, "marusia.savePicture")}
}

// Exec executes the request and unmarshals the response into MarusiaSavePictureResponse
func (r *MarusiaSavePictureRequest) Exec(ctx context.Context) (response response.MarusiaSavePictureResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaGetPicturesRequest defines the request for marusia.getPictures
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaGetPicturesRequest struct {
	BaseRequest
}

// NewMarusiaGetPicturesRequest creates a new request for marusia.getPictures
func NewMarusiaGetPicturesRequest(a *api.API, actor actor.Actor) *MarusiaGetPicturesRequest {
	return &MarusiaGetPicturesRequest{*NewMethodBaseRequest(a, actor, "marusia.getPictures")}
}

// Exec executes the request and unmarshals the response into MarusiaGetPicturesResponse
func (r *MarusiaGetPicturesRequest) Exec(ctx context.Context) (response response.MarusiaGetPicturesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaDeletePictureRequest defines the request for marusia.deletePicture
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaDeletePictureRequest struct {
	BaseRequest
}

// NewMarusiaDeletePictureRequest creates a new request for marusia.deletePicture
func NewMarusiaDeletePictureRequest(a *api.API, actor actor.Actor) *MarusiaDeletePictureRequest {
	return &MarusiaDeletePictureRequest{*NewMethodBaseRequest(a, actor, "marusia.deletePicture")}
}

// Exec executes the request and unmarshals the response into MarusiaDeletePictureResponse
func (r *MarusiaDeletePictureRequest) Exec(ctx context.Context) (response response.MarusiaDeletePictureResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaGetAudioUploadLinkRequest defines the request for marusia.getAudioUploadLink
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaGetAudioUploadLinkRequest struct {
	BaseRequest
}

// NewMarusiaGetAudioUploadLinkRequest creates a new request for marusia.getAudioUploadLink
func NewMarusiaGetAudioUploadLinkRequest(a *api.API, actor actor.Actor) *MarusiaGetAudioUploadLinkRequest {
	return &MarusiaGetAudioUploadLinkRequest{*NewMethodBaseRequest(a, actor, "marusia.getAudioUploadLink")}
}

// Exec executes the request and unmarshals the response into MarusiaGetAudioUploadLinkResponse
func (r *MarusiaGetAudioUploadLinkRequest) Exec(ctx context.Context) (response response.MarusiaGetAudioUploadLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaCreateAudioRequest defines the request for marusia.createAudio
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaCreateAudioRequest struct {
	BaseRequest
}

// NewMarusiaCreateAudioRequest creates a new request for marusia.createAudio
func NewMarusiaCreateAudioRequest(a *api.API, actor actor.Actor) *MarusiaCreateAudioRequest {
	return &MarusiaCreateAudioRequest{*NewMethodBaseRequest(a, actor, "marusia.createAudio")}
}

// Exec executes the request and unmarshals the response into MarusiaCreateAudioResponse
func (r *MarusiaCreateAudioRequest) Exec(ctx context.Context) (response response.MarusiaCreateAudioResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaGetAudiosRequest defines the request for marusia.getAudios
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaGetAudiosRequest struct {
	BaseRequest
}

// NewMarusiaGetAudiosRequest creates a new request for marusia.getAudios
func NewMarusiaGetAudiosRequest(a *api.API, actor actor.Actor) *MarusiaGetAudiosRequest {
	return &MarusiaGetAudiosRequest{*NewMethodBaseRequest(a, actor, "marusia.getAudios")}
}

// Exec executes the request and unmarshals the response into MarusiaGetAudiosResponse
func (r *MarusiaGetAudiosRequest) Exec(ctx context.Context) (response response.MarusiaGetAudiosResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarusiaDeleteAudioRequest defines the request for marusia.deleteAudio
//
//	Doc: https://dev.vk.com/ru/marusia/media-api
type MarusiaDeleteAudioRequest struct {
	BaseRequest
}

// NewMarusiaDeleteAudioRequest creates a new request for marusia.deleteAudio
func NewMarusiaDeleteAudioRequest(a *api.API, actor actor.Actor) *MarusiaDeleteAudioRequest {
	return &MarusiaDeleteAudioRequest{*NewMethodBaseRequest(a, actor, "marusia.deleteAudio")}
}

// Exec executes the request and unmarshals the response into MarusiaDeleteAudioResponse
func (r *MarusiaDeleteAudioRequest) Exec(ctx context.Context) (response response.MarusiaDeleteAudioResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
