package response

import (
	"go-vk-sdk/constants"
)

// Doc: https://dev.vk.com/ru/method/translations

type TranslationsTranslateResponse struct {
	BaseResponse
	Response struct {
		Texts      []string                          `json:"texts"`
		SourceLang constants.LanguageTranslationName `json:"source_lang"`
	} `json:"response"`
}
