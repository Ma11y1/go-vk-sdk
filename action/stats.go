package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Stats Doc: https://dev.vk.com/ru/method/stats
type Stats struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/ru/method/stats.get
func (s *Stats) Get(actor actor.Actor) *request.StatsGetRequest {
	return request.NewStatsGetRequest(s.api, actor)
}

// GetPostReach Doc: https://dev.vk.com/ru/method/stats.getPostReach
func (s *Stats) GetPostReach(actor actor.Actor) *request.StatsGetPostReachRequest {
	return request.NewStatsGetPostReachRequest(s.api, actor)
}

// TrackVisitor Doc: https://dev.vk.com/ru/method/stats.trackVisitor
func (s *Stats) TrackVisitor(actor actor.Actor) *request.StatsTrackVisitorRequest {
	return request.NewStatsTrackVisitorRequest(s.api, actor)
}
