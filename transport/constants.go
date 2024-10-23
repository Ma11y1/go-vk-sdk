package transport

import "time"

const (
	UserAgent                 string        = "user-agent-go-vk-sdk"
	SocketTimeout             time.Duration = time.Millisecond * 5000
	MaxIdleConnections        int           = 100
	MaxIdleConnectionsPerHost int           = 100
	MaxConnectionsPerHost     int           = 100
	IdleConnTimeout           time.Duration = time.Millisecond * 90000
	NetDialerTimeout          time.Duration = time.Millisecond * 10000
	NetDialerKeepALive        time.Duration = time.Millisecond * 10000
	TSLHandshakeTimeout       time.Duration = time.Millisecond * 5000
	ExpectContinueTimeout     time.Duration = time.Millisecond * 1000
	ForceHTTP2                bool          = true
	RetryAttempts             int           = 3
	RetryAttemptTimeout       time.Duration = time.Millisecond * 500
)
