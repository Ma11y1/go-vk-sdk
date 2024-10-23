package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"strconv"
	"strings"
)

// Doc: https://dev.vk.com/ru/method/account

// AccountBanRequest defines the request for account.ban
// Adds a user or group to the blacklist.
// If the specified user is a friend of the current user or has an incoming or outgoing friend request from him, then to add the user to the blacklist, your application must have the rights: friends
// Doc: https://dev.vk.com/method/account.ban
type AccountBanRequest struct {
	BaseRequest
}

// NewAccountBanRequest creates a new request for account.ban
func NewAccountBanRequest(a *api.API, actor actor.Actor) *AccountBanRequest {
	return &AccountBanRequest{*NewMethodBaseRequest(a, actor, "account.ban")}
}

// Exec executes the request and unmarshals the response into AccountBanResponse
func (r *AccountBanRequest) Exec(ctx context.Context) (response response.AccountBanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OwnerID ID of the user or community that will be added to the blacklist.
func (r *AccountBanRequest) OwnerID(id int) *AccountBanRequest {
	r.parameters.Set(constants.ParameterNameOwnerID, strconv.Itoa(id))
	return r
}

// AccountChangePasswordRequest defines the request for account.changePassword
// Allows you to change the user's password after successfully restoring access to the account via SMS using the auth.restore method.
// Doc: https://dev.vk.com/method/account.changePassword
type AccountChangePasswordRequest struct {
	BaseRequest
}

// NewAccountChangePasswordRequest creates a new request for account.changePassword
func NewAccountChangePasswordRequest(a *api.API, actor actor.Actor) *AccountChangePasswordRequest {
	return &AccountChangePasswordRequest{*NewMethodBaseRequest(a, actor, "account.changePassword")}
}

// Exec executes the request and unmarshals the response into AccountChangePasswordResponse
func (r *AccountChangePasswordRequest) Exec(ctx context.Context) (response response.AccountChangePasswordResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// RestoreSID sets the session identifier received during access recovery using the auth.restore method.
// Use this parameter when changing the password immediately after restoring access.
func (r *AccountChangePasswordRequest) RestoreSID(restoreSID string) *AccountChangePasswordRequest {
	r.parameters.Set(constants.ParameterNameRestoreSID, restoreSID)
	return r
}

// ChangePasswordHash sets the hash obtained after a successful Auth authorization using the code received via SMS.
// Use this parameter when changing the password immediately after restoring access.
func (r *AccountChangePasswordRequest) ChangePasswordHash(changePasswordHash string) *AccountChangePasswordRequest {
	r.parameters.Set(constants.ParameterNameChangePasswordHash, changePasswordHash)
	return r
}

// OldPassword sets the user's current password.
func (r *AccountChangePasswordRequest) OldPassword(oldPassword string) *AccountChangePasswordRequest {
	r.parameters.Set(constants.ParameterNameOldPassword, oldPassword)
	return r
}

// NewPassword sets the new password that will be established as the current one.
// This parameter is mandatory, and the minimum length is 6 characters.
func (r *AccountChangePasswordRequest) NewPassword(newPassword string) *AccountChangePasswordRequest {
	r.parameters.Set(constants.ParameterNameNewPassword, newPassword)
	return r
}

// AccountGetActiveOffersRequest defines the request for account.getActiveOffers.
// Returns a list of active advertising offers (offers), by completing which the user will be able to receive the corresponding number of votes to his account within the application.
// Doc: https://dev.vk.com/method/account.getActiveOffers
type AccountGetActiveOffersRequest struct {
	BaseRequest
}

// NewAccountGetActiveOffersRequest creates a new request for account.getActiveOffers
func NewAccountGetActiveOffersRequest(a *api.API, actor actor.Actor) *AccountGetActiveOffersRequest {
	return &AccountGetActiveOffersRequest{*NewMethodBaseRequest(a, actor, "account.getActiveOffers")}
}

// Exec executes the request and unmarshals the response into AccountGetActiveOffersResponse
func (r *AccountGetActiveOffersRequest) Exec(ctx context.Context) (response response.AccountGetActiveOffersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Offset The bias required to select a specific subset of offers. Minimum is 0. By default 0
func (r *AccountGetActiveOffersRequest) Offset(offset uint) *AccountGetActiveOffersRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(int(offset)))
	return r
}

// Count The number of offers you need to receive. Maximum is 100. Minimum is 0. By default 100
func (r *AccountGetActiveOffersRequest) Count(count uint) *AccountGetActiveOffersRequest {
	r.parameters.Set(constants.ParameterNameCount, strconv.Itoa(int(count)))
	return r
}

// AccountGetAppPermissionsRequest defines the request for account.getAppPermissions.
// The method gets your application's user settings.
// Doc: https://dev.vk.com/method/account.getAppPermissions
type AccountGetAppPermissionsRequest struct {
	BaseRequest
}

// NewAccountGetAppPermissionsRequest creates a new request for account.getAppPermissions
func NewAccountGetAppPermissionsRequest(a *api.API, actor actor.Actor) *AccountGetAppPermissionsRequest {
	return &AccountGetAppPermissionsRequest{*NewMethodBaseRequest(a, actor, "account.getAppPermissions")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountGetAppPermissionsRequest) Exec(ctx context.Context) (response response.AccountGetAppPermissionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UserID Required parameter. ID of the user whose settings you want to retrieve. Default value: current user ID. Minimum is 1. Entity - owner
func (r *AccountGetAppPermissionsRequest) UserID(id uint) *AccountGetAppPermissionsRequest {
	r.parameters.Set(constants.ParameterNameUserID, strconv.Itoa(int(id)))
	return r
}

// AccountGetBannedRequest defines the request for account.getBanned.
// Returns a list of blacklisted users.
// Doc: https://dev.vk.com/method/account.getBanned
type AccountGetBannedRequest struct {
	BaseRequest
}

// NewAccountGetBannedRequest creates a new request for account.getBanned
func NewAccountGetBannedRequest(a *api.API, actor actor.Actor) *AccountGetBannedRequest {
	return &AccountGetBannedRequest{*NewMethodBaseRequest(a, actor, "account.getBanned")}
}

// Exec executes the request and unmarshals the response into AccountGetBannedResponse
func (r *AccountGetBannedRequest) Exec(ctx context.Context) (response response.AccountGetBannedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Offset The offset required to sample a specific subset of the blacklist.
func (r *AccountGetBannedRequest) Offset(offset uint) *AccountGetBannedRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(int(offset)))
	return r
}

// Count The number of objects about which information must be returned.
func (r *AccountGetBannedRequest) Count(count uint) *AccountGetBannedRequest {
	r.parameters.Set(constants.ParameterNameCount, strconv.Itoa(int(count)))
	return r
}

// Fields sets the list of additional profile fields to return for each user in the blacklist.
func (r *AccountGetBannedRequest) Fields(fields ...string) *AccountGetBannedRequest {
	r.parameters.Set(constants.ParameterNameFields, strings.Join(fields, ","))
	return r
}

// AccountGetCountersRequest defines the request for account.getCounters.
// The method returns non-zero values of the user counters.
// Doc: https://dev.vk.com/method/account.getCounters
type AccountGetCountersRequest struct {
	BaseRequest
}

// NewAccountGetCountersRequest creates a new request for account.getCounters
func NewAccountGetCountersRequest(a *api.API, actor actor.Actor) *AccountGetCountersRequest {
	return &AccountGetCountersRequest{*NewMethodBaseRequest(a, actor, "account.getCounters")}
}

// Exec executes the request and unmarshals the response into AccountGetCountersResponse
func (r *AccountGetCountersRequest) Exec(ctx context.Context) (response response.AccountGetCountersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Filter sets the counters to return information about, separated by commas.
// This optional parameter allows you to specify which types of counters should be included in the response.
// Possible values for the filter include:
// - 'friends': New friend requests.
// - 'friends_suggestions': Suggested friends.
// - 'messages': New messages.
// - 'photos': New tags on photos.
// - 'videos': New tags on videos.
// - 'gifts': Gifts received.
// - 'events': Events notifications.
// - 'groups': Community invitations and updates.
// - 'notifications': Replies and other notifications.
// - 'sdk': Requests from mobile games using the VK SDK.
// - 'app_requests': Notifications from applications.
// - 'friends_recommendations': Friend recommendations.
// The values should be provided as a comma-separated list.
func (r *AccountGetCountersRequest) Filter(filter ...string) *AccountGetCountersRequest {
	r.parameters.Set(constants.ParameterNameFilter, strings.Join(filter, ","))
	return r
}

// AccountGetInfoRequest defines the request for account.getInfo
// Returns information about the current account.
// Doc: https://dev.vk.com/method/account.getInfo
type AccountGetInfoRequest struct {
	BaseRequest
}

// NewAccountGetInfoRequest creates a new request for account.getInfo
func NewAccountGetInfoRequest(a *api.API, actor actor.Actor) *AccountGetInfoRequest {
	return &AccountGetInfoRequest{*NewMethodBaseRequest(a, actor, "account.getInfo")}
}

// Exec executes the request and unmarshals the response into AccountGetInfoResponse
func (r *AccountGetInfoRequest) Exec(ctx context.Context) (response response.AccountGetInfoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Fields sets the list of fields to return. The default behavior is to return all fields.
// Possible values for fields include:
// - 'country': The user's country.
// - 'https_required': Indicates if HTTPS is required for requests.
// - 'own_posts_default': Default setting for showing only own posts on the user's wall.
// - 'no_wall_replies': Indicates if comments on the wall are disabled.
// - 'intro': Bitmask indicating if the user has completed the introductory tutorial.
// - 'lang': The user's language settings.
func (r *AccountGetInfoRequest) Fields(fields ...string) *AccountGetInfoRequest {
	r.parameters.Set(constants.ParameterNameFields, strings.Join(fields, ","))
	return r
}

// AccountGetProfileInfoRequest defines the request for account.getProfileInfo
// Returns information about the current profile.
// Doc: https://dev.vk.com/method/account.getProfileInfo
type AccountGetProfileInfoRequest struct {
	BaseRequest
}

// NewAccountGetProfileInfoRequest creates a new request for account.getProfileInfo
func NewAccountGetProfileInfoRequest(a *api.API, actor actor.Actor) *AccountGetProfileInfoRequest {
	return &AccountGetProfileInfoRequest{*NewMethodBaseRequest(a, actor, "account.getProfileInfo")}
}

// Exec executes the request and unmarshals the response into AccountGetProfileInfoResponse
func (r *AccountGetProfileInfoRequest) Exec(ctx context.Context) (response response.AccountGetProfileInfoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountGetPushSettingsRequest defines the request for account.getPushSettings
// Allows you to receive Push Notification settings.
// One of the parameters must be passed: "token" or "device_id".
// Doc: https://dev.vk.com/method/account.getPushSettings
type AccountGetPushSettingsRequest struct {
	BaseRequest
}

// NewAccountGetPushSettingsRequest creates a new request for account.getPushSettings
func NewAccountGetPushSettingsRequest(a *api.API, actor actor.Actor) *AccountGetPushSettingsRequest {
	return &AccountGetPushSettingsRequest{*NewMethodBaseRequest(a, actor, "account.getPushSettings")}
}

// Exec executes the request and unmarshals the response into AccountGetPushSettingsResponse
func (r *AccountGetPushSettingsRequest) Exec(ctx context.Context) (response response.AccountGetPushSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Token Device ID used to send notifications. (for mpns the identifier must be the url for sending notifications)
func (r *AccountGetPushSettingsRequest) Token(token string) *AccountGetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameToken, token)
	return r
}

// DeviceID Unique device identifier.
func (r *AccountGetPushSettingsRequest) DeviceID(id string) *AccountGetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, id)
	return r
}

// AccountRegisterDeviceRequest defines the request for account.registerDevice
// Subscribes an iOS, Android, Windows Phone or Mac device to receive push notifications.
// Doc: https://dev.vk.com/method/account.registerDevice
type AccountRegisterDeviceRequest struct {
	BaseRequest
}

// NewAccountRegisterDeviceRequest creates a new request for account.registerDevice
func NewAccountRegisterDeviceRequest(a *api.API, actor actor.Actor) *AccountRegisterDeviceRequest {
	return &AccountRegisterDeviceRequest{*NewMethodBaseRequest(a, actor, "account.registerDevice")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountRegisterDeviceRequest) Exec(ctx context.Context) (response response.AccountRegisterDeviceResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Token sets the unique identifier of the device used for sending notifications.
// For MPNS (Microsoft Push Notification Service), this identifier should be a url for sending notifications.
// This parameter is mandatory.
func (r *AccountRegisterDeviceRequest) Token(token string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameToken, token)
	return r
}

// DeviceModel sets the string name of the device model.
func (r *AccountRegisterDeviceRequest) DeviceModel(deviceModel string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameDeviceModel, deviceModel)
	return r
}

// DeviceYear sets the year of the device.
func (r *AccountRegisterDeviceRequest) DeviceYear(deviceYear int) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameDeviceYear, strconv.Itoa(deviceYear))
	return r
}

