package objects

import (
	"encoding/json"
	"go-vk-sdk/constants"
	"go-vk-sdk/errors"
)

// Doc:

type AdsAccesses struct {
	ClientID string `json:"client_id"`
	Role     string `json:"role"`
}

type AdsAccount struct {
	AccountID                   int     `json:"account_id"`
	AccountName                 string  `json:"account_name"`
	AccountType                 string  `json:"account_type"`
	AccountStatus               BoolInt `json:"account_status"`
	AccessRole                  string  `json:"access_role"`
	CanViewBudget               BoolInt `json:"can_view_budget"`
	AdNetworkAllowedPotentially BoolInt `json:"ad_network_allowed_potentially"`
}

type AdsAdLayout struct {
	ID             string      `json:"id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Video          BoolInt     `json:"video"`
	LinkType       string      `json:"link_type"`
	LinkURL        string      `json:"link_url"`
	LinkDomain     string      `json:"link_domain"`
	AdFormat       interface{} `json:"ad_format"`
	ImageSrc       string      `json:"image_src"`
	ImageSrc2x     string      `json:"image_src_2x"`
	PreviewLink    string      `json:"preview_link"`
	CampaignID     int         `json:"campaign_id"`
	GoalType       int         `json:"goal_type"`
	CostType       int         `json:"cost_type"`
	AgeRestriction string      `json:"age_restriction"`
}

type AdsCampaign struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Type      constants.AdsCampaignType `json:"type"`
	Status    int                       `json:"status"`
	AllLimit  string                    `json:"all_limit"`
	DayLimit  string                    `json:"day_limit"`
	StartTime int                       `json:"start_time"`
	StopTime  int                       `json:"stop_time"`
}

type AdsCategory struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Subcategories []IDName `json:"subcategories"`
}

type AdsClient struct {
	ID       int                           `json:"id"`        // Client ID
	Name     string                        `json:"name"`      // Client name. Must be between 3 and 60 characters.
	DayLimit int                           `json:"day_limit"` // Daily budget limit in rubles.
	AllLimit int                           `json:"all_limit"` // Total budget limit in rubles.
	OrdData  AdsClientOrdDataSpecification `json:"ord_data"`  // Advertising data operator details.
}

type AdsClientOrdDataSpecification struct {
	ClientType     string               `json:"client_type"`        // Client type (person, individual, legal)
	ClientName     string               `json:"client_name"`        // Full name or organization name of the client.
	ContractNumber string               `json:"contract_number"`    // Contract number.
	ContractDate   string               `json:"contract_date"`      // Contract date in DD.MM.YYYY format.
	ContractType   string               `json:"contract_type"`      // EventType of contract (service agreement, addendum, intermediary contract, etc.)
	ContractObject string               `json:"contract_object"`    // Description of the contract subject.
	WithVAT        bool                 `json:"with_vat"`           // Specifies if VAT is applicable in the contract.
	INN            string               `json:"inn"`                // Client's tax identification number (10 or 12 characters).
	Phone          string               `json:"phone"`              // Client's phone number (7 to 20 characters).
	Subagent       AdsClientOrdSubagent `json:"subagent,omitempty"` // Optional. Subcontractor details.
}

type AdsClientOrdSubagent struct {
	Type  string `json:"type"`  // Subcontractor type (person, individual, legal)
	Name  string `json:"name"`  // Subcontractor name.
	INN   string `json:"inn"`   // Subcontractor's tax identification number (10 or 12 characters).
	Phone string `json:"phone"` // Subcontractor's phone number (7 to 20 characters).
}

type AdsTargeting struct {
	AgeFrom              int     `json:"age_from"`
	AgeTo                int     `json:"age_to"`
	Apps                 string  `json:"apps"`
	AppsNot              string  `json:"apps_not"`
	Birthday             int     `json:"birthday"`
	Cities               string  `json:"cities"`
	CitiesNot            string  `json:"cities_not"`
	Country              int     `json:"country"`
	Districts            string  `json:"districts"`
	Groups               string  `json:"groups"`
	InterestCategories   string  `json:"interest_categories"`
	Interests            string  `json:"interests"`
	Paying               BoolInt `json:"paying"`
	Positions            string  `json:"positions"`
	Religions            string  `json:"religions"`
	RetargetingGroups    string  `json:"retargeting_groups"`
	RetargetingGroupsNot string  `json:"retargeting_groups_not"`
	SchoolFrom           int     `json:"school_from"`
	SchoolTo             int     `json:"school_to"`
	Schools              string  `json:"schools"`
	Sex                  int     `json:"sex"`
	Stations             string  `json:"stations"`
	Statuses             string  `json:"statuses"`
	Streets              string  `json:"streets"`
	Travellers           int     `json:"travellers"`
	UniFrom              int     `json:"uni_from"`
	UniTo                int     `json:"uni_to"`
	UserBrowsers         string  `json:"user_browsers"`
	UserDevices          string  `json:"user_devices"`
	UserOs               string  `json:"user_os"`
}

type AdsDemoStats struct {
	ID    int                     `json:"id"`
	Stats AdsDemoStatisticsFormat `json:"stats"`
	Type  string                  `json:"type"`
}

type AdsDemoStatisticsFormat struct {
	Age     []AdsStatsAge         `json:"age"`
	Cities  []AdsStatisticsCities `json:"cities"`
	Day     string                `json:"day"`
	Month   string                `json:"month"`
	Overall int                   `json:"overall"`
	Sex     []AdsStatisticsSex    `json:"sex"`
	SexAge  []AdsStatisticsSexAge `json:"sex_age"`
}

type AdsFloodStats struct {
	Left    int `json:"left"`
	Refresh int `json:"refresh"`
}

type AdsRejectReason struct {
	Comment string     `json:"comment"`
	Rules   []AdsRules `json:"rules"`
}

type AdsRules struct {
	Title      string          `json:"title"`
	Paragraphs []AdsParagraphs `json:"paragraphs"`
}

type AdsParagraphs struct {
	Paragraph string `json:"paragraph"`
}

type AdsStatistics struct {
	ID    int                   `json:"id"`
	Type  string                `json:"type"`
	Stats []AdsStatisticsFormat `json:"stats"`
}

type AdsStatisticsFormat struct {
	Day                         string          `json:"day"`
	DayFrom                     string          `json:"day_from"`
	DayTo                       string          `json:"day_to"`
	Month                       string          `json:"month"`
	Overall                     BoolInt         `json:"overall"`
	Spent                       int             `json:"spent"`
	Impressions                 int             `json:"impressions"`
	Clicks                      int             `json:"clicks"`
	Reach                       int             `json:"reach"`
	JoinRate                    int             `json:"join_rate"`
	UniqViewsCount              int             `json:"uniq_views_count"`     // number of unique views
	LinkExternalClicks          int             `json:"link_external_clicks"` // number of unique clicks on links in the post text, snippets, cards and buttons that lead users to resources outside of VKontakte
	ConversionCount             int             `json:"conversion_count"`
	ConversionCr                string          `json:"conversion_cr"`
	ConversionsExternal         int             `json:"conversions_external"`
	CTR                         string          `json:"ctr"`
	VideoClicksSite             int             `json:"video_clicks_site"`
	VideoViews                  int             `json:"video_views"`
	VideoViewsFull              int             `json:"video_views_full"`
	VideoViewsHalf              int             `json:"video_views_half"`
	VideoPlaysUnique100Percents int             `json:"video_plays_unique_100_percents"`
	VideoPlaysUnique10Seconds   int             `json:"video_plays_unique_10_seconds"`
	VideoPlaysUnique25Percents  int             `json:"video_plays_unique_25_percents"`
	VideoPlaysUnique3Seconds    int             `json:"video_plays_unique_3_seconds"`
	VideoPlaysUnique50Percents  int             `json:"video_plays_unique_50_percents"`
	VideoPlaysUnique75Percents  int             `json:"video_plays_unique_75_percents"`
	VideoPlaysUniqueStarted     int             `json:"video_plays_unique_started"`
	EffectiveCostPerClick       string          `json:"effective_cost_per_click"`
	EffectiveCostPerMessage     string          `json:"effective_cost_per_message"`
	EffectiveCostPerMille       string          `json:"effective_cost_per_mille"`
	EffectiveCPF                string          `json:"effective_cpf"`
	MessageSends                int             `json:"message_sends"`
	MessageSendsByAnyUser       int             `json:"message_sends_by_any_user"`
	MobileAppStat               []MobileAppStat `json:"mobile_app_stat"`
}

type MobileAppStat KeyValue

type AdsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type AdsStatisticsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Name            string  `json:"name"`
	Value           int     `json:"value"`
}

type AdsStatisticsSex struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type AdsStatisticsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`
	ImpressionsRate float64 `json:"impressions_rate"`
	Value           string  `json:"value"`
}

