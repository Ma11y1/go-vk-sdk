package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/pages

// PagesClearCacheRequest defines the request for pages.clearCache
//
// Clears the cache of specific external pages that may be attached to VK posts.
// After clearing the cache, the data about the page will be updated when a link is
// attached to a post again.
// Doc: https://dev.vk.com/method/pages.clearCache
type PagesClearCacheRequest struct {
	BaseRequest
}

// NewPagesClearCacheRequest creates a new request for pages.clearCache
func NewPagesClearCacheRequest(a *api.API, actor actor.Actor) *PagesClearCacheRequest {
	return &PagesClearCacheRequest{*NewMethodBaseRequest(a, actor, "pages.clearCache")}
}

// Exec executes the request and unmarshals the response into PagesClearCacheResponse
func (r *PagesClearCacheRequest) Exec(ctx context.Context) (response response.PagesClearCacheResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesGetRequest defines the request for pages.get
//
// Returns information about a wiki page.
// Doc: https://dev.vk.com/method/pages.get
type PagesGetRequest struct {
	BaseRequest
}

// NewPagesGetRequest creates a new request for pages.get
func NewPagesGetRequest(a *api.API, actor actor.Actor) *PagesGetRequest {
	return &PagesGetRequest{*NewMethodBaseRequest(a, actor, "pages.get")}
}

// Exec executes the request and unmarshals the response into PagesGetResponse
func (r *PagesGetRequest) Exec(ctx context.Context) (response response.PagesGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesGetHistoryRequest defines the request for pages.getHistory
//
// Returns a list of all old versions of a wiki page.
// Doc: https://dev.vk.com/method/pages.getHistory
type PagesGetHistoryRequest struct {
	BaseRequest
}

// NewPagesGetHistoryRequest creates a new request for pages.getHistory
func NewPagesGetHistoryRequest(a *api.API, actor actor.Actor) *PagesGetHistoryRequest {
	return &PagesGetHistoryRequest{*NewMethodBaseRequest(a, actor, "pages.getHistory")}
}

// Exec executes the request and unmarshals the response into PagesGetHistoryResponse
func (r *PagesGetHistoryRequest) Exec(ctx context.Context) (response response.PagesGetHistoryResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesGetTitlesRequest defines the request for pages.getTitles
//
// Returns a list of wiki pages in a group.
// Doc: https://dev.vk.com/method/pages.getTitles
type PagesGetTitlesRequest struct {
	BaseRequest
}

// NewPagesGetTitlesRequest creates a new request for pages.getTitles
func NewPagesGetTitlesRequest(a *api.API, actor actor.Actor) *PagesGetTitlesRequest {
	return &PagesGetTitlesRequest{*NewMethodBaseRequest(a, actor, "pages.getTitles")}
}

// Exec executes the request and unmarshals the response into PagesGetTitlesResponse
func (r *PagesGetTitlesRequest) Exec(ctx context.Context) (response response.PagesGetTitlesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesGetVersionRequest defines the request for pages.getVersion
//
// Returns the text of one of the old versions of a page.
// Doc: https://dev.vk.com/method/pages.getVersion
type PagesGetVersionRequest struct {
	BaseRequest
}

// NewPagesGetVersionRequest creates a new request for pages.getVersion
func NewPagesGetVersionRequest(a *api.API, actor actor.Actor) *PagesGetVersionRequest {
	return &PagesGetVersionRequest{*NewMethodBaseRequest(a, actor, "pages.getVersion")}
}

// Exec executes the request and unmarshals the response into PagesGetVersionResponse
func (r *PagesGetVersionRequest) Exec(ctx context.Context) (response response.PagesGetVersionResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesParseWikiRequest defines the request for pages.parseWiki
//
// Gets the HTML representation of wiki markup.
// Doc: https://dev.vk.com/method/pages.parseWiki
type PagesParseWikiRequest struct {
	BaseRequest
}

// NewPagesParseWikiRequest creates a new request for pages.parseWiki
func NewPagesParseWikiRequest(a *api.API, actor actor.Actor) *PagesParseWikiRequest {
	return &PagesParseWikiRequest{*NewMethodBaseRequest(a, actor, "pages.parseWiki")}
}

// Exec executes the request and unmarshals the response into PagesParseWikiResponse
func (r *PagesParseWikiRequest) Exec(ctx context.Context) (response response.PagesParseWikiResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesSaveRequest defines the request for pages.save
//
// Saves the text of a wiki page.
// Doc: https://dev.vk.com/method/pages.save
type PagesSaveRequest struct {
	BaseRequest
}

// NewPagesSaveRequest creates a new request for pages.save
func NewPagesSaveRequest(a *api.API, actor actor.Actor) *PagesSaveRequest {
	return &PagesSaveRequest{*NewMethodBaseRequest(a, actor, "pages.save")}
}

// Exec executes the request and unmarshals the response into PagesSaveResponse
func (r *PagesSaveRequest) Exec(ctx context.Context) (response response.PagesSaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PagesSaveAccessRequest defines the request for pages.saveAccess
//
// Saves new access settings for reading and editing a wiki page.
// Doc: https://dev.vk.com/method/pages.saveAccess
type PagesSaveAccessRequest struct {
	BaseRequest
}

// NewPagesSaveAccessRequest creates a new request for pages.saveAccess
func NewPagesSaveAccessRequest(a *api.API, actor actor.Actor) *PagesSaveAccessRequest {
	return &PagesSaveAccessRequest{*NewMethodBaseRequest(a, actor, "pages.saveAccess")}
}

// Exec executes the request and unmarshals the response into PagesSaveAccessResponse
func (r *PagesSaveAccessRequest) Exec(ctx context.Context) (response response.PagesSaveAccessResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
