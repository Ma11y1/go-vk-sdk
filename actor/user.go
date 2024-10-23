package actor

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	errors2 "go-vk-sdk/errors"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type User struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ID          int    `json:"user_id"`
	Email       string `json:"email,omitempty"`
	State       string `json:"state,omitempty"`
}

func NewUser(id int, token string) *User {
	return &User{
		AccessToken: token,
		ID:          id,
	}
}

// NewUserFromDirectAuthURL Example url blank from direct auth: https://oauth.vk.com/blank.html#success=1&access_token=vk1.a.xxxx&user_id=xxx
func NewUserFromDirectAuthURL(rawURL string) (*User, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, err
	}

	if q.Get("success") != "1" {
		return nil, errors.New("receiving token unsuccessfully")
	}

	if q.Get("failed") == "1" {
		return nil, errors.New("authorization failed")
	}

	user := &User{}
	user.ID, err = strconv.Atoi(q.Get("user_id"))
	if err != nil {
		return nil, err
	}

	user.AccessToken = q.Get("access_token")

	return user, nil
}

func NewUserFromURL(rawURL string) (*User, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, err
	}

	if errType := q.Get("errors"); errType != "" {
		return nil, &errors2.AuthCodeFlowError{
			Type:        errType,
			Reason:      q.Get("error_reason"),
			Description: q.Get("error_description"),
		}
	}

	expiresIn, _ := strconv.Atoi(q.Get("expires_in"))
	userID, _ := strconv.Atoi(q.Get("user_id"))
	user := &User{
		ID:          userID,
		AccessToken: q.Get("access_token"),
		ExpiresIn:   expiresIn,
		State:       q.Get("state"),
		Email:       q.Get("email"),
	}

	return user, nil
}

func NewUserFromJSON(data []byte) (*User, error) {
	var authError errors2.AuthCodeFlowError
	err := json.Unmarshal(data, &authError)
	if err != nil {
		return nil, err
	}

	if authError.Type != "" {
		return nil, &authError
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) GetAccessToken() string {
	return u.AccessToken
}

func (u *User) BuildCodeFlowAuthorizeURL(params *UserAuthorizeURLParams) (*url.URL, error) {
	return NewUserCodeFlowAuthorizeURL(params)
}

func (u *User) BuildImplicitFlowAuthorizeURL(params *UserAuthorizeURLParams) (*url.URL, error) {
	return NewUserImplicitFlowAuthorizeURL(params)
}

type UserVKID struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	UserID       int    `json:"user_id"`
	IDToken      string `json:"id_token"`
	Scope        int    `json:"scope"`
}

type UserAuthorizeURLParams struct {
	ClientID    int    `json:"client_id"`    // required
	RedirectURI string `json:"redirect_uri"` // required
	Display     string `json:"display"`      // page, popup, mobile
	Scope       int    `json:"scope"`
	State       string `json:"state"`
}

func userAuthorizeURLValues(params *UserAuthorizeURLParams) *url.Values {
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

	return &q
}

// NewUserCodeFlowAuthorizeURL Assembling url to request an authorization code, which is used to obtain a token in code flow auth
// If scope is less than 0 then it is not added
func NewUserCodeFlowAuthorizeURL(params *UserAuthorizeURLParams) (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, err
	}

	q := userAuthorizeURLValues(params)
	q.Set(constants.ParameterNameResponseType, constants.ParameterNameCode)

	u.RawQuery = q.Encode()

	return u, nil
}

// NewUserImplicitFlowAuthorizeURL Assembling url to request an authorization code, which is used to obtain a token in implicit flow auth
// If scope is less than 0 then it is not added
func NewUserImplicitFlowAuthorizeURL(params *UserAuthorizeURLParams) (*url.URL, error) {
	u, err := url.Parse(api.AuthEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, err
	}

	q := userAuthorizeURLValues(params)
	q.Set(constants.ParameterNameResponseType, constants.ParameterNameCode)

	u.RawQuery = q.Encode()

	return u, nil
}

type UserVKIDAuthorizeURLParams struct {
	ClientID      int    `json:"client_id"`    // required
	RedirectURI   string `json:"redirect_uri"` // required
	State         string `json:"state"`
	CodeChallenge string `json:"code_challenge"`
}

func (p *UserVKIDAuthorizeURLParams) GenerateCodeVerifier() string {
	b := make([]byte, vkidCodeVerifierLength)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = vkidGenerateCodeCharset[s.Intn(len(vkidGenerateCodeCharset))]
	}
	return string(b)
}

func (p *UserVKIDAuthorizeURLParams) GenerateCodeChallengeByCodeVerifier(codeVerifier string) {
	hash := sha256.Sum256([]byte(codeVerifier))
	p.CodeChallenge = base64.RawURLEncoding.EncodeToString(hash[:])
}

func (p *UserVKIDAuthorizeURLParams) GenerateCodeChallenge() {
	hash := sha256.Sum256([]byte(p.GenerateCodeVerifier()))
	p.CodeChallenge = base64.RawURLEncoding.EncodeToString(hash[:])
}

func (p *UserVKIDAuthorizeURLParams) GenerateState() {
	b := make([]byte, 16)
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = vkidGenerateCodeCharset[s.Intn(len(vkidGenerateCodeCharset))]
	}
	p.State = string(b)
}

// NewUserVKIDCodeFlowAuthorizeURL Assembling url to request an authorization code, which is used to obtain a token in vk id code flow auth
// If scope is less than 0 then it is not added
func NewUserVKIDCodeFlowAuthorizeURL(params *UserVKIDAuthorizeURLParams) (*url.URL, error) {
	u, err := url.Parse(api.AuthVKIDEndpoint + api.AuthAuthorizeMethod)
	if err != nil {
		return nil, err
	}

	q := url.Values{
		constants.ParameterNameClientID: {strconv.Itoa(params.ClientID)},
	}

	if params.RedirectURI == "" {
		q.Set(constants.ParameterNameRedirectURI, api.AuthRedirectURI)
	} else {
		q.Set(constants.ParameterNameRedirectURI, params.RedirectURI)
	}

	q.Set(constants.ParameterNameResponseType, constants.ParameterNameCode)
	q.Set(constants.ParameterNameState, params.State)
	q.Set(constants.ParameterNameCodeChallenge, params.CodeChallenge)
	q.Set(constants.ParameterNameCodeChallengeMethod, vkidCodeChallengeMethod)

	u.RawQuery = q.Encode()

	return u, nil
}
