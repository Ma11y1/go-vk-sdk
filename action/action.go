package action

import "go-vk-sdk/api"

type Action interface {
	SetAPI(*api.API)
	GetAPI() *api.API
}

type BaseAction struct {
	api *api.API
}

func (a *BaseAction) SetAPI(api *api.API) {
	if api != nil {
		a.api = api
	}
}

func (a *BaseAction) GetAPI() *api.API {
	return a.api
}
