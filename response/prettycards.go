package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/prettyCards

type PrettyCardsCreateResponse struct {
	BaseResponse
	Response struct {
		OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
		CardID  string `json:"card_id"`  // Card ID of created pretty card
	} `json:"response"`
}

type PrettyCardsDeleteResponse struct {
	BaseResponse
	Response struct {
		OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
		CardID  string `json:"card_id"`  // Card ID of created pretty card
		Error   string `json:"error"`    // Error reason if error happened
	} `json:"response"`
}

type PrettyCardsEditResponse struct {
	BaseResponse
	Response struct {
		OwnerID int    `json:"owner_id"` // Owner ID of created pretty card
		CardID  string `json:"card_id"`  // Card ID of created pretty card
	} `json:"response"`
}

type PrettyCardsGetResponse struct {
	BaseResponse
	Response struct {
		Count int                             `json:"count"` // Total number
		Items []objects.PrettyCardsPrettyCard `json:"items"`
	} `json:"response"`
}

type PrettyCardsGetByIDResponse struct {
	BaseResponse
	Response []objects.PrettyCardsPrettyCard `json:"response"`
}

type PrettyCardsGetUploadURLResponse struct {
	BaseResponse
	Response string `json:"response"`
}

type PrettyCardsUploadPhotoResponse struct {
	Photo   string `json:"photo"`
	ErrCode int    `json:"errcode"`
}
