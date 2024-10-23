package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/downloadedGames

// DownloadedGamesGetPaidStatusRequest defines the request for downloadedGames.getPaidStatus
//
// The method gets information whether the user bought the game or not.
// Doc: https://dev.vk.com/method/downloadedGames.getPaidStatus
type DownloadedGamesGetPaidStatusRequest struct {
	BaseRequest
}

// NewDownloadedGamesGetPaidStatusRequest creates a new request for downloadedGames.getPaidStatus
func NewDownloadedGamesGetPaidStatusRequest(a *api.API, actor actor.Actor) *DownloadedGamesGetPaidStatusRequest {
	return &DownloadedGamesGetPaidStatusRequest{*NewMethodBaseRequest(a, actor, "downloadedGames.getPaidStatus")}
}

// Exec executes the request and unmarshals the response into DownloadedGamesGetPaidStatusResponse
func (r *DownloadedGamesGetPaidStatusRequest) Exec(ctx context.Context) (response response.DownloadedGamesGetPaidStatusResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
