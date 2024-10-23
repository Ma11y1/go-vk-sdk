package api

import (
	"go-vk-sdk/logger"
	"go-vk-sdk/transport"
)

type API struct {
	Version        string           `json:"version"`
	MethodEndpoint string           `json:"method_endpoint"`
	AuthEndpoint   string           `json:"auth_endpoint"`
	Language       string           `json:"language"`
	Client         transport.Client `json:"client"`
}

func NewAPI() *API {
	api := &API{
		Version:        Version,
		MethodEndpoint: MethodEndpoint,
		AuthEndpoint:   AuthEndpoint, // Legacy OAuth or VKID
		Language:       LanguageDefault,
		Client:         transport.NewHTTPClient(),
	}

	return api
}

func (a *API) SetHostMethodEndpoint(host string) {
	if host == "" {
		logger.Log("API.SetHostMethodEndpoint", "empty host value")
		return
	}

	if host[len(host)-1] == '/' {
		a.MethodEndpoint = "https://" + host + "method"
	} else {
		a.MethodEndpoint = "https://" + host + "/method"
	}
}

func (a *API) SetHostAuthEndpoint(host string) {
	if host == "" {
		logger.Log("API.SetHostAuthEndpoint", "empty host value")
		return
	}

	a.MethodEndpoint = "https://" + host
	if a.MethodEndpoint[len(a.MethodEndpoint)-1] != '/' {
		a.MethodEndpoint += "/"
	}
}