type AdsTargetSettings struct{}

type AdsTargetSuggestions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AdsTargetSuggestionsRegions struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type AdsTargetSuggestionsCities struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type AdsTargetSuggestionsSchools struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Type   string `json:"type"`
	Parent string `json:"parent"`
}

type AdsTargetGroup struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	AudienceCount   int     `json:"audience_count"`
	Lifetime        int     `json:"lifetime"`
	LastUpdated     int     `json:"last_updated"`
	IsAudience      BoolInt `json:"is_audience"`
	IsShared        BoolInt `json:"is_shared"`
	FileSource      BoolInt `json:"file_source"`
	APISource       BoolInt `json:"api_source"`
	LookalikeSource BoolInt `json:"lookalike_source"`
	Domain          string  `json:"domain,omitempty"`
	Pixel           string  `json:"pixel,omitempty"`
}

type AdsTargetPixelInfo struct {
	Name          string `json:"name"`
	Pixel         string `json:"pixel"`
	CategoryID    int    `json:"category_id"`
	TargetPixelID int    `json:"target_pixel_id"`
	LastUpdated   int    `json:"last_updated"`
	Domain        string `json:"domain"`
}

type AdsUsers struct {
	Accesses []AdsAccesses `json:"accesses"`
	UserID   int           `json:"user_id"`
}

