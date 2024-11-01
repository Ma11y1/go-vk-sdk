package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

type Auth struct {
	BaseAction
}

// Restore Doc: https://dev.vk.com/ru/method/auth.restore
func (a *Auth) Restore(actor actor.Actor) *request.AuthRestoreRequest {
	return request.NewAuthRestoreRequest(a.GetAPI(), actor)
}

func (a *Auth) UserDirectAuth() *request.AuthUserDirectRequest {
	return request.NewAuthUserDirectAuthRequest(a.GetAPI())
}

func (a *Auth) UserCodeFlow() *request.AuthUserCodeFlowRequest {
	return request.NewAuthUserCodeFlowRequest(a.GetAPI())
}

func (a *Auth) GroupCodeFlow() *request.AuthGroupCodeFlowRequest {
	return request.NewAuthGroupCodeFlowRequest(a.GetAPI())
}

func (a *Auth) VKIDCodeFlow() *request.AuthVKIDCodeFlowRequest {
	return request.NewAuthVKIDCodeFlowRequest(a.GetAPI())
}
