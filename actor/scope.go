package actor

// Scopes for user auth
const (
	ScopeUserNotify        int = 1 << 0
	ScopeUserFriends       int = 1 << 1
	ScopeUserPhotos        int = 1 << 2
	ScopeUserAudio         int = 1 << 3
	ScopeUserVideo         int = 1 << 4
	ScopeUserStories       int = 1 << 6
	ScopeUserPages         int = 1 << 7
	ScopeUserMenu          int = 1 << 8
	ScopeUserWallmenu      int = 1 << 9
	ScopeUserStatus        int = 1 << 10
	ScopeUserNotes         int = 1 << 11
	ScopeUserMessages      int = 1 << 12
	ScopeUserWall          int = 1 << 13
	ScopeUserAds           int = 1 << 15
	ScopeUserOffline       int = 1 << 16
	ScopeUserDocs          int = 1 << 17
	ScopeUserGroups        int = 1 << 18
	ScopeUserNotifications int = 1 << 19
	ScopeUserStats         int = 1 << 20
	ScopeUserEmail         int = 1 << 22
	ScopeUserAdsweb        int = 1 << 23
	ScopeUserLeads         int = 1 << 24
	ScopeUserGroupMessages int = 1 << 25
	ScopeUserExchange      int = 1 << 26
	ScopeUserMarket        int = 1 << 27
	ScopeUserPhone         int = 1 << 28
)

// Scopes for Group auth
const (
	ScopeGroupStories   int = 1 << 0
	ScopeGroupPhotos    int = 1 << 2
	ScopeGroupAppWidget int = 1 << 6
	ScopeGroupMessages  int = 1 << 12
	ScopeGroupDocs      int = 1 << 17
	ScopeGroupManage    int = 1 << 18
)

// Doc: https://id.vk.com/about/business/go/docs/ru/vkid/latest/vk-id/connection/work-with-user-info/scopes
const (
	ScopeUserNameVKIDPersonalInfo  string = "vkid.personal_info" // Default. Last name, first name, gender, profile photo, date of birth. Basic permission that is used by default for all applications
	ScopeUserNameVKIDEmail         string = "vkid.personal_info"
	ScopeUserNameVKIDPhone         string = "vkid.personal_info"
	ScopeUserNameVKIDFriends       string = "vkid.personal_info"
	ScopeUserNameVKIDWall          string = "vkid.personal_info"
	ScopeUserNameVKIDGroups        string = "vkid.personal_info"
	ScopeUserNameVKIDStories       string = "vkid.personal_info"
	ScopeUserNameVKIDDocs          string = "vkid.personal_info"
	ScopeUserNameVKIDPhotos        string = "vkid.personal_info"
	ScopeUserNameVKIDAds           string = "vkid.personal_info"
	ScopeUserNameVKIDVideo         string = "vkid.personal_info"
	ScopeUserNameVKIDStatus        string = "vkid.personal_info"
	ScopeUserNameVKIDMarket        string = "vkid.personal_info"
	ScopeUserNameVKIDPages         string = "vkid.personal_info"
	ScopeUserNameVKIDNotifications string = "vkid.personal_info"
	ScopeUserNameVKIDStats         string = "vkid.personal_info"
	ScopeUserNameVKIDNotes         string = "vkid.personal_info"
)
