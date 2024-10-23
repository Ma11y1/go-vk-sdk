package response

// Doc: https://dev.vk.com/ru/method/calls

type CallsStartResponse struct {
	BaseResponse
	Response struct {
		JoinLink string `json:"join_link"`
		CallID   string `json:"call_id"`
	} `json:"response"`
}

type CallsForceFinish struct {
	BaseResponse
	Response int `json:"response"`
}