// DeviceID sets the unique identifier of the device.
// This parameter is mandatory.
func (r *AccountRegisterDeviceRequest) DeviceID(deviceID string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, deviceID)
	return r
}

// SystemVersion sets the string version of the device's operating system.
func (r *AccountRegisterDeviceRequest) SystemVersion(systemVersion string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameSystemVersion, systemVersion)
	return r
}

// NoText determines whether to include the text of the message in push notifications.
// 1 — do not include the message text in push notifications.
// 0 — (default) include the message text.
func (r *AccountRegisterDeviceRequest) NoText(noText bool) *AccountRegisterDeviceRequest {
	if noText {
		r.parameters.Set(constants.ParameterNameNoText, "1")
	} else {
		r.parameters.Set(constants.ParameterNameNoText, "0")
	}
	return r
}

// Subscribe sets the list of notification types to send.
// Possible values are: msg, friend, call, reply, mention, group, like.
// Defaults to: msg, friend.
func (r *AccountRegisterDeviceRequest) Subscribe(subscribe ...string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameSubscribe, strings.Join(subscribe, ","))
	return r
}

// Settings sets a serialized JSON objects that describes notification settings in a special format.
// If not provided, current settings will be applied. For new device IDs, the latest settings for the given device_id will be used.
// https://dev.vk.com/ru/reference/objects/push-settings
func (r *AccountRegisterDeviceRequest) Settings(settings string) *AccountRegisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameSettings, settings)
	return r
}

