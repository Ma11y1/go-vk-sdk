package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Widgets Doc: https://dev.vk.com/ru/method/widgets
type Widgets struct {
	BaseAction
}

// GetComments Doc: https://dev.vk.com/method/widgets.getComments
func (w *Widgets) GetComments(actor actor.Actor) *request.WidgetsGetCommentsRequest {
	return request.NewWidgetsGetCommentsRequest(w.api, actor)
}

// GetPages Doc: https://dev.vk.com/method/widgets.getPages
func (w *Widgets) GetPages(actor actor.Actor) *request.WidgetsGetPagesRequest {
	return request.NewWidgetsGetPagesRequest(w.api, actor)
}
