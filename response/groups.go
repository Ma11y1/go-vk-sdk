package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/method/groups

type GroupsAddAddressResponse struct {
	BaseResponse
	Response objects.GroupAddress `json:"response"`
}

type GroupsAddCallbackServerResponse struct {
	BaseResponse
	Response struct {
		ServerID int `json:"server_id"`
	} `json:"response"`
}

type GroupsAddLinkResponse struct {
	BaseResponse
	Response objects.GroupLink `json:"response"`
}

type GroupsApproveRequestResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsBanResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsCreateResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsDeleteAddressResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsDeleteCallbackServerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsDeleteLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsDisableOnlineResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsEditAddressResponse struct {
	BaseResponse
	Response objects.GroupAddress `json:"response"`
}

type GroupsEditCallbackServerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsEditLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsEditManagerResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsEnableOnlineResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsGetResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type GroupsGetExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                 `json:"count"`
		Items []objects.GroupFull `json:"items"`
	} `json:"response"`
}

type GroupsGetAddressesResponse struct {
	BaseResponse
	Response struct {
		Count int                    `json:"count"`
		Items []objects.GroupAddress `json:"items"`
	} `json:"response"`
}

type GroupsGetBannedResponse struct {
	BaseResponse
	Response struct {
		Count int                            `json:"count"`
		Items []objects.GroupOwnerXtrBanInfo `json:"items"`
	} `json:"response"`
}

type GroupsGetByIDResponse struct {
	BaseResponse
	Response struct {
		Groups   []objects.GroupFull        `json:"groups"`
		Profiles []objects.GroupProfileItem `json:"profiles"`
	} `json:"response"`
}

type GroupsGetCallbackConfirmationCodeResponse struct {
	BaseResponse
	Response struct {
		Code string `json:"code"`
	} `json:"response"`
}

type GroupsGetCallbackServersResponse struct {
	BaseResponse
	Response struct {
		Count int                           `json:"count"`
		Items []objects.GroupCallbackServer `json:"items"`
	} `json:"response"`
}

type GroupsGetCallbackSettingsResponse struct {
	BaseResponse
	Response objects.GroupCallbackSettings `json:"response"`
}

type GroupsGetCatalogResponse struct {
	BaseResponse
	Response struct {
		Count int                 `json:"count"`
		Items []objects.GroupFull `json:"items"`
	} `json:"response"`
}

type GroupsGetCatalogInfoResponse struct {
	BaseResponse
	Response struct {
		Enabled    objects.BoolInt         `json:"enabled"`
		Categories []objects.GroupCategory `json:"categories"`
	} `json:"response"`
}

type GroupsGetCatalogInfoExtendedResponse struct {
	BaseResponse
	Response struct {
		Enabled    objects.BoolInt             `json:"enabled"`
		Categories []objects.GroupCategoryFull `json:"categories"`
	} `json:"response"`
}

type GroupsGetInvitedUsersResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type GroupsGetInvitesResponse struct {
	BaseResponse
	Response struct {
		Count int                         `json:"count"`
		Items []objects.GroupXtrInvitedBy `json:"items"`
	} `json:"response"`
}

type GroupsGetInvitesExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                         `json:"count"`
		Items []objects.GroupXtrInvitedBy `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type GroupsGetLongPollServerResponse struct {
	BaseResponse
	Response objects.GroupLongPollServer `json:"response"`
}

type GroupsGetLongPollSettingsResponse struct {
	BaseResponse
	Response objects.GroupLongPollSettings `json:"response"`
}

type GroupsGetMembersResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}
type GroupsGetMembersFieldsResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}
type GroupsGetMembersFilterManagersResponse struct {
	BaseResponse
	Response struct {
		Count int                              `json:"count"`
		Items []objects.GroupMemberRoleXtrUser `json:"items"`
	} `json:"response"`
}
type GroupsGetOnlineStatusResponse struct {
	BaseResponse
	Response objects.GroupOnlineStatus `json:"response"`
}

type GroupsGetRequestsResponse struct {
	BaseResponse
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type GroupsGetRequestsFieldsResponse struct {
	BaseResponse
	Response struct {
		Count int                `json:"count"`
		Items []objects.UserFull `json:"items"`
	} `json:"response"`
}

type GroupsGetSettingsResponse struct {
	BaseResponse
	Response objects.GroupSettings `json:"response"`
}

type GroupsGetTagListResponse struct {
	BaseResponse
	Response []objects.GroupTag `json:"response"`
}

type GroupsGetTokenPermissionsResponse struct {
	BaseResponse
	Response objects.GroupTokenPermissions `json:"response"`
}

type GroupsInviteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsIsMemberResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsIsMemberExtendedResponse struct {
	BaseResponse
	Response struct {
		Invitation objects.BoolInt `json:"invitation"` // Information whether user has been invited to the group
		Member     objects.BoolInt `json:"member"`     // Information whether user is a member of the group
		Request    objects.BoolInt `json:"request"`    // Information whether user has send request to the group
		CanInvite  objects.BoolInt `json:"can_invite"` // Information whether user can be invite
		CanRecall  objects.BoolInt `json:"can_recall"` // Information whether user's invite to the group can be recalled
	} `json:"response"`
}
type GroupsIsMemberUserIDsExtendedResponse struct {
	BaseResponse
	Response []objects.GroupMemberStatusFull `json:"response"`
}

type GroupsIsMemberUserIDsResponse struct {
	BaseResponse
	Response []objects.GroupMemberStatus `json:"response"`
}

type GroupsJoinResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsLeaveResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsRemoveUserResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsReorderLinkResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsSearchResponse struct {
	BaseResponse
	Response struct {
		Count int                 `json:"count"`
		Items []objects.GroupFull `json:"items"`
	} `json:"response"`
}

type GroupsSetCallbackSettingsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsSetLongPollSettingsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsSetSettingsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsSetUserNoteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsTagAddResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsTagBindResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsTagDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsTagUpdateResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsToggleMarketResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type GroupsUnbanResponse struct {
	BaseResponse
	Response int `json:"response"`
}
