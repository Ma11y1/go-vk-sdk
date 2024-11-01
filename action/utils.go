package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Utils Doc: https://dev.vk.com/ru/method/utils
type Utils struct {
	BaseAction
}

// CheckLink Doc: https://dev.vk.com/method/utils.checkLink
func (u *Utils) CheckLink(actor actor.Actor) *request.UtilsCheckLinkRequest {
	return request.NewUtilsCheckLinkRequest(u.api, actor)
}

// DeleteFromLastShortened Doc: https://dev.vk.com/method/utils.deleteFromLastShortened
func (u *Utils) DeleteFromLastShortened(actor actor.Actor) *request.UtilsDeleteFromLastShortenedRequest {
	return request.NewUtilsDeleteFromLastShortenedRequest(u.api, actor)
}

// GetLastShortenedLinks Doc: https://dev.vk.com/method/utils.getLastShortenedLinks
func (u *Utils) GetLastShortenedLinks(actor actor.Actor) *request.UtilsGetLastShortenedLinksRequest {
	return request.NewUtilsGetLastShortenedLinksRequest(u.api, actor)
}

// GetLinkStats Doc: https://dev.vk.com/method/utils.getLinkStats
func (u *Utils) GetLinkStats(actor actor.Actor) *request.UtilsGetLinkStatsRequest {
	return request.NewUtilsGetLinkStatsRequest(u.api, actor)
}

// GetLinkStatsExtended Doc: https://dev.vk.com/method/utils.getLinkStats
func (u *Utils) GetLinkStatsExtended(actor actor.Actor) *request.UtilsGetLinkStatsExtendedRequest {
	return request.NewUtilsGetLinkStatsExtendedRequest(u.api, actor)
}

// GetServerTime Doc: https://dev.vk.com/method/utils.getServerTime
func (u *Utils) GetServerTime(actor actor.Actor) *request.UtilsGetServerTimeRequest {
	return request.NewUtilsGetServerTimeRequest(u.api, actor)
}

// GetShortLink Doc: https://dev.vk.com/method/utils.getShortLink
func (u *Utils) GetShortLink(actor actor.Actor) *request.UtilsGetShortLinkRequest {
	return request.NewUtilsGetShortLinkRequest(u.api, actor)
}

// ResolveScreenName Doc: https://dev.vk.com/method/utils.resolveScreenName
func (u *Utils) ResolveScreenName(actor actor.Actor) *request.UtilsResolveScreenNameRequest {
	return request.NewUtilsResolveScreenNameRequest(u.api, actor)
}
