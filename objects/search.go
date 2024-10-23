package objects

type SearchHint struct {
	Description string    `json:"description"`
	Global      int       `json:"global,omitempty"`
	Group       GroupFull `json:"group,omitempty"`
	Profile     UserFull  `json:"profile,omitempty"`
	Section     string    `json:"section"`
	Type        string    `json:"type"`
}
