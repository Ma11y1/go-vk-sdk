package actor

import (
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	internalError "go-vk-sdk/errors"
	"net/url"
	"strconv"
)

// User

// UserImplicitFlowAuthorizeURL Contains data for generating URL
type UserImplicitFlowAuthorizeURL struct {
	ClientID    int    `json:"client_id"`         // required
	RedirectURI string `json:"redirect_uri"`      // required
	Display     string `json:"display,omitempty"` // page, popup, mobile
	Scope       int    `json:"scope,omitempty"`
	State       string `json:"state,omitempty"`
	Version     string `json:"version"`
}

// Build Generates URL
func (a *UserImplicitFlowAuthorizeURL) Build() (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, internalError.Error("Actor.ImplicitFlow.UserImplicitFlowAuthorizeURL,Build()", err.Error())
	}

	q := url.Values{
		constants.ParameterNameResponseType: {constants.ParameterNameToken},
		constants.ParameterNameClientID:     {strconv.Itoa(a.ClientID)},
	}

	if a.RedirectURI == "" {
		q.Set(constants.ParameterNameRedirectURI, AuthRedirectURI)
	} else {
		q.Set(constants.ParameterNameRedirectURI, a.RedirectURI)
	}

	if a.Scope > 0 {
		q.Set(constants.ParameterNameScope, strconv.Itoa(a.Scope))
	}

	if a.Display != "" {
		q.Set(constants.ParameterNameDisplay, a.Display)
	}

	if a.State != "" {
		q.Set(constants.ParameterNameState, a.State)
	}

	if a.Version != "" {
		q.Set(constants.ParameterNameVersion, a.Version)
	} else {
		q.Set(constants.ParameterNameVersion, api.Version)
	}

	u.RawQuery = q.Encode()

	return u, nil
}

// Group

// GroupsImplicitFlowAuthorizeURL Contains data for generating URL
type GroupsImplicitFlowAuthorizeURL struct {
	ClientID    int    `json:"client_id"`    // required
	RedirectURI string `json:"redirect_uri"` // required
	GroupIDs    []int  `json:"group_ids"`
	Display     string `json:"display"` // page, popup, mobile
	Scope       int    `json:"scope"`
	State       string `json:"state"`
	Version     string `json:"version"`
}

// Build Generates URL
func (groups *GroupsImplicitFlowAuthorizeURL) Build() (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, internalError.Error("Actor.ImplicitFlow.GroupsImplicitFlowAuthorizeURL,Build()", err.Error())
	}

	q := url.Values{
		constants.ParameterNameResponseType: {constants.ParameterNameToken},
		constants.ParameterNameClientID:     {strconv.Itoa(groups.ClientID)},
	}

	if groups.RedirectURI == "" {
		q.Set(constants.ParameterNameRedirectURI, AuthRedirectURI)
	} else {
		q.Set(constants.ParameterNameRedirectURI, groups.RedirectURI)
	}

	if groups.Scope > 0 {
		q.Set(constants.ParameterNameScope, strconv.Itoa(groups.Scope))
	}

	if groups.Display != "" {
		q.Set(constants.ParameterNameDisplay, groups.Display)
	}

	if groups.State != "" {
		q.Set(constants.ParameterNameState, groups.State)
	}

	var groupIDs string

	for i, id := range groups.GroupIDs {
		if i != 0 {
			groupIDs += ","
		}

		groupIDs += strconv.Itoa(id)
	}

	q.Set(constants.ParameterNameGroupIDs, groupIDs)

	if groups.Version != "" {
		q.Set(constants.ParameterNameVersion, groups.Version)
	} else {
		q.Set(constants.ParameterNameVersion, api.Version)
	}

	u.RawQuery = q.Encode()

	return u, nil
}
