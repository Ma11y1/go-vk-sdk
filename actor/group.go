package actor

import (
	"encoding/json"
	"fmt"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/errors"
	"net/url"
	"strconv"
	"strings"
)

type Group struct {
	ID          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

type Groups struct {
	Groups    []Group `json:"groups"`
	ExpiresIn int     `json:"expires_in"`
}

func NewGroupTokensFromJSON(data []byte) (*Groups, error) {
	var e errors.AuthCodeFlowError

	err := json.Unmarshal(data, &e)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal error response: %w", err)
	}

	if e.Type != "" {
		return nil, &e
	}

	var groups Groups

	err = json.Unmarshal(data, &groups)
	if err != nil {
		return nil, err
	}

	return &groups, nil
}

func NewGroupTokensFromURL(u *url.URL) (*Groups, error) {
	v, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, fmt.Errorf("oauth: %w", err)
	}

	if errType := v.Get("error"); errType != "" {
		return nil, &errors.AuthCodeFlowError{
			Type:        errType,
			Reason:      v.Get("error_reason"),
			Description: v.Get("error_description"),
		}
	}

	t := &Groups{
		Groups: make([]Group, 0),
	}

	for key, vs := range v {
		if strings.HasPrefix(key, "access_token_") {
			groupID, _ := strconv.Atoi(key[13:])
			if len(vs) > 0 {
				accessToken := vs[0]
				t.Groups = append(t.Groups, Group{
					ID:          groupID,
					AccessToken: accessToken,
				})
			}
		}

		if key == "expires_in" {
			t.ExpiresIn, _ = strconv.Atoi(vs[0])
		}
	}

	return t, nil
}

func (u *Group) BuildCodeFlowAuthorizeURL(params *GroupAuthorizeURLParams) (*url.URL, error) {
	return NewGroupCodeFlowAuthorizeURL(params)
}

func (u *Group) BuildImplicitFlowAuthorizeURL(params *GroupAuthorizeURLParams) (*url.URL, error) {
	return NewGroupImplicitFlowAuthorizeURL(params)
}

type GroupAuthorizeURLParams struct {
	ClientID    int    `json:"client_id"`    // required
	RedirectURI string `json:"redirect_uri"` // required
	GroupIDs    []int  `json:"group_ids"`
	Display     string `json:"display"` // page, popup, mobile
	Scope       int    `json:"scope"`
	State       string `json:"state"`
}

func groupAuthorizeURLValues(params *GroupAuthorizeURLParams) *url.Values {
	q := url.Values{
		constants.ParameterNameClientID: {strconv.Itoa(params.ClientID)},
	}

	if params.RedirectURI == "" {
		q.Set(constants.ParameterNameRedirectURI, api.AuthRedirectURI)
	} else {
		q.Set(constants.ParameterNameRedirectURI, params.RedirectURI)
	}

	if params.Scope > 0 {
		q.Set(constants.ParameterNameScope, strconv.Itoa(params.Scope))
	}

	if params.Display != "" {
		q.Set(constants.ParameterNameDisplay, params.Display)
	}

	if params.State != "" {
		q.Set(constants.ParameterNameState, params.State)
	}

	var groupIDs string

	for i, id := range params.GroupIDs {
		if i != 0 {
			groupIDs += ","
		}

		groupIDs += strconv.Itoa(id)
	}

	q.Set(constants.ParameterNameGroupIDs, groupIDs)

	return &q
}

// NewGroupCodeFlowAuthorizeURL Assembling url to request an authorization code, which is used to obtain a token in code flow auth
// If scope is less than 0 then it is not added
func NewGroupCodeFlowAuthorizeURL(params *GroupAuthorizeURLParams) (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, err
	}

	q := groupAuthorizeURLValues(params)
	q.Set(constants.ParameterNameResponseType, constants.ParameterNameCode)

	u.RawQuery = groupAuthorizeURLValues(params).Encode()

	return u, nil
}

// NewGroupImplicitFlowAuthorizeURL Assembling url to request an authorization code, which is used to obtain a token in implicit flow auth
// If scope is less than 0 then it is not added
func NewGroupImplicitFlowAuthorizeURL(params *GroupAuthorizeURLParams) (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, err
	}

	q := groupAuthorizeURLValues(params)
	q.Set(constants.ParameterNameResponseType, constants.ParameterNameToken)

	u.RawQuery = q.Encode()

	return u, nil
}
