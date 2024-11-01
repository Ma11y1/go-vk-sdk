package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/leadForms
// Suitable for registering clients, pre-registration, subscribing to newsletters, requesting information, connecting services, placing orders and much more.
// You create application forms, and users leave their contact information. All you have to do is complete the application by contacting them in a convenient way.

// LeadFormsCreateRequest defines the request for leadForms.create
// Doc: https://dev.vk.com/ru/method/leadForms.create
type LeadFormsCreateRequest struct {
	BaseRequest
}

// NewLeadFormsCreateRequest creates a new request for leadForms.create
func NewLeadFormsCreateRequest(a *api.API, actor actor.Actor) *LeadFormsCreateRequest {
	return &LeadFormsCreateRequest{*NewMethodBaseRequest(a, actor, "leadForms.create")}
}

// Exec executes the request and unmarshals the response into LeadFormsCreateResponse
func (r *LeadFormsCreateRequest) Exec(ctx context.Context) (response response.LeadFormsCreateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsDeleteRequest defines the request for leadForms.delete
// Doc: https://dev.vk.com/ru/method/leadForms.delete
type LeadFormsDeleteRequest struct {
	BaseRequest
}

// NewLeadFormsDeleteRequest creates a new request for leadForms.delete
func NewLeadFormsDeleteRequest(a *api.API, actor actor.Actor) *LeadFormsDeleteRequest {
	return &LeadFormsDeleteRequest{*NewMethodBaseRequest(a, actor, "leadForms.delete")}
}

// Exec executes the request and unmarshals the response into LeadFormsDeleteResponse
func (r *LeadFormsDeleteRequest) Exec(ctx context.Context) (response response.LeadFormsDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsGetRequest defines the request for leadForms.get
// Doc: https://dev.vk.com/ru/method/leadForms.get
type LeadFormsGetRequest struct {
	BaseRequest
}

// NewLeadFormsGetRequest creates a new request for leadForms.get
func NewLeadFormsGetRequest(a *api.API, actor actor.Actor) *LeadFormsGetRequest {
	return &LeadFormsGetRequest{*NewMethodBaseRequest(a, actor, "leadForms.get")}
}

// Exec executes the request and unmarshals the response into LeadFormsGetResponse
func (r *LeadFormsGetRequest) Exec(ctx context.Context) (response response.LeadFormsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsGetLeadsRequest defines the request for leadForms.getLeads
// Doc: https://dev.vk.com/ru/method/leadForms.getLeads
type LeadFormsGetLeadsRequest struct {
	BaseRequest
}

// NewLeadFormsGetLeadsRequest creates a new request for leadForms.getLeads
func NewLeadFormsGetLeadsRequest(a *api.API, actor actor.Actor) *LeadFormsGetLeadsRequest {
	return &LeadFormsGetLeadsRequest{*NewMethodBaseRequest(a, actor, "leadForms.getLeads")}
}

// Exec executes the request and unmarshals the response into LeadFormsGetLeadsResponse
func (r *LeadFormsGetLeadsRequest) Exec(ctx context.Context) (response response.LeadFormsGetLeadsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsGetUploadURLRequest defines the request for leadForms.getUploadURL
// Doc: https://dev.vk.com/ru/method/leadForms.getUploadURL
type LeadFormsGetUploadURLRequest struct {
	BaseRequest
}

// NewLeadFormsGetUploadURLRequest creates a new request for leadForms.getUploadURL
func NewLeadFormsGetUploadURLRequest(a *api.API, actor actor.Actor) *LeadFormsGetUploadURLRequest {
	return &LeadFormsGetUploadURLRequest{*NewMethodBaseRequest(a, actor, "leadForms.getUploadURL")}
}

// Exec executes the request and unmarshals the response into LeadFormsGetUploadURLResponse
func (r *LeadFormsGetUploadURLRequest) Exec(ctx context.Context) (response response.LeadFormsGetUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsListRequest defines the request for leadForms.list
// Doc: https://dev.vk.com/ru/method/leadForms.list
type LeadFormsListRequest struct {
	BaseRequest
}

// NewLeadFormsListRequest creates a new request for leadForms.list
func NewLeadFormsListRequest(a *api.API, actor actor.Actor) *LeadFormsListRequest {
	return &LeadFormsListRequest{*NewMethodBaseRequest(a, actor, "leadForms.list")}
}

// Exec executes the request and unmarshals the response into LeadFormsListResponse
func (r *LeadFormsListRequest) Exec(ctx context.Context) (response response.LeadFormsListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// LeadFormsUpdateRequest defines the request for leadForms.update
// Doc: https://dev.vk.com/ru/method/leadForms.update
type LeadFormsUpdateRequest struct {
	BaseRequest
}

// NewLeadFormsUpdateRequest creates a new request for leadForms.update
func NewLeadFormsUpdateRequest(a *api.API, actor actor.Actor) *LeadFormsUpdateRequest {
	return &LeadFormsUpdateRequest{*NewMethodBaseRequest(a, actor, "leadForms.update")}
}

// Exec executes the request and unmarshals the response into LeadFormsUpdateResponse
func (r *LeadFormsUpdateRequest) Exec(ctx context.Context) (response response.LeadFormsUpdateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
