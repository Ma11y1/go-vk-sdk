package response

import (
	"encoding/json"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/errors"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/method/auth

type AuthDirectResponse struct {
	User  *actor.User
	Error *errors.AuthDirectError // {"access_token":"vk1.UMx","expires_in":86179,"user_id":111}
}

func (a *AuthDirectResponse) UnmarshalJSON(data []byte) error {
	var user actor.User

	err := json.Unmarshal(data, &user)
	if err == nil && user.AccessToken != "" {
		a.User = &user
		return nil
	}

	var directError errors.AuthDirectError
	err = json.Unmarshal(data, &directError)
	if err == nil && directError.Type != "" {
		a.Error = &directError
		return nil
	}

	return fmt.Errorf("unable to parse direct auth response: %s", string(data))
}

type AuthUserCodeFlowResponse struct {
	User  *actor.User               // {"access_token":"vk1.UMx","expires_in":86179,"user_id":111}
	Error *errors.AuthCodeFlowError // {"errors":"invalid_grant","error_description":"Code is invalid or expired."}
}

func (a *AuthUserCodeFlowResponse) UnmarshalJSON(data []byte) error {
	var user actor.User

	err := json.Unmarshal(data, &user)
	if err == nil && user.AccessToken != "" {
		a.User = &user
		return nil
	}

	var codeFlowError errors.AuthCodeFlowError
	err = json.Unmarshal(data, &codeFlowError)
	if err == nil && codeFlowError.Type != "" {
		a.Error = &codeFlowError
		return nil
	}

	return fmt.Errorf("unable to parse user code flow auth response: %s", string(data))
}

type AuthGroupCodeFlowResponse struct {
	Group *actor.Group
	Error *errors.AuthCodeFlowError
}

func (a *AuthGroupCodeFlowResponse) UnmarshalJSON(data []byte) error {
	var group actor.Group

	err := json.Unmarshal(data, &group)
	if err == nil && group.AccessToken != "" {
		a.Group = &group
		return nil
	}

	var directError errors.AuthCodeFlowError
	err = json.Unmarshal(data, &directError)
	if err == nil && directError.Type != "" {
		a.Error = &directError
		return nil
	}

	return fmt.Errorf("unable to parse groip code flow auth response: %s", string(data))
}

type AuthUserVKIDCodeFlowResponse struct {
	User  *actor.UserVKID
	Error *errors.AuthCodeFlowError
}

func (a *AuthUserVKIDCodeFlowResponse) UnmarshalJSON(data []byte) error {
	var user actor.UserVKID

	err := json.Unmarshal(data, &user)
	if err == nil && user.AccessToken != "" {
		a.User = &user
		return nil
	}

	var directError errors.AuthCodeFlowError
	err = json.Unmarshal(data, &directError)
	if err == nil && directError.Type != "" {
		a.Error = &directError
		return nil
	}

	return fmt.Errorf("unable to parse user code flow auth response: %s", string(data))
}

type AuthRestoreResponse struct {
	BaseResponse
	Response struct {
		Success objects.NumberFlagBool `json:"success"`
		SID     string                 `json:"sid"`
	} `json:"response"`
}
