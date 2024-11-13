package constants

const (
	ContentTypeFormURLEncoded string = "application/x-www-form-urlencoded"
	ContentTypeMultipartForm  string = "multipart/form-data"
	ContentTypeJSON           string = "application/json"
	AcceptEncodingGzip        string = "gzip"
)

// The supported values.
const (
	DisplayPage   string = "page"   // authorization form in a separate window
	DisplayPopup  string = "popup"  // a pop-up window
	DisplayMobile string = "mobile" // authorization for mobile devices (uses no Javascript).
)

// Filter for method account.getCounters
const (
	AccountCounterFilterFriends                string = "friends"
	AccountCounterFilterFriendsSuggestions     string = "friends_suggestions"
	AccountCounterFilterMessages               string = "messages"
	AccountCounterFilterPhotos                 string = "photos"
	AccountCounterFilterVideos                 string = "videos"
	AccountCounterFilterGifts                  string = "gifts"
	AccountCounterFilterEvents                 string = "events"
	AccountCounterFilterGroups                 string = "groups"
	AccountCounterFilterNotifications          string = "notifications"
	AccountCounterFilterSdk                    string = "sdk"
	AccountCounterFilterAppRequests            string = "app_requests"
	AccountCounterFilterFriendsRecommendations string = "friends_recommendations"
)

