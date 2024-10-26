package actor

import (
	internalError "go-vk-sdk/errors"
	"net/url"
	"strconv"
)

// NewUserDirectRawURL
//
//	Creating new user from  raw URL received during direct authorization
//	Example url blank from direct auth: https://oauth.vk.com/blank.html#success=1&access_token=vk1.a.xxxx&user_id=xxx
func NewUserDirectRawURL(rawURL string) (*User, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, internalError.Error("Actor.Direct.NewUserDirectRawURL()", err.Error())
	}
	return NewUserDirectURL(u)
}

// NewUserDirectURL
//
//	Creating new user from  URL received during direct authorization
//	Example url blank from direct auth: https://oauth.vk.com/blank.html#success=1&access_token=vk1.a.xxxx&user_id=xxx
func NewUserDirectURL(u *url.URL) (*User, error) {
	q, err := url.ParseQuery(u.Fragment)
	if err != nil {
		return nil, internalError.Error("Actor.Direct.NewUserDirectURL()", err.Error())
	}

	if q.Get("success") != "1" {
		return nil, internalError.Error("Actor.Direct.NewUserDirectURL()", "Receiving token unsuccessfully URL: "+u.String())
	}

	if q.Get("failed") == "1" {
		return nil, internalError.Error("Actor.Direct.NewUserDirectURL()", "Authorization failed URL: "+u.String())
	}

	user := &User{}
	user.ID, err = strconv.Atoi(q.Get("user_id"))
	if err != nil {
		return nil, internalError.Error("Actor.Direct.NewUserDirectURL()", err.Error())
	}

	user.AccessToken = q.Get("access_token")

	return user, nil
}
