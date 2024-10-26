package request

import (
	"bytes"
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/transport"
	"io"
	"net/http"
	"strings"
	"sync"
)

type Request interface {
	Get(ctx context.Context) (*http.Response, error)
	GetUnmarshal(ctx context.Context, target interface{}) error
	Post(ctx context.Context) (*http.Response, error)
	PostData(ctx context.Context, data io.Reader) (*http.Response, error)
	PostUnmarshal(ctx context.Context, target interface{}) error
	GetHeader() *http.Header
	GetParameters() Parameters
	SetParameters(Parameters)
}

// BaseRequest VK api methods (except for methods from the secure and ads sections) with a user access key can be accessed no more than 3 times per second.
// The Community Access key is limited to 20 requests per second. If your application logic involves calling several methods in a row, it makes sense to pay attention to the execute method.
// It allows you to make up to 25 calls to different methods within one request.
type BaseRequest struct {
	mtx        sync.RWMutex
	url        string
	header     *http.Header
	parameters Parameters
	transport  transport.Client
}

func NewAuthBaseRequest(api *api.API, method string) *BaseRequest {
	r := &BaseRequest{
		url:       api.AuthEndpoint + method,
		transport: api.Client,
		header: &http.Header{
			"Content-EventType": []string{constants.ContentTypeFormURLEncoded},
			"Accept-Encoding":   []string{constants.AcceptEncodingGzip},
		},
		parameters: NewBaseRequestParameters(),
	}

	r.Version(api.Version)
	r.Language(api.Language)

	return r
}

func NewMethodBaseRequest(api *api.API, actor actor.Actor, method string) *BaseRequest {
	r := &BaseRequest{
		url:       api.MethodEndpoint + method,
		transport: api.Client,
		header: &http.Header{
			"Content-EventType": []string{constants.ContentTypeFormURLEncoded},
			"Accept-Encoding":   []string{constants.AcceptEncodingGzip},
		},
		parameters: NewBaseRequestParameters(),
	}

	r.AccessToken(actor.GetAccessToken())
	r.Version(api.Version)
	r.Language(api.Language)

	return r
}

func NewUploadBaseRequest(api *api.API, actor actor.Actor) *BaseRequest {
	r := &BaseRequest{
		url:        "",
		transport:  api.Client,
		header:     &http.Header{},
		parameters: NewBaseRequestParameters(),
	}

	r.AccessToken(actor.GetAccessToken())
	r.Version(api.Version)
	r.Language(api.Language)

	return r
}

func (r *BaseRequest) Get(ctx context.Context) (*http.Response, error) {
	r.mtx.RLock()
	u := r.url

	if strings.Contains(u, "?") {
		u += "&" + r.parameters.BuildURLValuesEncode()
	} else {
		u += "?" + r.parameters.BuildURLValuesEncode()
	}

	r.mtx.RUnlock()

	resp, err := r.transport.Get(ctx, u, r.header)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *BaseRequest) GetUnmarshal(ctx context.Context, target interface{}) error {
	r.mtx.RLock()

	u := r.url

	if strings.Contains(u, "?") {
		u += "&" + r.parameters.BuildURLValuesEncode()
	} else {
		u += "?" + r.parameters.BuildURLValuesEncode()
	}

	httpResponse, err := r.transport.GetDecodeJSON(ctx, u, target, r.header)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close() // default http client api closes the response body itself
	return nil
}

func (r *BaseRequest) Post(ctx context.Context) (*http.Response, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.transport.Post(ctx, r.url, bytes.NewBufferString(r.parameters.BuildURLValuesEncode()), r.header)
}

func (r *BaseRequest) PostData(ctx context.Context, data io.Reader) (*http.Response, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.transport.Post(ctx, r.url, data, r.header)
}

func (r *BaseRequest) PostUnmarshal(ctx context.Context, target interface{}) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	httpResponse, err := r.transport.PostDecodeJSON(
		ctx,
		r.url,
		bytes.NewBufferString(r.parameters.BuildURLValuesEncode()),
		target,
		r.header,
	)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close() // default http client api closes the response body itself
	return nil
}

