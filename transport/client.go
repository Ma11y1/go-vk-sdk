package transport

import (
	"context"
	"io"
	"net/http"
	"time"
)

type Client interface {
	Get(ctx context.Context, url string, header *http.Header) (*http.Response, error)
	GetDecodeJSON(ctx context.Context, url string, target interface{}, header *http.Header) (*http.Response, error)
	Post(ctx context.Context, url string, body io.Reader, header *http.Header) (*http.Response, error)
	PostDecodeJSON(ctx context.Context, url string, body io.Reader, target interface{}, header *http.Header) (*http.Response, error)
	SetUserAgent(string)
	SetAttemptTimeout(t time.Duration)
	Close() error
	IsClose() bool
}
