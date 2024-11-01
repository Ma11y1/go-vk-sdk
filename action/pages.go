package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Pages Doc: https://dev.vk.com/ru/method/pages
type Pages struct {
	BaseAction
}

// ClearCache Doc: https://dev.vk.com/ru/method/pages.clearCache
func (p *Pages) ClearCache(actor actor.Actor) *request.PagesClearCacheRequest {
	return request.NewPagesClearCacheRequest(p.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/pages.get
func (p *Pages) Get(actor actor.Actor) *request.PagesGetRequest {
	return request.NewPagesGetRequest(p.api, actor)
}

// GetHistory Doc: https://dev.vk.com/ru/method/pages.getHistory
func (p *Pages) GetHistory(actor actor.Actor) *request.PagesGetHistoryRequest {
	return request.NewPagesGetHistoryRequest(p.api, actor)
}

// GetTitles Doc: https://dev.vk.com/ru/method/pages.getTitles
func (p *Pages) GetTitles(actor actor.Actor) *request.PagesGetTitlesRequest {
	return request.NewPagesGetTitlesRequest(p.api, actor)
}

// GetVersion Doc: https://dev.vk.com/ru/method/pages.getVersion
func (p *Pages) GetVersion(actor actor.Actor) *request.PagesGetVersionRequest {
	return request.NewPagesGetVersionRequest(p.api, actor)
}

// ParseWiki Doc: https://dev.vk.com/ru/method/pages.parseWiki
func (p *Pages) ParseWiki(actor actor.Actor) *request.PagesParseWikiRequest {
	return request.NewPagesParseWikiRequest(p.api, actor)
}

// Save Doc: https://dev.vk.com/ru/method/pages.save
func (p *Pages) Save(actor actor.Actor) *request.PagesSaveRequest {
	return request.NewPagesSaveRequest(p.api, actor)
}

// SaveAccess Doc: https://dev.vk.com/ru/method/pages.saveAccess
func (p *Pages) SaveAccess(actor actor.Actor) *request.PagesSaveAccessRequest {
	return request.NewPagesSaveAccessRequest(p.api, actor)
}
