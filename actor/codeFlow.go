package actor

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	internalError "go-vk-sdk/errors"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Global

type CodeFlowAuthorizeCode struct {
	Code string `json:"code"`
}

// NewCodeFlowAuthorizeCodeRawURL Creates an object from a URL containing an authorization code, containing data, or returning an error
func NewCodeFlowAuthorizeCodeRawURL(rawURL string) (*CodeFlowAuthorizeCode, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewCodeFlowAuthorizeCodeRawURL()", err.Error())
	}

	return NewCodeFlowAuthorizeCodeURL(u)
}

// NewCodeFlowAuthorizeCodeURL Creates an object from a URL containing an authorization code, containing data, or returning an error
func NewCodeFlowAuthorizeCodeURL(u *url.URL) (*CodeFlowAuthorizeCode, error) {
	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewCodeFlowAuthorizeCodeURL()", err.Error())
	}

	if errType := q.Get("error"); errType != "" {
		return nil, &internalError.AuthCodeFlowError{
			Type:        errType,
			Reason:      q.Get("error_reason"),
			Description: q.Get("error_description"),
		}
	}

	authCode := &CodeFlowAuthorizeCode{
		Code: q.Get("code"),
	}

	return authCode, nil
}

// User

// UserCodeFlowAuthorizeURL Contains data for generating a URL for obtaining an authorization code
type UserCodeFlowAuthorizeURL struct {
	ClientID    int    `json:"client_id"`         // required
	RedirectURI string `json:"redirect_uri"`      // required
	Display     string `json:"display,omitempty"` // page, popup, mobile
	Scope       int    `json:"scope,omitempty"`
	State       string `json:"state,omitempty"`
	Version     string `json:"version"`
}

// Build Generates URL to get authorization code
func (a *UserCodeFlowAuthorizeURL) Build() (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.UserCodeFlowAuthorizeURL.Build()", err.Error())
	}

	q := url.Values{
		constants.ParameterNameResponseType: {constants.ParameterNameCode},
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

// NewUserCodeFlowRawURL Creating a user from the url parameters received after successfully completing the authorization code and receiving an access token
func NewUserCodeFlowRawURL(rawURL string) (*User, error) {
	user := &User{}
	err := user.InitFromCodeFlowRawURL(rawURL)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewUserCodeFlowRawURL()", err.Error())
	}
	return user, nil
}

// NewUserCodeFlowURL Creating a user from the url parameters received after successfully completing the authorization code and receiving an access token
func NewUserCodeFlowURL(u *url.URL) (*User, error) {
	user := &User{}
	err := user.InitFromCodeFlowURL(u)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewUserCodeFlowURL()", err.Error())
	}
	return user, nil
}

// NewUserCodeFlowJSON Creating a new user from JSON data received after successfully exchanging the authorization code for an access token
func NewUserCodeFlowJSON(data []byte) (*User, error) {
	user := User{}
	err := user.InitFromCodeFlowJSON(data)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewUserCodeFlowJSON()", err.Error())
	}
	return &user, nil
}

// InitFromCodeFlowRawURL Initialization from the url parameters received after successfully completing the authorization code and receiving an access token
func (user *User) InitFromCodeFlowRawURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.User.InitFromCodeFlowRawURL()", err.Error())
	}

	return user.InitFromCodeFlowURL(u)
}

// InitFromCodeFlowURL Initialization from the url parameters received after successfully completing the authorization code and receiving an access token
func (user *User) InitFromCodeFlowURL(u *url.URL) error {
	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.User.InitFromCodeFlowURL()", err.Error())
	}

	if errType := q.Get("errors"); errType != "" {
		return &internalError.AuthCodeFlowError{
			Type:        errType,
			Reason:      q.Get("error_reason"),
			Description: q.Get("error_description"),
		}
	}

	user.ID, _ = strconv.Atoi(q.Get("user_id"))
	user.AccessToken = q.Get("access_token")
	user.ExpiresIn, _ = strconv.Atoi(q.Get("expires_in"))
	user.State = q.Get("state")
	user.Email = q.Get("email")

	return nil
}

func (user *User) InitFromCodeFlowJSON(data []byte) error {
	var authError internalError.AuthCodeFlowError
	err := json.Unmarshal(data, &authError)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.User.InitFromCodeFlowJSON()", err.Error())
	}

	if authError.Type != "" {
		return &authError
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.User.InitFromCodeFlowJSON()", err.Error())
	}

	return nil
}

// Group

type GroupsCodeFlowAuthorizeURL struct {
	ClientID    int    `json:"client_id"`    // required
	RedirectURI string `json:"redirect_uri"` // required
	GroupIDs    []int  `json:"group_ids"`
	Display     string `json:"display"` // page, popup, mobile
	Scope       int    `json:"scope"`
	State       string `json:"state"`
	Version     string `json:"version"`
}

// Build Generating URL to get authorization code
func (a *GroupsCodeFlowAuthorizeURL) Build() (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.GroupsCodeFlowAuthorizeURL.Build()", err.Error())
	}

	q := url.Values{
		constants.ParameterNameResponseType: {constants.ParameterNameCode},
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

	var groupIDs string

	for i, id := range a.GroupIDs {
		if i != 0 {
			groupIDs += ","
		}

		groupIDs += strconv.Itoa(id)
	}

	q.Set(constants.ParameterNameGroupIDs, groupIDs)

	if a.Version != "" {
		q.Set(constants.ParameterNameVersion, a.Version)
	} else {
		q.Set(constants.ParameterNameVersion, api.Version)
	}

	u.RawQuery = q.Encode()

	return u, nil
}