type AdsTargetStatistic struct {
	AudienceCount         int             `json:"audience_count"`          // Size of the target audience.
	RecommendedCPC        float64         `json:"recommended_cpc"`         // Recommended cost-per-click price (in rubles, with decimals), 0 if not applicable.
	RecommendedCPM        float64         `json:"recommended_cpm"`         // Recommended cost-per-thousand impressions price (in rubles, with decimals).
	RecommendedCPC50      float64         `json:"recommended_cpc_50"`      // Recommended CPC to reach 50% of the audience within the first week.
	RecommendedCPM50      float64         `json:"recommended_cpm_50"`      // Recommended CPM to reach 50% of the audience within the first week.
	RecommendedCPC70      float64         `json:"recommended_cpc_70"`      // Recommended CPC to reach 70% of the audience within the first week.
	RecommendedCPM70      float64         `json:"recommended_cpm_70"`      // Recommended CPM to reach 70% of the audience within the first week.
	RecommendedCPC90      float64         `json:"recommended_cpc_90"`      // Recommended CPC to reach 90% of the audience within the first week.
	RecommendedCPM90      float64         `json:"recommended_cpm_90"`      // Recommended CPM to reach 90% of the audience within the first week.
	RecommendedCPMPrecise map[int]float64 `json:"recommended_cpm_precise"` // Recommended CPM for precise reach percentages (if need_precise is set), with the key as the percentage (e.g., 5, 10, 15...95).
	RecommendedCPCPrecise map[int]float64 `json:"recommended_cpc_precise"` // Recommended CPC for precise reach percentages (if need_precise is set), with the key as the percentage (e.g., 5, 10, 15...95).
}

type AdsTargetStatsCriteria struct {
	Sex                  int    `json:"sex"`                    // Gender: 0 - any, 1 - female, 2 - male
	AgeFrom              int    `json:"age_from"`               // Minimum age (0 if not set)
	AgeTo                int    `json:"age_to"`                 // Maximum age (0 if not set)
	Birthday             int    `json:"birthday"`               // Birthday flag (1 - today, 2 - tomorrow, 4 - within the week)
	Country              int    `json:"country"`                // Country ID (0 if not set)
	Cities               string `json:"cities"`                 // Comma-separated city/region IDs, with regions prefixed by a minus sign
	CitiesNot            string `json:"cities_not"`             // Comma-separated city/region IDs to exclude, with regions prefixed by a minus sign
	GeoNear              string `json:"geo_near"`               // Geo-targeting points, format: "latitude,longitude,radius[,place name]"
	GeoPointType         string `json:"geo_point_type"`         // EventType of geo-location: "regular", "home", "work", "online"
	Statuses             string `json:"statuses"`               // Comma-separated relationship status IDs (1-8)
	Groups               string `json:"groups"`                 // Comma-separated community IDs
	GroupsNot            string `json:"groups_not"`             // Comma-separated community IDs to exclude
	Apps                 string `json:"apps"`                   // Comma-separated application IDs
	AppsNot              string `json:"apps_not"`               // Comma-separated application IDs to exclude
	Districts            string `json:"districts"`              // Comma-separated district IDs
	Stations             string `json:"stations"`               // Comma-separated metro station IDs
	Streets              string `json:"streets"`                // Comma-separated street IDs
	Schools              string `json:"schools"`                // Comma-separated school IDs
	Positions            string `json:"positions"`              // Comma-separated position/job title IDs
	Religions            string `json:"religions"`              // Comma-separated religion IDs
	InterestCategories   string `json:"interest_categories"`    // Comma-separated interest category IDs
	Interests            string `json:"interests"`              // Comma-separated interest IDs
	UserDevices          string `json:"user_devices"`           // Comma-separated user device IDs
	UserOS               string `json:"user_os"`                // Comma-separated operating system IDs
	UserBrowsers         string `json:"user_browsers"`          // Comma-separated browser IDs
	RetargetingGroups    string `json:"retargeting_groups"`     // Comma-separated retargeting group IDs
	RetargetingGroupsNot string `json:"retargeting_groups_not"` // Comma-separated retargeting group IDs to exclude
	Paying               int    `json:"paying"`                 // Payment filter: 1 - didn't use VKontakte currency, 2 - used VKontakte currency
	Travellers           int    `json:"travellers"`             // Traveller filter: 1 - include only travellers
	SchoolFrom           int    `json:"school_from"`            // Minimum school graduation year (0 if not set)
	SchoolTo             int    `json:"school_to"`              // Maximum school graduation year (0 if not set)
	UniFrom              int    `json:"uni_from"`               // Minimum university graduation year (0 if not set)
	UniTo                int    `json:"uni_to"`
}

