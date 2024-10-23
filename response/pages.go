package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/pages

type PagesClearCacheResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PagesGetResponse struct {
	BaseResponse
	Response objects.PagesWikipageFull `json:"response"`
}

type PagesGetHistoryResponse struct {
	BaseResponse
	Response []objects.PagesWikipageHistory `json:"response"`
}

type PagesGetTitlesResponse struct {
	BaseResponse
	Response []objects.PagesWikipageFull `json:"response"`
}

type PagesGetVersionResponse struct {
	BaseResponse
	Response objects.PagesWikipageFull `json:"response"`
}

type PagesParseWikiResponse struct {
	BaseResponse
	Response string `json:"response"`
}

type PagesSaveResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type PagesSaveAccessResponse struct {
	BaseResponse
	Response int `json:"response"`
}
