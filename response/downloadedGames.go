package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/downloadedGames

type DownloadedGamesGetPaidStatusResponse struct {
	IsPaid objects.NumberFlagBool `json:"is_paid"`
}