func (a *AdsTargetStatsCriteria) ToJSON() string {
	jsonData, err := json.Marshal(*a)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsAd struct {
	ID                      string      `json:"id"`
	Name                    string      `json:"name"`
	Status                  int         `json:"status"`
	Approved                string      `json:"approved"`
	AllLimit                string      `json:"all_limit"`
	Category1ID             string      `json:"category1_id"`
	Category2ID             string      `json:"category2_id"`
	Cpm                     string      `json:"cpm"`
	AdFormat                int         `json:"ad_format"`
	AdPlatform              interface{} `json:"ad_platform"`
	CampaignID              int         `json:"campaign_id"`
	CostType                int         `json:"cost_type"`
	Cpc                     int         `json:"cpc"`
	DisclaimerMedical       BoolInt     `json:"disclaimer_medical"`
	DisclaimerSpecialist    BoolInt     `json:"disclaimer_specialist"`
	DisclaimerSupplements   BoolInt     `json:"disclaimer_supplements"`
	Video                   BoolInt     `json:"video"`
	ImpressionsLimited      BoolInt     `json:"impressions_limited"`
	Autobidding             BoolInt     `json:"autobidding"`
	ImpressionsLimit        int         `json:"impressions_limit"`
	CreateTime              string      `json:"create_time"`
	UpdateTime              string      `json:"update_time"`
	GoalType                int         `json:"goal_type"`
	DayLimit                string      `json:"day_limit"`
	StartTime               string      `json:"start_time"`
	StopTime                string      `json:"stop_time"`
	AgeRestriction          string      `json:"age_restriction"`
	EventsRetargetingGroups interface{} `json:"events_retargeting_groups"`
	ImpressionsLimitPeriod  string      `json:"impressions_limit_period"`
}

type AdsLookalikeRequest struct {
	ID                       int                                 `json:"id"`                          // Unique ID of the lookalike audience search request (unique per account/client).
	CreateTime               int64                               `json:"create_time"`                 // Timestamp of when the request was created.
	UpdateTime               int64                               `json:"update_time"`                 // Timestamp of the last status update.
	Status                   constants.AdsLookalikeRequestStatus `json:"status"`                      // Current status of the request (search_in_progress, search_done, search_failed, save_in_progress, save_done, save_failed).
	ScheduledDeleteTime      int64                               `json:"scheduled_delete_time"`       // Timestamp of the scheduled deletion of the request.
	SourceType               string                              `json:"source_type"`                 // Source type of the original audience, always "retargeting_group".
	SourceRetargetingGroupID int                                 `json:"source_retargeting_group_id"` // ID of the retargeting group for the original audience.
	SourceName               string                              `json:"source_name"`                 // Name of the retargeting audience at the time of request creation.
	AudienceCount            int                                 `json:"audience_count"`              // Size of the original audience.
	SaveAudienceLevels       []AdsAudienceLevel                  `json:"save_audience_levels"`        // Available audience sizes for saving after a successful search.
}

type AdsAudienceLevel struct {
	Level         int `json:"level"`          // Level parameter used for saving the lookalike audience result.
	AudienceCount int `json:"audience_count"` // Size of the lookalike audience for this option.
}

type AdsPromotedPostReach struct {
	ID               int `json:"id"`
	Hide             int `json:"hide"`
	JoinGroup        int `json:"join_group"`
	Links            int `json:"links"`
	ReachSubscribers int `json:"reach_subscribers"`
	ReachTotal       int `json:"reach_total"`
	Report           int `json:"report"`
	ToGroup          int `json:"to_group"`
	Unsubscribe      int `json:"unsubscribe"`
	VideoViews100p   int `json:"video_views_100p"`
	VideoViews25p    int `json:"video_views_25p"`
	VideoViews3s     int `json:"video_views_3s"`
	VideoViews50p    int `json:"video_views_50p"`
	VideoViews75p    int `json:"video_views_75p"`
	VideoViewsStart  int `json:"video_views_start"`
}

type AdsMusician struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}

