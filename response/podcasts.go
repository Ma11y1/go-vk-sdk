package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/podcasts

type PodcastsSearchPodcastResponse struct {
	BaseResponse
	Response struct {
		Podcasts     []objects.PodcastsSearchExternalData `json:"podcasts"`
		ResultsTotal int                                  `json:"results_total"`
	} `json:"response"`
}
