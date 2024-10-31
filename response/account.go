package response

import "go-vk-sdk/objects"

// AccountBanResponse defines the response for account.ban
type AccountBanResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountChangePasswordResponse defines the response for account.changePassword
type AccountChangePasswordResponse struct {
	BaseResponse
	Response struct {
		Token string `json:"token"`
	} `json:"response"`
}

// AccountGetActiveOffersResponse defines the response for account.getActiveOffers
// An array consisting of the total number of special offers targeted to the current user (the first element),
// and a list of objects with information about the offers.
// If no special offer is targeted to the user, the array will contain element 0 (the number of special offers).
type AccountGetActiveOffersResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.AccountOffer `json:"items"`
	} `json:"response"`
}

// AccountGetAppPermissionsResponse defines the response for account.getAppPermissions
type AccountGetAppPermissionsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountGetBannedResponse defines the response for account.getBanned
type AccountGetBannedResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

// AccountGetCountersResponse defines the response for account.getCounters
type AccountGetCountersResponse struct {
	BaseResponse
	Response objects.AccountCounters `json:"response"`
}

// AccountGetInfoResponse defines the response for account.getInfo
type AccountGetInfoResponse struct {
	BaseResponse
	Response objects.AccountInfo `json:"response"`
}

// AccountGetProfileInfoResponse defines the response for account.getProfileInfo
type AccountGetProfileInfoResponse struct {
	BaseResponse
	Response objects.AccountProfileInfo `json:"response"`
}

// AccountGetPushSettingsResponse defines the response for account.getPushSettings
type AccountGetPushSettingsResponse struct {
	BaseResponse
	Response objects.AccountPushSettings `json:"response"`
}

// AccountRegisterDeviceResponse defines the response for account.registerDevice
type AccountRegisterDeviceResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountSaveProfileInfoResponse defines the response for account.saveProfileInfo
type AccountSaveProfileInfoResponse struct {
	BaseResponse
	Response struct {
		Changed     objects.BoolInt            `json:"changed"`
		NameRequest objects.AccountNameRequest `json:"name_request"`
	} `json:"response"`
}

// AccountSetInfoResponse defines the response for account.setInfo
type AccountSetInfoResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountSetOfflineResponse defines the response for account.setOffline
type AccountSetOfflineResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountSetOnlineResponse defines the response for account.setOnline
type AccountSetOnlineResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountSetPushSettingsResponse defines the response for account.setPushSettings
type AccountSetPushSettingsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountSetSilenceModeResponse defines the response for account.setSilenceMode
type AccountSetSilenceModeResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountUnbanResponse defines the response for account.unban
type AccountUnbanResponse struct {
	BaseResponse
	Response int `json:"response"`
}

// AccountUnregisterDeviceResponse defines the response for account.unregisterDevice
type AccountUnregisterDeviceResponse struct {
	BaseResponse
	Response int `json:"response"`
}
