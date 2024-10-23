package action

import (
	"go-vk-sdk/request"
)

type Auth struct {
	BaseAction
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
