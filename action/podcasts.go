package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Podcasts Doc: https://dev.vk.com/ru/method/podcasts
type Podcasts struct {
	BaseAction
}

// SearchPodcast Doc: https://dev.vk.com/ru/method/podcasts.searchPodcast
func (p *Podcasts) SearchPodcast(actor actor.Actor) *request.PodcastsSearchPodcastRequest {
	return request.NewPodcastsSearchPodcastRequest(p.api, actor)
}
