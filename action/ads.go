package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/request"
)

// Ads Doc: https://dev.vk.com/ru/method/ads
type Ads struct {
	BaseAction
}

func (a *Ads) GetAPI() *api.API {
	return a.api
}

// AdsAddOfficeUsers Doc: https://dev.vk.com/method/ads.addOfficeUsers
func (a *Account) AdsAddOfficeUsers(actor actor.Actor) *request.AdsAddOfficeUsersRequest {
	return request.NewAdsAddOfficeUsersRequest(a.api, actor)
}

// AdsCheckLink Doc: https://dev.vk.com/method/ads.checkLink
func (a *Account) AdsCheckLink(actor actor.Actor) *request.AdsCheckLinkRequest {
	return request.NewAdsCheckLinkRequest(a.api, actor)
}

// AdsCreateAds Doc: https://dev.vk.com/method/ads.createAds
func (a *Account) AdsCreateAds(actor actor.Actor) *request.AdsCreateAdsRequest {
	return request.NewAdsCreateAdsRequest(a.api, actor)
}

// AdsCreateCampaigns Doc: https://dev.vk.com/method/ads.createCampaigns
func (a *Account) AdsCreateCampaigns(actor actor.Actor) *request.AdsCreateCampaignsRequest {
	return request.NewAdsCreateCampaignsRequest(a.api, actor)
}

// AdsCreateClients Doc: https://dev.vk.com/method/ads.createClients
func (a *Account) AdsCreateClients(actor actor.Actor) *request.AdsCreateClientsRequest {
	return request.NewAdsCreateClientsRequest(a.api, actor)
}

// AdsCreateLookalikeRequest Doc: https://dev.vk.com/method/ads.createLookalikeRequest
func (a *Account) AdsCreateLookalikeRequest(actor actor.Actor) *request.AdsCreateLookalikeRequest {
	return request.NewAdsCreateLookalikeRequest(a.api, actor)
}

// AdsCreateTargetGroup Doc: https://dev.vk.com/method/ads.createTargetGroup
func (a *Account) AdsCreateTargetGroup(actor actor.Actor) *request.AdsCreateTargetGroupRequest {
	return request.NewAdsCreateTargetGroupRequest(a.api, actor)
}

// AdsCreateTargetPixel Doc: https://dev.vk.com/method/ads.createTargetPixel
func (a *Account) AdsCreateTargetPixel(actor actor.Actor) *request.AdsCreateTargetPixelRequest {
	return request.NewAdsCreateTargetPixelRequest(a.api, actor)
}

// AdsDeleteAds Doc: https://dev.vk.com/method/ads.deleteAds
func (a *Account) AdsDeleteAds(actor actor.Actor) *request.AdsDeleteAdsRequest {
	return request.NewAdsDeleteAdsRequest(a.api, actor)
}

// AdsDeleteCampaigns Doc: https://dev.vk.com/method/ads.deleteCampaigns
func (a *Account) AdsDeleteCampaigns(actor actor.Actor) *request.AdsDeleteCampaignsRequest {
	return request.NewAdsDeleteCampaignsRequest(a.api, actor)
}

// AdsDeleteClients Doc: https://dev.vk.com/method/ads.deleteClients
func (a *Account) AdsDeleteClients(actor actor.Actor) *request.AdsDeleteClientsRequest {
	return request.NewAdsDeleteClientsRequest(a.api, actor)
}

// AdsDeleteTargetGroup Doc: https://dev.vk.com/method/ads.deleteTargetGroup
func (a *Account) AdsDeleteTargetGroup(actor actor.Actor) *request.AdsDeleteTargetGroupRequest {
	return request.NewAdsDeleteTargetGroupRequest(a.api, actor)
}

// AdsDeleteTargetPixel Doc: https://dev.vk.com/method/ads.deleteTargetPixel
func (a *Account) AdsDeleteTargetPixel(actor actor.Actor) *request.AdsDeleteTargetPixelRequest {
	return request.NewAdsDeleteTargetPixelRequest(a.api, actor)
}

// AdsGetAccounts Doc: https://dev.vk.com/method/ads.getAccounts
func (a *Account) AdsGetAccounts(actor actor.Actor) *request.AdsGetAccountsRequest {
	return request.NewAdsGetAccountsRequest(a.api, actor)
}

// AdsGetAds Doc: https://dev.vk.com/method/ads.getAds
func (a *Account) AdsGetAds(actor actor.Actor) *request.AdsGetAdsRequest {
	return request.NewAdsGetAdsRequest(a.api, actor)
}

