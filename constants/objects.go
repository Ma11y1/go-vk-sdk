package constants

// Base sex
const (
	SexUnknown = iota
	SexFemale
	SexMale
)

// Privacy
const (
	PrivacyAll              string = "all"
	PrivacyOnlyMe           string = "only_me"
	PrivacyFriends          string = "friends"
	PrivacyFriendsOfFriends string = "friends_of_friends"
)

// Platform types
type Platform int

const (
	_                    Platform = iota
	PlatformMobile                // mobile web version
	PlatformIPhone                // iPhone
	PlatformIPad                  // iPad
	PlatformAndroid               // Android
	PlatformWindowsPhone          // Windows Phone
	PlatformWindows               // Windows 8
	PlatformFull                  // full web version
	PlatformOther                 // other apps
)

const (
	UserRelationNotSpecified      = iota // not specified
	UserRelationSingle                   // single
	UserRelationInRelationship           // in a relationship
	UserRelationEngaged                  // engaged
	UserRelationMarried                  // married
	UserRelationComplicated              // complicated
	UserRelationActivelySearching        // actively searching
	UserRelationInLove                   // in love
	UserRelationCivilUnion               // in a civil union
)

// Button action type.
const (
	ButtonText     string = "text"      // A button that sends a message with text specified in the label.
	ButtonVKPay    string = "vkpay"     // Opens the VK Pay window with predefined parameters. The button is called “Pay with VK Pay” (VK Pay is displayed as a logo). This button always stretches to the whole keyboard width.
	ButtonVKApp    string = "open_app"  // Opens a specified VK Apps app. This button always stretches to the whole keyboard width.
	ButtonLocation string = "location"  // Sends the location to the chat. This button always stretches to the whole keyboard width.
	ButtonOpenLink string = "open_link" // Opens the specified link.
	ButtonCallback string = "callback"  // Allows, without sending a message from the user, to receive a notification about pressing the button and perform the necessary action.
)

// Button color. This parameter is used only for buttons with the text and callback types.
const (
	Primary     string = "primary"
	Secondary   string = "secondary"
	Negative    string = "negative"
	Positive    string = "positive"
	ButtonBlue  string = Primary   // Blue button, indicates the main action. #5181B8
	ButtonWhite string = Secondary // Default white button. #FFFFFF
	ButtonRed   string = Negative  // Dangerous or negative action (cancel, delete etc.) #E64646
	ButtonGreen string = Positive  // Accept, agree. #4BB34B
)

// Conversations types.
const (
	PeerUser  string = "user"
	PeerChat  string = "chat"
	PeerGroup string = "group"
	PeerEmail string = "email"
)

// Ads
type AdsLinkStatus string

const (
	AdsLinkAllowed    AdsLinkStatus = "allowed"
	AdsLinkDisallowed AdsLinkStatus = "disallowed"
	AdsLinkInProgress AdsLinkStatus = "in_progress"
)

type AdsUserSpecificationRole string

const (
	AdsUserSpecificationRoleReports AdsUserSpecificationRole = "reports"
	AdsUserSpecificationRoleManager AdsUserSpecificationRole = "manager"
)

type AdsLinkType string

const (
	AdsLinkTypeCommunity AdsLinkType = "community"
	AdsLinkTypePost      AdsLinkType = "post"
	AdsLinkTypeVKApp     AdsLinkType = "application"
	AdsLinkTypeVideo     AdsLinkType = "video"
	AdsLinkTypeSite      AdsLinkType = "site"
)

type AdsCampaignType string

const (
	AdsCampaignTypeNormal        AdsCampaignType = "normal"
	AdsCampaignTypeVKAppsManaged AdsCampaignType = "vk_apps_managed "
	AdsCampaignTypeMobileApps    AdsCampaignType = "mobile_apps "
	AdsCampaignTypePromotedPosts AdsCampaignType = "promoted_posts "
	AdsCampaignTypeAdaptiveAds   AdsCampaignType = "adaptive_ads "
)

type AdsLookalikeRequestStatus string

const (
	AdsLookalikeRequestStatusSearchInProgress AdsLookalikeRequestStatus = "search_in_progress"
	AdsLookalikeRequestStatusSearchDone       AdsLookalikeRequestStatus = "search_done"
	AdsLookalikeRequestStatusSearchFailed     AdsLookalikeRequestStatus = "search_failed"
	AdsLookalikeRequestStatusSaveInProgress   AdsLookalikeRequestStatus = "save_in_progress"
	AdsLookalikeRequestStatusSaveDone         AdsLookalikeRequestStatus = "save_done"
	AdsLookalikeRequestStatusSaveFailed       AdsLookalikeRequestStatus = "save_failed"
)