// Sandbox is a flag intended for iOS devices.
// 1 — use the sandbox server for sending push notifications.
// 0 — do not use the sandbox server.
func (r *AccountRegisterDeviceRequest) Sandbox(sandbox bool) *AccountRegisterDeviceRequest {
	if sandbox {
		r.parameters.Set(constants.ParameterNameSandbox, "1")
	} else {
		r.parameters.Set(constants.ParameterNameSandbox, "0")
	}
	return r
}

// AccountSaveProfileInfoRequest defines the request for account.saveProfileInfo
// Edits the current profile information.
// Doc: https://dev.vk.com/method/account.saveProfileInfo
type AccountSaveProfileInfoRequest struct {
	BaseRequest
}

// NewAccountSaveProfileInfoRequest creates a new request for account.saveProfileInfo
func NewAccountSaveProfileInfoRequest(a *api.API, actor actor.Actor) *AccountSaveProfileInfoRequest {
	return &AccountSaveProfileInfoRequest{*NewMethodBaseRequest(a, actor, "account.saveProfileInfo")}
}

// Exec executes the request and unmarshals the response into AccountSaveProfileInfoResponse
func (r *AccountSaveProfileInfoRequest) Exec(ctx context.Context) (response response.AccountSaveProfileInfoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// FirstName sets the first name of the user. The name must start with a capital letter.
func (r *AccountSaveProfileInfoRequest) FirstName(firstName string) *AccountSaveProfileInfoRequest {
	// Ensure the first letter is uppercase
	if len(firstName) > 0 {
		firstName = string(firstName[0]-32) + firstName[1:]
	}
	r.parameters.Set(constants.ParameterNameFirstName, firstName)
	return r
}

// LastName sets the last name of the user. The last name must start with a capital letter.
func (r *AccountSaveProfileInfoRequest) LastName(lastName string) *AccountSaveProfileInfoRequest {
	// Ensure the first letter is uppercase
	if len(lastName) > 0 {
		lastName = string(lastName[0]-32) + lastName[1:]
	}
	r.parameters.Set(constants.ParameterNameLastName, lastName)
	return r
}

// MaidenName sets the maiden name of the user. This field is relevant only for female users.
func (r *AccountSaveProfileInfoRequest) MaidenName(maidenName string) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameMaidenName, maidenName)
	return r
}