// AdsGetAdsLayout Doc: https://dev.vk.com/method/ads.getAdsLayout
func (a *Account) AdsGetAdsLayout(actor actor.Actor) *request.AdsGetAdsLayoutRequest {
	return request.NewAdsGetAdsLayoutRequest(a.api, actor)
}

// AdsGetAdsTargeting Doc: https://dev.vk.com/method/ads.getAdsTargeting
func (a *Account) AdsGetAdsTargeting(actor actor.Actor) *request.AdsGetAdsTargetingRequest {
	return request.NewAdsGetAdsTargetingRequest(a.api, actor)
}

// AdsGetBudget Doc: https://dev.vk.com/method/ads.getBudget
func (a *Account) AdsGetBudget(actor actor.Actor) *request.AdsGetBudgetRequest {
	return request.NewAdsGetBudgetRequest(a.api, actor)
}

// AdsGetCampaigns Doc: https://dev.vk.com/method/ads.getCampaigns
func (a *Account) AdsGetCampaigns(actor actor.Actor) *request.AdsGetCampaignsRequest {
	return request.NewAdsGetCampaignsRequest(a.api, actor)
}

// AdsGetCategories Doc: https://dev.vk.com/method/ads.getCategories
func (a *Account) AdsGetCategories(actor actor.Actor) *request.AdsGetCategoriesRequest {
	return request.NewAdsGetCategoriesRequest(a.api, actor)
}

// AdsGetClients Doc: https://dev.vk.com/method/ads.getClients
func (a *Account) AdsGetClients(actor actor.Actor) *request.AdsGetClientsRequest {
	return request.NewAdsGetClientsRequest(a.api, actor)
}

// AdsGetDemographics Doc: https://dev.vk.com/method/ads.getDemographics
func (a *Account) AdsGetDemographics(actor actor.Actor) *request.AdsGetDemographicsRequest {
	return request.NewAdsGetDemographicsRequest(a.api, actor)
}

// AdsGetFloodStats Doc: https://dev.vk.com/method/ads.getFloodStats
func (a *Account) AdsGetFloodStats(actor actor.Actor) *request.AdsGetFloodStatsRequest {
	return request.NewAdsGetFloodStatsRequest(a.api, actor)
}

// AdsGetLookalikeRequests Doc: https://dev.vk.com/method/ads.getLookalikeRequests
func (a *Account) AdsGetLookalikeRequests(actor actor.Actor) *request.AdsGetLookalikeRequestsRequest {
	return request.NewAdsGetLookalikeRequestsRequest(a.api, actor)
}

// AdsGetMusicians Doc: https://dev.vk.com/method/ads.getMusicians
func (a *Account) AdsGetMusicians(actor actor.Actor) *request.AdsGetMusiciansRequest {
	return request.NewAdsGetMusiciansRequest(a.api, actor)
}

// AdsGetMusiciansByIds Doc: https://dev.vk.com/method/ads.getMusiciansByIds
func (a *Account) AdsGetMusiciansByIds(actor actor.Actor) *request.AdsGetMusiciansByIdsRequest {
	return request.NewAdsGetMusiciansByIdsRequest(a.api, actor)
}

// AdsGetOfficeUsers Doc: https://dev.vk.com/method/ads.getOfficeUsers
func (a *Account) AdsGetOfficeUsers(actor actor.Actor) *request.AdsGetOfficeUsersRequest {
	return request.NewAdsGetOfficeUsersRequest(a.api, actor)
}

// AdsGetPostsReach Doc: https://dev.vk.com/method/ads.getPostsReach
func (a *Account) AdsGetPostsReach(actor actor.Actor) *request.AdsGetPostsReachRequest {
	return request.NewAdsGetPostsReachRequest(a.api, actor)
}

// AdsGetRejectionReason Doc: https://dev.vk.com/method/ads.getRejectionReason
func (a *Account) AdsGetRejectionReason(actor actor.Actor) *request.AdsGetRejectionReasonRequest {
	return request.NewAdsGetRejectionReasonRequest(a.api, actor)
}

// AdsGetStatistics Doc: https://dev.vk.com/method/ads.getStatistics
func (a *Account) AdsGetStatistics(actor actor.Actor) *request.AdsGetStatisticsRequest {
	return request.NewAdsGetStatisticsRequest(a.api, actor)
}

// AdsGetSuggestions Doc: https://dev.vk.com/method/ads.getSuggestions
func (a *Account) AdsGetSuggestions(actor actor.Actor) *request.AdsGetSuggestionsRequest {
	return request.NewAdsGetSuggestionsRequest(a.api, actor)
}

// AdsGetTargetGroups Doc: https://dev.vk.com/method/ads.getTargetGroups
func (a *Account) AdsGetTargetGroups(actor actor.Actor) *request.AdsGetTargetGroupsRequest {
	return request.NewAdsGetTargetGroupsRequest(a.api, actor)
}

