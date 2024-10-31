package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/translations

// TranslationsTranslateRequest defines the request for translations.translate
//
//	The method can be called no more than 10 times per minute.
//	Doc: https://dev.vk.com/method/translations.translate
type TranslationsTranslateRequest struct {
	BaseRequest
}

// NewTranslationsTranslateRequest creates a new request for translations.translate
func NewTranslationsTranslateRequest(a *api.API, actor actor.Actor) *TranslationsTranslateRequest {
	return &TranslationsTranslateRequest{*NewMethodBaseRequest(a, actor, "translations.translate")}
}

// Exec executes the request and unmarshals the response into TranslationsTranslateResponse
func (r *TranslationsTranslateRequest) Exec(ctx context.Context) (response response.TranslationsTranslateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
