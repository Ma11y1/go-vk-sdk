package response

import (
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/method/utils

type UtilsCheckLinkResponse struct {
	BaseResponse
	Response objects.UtilsLinkChecked `json:"response"`
}

type UtilsDeleteFromLastShortenedResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type UtilsGetLastShortenedLinksResponse struct {
	BaseResponse
	Response struct {
		Count int                              `json:"count"`
		Items []objects.UtilsLastShortenedLink `json:"items"`
	} `json:"response"`
}

type UtilsGetLinkStatsResponse struct {
	BaseResponse
	Response objects.UtilsLinkStats `json:"response"`
}

type UtilsGetLinkStatsExtendedResponse struct {
	BaseResponse
	Response objects.UtilsLinkStatsExtended `json:"response"`
}

type UtilsGetServerTimeResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type UtilsGetShortLinkResponse struct {
	BaseResponse
	Response objects.UtilsShortLink `json:"response"`
}

type UtilsResolveScreenNameResponse struct {
	BaseResponse
	Response objects.UtilsDomainResolved `json:"response"`
}
