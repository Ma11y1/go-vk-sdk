package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Groups Doc: https://dev.vk.com/ru/method/groups
type Groups struct {
	BaseAction
}

// AddAddress Doc: https://dev.vk.com/ru/method/groups.addAddress
func (g *Groups) AddAddress(actor actor.Actor) *request.GroupsAddAddressRequest {
	return request.NewGroupsAddAddressRequest(g.api, actor)
}

// AddCallbackServer Doc: https://dev.vk.com/ru/method/groups.addCallbackServer
func (g *Groups) AddCallbackServer(actor actor.Actor) *request.GroupsAddCallbackServerRequest {
	return request.NewGroupsAddCallbackServerRequest(g.api, actor)
}

// AddLink Doc: https://dev.vk.com/ru/method/groups.addLink
func (g *Groups) AddLink(actor actor.Actor) *request.GroupsAddLinkRequest {
	return request.NewGroupsAddLinkRequest(g.api, actor)
}

// ApproveRequest Doc: https://dev.vk.com/ru/method/groups.approveRequest
func (g *Groups) ApproveRequest(actor actor.Actor) *request.GroupsApproveRequest {
	return request.NewGroupsApproveRequest(g.api, actor)
}

// Ban Doc: https://dev.vk.com/ru/method/groups.ban
func (g *Groups) Ban(actor actor.Actor) *request.GroupsBanRequest {
	return request.NewGroupsBanRequest(g.api, actor)
}

// Create Doc: https://dev.vk.com/ru/method/groups.create
func (g *Groups) Create(actor actor.Actor) *request.GroupsCreateRequest {
	return request.NewGroupsCreateRequest(g.api, actor)
}

// DeleteAddress Doc: https://dev.vk.com/ru/method/groups.deleteAddress
func (g *Groups) DeleteAddress(actor actor.Actor) *request.GroupsDeleteAddressRequest {
	return request.NewGroupsDeleteAddressRequest(g.api, actor)
}

// DeleteCallbackServer Doc: https://dev.vk.com/ru/method/groups.deleteCallbackServer
func (g *Groups) DeleteCallbackServer(actor actor.Actor) *request.GroupsDeleteCallbackServerRequest {
	return request.NewGroupsDeleteCallbackServerRequest(g.api, actor)
}

// DeleteLink Doc: https://dev.vk.com/ru/method/groups.deleteLink
func (g *Groups) DeleteLink(actor actor.Actor) *request.GroupsDeleteLinkRequest {
	return request.NewGroupsDeleteLinkRequest(g.api, actor)
}

