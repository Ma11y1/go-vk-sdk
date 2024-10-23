package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Account Doc: https://dev.vk.com/ru/method/account
type Account struct {
	BaseAction
}

// Ban Doc: https://dev.vk.com/method/account.ban
func (a *Account) Ban(actor actor.Actor) *request.AccountBanRequest {
	return request.NewAccountBanRequest(a.api, actor)
}

// Unban Doc: https://dev.vk.com/method/account.unban
func (a *Account) Unban(actor actor.Actor) *request.AccountUnbanRequest {
	return request.NewAccountUnbanRequest(a.api, actor)
}

// GetBanned Doc: https://dev.vk.com/method/account.getBanned
func (a *Account) GetBanned(actor actor.Actor) *request.AccountGetBannedRequest {
	return request.NewAccountGetBannedRequest(a.api, actor)
}

// ChangePassword Doc: https://dev.vk.com/method/account.changePassword
func (a *Account) ChangePassword(actor actor.Actor) *request.AccountChangePasswordRequest {
	return request.NewAccountChangePasswordRequest(a.api, actor)
}

// GetActiveOffers Doc: https://dev.vk.com/method/account.getActiveOffers
func (a *Account) GetActiveOffers(actor actor.Actor) *request.AccountGetActiveOffersRequest {
	return request.NewAccountGetActiveOffersRequest(a.api, actor)
}

// GetAppPermissions Doc: https://dev.vk.com/method/account.getAppPermissions
func (a *Account) GetAppPermissions(actor actor.Actor) *request.AccountGetAppPermissionsRequest {
	return request.NewAccountGetAppPermissionsRequest(a.api, actor)
}

// GetCounters Doc: https://dev.vk.com/method/account.getCounters
func (a *Account) GetCounters(actor actor.Actor) *request.AccountGetCountersRequest {
	return request.NewAccountGetCountersRequest(a.api, actor)
}

// GetInfo Doc: https://dev.vk.com/method/account.getInfo
func (a *Account) GetInfo(actor actor.Actor) *request.AccountGetInfoRequest {
	return request.NewAccountGetInfoRequest(a.api, actor)
}

// GetProfileInfo Doc: https://dev.vk.com/method/account.getProfileInfo
func (a *Account) GetProfileInfo(actor actor.Actor) *request.AccountGetProfileInfoRequest {
	return request.NewAccountGetProfileInfoRequest(a.api, actor)
}

// GetPushSettings Doc: https://dev.vk.com/method/account.getPushSettings
func (a *Account) GetPushSettings(actor actor.Actor) *request.AccountGetPushSettingsRequest {
	return request.NewAccountGetPushSettingsRequest(a.api, actor)
}

// SaveProfileInfo Doc: https://dev.vk.com/method/account.saveProfileInfo
func (a *Account) SaveProfileInfo(actor actor.Actor) *request.AccountSaveProfileInfoRequest {
	return request.NewAccountSaveProfileInfoRequest(a.api, actor)
}

// SetInfo Doc: https://dev.vk.com/method/account.setInfo
func (a *Account) SetInfo(actor actor.Actor) *request.AccountSetInfoRequest {
	return request.NewAccountSetInfoRequest(a.api, actor)
}

// SetOnline Doc: https://dev.vk.com/method/account.setOnline
func (a *Account) SetOnline(actor actor.Actor) *request.AccountSetOnlineRequest {
	return request.NewAccountSetOnlineRequest(a.api, actor)
}

// SetOffline Doc: https://dev.vk.com/method/account.setOffline
func (a *Account) SetOffline(actor actor.Actor) *request.AccountSetOfflineRequest {
	return request.NewAccountSetOfflineRequest(a.api, actor)
}

// SetPushSettings Doc: https://dev.vk.com/method/account.setPushSettings
func (a *Account) SetPushSettings(actor actor.Actor) *request.AccountSetPushSettingsRequest {
	return request.NewAccountSetPushSettingsRequest(a.api, actor)
}

// SetSilenceMode Doc: https://dev.vk.com/method/account.setSilenceMode
func (a *Account) SetSilenceMode(actor actor.Actor) *request.AccountSetSilenceModeRequest {
	return request.NewAccountSetSilenceModeRequest(a.api, actor)
}

// RegisterDevice Doc: https://dev.vk.com/method/account.registerDevice
func (a *Account) RegisterDevice(actor actor.Actor) *request.AccountRegisterDeviceRequest {
	return request.NewAccountRegisterDeviceRequest(a.api, actor)
}

// UnregisterDevice Doc: https://dev.vk.com/method/account.unregisterDevice
func (a *Account) UnregisterDevice(actor actor.Actor) *request.AccountUnregisterDeviceRequest {
	return request.NewAccountUnregisterDeviceRequest(a.api, actor)
}
