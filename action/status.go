package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Status Doc: https://dev.vk.com/ru/method/status
type Status struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/ru/method/status.get
func (a *Status) Get(actor actor.Actor) *request.StatusGetRequest {
	return request.NewStatusGetRequest(a.api, actor)
}

// Set Doc: https://dev.vk.com/ru/method/status.set
func (a *Status) Set(actor actor.Actor) *request.StatusSetRequest {
	return request.NewStatusSetRequest(a.api, actor)
}