type AdsStatisticIDType string

const (
	AdsStatisticIDTypeAd       AdsStatisticIDType = "ad"
	AdsStatisticIDTypeCampaign AdsStatisticIDType = "campaign"
	AdsStatisticIDTypeClient   AdsStatisticIDType = "client"
	AdsStatisticIDTypeOffice   AdsStatisticIDType = "office"
)

type AdsStatisticPeriod string

const (
	AdsStatisticPeriodDay     AdsStatisticPeriod = "day"
	AdsStatisticPeriodWeek    AdsStatisticPeriod = "week"
	AdsStatisticPeriodMonth   AdsStatisticPeriod = "month"
	AdsStatisticPeriodYear    AdsStatisticPeriod = "year"
	AdsStatisticPeriodOverall AdsStatisticPeriod = "overall "
)

type AdsTargetingStatisticAdFormat int

const (
	AdsTargetingStatisticAdFormatImgAndText    AdsTargetingStatisticAdFormat = 1
	AdsTargetingStatisticAdFormatBigImg        AdsTargetingStatisticAdFormat = 2
	AdsTargetingStatisticAdFormatExclusive     AdsTargetingStatisticAdFormat = 3
	AdsTargetingStatisticAdFormatSquareImage   AdsTargetingStatisticAdFormat = 4
	AdsTargetingStatisticAdFormatSpecialApp    AdsTargetingStatisticAdFormat = 7
	AdsTargetingStatisticAdFormatSpecialGroups AdsTargetingStatisticAdFormat = 8
	AdsTargetingStatisticAdFormatCommunityPost AdsTargetingStatisticAdFormat = 9
	AdsTargetingStatisticAdFormatAppShowcase   AdsTargetingStatisticAdFormat = 10
	AdsTargetingStatisticAdFormatAdaptive      AdsTargetingStatisticAdFormat = 11
)

type AdsUploadURLAdFormat int

const (
	AdsUploadURLAdFormatImgAndText  AdsUploadURLAdFormat = 1
	AdsUploadURLAdFormatBigImg      AdsUploadURLAdFormat = 2
	AdsUploadURLAdFormatExclusive   AdsUploadURLAdFormat = 3
	AdsUploadURLAdFormatSquareImage AdsUploadURLAdFormat = 4
	AdsUploadURLAdFormatNewsFeedApp AdsUploadURLAdFormat = 5
	AdsUploadURLAdFormatMobileApp   AdsUploadURLAdFormat = 6
	AdsUploadURLAdFormatAdaptive    AdsUploadURLAdFormat = 11
)

// AdsAgeRestriction Display the age limit mark on the ad
type AdsAgeRestriction int

const (
	AdsAgeRestrictionNone AdsAgeRestriction = iota
	AdsAgeRestriction0
	AdsAgeRestriction6
	AdsAgeRestriction12
	AdsAgeRestriction16
	AdsAgeRestriction18
)

// AppWidgets Doc: https://dev.vk.com/ru/method/appWidgets

type AppWidgetsImageType string

const (
	AppWidgetsImageType2424    AppWidgetsImageType = "24x24"
	AppWidgetsImageType50x50   AppWidgetsImageType = "50x50"
	AppWidgetsImageType160x160 AppWidgetsImageType = "160x160"
	AppWidgetsImageType160x240 AppWidgetsImageType = "160x240"
	AppWidgetsImageType510x128 AppWidgetsImageType = "510x128"
)

// Asr Doc: https://dev.vk.com/ru/method/asr

type AsrStatus string

const (
	AsrStatusProcessing       AsrStatus = "processing"        // the audio recording is being processed.
	AsrStatusFinished         AsrStatus = "finished"          // processing of the audio recording is finished.
	AsrStatusInternalError    AsrStatus = "internal_error"    // internal errors of the VK speech recognition service.
	AsrStatusTranscodingError AsrStatus = "transcoding_error" // error in transcoding the audio recording into the internal format. Try downloading the audio in a different supported format.
	AsrStatusRecognitionError AsrStatus = "recognition_error" // speech recognition error, difficulty in recognition. Try speaking more clearly or reducing background noise.
)

