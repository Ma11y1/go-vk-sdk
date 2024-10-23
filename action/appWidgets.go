package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// AppWidgets Doc: https://dev.vk.com/ru/method/appWidgets
type AppWidgets struct {
	BaseAction
}

// GetAppImageUploadServer Doc: https://dev.vk.com/method/appWidgets.getAppImageUploadServer
func (a *AppWidgets) GetAppImageUploadServer(actor actor.Actor) *request.AppWidgetsGetAppImageUploadServer {
	return request.NewAppWidgetsGetAppImageUploadServer(a.api, actor)
}

// GetAppImages Doc: https://dev.vk.com/method/appWidgets.getAppImages
func (a *AppWidgets) GetAppImages(actor actor.Actor) {
	return
}

// GetGroupImageUploadServer Doc: https://dev.vk.com/method/appWidgets.getGroupImageUploadServer
func (a *AppWidgets) GetGroupImageUploadServer(actor actor.Actor) {
	return
}

// GetGroupImages Doc: https://dev.vk.com/method/appWidgets.getGroupImages
func (a *AppWidgets) GetGroupImages(actor actor.Actor) {
	return
}

// GetImagesByID Doc: https://dev.vk.com/method/appWidgets.getImagesById
func (a *AppWidgets) GetImagesByID(actor actor.Actor) {
	return
}

// SaveAppImage Doc: https://dev.vk.com/method/appWidgets.saveAppImage
func (a *AppWidgets) SaveAppImage(actor actor.Actor) {
	return
}

// SaveGroupImage Doc: https://dev.vk.com/method/appWidgets.saveGroupImage
func (a *AppWidgets) SaveGroupImage(actor actor.Actor) {
	return
}

// Update Doc: https://dev.vk.com/method/appWidgets.update
func (a *AppWidgets) Update(actor actor.Actor) {
	return
}
