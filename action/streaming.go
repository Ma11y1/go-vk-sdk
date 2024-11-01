package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Streaming Doc: https://dev.vk.com/ru/method/streaming
type Streaming struct {
	BaseAction
}

// GetServerURL Doc: https://dev.vk.com/method/streaming.getServerUrl
func (s *Streaming) GetServerURL(actor actor.Actor) *request.StreamingGetServerURLRequest {
	return request.NewStreamingGetServerURLRequest(s.api, actor)
}

// GetStats Doc: https://dev.vk.com/method/streaming.getStats
func (s *Streaming) GetStats(actor actor.Actor) *request.StreamingGetStatsRequest {
	return request.NewStreamingGetStatsRequest(s.api, actor)
}

// GetStem Doc: https://dev.vk.com/method/streaming.getStem
func (s *Streaming) GetStem(actor actor.Actor) *request.StreamingGetStemRequest {
	return request.NewStreamingGetStemRequest(s.api, actor)
}
