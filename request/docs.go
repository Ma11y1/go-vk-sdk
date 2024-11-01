package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/docs

// DocsAddRequest defines the request for docs.add
//
// The method copies a document to the current user's documents.
// Doc: https://dev.vk.com/method/docs.add
type DocsAddRequest struct {
	BaseRequest
}

// NewDocsAddRequest creates a new request for docs.add
func NewDocsAddRequest(a *api.API, actor actor.Actor) *DocsAddRequest {
	return &DocsAddRequest{*NewMethodBaseRequest(a, actor, "docs.add")}
}

// Exec executes the request and unmarshals the response into DocsAddResponse
func (r *DocsAddRequest) Exec(ctx context.Context) (response response.DocsAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsDeleteRequest defines the request for docs.delete
//
// The method deletes a user's or group's document.
// Doc: https://dev.vk.com/method/docs.delete
type DocsDeleteRequest struct {
	BaseRequest
}

// NewDocsDeleteRequest creates a new request for docs.delete
func NewDocsDeleteRequest(a *api.API, actor actor.Actor) *DocsDeleteRequest {
	return &DocsDeleteRequest{*NewMethodBaseRequest(a, actor, "docs.delete")}
}

// Exec executes the request and unmarshals the response into DocsDeleteResponse
func (r *DocsDeleteRequest) Exec(ctx context.Context) (response response.DocsDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsEditRequest defines the request for docs.edit
//
// The method edits a user's or group's document.
// Doc: https://dev.vk.com/method/docs.edit
type DocsEditRequest struct {
	BaseRequest
}

// NewDocsEditRequest creates a new request for docs.edit
func NewDocsEditRequest(a *api.API, actor actor.Actor) *DocsEditRequest {
	return &DocsEditRequest{*NewMethodBaseRequest(a, actor, "docs.edit")}
}

// Exec executes the request and unmarshals the response into DocsEditResponse
func (r *DocsEditRequest) Exec(ctx context.Context) (response response.DocsEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetRequest defines the request for docs.get
//
// The method returns detailed information about user or community documents.
// Doc: https://dev.vk.com/method/docs.get
type DocsGetRequest struct {
	BaseRequest
}

// NewDocsGetRequest creates a new request for docs.get
func NewDocsGetRequest(a *api.API, actor actor.Actor) *DocsGetRequest {
	return &DocsGetRequest{*NewMethodBaseRequest(a, actor, "docs.get")}
}

// Exec executes the request and unmarshals the response into DocsGetResponse
func (r *DocsGetRequest) Exec(ctx context.Context) (response response.DocsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetByIDRequest defines the request for docs.getById
//
// The method returns information about documents by their IDs.
// Doc: https://dev.vk.com/method/docs.getById
type DocsGetByIDRequest struct {
	BaseRequest
}

// NewDocsGetByIDRequest creates a new request for docs.getById
func NewDocsGetByIDRequest(a *api.API, actor actor.Actor) *DocsGetByIDRequest {
	return &DocsGetByIDRequest{*NewMethodBaseRequest(a, actor, "docs.getById")}
}

// Exec executes the request and unmarshals the response into DocsGetByIDResponse
func (r *DocsGetByIDRequest) Exec(ctx context.Context) (response response.DocsGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetMessagesUploadServerRequest defines the request for docs.getMessagesUploadServer
//
// The method gets the server address for uploading a file in a private message.
// Doc: https://dev.vk.com/method/docs.getMessagesUploadServer
type DocsGetMessagesUploadServerRequest struct {
	BaseRequest
}

// NewDocsGetMessagesUploadServerRequest creates a new request for docs.getMessagesUploadServer
func NewDocsGetMessagesUploadServerRequest(a *api.API, actor actor.Actor) *DocsGetMessagesUploadServerRequest {
	return &DocsGetMessagesUploadServerRequest{*NewMethodBaseRequest(a, actor, "docs.getMessagesUploadServer")}
}

// Exec executes the request and unmarshals the response into DocsGetMessagesUploadServerResponse
func (r *DocsGetMessagesUploadServerRequest) Exec(ctx context.Context) (response response.DocsGetMessagesUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetTypesRequest defines the request for docs.getTypes
//
// The method returns available document types for the user.
// Doc: https://dev.vk.com/method/docs.getTypes
type DocsGetTypesRequest struct {
	BaseRequest
}

// NewDocsGetTypesRequest creates a new request for docs.getTypes
func NewDocsGetTypesRequest(a *api.API, actor actor.Actor) *DocsGetTypesRequest {
	return &DocsGetTypesRequest{*NewMethodBaseRequest(a, actor, "docs.getTypes")}
}

// Exec executes the request and unmarshals the response into DocsGetTypesResponse
func (r *DocsGetTypesRequest) Exec(ctx context.Context) (response response.DocsGetTypesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetUploadServerRequest defines the request for docs.getUploadServer
//
// The method gets the server address for uploading a file.
// Doc: https://dev.vk.com/method/docs.getUploadServer
type DocsGetUploadServerRequest struct {
	BaseRequest
}

// NewDocsGetUploadServerRequest creates a new request for docs.getUploadServer
func NewDocsGetUploadServerRequest(a *api.API, actor actor.Actor) *DocsGetUploadServerRequest {
	return &DocsGetUploadServerRequest{*NewMethodBaseRequest(a, actor, "docs.getUploadServer")}
}

// Exec executes the request and unmarshals the response into DocsGetUploadServerResponse
func (r *DocsGetUploadServerRequest) Exec(ctx context.Context) (response response.DocsGetUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsGetWallUploadServerRequest defines the request for docs.getWallUploadServer
//
// The method gets the server address for uploading a document to the Sent folder for further posting on the wall or in a private message.
// Doc: https://dev.vk.com/method/docs.getWallUploadServer
type DocsGetWallUploadServerRequest struct {
	BaseRequest
}

// NewDocsGetWallUploadServerRequest creates a new request for docs.getWallUploadServer
func NewDocsGetWallUploadServerRequest(a *api.API, actor actor.Actor) *DocsGetWallUploadServerRequest {
	return &DocsGetWallUploadServerRequest{*NewMethodBaseRequest(a, actor, "docs.getWallUploadServer")}
}

// Exec executes the request and unmarshals the response into DocsGetWallUploadServerResponse
func (r *DocsGetWallUploadServerRequest) Exec(ctx context.Context) (response response.DocsGetWallUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsSaveRequest defines the request for docs.save
//
// The method saves the file after a successful upload to the server.
// Doc: https://dev.vk.com/method/docs.save
type DocsSaveRequest struct {
	BaseRequest
}

// NewDocsSaveRequest creates a new request for docs.save
func NewDocsSaveRequest(a *api.API, actor actor.Actor) *DocsSaveRequest {
	return &DocsSaveRequest{*NewMethodBaseRequest(a, actor, "docs.save")}
}

// Exec executes the request and unmarshals the response into DocsSaveResponse
func (r *DocsSaveRequest) Exec(ctx context.Context) (response response.DocsSaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DocsSearchRequest defines the request for docs.search
//
// The method returns search results for documents.
// Doc: https://dev.vk.com/method/docs.search
type DocsSearchRequest struct {
	BaseRequest
}

// NewDocsSearchRequest creates a new request for docs.search
func NewDocsSearchRequest(a *api.API, actor actor.Actor) *DocsSearchRequest {
	return &DocsSearchRequest{*NewMethodBaseRequest(a, actor, "docs.search")}
}

// Exec executes the request and unmarshals the response into DocsSearchResponse
func (r *DocsSearchRequest) Exec(ctx context.Context) (response response.DocsSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