// NewGroupsCodeFlowRawURL Creating a groups from the url parameters received after successfully completing the authorization code and receiving an access token
func NewGroupsCodeFlowRawURL(rawURL string) (*Groups, error) {
	groups := &Groups{}
	err := groups.InitFromCodeFlowRawURL(rawURL)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewGroupsCodeFlowRawURL()", err.Error())
	}
	return groups, nil
}

// NewGroupsCodeFlowURL Creating a groups from the url parameters received after successfully completing the authorization code and receiving an access token
func NewGroupsCodeFlowURL(u *url.URL) (*Groups, error) {
	groups := &Groups{}
	err := groups.InitFromCodeFlowURL(u)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewGroupsCodeFlowURL()", err.Error())
	}
	return groups, nil
}

// NewGroupsCodeFlowJSON Creating a new groups from JSON data received after successfully exchanging the authorization code for an access token
func NewGroupsCodeFlowJSON(data []byte) (*Groups, error) {
	groups := &Groups{}
	err := groups.InitFromCodeFlowJSON(data)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.NewGroupsCodeFlowJSON()", err.Error())
	}
	return groups, nil
}

// InitFromCodeFlowRawURL Initialization from the url parameters received after successfully completing the authorization code and receiving an access token
func (groups *Groups) InitFromCodeFlowRawURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.Groups.InitFromCodeFlowRawURL()", err.Error())
	}

	return groups.InitFromCodeFlowURL(u)
}

// InitFromCodeFlowURL Initialization from the url parameters received after successfully completing the authorization code and receiving an access token
func (groups *Groups) InitFromCodeFlowURL(u *url.URL) error {
	v, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.Groups.InitFromCodeFlowURL()", err.Error())
	}

	if errType := v.Get("error"); errType != "" {
		return &internalError.AuthCodeFlowError{
			Type:        errType,
			Reason:      v.Get("error_reason"),
			Description: v.Get("error_description"),
		}
	}

	for key, vs := range v {
		if strings.HasPrefix(key, "access_token_") {
			groupID, _ := strconv.Atoi(key[13:])
			if len(vs) > 0 {
				accessToken := vs[0]
				groups.Groups = append(groups.Groups, Group{
					ID:          groupID,
					AccessToken: accessToken,
				})
			}
		}

		if key == "expires_in" {
			groups.ExpiresIn, _ = strconv.Atoi(vs[0])
		}
	}

	return nil
}

func (groups *Groups) InitFromCodeFlowJSON(data []byte) error {
	var authError internalError.AuthCodeFlowError

	err := json.Unmarshal(data, &authError)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.Groups.InitFromCodeFlowJSON()", err.Error())
	}

	if authError.Type != "" {
		return &authError
	}

	err = json.Unmarshal(data, &groups)
	if err != nil {
		return internalError.Error("Actor.CodeFlow.Groups.InitFromCodeFlowJSON()", err.Error())
	}

	return nil
}

// VKID

type UserVKIDCodeFlowAuthorizeURL struct {
	ClientID      int    `json:"client_id"`    // required
	RedirectURI   string `json:"redirect_uri"` // required
	State         string `json:"state,omitempty"`
	CodeChallenge string `json:"code_challenge,omitempty"`
}

func (*UserVKIDCodeFlowAuthorizeURL) GenerateCodeVerifier() string {
	b := make([]byte, vkidCodeVerifierLength)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = vkidGenerateCodeCharset[s.Intn(len(vkidGenerateCodeCharset))]
	}
	return string(b)
}

func (user *UserVKIDCodeFlowAuthorizeURL) GenerateCodeChallengeByCodeVerifier(codeVerifier string) {
	hash := sha256.Sum256([]byte(codeVerifier))
	user.CodeChallenge = base64.RawURLEncoding.EncodeToString(hash[:])
}

func (user *UserVKIDCodeFlowAuthorizeURL) GenerateCodeChallenge() {
	hash := sha256.Sum256([]byte(user.GenerateCodeVerifier()))
	user.CodeChallenge = base64.RawURLEncoding.EncodeToString(hash[:])
}

func (user *UserVKIDCodeFlowAuthorizeURL) GenerateState() {
	b := make([]byte, 16)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = vkidGenerateCodeCharset[s.Intn(len(vkidGenerateCodeCharset))]
	}
	user.State = string(b)
}

// Build Assembling url to request an authorization code, which is used to obtain a token in vk id code flow auth
// If scope is less than 0 then it is not added
func (user *UserVKIDCodeFlowAuthorizeURL) Build() (*url.URL, error) {
	u, err := url.Parse(api.AuthVKIDEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, internalError.Error("Actor.CodeFlow.UserVKIDCodeFlowAuthorizeURL.Build()", err.Error())
	}

	q := url.Values{
		constants.ParameterNameClientID: {strconv.Itoa(user.ClientID)},
	}

	if user.RedirectURI == "" {
		q.Set(constants.ParameterNameRedirectURI, AuthRedirectURI)
	} else {
		q.Set(constants.ParameterNameRedirectURI, user.RedirectURI)
	}

	q.Set(constants.ParameterNameResponseType, constants.ParameterNameCode)
	q.Set(constants.ParameterNameState, user.State)
	q.Set(constants.ParameterNameCodeChallenge, user.CodeChallenge)
	q.Set(constants.ParameterNameCodeChallengeMethod, vkidCodeChallengeMethod)

	u.RawQuery = q.Encode()

	return u, nil
}