// ScreenName sets the short name of the user's page.
func (r *AccountSaveProfileInfoRequest) ScreenName(screenName string) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameScreenName, screenName)
	return r
}

// CancelRequestID sets the ID of the name change request that should be canceled. If this parameter is passed, all other parameters are ignored.
func (r *AccountSaveProfileInfoRequest) CancelRequestID(id uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameCancelRequestID, strconv.Itoa(int(id)))
	return r
}

// Sex sets the gender of the user. Possible values:
// 1 — female;
// 2 — male.
func (r *AccountSaveProfileInfoRequest) Sex(sex uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameSex, strconv.Itoa(int(sex)))
	return r
}

// Relation sets the relationship status of the user. Possible values:
// 1 — single;
// 2 — in a relationship;
// 3 — engaged;
// 4 — married;
// 5 — it's complicated;
// 6 — actively searching;
// 7 — in love;
// 8 — in a civil union;
// 0 — not specified.
func (r *AccountSaveProfileInfoRequest) Relation(relation uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameRelation, strconv.Itoa(int(relation)))
	return r
}

// RelationPartnerID sets the ID of the user with whom the relationship status is associated.
func (r *AccountSaveProfileInfoRequest) RelationPartnerID(id int) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameRelationPartnerID, strconv.Itoa(id))
	return r
}