type AsrProcessModel string

const (
	AsrModelNeutral     AsrProcessModel = "neutral"
	AsrModelSpontaneous AsrProcessModel = "spontaneous"
)

// App Doc: https://dev.vk.com/ru/reference/objects/app
const (
	AppTypeApp          = "app"
	AppTypeGame         = "game"
	AppTypeSite         = "site"
	AppTypeStandalone   = "standalone"
	AppTypeVkApp        = "vk_app"
	AppTypeCommunityApp = "community_app"
	AppTypeHTML5Game    = "html5_game"
)

const (
	AppsLeaderboardTypeNotSupported = iota
	AppsLeaderboardTypeLevels
	AppsLeaderboardTypePoints
)

type AppsScreenOrientation int

const (
	AppScreenOrientationBoth AppsScreenOrientation = iota
	AppScreenOrientationLandscape
	AppScreenOrientationPortrait
)

// Friends status
const (
	FriendStatusNotFriend        = iota // not a friend
	FriendStatusOutComingRequest        // out coming request
	FriendStatusInComingRequest         // incoming request
	FriendStatusIsFriend                // is friend
)

// Gifts
const (
	GiftPrivacyForAll = iota
	GiftPrivacyNameForAll
	GiftPrivacyRecipientOnly
)

// Group
const (
	GroupWorkStatusNoInformation     = "no_information"
	GroupWorkStatusTemporarilyClosed = "temporarily_closed"
	GroupWorkStatusAlwaysOpened      = "always_opened"
	GroupWorkStatusTimetable         = "timetable"
	GroupWorkStatusForeverClosed     = "forever_closed"
)

const (
	GroupAdminLevelModerator = iota
	GroupAdminLevelEditor
	GroupAdminLevelAdministrator
)

const (
	GroupMainSectionAbsent = iota
	GroupMainSectionPhotos
	GroupMainSectionTopics
	GroupMainSectionAudio
	GroupMainSectionVideo
	GroupMainSectionMarket
)

const (
	GroupMemberStatusNotMember = iota
	GroupMemberStatusMember
	GroupMemberStatusNotSure
	GroupMemberStatusDeclined
	GroupMemberStatusHasSentRequest
	GroupMemberStatusInvited
)

const (
	GroupGroupOpen = iota
	GroupGroupClosed
	GroupGroupPrivate
)

const (
	GroupAgeLimitsNo = iota
	GroupAgeLimitsOver16
	GroupAgeLimitsOver18
)

const (
	GroupTypeGroup = "group"
	GroupTypePage  = "page"
	GroupTypeEvent = "event"
)
const (
	GroupBanReasonOther = iota
	GroupBanReasonSpam
	GroupBanReasonVerbalAbuse
	GroupBanReasonStrongLanguage
	GroupBanReasonFlood
)
const (
	GroupGroupPhotosDisabled = iota
	GroupGroupPhotosOpen
	GroupGroupPhotosLimited
)

const (
	_ = iota
	GroupSubjectAuto
	GroupSubjectActivityHolidays
	GroupSubjectBusiness
	GroupSubjectPets
	GroupSubjectHealth
	GroupSubjectDatingAndCommunication
	GroupSubjectGames
	GroupSubjectIt
	GroupSubjectCinema
	GroupSubjectBeautyAndFashion
	GroupSubjectCooking
	GroupSubjectArtAndCulture
	GroupSubjectLiterature
	GroupSubjectMobileServicesAndInternet
	GroupSubjectMusic
	GroupSubjectScienceAndTechnology
	GroupSubjectRealEstate
	GroupSubjectNewsAndMedia
	GroupSubjectSecurity
	GroupSubjectEducation
	GroupSubjectHomeAndRenovations
	GroupSubjectPolitics
	GroupSubjectFood
	GroupSubjectIndustry
	GroupSubjectTravel
	GroupSubjectWork
	GroupSubjectEntertainment
	GroupSubjectReligion
	GroupSubjectFamily
	GroupSubjectSports
	GroupSubjectInsurance
	GroupSubjectTelevision
	GroupSubjectGoodsAndServices
	GroupSubjectHobbies
	GroupSubjectFinance
	GroupSubjectPhoto
	GroupSubjectEsoterics
	GroupSubjectElectronicsAndAppliances
	GroupSubjectErotic
	GroupSubjectHumor
	GroupSubjectSocietyHumanities
	GroupSubjectDesignAndGraphics
)

