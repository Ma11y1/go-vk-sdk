package response

import (
	"go-vk-sdk/constants"
	"go-vk-sdk/errors"
	"go-vk-sdk/objects"
)

// Doc: https://dev.vk.com/ru/method/ads

type AdsAddOfficeUsersResponse struct {
	BaseResponse
	Response []objects.AdsAddOfficeUsersItem `json:"response"`
}

type AdsCheckLinkResponse struct {
	BaseResponse
	Response struct {
		Status      constants.AdsLinkStatus `json:"status"`
		Description string                  `json:"description,omitempty"`  // (if status = disallowed) — description of the reason
		RedirectURL string                  `json:"redirect_url,omitempty"` // (if the end link differs from original and status = allowed) — end link.
	} `json:"response"`
}

type AdsCreateAdsResponse []struct {
	BaseResponse
	Response struct {
		ID int `json:"id"`
		errors.AdsAPIError
	} `json:"response"`
}

type AdsCreateCampaignsResponse []struct {
	BaseResponse
	Response struct {
		ID int `json:"id"`
		errors.AdsAPIError
	} `json:"response"`
}

type AdsCreateClientsResponse []struct {
	BaseResponse
	Response struct {
		ID int `json:"id"`
		errors.AdsAPIError
	} `json:"response"`
}

type AdsCreateLookalikeRequestResponse struct {
	BaseResponse
	Response struct {
		RequestID int `json:"request_id"`
	} `json:"response"`
}

type AdsCreateTargetGroupResponse struct {
	BaseResponse
	Response struct {
		ID int `json:"id"`
	} `json:"response"`
}

type AdsCreateTargetPixelResponse struct {
	BaseResponse
	Response struct {
		ID    int    `json:"id"`
		Pixel string `json:"pixel"`
	} `json:"response"`
}

type AdsDeleteAdsResponse struct {
	BaseResponse
	Response []errors.NumberAPIError `json:"response"`
}

type AdsDeleteCampaignsResponse struct {
	BaseResponse
	Response []errors.NumberAPIError `json:"response"`
}

type AdsDeleteClientsResponse struct {
	BaseResponse
	Response []errors.NumberAPIError `json:"response"`
}

type AdsDeleteTargetGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AdsDeleteTargetPixelResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AdsGetAccountsResponse struct {
	BaseResponse
	Response []objects.AdsAccount `json:"response"`
}

type AdsGetAdsResponse struct {
	BaseResponse
	Response []objects.AdsAd `json:"response"`
}

type AdsGetAdsLayoutResponse struct {
	BaseResponse
	Response []objects.AdsAdLayout `json:"response"`
}

type AdsGetAdsTargetingResponse struct {
	BaseResponse
	Response []objects.AdsTargeting `json:"response"`
}

type AdsGetBudgetResponse struct {
	BaseResponse
	Response string `json:"response"`
}

type AdsGetCampaignsResponse struct {
	BaseResponse
	Response []objects.AdsCampaign `json:"response"`
}

type AdsGetCategoriesResponse struct {
	BaseResponse
	Response struct {
		V1 []objects.AdsCategory `json:"v1"`
		V2 []objects.AdsCategory `json:"v2"`
	} `json:"response"`
}

type AdsGetClientsResponse struct {
	BaseResponse
	Response []objects.AdsClient `json:"response"`
}

type AdsGetDemographicsResponse struct {
	BaseResponse
	Response []objects.AdsDemoStats `json:"response"`
}

type AdsGetFloodStatsResponse struct {
	BaseResponse
	Response objects.AdsFloodStats `json:"response"`
}

type AdsGetLookalikeRequestsResponse struct {
	BaseResponse
	Response struct {
		Count int                           `json:"count"`
		Items []objects.AdsLookalikeRequest `json:"items"`
	} `json:"response"`
}

type AdsGetMusiciansResponse struct {
	BaseResponse
	Response []objects.AdsMusician `json:"response"`
}

type AdsGetMusiciansByIdsResponse struct {
	BaseResponse
	Response []objects.AdsMusician `json:"response"`
}

type AdsGetOfficeUsersResponse struct {
	BaseResponse
	Response []objects.AdsOfficeUser `json:"response"`
}

type AdsGetPostsReachResponse struct {
	BaseResponse
	Response []objects.AdsPromotedPostReach `json:"response"`
}

type AdsGetRejectionReasonResponse struct {
	BaseResponse
	Response objects.AdsRejectReason `json:"response"`
}

type AdsGetStatisticsResponse struct {
	BaseResponse
	Response objects.AdsStatistics `json:"response"`
}

type AdsGetSuggestionsResponse struct {
	BaseResponse
	Response objects.AdsTargetSuggestions `json:"response"`
}

type AdsGetSuggestionsRegionsResponse struct {
	BaseResponse
	Response objects.AdsTargetSuggestionsRegions `json:"response"`
}

type AdsGetSuggestionsCitiesResponse struct {
	BaseResponse
	Response objects.AdsTargetSuggestionsCities `json:"response"`
}

type AdsGetSuggestionsSchoolsResponse struct {
	BaseResponse
	Response objects.AdsTargetSuggestionsSchools `json:"response"`
}

type AdsGetTargetGroupsResponse struct {
	BaseResponse
	Response objects.AdsTargetGroup `json:"response"`
}

type AdsGetTargetPixelsResponse struct {
	BaseResponse
	Response []objects.AdsTargetPixelInfo `json:"response"`
}

type AdsGetTargetingStatisticResponse struct {
	BaseResponse
	Response objects.AdsTargetStatistic `json:"response"`
}

type AdsGetUploadURLResponse struct {
	BaseResponse
	Response string `json:"response"`
}

type AdsGetVideoUploadURLResponse struct {
	BaseResponse
	Response string `json:"response"`
}

type AdsImportTargetContactsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AdsRemoveOfficeUsersResponse struct {
	BaseResponse
	Response []bool `json:"response"`
}

type AdsRemoveTargetContactsResponse struct {
	BaseResponse
	Response objects.NumberFlagBool `json:"response"`
}

type AdsSaveLookalikeRequestResultResponse struct {
	BaseResponse
	Response struct {
		AudienceCount      int `json:"audience_count"`
		RetargetingGroupID int `json:"retargeting_group_id"`
	} `json:"response"`
}

type AdsShareTargetGroupResponse struct {
	BaseResponse
	Response struct {
		ID int `json:"id"`
	} `json:"response"`
}

type AdsUpdateAdsResponse struct {
	BaseResponse
	Response []struct {
		ID int `json:"id"`
		errors.AdsAPIError
	} `json:"response"`
}

type AdsUpdateCampaignsResponse struct {
	BaseResponse
	Response []struct {
		ID int `json:"id"`
		errors.AdsAPIError
	} `json:"response"`
}

type AdsUpdateClientsResponse struct {
	BaseResponse
	Response []struct {
		ID int `json:"id"`
		errors.AdsAPIError
	}
}

type AdsUpdateOfficeUsersResponse struct {
	BaseResponse
	Response []struct {
		UserID    int             `json:"user_id"`
		IsSuccess bool            `json:"is_success"`
		Error     errors.APIError `json:"error"`
	} `json:"response"`
}

type AdsUpdateTargetGroupResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type AdsUpdateTargetPixelResponse struct {
	BaseResponse
	Response int `json:"response"`
}
