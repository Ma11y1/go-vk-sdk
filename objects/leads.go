package objects

type LeadChecked struct {
	Reason    string `json:"reason"`
	Result    string `json:"result"`
	Sid       string `json:"sid"`
	StartLink string `json:"start_link"`
}

type LeadComplete struct {
	Cost     int     `json:"cost"`
	Limit    int     `json:"limit"`
	Spent    int     `json:"spent"`
	Success  BoolInt `json:"success"`
	TestMode BoolInt `json:"test_mode"`
}

type LeadEntry struct {
	Aid       int     `json:"aid"`
	Comment   string  `json:"comment"`
	Date      int     `json:"date"`
	Sid       string  `json:"sid"`
	StartDate int     `json:"start_date"`
	Status    int     `json:"status"`
	TestMode  BoolInt `json:"test_mode"`
	UID       int     `json:"uid"`
}

type LeadLead struct {
	Completed   int          `json:"completed"`
	Cost        int          `json:"cost"`
	Days        LeadLeadDays `json:"days"`
	Impressions int          `json:"impressions"`
	Limit       int          `json:"limit"`
	Spent       int          `json:"spent"`
	Started     int          `json:"started"`
}

type LeadLeadDays struct {
	Completed   int `json:"completed"`
	Impressions int `json:"impressions"`
	Spent       int `json:"spent"`
	Started     int `json:"started"`
}

type LeadStart struct {
	TestMode BoolInt `json:"test_mode"`
	VkSid    string  `json:"vk_sid"`
}
