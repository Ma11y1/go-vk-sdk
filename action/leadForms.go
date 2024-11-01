package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// LeadForms Doc: https://dev.vk.com/ru/method/leadForms
type LeadForms struct {
	BaseAction
}

// Create Doc: https://dev.vk.com/ru/method/leadForms.create
func (l *LeadForms) Create(actor actor.Actor) *request.LeadFormsCreateRequest {
	return request.NewLeadFormsCreateRequest(l.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/leadForms.delete
func (l *LeadForms) Delete(actor actor.Actor) *request.LeadFormsDeleteRequest {
	return request.NewLeadFormsDeleteRequest(l.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/leadForms.get
func (l *LeadForms) Get(actor actor.Actor) *request.LeadFormsGetRequest {
	return request.NewLeadFormsGetRequest(l.api, actor)
}

// GetLeads Doc: https://dev.vk.com/ru/method/leadForms.getLeads
func (l *LeadForms) GetLeads(actor actor.Actor) *request.LeadFormsGetLeadsRequest {
	return request.NewLeadFormsGetLeadsRequest(l.api, actor)
}

// GetUploadURL Doc: https://dev.vk.com/ru/method/leadForms.getUploadURL
func (l *LeadForms) GetUploadURL(actor actor.Actor) *request.LeadFormsGetUploadURLRequest {
	return request.NewLeadFormsGetUploadURLRequest(l.api, actor)
}

// List Doc: https://dev.vk.com/ru/method/leadForms.list
func (l *LeadForms) List(actor actor.Actor) *request.LeadFormsListRequest {
	return request.NewLeadFormsListRequest(l.api, actor)
}

// Update Doc: https://dev.vk.com/ru/method/leadForms.update
func (l *LeadForms) Update(actor actor.Actor) *request.LeadFormsUpdateRequest {
	return request.NewLeadFormsUpdateRequest(l.api, actor)
}
