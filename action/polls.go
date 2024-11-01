package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Polls Doc: https://dev.vk.com/ru/method/polls
type Polls struct {
	BaseAction
}

// AddVote Doc: https://dev.vk.com/ru/method/polls.addVote
func (p *Polls) AddVote(actor actor.Actor) *request.PollsAddVoteRequest {
	return request.NewPollsAddVoteRequest(p.api, actor)
}

// Create Doc: https://dev.vk.com/ru/method/polls.create
func (p *Polls) Create(actor actor.Actor) *request.PollsCreateRequest {
	return request.NewPollsCreateRequest(p.api, actor)
}

// DeleteVote Doc: https://dev.vk.com/ru/method/polls.deleteVote
func (p *Polls) DeleteVote(actor actor.Actor) *request.PollsDeleteVoteRequest {
	return request.NewPollsDeleteVoteRequest(p.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/polls.edit
func (p *Polls) Edit(actor actor.Actor) *request.PollsEditRequest {
	return request.NewPollsEditRequest(p.api, actor)
}

// GetBackgrounds Doc: https://dev.vk.com/ru/method/polls.getBackgrounds
func (p *Polls) GetBackgrounds(actor actor.Actor) *request.PollsGetBackgroundsRequest {
	return request.NewPollsGetBackgroundsRequest(p.api, actor)
}

// GetByID Doc: https://dev.vk.com/ru/method/polls.getById
func (p *Polls) GetByID(actor actor.Actor) *request.PollsGetByIDRequest {
	return request.NewPollsGetByIDRequest(p.api, actor)
}

// GetPhotoUploadServer Doc: https://dev.vk.com/ru/method/polls.getPhotoUploadServer
func (p *Polls) GetPhotoUploadServer(actor actor.Actor) *request.PollsGetPhotoUploadServerRequest {
	return request.NewPollsGetPhotoUploadServerRequest(p.api, actor)
}

// GetVoters Doc: https://dev.vk.com/ru/method/polls.getVoters
func (p *Polls) GetVoters(actor actor.Actor) *request.PollsGetVotersRequest {
	return request.NewPollsGetVotersRequest(p.api, actor)
}

// SavePhoto Doc: https://dev.vk.com/ru/method/polls.savePhoto
func (p *Polls) SavePhoto(actor actor.Actor) *request.PollsSavePhotoRequest {
	return request.NewPollsSavePhotoRequest(p.api, actor)
}
