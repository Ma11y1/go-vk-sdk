package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/downloadedGames

type DownloadedGamesGetPaidStatusResponse struct {
	IsPaid objects.BoolInt `json:"is_paid"`
}
