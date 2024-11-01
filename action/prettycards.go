package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// PrettyCards Doc: https://dev.vk.com/ru/method/prettyCards
type PrettyCards struct {
	BaseAction
}

// Create Doc: https://dev.vk.com/ru/method/prettyCards.create
func (p *PrettyCards) Create(actor actor.Actor) *request.PrettyCardsCreateRequest {
	return request.NewPrettyCardsCreateRequest(p.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/prettyCards.delete
func (p *PrettyCards) Delete(actor actor.Actor) *request.PrettyCardsDeleteRequest {
	return request.NewPrettyCardsDeleteRequest(p.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/prettyCards.edit
func (p *PrettyCards) Edit(actor actor.Actor) *request.PrettyCardsEditRequest {
	return request.NewPrettyCardsEditRequest(p.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/prettyCards.get
func (p *PrettyCards) Get(actor actor.Actor) *request.PrettyCardsGetRequest {
	return request.NewPrettyCardsGetRequest(p.api, actor)
}

// GetByID Doc: https://dev.vk.com/ru/method/prettyCards.getById
func (p *PrettyCards) GetByID(actor actor.Actor) *request.PrettyCardsGetByIDRequest {
	return request.NewPrettyCardsGetByIDRequest(p.api, actor)
}

// GetUploadURL Doc: https://dev.vk.com/ru/method/prettyCards.getUploadURL
func (p *PrettyCards) GetUploadURL(actor actor.Actor) *request.PrettyCardsGetUploadURLRequest {
	return request.NewPrettyCardsGetUploadURLRequest(p.api, actor)
}