// BDate sets the user's birth date in the format DD.MM.YYYY, e.g., 15.11.1984.
func (r *AccountSaveProfileInfoRequest) BDate(bdate string) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameBDate, bdate)
	return r
}

// BDateVisibility sets the visibility of the user's birth date. Possible values:
// 1 — show birth date;
// 2 — show only month and day;
// 0 — do not show birth date.
func (r *AccountSaveProfileInfoRequest) BDateVisibility(visibility uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameBDateVisibility, strconv.Itoa(int(visibility)))
	return r
}

// HomeTown sets the user's hometown.
func (r *AccountSaveProfileInfoRequest) HomeTown(homeTown string) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameHomeTown, homeTown)
	return r
}

// CountryID sets the ID of the user's country.
func (r *AccountSaveProfileInfoRequest) CountryID(countryID uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameCountryID, strconv.Itoa(int(countryID)))
	return r
}

// CityID sets the ID of the user's city.
func (r *AccountSaveProfileInfoRequest) CityID(cityID uint) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameCityID, strconv.Itoa(int(cityID)))
	return r
}

// Status sets the status message of the user, which can also be changed using the status.set method.
func (r *AccountSaveProfileInfoRequest) Status(status string) *AccountSaveProfileInfoRequest {
	r.parameters.Set(constants.ParameterNameStatus, status)
	return r
}

// AccountSetInfoRequest defines the request for account.setInfo
// Allows you to edit information about the current account.
// Doc: https://dev.vk.com/method/account.setInfo
type AccountSetInfoRequest struct {
	BaseRequest
}

