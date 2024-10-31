package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/podcasts

// PodcastsSearchPodcastRequest defines the request for podcasts.searchPodcast
//
// Doc: https://dev.vk.com/ru/method/podcasts.searchPodcast
type PodcastsSearchPodcastRequest struct {
	BaseRequest
}

// NewPodcastsSearchPodcastRequest creates a new request for podcasts.searchPodcast
func NewPodcastsSearchPodcastRequest(a *api.API, actor actor.Actor) *PodcastsSearchPodcastRequest {
	return &PodcastsSearchPodcastRequest{*NewMethodBaseRequest(a, actor, "podcasts.searchPodcast")}
}

// Exec executes the request and unmarshals the response into PodcastsSearchPodcastResponse
func (r *PodcastsSearchPodcastRequest) Exec(ctx context.Context) (response response.PodcastsSearchPodcastResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