// AdsGetTargetPixels Doc: https://dev.vk.com/method/ads.getTargetPixels
func (a *Account) AdsGetTargetPixels(actor actor.Actor) *request.AdsGetTargetPixelsRequest {
	return request.NewAdsGetTargetPixelsRequest(a.api, actor)
}

// AdsGetTargetingStats Doc: https://dev.vk.com/method/ads.getTargetingStats
func (a *Account) AdsGetTargetingStats(actor actor.Actor) *request.AdsGetTargetingStatsRequest {
	return request.NewAdsGetTargetingStatsRequest(a.api, actor)
}

// AdsGetUploadURL Doc: https://dev.vk.com/method/ads.getUploadURL
func (a *Account) AdsGetUploadURL(actor actor.Actor) *request.AdsGetUploadURLRequest {
	return request.NewAdsGetUploadURLRequest(a.api, actor)
}

// AdsGetVideoUploadURL Doc: https://dev.vk.com/method/ads.getVideoUploadURL
func (a *Account) AdsGetVideoUploadURL(actor actor.Actor) *request.AdsGetVideoUploadURLRequest {
	return request.NewAdsGetVideoUploadURLRequest(a.api, actor)
}

// AdsImportTargetContacts Doc: https://dev.vk.com/method/ads.importTargetContacts
func (a *Account) AdsImportTargetContacts(actor actor.Actor) *request.AdsImportTargetContactsRequest {
	return request.NewAdsImportTargetContactsRequest(a.api, actor)
}

// AdsRemoveOfficeUsers Doc: https://dev.vk.com/method/ads.removeOfficeUsers
func (a *Account) AdsRemoveOfficeUsers(actor actor.Actor) *request.AdsRemoveOfficeUsersRequest {
	return request.NewAdsRemoveOfficeUsersRequest(a.api, actor)
}

// AdsRemoveTargetContacts Doc: https://dev.vk.com/method/ads.removeTargetContacts
func (a *Account) AdsRemoveTargetContacts(actor actor.Actor) *request.AdsRemoveTargetContactsRequest {
	return request.NewAdsRemoveTargetContactsRequest(a.api, actor)
}

// AdsSaveLookalikeRequestResult Doc: https://dev.vk.com/method/ads.saveLookalikeRequestResult
func (a *Account) AdsSaveLookalikeRequestResult(actor actor.Actor) *request.AdsSaveLookalikeRequestResultRequest {
	return request.NewAdsSaveLookalikeRequestResultRequest(a.api, actor)
}

// AdsShareTargetGroup Doc: https://dev.vk.com/method/ads.shareTargetGroup
func (a *Account) AdsShareTargetGroup(actor actor.Actor) *request.AdsShareTargetGroupRequest {
	return request.NewAdsShareTargetGroupRequest(a.api, actor)
}

// AdsUpdateAds Doc: https://dev.vk.com/method/ads.updateAds
func (a *Account) AdsUpdateAds(actor actor.Actor) *request.AdsUpdateAdsRequest {
	return request.NewAdsUpdateAdsRequest(a.api, actor)
}

// AdsUpdateCampaigns Doc: https://dev.vk.com/method/ads.updateCampaigns
func (a *Account) AdsUpdateCampaigns(actor actor.Actor) *request.AdsUpdateCampaignsRequest {
	return request.NewAdsUpdateCampaignsRequest(a.api, actor)
}

// AdsUpdateClients Doc: https://dev.vk.com/method/ads.updateClients
func (a *Account) AdsUpdateClients(actor actor.Actor) *request.AdsUpdateClientsRequest {
	return request.NewAdsUpdateClientsRequest(a.api, actor)
}

// AdsUpdateOfficeUsers Doc: https://dev.vk.com/method/ads.updateOfficeUsers
func (a *Account) AdsUpdateOfficeUsers(actor actor.Actor) *request.AdsUpdateOfficeUsersRequest {
	return request.NewAdsUpdateOfficeUsersRequest(a.api, actor)
}

// AdsUpdateTargetGroup Doc: https://dev.vk.com/method/ads.updateTargetGroup
func (a *Account) AdsUpdateTargetGroup(actor actor.Actor) *request.AdsUpdateTargetGroupRequest {
	return request.NewAdsUpdateTargetGroupRequest(a.api, actor)
}

// AdsUpdateTargetPixel Doc: https://dev.vk.com/method/ads.updateTargetPixel
func (a *Account) AdsUpdateTargetPixel(actor actor.Actor) *request.AdsUpdateTargetPixelRequest {
	return request.NewAdsUpdateTargetPixelRequest(a.api, actor)
}
