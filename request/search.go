package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/search

// SearchGetHintsRequest defines the request for search.getHints
//
// The method allows you to get quick search results using an arbitrary substring.
// Doc: https://dev.vk.com/method/search.getHints
type SearchGetHintsRequest struct {
	BaseRequest
}

// NewSearchGetHintsRequest creates a new request for search.getHints
func NewSearchGetHintsRequest(a *api.API, actor actor.Actor) *SearchGetHintsRequest {
	return &SearchGetHintsRequest{*NewMethodBaseRequest(a, actor, "search.getHints")}
}

// Exec executes the request and unmarshals the response into SearchGetHintsResponse
func (r *SearchGetHintsRequest) Exec(ctx context.Context) (response response.SearchGetHintsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
