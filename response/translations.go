package response

import (
	"go-vk-sdk/constants"
)

// Doc: https://dev.vk.com/ru/method/translations

type TranslationsTranslateResponse struct {
	BaseResponse
	Response struct {
		Texts      []string                          `json:"texts"`
		SourceLang constants.LanguageTranslationType `json:"source_lang"`
	} `json:"response"`
}