// Name parameters
const (
	ParameterNameVersion                string = "v"                         // type=string
	ParameterNameVersionLP              string = "version"                   // type=int
	ParameterNameLPVersion              string = "lp_version"                // type=int
	ParameterNameLang                   string = "lang"                      // lang=ru type=string
	ParameterNameCaptchaSID             string = "captcha_sid"               // captcha_sid=12345 type=string
	ParameterNameCaptchaKey             string = "captcha_key"               // captcha_key=12345 type=string
	ParameterNameClientID               string = "client_id"                 // type=int
	ParameterNameAccountID              string = "account_id"                // type=int
	ParameterNameData                   string = "data"                      // type=string
	ParameterNameClientSecret           string = "client_secret"             // type=string
	ParameterNameSecretKey              string = "secret_key"                // type=string
	ParameterNameGrandType              string = "grant_type"                // type=string
	ParameterNameAuthorizationCode      string = "authorization_code"        // type=string
	ParameterNameResponseType           string = "response_type"             // type=string
	ParameterNameCode                   string = "code"                      // type=string
	ParameterNameCodeChallenge          string = "code_challenge"            // type=string
	ParameterNameCodeChallengeMethod    string = "code_challenge_method"     // type=string
	ParameterNameDisplay                string = "display"                   // type=string
	ParameterNameUsername               string = "username"                  // type=string
	ParameterNamePassword               string = "password"                  // type=string
	ParameterNameScope                  string = "scope"                     // type=string
	ParameterName2faSupported           string = "2fa_supported"             // type=0, 1
	ParameterNameForceSMS               string = "force_sms"                 // type=0, 1
	ParameterNameTitle                  string = "title"                     //type=string
	ParameterNameTestRedirectURI        string = "test_redirect_uri"         // type=0, 1
	ParameterNameConfirm                string = "confirm"                   // confirm=12345 type=string
	ParameterNameRedirectURI            string = "redirect_uri"              // redirect_uri=URI type=string
	ParameterNameToken                  string = "token"                     // type=string
	ParameterNameOwnerID                string = "owner_id"                  // owner_id=int type=int
	ParameterNameUserID                 string = "user_id"                   // user_id=int type=int, uint
	ParameterNameUserIDs                string = "user_ids"                  // type=string
	ParameterNameRetargetingGroupID     string = "retargeting_group_id"      // user_id=int type=int, uint
	ParameterNameRestoreSID             string = "restore_sid"               // restore_sid=123123123 type=string
	ParameterNameChangePasswordHash     string = "change_password_hash"      // type=string
	ParameterNameOldPassword            string = "old_password"              // type=string
	ParameterNameNewPassword            string = "new_password"              // type=string
	ParameterNameOffset                 string = "offset"                    // type=uint,int
	ParameterNameCount                  string = "count"                     // type=uint,int
	ParameterNameFields                 string = "fields"                    // type=string
	ParameterNameFilter                 string = "filter"                    // type=string
	ParameterNameServerID               string = "server_id"                 // type=uint
	ParameterNameServerIDs              string = "server_ids"                // type=string
	ParameterNameDeviceID               string = "device_id"                 // type=string
	ParameterNameDeviceModel            string = "device_model"              // type=string
	ParameterNameDeviceYear             string = "device_year"               // type=int
	ParameterNameNoText                 string = "no_text"                   // type=0,1
	ParameterNameSystemVersion          string = "system_version"            // type=string
	ParameterNameSubscribe              string = "subscribe"                 // type=string
	ParameterNameSettings               string = "settings"                  // type=string
	ParameterNameState                  string = "state"                     // type=string
	ParameterNameGroupIDs               string = "group_ids"                 // type=string
	ParameterNameSandbox                string = "sandbox"                   // type=uint 0,1
	ParameterNameFirstName              string = "first_name"                // type=string
	ParameterNameLastName               string = "last_name"                 // type=string
	ParameterNameMaidenName             string = "maiden_name"               // type=string
	ParameterNameScreenName             string = "screen_name"               // type=string
	ParameterNameCancelRequestID        string = "cancel_request_id"         // type=uint
	ParameterNameSex                    string = "sex"                       // type=uint
	ParameterNameRelation               string = "relation"                  // type=uint
	ParameterNameRelationPartnerID      string = "relation_partner_id"       // type=int
	ParameterNameBDate                  string = "bdate"                     // type=string
	ParameterNameBDateVisibility        string = "bdate_visibility"          // type=uint
	ParameterNameHomeTown               string = "home_town"                 // type=string
	ParameterNameCountryID              string = "country_id"                // type=uint
	ParameterNameCityID                 string = "city_id"                   // type=uint
	ParameterNameStatus                 string = "status"                    // type=string
	ParameterNameIntro                  string = "intro"                     // type=uint
	ParameterNameOwnPostsDefault        string = "own_posts_default"         // type=0,1
	ParameterNameNoWallReplies          string = "no_wall_replies"           // type=0,1
	ParameterNameName                   string = "name"                      // type=string
	ParameterNameValue                  string = "value"                     // type=string
	ParameterNameVoIP                   string = "voip"                      // type=0,1
	ParameterNameKey                    string = "key"                       //type=string
	ParameterNameTime                   string = "time"                      //type=int
	ParameterNameChatID                 string = "chat_id"                   //type=int
	ParameterNamePeerID                 string = "peer_id"                   //type=int
	ParameterNameSound                  string = "sound"                     //type=0,1
	ParameterNameURL                    string = "url"                       //type=string
	ParameterNameLinkType               string = "link_type"                 //type=string
	ParameterNameLinkURL                string = "link_url"                  //type=string
	ParameterNameLinkDomain             string = "link_domain"               //type=string
	ParameterNameCampaignID             string = "campaign_id"               //type=string
	ParameterNameSourceType             string = "source_type"               //type=string
	ParameterNameLifeTime               string = "lifetime"                  //type=string
	ParameterNameTargetGroupID          string = "target_group_id"           //type=int
	ParameterNameTargetPixelID          string = "target_pixel_id"           //type=int
	ParameterNameTargetPixelRules       string = "target_pixel_rules"        //type=string (json:map[string]interface{})
	ParameterNameDomain                 string = "domain"                    //type=string
	ParameterNameCategoryID             string = "category_id"               //type=int
	ParameterNameIDs                    string = "ids"                       //type=[]int, string
	ParameterNameIDsType                string = "ids_type"                  //type=string
	ParameterNameIncludeDeleted         string = "include_deleted"           //type=int
	ParameterNameOnlyDeleted            string = "only_deleted"              //type=int
	ParameterNameCampaignIDs            string = "campaign_ids"              //type=int
	ParameterNameAdID                   string = "ad_id"                     //type=int
	ParameterNameAdIDs                  string = "ad_ids"                    //type=int
	ParameterNameAdFormat               string = "ad_format"                 //type=int
	ParameterNameLimit                  string = "limit"                     //type=int
	ParameterNameLevel                  string = "level"                     //type=int
	ParameterNamePeriod                 string = "period"                    //type=string
	ParameterNameDateFrom               string = "date_from"                 //type=string
	ParameterNameDateTo                 string = "date_to"                   //type=string
	ParameterNameRequestID              string = "requests_id"               //type=int
	ParameterNameRequestIDs             string = "requests_ids"              //type=string
	ParameterNameSortBy                 string = "sort_by"                   //type=string
	ParameterNameArtistName             string = "artist_name"               //type=string
	ParameterNameStatsFields            string = "stats_fields"              //type=string
	ParameterNameCountry                string = "country"                   //type=int
	ParameterNameCities                 string = "cities"                    //type=string
	ParameterNameQuery                  string = "q"                         //type=string
	ParameterNameCriteria               string = "criteria"                  //type=string
	ParameterNameAdPlatformNoWall       string = "ad_platform_no_wall"       //type=string
	ParameterNameAdPlatformNoAdNetwork  string = "ad_platform_no_ad_network" //type=string
	ParameterNameNeedPrecise            string = "need_precise"              //type=string
	ParameterNameImpressionsLimitPeriod string = "impressions_limit_period"  //type=int
	ParameterNameIcon                   string = "icon"                      //type=1
	ParameterNameContacts               string = "contacts"                  //type=string
	ParameterNameShareWithClientID      string = "share_with_client_id"      //type=string
	ParameterNameImageType              string = "image_type"                //type=string
	ParameterNameImages                 string = "images"                    //type=string
	ParameterNameImage                  string = "image"                     //type=string
	ParameterNameHash                   string = "hash"                      //type=string
	ParameterNameExtended               string = "extended"                  //type=1
	ParameterNameTaskID                 string = "task_id"                   //type=string
	ParameterNameAudio                  string = "audio"                     //type=string
	ParameterNameModel                  string = "model"                     //type=string
	ParameterNameAlbumID                string = "album_id"                  //type=int
	ParameterNameGroupID                string = "group_id"                  //type=int
	ParameterNameFieldSquareCrop        string = "_square_crop"              // type=string
	ParameterNameAct                    string = "act"                       //type=int
	ParameterNameTs                     string = "ts"                        //type=int
	ParameterNameWait                   string = "wait"                      //type=int
	ParameterNameMode                   string = "mode"                      //type=int
	ParameterNameNeedPts                string = "need_pts"                  //type=1
	ParameterNameEnabled                string = "enabled"                   //type=1
	ParameterNameAPIVersion             string = "api_version"               //type=string
)
