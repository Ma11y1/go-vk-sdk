package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Calls Doc: https://dev.vk.com/ru/method/calls
type Calls struct {
	BaseAction
}

// ForceFinish Doc: https://dev.vk.com/ru/method/calls.forceFinish
func (c *Calls) ForceFinish(actor actor.Actor) *request.CallsForceFinishRequest {
	return request.NewCallsForceFinishRequest(c.api, actor)
}

// Start Doc: https://dev.vk.com/ru/method/calls.start
func (c *Calls) Start(actor actor.Actor) *request.CallsStartRequest {
	return request.NewCallsStartRequest(c.api, actor)
}
