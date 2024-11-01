package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// UtilsCheckLinkRequest defines the request for utils.checkLink
//
// Checks if the link is malicious.
// Doc: https://dev.vk.com/method/utils.checkLink
type UtilsCheckLinkRequest struct {
	BaseRequest
}

// NewUtilsCheckLinkRequest creates a new request for utils.checkLink
func NewUtilsCheckLinkRequest(a *api.API, actor actor.Actor) *UtilsCheckLinkRequest {
	return &UtilsCheckLinkRequest{*NewMethodBaseRequest(a, actor, "utils.checkLink")}
}

// Exec executes the request and unmarshals the response into UtilsCheckLinkResponse
func (r *UtilsCheckLinkRequest) Exec(ctx context.Context) (response response.UtilsCheckLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsDeleteFromLastShortenedRequest defines the request for utils.deleteFromLastShortened
//
// Deletes a shortened link from the last shortened links list.
// Doc: https://dev.vk.com/method/utils.deleteFromLastShortened
type UtilsDeleteFromLastShortenedRequest struct {
	BaseRequest
}

// NewUtilsDeleteFromLastShortenedRequest creates a new request for utils.deleteFromLastShortened
func NewUtilsDeleteFromLastShortenedRequest(a *api.API, actor actor.Actor) *UtilsDeleteFromLastShortenedRequest {
	return &UtilsDeleteFromLastShortenedRequest{*NewMethodBaseRequest(a, actor, "utils.deleteFromLastShortened")}
}

// Exec executes the request and unmarshals the response into UtilsDeleteFromLastShortenedResponse
func (r *UtilsDeleteFromLastShortenedRequest) Exec(ctx context.Context) (response response.UtilsDeleteFromLastShortenedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsGetLastShortenedLinksRequest defines the request for utils.getLastShortenedLinks
//
// Retrieves the last shortened links.
// Doc: https://dev.vk.com/method/utils.getLastShortenedLinks
type UtilsGetLastShortenedLinksRequest struct {
	BaseRequest
}

// NewUtilsGetLastShortenedLinksRequest creates a new request for utils.getLastShortenedLinks
func NewUtilsGetLastShortenedLinksRequest(a *api.API, actor actor.Actor) *UtilsGetLastShortenedLinksRequest {
	return &UtilsGetLastShortenedLinksRequest{*NewMethodBaseRequest(a, actor, "utils.getLastShortenedLinks")}
}

// Exec executes the request and unmarshals the response into UtilsGetLastShortenedLinksResponse
func (r *UtilsGetLastShortenedLinksRequest) Exec(ctx context.Context) (response response.UtilsGetLastShortenedLinksResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsGetLinkStatsRequest defines the request for utils.getLinkStats
//
// Retrieves link statistics.
// Doc: https://dev.vk.com/method/utils.getLinkStats
type UtilsGetLinkStatsRequest struct {
	BaseRequest
}

// NewUtilsGetLinkStatsRequest creates a new request for utils.getLinkStats
func NewUtilsGetLinkStatsRequest(a *api.API, actor actor.Actor) *UtilsGetLinkStatsRequest {
	return &UtilsGetLinkStatsRequest{*NewMethodBaseRequest(a, actor, "utils.getLinkStats")}
}

// Exec executes the request and unmarshals the response into UtilsGetLinkStatsResponse
func (r *UtilsGetLinkStatsRequest) Exec(ctx context.Context) (response response.UtilsGetLinkStatsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsGetLinkStatsExtendedRequest defines the request for utils.getLinkStats
//
// Retrieves link statistics.
// Doc: https://dev.vk.com/method/utils.getLinkStats
type UtilsGetLinkStatsExtendedRequest struct {
	BaseRequest
}

// NewUtilsGetLinkStatsExtendedRequest creates a new request for utils.getLinkStats
func NewUtilsGetLinkStatsExtendedRequest(a *api.API, actor actor.Actor) *UtilsGetLinkStatsExtendedRequest {
	return &UtilsGetLinkStatsExtendedRequest{*NewMethodBaseRequest(a, actor, "utils.getLinkStats")}
}

// Exec executes the request and unmarshals the response into UtilsGetLinkStatsExtendedResponse
func (r *UtilsGetLinkStatsExtendedRequest) Exec(ctx context.Context) (response response.UtilsGetLinkStatsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsGetServerTimeRequest defines the request for utils.getServerTime
//
// Retrieves the current server time.
// Doc: https://dev.vk.com/method/utils.getServerTime
type UtilsGetServerTimeRequest struct {
	BaseRequest
}

// NewUtilsGetServerTimeRequest creates a new request for utils.getServerTime
func NewUtilsGetServerTimeRequest(a *api.API, actor actor.Actor) *UtilsGetServerTimeRequest {
	return &UtilsGetServerTimeRequest{*NewMethodBaseRequest(a, actor, "utils.getServerTime")}
}

// Exec executes the request and unmarshals the response into UtilsGetServerTimeResponse
func (r *UtilsGetServerTimeRequest) Exec(ctx context.Context) (response response.UtilsGetServerTimeResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsGetShortLinkRequest defines the request for utils.getShortLink
//
// Generates a shortened link.
// Doc: https://dev.vk.com/method/utils.getShortLink
type UtilsGetShortLinkRequest struct {
	BaseRequest
}

// NewUtilsGetShortLinkRequest creates a new request for utils.getShortLink
func NewUtilsGetShortLinkRequest(a *api.API, actor actor.Actor) *UtilsGetShortLinkRequest {
	return &UtilsGetShortLinkRequest{*NewMethodBaseRequest(a, actor, "utils.getShortLink")}
}

// Exec executes the request and unmarshals the response into UtilsGetShortLinkResponse
func (r *UtilsGetShortLinkRequest) Exec(ctx context.Context) (response response.UtilsGetShortLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UtilsResolveScreenNameRequest defines the request for utils.resolveScreenName
//
// Resolves screen name to object ID.
// Doc: https://dev.vk.com/method/utils.resolveScreenName
type UtilsResolveScreenNameRequest struct {
	BaseRequest
}

// NewUtilsResolveScreenNameRequest creates a new request for utils.resolveScreenName
func NewUtilsResolveScreenNameRequest(a *api.API, actor actor.Actor) *UtilsResolveScreenNameRequest {
	return &UtilsResolveScreenNameRequest{*NewMethodBaseRequest(a, actor, "utils.resolveScreenName")}
}

// Exec executes the request and unmarshals the response into UtilsResolveScreenNameResponse
func (r *UtilsResolveScreenNameRequest) Exec(ctx context.Context) (response response.UtilsResolveScreenNameResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
