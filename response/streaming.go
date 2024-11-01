package response

// Doc: https://dev.vk.com/ru/method/streaming

type StreamingGetServerURLResponse struct {
	BaseResponse
	Response struct {
		Endpoint string `json:"endpoint"`
		Key      string `json:"key"`
	} `json:"response"`
}

type StreamingGetSettingsResponse struct {
	BaseResponse
	Response struct {
		MonthlyLimit string `json:"monthly_limit"`
	} `json:"response"`
}

type StreamingGetStatsResponse struct {
	BaseResponse
	Response []struct {
		EventType string `json:"event_type"`
		Stats     []struct {
			Timestamp int `json:"timestamp"`
			Value     int `json:"value"`
		} `json:"stats"`
	} `json:"response"`
}

type StreamingGetStemResponse struct {
	BaseResponse
	Response struct {
		Stem string `json:"stem"`
	} `json:"response"`
}
