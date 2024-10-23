package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/stats

type StatsGetResponse struct {
	BaseResponse
	Response []objects.StatsPeriod `json:"response"`
}

type StatsGetPostReachResponse struct {
	BaseResponse
	Response []objects.StatsWallpostStat `json:"response"`
}

type StatsTrackVisitorResponse struct {
	BaseResponse
	Response int `json:"response"`
}
