package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/method/secure

// SecureAddAppEventRequest defines the request for secure.addAppEvent
//
// Adds information about user achievements in the application.
// Doc: https://dev.vk.com/method/secure.addAppEvent
type SecureAddAppEventRequest struct {
	BaseRequest
}

// NewSecureAddAppEventRequest creates a new request for secure.addAppEvent
func NewSecureAddAppEventRequest(a *api.API, actor actor.Actor) *SecureAddAppEventRequest {
	return &SecureAddAppEventRequest{*NewMethodBaseRequest(a, actor, "secure.addAppEvent")}
}

// Exec executes the request and unmarshals the response into SecureAddAppEventResponse
func (r *SecureAddAppEventRequest) Exec(ctx context.Context) (response response.SecureAddAppEventResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureCheckTokenRequest defines the request for secure.checkToken
//
// Verifies that the user's access token is issued to the application associated with the passed service access key.
// Doc: https://dev.vk.com/method/secure.checkToken
type SecureCheckTokenRequest struct {
	BaseRequest
}

// NewSecureCheckTokenRequest creates a new request for secure.checkToken
func NewSecureCheckTokenRequest(a *api.API, actor actor.Actor) *SecureCheckTokenRequest {
	return &SecureCheckTokenRequest{*NewMethodBaseRequest(a, actor, "secure.checkToken")}
}

// Exec executes the request and unmarshals the response into SecureCheckTokenResponse
func (r *SecureCheckTokenRequest) Exec(ctx context.Context) (response response.SecureCheckTokenResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureGetAppBalanceRequest defines the request for secure.getAppBalance
//
// Returns the application's payment balance in hundredths of a voice.
// Doc: https://dev.vk.com/method/secure.getAppBalance
type SecureGetAppBalanceRequest struct {
	BaseRequest
}

// NewSecureGetAppBalanceRequest creates a new request for secure.getAppBalance
func NewSecureGetAppBalanceRequest(a *api.API, actor actor.Actor) *SecureGetAppBalanceRequest {
	return &SecureGetAppBalanceRequest{*NewMethodBaseRequest(a, actor, "secure.getAppBalance")}
}

// Exec executes the request and unmarshals the response into SecureGetAppBalanceResponse
func (r *SecureGetAppBalanceRequest) Exec(ctx context.Context) (response response.SecureGetAppBalanceResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureGetSMSHistoryRequest defines the request for secure.getSMSHistory
//
// Retrieves the list of SMS notifications sent by the application.
// Doc: https://dev.vk.com/method/secure.getSMSHistory
type SecureGetSMSHistoryRequest struct {
	BaseRequest
}

// NewSecureGetSMSHistoryRequest creates a new request for secure.getSMSHistory
func NewSecureGetSMSHistoryRequest(a *api.API, actor actor.Actor) *SecureGetSMSHistoryRequest {
	return &SecureGetSMSHistoryRequest{*NewMethodBaseRequest(a, actor, "secure.getSMSHistory")}
}

// Exec executes the request and unmarshals the response into SecureGetSMSHistoryResponse
func (r *SecureGetSMSHistoryRequest) Exec(ctx context.Context) (response response.SecureGetSMSHistoryResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureGetTransactionsHistoryRequest defines the request for secure.getTransactionsHistory
//
// Retrieves the history of transactions involving voice transfers between users and the application.
// Doc: https://dev.vk.com/method/secure.getTransactionsHistory
type SecureGetTransactionsHistoryRequest struct {
	BaseRequest
}

// NewSecureGetTransactionsHistoryRequest creates a new request for secure.getTransactionsHistory
func NewSecureGetTransactionsHistoryRequest(a *api.API, actor actor.Actor) *SecureGetTransactionsHistoryRequest {
	return &SecureGetTransactionsHistoryRequest{*NewMethodBaseRequest(a, actor, "secure.getTransactionsHistory")}
}

// Exec executes the request and unmarshals the response into SecureGetTransactionsHistoryResponse
func (r *SecureGetTransactionsHistoryRequest) Exec(ctx context.Context) (response response.SecureGetTransactionsHistoryResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureGetUserLevelRequest defines the request for secure.getUserLevel
//
// Returns the previously set game level of one or more users in the application.
// Doc: https://dev.vk.com/method/secure.getUserLevel
type SecureGetUserLevelRequest struct {
	BaseRequest
}

// NewSecureGetUserLevelRequest creates a new request for secure.getUserLevel
func NewSecureGetUserLevelRequest(a *api.API, actor actor.Actor) *SecureGetUserLevelRequest {
	return &SecureGetUserLevelRequest{*NewMethodBaseRequest(a, actor, "secure.getUserLevel")}
}

// Exec executes the request and unmarshals the response into SecureGetUserLevelResponse
func (r *SecureGetUserLevelRequest) Exec(ctx context.Context) (response response.SecureGetUserLevelResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureGiveEventStickerRequest defines the request for secure.giveEventSticker
//
// Issues a sticker to the user and unlocks a game achievement.
// Doc: https://dev.vk.com/method/secure.giveEventSticker
type SecureGiveEventStickerRequest struct {
	BaseRequest
}

// NewSecureGiveEventStickerRequest creates a new request for secure.giveEventSticker
func NewSecureGiveEventStickerRequest(a *api.API, actor actor.Actor) *SecureGiveEventStickerRequest {
	return &SecureGiveEventStickerRequest{*NewMethodBaseRequest(a, actor, "secure.giveEventSticker")}
}

// Exec executes the request and unmarshals the response into SecureGiveEventStickerResponse
func (r *SecureGiveEventStickerRequest) Exec(ctx context.Context) (response response.SecureGiveEventStickerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureSendNotificationRequest defines the request for secure.sendNotification
//
// Sends a notification to one or more users.
// Doc: https://dev.vk.com/method/secure.sendNotification
type SecureSendNotificationRequest struct {
	BaseRequest
}

// NewSecureSendNotificationRequest creates a new request for secure.sendNotification
func NewSecureSendNotificationRequest(a *api.API, actor actor.Actor) *SecureSendNotificationRequest {
	return &SecureSendNotificationRequest{*NewMethodBaseRequest(a, actor, "secure.sendNotification")}
}

// Exec executes the request and unmarshals the response into SecureSendNotificationResponse
func (r *SecureSendNotificationRequest) Exec(ctx context.Context) (response response.SecureSendNotificationResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureSendSMSNotificationRequest defines the request for secure.sendSMSNotification
//
// Sends an SMS notification to the user’s mobile phone.
// Doc: https://dev.vk.com/method/secure.sendSMSNotification
type SecureSendSMSNotificationRequest struct {
	BaseRequest
}

// NewSecureSendSMSNotificationRequest creates a new request for secure.sendSMSNotification
func NewSecureSendSMSNotificationRequest(a *api.API, actor actor.Actor) *SecureSendSMSNotificationRequest {
	return &SecureSendSMSNotificationRequest{*NewMethodBaseRequest(a, actor, "secure.sendSMSNotification")}
}

// Exec executes the request and unmarshals the response into SecureSendSMSNotificationResponse
func (r *SecureSendSMSNotificationRequest) Exec(ctx context.Context) (response response.SecureSendSMSNotificationResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// SecureSetCounterRequest defines the request for secure.setCounter
//
// Sets a counter displayed in bold in the user's left menu.
// Doc: https://dev.vk.com/method/secure.setCounter
type SecureSetCounterRequest struct {
	BaseRequest
}

// NewSecureSetCounterRequest creates a new request for secure.setCounter
func NewSecureSetCounterRequest(a *api.API, actor actor.Actor) *SecureSetCounterRequest {
	return &SecureSetCounterRequest{*NewMethodBaseRequest(a, actor, "secure.setCounter")}
}

// Exec executes the request and unmarshals the response into SecureSetCounterResponse
func (r *SecureSetCounterRequest) Exec(ctx context.Context) (response response.SecureSetCounterResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
