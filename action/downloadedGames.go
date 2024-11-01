package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// DownloadedGames Doc: https://dev.vk.com/ru/method/downloadedGames
type DownloadedGames struct {
	BaseAction
}

// GetPaidStatus Doc: https://dev.vk.com/ru/method/downloadedGames.getPaidStatus
func (a *DownloadedGames) GetPaidStatus(actor actor.Actor) *request.DownloadedGamesGetPaidStatusRequest {
	return request.NewDownloadedGamesGetPaidStatusRequest(a.api, actor)
}
