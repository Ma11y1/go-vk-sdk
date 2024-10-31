package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/storage

// StorageGetRequest defines the request for storage.get
//
// Returns the value of the variable specified by the key parameter.
// Doc: https://dev.vk.com/method/storage.get
type StorageGetRequest struct {
	BaseRequest
}

// NewStorageGetRequest creates a new request for storage.get
func NewStorageGetRequest(a *api.API, actor actor.Actor) *StorageGetRequest {
	return &StorageGetRequest{*NewMethodBaseRequest(a, actor, "storage.get")}
}

// Exec executes the request and unmarshals the response into StorageGetResponse
func (r *StorageGetRequest) Exec(ctx context.Context) (response response.StorageGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StorageGetKeysRequest defines the request for storage.getKeys
//
// Returns the names of all variables.
// Doc: https://dev.vk.com/method/storage.getKeys
type StorageGetKeysRequest struct {
	BaseRequest
}

// NewStorageGetKeysRequest creates a new request for storage.getKeys
func NewStorageGetKeysRequest(a *api.API, actor actor.Actor) *StorageGetKeysRequest {
	return &StorageGetKeysRequest{*NewMethodBaseRequest(a, actor, "storage.getKeys")}
}

// Exec executes the request and unmarshals the response into StorageGetKeysResponse
func (r *StorageGetKeysRequest) Exec(ctx context.Context) (response response.StorageGetKeysResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// StorageSetRequest defines the request for storage.set
//
// Sets the value of the variable specified by the key parameter. Variables are stored indefinitely.
// Doc: https://dev.vk.com/method/storage.set
type StorageSetRequest struct {
	BaseRequest
}

// NewStorageSetRequest creates a new request for storage.set
func NewStorageSetRequest(a *api.API, actor actor.Actor) *StorageSetRequest {
	return &StorageSetRequest{*NewMethodBaseRequest(a, actor, "storage.set")}
}

// Exec executes the request and unmarshals the response into StorageSetResponse
func (r *StorageSetRequest) Exec(ctx context.Context) (response response.StorageSetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
