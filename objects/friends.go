package objects

type FriendsStatus struct {
	FriendStatus   int     `json:"friend_status"`
	ReadState      BoolInt `json:"read_state"`
	RequestMessage string  `json:"request_message"`
	Sign           string  `json:"sign"`
	UserID         int     `json:"user_id"`
}

type FriendsList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FriendsRequests struct {
	UserFull
	From      string        `json:"from"`
	Mutual    RequestMutual `json:"mutual"`
	UserID    int           `json:"user_id"`
	TrackCode string        `json:"track_code"`
}

type FriendsRequestsXtrMessage struct {
	FriendsRequests
	Message string `json:"message"`
}

type FriendsUserXtrLists struct {
	UserFull
	Lists []int `json:"lists"`
}

type FriendsUserXtrPhone struct {
	UserFull
	Phone string `json:"phone"`
}
