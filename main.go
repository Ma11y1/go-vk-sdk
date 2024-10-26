package main

import (
	"context"
	"fmt"
	"go-vk-sdk/action"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/errors"
	"go-vk-sdk/longPollUser"
	"go-vk-sdk/request"
)

// TODO - Change all generated errors to errors.InternalError or call errors.Error() function
// TODO - Complete the methods of action
// TODO - Complete request structures for all API methods
// TODO - Test all requests
// TODO - Make display of logger information using logger.Log()

const (
	APP_ID        = 11
	APP_SECRET    = "11"
	SERVICE_TOKEN = "Service-Token-123456"
	CLIENT_ID     = 22222222222
	CLIENT_TOKEN  = "Client-Token"
)

func main() {
	user := DirectAuth()
	list := GetFriendsList(user)
	fmt.Println(list)
}

func DirectAuth() *actor.User {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	requestDirectAuth := actions.Auth.UserDirectAuth().
		ClientID(APP_ID).
		ClientSecret(APP_SECRET).
		Scope(actor.ScopeUserMessages + actor.ScopeUserFriends).
		Username("+7username").
		Password("pwd")

	responseDirectAuth, err := requestDirectAuth.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseDirectAuth.Error != nil { // if vk server error auth
		fmt.Println(responseDirectAuth.Error.Description)
		return nil
	}

	return responseDirectAuth.User
}

func CodeFlowAuth() *actor.User {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	// Create URL for get auth code
	getAuthCodeURL, err := (&actor.UserCodeFlowAuthorizeURL{
		ClientID:    APP_ID,
		RedirectURI: actor.AuthRedirectURI,
		Scope:       actor.ScopeUserFriends,
	}).Build()
	if err != nil {
		panic(err)
	}

	fmt.Println(getAuthCodeURL)

	code, err := actor.NewCodeFlowAuthorizeCodeRawURL("https://redirect.domain#code=code12345")
	if err != nil {
		panic(err)
	}

	// Swap auth code to token
	response, err := actions.Auth.UserCodeFlow().
		ClientID(APP_ID).
		ClientSecret(APP_SECRET).
		Code(code.Code).
		RedirectURI(actor.AuthRedirectURI).
		Exec(context.Background())
	if err != nil {
		panic(err)
	}

	if response.Error.Type != "" {
		panic(response.Error.Error())
	}

	return response.User
}

func GetFriendsList(user *actor.User) []int {
	a := api.NewAPI()

	// actions have not yet been completed
	requestGetProfileInfo := request.NewFriendsGetRequest(a, user)

	responseGetProfileInfo, err := requestGetProfileInfo.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseGetProfileInfo.Error.Code != errors.None {
		fmt.Println(responseGetProfileInfo.Error.Message)
		return nil
	} else {
		return responseGetProfileInfo.Response.Items
	}
}

func LongPoll(user *actor.User) {
	a := api.NewAPI()

	lp := longPollUser.NewLongPoll(a, user, longPollUser.ExtraOptionsModeReceiveAttachments)

	for update := range lp.Updates() {
		switch update.Type {
		case longPollUser.EventTypeMessageNew:
			fmt.Println(update.Event.(longPollUser.EventMessageNew))
		}
	}
}
