package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Storage Doc: https://dev.vk.com/ru/method/storage
type Storage struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/ru/method/storage.get
func (a *Storage) Get(actor actor.Actor) *request.StorageGetRequest {
	return request.NewStorageGetRequest(a.api, actor)
}

// GetKeys Doc: https://dev.vk.com/ru/method/storage.getKeys
func (a *Storage) GetKeys(actor actor.Actor) *request.StorageGetKeysRequest {
	return request.NewStorageGetKeysRequest(a.api, actor)
}

// Set Doc: https://dev.vk.com/ru/method/storage.set
func (a *Storage) Set(actor actor.Actor) *request.StorageSetRequest {
	return request.NewStorageSetRequest(a.api, actor)
}