// NewAccountSetInfoRequest creates a new request for account.setInfo
func NewAccountSetInfoRequest(a *api.API, actor actor.Actor) *AccountSetInfoRequest {
	return &AccountSetInfoRequest{*NewMethodBaseRequest(a, actor, "account.setInfo")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountSetInfoRequest) Exec(ctx context.Context) (response response.AccountSetInfoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Intro sets the bitmask responsible for completing the onboarding process in mobile clients.
func (r *AccountSetInfoRequest) Intro(intro uint) *AccountSetInfoRequest {
	r.parameters.Set(constants.ParameterNameIntro, strconv.Itoa(int(intro)))
	return r
}

// OwnPostsDefault sets the default behavior for displaying posts on the user's wall.
// 1 – only the user's own posts should be displayed by default;
// 0 – all posts should be displayed.
func (r *AccountSetInfoRequest) OwnPostsDefault(ownPostsDefault bool) *AccountSetInfoRequest {
	if ownPostsDefault {
		r.parameters.Set(constants.ParameterNameOwnPostsDefault, "1")
	} else {
		r.parameters.Set(constants.ParameterNameOwnPostsDefault, "0")
	}
	return r
}

// NoWallReplies sets the comment policy for the user's wall posts.
// 1 – disable commenting on wall posts;
// 0 – allow commenting.
func (r *AccountSetInfoRequest) NoWallReplies(noWallReplies bool) *AccountSetInfoRequest {
	if noWallReplies {
		r.parameters.Set(constants.ParameterNameNoWallReplies, "1")
	} else {
		r.parameters.Set(constants.ParameterNameNoWallReplies, "0")
	}
	return r
}

// Name sets the name of the setting.
func (r *AccountSetInfoRequest) Name(name string) *AccountSetInfoRequest {
	r.parameters.Set(constants.ParameterNameName, name)
	return r
}

// Value sets the value of the setting.
func (r *AccountSetInfoRequest) Value(value string) *AccountSetInfoRequest {
	r.parameters.Set(constants.ParameterNameValue, value)
	return r
}

// AccountSetOfflineRequest defines the request for account.setOffline
// Marks the current user as offline (only in the current application).
// Doc: https://dev.vk.com/method/account.setOffline
type AccountSetOfflineRequest struct {
	BaseRequest
}

// NewAccountSetOfflineRequest creates a new request for account.setOffline
func NewAccountSetOfflineRequest(a *api.API, actor actor.Actor) *AccountSetOfflineRequest {
	return &AccountSetOfflineRequest{*NewMethodBaseRequest(a, actor, "account.setOffline")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountSetOfflineRequest) Exec(ctx context.Context) (response response.AccountSetOfflineResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountSetOnlineRequest defines the request for account.setOnline
// Marks the current user as online for 5 minutes.
// Doc: https://dev.vk.com/method/account.setOnline
type AccountSetOnlineRequest struct {
	BaseRequest
}

// NewAccountSetOnlineRequest creates a new request for account.setOnline
func NewAccountSetOnlineRequest(a *api.API, actor actor.Actor) *AccountSetOnlineRequest {
	return &AccountSetOnlineRequest{*NewMethodBaseRequest(a, actor, "account.setOnline")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountSetOnlineRequest) Exec(ctx context.Context) (response response.AccountSetOnlineResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// VoIP Are video calls possible for this device?
func (r *AccountSetOnlineRequest) VoIP(value bool) *AccountSetOnlineRequest {
	if value {
		r.parameters.Set(constants.ParameterNameVoIP, "1")
	} else {
		r.parameters.Set(constants.ParameterNameVoIP, "0")
	}
	return r
}

// AccountSetPushSettingsRequest defines the request for account.setPushSettings
// Changes the push notification setting.
// Doc: https://dev.vk.com/method/account.setPushSettings
type AccountSetPushSettingsRequest struct {
	BaseRequest
}

// NewAccountSetPushSettingsRequest creates a new request for account.setPushSettings
func NewAccountSetPushSettingsRequest(a *api.API, actor actor.Actor) *AccountSetPushSettingsRequest {
	return &AccountSetPushSettingsRequest{*NewMethodBaseRequest(a, actor, "account.setPushSettings")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountSetPushSettingsRequest) Exec(ctx context.Context) (response response.AccountSetPushSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// DeviceID Required parameter. Sets the unique device identifier.
func (r *AccountSetPushSettingsRequest) DeviceID(deviceID string) *AccountSetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, deviceID)
	return r
}

// Settings sets the serialized JSON objects describing the notification settings https://dev.vk.com/ru/reference/objects/push-settings. Example: request.Settings(`{"msg": {"sound": "default", "vibrate": 1}}`)
func (r *AccountSetPushSettingsRequest) Settings(settings string) *AccountSetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameSettings, settings)
	return r
}

// Key sets the notification key.
func (r *AccountSetPushSettingsRequest) Key(key string) *AccountSetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameKey, key)
	return r
}

// Value sets the new value for the notification in a special format https://dev.vk.com/ru/reference/objects/push-settings. Example: request.Value("msg")
func (r *AccountSetPushSettingsRequest) Value(value string) *AccountSetPushSettingsRequest {
	r.parameters.Set(constants.ParameterNameValue, value)
	return r
}

// AccountSetSilenceModeRequest defines the request for account.setSilenceMode
// Disables push notifications for a specified period of time.
// Doc: https://dev.vk.com/method/account.setSilenceMode
type AccountSetSilenceModeRequest struct {
	BaseRequest
}

// NewAccountSetSilenceModeRequest creates a new request for account.setSilenceMode
func NewAccountSetSilenceModeRequest(a *api.API, actor actor.Actor) *AccountSetSilenceModeRequest {
	return &AccountSetSilenceModeRequest{*NewMethodBaseRequest(a, actor, "account.setSilenceMode")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountSetSilenceModeRequest) Exec(ctx context.Context) (response response.AccountSetSilenceModeResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Token sets the device identifier for the push notification service.
func (r *AccountSetSilenceModeRequest) Token(token string) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNameToken, token)
	return r
}

// DeviceID sets the unique device identifier.
func (r *AccountSetSilenceModeRequest) DeviceID(deviceID string) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, deviceID)
	return r
}

