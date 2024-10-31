package objects

type AccountOffer struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Tag              string `json:"tag"`
	Price            int    `json:"price"`
	Img              string `json:"img"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Instruction      string `json:"instruction"`
	InstructionHTML  string `json:"instruction_html"`
}

type AccountNameRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

type AccountPushItem struct {
	PeerID               int     `json:"peer_id"`
	Sound                int     `json:"sound"`
	DisabledUntil        int     `json:"disabled_until"`
	DisabledMentions     BoolInt `json:"disabled_mentions"`
	DisabledMassMentions BoolInt `json:"disabled_mass_mentions"`
}

type AccountPushConversations struct {
	Count int                `json:"count"`
	Items []*AccountPushItem `json:"items"`
}

type AccountPushParams struct {
	AppRequest     []string `json:"app_request"`
	Birthday       []string `json:"birthday"`
	Chat           []string `json:"chat"`
	Comment        []string `json:"comment"`
	EventSoon      []string `json:"event_soon"`
	Friend         []string `json:"friend"`
	FriendAccepted []string `json:"friend_accepted"`
	FriendFound    []string `json:"friend_found"`
	GroupAccepted  []string `json:"group_accepted"`
	GroupInvite    []string `json:"group_invite"`
	Like           []string `json:"like"`
	Mention        []string `json:"mention"`
	Msg            []string `json:"msg"`
	NewPost        []string `json:"new_post"`
	PhotosTag      []string `json:"photos_tag"`
	Reply          []string `json:"reply"`
	Repost         []string `json:"repost"`
	SdkOpen        []string `json:"sdk_open"`
	WallPost       []string `json:"wall_post"`
	WallPublish    []string `json:"wall_publish"`
}

type AccountCounters struct {
	Friends                int `json:"friends"`
	FriendsSuggestions     int `json:"friends_suggestions"`
	FriendsRecommendations int `json:"friends_recommendations"`
	Messages               int `json:"messages"`
	Photos                 int `json:"photos"`
	Videos                 int `json:"videos"`
	Gifts                  int `json:"gifts"`
	Events                 int `json:"events"`
	Groups                 int `json:"groups"`
	SDK                    int `json:"sdk"`
	Notifications          int `json:"notifications"`
	AppRequests            int `json:"app_requests"`
	MenuDiscoverBadge      int `json:"menu_discover_badge"`
	MenuClipsBadge         int `json:"menu_clips_badge"`
	Faves                  int `json:"faves"`
	Calls                  int `json:"calls"`
}

type AccountInfo struct {
	Country                   string            `json:"country"`
	HTTPSRequired             BoolInt           `json:"https_required"`
	TwoFactorRequired         BoolInt           `json:"2fa_required"`
	OwnPostsDefault           BoolInt           `json:"own_posts_default"`
	NoWallReplies             BoolInt           `json:"no_wall_replies"`
	Intro                     BoolInt           `json:"intro"`
	Lang                      int               `json:"lang"`
	EuUser                    BoolInt           `json:"eu_user"`
	CommunityComments         BoolInt           `json:"community_comments"`
	IsLiveStreamingEnabled    BoolInt           `json:"is_live_streaming_enabled"`
	IsNewLiveStreamingEnabled BoolInt           `json:"is_new_live_streaming_enabled"`
	LinkRedirects             map[string]string `json:"link_redirects"`
	VkPayEndpointV2           string            `json:"vk_pay_endpoint_v2"`
}

type AccountPushSettings struct {
	Conversations AccountPushConversations `json:"conversations"`
	Settings      AccountPushParams        `json:"settings"`
	Disabled      BoolInt                  `json:"disabled"`
	DisabledUntil int                      `json:"disabled_until"`
}

type AccountProfileInfo struct {
	ID               int                `json:"id"` // always return 0?
	FirstName        string             `json:"first_name"`
	LastName         string             `json:"last_name"`
	MaidenName       string             `json:"maiden_name"`
	ScreenName       string             `json:"screen_name"`
	Sex              int                `json:"sex"`
	Relation         int                `json:"relation"`
	RelationPartner  UsersUserMin       `json:"relation_partner"`
	RelationPending  BoolInt            `json:"relation_pending"`
	RelationRequests []UsersUserMin     `json:"relation_requests"`
	Bdate            string             `json:"bdate"`
	BdateVisibility  int                `json:"bdate_visibility"`
	HomeTown         string             `json:"home_town"`
	Country          Country            `json:"country"`
	City             City               `json:"city"`
	NameRequest      AccountNameRequest `json:"name_request"`
	Status           string             `json:"status"`
	Phone            string             `json:"phone"`
}
