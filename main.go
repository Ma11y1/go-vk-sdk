package main

import (
	"context"
	"fmt"
	"go-vk-sdk/action"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/errors"
	"go-vk-sdk/longPollUser"
)

// TODO - Change all generated errors to errors.InternalError or call errors.Error() function
// TODO - Complete the methods of action
// TODO - Complete request structures for all API methods
// TODO - Test all requests
// TODO - Make display of logger information using logger.Log()

const (
	APP_ID        = 111111
	APP_SECRET    = "Secret-Secret"
	SERVICE_TOKEN = "Service-Token-123456"
	CLIENT_ID     = 22222222222
	CLIENT_TOKEN  = "Client-Token-123456"
)

func main() {
	a := api.NewAPI()
	actions := action.NewRouter(a)

	requestDirectAuth := actions.Auth.UserDirectAuth().
		ClientID(APP_ID).
		ClientSecret(APP_SECRET).
		Scope(actor.ScopeUserMessages + actor.ScopeUserFriends).
		Username("+7username").
		Password("pwd")
	//Password("pwd").Code("1111")

	responseDirectAuth, err := requestDirectAuth.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseDirectAuth.Error != nil { // if vk server error auth
		fmt.Println(responseDirectAuth.Error.Description)
	}

	user := responseDirectAuth.User

	requestGetProfileInfo := actions.Account.GetProfileInfo(user)
	responseGetProfileInfo, err := requestGetProfileInfo.Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	if responseGetProfileInfo.Error.Code != errors.None {
		fmt.Println(responseGetProfileInfo.Error.Message)
	} else {
		fmt.Println(responseGetProfileInfo.Response)
	}

	// LongPoll
	lp := longPollUser.NewLongPoll(a, user, longPollUser.ExtraOptionsModeReceiveAttachments)

	for update := range lp.Updates() {
		switch update.Type {
		case longPollUser.EventTypeMessageNew:
			fmt.Println(update.Event.(longPollUser.EventMessageNew))
		}
	}
}
