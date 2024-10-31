package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/stats

// StatsGetRequest defines the request for stats.get
//
// Returns community or application statistics.
// Doc: https://dev.vk.com/method/stats.get
type StatsGetRequest struct {
	BaseRequest
}

// NewStatsGetRequest creates a new request for stats.get
func NewStatsGetRequest(a *api.API, actor actor.Actor) *StatsGetRequest {
	return &StatsGetRequest{*NewMethodBaseRequest(a, actor, "stats.get")}
}

// Exec executes the request and unmarshals the response into StatsGetResponse
func (r *StatsGetRequest) Exec(ctx context.Context) (response response.StatsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StatsGetPostReachRequest defines the request for stats.getPostReach
//
// Returns statistics for a wall post.
// Doc: https://dev.vk.com/method/stats.getPostReach
type StatsGetPostReachRequest struct {
	BaseRequest
}

// NewStatsGetPostReachRequest creates a new request for stats.getPostReach
func NewStatsGetPostReachRequest(a *api.API, actor actor.Actor) *StatsGetPostReachRequest {
	return &StatsGetPostReachRequest{*NewMethodBaseRequest(a, actor, "stats.getPostReach")}
}

// Exec executes the request and unmarshals the response into StatsGetPostReachResponse
func (r *StatsGetPostReachRequest) Exec(ctx context.Context) (response response.StatsGetPostReachResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StatsTrackVisitorRequest defines the request for stats.trackVisitor
//
// Adds data about the current session to the application's visit statistics.
// Doc: https://dev.vk.com/method/stats.trackVisitor
type StatsTrackVisitorRequest struct {
	BaseRequest
}

// NewStatsTrackVisitorRequest creates a new request for stats.trackVisitor
func NewStatsTrackVisitorRequest(a *api.API, actor actor.Actor) *StatsTrackVisitorRequest {
	return &StatsTrackVisitorRequest{*NewMethodBaseRequest(a, actor, "stats.trackVisitor")}
}

// Exec executes the request and unmarshals the response into StatsTrackVisitorResponse
func (r *StatsTrackVisitorRequest) Exec(ctx context.Context) (response response.StatsTrackVisitorResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
