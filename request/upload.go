package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/objects"
	"go-vk-sdk/response"
	"io"
	"mime/multipart"
)

// Doc: https://dev.vk.com/ru/api/upload/overview

func uploadFile(request *BaseRequest, ctx context.Context, file *objects.UploadFile, response interface{}) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	part, err := writer.CreateFormFile(file.FieldName, file.Name)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file.Data)
	if err != nil {
		return err
	}

	request.SetContentType(writer.FormDataContentType())

	resp, err := request.PostData(ctx, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}

func uploadFiles(request *BaseRequest, ctx context.Context, files *[]objects.UploadFile, response interface{}) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	for i, file := range *files {
		part, err := writer.CreateFormFile(fmt.Sprintf("%s%d", file.FieldName, i+1), file.Name)
		if err != nil {
			return err
		}

		_, err = io.Copy(part, file.Data)
		if err != nil {
			return err
		}
	}

	request.SetContentType(writer.FormDataContentType())

	resp, err := request.PostData(ctx, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}

func uploadFileSquareCrop(request *BaseRequest, ctx context.Context, file *objects.UploadFile, crop string, response interface{}) error {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	part, err := writer.CreateFormFile(file.FieldName, file.Name)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file.Data)
	if err != nil {
		return err
	}

	request.SetContentType(writer.FormDataContentType())

	if crop != "" {
		err = writer.WriteField("_square_crop", crop)
		if err != nil {
			return err
		}
	}

	resp, err := request.PostData(ctx, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}

// UploadPhotoAlbumRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//
//	Restrictions:
//	• No more than 5 photos per request.
//	• The sum of height and width is no more than 14,000 pixels.
//	• The file is no more than 50 MB in size.
//	• Aspect ratio of at least 1:20.
//	Note. Photos can only be uploaded to a user-created album
//	To get the photo upload address, call the photos.getUploadServer method.
//	Doc: https://dev.vk.com/ru/api/upload/album-photos
type UploadPhotoAlbumRequest struct {
	BaseRequest
}

func NewUploadPhotoAlbumRequest(a *api.API, actor actor.Actor) *UploadPhotoAlbumRequest {
	return &UploadPhotoAlbumRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadPhotoAlbumRequest) Exec(ctx context.Context, files *[]objects.UploadFile) (response response.PhotosUploadPhotoAlbumResponse, err error) {
	err = uploadFiles(&r.BaseRequest, ctx, files, &response)
	return
}

// UploadPhotoWallRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//
//	Restrictions:
//	- The sum of height and width is no more than 14,000 pixels.
//	- The file is no more than 50 MB in size.
//	- Aspect ratio of at least 1:20
//	To get the address for uploading photos, call the photos.getWallUploadServer method
//	Doc: https://dev.vk.com/ru/api/upload/wall-photo
type UploadPhotoWallRequest struct {
	BaseRequest
}