type AdsOfficeUser struct {
	ClientID int    `json:"client_id"` // for admin client_id is not specified
	Role     string `json:"role"`      // admin, manager или reports
}

type AdsUpdateAdsStatus struct {
	ID int `json:"id"`
	errors.AdsAPIError
}

type AdsAddOfficeUsersItem struct {
	OK    BoolInt
	Error errors.AdsAPIError
}

func (a *AdsAddOfficeUsersItem) UnmarshalJSON(data []byte) (err error) {
	if a.OK.UnmarshalJSON(data) != nil {
		err = json.Unmarshal(data, &a.Error)
		if err != nil {
			return err
		}
	}

	return
}

// AdsUserSpecification Doc: https://dev.vk.com/ru/method/ads.createAds
type AdsUserSpecification struct {
	UserID                  int                                `json:"user_id"`                     // required The ID of the user being added as an administrator or observer.
	Role                    constants.AdsUserSpecificationRole `json:"role"`                        // required  parameter, string Flag describing the type of authority: reports - observer; manager - administrator.
	ClientID                int                                `json:"client_id"`                   // required An array of client identifiers. Maximum allowed number of clients added to a user is 100.
	ViewBudget              BoolInt                            `json:"view_budget"`                 // optional Whether to show the budget to the user.
	GrantAccessToAllClients bool                               `json:"grant_access_to_all_clients"` // optional Access to all current and new clients of this account.
}

