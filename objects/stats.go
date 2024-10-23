package objects

type StatsActivity struct {
	Comments     int `json:"comments"`
	Copies       int `json:"copies"`
	Hidden       int `json:"hidden"`
	Likes        int `json:"likes"`
	Subscribed   int `json:"subscribed"`
	Unsubscribed int `json:"unsubscribed"`
}

type StatsCity struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type StatsCountry struct {
	Code  string `json:"code"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type StatsPeriod struct {
	Activity   StatsActivity `json:"activity"`
	PeriodFrom int           `json:"period_from"`
	PeriodTo   int           `json:"period_to"`
	Reach      StatsReach    `json:"reach"`
	Visitors   StatsViews    `json:"visitors"`
}

type StatsReach struct {
	Age              []StatsSexAge  `json:"age"`
	Cities           []StatsCity    `json:"cities"`
	Countries        []StatsCountry `json:"countries"`
	MobileReach      int            `json:"mobile_reach"`
	Reach            int            `json:"reach"`
	ReachSubscribers int            `json:"reach_subscribers"`
	Sex              []StatsSexAge  `json:"sex"`
	SexAge           []StatsSexAge  `json:"sex_age"`
}

type StatsSexAge struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type StatsViews struct {
	Age         []StatsSexAge  `json:"age"`
	Cities      []StatsCity    `json:"cities"`
	Countries   []StatsCountry `json:"countries"`
	MobileViews int            `json:"mobile_views"`
	Sex         []StatsSexAge  `json:"sex"`
	SexAge      []StatsSexAge  `json:"sex_age"`
	Views       int            `json:"views"`
	Visitors    int            `json:"visitors"`
}

type StatsWallpostStat struct {
	PostID           int `json:"post_id"`
	Hide             int `json:"hide"`
	JoinGroup        int `json:"join_group"`
	Links            int `json:"links"`
	ReachSubscribers int `json:"reach_subscribers"`
	ReachTotal       int `json:"reach_total"`
	ReachViral       int `json:"reach_viral"`
	ReachAds         int `json:"reach_ads"`
	Report           int `json:"report"`
	ToGroup          int `json:"to_group"`
	Unsubscribe      int `json:"unsubscribe"`
	AdViews          int `json:"ad_views"`
	AdSubscribers    int `json:"ad_subscribers"`
	AdHide           int `json:"ad_hide"`
	AdUnsubscribe    int `json:"ad_unsubscribe"`
	AdLinks          int `json:"ad_links"`
	AdToGroup        int `json:"ad_to_group"`
	AdJoinGroup      int `json:"ad_join_group"`
	AdCoverage       int `json:"ad_coverage"`
	AdReport         int `json:"ad_report"`
}