func NewUploadPhotoWallRequest(a *api.API, actor actor.Actor) *UploadPhotoWallRequest {
	return &UploadPhotoWallRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPhotoWallRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoWallResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoOwnerRequest
//
// Acceptable formats: JPG, PNG, GIF.
//
//	Restrictions:
//	• Size no less than 200×200 px.
//	• The minimum photo size is 795×265 px.
//	• Recommended size is 1590×530 px.
//	• The sum of height and width is no more than 14,000 pixels.
//	• The file is no more than 50 MB in size.
//	• Aspect ratio: 0.25 to 3.
//	• You can download no more than 1,500 covers per day.
//	To get the photo upload address, call the photos.getOwnerPhotoUploadServer method.
//	Doc: https://dev.vk.com/ru/api/upload/main-photo-in-profile
type UploadPhotoOwnerRequest struct {
	BaseRequest
}

func NewUploadPhotoOwnerRequest(a *api.API, actor actor.Actor) *UploadPhotoOwnerRequest {
	return &UploadPhotoOwnerRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPhotoOwnerRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoOwnerResponse, err error) {
	err = uploadFileSquareCrop(&r.BaseRequest, ctx, file, r.parameters.Get(constants.ParameterNameFieldSquareCrop), &response)
	return
}

func (r *UploadPhotoOwnerRequest) SquareCrop(crop string) *UploadPhotoOwnerRequest {
	r.parameters.Set(constants.ParameterNameFieldSquareCrop, crop)
	return r
}

// UploadPhotoMessagesRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//
//	Limits:
//	- width+height not more than 14000 px
//	- file size up to 50 Mb
//	- aspect ratio of at least 1:20.
//	Doc: https://dev.vk.com/ru/api/upload/photo-in-message
type UploadPhotoMessagesRequest struct {
	BaseRequest
}

func NewUploadPhotoMessagesRequest(a *api.API, actor actor.Actor) *UploadPhotoMessagesRequest {
	return &UploadPhotoMessagesRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPhotoMessagesRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoMessageResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoChatRequest
//
// Acceptable formats: JPG, PNG, GIF.
//
//	Limits:
//	- size not less than 200x200px
//	- width+height not more than 14000 px
//	- file size up to 50 Mb, aspect ratio of at least 1:20
//	Doc: https://dev.vk.com/ru/api/upload/main-photo-in-chat
type UploadPhotoChatRequest struct {
	BaseRequest
}

func NewUploadPhotoChatRequest(a *api.API, actor actor.Actor) *UploadPhotoChatRequest {
	return &UploadPhotoChatRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPhotoChatRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoChatResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoMarketRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	The minimum photo size is 400x400 pixels.
//	The sum of height and angle is no more than 14,000 pixels.
//	The file size is no more than 50 MB.
//	Aspect ratio - no more than 1:20.
//	To optimize the load on servers and subsystems, VK uses different addresses for downloading files.
//	To get the upload address for your image, send an api request to market.getProductPhotoUploadServer.
//	Doc: https://dev.vk.com/ru/api/upload/photo-in-market
type UploadPhotoMarketRequest struct {
	BaseRequest
}

func NewUploadPhotoMarketRequest(a *api.API, actor actor.Actor) *UploadPhotoMarketRequest {
	return &UploadPhotoMarketRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadPhotoMarketRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoMarketResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoMarketAlbumRequest Acceptable formats: JPG, PNG, GIF.
//
// Acceptable formats: JPG, PNG, GIF.
//
//	Limits:
//	Size no less than 1280×720 px
//	The sum of height and width is no more than 14,000 pixels
//	The file is no more than 50 MB in size
//	Aspect ratio of at least 1:20
//	To optimize the load on servers and subsystems, VK uses different addresses for downloading files.
//	To get the upload address for your image, send an api request to market.getProductPhotoUploadServer.
//	Doc: https://dev.vk.com/ru/api/upload/main-photo-in-market
type UploadPhotoMarketAlbumRequest struct {
	BaseRequest
}

func NewUploadPhotoMarketAlbumRequest(a *api.API, actor actor.Actor) *UploadPhotoMarketAlbumRequest {
	return &UploadPhotoMarketAlbumRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadPhotoMarketAlbumRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoMarketAlbumResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadVideoRequest
//
// Acceptable formats: AVI, MP4, 3GP, MPEG, MOV, MP3, FLV, WMV.
//
//	Doc: https://dev.vk.com/ru/api/upload/video-in-profile
type UploadVideoRequest struct {
	BaseRequest
}

func NewUploadVideoRequest(a *api.API, actor actor.Actor) *UploadVideoRequest {
	return &UploadVideoRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "video"
func (r *UploadVideoRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadVideoResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadDocumentRequest
//
//	Acceptable formats: any except MP3 and executable files.
//	Limitations: file size no more than 200 MB.
//	To get the document download address, call one of the following methods:
//	- docs.getUploadServer - to upload a document to the Files section.
//	- docs.getWallUploadServer - for uploading a document to the wall.
//	- docs.getMessagesUploadServer - for uploading a document to a private message.
//	If you want to upload a document to a community, pass the community ID in the group_id parameter.
//	Doc: https://dev.vk.com/ru/api/upload/document-in-profile
type UploadDocumentRequest struct {
	BaseRequest
}

func NewUploadDocumentRequest(a *api.API, actor actor.Actor) *UploadDocumentRequest {
	return &UploadDocumentRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadDocumentRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadDocumentResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoOwnerCoverRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	Limitations:
//	- The minimum photo size is 795×265 px.
//	- Recommended size is 1590×530 px.
//	- The sum of height and width is no more than 14,000 pixels.
//	- The file is no more than 50 MB in size.
//	- Aspect ratio of at least 1:20.
//	- You can download no more than 1,500 covers per day.
//	To get the photo upload address, call the photos.getOwnerCoverPhotoUploadServer method
//	Doc: https://dev.vk.com/ru/api/upload/main-photo-in-group
type UploadPhotoOwnerCoverRequest struct {
	BaseRequest
}

func NewUploadPhotoOwnerCoverRequest(a *api.API, actor actor.Actor) *UploadPhotoOwnerCoverRequest {
	return &UploadPhotoOwnerCoverRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPhotoOwnerCoverRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoOwnerResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadAudioMessageRequest
//
//	Acceptable formats: OGG, OPUS.
//	Limitations:
//	- Sampling frequency - 16 kHz.
//	- Bitrate - 16 Kbps.
//	- Duration: no more than 60 minutes.
//	- Format: mono sound
//	To get the address for downloading an audio message, call the docs.getMessagesUploadServer method with the type=audio_message parameter.
//	Doc: https://dev.vk.com/ru/api/upload/audio-record
type UploadAudioMessageRequest struct {
	BaseRequest
}

func NewUploadAudioMessageRequest(a *api.API, actor actor.Actor) *UploadAudioMessageRequest {
	return &UploadAudioMessageRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadAudioMessageRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.AudioUploadAudioMessageResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPhotoStoriesRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	Limitations: sum of with and height no more than 14000px, file size no more than 10 MB
//	To get the image upload address, call the stories.getPhotoUploadServer method.
//	Doc: https://dev.vk.com/ru/api/upload/story-in-profile
type UploadPhotoStoriesRequest struct {
	BaseRequest
}

func NewUploadPhotoStoriesRequest(a *api.API, actor actor.Actor) *UploadPhotoStoriesRequest {
	return &UploadPhotoStoriesRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadPhotoStoriesRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PhotosUploadPhotoStoriesResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadVideoStoriesRequest
//
//	Acceptable formats: H.264, AAC, MP4
//	Limitations: h264 video, aac audio, maximum 720х1280, 30fps.
//	To get the video upload address, call the stories.getVideoUploadServer method.
//	Doc: https://dev.vk.com/ru/api/upload/story-in-profile
type UploadVideoStoriesRequest struct {
	BaseRequest
}

func NewUploadVideoStoriesRequest(a *api.API, actor actor.Actor) *UploadVideoStoriesRequest {
	return &UploadVideoStoriesRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "video"
func (r *UploadVideoStoriesRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.VideosUploadVideoStoriesResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPollsPhotoRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	Limitations: minimum photo size 795x200px, width+height not more than 14000px, file size up to 50 MB. Recommended size: 1590x400px.
//	To get the photo upload address, call polls.getPhotoUploadServer method.
type UploadPollsPhotoRequest struct {
	BaseRequest
}

func NewUploadPollsPhotoRequest(a *api.API, actor actor.Actor) *UploadPollsPhotoRequest {
	return &UploadPollsPhotoRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadPollsPhotoRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PollsUploadPhotoResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadPrettyCardsPhotoRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	To get the photo upload address, call prettyCards.getUploadURL method.
type UploadPrettyCardsPhotoRequest struct {
	BaseRequest
}

func NewUploadPrettyCardsPhotoRequest(a *api.API, actor actor.Actor) *UploadPrettyCardsPhotoRequest {
	return &UploadPrettyCardsPhotoRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadPrettyCardsPhotoRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.PrettyCardsUploadPhotoResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadLeadFormsPhotoRequest
//
//	Acceptable formats: JPG, PNG, GIF.
//	To get the photo upload address, call leadForms.getUploadURL method.
type UploadLeadFormsPhotoRequest struct {
	BaseRequest
}

func NewUploadLeadFormsPhotoRequest(a *api.API, actor actor.Actor) *UploadLeadFormsPhotoRequest {
	return &UploadLeadFormsPhotoRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadLeadFormsPhotoRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.LeadFormsUploadPhotoResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadAppImageRequest uploading Image into App collection for community app widgets.
type UploadAppImageRequest struct {
	BaseRequest
}

func NewUploadAppImageRequest(a *api.API, actor actor.Actor) *UploadAppImageRequest {
	return &UploadAppImageRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "image"
func (r *UploadAppImageRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.AppUploadWidgetsImageResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadAppGroupImageRequest uploading Image into Community collection for community app widgets.
type UploadAppGroupImageRequest struct {
	BaseRequest
}

func NewUploadAppGroupImageRequest(a *api.API, actor actor.Actor) *UploadAppGroupImageRequest {
	return &UploadAppGroupImageRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "image"
func (r *UploadAppGroupImageRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.AppUploadWidgetsGroupImageResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadMarussiaPictureRequest
//
//	Limits: height not more than 600 px, aspect ratio of at least 2:1.
type UploadMarussiaPictureRequest struct {
	BaseRequest
}

func NewUploadMarussiaPictureRequest(a *api.API, actor actor.Actor) *UploadMarussiaPictureRequest {
	return &UploadMarussiaPictureRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "photo"
func (r *UploadMarussiaPictureRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.MarusiaUploadPictureResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}

// UploadMarussiaAudioRequest
//
//	https://dev.vk.com/ru/marusia/media-api
type UploadMarussiaAudioRequest struct {
	BaseRequest
}

func NewUploadMarussiaAudioRequest(a *api.API, actor actor.Actor) *UploadMarussiaAudioRequest {
	return &UploadMarussiaAudioRequest{*NewUploadBaseRequest(a, actor)}
}

// Exec UploadFile.FieldName = "file"
func (r *UploadMarussiaAudioRequest) Exec(ctx context.Context, file *objects.UploadFile) (response response.MarusiaUploadAudioResponse, err error) {
	err = uploadFile(&r.BaseRequest, ctx, file, &response)
	return
}
