package transport

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	internalErrors "go-vk-sdk/errors"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"
)

type HTTPClientParameters struct {
	Transport       http.RoundTripper `json:"transport,omitempty"`
	SocketTimeout   time.Duration     `json:"socketTimeout,omitempty"`
	UserAgent       string            `json:"user_agent,omitempty"`
	AttemptsRequest int               `json:"attempts_request,omitempty"`
	AttemptTimeout  time.Duration     `json:"attempt_timeout,omitempty"`
}

type HTTPClient struct {
	http.Client
	mtx             sync.RWMutex
	userAgent       string
	attemptsRequest int
	attemptTimeout  time.Duration
	isClose         bool
}

func NewHTTPClient() *HTTPClient {
	jar, _ := cookiejar.New(nil)

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   NetDialerTimeout,
			KeepAlive: NetDialerKeepALive,
		}).DialContext,
		MaxIdleConns:          MaxIdleConnections,
		MaxConnsPerHost:       MaxConnectionsPerHost,
		MaxIdleConnsPerHost:   MaxIdleConnectionsPerHost,
		IdleConnTimeout:       IdleConnTimeout,
		TLSHandshakeTimeout:   TSLHandshakeTimeout,
		ExpectContinueTimeout: ExpectContinueTimeout,
		ForceAttemptHTTP2:     ForceHTTP2,
	}

	return &HTTPClient{
		mtx: sync.RWMutex{},
		Client: http.Client{
			Jar:       jar,
			Timeout:   SocketTimeout,
			Transport: transport,
		},
		userAgent:       UserAgent,
		attemptsRequest: RetryAttempts,
		attemptTimeout:  RetryAttemptTimeout,
	}
}

func NewHTTPClientParameters(p *HTTPClientParameters) *HTTPClient {
	client := NewHTTPClient()

	if p.Transport != nil {
		client.Transport = p.Transport
	}

	if p.SocketTimeout > 0 {
		client.Timeout = p.SocketTimeout
	}

	if p.UserAgent != "" {
		client.userAgent = p.UserAgent
	}

	if p.AttemptsRequest > 0 {
		client.attemptsRequest = p.AttemptsRequest
	}

	if p.AttemptTimeout > 0 {
		client.attemptTimeout = p.AttemptTimeout
	}

	return client
}

func (c *HTTPClient) request(ctx context.Context, req *http.Request) (*http.Response, error) {
	if c.isClose {
		return nil, internalErrors.ErrorLog("Transport.HttpClient.request()", "client is already closed")
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept-Encoding", "gzip")

	var err error
	var resp *http.Response

	for i := 0; i < c.attemptsRequest; i++ {
		if ctx.Err() != nil {
			return nil, internalErrors.ErrorLog("Transport.HttpClient.request()", "close context "+ctx.Err().Error())
		}

		resp, err = c.Do(req)
		if err != nil {
			time.Sleep(c.attemptTimeout)
			continue
		}

		if resp.StatusCode >= 500 {
			resp.Body.Close()
			err = internalErrors.ErrorLog("Transport.HttpClient.request()", fmt.Sprintf("unexpected status code %d\n", resp.StatusCode))
			time.Sleep(c.attemptTimeout)
			continue
		}

		if resp.Header.Get("Content-Encoding") == "gzip" {
			gzipReader, err := gzip.NewReader(resp.Body)
			if err != nil {
				resp.Body.Close()
				return nil, internalErrors.ErrorLog("Transport.HttpClient.request()", "error decompressing data "+err.Error())
			}

			// Original body objects is replaced with a custom ReadCloser,
			// which stores the reader of the original body objects and a Gzip reader,
			// so that when closing all readers are closed correctly
			resp.Body = &ReadCloser{Origin: resp.Body, Encode: gzipReader}
		}

		return resp, nil
	}

	return nil, err
}

func (c *HTTPClient) Get(ctx context.Context, url string, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, internalErrors.ErrorLog("Transport.HttpClient.Get()", "error create HTTP request with context "+err.Error())
	}

	if header != nil {
		req.Header = *header
	}

	return c.request(ctx, req)
}

func (c *HTTPClient) GetDecodeJSON(ctx context.Context, url string, target interface{}, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, internalErrors.ErrorLog("Transport.HttpClient.GetDecodeJSON()", "error create HTTP request with context "+err.Error())
	}

	if header != nil {
		req.Header = *header
	}

	res, err := c.request(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return res, internalErrors.ErrorLog("Transport.HttpClient.GetDecodeJSON()", "error decoding JSON "+err.Error())
	}

	return res, err
}

func (c *HTTPClient) Post(ctx context.Context, url string, body io.Reader, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, internalErrors.ErrorLog("Transport.HttpClient.Post()", "error create HTTP request with context "+err.Error())
	}

	if header != nil {
		req.Header = *header
	}

	return c.request(ctx, req)
}

func (c *HTTPClient) PostDecodeJSON(ctx context.Context, url string, body io.Reader, target interface{}, header *http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, internalErrors.ErrorLog("Transport.HttpClient.PostDecodeJSON()", "error create HTTP request with context "+err.Error())
	}

	if header != nil {
		req.Header = *header
	}

	res, err := c.request(ctx, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return res, internalErrors.ErrorLog("Transport.HttpClient.PostDecodeJSON()", "error decoding JSON "+err.Error())
	}

	return res, err
}

func (c *HTTPClient) SetUserAgent(ua string) {
	if ua != "" {
		c.mtx.Lock()
		c.userAgent = ua
		c.mtx.Unlock()
	}
}

func (c *HTTPClient) SetAttemptTimeout(t time.Duration) {
	if t > 0 {
		c.mtx.Lock()
		c.attemptTimeout = t
		c.mtx.Unlock()
	}
}

func (c *HTTPClient) Close() error {
	if c.isClose {
		return internalErrors.ErrorLog("Transport.HttpClient.PostDecodeJSON()", "client is already closed")
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if transport, ok := c.Transport.(*http.Transport); ok {
		transport.CloseIdleConnections()
		transport.DisableKeepAlives = true
	}

	c.isClose = true

	return nil
}

func (c *HTTPClient) IsClose() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.isClose
}
