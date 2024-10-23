package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/gifts

// GiftsGet defines the request for gifts.get
//
// Returns a list of the user's received gifts.
// Doc: https://dev.vk.com/method/gifts.get
type GiftsGet struct {
	BaseRequest
}

// NewGiftsGetRequest creates a new request for gifts.get
func NewGiftsGetRequest(a *api.API, actor actor.Actor) *GiftsGet {
	return &GiftsGet{*NewMethodBaseRequest(a, actor, "gifts.get")}
}

// Exec executes the request and unmarshals the response into GiftsGetResponse
func (r *GiftsGet) Exec(ctx context.Context) (response response.GiftsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
