package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Store Doc: https://dev.vk.com/ru/method/store
type Store struct {
	BaseAction
}

// AddStickersToFavorite Doc: https://dev.vk.com/method/store.addStickersToFavorite
func (s *Store) AddStickersToFavorite(actor actor.Actor) *request.StoreAddStickersToFavoriteRequest {
	return request.NewStoreAddStickersToFavoriteRequest(s.api, actor)
}

// GetFavoriteStickers Doc: https://dev.vk.com/method/store.getFavoriteStickers
func (s *Store) GetFavoriteStickers(actor actor.Actor) *request.StoreGetFavoriteStickersRequest {
	return request.NewStoreGetFavoriteStickersRequest(s.api, actor)
}

// GetProducts Doc: https://dev.vk.com/method/store.getProducts
func (s *Store) GetProducts(actor actor.Actor) *request.StoreGetProductsRequest {
	return request.NewStoreGetProductsRequest(s.api, actor)
}

// GetStickersKeywords Doc: https://dev.vk.com/method/store.getStickersKeywords
func (s *Store) GetStickersKeywords(actor actor.Actor) *request.StoreGetStickersKeywordsRequest {
	return request.NewStoreGetStickersKeywordsRequest(s.api, actor)
}

// RemoveStickersFromFavorite Doc: https://dev.vk.com/method/store.removeStickersFromFavorite
func (s *Store) RemoveStickersFromFavorite(actor actor.Actor) *request.StoreRemoveStickersFromFavoriteRequest {
	return request.NewStoreRemoveStickersFromFavoriteRequest(s.api, actor)
}