func (u *AdsUserSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsUserSpecifications []AdsUserSpecification

func (u *AdsUserSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// AdsAdSpecification Doc: https://dev.vk.com/ru/method/ads.createAds
type AdsAdSpecification struct {
	CampaignID                int      `json:"campaign_id"`                            // Required. Integer. The ID of the campaign where the ad will be created.
	AdFormat                  int      `json:"ad_format"`                              // Required. Integer. The format of the ad (1: Image and text, 2: Large image, etc.).
	Autobidding               BoolInt  `json:"autobidding,omitempty"`                  // Optional. Flag (0 or 1). Automatic bid management (0: disabled, 1: enabled).
	AutobiddingMaxCost        string   `json:"autobidding_max_cost,omitempty"`         // Optional. Max automatic bid cost. Can be integer or string with decimal places.
	CostType                  int      `json:"cost_type"`                              // Required. Payment method (0: per click, 1: per impression, 3: optimized per action).
	CPC                       float64  `json:"cpc,omitempty"`                          // Required if cost_type is 0. Price per click in rubles.
	CPM                       float64  `json:"cpm,omitempty"`                          // Required if cost_type is 1. Price per 1000 impressions in rubles.
	OCPM                      float64  `json:"ocpm,omitempty"`                         // Required if cost_type is 3. Price per action for oCPM.
	GoalType                  int      `json:"goal_type,omitempty"`                    // Required if cost_type is 3. EventType of goal.
	ImpressionsLimit          int      `json:"impressions_limit,omitempty"`            // Optional. Impression limit per user if ad_format = 9 or 11 and cost_type = 1.
	ImpressionsLimited        BoolInt  `json:"impressions_limited,omitempty"`          // Optional. Flag (0 or 1). Impression limit per user (0: no limit, 1: max 100 per user).
	ImpressionsLimitPeriod    int      `json:"impressions_limit_period,omitempty"`     // Optional. Impression limit period in seconds, must be a multiple of 86400.
	AdPlatform                string   `json:"ad_platform,omitempty"`                  // Optional. Platforms to display the ad (0: VK and partners, 1: VK only).
	AdPlatformNoWall          BoolInt  `json:"ad_platform_no_wall,omitempty"`          // Optional. Only for ad_format = 9 or 11. 1: Do not display on community walls.
	AdPlatformNoAdNetwork     BoolInt  `json:"ad_platform_no_ad_network,omitempty"`    // Optional. 1: Do not display in ad network.
	PublisherPlatforms        string   `json:"publisher_platforms,omitempty"`          // Optional. Platforms to display the ad (all, social, vk).
	DayLimit                  float64  `json:"day_limit,omitempty"`                    // Optional. Daily limit in rubles.
	AllLimit                  float64  `json:"all_limit,omitempty"`                    // Optional. Total limit in rubles.
	Category1ID               int      `json:"category1_id,omitempty"`                 // Optional. ID of the ad category.
	Category2ID               int      `json:"category2_id,omitempty"`                 // Optional. ID of the secondary ad category.
	AgeRestriction            int      `json:"age_restriction,omitempty"`              // Optional. Age restriction flag (0: No mark, 1: 0+, 2: 6+, etc.).
	Status                    BoolInt  `json:"status,omitempty"`                       // Optional. Ad status (0: paused, 1: running).
	Name                      string   `json:"name,omitempty"`                         // Optional. Name of the ad (3-60 characters).
	Title                     string   `json:"title,omitempty"`                        // Optional. Title of the ad (3-33 characters, 25 for ad_format = 11).
	Description               string   `json:"description,omitempty"`                  // Optional. Ad description (3-70 characters, 90 for ad_format = 11).
	LinkURL                   string   `json:"link_url"`                               // Required. url of the advertised objects.
	LinkDomain                string   `json:"link_domain,omitempty"`                  // Optional. Domain of the advertised objects.
	LinkTitle                 string   `json:"link_title,omitempty"`                   // Optional. Title next to the link button (3-25 characters).
	LinkButton                string   `json:"link_button,omitempty"`                  // Optional. Button ID for the ad.
	Photo                     string   `json:"photo,omitempty"`                        // Optional. Main imag for certain ad formats.
	PhotoIcon                 string   `json:"photo_icon,omitempty"`                   // Optional. Logo for certain ad formats.
	Video                     string   `json:"video,omitempty"`                        // Optional. Main video for adaptive format.
	RepeatVideo               BoolInt  `json:"repeat_video,omitempty"`                 // Optional. Flag (0 or 1) to loop the video.
	DisclaimerMedical         BoolInt  `json:"disclaimer_medical,omitempty"`           // Optional. Flag (0 or 1) to display a medical disclaimer.
	DisclaimerSpecialist      BoolInt  `json:"disclaimer_specialist,omitempty"`        // Optional. Flag (0 or 1) to display a specialist consultation disclaimer.
	DisclaimerSupplements     BoolInt  `json:"disclaimer_supplements,omitempty"`       // Optional. Flag (0 or 1) to display a supplement disclaimer.
	WeeklyScheduleHours       []string `json:"weekly_schedule_hours,omitempty"`        // Optional. Array for weekly display schedule, 7 strings of 0/1 (each string represents a day).
	WeeklyScheduleUseHolidays BoolInt  `json:"weekly_schedule_use_holidays,omitempty"` // Optional. Flag (0 or 1) to apply Sunday schedule to holidays.
	StatsURL                  string   `json:"stats_url,omitempty"`                    // Optional. External statistics pixel.
	StatsURL2                 string   `json:"stats_url2,omitempty"`                   // Optional. Additional external statistics pixel.
}

func (u *AdsAdSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsAdSpecifications []AdsAdSpecification

func (u *AdsAdSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsCampaignSpecification struct {
	ClientID  int     `json:"client_id"`            // Required. Integer. Only for advertising agencies. The client ID for whom the campaign will be created.
	Type      string  `json:"type"`                 // Required. Type of campaign (normal, promoted_posts, adaptive_ads).
	Name      string  `json:"name"`                 // Required. The name of the advertising campaign (3-60 characters).
	DayLimit  float64 `json:"day_limit,omitempty"`  // Optional. Daily budget limit in rubles. Must be a positive number.
	AllLimit  float64 `json:"all_limit,omitempty"`  // Optional. Total budget limit in rubles. Must be a positive number.
	StartTime int     `json:"start_time,omitempty"` // Optional. Start time of the campaign in Unix time (positive integer).
	StopTime  int     `json:"stop_time,omitempty"`  // Optional. End time of the campaign in Unix time (positive integer).
	Status    BoolInt `json:"status,omitempty"`     // Optional. Campaign status (0: stopped, 1: running).
}

func (u *AdsCampaignSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsCampaignSpecifications []AdsCampaignSpecification

func (u *AdsCampaignSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsClientSpecification struct {
	Name     string                  `json:"name"`      // Required. Name of the client. Must be between 3 and 60 characters.
	DayLimit int                     `json:"day_limit"` // Required. Daily budget limit in rubles.
	AllLimit int                     `json:"all_limit"` // Required. Total budget limit in rubles.
	OrdData  AdsOrdDataSpecification `json:"ord_data"`  // Required. JSON string containing data of the advertising data operator (ord_data_specification).
}

func (u *AdsClientSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsClientSpecifications []AdsClientSpecification

func (u *AdsClientSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// AdsOrdDataSpecification represents the advertising data operator.
type AdsOrdDataSpecification struct {
	ClientType        string         `json:"client_type"`         // Required. EventType of client (person, individual, legal).
	ClientName        string         `json:"client_name"`         // Required. Full name of the client or the client’s organization.
	INN               string         `json:"inn"`                 // Required. Client's INN (10 or 12 characters).
	Phone             string         `json:"phone"`               // Required. Client's phone number (7 to 20 characters).
	Subagent          AdsOrdSubagent `json:"subagent,omitempty"`  // Optional. JSON objects with subagent information (ord_subagent).
	ContractNumber    string         `json:"contract_number"`     // Required. Contract number.
	ContractDate      string         `json:"contract_date"`       // Required. Contract date in DD.MM.YYYY format.
	ContractType      string         `json:"contract_type"`       // Required. EventType of contract (e.g., "Service Agreement", "Supplementary Agreement").
	ContractObject    string         `json:"contract_object"`     // Required. Subject of the contract (a brief description or copy from the "Subject of Contract" clause).
	WithVAT           bool           `json:"with_vat"`            // Required. Whether VAT is applicable under this contract (true or false).
	IsSubagentEnabled bool           `json:"is_subagent_enabled"` // false if you need to delete data about an existing contractor

}

// AdsOrdSubagent represents the subcontractor (ord_subagent).
type AdsOrdSubagent struct {
	Type  string `json:"type"`  // Required. Type of subcontractor (person, individual, legal).
	Name  string `json:"name"`  // Required. Name of the subcontractor.
	INN   string `json:"inn"`   // Required. Subcontractor's INN (10 or 12 characters).
	Phone string `json:"phone"` // Required. Subcontractor's phone number (7 to 20 characters).
}

type AdsAdEditSpecification struct {
	AdID                      int                    `json:"ad_id"`                                  // Required. ID of the ad being edited.
	CPC                       float64                `json:"cpc,omitempty"`                          // Optional. Price per click, specified in rubles with decimals (if cost_type = 0).
	CPM                       float64                `json:"cpm,omitempty"`                          // Optional. Price per thousand impressions, specified in rubles with decimals (if cost_type = 1).
	OCPM                      float64                `json:"ocpm,omitempty"`                         // Optional. Price per action for oCPM, specified in rubles with decimals (if cost_type = 3).
	ImpressionsLimit          int                    `json:"impressions_limit,omitempty"`            // Optional. Limit on impressions per user if ad_format = 9 or 11 and cost_type = 1 (valid values: 1, 2, 3, 5, 10, 15, 20).
	ImpressionsLimited        BoolInt                `json:"impressions_limited,omitempty"`          // Optional. Flag indicating whether the ad has a limit of impressions per user (0 = no limit, 1 = up to 100 impressions per user).
	ImpressionsLimitPeriod    int                    `json:"impressions_limit_period,omitempty"`     // Optional. Period for resetting impression count per user, in seconds (if ad_format = 9 or 11 and cost_type = 1).
	AdPlatform                string                 `json:"ad_platform,omitempty"`                  // Optional. Platforms for displaying the ad, depending on the ad_format and cost_type.
	AdPlatformNoWall          int                    `json:"ad_platform_no_wall,omitempty"`          // Optional. For ad_format = 9 or 11. Flag indicating whether the ad is displayed on community walls (1 = do not display).
	AdPlatformNoAdNetwork     int                    `json:"ad_platform_no_ad_network,omitempty"`    // Optional. For ad_format = 9 or 11. Flag indicating whether the ad is displayed on ad networks (1 = do not display).
	PublisherPlatforms        string                 `json:"publisher_platforms,omitempty"`          // Optional. Platforms where the ad will be shown (all, social, vk).
	DayLimit                  float64                `json:"day_limit,omitempty"`                    // Optional. Daily budget in rubles.
	AllLimit                  float64                `json:"all_limit,omitempty"`                    // Optional. Total budget in rubles.
	Category1ID               int                    `json:"category1_id,omitempty"`                 // Optional. ID of the primary ad category or subcategory.
	Category2ID               int                    `json:"category2_id,omitempty"`                 // Optional. ID of the secondary ad category or subcategory.
	AgeRestriction            int                    `json:"age_restriction,omitempty"`              // Optional. Age restriction tag for the ad (0 = no tag, 1 = 0+, 2 = 6+, etc.).
	Status                    BoolInt                `json:"status,omitempty"`                       // Optional. Status of the ad (0 = stopped, 1 = active).
	Name                      string                 `json:"name,omitempty"`                         // Optional. Name of the ad for internal use (3 to 60 characters).
	Title                     string                 `json:"title,omitempty"`                        // Optional. Title of the ad (3 to 33 characters, up to 25 for ad_format = 11).
	Description               string                 `json:"description,omitempty"`                  // Optional. Description of the ad (3 to 70 characters, up to 90 for ad_format = 11).
	LinkURL                   string                 `json:"link_url,omitempty"`                     // Required. url of the advertised objects.
	LinkDomain                string                 `json:"link_domain,omitempty"`                  // Optional. Domain of the advertised objects.
	LinkTitle                 string                 `json:"link_title,omitempty"`                   // Optional. Title next to the button/link.
	LinkButton                string                 `json:"link_button,omitempty"`                  // Optional. Button ID for the ad.
	Photo                     string                 `json:"photo,omitempty"`                        // Optional. Main image for certain ad formats.
	PhotoIcon                 string                 `json:"photo_icon,omitempty"`                   // Optional. Logo for certain ad formats.
	Video                     string                 `json:"video,omitempty"`                        // Optional. Main video for adaptive format ads.
	RepeatVideo               BoolInt                `json:"repeat_video,omitempty"`                 // Optional. Flag to loop the video (0 = no, 1 = yes).
	DisclaimerMedical         BoolInt                `json:"disclaimer_medical,omitempty"`           // Optional. Display a medical disclaimer ("Consult a specialist") (1 = enabled).
	DisclaimerSpecialist      BoolInt                `json:"disclaimer_specialist,omitempty"`        // Optional. Display a specialist disclaimer ("Consult a specialist") (1 = enabled).
	DisclaimerSupplements     BoolInt                `json:"disclaimer_supplements,omitempty"`       // Optional. Display a disclaimer for supplements ("Not a medicinal product") (1 = enabled).
	StartTime                 int64                  `json:"start_time,omitempty"`                   // Optional. Start time for the ad in Unix time (0 to reset).
	StopTime                  int64                  `json:"stop_time,omitempty"`                    // Optional. Stop time for the ad in Unix time (0 to reset).
	WeeklyScheduleHours       [7]string              `json:"weekly_schedule_hours,omitempty"`        // Optional. Schedule for showing the ad by hours, one string per day (0 = don't show, 1 = show).
	WeeklyScheduleUseHolidays BoolInt                `json:"weekly_schedule_use_holidays,omitempty"` // Optional. Flag to use Sunday schedule on holidays (0 = no, 1 = yes).
	StatsURL                  string                 `json:"stats_url,omitempty"`                    // Optional. External stats pixel url.
	StatsURL2                 string                 `json:"stats_url2,omitempty"`                   // Optional. Additional external stats pixel url.
	AutoBidding               BoolInt                `json:"autobidding,omitempty"`                  // Optional. Flag for automatic bid management (0 = off, 1 = on).
	AutoBiddingMaxCost        interface{}            `json:"autobidding_max_cost,omitempty"`         // Optional. Maximum automatic bid limit (rubles with decimals), 0 to remove limit.
	TargetingSettings         map[string]interface{} `json:"targeting_settings,omitempty"`           // Optional. Fields for targeting settings.
}

func (u *AdsAdEditSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsAdEditSpecifications []AdsAdEditSpecification

func (u *AdsAdEditSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsCampaignModSpecification struct {
	CampaignID int     `json:"campaign_id"`
	Name       string  `json:"name"`
	DayLimit   int     `json:"day_limit"`
	AllLimit   int     `json:"all_limit"`
	StartTime  int64   `json:"start_time"`
	StopTime   int64   `json:"stop_time"`
	Status     BoolInt `json:"status"`
}

func (u *AdsCampaignModSpecification) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

type AdsCampaignModSpecifications []AdsCampaignModSpecification

func (u *AdsCampaignModSpecifications) ToJSON() string {
	jsonData, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