const (
	GroupTopicsDisabled = iota
	GroupTopicsOpen
	GroupTopicsLimited
)

const (
	GroupDocsDisabled = iota
	GroupDocsOpen
	GroupDocsLimited
)

const (
	GroupAudioDisabled = iota
	GroupAudioOpen
	GroupAudioLimited
)

const (
	GroupVideoDisabled = iota
	GroupVideoOpen
	GroupVideoLimited
)

const (
	GroupWallDisabled = iota
	GroupWallOpen
	GroupWallLimited
	GroupWallClosed
)

const (
	GroupWikiDisabled = iota
	GroupWikiOpen
	GroupWikiLimited
)

type GroupActionType string

const (
	GroupActionTypeOpenURL      GroupActionType = "open_url"
	GroupActionTypeSendEmail    GroupActionType = "send_email"
	GroupActionTypeCallPhone    GroupActionType = "call_phone"
	GroupActionTypeCallVK       GroupActionType = "call_vk"
	GroupActionTypeOpenGroupApp GroupActionType = "open_group_app"
	GroupActionTypeOpenApp      GroupActionType = "open_app"
)

type GroupMarketType string

const (
	GroupMarketBasic    GroupMarketType = "basic"
	GroupMarketAdvanced GroupMarketType = "advanced"
)
const (
	GroupRoleModerator     = "moderator"
	GroupRoleEditor        = "editor"
	GroupRoleAdministrator = "administrator"
	GroupRoleCreator       = "creator"
)
const (
	GroupOnlineStatusTypeNone       = "none"
	GroupOnlineStatusTypeOnline     = "online"
	GroupOnlineStatusTypeAnswerMark = "answer_mark"
)

// Market
type MarketOrderStatus int

const (
	MarketOrderNew MarketOrderStatus = iota
	MarketOrderPending
	MarketOrderProcessing
	MarketOrderShipped
	MarketOrderComplete
	MarketOrderCanceled
	MarketOrderRefund
)

const (
	MarketItemAvailable = iota
	MarketItemRemoved
	MarketItemUnavailable
)

// Orders
type OrderState string

const (
	OrderStateChargeable OrderState = "chargeable" // unconfirmed order. Orders end up in this state if the stoe does not process notifications.
	OrderStateDeclined   OrderState = "declined"   // a canceled order at the stage of obtaining information about the product, for example, error 20 was received, “Product does not exist.” An order can enter this state from the chargeable state.
	OrderStateCancelled  OrderState = "cancelled"  // canceled order. An order can enter this state from the chargeable state.
	OrderStateCharged    OrderState = "charged"    // paid order. An order can enter this state from the chargeable state, or immediately after payments, if the application processes notifications.
	OrderStateRefunded   OrderState = "refunded"   // order canceled after payments, votes returned to the user.
)

// Pages

const (
	PagesPrivacyCommunityManagers = iota // community managers only
	PagesPrivacyCommunityMembers         // community members only
	PagesPrivacyEveryone                 // everyone
)

// Messages

const (
	MessagesCallEndStateCanceledByInitiator string = "canceled_by_initiator"
	MessagesCallEndStateCanceledByReceiver  string = "canceled_by_receiver"
	MessagesCallEndStateReached             string = "reached"
)

const (
	MessagesChatPhotoUpdate        string = "chat_photo_update"
	MessagesChatPhotoRemove        string = "chat_photo_remove"
	MessagesChatCreate             string = "chat_create"
	MessagesChatTitleUpdate        string = "chat_title_update"
	MessagesChatInviteUser         string = "chat_invite_user"
	MessagesChatKickUser           string = "chat_kick_user"
	MessagesChatPinMessage         string = "chat_pin_message"
	MessagesChatUnpinMessage       string = "chat_unpin_message"
	MessagesChatInviteUserByLink   string = "chat_invite_user_by_link"
	MessagesAcceptedMessageRequest string = "accepted_message_request"
)

type MessagesChatPermission string

const (
	MessagesOwnerChatPermission          MessagesChatPermission = "owner"
	MessagesOwnerAndAdminsChatPermission MessagesChatPermission = "owner_and_admins"
	MessagesAllChatPermission            MessagesChatPermission = "all"
)

