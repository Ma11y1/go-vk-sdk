package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Docs Doc: https://dev.vk.com/ru/method/docs
type Docs struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/ru/method/docs.add
func (d *Docs) Add(actor actor.Actor) *request.DocsAddRequest {
	return request.NewDocsAddRequest(d.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/docs.delete
func (d *Docs) Delete(actor actor.Actor) *request.DocsDeleteRequest {
	return request.NewDocsDeleteRequest(d.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/docs.edit
func (d *Docs) Edit(actor actor.Actor) *request.DocsEditRequest {
	return request.NewDocsEditRequest(d.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/docs.get
func (d *Docs) Get(actor actor.Actor) *request.DocsGetRequest {
	return request.NewDocsGetRequest(d.api, actor)
}

// GetByID Doc: https://dev.vk.com/ru/method/docs.getById
func (d *Docs) GetByID(actor actor.Actor) *request.DocsGetByIDRequest {
	return request.NewDocsGetByIDRequest(d.api, actor)
}

// GetMessagesUploadServer Doc: https://dev.vk.com/ru/method/docs.getMessagesUploadServer
func (d *Docs) GetMessagesUploadServer(actor actor.Actor) *request.DocsGetMessagesUploadServerRequest {
	return request.NewDocsGetMessagesUploadServerRequest(d.api, actor)
}

// GetTypes Doc: https://dev.vk.com/ru/method/docs.getTypes
func (d *Docs) GetTypes(actor actor.Actor) *request.DocsGetTypesRequest {
	return request.NewDocsGetTypesRequest(d.api, actor)
}

// GetUploadServer Doc: https://dev.vk.com/ru/method/docs.getUploadServer
func (d *Docs) GetUploadServer(actor actor.Actor) *request.DocsGetUploadServerRequest {
	return request.NewDocsGetUploadServerRequest(d.api, actor)
}

// GetWallUploadServer Doc: https://dev.vk.com/ru/method/docs.getWallUploadServer
func (d *Docs) GetWallUploadServer(actor actor.Actor) *request.DocsGetWallUploadServerRequest {
	return request.NewDocsGetWallUploadServerRequest(d.api, actor)
}

// Save Doc: https://dev.vk.com/ru/method/docs.save
func (d *Docs) Save(actor actor.Actor) *request.DocsSaveRequest {
	return request.NewDocsSaveRequest(d.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/docs.search
func (d *Docs) Search(actor actor.Actor) *request.DocsSearchRequest {
	return request.NewDocsSearchRequest(d.api, actor)
}