// Time sets the duration in seconds for which notifications should be muted.
// Use -1 to disable notifications permanently.
func (r *AccountSetSilenceModeRequest) Time(time int) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNameTime, strconv.Itoa(time))
	return r
}

// ChatID sets the identifier of the chat for which notifications should be muted.
func (r *AccountSetSilenceModeRequest) ChatID(chatID int) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNameChatID, strconv.Itoa(chatID))
	return r
}

// UserID sets the identifier of the user for whom notifications should be muted.
func (r *AccountSetSilenceModeRequest) UserID(userID int) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNameUserID, strconv.Itoa(userID))
	return r
}

// PeerID sets the destination identifier.
// For a user: user ID.
// For a group chat: 2000000000 + chat ID.
// For a community: -community ID.
func (r *AccountSetSilenceModeRequest) PeerID(peerID int) *AccountSetSilenceModeRequest {
	r.parameters.Set(constants.ParameterNamePeerID, strconv.Itoa(peerID))
	return r
}

// Sound sets the sound settings for this conversation.
// Use 1 to enable sound, 0 to disable it (works only if peer_id is a group chat or user ID).
func (r *AccountSetSilenceModeRequest) Sound(sound bool) *AccountSetSilenceModeRequest {
	if sound {
		r.parameters.Set(constants.ParameterNameSound, "1")
	} else {
		r.parameters.Set(constants.ParameterNameSound, "0")
	}
	return r
}

// AccountUnbanRequest defines the request for account.unban
// Removes a user or group from the blacklist.
// Doc: https://dev.vk.com/method/account.unban
type AccountUnbanRequest struct {
	BaseRequest
}

// NewAccountUnbanRequest creates a new request for account.unban
func NewAccountUnbanRequest(a *api.API, actor actor.Actor) *AccountUnbanRequest {
	return &AccountUnbanRequest{*NewMethodBaseRequest(a, actor, "account.unban")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountUnbanRequest) Exec(ctx context.Context) (response response.AccountUnbanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// OwnerID The ID of the user or group to remove from the blacklist.
func (r *AccountUnbanRequest) OwnerID(ownerID int) *AccountUnbanRequest {
	r.parameters.Set(constants.ParameterNameOwnerID, strconv.Itoa(ownerID))
	return r
}

// AccountUnregisterDeviceRequest defines the request for account.unregisterDevice
// Unsubscribes the device from Push notifications.
// Doc: https://dev.vk.com/method/account.unregisterDevice
type AccountUnregisterDeviceRequest struct {
	BaseRequest
}

// NewAccountUnregisterDeviceRequest creates a new request for account.unregisterDevice
func NewAccountUnregisterDeviceRequest(a *api.API, actor actor.Actor) *AccountUnregisterDeviceRequest {
	return &AccountUnregisterDeviceRequest{*NewMethodBaseRequest(a, actor, "account.unregisterDevice")}
}

// Exec executes the request and unmarshals the response into an int
func (r *AccountUnregisterDeviceRequest) Exec(ctx context.Context) (response response.AccountUnregisterDeviceResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Token sets the device identifier used for sending notifications.
func (r *AccountUnregisterDeviceRequest) Token(token string) *AccountUnregisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameToken, token)
	return r
}

// DeviceID sets the unique device identifier.
func (r *AccountUnregisterDeviceRequest) DeviceID(deviceID string) *AccountUnregisterDeviceRequest {
	r.parameters.Set(constants.ParameterNameDeviceID, deviceID)
	return r
}

// Sandbox sets the flag for iOS devices.
// Use 1 to unregister the device using the sandbox server for sending push notifications.
// Use 0 to unregister the device that does not use the sandbox server.
func (r *AccountUnregisterDeviceRequest) Sandbox(sandbox bool) *AccountUnregisterDeviceRequest {
	if sandbox {
		r.parameters.Set(constants.ParameterNameSandbox, "1")
	} else {
		r.parameters.Set(constants.ParameterNameSandbox, "0")
	}
	return r
}
