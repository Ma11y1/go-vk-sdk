package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Secure Doc: https://dev.vk.com/ru/method/secure
type Secure struct {
	BaseAction
}

// AddAppEvent Doc: https://dev.vk.com/ru/method/secure.addAppEvent
func (s *Secure) AddAppEvent(actor actor.Actor) *request.SecureAddAppEventRequest {
	return request.NewSecureAddAppEventRequest(s.api, actor)
}

// CheckToken Doc: https://dev.vk.com/ru/method/secure.checkToken
func (s *Secure) CheckToken(actor actor.Actor) *request.SecureCheckTokenRequest {
	return request.NewSecureCheckTokenRequest(s.api, actor)
}

// GetAppBalance Doc: https://dev.vk.com/ru/method/secure.getAppBalance
func (s *Secure) GetAppBalance(actor actor.Actor) *request.SecureGetAppBalanceRequest {
	return request.NewSecureGetAppBalanceRequest(s.api, actor)
}

// GetSMSHistory Doc: https://dev.vk.com/ru/method/secure.getSMSHistory
func (s *Secure) GetSMSHistory(actor actor.Actor) *request.SecureGetSMSHistoryRequest {
	return request.NewSecureGetSMSHistoryRequest(s.api, actor)
}

// GetTransactionsHistory Doc: https://dev.vk.com/ru/method/secure.getTransactionsHistory
func (s *Secure) GetTransactionsHistory(actor actor.Actor) *request.SecureGetTransactionsHistoryRequest {
	return request.NewSecureGetTransactionsHistoryRequest(s.api, actor)
}

// GetUserLevel Doc: https://dev.vk.com/ru/method/secure.getUserLevel
func (s *Secure) GetUserLevel(actor actor.Actor) *request.SecureGetUserLevelRequest {
	return request.NewSecureGetUserLevelRequest(s.api, actor)
}

// GiveEventSticker Doc: https://dev.vk.com/ru/method/secure.giveEventSticker
func (s *Secure) GiveEventSticker(actor actor.Actor) *request.SecureGiveEventStickerRequest {
	return request.NewSecureGiveEventStickerRequest(s.api, actor)
}

// SendNotification Doc: https://dev.vk.com/ru/method/secure.sendNotification
func (s *Secure) SendNotification(actor actor.Actor) *request.SecureSendNotificationRequest {
	return request.NewSecureSendNotificationRequest(s.api, actor)
}

// SendSMSNotification Doc: https://dev.vk.com/ru/method/secure.sendSMSNotification
func (s *Secure) SendSMSNotification(actor actor.Actor) *request.SecureSendSMSNotificationRequest {
	return request.NewSecureSendSMSNotificationRequest(s.api, actor)
}

// SetCounter Doc: https://dev.vk.com/ru/method/secure.setCounter
func (s *Secure) SetCounter(actor actor.Actor) *request.SecureSetCounterRequest {
	return request.NewSecureSetCounterRequest(s.api, actor)
}
