package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Search Doc: https://dev.vk.com/ru/method/search
type Search struct {
	BaseAction
}

// GetHints Doc: https://dev.vk.com/ru/method/search.getHints
func (a *Search) GetHints(actor actor.Actor) *request.SearchGetHintsRequest {
	return request.NewSearchGetHintsRequest(a.api, actor)
}
