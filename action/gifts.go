package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Gifts Doc: https://dev.vk.com/ru/method/gifts
type Gifts struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/ru/method/gifts.get
func (a *Gifts) Get(actor actor.Actor) *request.GiftsGet {
	return request.NewGiftsGetRequest(a.api, actor)
}