// DisableOnline Doc: https://dev.vk.com/ru/method/groups.disableOnline
func (g *Groups) DisableOnline(actor actor.Actor) *request.GroupsDisableOnlineRequest {
	return request.NewGroupsDisableOnlineRequest(g.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/groups.edit
func (g *Groups) Edit(actor actor.Actor) *request.GroupsEditRequest {
	return request.NewGroupsEditRequest(g.api, actor)
}

// EditAddress Doc: https://dev.vk.com/ru/method/groups.editAddress
func (g *Groups) EditAddress(actor actor.Actor) *request.GroupsEditAddressRequest {
	return request.NewGroupsEditAddressRequest(g.api, actor)
}

// EditCallbackServer Doc: https://dev.vk.com/ru/method/groups.editCallbackServer
func (g *Groups) EditCallbackServer(actor actor.Actor) *request.GroupsEditCallbackServerRequest {
	return request.NewGroupsEditCallbackServerRequest(g.api, actor)
}

// EditLink Doc: https://dev.vk.com/ru/method/groups.editLink
func (g *Groups) EditLink(actor actor.Actor) *request.GroupsEditLinkRequest {
	return request.NewGroupsEditLinkRequest(g.api, actor)
}

// EditManager Doc: https://dev.vk.com/ru/method/groups.editManager
func (g *Groups) EditManager(actor actor.Actor) *request.GroupsEditManagerRequest {
	return request.NewGroupsEditManagerRequest(g.api, actor)
}

// EnableOnline Doc: https://dev.vk.com/ru/method/groups.enableOnline
func (g *Groups) EnableOnline(actor actor.Actor) *request.GroupsEnableOnlineRequest {
	return request.NewGroupsEnableOnlineRequest(g.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/groups.get
func (g *Groups) Get(actor actor.Actor) *request.GroupsGetRequest {
	return request.NewGroupsGetRequest(g.api, actor)
}

// GetExtended Doc: https://dev.vk.com/ru/method/groups.get
func (g *Groups) GetExtended(actor actor.Actor) *request.GroupsGetExtendedRequest {
	return request.NewGroupsGetExtendedRequest(g.api, actor)
}

// GetAddresses Doc: https://dev.vk.com/ru/method/groups.getAddresses
func (g *Groups) GetAddresses(actor actor.Actor) *request.GroupsGetAddressesRequest {
	return request.NewGroupsGetAddressesRequest(g.api, actor)
}

// GetBanned Doc: https://dev.vk.com/ru/method/groups.getBanned
func (g *Groups) GetBanned(actor actor.Actor) *request.GroupsGetBannedRequest {
	return request.NewGroupsGetBannedRequest(g.api, actor)
}

// GetByID Doc: https://dev.vk.com/ru/method/groups.getById
func (g *Groups) GetByID(actor actor.Actor) *request.GroupsGetByIDRequest {
	return request.NewGroupsGetByIDRequest(g.api, actor)
}

// GetCallbackConfirmationCode Doc: https://dev.vk.com/ru/method/groups.getCallbackConfirmationCode
func (g *Groups) GetCallbackConfirmationCode(actor actor.Actor) *request.GroupsGetCallbackConfirmationCodeRequest {
	return request.NewGroupsGetCallbackConfirmationCodeRequest(g.api, actor)
}

// GetCallbackServers Doc: https://dev.vk.com/ru/method/groups.getCallbackServers
func (g *Groups) GetCallbackServers(actor actor.Actor) *request.GroupsGetCallbackServersRequest {
	return request.NewGroupsGetCallbackServersRequest(g.api, actor)
}

// GetCallbackSettings Doc: https://dev.vk.com/ru/method/groups.getCallbackSettings
func (g *Groups) GetCallbackSettings(actor actor.Actor) *request.GroupsGetCallbackSettingsRequest {
	return request.NewGroupsGetCallbackSettingsRequest(g.api, actor)
}

// GetCatalogInfo Doc: https://dev.vk.com/ru/method/groups.getCatalogInfo
func (g *Groups) GetCatalogInfo(actor actor.Actor) *request.GroupsGetCatalogInfoRequest {
	return request.NewGroupsGetCatalogInfoRequest(g.api, actor)
}

// GetCatalogInfoExtended Doc: https://dev.vk.com/ru/method/groups.getCatalogInfo
func (g *Groups) GetCatalogInfoExtended(actor actor.Actor) *request.GroupsGetCatalogInfoExtendedRequest {
	return request.NewGroupsGetCatalogInfoExtendedRequest(g.api, actor)
}

// GetInvitedUsers Doc: https://dev.vk.com/ru/method/groups.getInvitedUsers
func (g *Groups) GetInvitedUsers(actor actor.Actor) *request.GroupsGetInvitedUsersRequest {
	return request.NewGroupsGetInvitedUsersRequest(g.api, actor)
}

// GetInvites Doc: https://dev.vk.com/ru/method/groups.getInvites
func (g *Groups) GetInvites(actor actor.Actor) *request.GroupsGetInvitesRequest {
	return request.NewGroupsGetInvitesRequest(g.api, actor)
}

// GetInvitesExtended Doc: https://dev.vk.com/ru/method/groups.getInvites
func (g *Groups) GetInvitesExtended(actor actor.Actor) *request.GroupsGetInvitesExtendedRequest {
	return request.NewGroupsGetInvitesExtendedRequest(g.api, actor)
}

// GetLongPollServer Doc: https://dev.vk.com/ru/method/groups.getLongPollServer
func (g *Groups) GetLongPollServer(actor actor.Actor) *request.GroupsGetLongPollServerRequest {
	return request.NewGroupsGetLongPollServerRequest(g.api, actor)
}

// GetLongPollSettings Doc: https://dev.vk.com/ru/method/groups.getLongPollSettings
func (g *Groups) GetLongPollSettings(actor actor.Actor) *request.GroupsGetLongPollSettingsRequest {
	return request.NewGroupsGetLongPollSettingsRequest(g.api, actor)
}

// GetMembers Doc: https://dev.vk.com/ru/method/groups.getMembers
func (g *Groups) GetMembers(actor actor.Actor) *request.GroupsGetMembersRequest {
	return request.NewGroupsGetMembersRequest(g.api, actor)
}

// GetOnlineStatus Doc: https://dev.vk.com/ru/method/groups.getOnlineStatus
func (g *Groups) GetOnlineStatus(actor actor.Actor) *request.GroupsGetOnlineStatusRequest {
	return request.NewGroupsGetOnlineStatusRequest(g.api, actor)
}

// GetRequests Doc: https://dev.vk.com/ru/method/groups.getRequests
func (g *Groups) GetRequests(actor actor.Actor) *request.GroupsGetRequestsRequest {
	return request.NewGroupsGetRequestsRequest(g.api, actor)
}

// GetSettings Doc: https://dev.vk.com/ru/method/groups.getSettings
func (g *Groups) GetSettings(actor actor.Actor) *request.GroupsGetSettingsRequest {
	return request.NewGroupsGetSettingsRequest(g.api, actor)
}

// GetTagList Doc: https://dev.vk.com/ru/method/groups.getTagList
func (g *Groups) GetTagList(actor actor.Actor) *request.GroupsGetTagListRequest {
	return request.NewGroupsGetTagListRequest(g.api, actor)
}

// GetTokenPermissions Doc: https://dev.vk.com/ru/method/groups.getTokenPermissions
func (g *Groups) GetTokenPermissions(actor actor.Actor) *request.GroupsGetTokenPermissionsRequest {
	return request.NewGroupsGetTokenPermissionsRequest(g.api, actor)
}

// Invite Doc: https://dev.vk.com/ru/method/groups.invite
func (g *Groups) Invite(actor actor.Actor) *request.GroupsInviteRequest {
	return request.NewGroupsInviteRequest(g.api, actor)
}

// IsMember Doc: https://dev.vk.com/ru/method/groups.isMember
func (g *Groups) IsMember(actor actor.Actor) *request.GroupsIsMemberRequest {
	return request.NewGroupsIsMemberRequest(g.api, actor)
}

// IsMemberExtended Doc: https://dev.vk.com/ru/method/groups.isMember
func (g *Groups) IsMemberExtended(actor actor.Actor) *request.GroupsIsMemberExtendedRequest {
	return request.NewGroupsIsMemberExtendedRequest(g.api, actor)
}

// IsMemberUserIDs Doc: https://dev.vk.com/ru/method/groups.isMember
func (g *Groups) IsMemberUserIDs(actor actor.Actor) *request.GroupsIsMemberUserIDsRequest {
	return request.NewGroupsIsMemberUserIDsRequest(g.api, actor)
}

// IsMemberUserIDsExtended Doc: https://dev.vk.com/ru/method/groups.isMember
func (g *Groups) IsMemberUserIDsExtended(actor actor.Actor) *request.GroupsIsMemberUserIDsExtendedRequest {
	return request.NewGroupsIsMemberUserIDsExtendedRequest(g.api, actor)
}

// Join Doc: https://dev.vk.com/ru/method/groups.join
func (g *Groups) Join(actor actor.Actor) *request.GroupsJoinRequest {
	return request.NewGroupsJoinRequest(g.api, actor)
}

// Leave Doc: https://dev.vk.com/ru/method/groups.leave
func (g *Groups) Leave(actor actor.Actor) *request.GroupsLeaveRequest {
	return request.NewGroupsLeaveRequest(g.api, actor)
}

// RemoveUser Doc: https://dev.vk.com/ru/method/groups.removeUser
func (g *Groups) RemoveUser(actor actor.Actor) *request.GroupsRemoveUserRequest {
	return request.NewGroupsRemoveUserRequest(g.api, actor)
}

// ReorderLink Doc: https://dev.vk.com/ru/method/groups.reorderLink
func (g *Groups) ReorderLink(actor actor.Actor) *request.GroupsReorderLinkRequest {
	return request.NewGroupsReorderLinkRequest(g.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/groups.search
func (g *Groups) Search(actor actor.Actor) *request.GroupsSearchRequest {
	return request.NewGroupsSearchRequest(g.api, actor)
}

// SetCallbackSettings Doc: https://dev.vk.com/ru/method/groups.setCallbackSettings
func (g *Groups) SetCallbackSettings(actor actor.Actor) *request.GroupsSetCallbackSettingsRequest {
	return request.NewGroupsSetCallbackSettingsRequest(g.api, actor)
}

// SetLongPollSettings Doc: https://dev.vk.com/ru/method/groups.setLongPollSettings
func (g *Groups) SetLongPollSettings(actor actor.Actor) *request.GroupsSetLongPollSettingsRequest {
	return request.NewGroupsSetLongPollSettingsRequest(g.api, actor)
}

// SetSettings Doc: https://dev.vk.com/ru/method/groups.setSettings
func (g *Groups) SetSettings(actor actor.Actor) *request.GroupsSetSettingsRequest {
	return request.NewGroupsSetSettingsRequest(g.api, actor)
}

// SetUserNote Doc: https://dev.vk.com/ru/method/groups.setUserNote
func (g *Groups) SetUserNote(actor actor.Actor) *request.GroupsSetUserNoteRequest {
	return request.NewGroupsSetUserNoteRequest(g.api, actor)
}

// TagAdd Doc: https://dev.vk.com/ru/method/groups.tagAdd
func (g *Groups) TagAdd(actor actor.Actor) *request.GroupsTagAddRequest {
	return request.NewGroupsTagAddRequest(g.api, actor)
}

// TagBind Doc: https://dev.vk.com/ru/method/groups.tagBind
func (g *Groups) TagBind(actor actor.Actor) *request.GroupsTagBindRequest {
	return request.NewGroupsTagBindRequest(g.api, actor)
}

// TagDelete Doc: https://dev.vk.com/ru/method/groups.tagDelete
func (g *Groups) TagDelete(actor actor.Actor) *request.GroupsTagDeleteRequest {
	return request.NewGroupsTagDeleteRequest(g.api, actor)
}

// TagUpdate Doc: https://dev.vk.com/ru/method/groups.tagUpdate
func (g *Groups) TagUpdate(actor actor.Actor) *request.GroupsTagUpdateRequest {
	return request.NewGroupsTagUpdateRequest(g.api, actor)
}

// ToggleMarket Doc: https://dev.vk.com/ru/method/groups.toggleMarket
func (g *Groups) ToggleMarket(actor actor.Actor) *request.GroupsToggleMarketRequest {
	return request.NewGroupsToggleMarketRequest(g.api, actor)
}

// Unban Doc: https://dev.vk.com/ru/method/groups.unban
func (g *Groups) Unban(actor actor.Actor) *request.GroupsUnbanRequest {
	return request.NewGroupsUnbanRequest(g.api, actor)
}
