package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"strconv"
)

// AuthUserDirectRequest defines the request for direct auth
type AuthUserDirectRequest struct {
	BaseRequest
}

func NewAuthUserDirectAuthRequest(a *api.API) *AuthUserDirectRequest {
	r := &AuthUserDirectRequest{*NewAuthBaseRequest(a, api.AuthDirectMethod)}
	r.parameters.Set(constants.ParameterNameGrandType, constants.ParameterNamePassword)
	return r
}

func (r *AuthUserDirectRequest) Exec(ctx context.Context) (response response.AuthDirectResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AuthUserDirectRequest) ClientID(id int) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

func (r *AuthUserDirectRequest) ClientSecret(secret string) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNameClientSecret, secret)
	return r
}

func (r *AuthUserDirectRequest) Username(name string) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNameUsername, name)
	return r
}

func (r *AuthUserDirectRequest) Password(password string) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNamePassword, password)
	return r
}

func (r *AuthUserDirectRequest) Scope(scope int) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNameScope, strconv.Itoa(scope))
	return r
}

func (r *AuthUserDirectRequest) TestRedirectURI(v bool) *AuthUserDirectRequest {
	if v {
		r.parameters.Set(constants.ParameterNameTestRedirectURI, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameTestRedirectURI)
	}
	return r
}

func (r *AuthUserDirectRequest) TwoFactorSupported(v bool) *AuthUserDirectRequest {
	if v {
		r.parameters.Set(constants.ParameterName2faSupported, "1")
	} else {
		r.parameters.Remove(constants.ParameterName2faSupported)
	}
	return r
}

func (r *AuthUserDirectRequest) ForceSMS(v bool) *AuthUserDirectRequest {
	if v {
		r.parameters.Set(constants.ParameterNameForceSMS, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameForceSMS)
	}
	return r
}

func (r *AuthUserDirectRequest) Code(code string) *AuthUserDirectRequest {
	r.parameters.Set(constants.ParameterNameCode, code)
	return r
}

// AuthUserCodeFlowRequest defines the request auth code flow
type AuthUserCodeFlowRequest struct {
	BaseRequest
}

func NewAuthUserCodeFlowRequest(a *api.API) *AuthUserCodeFlowRequest {
	return &AuthUserCodeFlowRequest{*NewAuthBaseRequest(a, api.AuthCodeFlowTokenMethod)}
}

func (r *AuthUserCodeFlowRequest) Exec(ctx context.Context) (response response.AuthUserCodeFlowResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AuthUserCodeFlowRequest) ClientID(id int) *AuthUserCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

func (r *AuthUserCodeFlowRequest) ClientSecret(secret string) *AuthUserCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameClientSecret, secret)
	return r
}

func (r *AuthUserCodeFlowRequest) RedirectURI(uri string) *AuthUserCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameRedirectURI, uri)
	return r
}

func (r *AuthUserCodeFlowRequest) Code(code string) *AuthUserCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameCode, code)
	return r
}

// AuthGroupCodeFlowRequest defines the request auth code flow
//
//	Used to exchange an authorization code for an access and refresh token
type AuthGroupCodeFlowRequest struct {
	BaseRequest
}

func NewAuthGroupCodeFlowRequest(a *api.API) *AuthGroupCodeFlowRequest {
	return &AuthGroupCodeFlowRequest{*NewAuthBaseRequest(a, api.AuthCodeFlowTokenMethod)}
}

func (r *AuthGroupCodeFlowRequest) Exec(ctx context.Context) (response response.AuthGroupCodeFlowResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AuthGroupCodeFlowRequest) ClientID(id int) *AuthGroupCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

func (r *AuthGroupCodeFlowRequest) ClientSecret(secret string) *AuthGroupCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameClientSecret, secret)
	return r
}

func (r *AuthGroupCodeFlowRequest) RedirectURI(uri string) *AuthGroupCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameRedirectURI, uri)
	return r
}

func (r *AuthGroupCodeFlowRequest) Code(code string) *AuthGroupCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameCode, code)
	return r
}

// AuthRestoreRequest defines the request for auth.restore
//
//	Allows you to restore access to your account using a code received via SMS.
//	This method is only available to applications that have access to Direct Authorization.
//	Doc: https://dev.vk.com/ru/method/auth.restore
type AuthRestoreRequest struct {
	BaseRequest
}

// NewAuthRestoreRequest creates a new request for auth.restore
func NewAuthRestoreRequest(a *api.API, actor actor.Actor) *AuthRestoreRequest {
	return &AuthRestoreRequest{*NewMethodBaseRequest(a, actor, "auth.restore")}
}

// Exec executes the request and unmarshals the response into AccountBanResponse
func (r *AuthRestoreRequest) Exec(ctx context.Context) (response response.AuthRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AuthVKIDCodeFlowRequest
//
//	Used to exchange an authorization code for an access and refresh token
//	Doc: https://id.vk.com/about/business/go/docs/ru/vkid/latest/vk-id/connection/start-integration/auth-flow-web
type AuthVKIDCodeFlowRequest struct {
	BaseRequest
}

func NewAuthVKIDCodeFlowRequest(a *api.API) *AuthVKIDCodeFlowRequest {
	r := &AuthVKIDCodeFlowRequest{*NewAuthBaseRequest(a, api.AuthVKIDCodeFlowTokenMethod)}
	r.parameters.Set(constants.ParameterNameGrandType, constants.ParameterNameAuthorizationCode)
	return r
}

func (r *AuthVKIDCodeFlowRequest) Exec(ctx context.Context) (response response.AuthUserVKIDCodeFlowResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AuthVKIDCodeFlowRequest) Code(code string) *AuthVKIDCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameCode, code)
	return r
}

func (r *AuthVKIDCodeFlowRequest) DeviceID(id string) *AuthVKIDCodeFlowRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, id)
	return r
}
