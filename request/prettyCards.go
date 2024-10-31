package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/prettyCards

// PrettyCardsCreateRequest defines the request for prettyCards.create
//
// Creates a carousel card. The created card must be manually added to a carousel.
// Doc: https://dev.vk.com/method/prettyCards.create
type PrettyCardsCreateRequest struct {
	BaseRequest
}

// NewPrettyCardsCreateRequest creates a new request for prettyCards.create
func NewPrettyCardsCreateRequest(a *api.API, actor actor.Actor) *PrettyCardsCreateRequest {
	return &PrettyCardsCreateRequest{*NewMethodBaseRequest(a, actor, "prettyCards.create")}
}

// Exec executes the request and unmarshals the response into PrettyCardsCreateResponse
func (r *PrettyCardsCreateRequest) Exec(ctx context.Context) (response response.PrettyCardsCreateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PrettyCardsDeleteRequest defines the request for prettyCards.delete
//
// Deletes a carousel card.
// Doc: https://dev.vk.com/method/prettyCards.delete
type PrettyCardsDeleteRequest struct {
	BaseRequest
}

// NewPrettyCardsDeleteRequest creates a new request for prettyCards.delete
func NewPrettyCardsDeleteRequest(a *api.API, actor actor.Actor) *PrettyCardsDeleteRequest {
	return &PrettyCardsDeleteRequest{*NewMethodBaseRequest(a, actor, "prettyCards.delete")}
}

// Exec executes the request and unmarshals the response into PrettyCardsDeleteResponse
func (r *PrettyCardsDeleteRequest) Exec(ctx context.Context) (response response.PrettyCardsDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PrettyCardsEditRequest defines the request for prettyCards.edit
//
// Edits a carousel card.
// Doc: https://dev.vk.com/method/prettyCards.edit
type PrettyCardsEditRequest struct {
	BaseRequest
}

// NewPrettyCardsEditRequest creates a new request for prettyCards.edit
func NewPrettyCardsEditRequest(a *api.API, actor actor.Actor) *PrettyCardsEditRequest {
	return &PrettyCardsEditRequest{*NewMethodBaseRequest(a, actor, "prettyCards.edit")}
}

// Exec executes the request and unmarshals the response into PrettyCardsEditResponse
func (r *PrettyCardsEditRequest) Exec(ctx context.Context) (response response.PrettyCardsEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PrettyCardsGetRequest defines the request for prettyCards.get
//
// Returns unused carousel cards for the owner.
// Doc: https://dev.vk.com/method/prettyCards.get
type PrettyCardsGetRequest struct {
	BaseRequest
}

// NewPrettyCardsGetRequest creates a new request for prettyCards.get
func NewPrettyCardsGetRequest(a *api.API, actor actor.Actor) *PrettyCardsGetRequest {
	return &PrettyCardsGetRequest{*NewMethodBaseRequest(a, actor, "prettyCards.get")}
}

// Exec executes the request and unmarshals the response into PrettyCardsGetResponse
func (r *PrettyCardsGetRequest) Exec(ctx context.Context) (response response.PrettyCardsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PrettyCardsGetByIDRequest defines the request for prettyCards.getById
//
// Returns information about a carousel card by its ID.
// Doc: https://dev.vk.com/method/prettyCards.getById
type PrettyCardsGetByIDRequest struct {
	BaseRequest
}

// NewPrettyCardsGetByIdRequest creates a new request for prettyCards.getById
func NewPrettyCardsGetByIdRequest(a *api.API, actor actor.Actor) *PrettyCardsGetByIDRequest {
	return &PrettyCardsGetByIDRequest{*NewMethodBaseRequest(a, actor, "prettyCards.getById")}
}

// Exec executes the request and unmarshals the response into PrettyCardsGetByIDResponse
func (r *PrettyCardsGetByIDRequest) Exec(ctx context.Context) (response response.PrettyCardsGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// PrettyCardsGetUploadURLRequest defines the request for prettyCards.getUploadURL
//
// Returns the URL for uploading a photo for a carousel card.
// Doc: https://dev.vk.com/method/prettyCards.getUploadURL
type PrettyCardsGetUploadURLRequest struct {
	BaseRequest
}

// NewPrettyCardsGetUploadURLRequest creates a new request for prettyCards.getUploadURL
func NewPrettyCardsGetUploadURLRequest(a *api.API, actor actor.Actor) *PrettyCardsGetUploadURLRequest {
	return &PrettyCardsGetUploadURLRequest{*NewMethodBaseRequest(a, actor, "prettyCards.getUploadURL")}
}

// Exec executes the request and unmarshals the response into PrettyCardsGetUploadURLResponse
func (r *PrettyCardsGetUploadURLRequest) Exec(ctx context.Context) (response response.PrettyCardsGetUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
