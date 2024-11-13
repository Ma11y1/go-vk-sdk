package transport

import (
	"context"
	"errors"
	"fmt"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"net/http"
)

type CallbackServer interface {
	Run() error
	Stop(context.Context) error
	SetAddress(string) error
	SetHandler(http.Handler) error
	IsRunning() bool
}

type BaseCallbackServer struct {
	server    *http.Server
	handler   http.Handler
	address   string
	isRunning bool
}

func NewBaseCallbackServer(address string) *BaseCallbackServer {
	return &BaseCallbackServer{
		server: &http.Server{
			Addr: address,
		},
		handler:   nil,
		address:   address,
		isRunning: false,
	}
}

func NewBaseCallbackServerWithHandler(address string, handler http.Handler) *BaseCallbackServer {
	s := NewBaseCallbackServer(address)
	s.handler = handler
	s.server.Handler = handler
	return s
}

func (s *BaseCallbackServer) Run() error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.Run()", "Server already running")
	}

	s.isRunning = true
	go func() {
		logger.Log("BaseCallbackServer.Run()", fmt.Sprintf("Server is running at address: %s", s.address))

		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log("BaseCallbackServer.Run()", "Error listen and serve server: "+err.Error())
			s.isRunning = false
			return
		}
	}()

	return nil
}

func (s *BaseCallbackServer) Stop(ctx context.Context) error {
	if !s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.Stop()", "Server is not running")
	}

	s.isRunning = false

	err := s.server.Shutdown(ctx)
	if err != nil {
		return internalErrors.ErrorLog("BaseCallbackServer.Stop()", "Server stop error: "+err.Error())
	}

	logger.Log("BaseCallbackServer.Stop()", "Server was stopped at address: "+s.address)

	s.server = &http.Server{
		Addr:    s.address,
		Handler: s.handler,
	}

	return nil
}

func (s *BaseCallbackServer) SetAddress(address string) error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "Cannot change address while server is running")
	}

	s.address = address
	s.server.Addr = address

	return nil
}

func (s *BaseCallbackServer) SetHandler(handler http.Handler) error {
	if s.isRunning {
		return internalErrors.ErrorLog("BaseCallbackServer.SetHandler()", "Cannot change handler while server is running")
	}

	s.handler = handler
	s.server.Handler = handler

	return nil
}

func (s *BaseCallbackServer) IsRunning() bool {
	return s.isRunning
}
