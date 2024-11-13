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

// Doc: https://dev.vk.com/method/groups

// GroupsAddAddressRequest defines the request for groups.addAddress
//
// The method allows adding an address to a community. The list of addresses can be retrieved with the method groups.getAddresses.
// You must be an admin of the community to use this method.
// Doc: https://dev.vk.com/method/groups.addAddress
type GroupsAddAddressRequest struct {
	BaseRequest
}

// NewGroupsAddAddressRequest creates a new request for groups.addAddress
func NewGroupsAddAddressRequest(a *api.API, actor actor.Actor) *GroupsAddAddressRequest {
	return &GroupsAddAddressRequest{*NewMethodBaseRequest(a, actor, "groups.addAddress")}
}

// Exec executes the request and unmarshals the response into GroupsAddAddressResponse
func (r *GroupsAddAddressRequest) Exec(ctx context.Context) (response response.GroupsAddAddressResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsAddCallbackServerRequest defines the request for groups.addCallbackServer
//
// The method adds a server for Callback api in the community.
// Doc: https://dev.vk.com/method/groups.addCallbackServer
type GroupsAddCallbackServerRequest struct {
	BaseRequest
}

// NewGroupsAddCallbackServerRequest creates a new request for groups.addCallbackServer
func NewGroupsAddCallbackServerRequest(a *api.API, actor actor.Actor) *GroupsAddCallbackServerRequest {
	return &GroupsAddCallbackServerRequest{*NewMethodBaseRequest(a, actor, "groups.addCallbackServer")}
}

// Exec executes the request and unmarshals the response into GroupsAddCallbackServerResponse
func (r *GroupsAddCallbackServerRequest) Exec(ctx context.Context) (response response.GroupsAddCallbackServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupID required
func (r *GroupsAddCallbackServerRequest) GroupID(id int) *GroupsAddCallbackServerRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// URL required
func (r *GroupsAddCallbackServerRequest) URL(url string) *GroupsAddCallbackServerRequest {
	r.parameters.Set(constants.ParameterNameURL, url)
	return r
}

// Title required. length < 15
func (r *GroupsAddCallbackServerRequest) Title(title string) *GroupsAddCallbackServerRequest {
	if len(title) < 15 {
		r.parameters.Set(constants.ParameterNameTitle, title)
	}
	return r
}

// SecretKey required. length < 50
func (r *GroupsAddCallbackServerRequest) SecretKey(key string) *GroupsAddCallbackServerRequest {
	if len(key) < 50 {
		r.parameters.Set(constants.ParameterNameSecretKey, key)
	}
	return r
}

// GroupsAddLinkRequest defines the request for groups.addLink
//
// The method allows adding links to the community.
// Doc: https://dev.vk.com/method/groups.addLink
type GroupsAddLinkRequest struct {
	BaseRequest
}

// NewGroupsAddLinkRequest creates a new request for groups.addLink
func NewGroupsAddLinkRequest(a *api.API, actor actor.Actor) *GroupsAddLinkRequest {
	return &GroupsAddLinkRequest{*NewMethodBaseRequest(a, actor, "groups.addLink")}
}

// Exec executes the request and unmarshals the response into GroupsAddLinkResponse
func (r *GroupsAddLinkRequest) Exec(ctx context.Context) (response response.GroupsAddLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsApproveRequest defines the request for groups.approveRequest
//
// The method allows approving a user's request to join the community.
// Doc: https://dev.vk.com/method/groups.approveRequest
type GroupsApproveRequest struct {
	BaseRequest
}

// NewGroupsApproveRequest creates a new request for groups.approveRequest
func NewGroupsApproveRequest(a *api.API, actor actor.Actor) *GroupsApproveRequest {
	return &GroupsApproveRequest{*NewMethodBaseRequest(a, actor, "groups.approveRequest")}
}

// Exec executes the request and unmarshals the response into GroupsApproveRequestResponse
func (r *GroupsApproveRequest) Exec(ctx context.Context) (response response.GroupsApproveRequestResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsBanRequest defines the request for groups.ban
//
// The method adds a user or group to the community's blacklist.
// Doc: https://dev.vk.com/method/groups.ban
type GroupsBanRequest struct {
	BaseRequest
}

// NewGroupsBanRequest creates a new request for groups.ban
func NewGroupsBanRequest(a *api.API, actor actor.Actor) *GroupsBanRequest {
	return &GroupsBanRequest{*NewMethodBaseRequest(a, actor, "groups.ban")}
}

// Exec executes the request and unmarshals the response into GroupsBanResponse
func (r *GroupsBanRequest) Exec(ctx context.Context) (response response.GroupsBanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsCreateRequest defines the request for groups.create
//
// The method creates a new community.
// Doc: https://dev.vk.com/method/groups.create
type GroupsCreateRequest struct {
	BaseRequest
}

// NewGroupsCreateRequest creates a new request for groups.create
func NewGroupsCreateRequest(a *api.API, actor actor.Actor) *GroupsCreateRequest {
	return &GroupsCreateRequest{*NewMethodBaseRequest(a, actor, "groups.create")}
}

// Exec executes the request and unmarshals the response into GroupsCreateResponse
func (r *GroupsCreateRequest) Exec(ctx context.Context) (response response.GroupsCreateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsDeleteAddressRequest defines the request for groups.deleteAddress
//
// The method deletes an address from the community.
// Doc: https://dev.vk.com/method/groups.deleteAddress
type GroupsDeleteAddressRequest struct {
	BaseRequest
}

// NewGroupsDeleteAddressRequest creates a new request for groups.deleteAddress
func NewGroupsDeleteAddressRequest(a *api.API, actor actor.Actor) *GroupsDeleteAddressRequest {
	return &GroupsDeleteAddressRequest{*NewMethodBaseRequest(a, actor, "groups.deleteAddress")}
}

// Exec executes the request and unmarshals the response into GroupsDeleteAddressResponse
func (r *GroupsDeleteAddressRequest) Exec(ctx context.Context) (response response.GroupsDeleteAddressResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsDeleteCallbackServerRequest defines the request for groups.deleteCallbackServer
//
// The method deletes a Callback api server from the community.
// Doc: https://dev.vk.com/method/groups.deleteCallbackServer
type GroupsDeleteCallbackServerRequest struct {
	BaseRequest
}

// NewGroupsDeleteCallbackServerRequest creates a new request for groups.deleteCallbackServer
func NewGroupsDeleteCallbackServerRequest(a *api.API, actor actor.Actor) *GroupsDeleteCallbackServerRequest {
	return &GroupsDeleteCallbackServerRequest{*NewMethodBaseRequest(a, actor, "groups.deleteCallbackServer")}
}

// Exec executes the request and unmarshals the response into GroupsDeleteCallbackServerResponse
func (r *GroupsDeleteCallbackServerRequest) Exec(ctx context.Context) (response response.GroupsDeleteCallbackServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupID required
func (r *GroupsDeleteCallbackServerRequest) GroupID(id int) *GroupsDeleteCallbackServerRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// ServerID Server ID to delete. Required
func (r *GroupsDeleteCallbackServerRequest) ServerID(id int) *GroupsDeleteCallbackServerRequest {
	r.parameters.Set(constants.ParameterNameServerID, strconv.Itoa(id))
	return r
}

// GroupsDeleteLinkRequest defines the request for groups.deleteLink
//
// The method allows deleting links from the community.
// Doc: https://dev.vk.com/method/groups.deleteLink
type GroupsDeleteLinkRequest struct {
	BaseRequest
}

// NewGroupsDeleteLinkRequest creates a new request for groups.deleteLink
func NewGroupsDeleteLinkRequest(a *api.API, actor actor.Actor) *GroupsDeleteLinkRequest {
	return &GroupsDeleteLinkRequest{*NewMethodBaseRequest(a, actor, "groups.deleteLink")}
}

// Exec executes the request and unmarshals the response into GroupsDeleteLinkResponse
func (r *GroupsDeleteLinkRequest) Exec(ctx context.Context) (response response.GroupsDeleteLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsDisableOnlineRequest defines the request for groups.disableOnline
//
// The method disables the "online" status in the community.
// Doc: https://dev.vk.com/method/groups.disableOnline
type GroupsDisableOnlineRequest struct {
	BaseRequest
}

// NewGroupsDisableOnlineRequest creates a new request for groups.disableOnline
func NewGroupsDisableOnlineRequest(a *api.API, actor actor.Actor) *GroupsDisableOnlineRequest {
	return &GroupsDisableOnlineRequest{*NewMethodBaseRequest(a, actor, "groups.disableOnline")}
}

// Exec executes the request and unmarshals the response into GroupsDisableOnlineResponse
func (r *GroupsDisableOnlineRequest) Exec(ctx context.Context) (response response.GroupsDisableOnlineResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEditRequest defines the request for groups.edit
//
// The method allows editing community settings.
// Doc: https://dev.vk.com/method/groups.edit
type GroupsEditRequest struct {
	BaseRequest
}

// NewGroupsEditRequest creates a new request for groups.edit
func NewGroupsEditRequest(a *api.API, actor actor.Actor) *GroupsEditRequest {
	return &GroupsEditRequest{*NewMethodBaseRequest(a, actor, "groups.edit")}
}

// Exec executes the request and unmarshals the response into GroupsEditResponse
func (r *GroupsEditRequest) Exec(ctx context.Context) (response response.GroupsEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEditAddressRequest defines the request for groups.editAddress
//
// The method allows editing an address in the community. Use groups.getAddresses to retrieve the list of addresses.
// Doc: https://dev.vk.com/method/groups.editAddress
type GroupsEditAddressRequest struct {
	BaseRequest
}

// NewGroupsEditAddressRequest creates a new request for groups.editAddress
func NewGroupsEditAddressRequest(a *api.API, actor actor.Actor) *GroupsEditAddressRequest {
	return &GroupsEditAddressRequest{*NewMethodBaseRequest(a, actor, "groups.editAddress")}
}

// Exec executes the request and unmarshals the response into GroupsEditAddressResponse
func (r *GroupsEditAddressRequest) Exec(ctx context.Context) (response response.GroupsEditAddressResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEditCallbackServerRequest defines the request for groups.editCallbackServer
//
// The method edits a Callback api server in the community.
// Doc: https://dev.vk.com/method/groups.editCallbackServer
type GroupsEditCallbackServerRequest struct {
	BaseRequest
}

// NewGroupsEditCallbackServerRequest creates a new request for groups.editCallbackServer
func NewGroupsEditCallbackServerRequest(a *api.API, actor actor.Actor) *GroupsEditCallbackServerRequest {
	return &GroupsEditCallbackServerRequest{*NewMethodBaseRequest(a, actor, "groups.editCallbackServer")}
}

// Exec executes the request and unmarshals the response into GroupsEditCallbackServerResponse
func (r *GroupsEditCallbackServerRequest) Exec(ctx context.Context) (response response.GroupsEditCallbackServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEditLinkRequest defines the request for groups.editLink
//
// The method allows editing links in the community.
// Doc: https://dev.vk.com/method/groups.editLink
type GroupsEditLinkRequest struct {
	BaseRequest
}

// NewGroupsEditLinkRequest creates a new request for groups.editLink
func NewGroupsEditLinkRequest(a *api.API, actor actor.Actor) *GroupsEditLinkRequest {
	return &GroupsEditLinkRequest{*NewMethodBaseRequest(a, actor, "groups.editLink")}
}

// Exec executes the request and unmarshals the response into GroupsEditLinkResponse
func (r *GroupsEditLinkRequest) Exec(ctx context.Context) (response response.GroupsEditLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEditManagerRequest defines the request for groups.editManager
//
// The method allows assigning or removing a manager in the community or changing their role.
// Doc: https://dev.vk.com/method/groups.editManager
type GroupsEditManagerRequest struct {
	BaseRequest
}

// NewGroupsEditManagerRequest creates a new request for groups.editManager
func NewGroupsEditManagerRequest(a *api.API, actor actor.Actor) *GroupsEditManagerRequest {
	return &GroupsEditManagerRequest{*NewMethodBaseRequest(a, actor, "groups.editManager")}
}

// Exec executes the request and unmarshals the response into GroupsEditManagerResponse
func (r *GroupsEditManagerRequest) Exec(ctx context.Context) (response response.GroupsEditManagerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsEnableOnlineRequest defines the request for groups.enableOnline
//
// The method enables the "online" status in the community.
// Doc: https://dev.vk.com/method/groups.enableOnline
type GroupsEnableOnlineRequest struct {
	BaseRequest
}

// NewGroupsEnableOnlineRequest creates a new request for groups.enableOnline
func NewGroupsEnableOnlineRequest(a *api.API, actor actor.Actor) *GroupsEnableOnlineRequest {
	return &GroupsEnableOnlineRequest{*NewMethodBaseRequest(a, actor, "groups.enableOnline")}
}

// Exec executes the request and unmarshals the response into GroupsEnableOnlineResponse
func (r *GroupsEnableOnlineRequest) Exec(ctx context.Context) (response response.GroupsEnableOnlineResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetRequest defines the request for groups.get
//
// The method returns a list of communities for the specified user.
// Doc: https://dev.vk.com/method/groups.get
type GroupsGetRequest struct {
	BaseRequest
}

// NewGroupsGetRequest creates a new request for groups.get
func NewGroupsGetRequest(a *api.API, actor actor.Actor) *GroupsGetRequest {
	return &GroupsGetRequest{*NewMethodBaseRequest(a, actor, "groups.get")}
}

// Exec executes the request and unmarshals the response into GroupsGetResponse
func (r *GroupsGetRequest) Exec(ctx context.Context) (response response.GroupsGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetExtendedRequest defines the request for groups.get
//
// The method returns a list of communities for the specified user.
// Doc: https://dev.vk.com/method/groups.get
type GroupsGetExtendedRequest struct {
	BaseRequest
}

// NewGroupsGetExtendedRequest creates a new request for groups.get
func NewGroupsGetExtendedRequest(a *api.API, actor actor.Actor) *GroupsGetExtendedRequest {
	r := &GroupsGetExtendedRequest{*NewMethodBaseRequest(a, actor, "groups.get")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into GroupsGetExtendedResponse
func (r *GroupsGetExtendedRequest) Exec(ctx context.Context) (response response.GroupsGetExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetAddressesRequest defines the request for groups.getAddresses
//
// The method returns the addresses of a specified community.
// Doc: https://dev.vk.com/method/groups.getAddresses
type GroupsGetAddressesRequest struct {
	BaseRequest
}

// NewGroupsGetAddressesRequest creates a new request for groups.getAddresses
func NewGroupsGetAddressesRequest(a *api.API, actor actor.Actor) *GroupsGetAddressesRequest {
	return &GroupsGetAddressesRequest{*NewMethodBaseRequest(a, actor, "groups.getAddresses")}
}

// Exec executes the request and unmarshals the response into GroupsGetAddressesResponse
func (r *GroupsGetAddressesRequest) Exec(ctx context.Context) (response response.GroupsGetAddressesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetBannedRequest defines the request for groups.getBanned
//
// The method returns a list of banned users and communities in the community.
// Doc: https://dev.vk.com/method/groups.getBanned
type GroupsGetBannedRequest struct {
	BaseRequest
}

// NewGroupsGetBannedRequest creates a new request for groups.getBanned
func NewGroupsGetBannedRequest(a *api.API, actor actor.Actor) *GroupsGetBannedRequest {
	return &GroupsGetBannedRequest{*NewMethodBaseRequest(a, actor, "groups.getBanned")}
}

// Exec executes the request and unmarshals the response into GroupsGetBannedResponse
func (r *GroupsGetBannedRequest) Exec(ctx context.Context) (response response.GroupsGetBannedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetByIDRequest defines the request for groups.getById
//
// The method returns information about a specified community or several communities.
// Doc: https://dev.vk.com/method/groups.getById
type GroupsGetByIDRequest struct {
	BaseRequest
}

// NewGroupsGetByIDRequest creates a new request for groups.getById
func NewGroupsGetByIDRequest(a *api.API, actor actor.Actor) *GroupsGetByIDRequest {
	return &GroupsGetByIDRequest{*NewMethodBaseRequest(a, actor, "groups.getById")}
}

// Exec executes the request and unmarshals the response into GroupsGetByIdResponse
func (r *GroupsGetByIDRequest) Exec(ctx context.Context) (response response.GroupsGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *GroupsGetByIDRequest) GroupID(id string) *GroupsGetByIDRequest {
	r.parameters.Set(constants.ParameterNameGroupID, id)
	return r
}

// GroupIDs quantity ids < 501
func (r *GroupsGetByIDRequest) GroupIDs(ids string) *GroupsGetByIDRequest {
	r.parameters.Set(constants.ParameterNameGroupIDs, ids)
	return r
}

// GroupIDsArr quantity ids < 501
func (r *GroupsGetByIDRequest) GroupIDsArr(ids []string) *GroupsGetByIDRequest {
	if len(ids) < 501 {
		r.parameters.Set(constants.ParameterNameGroupIDs, strings.Join(ids, ","))
	}
	return r
}

// Fields activity, ban_info, can_post, can_see_all_posts, city, contacts, counters,
// country, cover, description, finish_date, fixed_post, links, market, members_count, place,
// site, start_date, status, verified, wiki_page,
func (r *GroupsGetByIDRequest) Fields(fields string) *GroupsGetByIDRequest {
	r.parameters.Set(constants.ParameterNameFields, fields)
	return r
}

// FieldsArr activity, ban_info, can_post, can_see_all_posts, city, contacts, counters,
// country, cover, description, finish_date, fixed_post, links, market, members_count, place,
// site, start_date, status, verified, wiki_page,
func (r *GroupsGetByIDRequest) FieldsArr(fields []string) *GroupsGetByIDRequest {
	r.parameters.Set(constants.ParameterNameFields, strings.Join(fields, ","))
	return r
}

// GroupsGetCallbackConfirmationCodeRequest defines the request for groups.getCallbackConfirmationCode
//
// The method returns a string needed to confirm the Callback api server address.
// Doc: https://dev.vk.com/method/groups.getCallbackConfirmationCode
type GroupsGetCallbackConfirmationCodeRequest struct {
	BaseRequest
}

// NewGroupsGetCallbackConfirmationCodeRequest creates a new request for groups.getCallbackConfirmationCode
func NewGroupsGetCallbackConfirmationCodeRequest(a *api.API, actor actor.Actor) *GroupsGetCallbackConfirmationCodeRequest {
	return &GroupsGetCallbackConfirmationCodeRequest{*NewMethodBaseRequest(a, actor, "groups.getCallbackConfirmationCode")}
}

// Exec executes the request and unmarshals the response into GroupsGetCallbackConfirmationCodeResponse
func (r *GroupsGetCallbackConfirmationCodeRequest) Exec(ctx context.Context) (response response.GroupsGetCallbackConfirmationCodeResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *GroupsGetCallbackConfirmationCodeRequest) GroupID(id int) *GroupsGetCallbackConfirmationCodeRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// GroupsGetCallbackServersRequest defines the request for groups.getCallbackServers
//
// The method retrieves information about Callback api servers in the community.
// Doc: https://dev.vk.com/method/groups.getCallbackServers
type GroupsGetCallbackServersRequest struct {
	BaseRequest
}

// NewGroupsGetCallbackServersRequest creates a new request for groups.getCallbackServers
func NewGroupsGetCallbackServersRequest(a *api.API, actor actor.Actor) *GroupsGetCallbackServersRequest {
	return &GroupsGetCallbackServersRequest{*NewMethodBaseRequest(a, actor, "groups.getCallbackServers")}
}

// Exec executes the request and unmarshals the response into GroupsGetCallbackServersResponse
func (r *GroupsGetCallbackServersRequest) Exec(ctx context.Context) (response response.GroupsGetCallbackServersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupID required
func (r *GroupsGetCallbackServersRequest) GroupID(id int) *GroupsGetCallbackServersRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// ServerIDs Identifiers of the servers about which you want to obtain data. By default, all servers are returned.
func (r *GroupsGetCallbackServersRequest) ServerIDs(ids string) *GroupsGetCallbackServersRequest {
	r.parameters.Set(constants.ParameterNameServerIDs, ids)
	return r
}

// ServerIDsArr Identifiers of the servers about which you want to obtain data. By default, all servers are returned.
func (r *GroupsGetCallbackServersRequest) ServerIDsArr(ids []string) *GroupsGetCallbackServersRequest {
	r.parameters.Set(constants.ParameterNameServerIDs, strings.Join(ids, ","))
	return r
}

// GroupsGetCallbackSettingsRequest defines the request for groups.getCallbackSettings
//
// The method retrieves the notification settings for the Callback api in the community.
// Doc: https://dev.vk.com/method/groups.getCallbackSettings
type GroupsGetCallbackSettingsRequest struct {
	BaseRequest
}

// NewGroupsGetCallbackSettingsRequest creates a new request for groups.getCallbackSettings
func NewGroupsGetCallbackSettingsRequest(a *api.API, actor actor.Actor) *GroupsGetCallbackSettingsRequest {
	return &GroupsGetCallbackSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.getCallbackSettings")}
}

// Exec executes the request and unmarshals the response into GroupsGetCallbackSettingsResponse
func (r *GroupsGetCallbackSettingsRequest) Exec(ctx context.Context) (response response.GroupsGetCallbackSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetCatalogInfoRequest defines the request for groups.getCatalogInfo
//
// The method returns a list of categories for the community catalog.
// Doc: https://dev.vk.com/method/groups.getCatalogInfo
type GroupsGetCatalogInfoRequest struct {
	BaseRequest
}

// NewGroupsGetCatalogInfoRequest creates a new request for groups.getCatalogInfo
func NewGroupsGetCatalogInfoRequest(a *api.API, actor actor.Actor) *GroupsGetCatalogInfoRequest {
	return &GroupsGetCatalogInfoRequest{*NewMethodBaseRequest(a, actor, "groups.getCatalogInfo")}
}

// Exec executes the request and unmarshals the response into GroupsGetCatalogInfoResponse
func (r *GroupsGetCatalogInfoRequest) Exec(ctx context.Context) (response response.GroupsGetCatalogInfoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetCatalogInfoExtendedRequest defines the request for groups.getCatalogInfo
//
// The method returns a list of categories for the community catalog.
// Doc: https://dev.vk.com/method/groups.getCatalogInfo
type GroupsGetCatalogInfoExtendedRequest struct {
	BaseRequest
}

// NewGroupsGetCatalogInfoExtendedRequest creates a new request for groups.getCatalogInfo
func NewGroupsGetCatalogInfoExtendedRequest(a *api.API, actor actor.Actor) *GroupsGetCatalogInfoExtendedRequest {
	r := &GroupsGetCatalogInfoExtendedRequest{*NewMethodBaseRequest(a, actor, "groups.getCatalogInfo")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into GroupsGetCatalogInfoExtendedResponse
func (r *GroupsGetCatalogInfoExtendedRequest) Exec(ctx context.Context) (response response.GroupsGetCatalogInfoExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetInvitedUsersRequest defines the request for groups.getInvitedUsers
//
// The method returns a list of users invited to the community by administrators.
// Doc: https://dev.vk.com/method/groups.getInvitedUsers
type GroupsGetInvitedUsersRequest struct {
	BaseRequest
}

// NewGroupsGetInvitedUsersRequest creates a new request for groups.getInvitedUsers
func NewGroupsGetInvitedUsersRequest(a *api.API, actor actor.Actor) *GroupsGetInvitedUsersRequest {
	return &GroupsGetInvitedUsersRequest{*NewMethodBaseRequest(a, actor, "groups.getInvitedUsers")}
}

// Exec executes the request and unmarshals the response into GroupsGetInvitedUsersResponse
func (r *GroupsGetInvitedUsersRequest) Exec(ctx context.Context) (response response.GroupsGetInvitedUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetInvitesRequest defines the request for groups.getInvites
//
// The method returns a list of invitations to communities and events for the current user.
// Doc: https://dev.vk.com/method/groups.getInvites
type GroupsGetInvitesRequest struct {
	BaseRequest
}

// NewGroupsGetInvitesRequest creates a new request for groups.getInvites
func NewGroupsGetInvitesRequest(a *api.API, actor actor.Actor) *GroupsGetInvitesRequest {
	return &GroupsGetInvitesRequest{*NewMethodBaseRequest(a, actor, "groups.getInvites")}
}

// Exec executes the request and unmarshals the response into GroupsGetInvitesResponse
func (r *GroupsGetInvitesRequest) Exec(ctx context.Context) (response response.GroupsGetInvitesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetInvitesExtendedRequest defines the request for groups.getInvites
//
// The method returns a list of invitations to communities and events for the current user.
// Doc: https://dev.vk.com/method/groups.getInvites
type GroupsGetInvitesExtendedRequest struct {
	BaseRequest
}

// NewGroupsGetInvitesExtendedRequest creates a new request for groups.getInvites
func NewGroupsGetInvitesExtendedRequest(a *api.API, actor actor.Actor) *GroupsGetInvitesExtendedRequest {
	r := &GroupsGetInvitesExtendedRequest{*NewMethodBaseRequest(a, actor, "groups.getInvites")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into GroupsGetInvitesExtendedResponse
func (r *GroupsGetInvitesExtendedRequest) Exec(ctx context.Context) (response response.GroupsGetInvitesExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetLongPollServerRequest defines the request for groups.getLongPollServer
//
// The method returns data for connecting to Bots Longpoll api.
// Doc: https://dev.vk.com/method/groups.getLongPollServer
type GroupsGetLongPollServerRequest struct {
	BaseRequest
}

// NewGroupsGetLongPollServerRequest creates a new request for groups.getLongPollServer
func NewGroupsGetLongPollServerRequest(a *api.API, actor actor.Actor) *GroupsGetLongPollServerRequest {
	return &GroupsGetLongPollServerRequest{*NewMethodBaseRequest(a, actor, "groups.getLongPollServer")}
}

// Exec executes the request and unmarshals the response into GroupsGetLongPollServerResponse
func (r *GroupsGetLongPollServerRequest) Exec(ctx context.Context) (response response.GroupsGetLongPollServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *GroupsGetLongPollServerRequest) GroupID(id int) *GroupsGetLongPollServerRequest {
	r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	return r
}

// GroupsGetLongPollSettingsRequest defines the request for groups.getLongPollSettings
//
// The method retrieves the Bots Longpoll api settings for the community.
// Doc: https://dev.vk.com/method/groups.getLongPollSettings
type GroupsGetLongPollSettingsRequest struct {
	BaseRequest
}

// NewGroupsGetLongPollSettingsRequest creates a new request for groups.getLongPollSettings
func NewGroupsGetLongPollSettingsRequest(a *api.API, actor actor.Actor) *GroupsGetLongPollSettingsRequest {
	return &GroupsGetLongPollSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.getLongPollSettings")}
}

// Exec executes the request and unmarshals the response into GroupsGetLongPollSettingsResponse
func (r *GroupsGetLongPollSettingsRequest) Exec(ctx context.Context) (response response.GroupsGetLongPollSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetMembersRequest defines the request for groups.getMembers
//
// The method returns a list of members in the community.
// Doc: https://dev.vk.com/method/groups.getMembers
type GroupsGetMembersRequest struct {
	BaseRequest
}

// NewGroupsGetMembersRequest creates a new request for groups.getMembers
func NewGroupsGetMembersRequest(a *api.API, actor actor.Actor) *GroupsGetMembersRequest {
	return &GroupsGetMembersRequest{*NewMethodBaseRequest(a, actor, "groups.getMembers")}
}

// Exec executes the request and unmarshals the response into GroupsGetMembersResponse
func (r *GroupsGetMembersRequest) Exec(ctx context.Context) (response response.GroupsGetMembersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetOnlineStatusRequest defines the request for groups.getOnlineStatus
//
// The method retrieves the online status of the community.
// Doc: https://dev.vk.com/method/groups.getOnlineStatus
type GroupsGetOnlineStatusRequest struct {
	BaseRequest
}

// NewGroupsGetOnlineStatusRequest creates a new request for groups.getOnlineStatus
func NewGroupsGetOnlineStatusRequest(a *api.API, actor actor.Actor) *GroupsGetOnlineStatusRequest {
	return &GroupsGetOnlineStatusRequest{*NewMethodBaseRequest(a, actor, "groups.getOnlineStatus")}
}

// Exec executes the request and unmarshals the response into GroupsGetOnlineStatusResponse
func (r *GroupsGetOnlineStatusRequest) Exec(ctx context.Context) (response response.GroupsGetOnlineStatusResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetRequestsRequest defines the request for groups.getRequests
//
// The method returns a list of membership requests for the community.
// Doc: https://dev.vk.com/method/groups.getRequests
type GroupsGetRequestsRequest struct {
	BaseRequest
}

// NewGroupsGetRequestsRequest creates a new request for groups.getRequests
func NewGroupsGetRequestsRequest(a *api.API, actor actor.Actor) *GroupsGetRequestsRequest {
	return &GroupsGetRequestsRequest{*NewMethodBaseRequest(a, actor, "groups.getRequests")}
}

// Exec executes the request and unmarshals the response into GroupsGetRequestsResponse
func (r *GroupsGetRequestsRequest) Exec(ctx context.Context) (response response.GroupsGetRequestsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetSettingsRequest defines the request for groups.getSettings
//
// The method retrieves data needed to display the community's settings page.
// Doc: https://dev.vk.com/method/groups.getSettings
type GroupsGetSettingsRequest struct {
	BaseRequest
}

// NewGroupsGetSettingsRequest creates a new request for groups.getSettings
func NewGroupsGetSettingsRequest(a *api.API, actor actor.Actor) *GroupsGetSettingsRequest {
	return &GroupsGetSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.getSettings")}
}

// Exec executes the request and unmarshals the response into GroupsGetSettingsResponse
func (r *GroupsGetSettingsRequest) Exec(ctx context.Context) (response response.GroupsGetSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetTagListRequest defines the request for groups.getTagList
//
// The method returns a list of tags in the community.
// Doc: https://dev.vk.com/method/groups.getTagList
type GroupsGetTagListRequest struct {
	BaseRequest
}

// NewGroupsGetTagListRequest creates a new request for groups.getTagList
func NewGroupsGetTagListRequest(a *api.API, actor actor.Actor) *GroupsGetTagListRequest {
	return &GroupsGetTagListRequest{*NewMethodBaseRequest(a, actor, "groups.getTagList")}
}

// Exec executes the request and unmarshals the response into GroupsGetTagListResponse
func (r *GroupsGetTagListRequest) Exec(ctx context.Context) (response response.GroupsGetTagListResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsGetTokenPermissionsRequest defines the request for groups.getTokenPermissions
//
// The method returns the access rights for the community's access token.
// Doc: https://dev.vk.com/method/groups.getTokenPermissions
type GroupsGetTokenPermissionsRequest struct {
	BaseRequest
}

// NewGroupsGetTokenPermissionsRequest creates a new request for groups.getTokenPermissions
func NewGroupsGetTokenPermissionsRequest(a *api.API, actor actor.Actor) *GroupsGetTokenPermissionsRequest {
	return &GroupsGetTokenPermissionsRequest{*NewMethodBaseRequest(a, actor, "groups.getTokenPermissions")}
}

// Exec executes the request and unmarshals the response into GroupsGetTokenPermissionsResponse
func (r *GroupsGetTokenPermissionsRequest) Exec(ctx context.Context) (response response.GroupsGetTokenPermissionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsInviteRequest defines the request for groups.invite
//
// The method allows inviting friends to the community.
// Doc: https://dev.vk.com/method/groups.invite
type GroupsInviteRequest struct {
	BaseRequest
}

// NewGroupsInviteRequest creates a new request for groups.invite
func NewGroupsInviteRequest(a *api.API, actor actor.Actor) *GroupsInviteRequest {
	return &GroupsInviteRequest{*NewMethodBaseRequest(a, actor, "groups.invite")}
}

// Exec executes the request and unmarshals the response into GroupsInviteResponse
func (r *GroupsInviteRequest) Exec(ctx context.Context) (response response.GroupsInviteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsIsMemberRequest defines the request for groups.isMember
//
// The method returns information on whether a user is a member of the community.
// Doc: https://dev.vk.com/method/groups.isMember
type GroupsIsMemberRequest struct {
	BaseRequest
}

// NewGroupsIsMemberRequest creates a new request for groups.isMember
func NewGroupsIsMemberRequest(a *api.API, actor actor.Actor) *GroupsIsMemberRequest {
	return &GroupsIsMemberRequest{*NewMethodBaseRequest(a, actor, "groups.isMember")}
}

// Exec executes the request and unmarshals the response into GroupsIsMemberResponse
func (r *GroupsIsMemberRequest) Exec(ctx context.Context) (response response.GroupsIsMemberResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsIsMemberExtendedRequest defines the request for groups.isMember
//
// The method returns information on whether a user is a member of the community.
// Doc: https://dev.vk.com/method/groups.isMember
type GroupsIsMemberExtendedRequest struct {
	BaseRequest
}

// NewGroupsIsMemberExtendedRequest creates a new request for groups.isMember
func NewGroupsIsMemberExtendedRequest(a *api.API, actor actor.Actor) *GroupsIsMemberExtendedRequest {
	r := &GroupsIsMemberExtendedRequest{*NewMethodBaseRequest(a, actor, "groups.isMember")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into GroupsIsMemberExtendedResponse
func (r *GroupsIsMemberExtendedRequest) Exec(ctx context.Context) (response response.GroupsIsMemberExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsIsMemberUserIDsRequest defines the request for groups.isMember
//
// The method returns information on whether a user is a member of the community.
// Doc: https://dev.vk.com/method/groups.isMember
type GroupsIsMemberUserIDsRequest struct {
	BaseRequest
}

// NewGroupsIsMemberUserIDsRequest creates a new request for groups.isMember
func NewGroupsIsMemberUserIDsRequest(a *api.API, actor actor.Actor) *GroupsIsMemberUserIDsRequest {
	return &GroupsIsMemberUserIDsRequest{*NewMethodBaseRequest(a, actor, "groups.isMember")}
}

// Exec executes the request and unmarshals the response into GroupsIsMemberUserIDsResponse
func (r *GroupsIsMemberUserIDsRequest) Exec(ctx context.Context) (response response.GroupsIsMemberUserIDsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsIsMemberUserIDsExtendedRequest defines the request for groups.isMember
//
// The method returns information on whether a user is a member of the community.
// Doc: https://dev.vk.com/method/groups.isMember
type GroupsIsMemberUserIDsExtendedRequest struct {
	BaseRequest
}

// NewGroupsIsMemberUserIDsExtendedRequest creates a new request for groups.isMember
func NewGroupsIsMemberUserIDsExtendedRequest(a *api.API, actor actor.Actor) *GroupsIsMemberUserIDsExtendedRequest {
	r := &GroupsIsMemberUserIDsExtendedRequest{*NewMethodBaseRequest(a, actor, "groups.isMember")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into GroupsIsMemberUserIDsExtendedResponse
func (r *GroupsIsMemberUserIDsExtendedRequest) Exec(ctx context.Context) (response response.GroupsIsMemberUserIDsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// UserIDs count users < 500
func (r *GroupsIsMemberUserIDsExtendedRequest) UserIDs(ids string) *GroupsIsMemberUserIDsExtendedRequest {
	_ = r.parameters.Set(constants.ParameterNameUserIDs, ids)
	return r
}

// GroupsJoinRequest defines the request for groups.join
//
// The method allows a user to join a group, public page, or confirm participation in an event.
// Doc: https://dev.vk.com/method/groups.join
type GroupsJoinRequest struct {
	BaseRequest
}

// NewGroupsJoinRequest creates a new request for groups.join
func NewGroupsJoinRequest(a *api.API, actor actor.Actor) *GroupsJoinRequest {
	return &GroupsJoinRequest{*NewMethodBaseRequest(a, actor, "groups.join")}
}

// Exec executes the request and unmarshals the response into GroupsJoinResponse
func (r *GroupsJoinRequest) Exec(ctx context.Context) (response response.GroupsJoinResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsLeaveRequest defines the request for groups.leave
//
// The method allows a user to leave a community or decline an invitation to a community.
// Doc: https://dev.vk.com/method/groups.leave
type GroupsLeaveRequest struct {
	BaseRequest
}

// NewGroupsLeaveRequest creates a new request for groups.leave
func NewGroupsLeaveRequest(a *api.API, actor actor.Actor) *GroupsLeaveRequest {
	return &GroupsLeaveRequest{*NewMethodBaseRequest(a, actor, "groups.leave")}
}

// Exec executes the request and unmarshals the response into GroupsLeaveResponse
func (r *GroupsLeaveRequest) Exec(ctx context.Context) (response response.GroupsLeaveResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsRemoveUserRequest defines the request for groups.removeUser
//
// The method allows removing a user from a community or rejecting their membership request.
// Doc: https://dev.vk.com/method/groups.removeUser
type GroupsRemoveUserRequest struct {
	BaseRequest
}

// NewGroupsRemoveUserRequest creates a new request for groups.removeUser
func NewGroupsRemoveUserRequest(a *api.API, actor actor.Actor) *GroupsRemoveUserRequest {
	return &GroupsRemoveUserRequest{*NewMethodBaseRequest(a, actor, "groups.removeUser")}
}

// Exec executes the request and unmarshals the response into GroupsRemoveUserResponse
func (r *GroupsRemoveUserRequest) Exec(ctx context.Context) (response response.GroupsRemoveUserResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsReorderLinkRequest defines the request for groups.reorderLink
//
// The method allows changing the order of a link in the community's link list.
// Doc: https://dev.vk.com/method/groups.reorderLink
type GroupsReorderLinkRequest struct {
	BaseRequest
}

// NewGroupsReorderLinkRequest creates a new request for groups.reorderLink
func NewGroupsReorderLinkRequest(a *api.API, actor actor.Actor) *GroupsReorderLinkRequest {
	return &GroupsReorderLinkRequest{*NewMethodBaseRequest(a, actor, "groups.reorderLink")}
}

// Exec executes the request and unmarshals the response into GroupsReorderLinkResponse
func (r *GroupsReorderLinkRequest) Exec(ctx context.Context) (response response.GroupsReorderLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsSearchRequest defines the request for groups.search
//
// The method searches for communities by a specified query string.
// Doc: https://dev.vk.com/method/groups.search
type GroupsSearchRequest struct {
	BaseRequest
}

// NewGroupsSearchRequest creates a new request for groups.search
func NewGroupsSearchRequest(a *api.API, actor actor.Actor) *GroupsSearchRequest {
	return &GroupsSearchRequest{*NewMethodBaseRequest(a, actor, "groups.search")}
}

// Exec executes the request and unmarshals the response into GroupsSearchResponse
func (r *GroupsSearchRequest) Exec(ctx context.Context) (response response.GroupsSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsSetCallbackSettingsRequest defines the request for groups.setCallbackSettings
//
// The method sets up notification settings for Callback api in the community.
// Doc: https://dev.vk.com/method/groups.setCallbackSettings
type GroupsSetCallbackSettingsRequest struct {
	BaseRequest
}

// NewGroupsSetCallbackSettingsRequest creates a new request for groups.setCallbackSettings
func NewGroupsSetCallbackSettingsRequest(a *api.API, actor actor.Actor) *GroupsSetCallbackSettingsRequest {
	return &GroupsSetCallbackSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.setCallbackSettings")}
}

// Exec executes the request and unmarshals the response into GroupsSetCallbackSettingsResponse
func (r *GroupsSetCallbackSettingsRequest) Exec(ctx context.Context) (response response.GroupsSetCallbackSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *GroupsSetCallbackSettingsRequest) GroupID(id int) *GroupsSetCallbackSettingsRequest {
	if id > 0 {
		r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	}
	return r
}

func (r *GroupsSetCallbackSettingsRequest) ServerID(id int) *GroupsSetCallbackSettingsRequest {
	if id > 0 {
		r.parameters.Set(constants.ParameterNameServerID, strconv.Itoa(id))
	}
	return r
}

func (r *GroupsSetCallbackSettingsRequest) APIVersion(v string) *GroupsSetCallbackSettingsRequest {
	r.parameters.Set(constants.ParameterNameAPIVersion, v)
	return r
}

func (r *GroupsSetCallbackSettingsRequest) SetEvent(event string, isEnable bool) *GroupsSetCallbackSettingsRequest {
	if isEnable {
		r.parameters.Set(event, "1")
	} else {
		r.parameters.Set(event, "0")
	}
	return r
}

func (r *GroupsSetCallbackSettingsRequest) ResetEvents() *GroupsSetCallbackSettingsRequest {
	gid := r.parameters.Get(constants.ParameterNameGroupID)
	sid := r.parameters.Get(constants.ParameterNameServerID)
	v := r.parameters.Get(constants.ParameterNameAPIVersion)

	r.ResetParameters()

	r.parameters.SetIfNotEmpty(constants.ParameterNameGroupID, gid)
	r.parameters.SetIfNotEmpty(constants.ParameterNameServerID, sid)
	r.parameters.SetIfNotEmpty(constants.ParameterNameAPIVersion, v)

	return r
}

// GroupsSetLongPollSettingsRequest defines the request for groups.setLongPollSettings
//
// The method sets the Bots Long Poll api settings for the community.
// Doc: https://dev.vk.com/method/groups.setLongPollSettings
type GroupsSetLongPollSettingsRequest struct {
	BaseRequest
}

// NewGroupsSetLongPollSettingsRequest creates a new request for groups.setLongPollSettings
func NewGroupsSetLongPollSettingsRequest(a *api.API, actor actor.Actor) *GroupsSetLongPollSettingsRequest {
	return &GroupsSetLongPollSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.setLongPollSettings")}
}

// Exec executes the request and unmarshals the response into GroupsSetLongPollSettingsResponse
func (r *GroupsSetLongPollSettingsRequest) Exec(ctx context.Context) (response response.GroupsSetLongPollSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *GroupsSetLongPollSettingsRequest) GroupID(id int) *GroupsSetLongPollSettingsRequest {
	if id > 0 {
		r.parameters.Set(constants.ParameterNameGroupID, strconv.Itoa(id))
	}
	return r
}

func (r *GroupsSetLongPollSettingsRequest) Enabled(v bool) *GroupsSetLongPollSettingsRequest {
	if v {
		r.parameters.Set(constants.ParameterNameEnabled, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameEnabled)
	}
	return r
}

func (r *GroupsSetLongPollSettingsRequest) APIVersion(v string) *GroupsSetLongPollSettingsRequest {
	r.parameters.Set(constants.ParameterNameAPIVersion, v)
	return r
}

func (r *GroupsSetLongPollSettingsRequest) SetEvent(event string, isEnable bool) *GroupsSetLongPollSettingsRequest {
	if isEnable {
		r.parameters.Set(event, "1")
	} else {
		r.parameters.Set(event, "0")
	}
	return r
}

func (r *GroupsSetLongPollSettingsRequest) ResetEvents() *GroupsSetLongPollSettingsRequest {
	gid := r.parameters.Get(constants.ParameterNameGroupID)
	e := r.parameters.Get(constants.ParameterNameEnabled)
	v := r.parameters.Get(constants.ParameterNameAPIVersion)

	r.ResetParameters()

	r.parameters.SetIfNotEmpty(constants.ParameterNameGroupID, gid)
	r.parameters.SetIfNotEmpty(constants.ParameterNameEnabled, e)
	r.parameters.SetIfNotEmpty(constants.ParameterNameAPIVersion, v)

	return r
}

// GroupsSetSettingsRequest defines the request for groups.setSettings
//
// The method sets the settings of the community.
// Doc: https://dev.vk.com/method/groups.setSettings
type GroupsSetSettingsRequest struct {
	BaseRequest
}

// NewGroupsSetSettingsRequest creates a new request for groups.setSettings
func NewGroupsSetSettingsRequest(a *api.API, actor actor.Actor) *GroupsSetSettingsRequest {
	return &GroupsSetSettingsRequest{*NewMethodBaseRequest(a, actor, "groups.setSettings")}
}

// Exec executes the request and unmarshals the response into GroupsSetSettingsResponse
func (r *GroupsSetSettingsRequest) Exec(ctx context.Context) (response response.GroupsSetSettingsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsSetUserNoteRequest defines the request for groups.setUserNote
//
// The method allows creating or editing a note about a user during correspondence with the community.
// Doc: https://dev.vk.com/method/groups.setUserNote
type GroupsSetUserNoteRequest struct {
	BaseRequest
}

// NewGroupsSetUserNoteRequest creates a new request for groups.setUserNote
func NewGroupsSetUserNoteRequest(a *api.API, actor actor.Actor) *GroupsSetUserNoteRequest {
	return &GroupsSetUserNoteRequest{*NewMethodBaseRequest(a, actor, "groups.setUserNote")}
}

// Exec executes the request and unmarshals the response into GroupsSetUserNoteResponse
func (r *GroupsSetUserNoteRequest) Exec(ctx context.Context) (response response.GroupsSetUserNoteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsTagAddRequest defines the request for groups.tagAdd
//
// The method allows adding a new tag to the community.
// Doc: https://dev.vk.com/method/groups.tagAdd
type GroupsTagAddRequest struct {
	BaseRequest
}

// NewGroupsTagAddRequest creates a new request for groups.tagAdd
func NewGroupsTagAddRequest(a *api.API, actor actor.Actor) *GroupsTagAddRequest {
	return &GroupsTagAddRequest{*NewMethodBaseRequest(a, actor, "groups.tagAdd")}
}

// Exec executes the request and unmarshals the response into GroupsTagAddResponse
func (r *GroupsTagAddRequest) Exec(ctx context.Context) (response response.GroupsTagAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsTagBindRequest defines the request for groups.tagBind
//
// The method allows binding or unbinding community tags to conversations.
// Doc: https://dev.vk.com/method/groups.tagBind
type GroupsTagBindRequest struct {
	BaseRequest
}

// NewGroupsTagBindRequest creates a new request for groups.tagBind
func NewGroupsTagBindRequest(a *api.API, actor actor.Actor) *GroupsTagBindRequest {
	return &GroupsTagBindRequest{*NewMethodBaseRequest(a, actor, "groups.tagBind")}
}

// Exec executes the request and unmarshals the response into GroupsTagBindResponse
func (r *GroupsTagBindRequest) Exec(ctx context.Context) (response response.GroupsTagBindResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsTagDeleteRequest defines the request for groups.tagDelete
//
// The method allows deleting a community tag.
// Doc: https://dev.vk.com/method/groups.tagDelete
type GroupsTagDeleteRequest struct {
	BaseRequest
}

// NewGroupsTagDeleteRequest creates a new request for groups.tagDelete
func NewGroupsTagDeleteRequest(a *api.API, actor actor.Actor) *GroupsTagDeleteRequest {
	return &GroupsTagDeleteRequest{*NewMethodBaseRequest(a, actor, "groups.tagDelete")}
}

// Exec executes the request and unmarshals the response into GroupsTagDeleteResponse
func (r *GroupsTagDeleteRequest) Exec(ctx context.Context) (response response.GroupsTagDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsTagUpdateRequest defines the request for groups.tagUpdate
//
// The method allows renaming an existing community tag.
// Doc: https://dev.vk.com/method/groups.tagUpdate
type GroupsTagUpdateRequest struct {
	BaseRequest
}

// NewGroupsTagUpdateRequest creates a new request for groups.tagUpdate
func NewGroupsTagUpdateRequest(a *api.API, actor actor.Actor) *GroupsTagUpdateRequest {
	return &GroupsTagUpdateRequest{*NewMethodBaseRequest(a, actor, "groups.tagUpdate")}
}

// Exec executes the request and unmarshals the response into GroupsTagUpdateResponse
func (r *GroupsTagUpdateRequest) Exec(ctx context.Context) (response response.GroupsTagUpdateResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsToggleMarketRequest defines the request for groups.toggleMarket
//
// The method toggles the market functionality in a selected community.
// Doc: https://dev.vk.com/method/groups.toggleMarket
type GroupsToggleMarketRequest struct {
	BaseRequest
}

// NewGroupsToggleMarketRequest creates a new request for groups.toggleMarket
func NewGroupsToggleMarketRequest(a *api.API, actor actor.Actor) *GroupsToggleMarketRequest {
	return &GroupsToggleMarketRequest{*NewMethodBaseRequest(a, actor, "groups.toggleMarket")}
}

// Exec executes the request and unmarshals the response into GroupsToggleMarketResponse
func (r *GroupsToggleMarketRequest) Exec(ctx context.Context) (response response.GroupsToggleMarketResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// GroupsUnbanRequest defines the request for groups.unban
//
// The method unbans a user or group from the community's blacklist.
// Doc: https://dev.vk.com/method/groups.unban
type GroupsUnbanRequest struct {
	BaseRequest
}

// NewGroupsUnbanRequest creates a new request for groups.unban
func NewGroupsUnbanRequest(a *api.API, actor actor.Actor) *GroupsUnbanRequest {
	return &GroupsUnbanRequest{*NewMethodBaseRequest(a, actor, "groups.unban")}
}

// Exec executes the request and unmarshals the response into GroupsUnbanResponse
func (r *GroupsUnbanRequest) Exec(ctx context.Context) (response response.GroupsUnbanResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
