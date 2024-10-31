package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/store

// StoreAddStickersToFavoriteRequest defines the request for store.addStickersToFavorite
//
// Adds a sticker to favorites.
// Doc: https://dev.vk.com/method/store.addStickersToFavorite
type StoreAddStickersToFavoriteRequest struct {
	BaseRequest
}

// NewStoreAddStickersToFavoriteRequest creates a new request for store.addStickersToFavorite
func NewStoreAddStickersToFavoriteRequest(a *api.API, actor actor.Actor) *StoreAddStickersToFavoriteRequest {
	return &StoreAddStickersToFavoriteRequest{*NewMethodBaseRequest(a, actor, "store.addStickersToFavorite")}
}

// Exec executes the request and unmarshals the response into StoreAddStickersToFavoriteResponse
func (r *StoreAddStickersToFavoriteRequest) Exec(ctx context.Context) (response response.StoreAddStickersToFavoriteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoreGetFavoriteStickersRequest defines the request for store.getFavoriteStickers
//
// Returns a list of favorite stickers.
// Doc: https://dev.vk.com/method/store.getFavoriteStickers
type StoreGetFavoriteStickersRequest struct {
	BaseRequest
}

// NewStoreGetFavoriteStickersRequest creates a new request for store.getFavoriteStickers
func NewStoreGetFavoriteStickersRequest(a *api.API, actor actor.Actor) *StoreGetFavoriteStickersRequest {
	return &StoreGetFavoriteStickersRequest{*NewMethodBaseRequest(a, actor, "store.getFavoriteStickers")}
}

// Exec executes the request and unmarshals the response into StoreGetFavoriteStickersResponse
func (r *StoreGetFavoriteStickersRequest) Exec(ctx context.Context) (response response.StoreGetFavoriteStickersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoreGetProductsRequest defines the request for store.getProducts
//
// Returns a list of products.
// Doc: https://dev.vk.com/method/store.getProducts
type StoreGetProductsRequest struct {
	BaseRequest
}

// NewStoreGetProductsRequest creates a new request for store.getProducts
func NewStoreGetProductsRequest(a *api.API, actor actor.Actor) *StoreGetProductsRequest {
	return &StoreGetProductsRequest{*NewMethodBaseRequest(a, actor, "store.getProducts")}
}

// Exec executes the request and unmarshals the response into StoreGetProductsResponse
func (r *StoreGetProductsRequest) Exec(ctx context.Context) (response response.StoreGetProductsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoreGetStickersKeywordsRequest defines the request for store.getStickersKeywords
//
// Returns a list of keywords for stickers.
// Doc: https://dev.vk.com/method/store.getStickersKeywords
type StoreGetStickersKeywordsRequest struct {
	BaseRequest
}

// NewStoreGetStickersKeywordsRequest creates a new request for store.getStickersKeywords
func NewStoreGetStickersKeywordsRequest(a *api.API, actor actor.Actor) *StoreGetStickersKeywordsRequest {
	return &StoreGetStickersKeywordsRequest{*NewMethodBaseRequest(a, actor, "store.getStickersKeywords")}
}

// Exec executes the request and unmarshals the response into StoreGetStickersKeywordsResponse
func (r *StoreGetStickersKeywordsRequest) Exec(ctx context.Context) (response response.StoreGetStickersKeywordsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StoreRemoveStickersFromFavoriteRequest defines the request for store.removeStickersFromFavorite
//
// Removes a sticker from favorites.
// Doc: https://dev.vk.com/method/store.removeStickersFromFavorite
type StoreRemoveStickersFromFavoriteRequest struct {
	BaseRequest
}

// NewStoreRemoveStickersFromFavoriteRequest creates a new request for store.removeStickersFromFavorite
func NewStoreRemoveStickersFromFavoriteRequest(a *api.API, actor actor.Actor) *StoreRemoveStickersFromFavoriteRequest {
	return &StoreRemoveStickersFromFavoriteRequest{*NewMethodBaseRequest(a, actor, "store.removeStickersFromFavorite")}
}

// Exec executes the request and unmarshals the response into StoreRemoveStickersFromFavoriteResponse
func (r *StoreRemoveStickersFromFavoriteRequest) Exec(ctx context.Context) (response response.StoreRemoveStickersFromFavoriteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
