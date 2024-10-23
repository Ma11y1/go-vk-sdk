package errors

const (
	ValidationSMS string = "2fa_sms"
	ValidationApp string = "2fa_app"
)

const (
	AuthInvalidRequest          string = "invalid_request"
	AuthUnauthorizedClient      string = "unauthorized_client" // app is not allowed to request an authorization code.
	AuthUnsupportedResponseType string = "unsupported_response_type"
	AuthInvalidScope            string = "invalid_scope"
	AuthInvalidGrant            string = "invalid_grant"
	AuthInvalidClient           string = "invalid_client"
	AuthTemporarilyUnavailable  string = "temporarily_unavailable"
	AuthAccessDenied            string = "access_denied"
	AuthNeedValidation          string = "need_validation"
	AuthNeedCaptcha             string = "need_captcha"
	AuthUserDenied              string = "user_denied"
)

// Doc: https://dev.vk.com/ru/reference/s
const (
	None                                              int = 0
	UnknownCode                                       int = 1    // Unknown s occurred
	DisabledCode                                      int = 2    // App is disabled, enable your app or use test mode (test_mode=1)
	MethodCode                                        int = 3    // Unknown method passed. http://vk.com/dev/methods
	SignatureCode                                     int = 4    // Incorrect signature
	AuthCode                                          int = 5    // Auth s
	TooManyRequestCode                                int = 6    // Too many requests per second
	PermissionCode                                    int = 7    // Permission to perform this action is denied
	InvalidRequestCode                                int = 8    // Invalid query
	FloodCode                                         int = 9    // Flood control
	InternalServerCode                                int = 10   // Internal server s
	EnabledInTestCode                                 int = 11   // In test mode application should be disabled or user should be authorized
	CompileCode                                       int = 12   // Unable to compile code
	RuntimeCode                                       int = 13   // Runtime s occurred during code invocation
	CaptchaCode                                       int = 14   // Captcha needed
	AccessCode                                        int = 15   // Access denied
	AuthHTTPSCode                                     int = 16   // HTTP authorization failed, use secure connection
	AuthValidationCode                                int = 17   // Validation required. https://dev.vk.com/ru/api/validation-required-
	UserDeletedCode                                   int = 18   // OAuthUser was deleted or banned. https://dev.vk.com/ru/api/validation-required-
	BlockedCode                                       int = 19   // Content blocked. https://dev.vk.com/ru/api/validation-required-
	MethodPermissionCode                              int = 20   // Permission to perform this action is denied for non-standalone applications
	MethodAdsCode                                     int = 21   // Permission to perform this action is allowed only for standalone and OpenAPI applications
	UploadCode                                        int = 22   // Upload s. Permission to perform this action is allowed only for standalone and OpenAPI applications
	MethodDisabledCode                                int = 23   // This method was disabled. http://vk.com/dev/methods
	NeedConfirmationCode                              int = 24   // Confirmation required. https://dev.vk.com/ru/api/confirmation-required-
	NeedTokenConfirmationCode                         int = 25   // Confirmation required. AccessToken confirmation required. https://dev.vk.com/ru/api/confirmation-required-
	GroupAuthCode                                     int = 27   // Confirmation required. Group authorization failed. https://dev.vk.com/ru/api/confirmation-required-
	AppAuthCode                                       int = 28   // Confirmation required. Application authorization failed. https://dev.vk.com/ru/api/confirmation-required-
	RateLimitCode                                     int = 29   // Rate limit reached. https://dev.vk.com/ru/reference/roadmap
	PrivateProfileCode                                int = 30   // This profile is private Rate limit reached. https://dev.vk.com/ru/reference/roadmap
	ClientVersionDeprecatedCode                       int = 34   // TransportClient version deprecated
	ExecutionTimeoutCode                              int = 36   // Method execution was interrupted due to timeout
	UserBannedCode                                    int = 37   // OAuthUser was banned
	UnknownApplicationCode                            int = 38   // Unknown application
	UnknownUserCode                                   int = 39   // Unknown user
	UnknownGroupCode                                  int = 40   // Unknown group
	AdditionalSignupRequiredCode                      int = 41   // Additional signup required
	IPNotAllowedCode                                  int = 42   // IP is not allowed
	ParamCode                                         int = 100  // One of the parameters specified was missing or invalid
	ParamAPIIDCode                                    int = 101  // Invalid application api ID. http://vk.com/apps?act=settings
	LimitsCode                                        int = 103  // Invalid application api ID. Out of limits
	NotFoundCode                                      int = 104  // Invalid application api ID. Not found
	SaveFileCode                                      int = 105  // Invalid application api ID. Couldn't save file
	ActionFailedCode                                  int = 106  // Invalid application api ID. Unable to process action
	ParamUserIDCode                                   int = 113  // Invalid user id
	ParamAlbumIDCode                                  int = 114  // Invalid user id. Invalid album id
	ParamServerCode                                   int = 118  // Invalid user id. Invalid server
	ParamTitleCode                                    int = 119  // Invalid user id. Invalid title
	ParamPhotosCode                                   int = 122  // Invalid user id. Invalid photos
	ParamHashCode                                     int = 121  // Invalid user id. Invalid hash
	ParamPhotoCode                                    int = 129  // Invalid user id. Invalid photo
	ParamGroupIDCode                                  int = 125  // Invalid user id. Invalid group id
	ParamPageIDCode                                   int = 140  // Invalid user id. Page not found
	AccessPageCode                                    int = 141  // Invalid user id. Access to page denied
	MobileNotActivatedCode                            int = 146  // The mobile number of the user is unknown
	InsufficientFundsCode                             int = 147  // Application has insufficient funds
	AccessMenuCode                                    int = 148  // Access to the menu of the user denied
	ParamTimestampCode                                int = 150  // Invalid timestamp
	FriendsListIDCode                                 int = 171  // Invalid list id
	FriendsListLimitCode                              int = 173  // Reached the maximum number of lists
	FriendsAddYourselfCode                            int = 174  // Cannot add user himself as friend
	FriendsAddInEnemyCode                             int = 175  // Cannot add this user to friends as they have put you on their blacklist
	FriendsAddEnemyCode                               int = 176  // Cannot add this user to friends as you put him on blacklist
	FriendsAddNotFoundCode                            int = 177  // Cannot add this user to friends as user not found
	ParamNoteIDCode                                   int = 180  // Cannot add this user to friends as user not found. Note not found
	AccessNoteCode                                    int = 181  // Cannot add this user to friends as user not found. Access to note denied
	AccessNoteCommentCode                             int = 182  // Cannot add this user to friends as user not found. You can't comment this note
	AccessCommentCode                                 int = 183  // Cannot add this user to friends as user not found. Access to comment denied
	AccessAlbumCode                                   int = 200  // Access to album denied
	AccessAudioCode                                   int = 201  // Access to audio denied
	AccessGroupCode                                   int = 203  // Access to group denied
	AccessVideoCode                                   int = 204  // Access denied
	AccessMarketCode                                  int = 205  // Access denied
	WallAccessPostCode                                int = 210  // Access to wall's post denied
	WallAccessCommentCode                             int = 211  // Access to wall's comment denied
	WallAccessRepliesCode                             int = 212  // Access to post comments denied
	WallAccessAddReplyCode                            int = 213  // Access to status replies denied
	WallAddPostCode                                   int = 214  // Access to adding post denied
	WallAdsPublishedCode                              int = 219  // Advertisement post was recently added
	WallTooManyRecipientsCode                         int = 220  // Too many recipients
	StatusNoAudioCode                                 int = 221  // OAuthUser disabled track name broadcast
	WallLinksForbiddenCode                            int = 222  // Hyperlinks are forbidden
	WallReplyOwnerFloodCode                           int = 223  // Too many replies
	WallAdsPostLimitReachedCode                       int = 224  // Too many ads posts
	DonutDisabledCode                                 int = 225  // Donut is disabled
	LikesReactionCanNotBeAppliedCode                  int = 232  // Reaction can not be applied to the temp
	PollsAccessCode                                   int = 250  // Access to poll denied
	PollsAnswerIDCode                                 int = 252  // Invalid answer id
	PollsPollIDCode                                   int = 251  // Invalid poll id
	PollsAccessWithoutVoteCode                        int = 253  // Access denied, please vote first
	AccessGroupsCode                                  int = 260  // Access to the groups list is denied due to the user's privacy settings
	AlbumFullCode                                     int = 300  // This album is full
	AlbumsLimitCode                                   int = 302  // This album is full. Albums number limit is reached
	VotesPermissionCode                               int = 500  // Permission denied. You must enable votes processing in application. http://vk.com/editapp?id={API_ID}&section=payments
	VotesCode                                         int = 503  // Not enough votes
	NotEnoughMoneyCode                                int = 504  // Not enough money on owner's balance
	AdsPermissionCode                                 int = 600  // Permission denied. You have no access to operations specified with given temp
	WeightedFloodCode                                 int = 601  // Permission denied. You have requested too many action this day. Try later
	AdsPartialSuccessCode                             int = 602  // Some part of the query has not been completed
	AdsSpecificCode                                   int = 603  // Some ads s occurred
	AdsDomainInvalidCode                              int = 604  // Invalid domain
	AdsDomainForbiddenCode                            int = 605  // Domain is forbidden
	AdsDomainReservedCode                             int = 606  // Domain is reserved
	AdsDomainOccupiedCode                             int = 607  // Domain is occupied
	AdsDomainActiveCode                               int = 608  // Domain is active
	AdsDomainAppInvalidCode                           int = 609  // Domain app is invalid
	AdsDomainAppForbiddenCode                         int = 610  // Domain app is forbidden
	AdsApplicationMustBeVerifiedCode                  int = 611  // Application must be verified
	AdsApplicationMustBeInDomainsListCode             int = 612  // Application must be in domains list of site of ad unit
	AdsApplicationBlockedCode                         int = 613  // Application is blocked
	AdsDomainTypeForbiddenInCurrentOfficeCode         int = 614  // Domain of type specified is forbidden in current office type
	AdsDomainGroupInvalidCode                         int = 615  // Domain group is invalid
	AdsDomainGroupForbiddenCode                       int = 616  // Domain group is forbidden
	AdsDomainAppBlockedCode                           int = 617  // Domain app is blocked AdsDomainGroupNotOpen int = 618 	// Domain group is not open
	AdsDomainGroupNotPossibleJoinedCode               int = 619  // Domain group is not possible to be joined to adsweb
	AdsDomainGroupBlockedCode                         int = 620  // Domain group is blocked
	AdsDomainGroupLinksForbiddenCode                  int = 621  // Domain group has restriction: links are forbidden
	AdsDomainGroupExcludedFromSearchCode              int = 622  // Domain group has restriction: excluded from search
	AdsDomainGroupCoverForbiddenCode                  int = 623  // Domain group has restriction: cover is forbidden
	AdsDomainGroupWrongCategoryCode                   int = 624  // Domain group has wrong category
	AdsDomainGroupWrongNameCode                       int = 625  // Domain group has wrong name
	AdsDomainGroupLowPostsReachCode                   int = 626  // Domain group has low posts reach
	AdsDomainGroupWrongClassCode                      int = 627  // Domain group has wrong class
	AdsDomainGroupCreatedRecentlyCode                 int = 628  // Domain group is created recently
	AdsObjectDeletedCode                              int = 629  // Object deleted
	AdsLookalikeRequestAlreadyInProgressCode          int = 630  // Lookalike query with same source already in progress
	AdsLookalikeRequestsLimitCode                     int = 631  // Max count of lookalike requests per day reached
	AdsAudienceTooSmallCode                           int = 632  // Given audience is too small
	AdsAudienceTooLargeCode                           int = 633  // Given audience is too large
	AdsLookalikeAudienceSaveAlreadyInProgressCode     int = 634  // Lookalike query audience save already in progress
	AdsLookalikeSavesLimitCode                        int = 635  // Max count of lookalike query audience saves per day reached
	AdsRetargetingGroupsLimitCode                     int = 636  // Max count of retargeting groups reached
	AdsDomainGroupActiveNemesisPunishmentCode         int = 637  // Domain group has active nemesis punishment
	GroupChangeCreatorCode                            int = 700  // Cannot edit creator role
	GroupNotInClubCode                                int = 701  // OAuthUser should be in club
	GroupTooManyOfficersCode                          int = 702  // Too many officers in club
	GroupNeed2faCode                                  int = 703  // You need to enable 2FA for this action
	GroupHostNeed2faCode                              int = 704  // OAuthUser needs to enable 2FA for this action
	GroupTooManyAddressesCode                         int = 706  // Too many addresses in club
	GroupAppIsNotInstalledInCommunityCode             int = 711  // Application is not installed in community
	GroupInvalidInviteLinkCode                        int = 714  // Invite link is invalid - expired, deleted or not exists
	VideoAlreadyAddedCode                             int = 800  // This video is already added
	VideoCommentsClosedCode                           int = 801  // Comments for this video are closed
	MessagesUserBlockedCode                           int = 900  // Can't send messages for users from blacklist
	MessagesDenySendCode                              int = 901  // Can't send messages for users without permission
	MessagesPrivacyCode                               int = 902  // Can't send messages to this user due to their privacy settings
	MessagesTooOldPtsCode                             int = 907  // Value of ts or pts is too old
	MessagesTooNewPtsCode                             int = 908  // Value of ts or pts is too new
	MessagesEditExpiredCode                           int = 909  // Can't edit this message, because it's too old
	MessagesTooBigCode                                int = 910  // Can't sent this message, because it's too big
	MessagesKeyboardInvalidCode                       int = 911  // Keyboard format is invalid
	MessagesChatBotFeatureCode                        int = 912  // This is a chat bot feature, change this status in settings
	MessagesTooLongForwardsCode                       int = 913  // Too many forwarded messages
	MessagesTooLongMessageCode                        int = 914  // EventType is too long
	MessagesChatUserNoAccessCode                      int = 917  // You don't have access to this chat
	MessagesCantSeeInviteLinkCode                     int = 919  // You can't see invite link for this chat
	MessagesEditKindDisallowedCode                    int = 920  // Can't edit this kind of message
	MessagesCantFwdCode                               int = 921  // Can't forward these messages
	MessagesCantDeleteForAllCode                      int = 924  // Can't delete this message for everybody
	MessagesChatNotAdminCode                          int = 925  // You are not admin of this chat
	MessagesChatNotExistCode                          int = 927  // Chat does not exist
	MessagesCantChangeInviteLinkCode                  int = 931  // You can't change invite link for this chat
	MessagesGroupPeerAccessCode                       int = 932  // Your community can't interact with this peer
	MessagesChatUserNotInChatCode                     int = 935  // OAuthUser not found in chat
	MessagesContactNotFoundCode                       int = 936  // Contact not found
	MessagesMessageRequestAlreadySendCode             int = 939  // EventType query already send
	MessagesTooManyPostsCode                          int = 940  // Too many posts in messages
	MessagesCantPinOneTimeStoryCode                   int = 942  // Cannot pin one-time story
	MessagesCantUseIntentCode                         int = 943  // Cannot use this intent
	MessagesLimitIntentCode                           int = 944  // Limits overflow for this intent
	MessagesChatDisabledCode                          int = 945  // Chat was disabled
	MessagesChatNotSupportedCode                      int = 946  // Chat not support
	MessagesMemberAccessToGroupDeniedCode             int = 947  // Can't add user to chat, because user has no access to group
	MessagesEditPinnedCode                            int = 949  // Can't edit pinned message yet
	MessagesReplyTimedOutCode                         int = 950  // Can't send message, reply timed out
	MessagesAccessDonutChatCode                       int = 962  // You can't access donut chat without subscription
	MessagesAccessWorkChatCode                        int = 967  // This user can't be added to the work chat, as they aren't an employe
	MessagesCantForwardedCode                         int = 969  // EventType cannot be forwarded
	MessagesPinExpiringMessageCode                    int = 970  // Cannot pin an expiring message
	ParamPhoneCode                                    int = 1000 // Invalid phone number
	PhoneAlreadyUsedCode                              int = 1004 // This phone number is used by another user
	AuthFloodCode                                     int = 1105 // Too many auth attempts, try again later
	AuthDelayCode                                     int = 1112 // Processing. Try later
	AnonymousTokenExpiredCode                         int = 1114 // Anonymous token has expired
	AnonymousTokenInvalidCode                         int = 1116 // Anonymous token is invalid
	AuthAccessTokenHasExpiredCode                     int = 1117 // Access token has expired
	AuthAnonymousTokenIPMismatchCode                  int = 1118 // Anonymous token ip mismatch
	ParamDocIDCode                                    int = 1150 // Invalid document id
	ParamDocDeleteAccessCode                          int = 1151 // Access to document deleting is denied
	ParamDocTitleCode                                 int = 1152 // Invalid document title
	ParamDocAccessCode                                int = 1153 // Access to document is denied
	PhotoChangedCode                                  int = 1160 // Original photo was changed
	TooManyListsCode                                  int = 1170 // Too many feed lists
	AppsAlreadyUnlockedCode                           int = 1251 // This achievement is already unlocked
	AppsSubscriptionNotFoundCode                      int = 1256 // Subscription not found
	AppsSubscriptionInvalidStatusCode                 int = 1257 // Subscription is in invalid status
	InvalidAddressCode                                int = 1260 // Invalid screen name
	CommunitiesCatalogDisabledCode                    int = 1310 // Catalog is not available for this user
	CommunitiesCategoriesDisabledCode                 int = 1311 // Catalog categories are not available for this user
	MarketRestoreTooLateCode                          int = 1400 // Too late for restore
	MarketCommentsClosedCode                          int = 1401 // Comments for this market are closed
	MarketAlbumNotFoundCode                           int = 1402 // Album not found
	MarketItemNotFoundCode                            int = 1403 // Item not found
	MarketItemAlreadyAddedCode                        int = 1404 // Item already added to album
	MarketTooManyItemsCode                            int = 1405 // Too many items
	MarketTooManyItemsInAlbumCode                     int = 1406 // Too many items in album
	MarketTooManyAlbumsCode                           int = 1407 // Too many albums
	MarketItemHasBadLinksCode                         int = 1408 // Item has bad links in description
	MarketShopNotEnabledCode                          int = 1409 // Extended market not enabled
	MarketGroupingItemsWithDifferentPropertiesCode    int = 1412 // Grouping items with different properties
	MarketGroupingAlreadyHasSuchVariantCode           int = 1413 // Grouping already has such variant
	MarketVariantNotFoundCode                         int = 1416 // Variant not found
	MarketPropertyNotFoundCode                        int = 1417 // Property not found
	MarketGroupingMustContainMoreThanOneItemCode      int = 1425 // Grouping must have two or more items
	MarketGroupingItemsMustHaveDistinctPropertiesCode int = 1426 // Item must have distinct properties
	MarketOrdersNoCartItemsCode                       int = 1427 // Cart is empty
	MarketInvalidDimensionsCode                       int = 1429 // Specify width, length, height and weight all together
	MarketCantChangeVkpayStatusCode                   int = 1430 // VK Pay status can not be changed
	MarketShopAlreadyEnabledCode                      int = 1431 // Market was already enabled in this group
	MarketShopAlreadyDisabledCode                     int = 1432 // Market was already disabled in this group
	MarketPhotosCropInvalidFormatCode                 int = 1433 // Invalid image crop format
	MarketPhotosCropOverflowCode                      int = 1434 // Crop bottom right corner is outside of the image
	MarketPhotosCropSizeTooLowCode                    int = 1435 // Crop size is less than the minimum
	MarketNotEnabledCode                              int = 1438 // Market not enabled
	MarketCartEmptyCode                               int = 1427 // Cart is empty
	MarketSpecifyDimensionsCode                       int = 1429 // Specify width, length, height and weight all together
	VKPayStatusCode                                   int = 1430 // VK Pay status can not be changed
	MarketAlreadyEnabledCode                          int = 1431 // Market was already enabled in this group
	MarketAlreadyDisabledCode                         int = 1432 // Market was already disabled in this group
	MainAlbumCantHiddenCode                           int = 1446 // Main album can not be hidden
	StoryExpiredCode                                  int = 1600 // Story has already expired
	StoryIncorrectReplyPrivacyCode                    int = 1602 // Incorrect reply privacy
	PrettyCardsCardNotFoundCode                       int = 1900 // Card not found
	PrettyCardsTooManyCardsCode                       int = 1901 // Too many cards
	PrettyCardsCardIsConnectedToPostCode              int = 1902 // Card is connected to post
	CallbackServersLimitCode                          int = 2000 // Servers number limit is reached
	StickersNotPurchasedCode                          int = 2100 // Stickers are not purchased
	StickersTooManyFavoritesCode                      int = 2101 // Too many favorite stickers
	StickersNotFavoriteCode                           int = 2102 // Stickers are not favorite
	WallCheckLinkCantDetermineSourceCode              int = 3102 // Specified link is incorrect (can't find source)
	RecaptchaCode                                     int = 3300 // Recaptcha needed
	PhoneValidationCode                               int = 3301 // Phone validation needed
	PasswordValidationCode                            int = 3302 // Password validation needed
	OtpAppValidationCode                              int = 3303 // Otp app validation needed
	EmailConfirmationCode                             int = 3304 // Email confirmation needed
	AssertVotesCode                                   int = 3305 // Assert votes
	TokenExtensionCode                                int = 3609 // AccessToken extension required
	UserDeactivatedCode                               int = 3610 // OAuthUser is deactivated
	ServiceDeactivatedCode                            int = 3611 // Service is deactivated for user
	AliExpressTagCode                                 int = 3800 // Can't set AliExpress tag to this code of temp
	InvalidUploadResponseCode                         int = 5701 // Invalid upload response
	InvalidUploadHashCode                             int = 5702 // Invalid upload hash
	InvalidUploadUserCode                             int = 5703 // Invalid upload user
	InvalidUploadGroupCode                            int = 5704 // Invalid upload group
	InvalidCropDataCode                               int = 5705 // Invalid crop data
	ToSmallAvatarCode                                 int = 5706 // To small avatar
	PhotoNotFoundCode                                 int = 5708 // Photo not found
	InvalidPhotoCode                                  int = 5709 // Invalid Photo
	InvalidHashCode                                   int = 5710 // Invalid hash
)