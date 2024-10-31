package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/streaming

// StreamingGetServerUrlRequest defines the request for streaming.getServerUrl
//
// Retrieves data for connecting to the Streaming API.
// Doc: https://dev.vk.com/method/streaming.getServerUrl
type StreamingGetServerUrlRequest struct {
	BaseRequest
}

// NewStreamingGetServerUrlRequest creates a new request for streaming.getServerUrl
func NewStreamingGetServerUrlRequest(a *api.API, actor actor.Actor) *StreamingGetServerUrlRequest {
	return &StreamingGetServerUrlRequest{*NewMethodBaseRequest(a, actor, "streaming.getServerUrl")}
}

// Exec executes the request and unmarshals the response into StreamingGetServerUrlResponse
func (r *StreamingGetServerUrlRequest) Exec(ctx context.Context) (response response.StreamingGetServerURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StreamingGetStatsRequest defines the request for streaming.getStats
//
// Retrieves statistics for prepared and delivered Streaming API events.
// Doc: https://dev.vk.com/method/streaming.getStats
type StreamingGetStatsRequest struct {
	BaseRequest
}

// NewStreamingGetStatsRequest creates a new request for streaming.getStats
func NewStreamingGetStatsRequest(a *api.API, actor actor.Actor) *StreamingGetStatsRequest {
	return &StreamingGetStatsRequest{*NewMethodBaseRequest(a, actor, "streaming.getStats")}
}

// Exec executes the request and unmarshals the response into StreamingGetStatsResponse
func (r *StreamingGetStatsRequest) Exec(ctx context.Context) (response response.StreamingGetStatsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StreamingGetStemRequest defines the request for streaming.getStem
//
// Retrieves the stem of a word.
// Doc: https://dev.vk.com/method/streaming.getStem
type StreamingGetStemRequest struct {
	BaseRequest
}

// NewStreamingGetStemRequest creates a new request for streaming.getStem
func NewStreamingGetStemRequest(a *api.API, actor actor.Actor) *StreamingGetStemRequest {
	return &StreamingGetStemRequest{*NewMethodBaseRequest(a, actor, "streaming.getStem")}
}

// Exec executes the request and unmarshals the response into StreamingGetStemResponse
func (r *StreamingGetStemRequest) Exec(ctx context.Context) (response response.StreamingGetStemResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
