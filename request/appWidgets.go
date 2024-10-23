package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"strconv"
)

// Doc: https://dev.vk.com/method/appWidgets

// AppWidgetsGetAppImageUploadServer defines the request for appWidgets.getAppImageUploadServer
// Allows you to get the address for uploading a photo to the application collection for community application widgets - https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.getAppImageUploadServer
type AppWidgetsGetAppImageUploadServer struct {
	BaseRequest
}

// NewAppWidgetsGetAppImageUploadServer creates a new request for appWidgets.getAppImageUploadServer
func NewAppWidgetsGetAppImageUploadServer(a *api.API, actor actor.Actor) *AppWidgetsGetAppImageUploadServer {
	return &AppWidgetsGetAppImageUploadServer{*NewMethodBaseRequest(a, actor, "appWidgets.getAppImageUploadServer")}
}

// Exec executes the request and unmarshals the response into AppWidgetsGetAppImageUploadServerResponse
func (r *AppWidgetsGetAppImageUploadServer) Exec(ctx context.Context) (response response.AppWidgetsGetAppImageUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// ImageType - method to set the required parameter image_type.
// Image type. Possible values:
// 24x24
// 50x50
// 160x160
// 160x240
// 510x128
func (r *AppWidgetsGetAppImageUploadServer) ImageType(t constants.AppWidgetsImageType) *AppWidgetsGetAppImageUploadServer {
	r.parameters.Set(constants.ParameterNameImageType, string(t))
	return r
}

// AppWidgetsGetAppImages defines the request for appWidgets.getAppImages
// Allows you to get a collection of images uploaded for the application in https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.getAppImages
type AppWidgetsGetAppImages struct {
	BaseRequest
}

// NewAppWidgetsGetAppImages creates a new request for appWidgets.getAppImages
func NewAppWidgetsGetAppImages(a *api.API, actor actor.Actor) *AppWidgetsGetAppImages {
	return &AppWidgetsGetAppImages{*NewMethodBaseRequest(a, actor, "appWidgets.getAppImages")}
}

// Exec executes the request and unmarshals the response into AppWidgetsGetAppImagesResponse
func (r *AppWidgetsGetAppImages) Exec(ctx context.Context) (response response.AppWidgetsGetAppImagesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Offset - method to set the required parameter offset.
// offset to produce a specific subset of results.
func (r *AppWidgetsGetAppImages) Offset(offset int) *AppWidgetsGetAppImages {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// Count - method to set the required parameter count.
// the maximum number of results to obtain.
func (r *AppWidgetsGetAppImages) Count(count int) *AppWidgetsGetAppImages {
	r.parameters.Set(constants.ParameterNameCount, strconv.Itoa(count))
	return r
}

// ImageType - method to set the required parameter image_type.
// Image type. Possible values:
// 24x24
// 50x50
// 160x160
// 160x240
// 510x128
func (r *AppWidgetsGetAppImages) ImageType(t constants.AppWidgetsImageType) *AppWidgetsGetAppImages {
	r.parameters.Set(constants.ParameterNameImageType, string(t))
	return r
}

// AppWidgetsGetGroupImageUploadServer defines the request for appWidgets.getGroupImageUploadServer
// Allows you to get the address for uploading a photo to the community collection for community application widgets - https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.getGroupImageUploadServer
type AppWidgetsGetGroupImageUploadServer struct {
	BaseRequest
}

// NewAppWidgetsGetGroupImageUploadServer creates a new request for appWidgets.getGroupImageUploadServer
func NewAppWidgetsGetGroupImageUploadServer(a *api.API, actor actor.Actor) *AppWidgetsGetGroupImageUploadServer {
	return &AppWidgetsGetGroupImageUploadServer{*NewMethodBaseRequest(a, actor, "appWidgets.getGroupImageUploadServer")}
}

// Exec executes the request and unmarshals the response into AppWidgetsGetGroupImageUploadServerResponse
func (r *AppWidgetsGetGroupImageUploadServer) Exec(ctx context.Context) (response response.AppWidgetsGetGroupImageUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// ImageType - method to set the required parameter image_type.
// Image type. Possible values:
// 24x24
// 50x50
// 160x160
// 160x240
// 510x128
func (r *AppWidgetsGetGroupImageUploadServer) ImageType(t constants.AppWidgetsImageType) *AppWidgetsGetGroupImageUploadServer {
	r.parameters.Set(constants.ParameterNameImageType, string(t))
	return r
}

// AppWidgetsGetGroupImages defines the request for appWidgets.getGroupImages
// Allows you to get a collection of images uploaded for the application in community application widgets - https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.getGroupImages
type AppWidgetsGetGroupImages struct {
	BaseRequest
}

// NewAppWidgetsGetGroupImages creates a new request for appWidgets.getGroupImages
func NewAppWidgetsGetGroupImages(a *api.API, actor actor.Actor) *AppWidgetsGetGroupImages {
	return &AppWidgetsGetGroupImages{*NewMethodBaseRequest(a, actor, "appWidgets.getGroupImages")}
}

// Exec executes the request and unmarshals the response into AppWidgetsGetGroupImagesResponse
func (r *AppWidgetsGetGroupImages) Exec(ctx context.Context) (response response.AppWidgetsGetGroupImagesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Offset - method to set the required parameter offset.
// offset to produce a specific subset of results.
func (r *AppWidgetsGetGroupImages) Offset(offset int) *AppWidgetsGetGroupImages {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// Count - method to set the required parameter count.
// the maximum number of results to obtain.
func (r *AppWidgetsGetGroupImages) Count(count int) *AppWidgetsGetGroupImages {
	r.parameters.Set(constants.ParameterNameCount, strconv.Itoa(count))
	return r
}

// ImageType - method to set the required parameter image_type.
// Image type. Possible values:
// 24x24
// 50x50
// 160x160
// 160x240
// 510x128
func (r *AppWidgetsGetGroupImages) ImageType(t constants.AppWidgetsImageType) *AppWidgetsGetGroupImages {
	r.parameters.Set(constants.ParameterNameImageType, string(t))
	return r
}

// AppWidgetsGetImagesByID defines the request for appWidgets.getImagesById
// Allows you to get an image for community application widgets by its ID.
// Doc: https://dev.vk.com/method/appWidgets.getImagesById
type AppWidgetsGetImagesByID struct {
	BaseRequest
}

// NewAppWidgetsAppWidgetsGetImagesByID creates a new request for appWidgets.getImagesById
func NewAppWidgetsAppWidgetsGetImagesByID(a *api.API, actor actor.Actor) *AppWidgetsGetImagesByID {
	return &AppWidgetsGetImagesByID{*NewMethodBaseRequest(a, actor, "appWidgets.getImagesById")}
}

// Exec executes the request and unmarshals the response into AppWidgetsGetImagesByID
func (r *AppWidgetsGetImagesByID) Exec(ctx context.Context) (response response.AppWidgetsGetImagesByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Images - method to set the required parameter images. List of image IDs to retrieve.
func (r *AppWidgetsGetImagesByID) Images(images string) *AppWidgetsGetImagesByID {
	r.parameters.Set(constants.ParameterNameImages, images)
	return r
}

// AppWidgetsSaveAppImage defines the request for appWidgets.saveAppImage
// Allows you to save the image to the application collection for community application widgets after uploading to the server.
// https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.saveAppImage
type AppWidgetsSaveAppImage struct {
	BaseRequest
}

// NewAppWidgetsAppWidgetsSaveAppImage creates a new request for appWidgets.saveAppImage
func NewAppWidgetsAppWidgetsSaveAppImage(a *api.API, actor actor.Actor) *AppWidgetsSaveAppImage {
	return &AppWidgetsSaveAppImage{*NewMethodBaseRequest(a, actor, "appWidgets.saveAppImage")}
}

// Exec executes the request and unmarshals the response into AppWidgetsSaveAppImage
func (r *AppWidgetsSaveAppImage) Exec(ctx context.Context) (response response.AppWidgetsSaveAppImageResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Hash - method to set the required parameter hash. The hash parameter received after uploading to the server.
func (r *AppWidgetsSaveAppImage) Hash(images string) *AppWidgetsSaveAppImage {
	r.parameters.Set(constants.ParameterNameHash, images)
	return r
}

// Image - method to set the required parameter image. The image parameter received after uploading to the server.
func (r *AppWidgetsSaveAppImage) Image(images string) *AppWidgetsSaveAppImage {
	r.parameters.Set(constants.ParameterNameImage, images)
	return r
}

// AppWidgetsSaveGroupImage defines the request for appWidgets.saveGroupImage
// Allows you to save an image to the community collection for community app widgets. after uploading to the server.
// https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.saveGroupImage
type AppWidgetsSaveGroupImage struct {
	BaseRequest
}

// NewAppWidgetsSaveGroupImage creates a new request for appWidgets.saveGroupImage
func NewAppWidgetsSaveGroupImage(a *api.API, actor actor.Actor) *AppWidgetsSaveGroupImage {
	return &AppWidgetsSaveGroupImage{*NewMethodBaseRequest(a, actor, "appWidgets.saveGroupImage")}
}

// Exec executes the request and unmarshals the response into AppWidgetsSaveAppImage
func (r *AppWidgetsSaveGroupImage) Exec(ctx context.Context) (response response.AppWidgetsSaveGroupImageResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Hash - method to set the required parameter hash. The hash parameter received after uploading to the server.
func (r *AppWidgetsSaveGroupImage) Hash(images string) *AppWidgetsSaveGroupImage {
	r.parameters.Set(constants.ParameterNameHash, images)
	return r
}

// Image - method to set the required parameter image. The image parameter received after uploading to the server.
func (r *AppWidgetsSaveGroupImage) Image(images string) *AppWidgetsSaveGroupImage {
	r.parameters.Set(constants.ParameterNameImage, images)
	return r
}

// AppWidgetsUpdate defines the request for appWidgets.update
// Allows you to update the community application widget.
// https://dev.vk.com/ru/api/community-apps-widgets/getting-started
// Doc: https://dev.vk.com/method/appWidgets.update
type AppWidgetsUpdate struct {
	BaseRequest
}

// NewAppWidgetsUpdate creates a new request for appWidgets.update
func NewAppWidgetsUpdate(a *api.API, actor actor.Actor) *AppWidgetsUpdate {
	return &AppWidgetsUpdate{*NewMethodBaseRequest(a, actor, "appWidgets.update")}
}

// Exec executes the request and unmarshals the response into AppWidgetsUpdateResponse
func (r *AppWidgetsUpdate) Exec(ctx context.Context) (response response.AppWidgetsUpdateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Code - method to set the required parameter code. widget code. For a detailed description, see https://dev.vk.com/ru/api/community-apps-widgets/getting-started
func (r *AppWidgetsUpdate) Code(code string) *AppWidgetsUpdate {
	r.parameters.Set(constants.ParameterNameHash, code)
	return r
}

// Type - method to set the required parameter type. widget type. For a list of all available types, see https://dev.vk.com/ru/reference/objects/app-widget
func (r *AppWidgetsUpdate) Type(t string) *AppWidgetsUpdate {
	r.parameters.Set(constants.ParameterNameImage, t)
	return r
}