// Command for MessagesBasePayload.
const (
	MessagesCommandNotSupportedButton string = "not_supported_button"
)

// Stories

type StoriesStoryType string

const (
	StoriesStoryPhoto          StoriesStoryType = "photo"
	StoriesStoryVideo          StoriesStoryType = "video"
	StoriesStoryLiveActive     StoriesStoryType = "live_active"
	StoriesStoryLiveFinished   StoriesStoryType = "live_finished"
	StoriesStoryBirthdayInvite StoriesStoryType = "birthday_invite"
)

type StoriesFeedItemType string

const (
	StoriesFeedItemStories   StoriesFeedItemType = "stories"
	StoriesFeedItemCommunity StoriesFeedItemType = "community_grouped_stories"
	StoriesFeedItemApp       StoriesFeedItemType = "app_grouped_stories"
)

const (
	ClickableStickerPost       string = "post"
	ClickableStickerSticker    string = "sticker"
	ClickableStickerPlace      string = "place"
	ClickableStickerQuestion   string = "question"
	ClickableStickerMention    string = "mention"
	ClickableStickerHashtag    string = "hashtag"
	ClickableStickerMarketItem string = "market_item"
	ClickableStickerLink       string = "link"
	ClickableStickerStoryReply string = "story_reply"
	ClickableStickerOwner      string = "owner"
	ClickableStickerPoll       string = "poll"
	ClickableStickerMusic      string = "music"
	ClickableStickerApp        string = "app"
	ClickableStickerTime       string = "time"
	ClickableStickerEmoji      string = "emoji"
	ClickableStickerGeo        string = "geo"
	ClickableStickerText       string = "text"
)

const (
	ClickableStickerSubtypeMarketItem        string = "market_item"
	ClickableStickerSubtypeAliexpressProduct string = "aliexpress_product"
)

const (
	ClickableStickerTransparent   string = "transparent"
	ClickableStickerBlueGradient  string = "blue_gradient"
	ClickableStickerRedGradient   string = "red_gradient"
	ClickableStickerUnderline     string = "underline"
	ClickableStickerBlue          string = "blue"
	ClickableStickerGreen         string = "green"
	ClickableStickerWhite         string = "white"
	ClickableStickerQuestionReply string = "question_reply"
	ClickableStickerLight         string = "light"
	ClickableStickerImpressive    string = "impressive"
)

// Utils

const (
	UtilsDomainResolvedTypeUser        string = "user"
	UtilsDomainResolvedTypeGroup       string = "group"
	UtilsDomainResolvedTypeApplication string = "application"
	UtilsDomainResolvedTypePage        string = "page"
	UtilsDomainResolvedTypeVkApp       string = "vk_app"
)

const (
	UtilsLinkCheckedStatusNotBanned  = "not_banned"
	UtilsLinkCheckedStatusBanned     = "banned"
	UtilsLinkCheckedStatusProcessing = "processing"
)

// Wall

const (
	WallPostSourceTypeVk     = "vk"
	WallPostSourceTypeWidget = "widget"
	WallPostSourceTypeAPI    = "api"
	WallPostSourceTypeRss    = "rss"
	WallPostSourceTypeSms    = "sms"
)

const (
	WallPostTypePost     = "post"
	WallPostTypeCopy     = "copy"
	WallPostTypeReply    = "reply"
	WallPostTypePostpone = "postpone"
	WallPostTypeSuggest  = "suggest"
)

// Attachment

const (
	AttachmentTypePhoto       = "photo"
	AttachmentTypePostedPhoto = "posted_photo"
	AttachmentTypeAudio       = "audio"
	AttachmentTypeVideo       = "video"
	AttachmentTypeDoc         = "doc"
	AttachmentTypeLink        = "link"
	AttachmentTypeGraffiti    = "graffiti"
	AttachmentTypeNote        = "note"
	AttachmentTypeApp         = "app"
	AttachmentTypePoll        = "poll"
	AttachmentTypePage        = "page"
	AttachmentTypeAlbum       = "album"
	AttachmentTypePhotosList  = "photos_list"
	AttachmentTypeMarketAlbum = "market_album"
	AttachmentTypeMarket      = "market"
	AttachmentTypeEvent       = "event"
	AttachmentTypeWall        = "wall"
	AttachmentTypeStory       = "story"
	AttachmentTypePodcast     = "podcast"
)
