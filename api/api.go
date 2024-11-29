package api

import (
	internalError "go-vk-sdk/errors"
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

func (a *API) SetHostMethodEndpoint(host string) error {
	if host == "" {
		return internalError.Error("API.SetHostMethodEndpoint()", "empty host value "+host)
	}

	if host[len(host)-1] == '/' {
		a.MethodEndpoint = "https://" + host + "method"
	} else {
		a.MethodEndpoint = "https://" + host + "/method"
	}

	return nil
}

func (a *API) SetHostAuthEndpoint(host string) error {
	if host == "" {
		return internalError.Error("API.SetHostAuthEndpoint()", "empty host value "+host)
	}

	a.MethodEndpoint = "https://" + host
	if a.MethodEndpoint[len(a.MethodEndpoint)-1] != '/' {
		a.MethodEndpoint += "/"
	}

	return nil
}
