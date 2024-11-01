package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Translations Doc: https://dev.vk.com/ru/method/users
type Translations struct {
	BaseAction
}

// Translate Doc: https://dev.vk.com/ru/method/users.translate
func (a *Translations) Translate(actor actor.Actor) *request.TranslationsTranslateRequest {
	return request.NewTranslationsTranslateRequest(a.api, actor)
}