func (r *BaseRequest) GetHeader() *http.Header {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.header
}

func (r *BaseRequest) GetParameters() Parameters {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.parameters
}

func (r *BaseRequest) SetParameters(params Parameters) {
	if params != nil {
		r.mtx.Lock()
		r.parameters = params
		r.mtx.Unlock()
	}
}

func (r *BaseRequest) ResetParameters() *BaseRequest {
	r.mtx.Lock()
	version := r.parameters.Get(constants.ParameterNameVersion)
	lang := r.parameters.Get(constants.ParameterNameLang)

	r.parameters.Clear()

	_ = r.parameters.SetIfNotEmpty(constants.ParameterNameVersion, version)
	_ = r.parameters.SetIfNotEmpty(constants.ParameterNameLang, lang)
	r.mtx.Unlock()

	return r
}

func (r *BaseRequest) GetURL() string {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.url
}

func (r *BaseRequest) SetURL(url string) {
	if url != "" {
		r.mtx.Lock()
		r.url = url
		r.mtx.Unlock()
	}
}

func (r *BaseRequest) GetTransport() transport.Client {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	return r.transport
}

func (r *BaseRequest) SetTransport(t transport.Client) {
	if t != nil {
		r.mtx.Lock()
		r.transport = t
		r.mtx.Unlock()
	}
}

func (r *BaseRequest) AccessToken(token string) *BaseRequest {
	if token != "" {
		r.mtx.Lock()
		r.header.Set("Authorization", "Bearer "+token)
		r.mtx.Unlock()
	}
	return r
}

func (r *BaseRequest) SetContentType(t string) *BaseRequest {
	if t != "" {
		r.mtx.Lock()
		r.header.Set("Content-Type", t)
		r.mtx.Unlock()
	}
	return r
}

func (r *BaseRequest) SetAcceptEncoding(a string) *BaseRequest {
	if a != "" {
		r.mtx.Lock()
		r.header.Set("Accept-Encoding", a)
		r.mtx.Unlock()
	}
	return r
}

func (r *BaseRequest) Version(v string) *BaseRequest {
	r.mtx.Lock()
	_ = r.parameters.SetIfNotEmpty(constants.ParameterNameVersion, v)
	r.mtx.Unlock()
	return r
}

func (r *BaseRequest) Language(lang string) *BaseRequest {
	r.mtx.Lock()
	_ = r.parameters.SetIfNotEmpty(constants.ParameterNameLang, lang)
	r.mtx.Unlock()
	return r
}

// Captcha Doc: https://dev.vk.com/ru/api/captcha-error
func (r *BaseRequest) Captcha(sid string, key string) *BaseRequest {
	if sid != "" && key != "" {
		r.mtx.Lock()
		_ = r.parameters.Set(constants.ParameterNameCaptchaSID, sid)
		_ = r.parameters.Set(constants.ParameterNameCaptchaKey, key)
		r.mtx.Unlock()
	}
	return r
}

func (r *BaseRequest) ClearCaptcha() *BaseRequest {
	r.mtx.Lock()
	r.parameters.Remove(constants.ParameterNameCaptchaSID)
	r.parameters.Remove(constants.ParameterNameCaptchaKey)
	r.mtx.Unlock()
	return r
}

// Confirm Doc: https://dev.vk.com/ru/api/confirmation-required-error
func (r *BaseRequest) Confirm(confirm string) *BaseRequest {
	r.mtx.Lock()
	_ = r.parameters.SetIfNotEmpty(constants.ParameterNameConfirm, confirm)
	r.mtx.Unlock()
	return r
}

func (r *BaseRequest) ClearConfirm() *BaseRequest {
	r.mtx.Lock()
	r.parameters.Remove(constants.ParameterNameConfirm)
	r.mtx.Unlock()
	return r
}
