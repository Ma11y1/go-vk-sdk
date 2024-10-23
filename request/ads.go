package request

import (
	"context"
	"encoding/json"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/objects"
	"go-vk-sdk/response"
	"strconv"
	"strings"
)

// Doc: https://dev.vk.com/ru/method/ads

// AdsAddOfficeUsersRequest defines the request for ads.addOfficeUsers
// Adds administrators and/or observers to the advertising account.
// Doc: https://dev.vk.com/method/ads.addOfficeUsers
type AdsAddOfficeUsersRequest struct {
	BaseRequest
}

// NewAdsAddOfficeUsersRequest creates a new request for ads.getProfileInfo
func NewAdsAddOfficeUsersRequest(a *api.API, actor actor.Actor) *AdsAddOfficeUsersRequest {
	return &AdsAddOfficeUsersRequest{*NewMethodBaseRequest(a, actor, "ads.addOfficeUsers")}
}

// Exec executes the request and unmarshals the response into AdsAddOfficeUsersResponse
func (r *AdsAddOfficeUsersRequest) Exec(ctx context.Context) (response response.AdsAddOfficeUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsAddOfficeUsersRequest) AccountID(id int) *AdsAddOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required A serialized JSON array of objects describing the administrators to be added. For a description of user_specification objects, see below
func (r *AdsAddOfficeUsersRequest) Data(us *objects.AdsUserSpecifications) *AdsAddOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required A serialized JSON array of objects describing the administrators to be added. For a description of user_specification objects, see below
func (r *AdsAddOfficeUsersRequest) DataString(us string) *AdsAddOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsCheckLinkRequest defines the request for ads.checkLink
// Checks the link to the advertised objects.
// Doc: https://dev.vk.com/method/ads.checkLink
type AdsCheckLinkRequest struct {
	BaseRequest
}

// NewAdsCheckLinkRequest creates a new request for ads.checkLink
func NewAdsCheckLinkRequest(a *api.API, actor actor.Actor) *AdsCheckLinkRequest {
	return &AdsCheckLinkRequest{*NewMethodBaseRequest(a, actor, "ads.checkLink")}
}

// Exec executes the request and unmarshals the response into AdsCheckLinkResponse
func (r *AdsCheckLinkRequest) Exec(ctx context.Context) (response response.AdsCheckLinkResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCheckLinkRequest) AccountID(id int) *AdsCheckLinkRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// LinkType Required EventType of advertised objects:
// community - community;
// post — post in the community;
// application - VK app;
// video - video recording;
// site — external site.
func (r *AdsCheckLinkRequest) LinkType(linkType constants.AdsLinkType) *AdsCheckLinkRequest {
	r.parameters.Set(constants.ParameterNameLinkType, string(linkType))
	return r
}

// LinkURL Required Link to the advertised item
func (r *AdsCheckLinkRequest) LinkURL(url string) *AdsCheckLinkRequest {
	r.parameters.Set(constants.ParameterNameLinkURL, url)
	return r
}

// CampaignID ID of the campaign in which the ad will be created.
func (r *AdsCheckLinkRequest) CampaignID(id int) *AdsCheckLinkRequest {
	r.parameters.Set(constants.ParameterNameCampaignID, strconv.Itoa(id))
	return r
}

// AdsCreateAdsRequest defines the request for ads.createAds
// Creates new ads.
// Doc: https://dev.vk.com/method/ads.createAds
type AdsCreateAdsRequest struct {
	BaseRequest
}

// NewAdsCreateAdsRequest creates a new request for ads.createAds
func NewAdsCreateAdsRequest(a *api.API, actor actor.Actor) *AdsCreateAdsRequest {
	return &AdsCreateAdsRequest{*NewMethodBaseRequest(a, actor, "ads.createAds")}
}

// Exec executes the request and unmarshals the response into AdsCreateAdsResponse
func (r *AdsCreateAdsRequest) Exec(ctx context.Context) (response response.AdsCreateAdsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCreateAdsRequest) AccountID(id int) *AdsCreateAdsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
// For a description of ad_specification objects, see below.
// The maximum allowed number of advertisements created using one request is 5.
// If you need more, you can use execute.
func (r *AdsCreateAdsRequest) Data(us *objects.AdsAdSpecifications) *AdsCreateAdsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
// For a description of ad_specification objects, see below.
// The maximum allowed number of advertisements created using one request is 5.
// If you need more, you can use execute.
func (r *AdsCreateAdsRequest) DataString(us string) *AdsCreateAdsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsCreateCampaignsRequest defines the request for ads.createCampaigns
// Creates new campaigns.
// Doc: https://dev.vk.com/method/ads.createCampaigns
type AdsCreateCampaignsRequest struct {
	BaseRequest
}

// NewAdsCreateCampaignsRequest creates a new request for ads.createCampaigns
func NewAdsCreateCampaignsRequest(a *api.API, actor actor.Actor) *AdsCreateCampaignsRequest {
	return &AdsCreateCampaignsRequest{*NewMethodBaseRequest(a, actor, "ads.createCampaigns")}
}

// Exec executes the request and unmarshals the response into AdsCreateCampaignsResponse
func (r *AdsCreateCampaignsRequest) Exec(ctx context.Context) (response response.AdsCreateCampaignsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCreateCampaignsRequest) AccountID(id int) *AdsCreateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
func (r *AdsCreateCampaignsRequest) Data(us *objects.AdsCampaignSpecifications) *AdsCreateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
func (r *AdsCreateCampaignsRequest) DataString(us string) *AdsCreateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsCreateClientsRequest defines the request for ads.createClients
// The method creates advertising agency clients. Available only to advertising agencies.
// Doc: https://dev.vk.com/method/ads.createClients
type AdsCreateClientsRequest struct {
	BaseRequest
}

// NewAdsCreateClientsRequest creates a new request for ads.createClients
func NewAdsCreateClientsRequest(a *api.API, actor actor.Actor) *AdsCreateClientsRequest {
	return &AdsCreateClientsRequest{*NewMethodBaseRequest(a, actor, "ads.createClients")}
}

// Exec executes the request and unmarshals the response into AdsCreateClientsResponse
func (r *AdsCreateClientsRequest) Exec(ctx context.Context) (response response.AdsCreateClientsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCreateClientsRequest) AccountID(id int) *AdsCreateClientsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
func (r *AdsCreateClientsRequest) Data(us *objects.AdsClientSpecifications) *AdsCreateClientsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
func (r *AdsCreateClientsRequest) DataString(us string) *AdsCreateClientsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsCreateLookalikeRequest defines the request for ads.createLookalikeRequest
// Creates a lookalike audience request.
// Doc: https://dev.vk.com/method/ads.createLookalikeRequest
type AdsCreateLookalikeRequest struct {
	BaseRequest
}

// NewAdsCreateLookalikeRequest creates a new request for ads.createLookalikeRequest
func NewAdsCreateLookalikeRequest(a *api.API, actor actor.Actor) *AdsCreateLookalikeRequest {
	return &AdsCreateLookalikeRequest{*NewMethodBaseRequest(a, actor, "ads.createLookalikeRequest")}
}

// Exec executes the request and unmarshals the response into AdsCreateLookalikeRequestResponse
func (r *AdsCreateLookalikeRequest) Exec(ctx context.Context) (response response.AdsCreateLookalikeRequestResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCreateLookalikeRequest) AccountID(id int) *AdsCreateLookalikeRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID Advertising client ID
// Only for advertising agencies.
// ID of the client for which the audience will be created.
func (r *AdsCreateLookalikeRequest) ClientID(id int) *AdsCreateLookalikeRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// SourceType The source type of the original audience. At the moment, retargeting_group can take a single value.
func (r *AdsCreateLookalikeRequest) SourceType(t string) *AdsCreateLookalikeRequest {
	r.parameters.Set(constants.ParameterNameSourceType, t)
	return r
}

// RetargetingGroupID Required Advertising account ID
func (r *AdsCreateLookalikeRequest) RetargetingGroupID(id int) *AdsCreateLookalikeRequest {
	r.parameters.Set(constants.ParameterNameRetargetingGroupID, strconv.Itoa(id))
	return r
}

// AdsCreateTargetGroupRequest defines the request for ads.createTargetGroup
// Creates an audience for retargeting advertisements to users who visited the advertiser’s website (viewed product information, registered, etc.).
// Doc: https://dev.vk.com/method/ads.createTargetGroup
type AdsCreateTargetGroupRequest struct {
	BaseRequest
}

// NewAdsCreateTargetGroupRequest creates a new request for ads.createTargetGroup
func NewAdsCreateTargetGroupRequest(a *api.API, actor actor.Actor) *AdsCreateTargetGroupRequest {
	return &AdsCreateTargetGroupRequest{*NewMethodBaseRequest(a, actor, "ads.createTargetGroup")}
}

// Exec executes the request and unmarshals the response into AdsCreateTargetGroupResponse
func (r *AdsCreateTargetGroupRequest) Exec(ctx context.Context) (response response.AdsCreateTargetGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID Required Advertising account ID
func (r *AdsCreateTargetGroupRequest) AccountID(id int) *AdsCreateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the parameter client_id (for agencies only)
func (r *AdsCreateTargetGroupRequest) ClientID(id int) *AdsCreateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// Name - method to set the required parameter name
func (r *AdsCreateTargetGroupRequest) Name(name string) *AdsCreateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameName, name)
	return r
}

// Lifetime - method to set the required parameter lifetime
func (r *AdsCreateTargetGroupRequest) Lifetime(days int) *AdsCreateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameLifeTime, strconv.Itoa(days))
	return r
}

// TargetPixelID - method to set the parameter target_pixel_id
func (r *AdsCreateTargetGroupRequest) TargetPixelID(id int) *AdsCreateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameTargetPixelID, strconv.Itoa(id))
	return r
}

// TargetPixelRules - method to set the parameter target_pixel_rules
// [
//
//	{"type": args},
//	{"type": args},
//	...
//	{"type": args}
//
// ](
//
//	{"type": args},
//	{"type": args},
//	...
//	{"type": args}
//
// )
func (r *AdsCreateTargetGroupRequest) TargetPixelRules(rules []map[string]interface{}) *AdsCreateTargetGroupRequest {
	rulesJSON, _ := json.Marshal(rules)
	r.parameters.Set(constants.ParameterNameTargetPixelRules, string(rulesJSON))
	return r
}

// AdsCreateTargetPixelRequest defines the request for ads.createTargetPixel
// Creates a target pixel.
// Doc: https://dev.vk.com/method/ads.createTargetPixel
type AdsCreateTargetPixelRequest struct {
	BaseRequest
}

// NewAdsCreateTargetPixelRequest creates a new request for ads.createTargetPixel
func NewAdsCreateTargetPixelRequest(a *api.API, actor actor.Actor) *AdsCreateTargetPixelRequest {
	return &AdsCreateTargetPixelRequest{*NewMethodBaseRequest(a, actor, "ads.createTargetPixel")}
}

// Exec executes the request and unmarshals the response into AdsCreateTargetPixelResponse
func (r *AdsCreateTargetPixelRequest) Exec(ctx context.Context) (response response.AdsCreateTargetPixelResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsCreateTargetPixelRequest) AccountID(id int) *AdsCreateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the parameter client_id (for agencies only)
func (r *AdsCreateTargetPixelRequest) ClientID(id int) *AdsCreateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// Name - method to set the required parameter name. Max length 64
func (r *AdsCreateTargetPixelRequest) Name(name string) *AdsCreateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameName, name)
	return r
}

// Domain - method to set the required parameter domain. The domain of the site on which the pixel will be placed.
func (r *AdsCreateTargetPixelRequest) Domain(name string) *AdsCreateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameDomain, name)
	return r
}

// CategoryID - method to set the required parameter category_id. Category ID of the site on which the pixel will be placed. To get a list of possible IDs, use the ads.getSuggestions method (interest_categories_v2 section).
func (r *AdsCreateTargetPixelRequest) CategoryID(id int) *AdsCreateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameCategoryID, strconv.Itoa(id))
	return r
}

// AdsDeleteAdsRequest defines the request for ads.deleteAds
// Archives advertisements.
// Doc: https://dev.vk.com/method/ads.deleteAds
type AdsDeleteAdsRequest struct {
	BaseRequest
}

// NewAdsDeleteAdsRequest creates a new request for ads.deleteAds
func NewAdsDeleteAdsRequest(a *api.API, actor actor.Actor) *AdsDeleteAdsRequest {
	return &AdsDeleteAdsRequest{*NewMethodBaseRequest(a, actor, "ads.deleteAds")}
}

// Exec executes the request and unmarshals the response into AdsDeleteAdsResponse
func (r *AdsDeleteAdsRequest) Exec(ctx context.Context) (response response.AdsDeleteAdsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsDeleteAdsRequest) AccountID(id int) *AdsDeleteAdsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDs - method to set the required parameter IDs. A serialized JSON array containing ad IDs.
func (r *AdsDeleteAdsRequest) IDs(ids []int) *AdsDeleteAdsRequest {
	if len(ids) > 0 {
		adsJSON, _ := json.Marshal(ids)
		r.parameters.Set(constants.ParameterNameIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameIDs, "null")
	}
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsDeleteAdsRequest) IDsString(ids string) *AdsDeleteAdsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// AdsDeleteCampaignsRequest defines the request for ads.deleteCampaigns
// Archives advertising campaigns.
// Doc: https://dev.vk.com/method/ads.deleteCampaigns
type AdsDeleteCampaignsRequest struct {
	BaseRequest
}

// NewAdsDeleteCampaignsRequest creates a new request for ads.deleteCampaigns
func NewAdsDeleteCampaignsRequest(a *api.API, actor actor.Actor) *AdsDeleteCampaignsRequest {
	return &AdsDeleteCampaignsRequest{*NewMethodBaseRequest(a, actor, "ads.deleteCampaigns")}
}

// Exec executes the request and unmarshals the response into AdsDeleteCampaignsResponse
func (r *AdsDeleteCampaignsRequest) Exec(ctx context.Context) (response response.AdsDeleteCampaignsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsDeleteCampaignsRequest) AccountID(id int) *AdsDeleteCampaignsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDs - method to set the required parameter IDs. A serialized JSON array containing ad IDs.
func (r *AdsDeleteCampaignsRequest) IDs(ids []int) *AdsDeleteCampaignsRequest {
	if len(ids) > 0 {
		adsJSON, _ := json.Marshal(ids)
		r.parameters.Set(constants.ParameterNameIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameIDs, "null")
	}
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsDeleteCampaignsRequest) IDsString(ids string) *AdsDeleteCampaignsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// AdsDeleteClientsRequest defines the request for ads.deleteClients
// Archives advertising agency clients.
// Doc: https://dev.vk.com/method/ads.deleteClients
type AdsDeleteClientsRequest struct {
	BaseRequest
}

// NewAdsDeleteClientsRequest creates a new request for ads.deleteClients
func NewAdsDeleteClientsRequest(a *api.API, actor actor.Actor) *AdsDeleteClientsRequest {
	return &AdsDeleteClientsRequest{*NewMethodBaseRequest(a, actor, "ads.deleteClients")}
}

// Exec executes the request and unmarshals the response into AdsDeleteClientsResponse
func (r *AdsDeleteClientsRequest) Exec(ctx context.Context) (response response.AdsDeleteClientsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsDeleteClientsRequest) AccountID(id int) *AdsDeleteClientsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDs - method to set the required parameter IDs. A serialized JSON array containing ad IDs.
func (r *AdsDeleteClientsRequest) IDs(ids []int) *AdsDeleteClientsRequest {
	if len(ids) > 0 {
		adsJSON, _ := json.Marshal(ids)
		r.parameters.Set(constants.ParameterNameIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameIDs, "null")
	}
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsDeleteClientsRequest) IDsString(ids string) *AdsDeleteClientsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// AdsDeleteTargetGroupRequest defines the request for ads.deleteTargetGroup
// Removes a retargeting audience.
// Doc: https://dev.vk.com/method/ads.deleteTargetGroup
type AdsDeleteTargetGroupRequest struct {
	BaseRequest
}

// NewAdsDeleteTargetGroupRequest creates a new request for ads.deleteTargetGroup
func NewAdsDeleteTargetGroupRequest(a *api.API, actor actor.Actor) *AdsDeleteTargetGroupRequest {
	return &AdsDeleteTargetGroupRequest{*NewMethodBaseRequest(a, actor, "ads.deleteTargetGroup")}
}

// Exec executes the request and unmarshals the response into AdsDeleteTargetGroupResponse
func (r *AdsDeleteTargetGroupRequest) Exec(ctx context.Context) (response response.AdsDeleteTargetGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsDeleteTargetGroupRequest) AccountID(id int) *AdsDeleteTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the parameter client_id (for agencies only)
func (r *AdsDeleteTargetGroupRequest) ClientID(id int) *AdsDeleteTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the parameter target_group_id. External identifier.
func (r *AdsDeleteTargetGroupRequest) TargetGroupID(id int) *AdsDeleteTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// AdsDeleteTargetPixelRequest defines the request for ads.deleteTargetPixel
// Removes a retargeting pixel.
// Doc: https://dev.vk.com/method/ads.deleteTargetPixel
type AdsDeleteTargetPixelRequest struct {
	BaseRequest
}

// NewAdsDeleteTargetPixelRequest creates a new request for ads.deleteTargetPixel
func NewAdsDeleteTargetPixelRequest(a *api.API, actor actor.Actor) *AdsDeleteTargetPixelRequest {
	return &AdsDeleteTargetPixelRequest{*NewMethodBaseRequest(a, actor, "ads.deleteTargetPixel")}
}

// Exec executes the request and unmarshals the response into AdsDeleteTargetPixelResponse
func (r *AdsDeleteTargetPixelRequest) Exec(ctx context.Context) (response response.AdsDeleteTargetPixelResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account ID
func (r *AdsDeleteTargetPixelRequest) AccountID(id int) *AdsDeleteTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the parameter client_id (for agencies only)
func (r *AdsDeleteTargetPixelRequest) ClientID(id int) *AdsDeleteTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the parameter target_group_id. External identifier.
func (r *AdsDeleteTargetPixelRequest) TargetGroupID(id int) *AdsDeleteTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// AdsGetAccountsRequest defines the request for ads.getAccounts
// Returns advertising accounts.
// Doc: https://dev.vk.com/method/ads.getAccounts
type AdsGetAccountsRequest struct {
	BaseRequest
}

// NewAdsGetAccountsRequest creates a new request for ads.getAccounts
func NewAdsGetAccountsRequest(a *api.API, actor actor.Actor) *AdsGetAccountsRequest {
	return &AdsGetAccountsRequest{*NewMethodBaseRequest(a, actor, "ads.getAccounts")}
}

// Exec executes the request and unmarshals the response into AdsGetAccountsResponse
func (r *AdsGetAccountsRequest) Exec(ctx context.Context) (response response.AdsGetAccountsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdsGetAdsRequest defines the request for ads.getAds
// Returns a list of advertisements.
// If the campaign_ids and ad_ids filters are enabled at the same time, then ads are displayed according to the following principles:
// 1.If the ad ID is specified in ad_ids, then it is displayed.
// 2.If the ad belongs to an advertising campaign whose ID is specified in campaign_ids, then it is displayed.
// 3.If there is an ad with an ID specified in ad_ids and it belongs to a campaign whose ID is specified in campaign_ids, then this advertising campaign is ignored, i.e. Rule 2 doesn't work for her.
// Displays no more than 2,000 ads from no more than 2,000 campaigns.
// Doc: https://dev.vk.com/method/ads.getAds
type AdsGetAdsRequest struct {
	BaseRequest
}

// NewAdsGetAdsRequest creates a new request for ads.getAds
func NewAdsGetAdsRequest(a *api.API, actor actor.Actor) *AdsGetAdsRequest {
	return &AdsGetAdsRequest{*NewMethodBaseRequest(a, actor, "ads.getAds")}
}

// Exec executes the request and unmarshals the response into AdsGetAdsResponse
func (r *AdsGetAdsRequest) Exec(ctx context.Context) (response response.AdsGetAdsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetAdsRequest) AccountID(id int) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (for agencies only)
func (r *AdsGetAdsRequest) ClientID(id int) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// IncludeDeleted - method to set the optional parameter include_deleted (0 for active ads, 1 for all ads)
func (r *AdsGetAdsRequest) IncludeDeleted(flag bool) *AdsGetAdsRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "0")
	}
	return r
}

// OnlyDeleted - method to set the optional parameter only_deleted (1 for only deleted ads)
func (r *AdsGetAdsRequest) OnlyDeleted(flag bool) *AdsGetAdsRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "0")
	}
	return r
}

// CampaignIDs - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsRequest) CampaignIDs(campaigns []int) *AdsGetAdsRequest {
	if len(campaigns) > 0 {
		campaignsJSON, _ := json.Marshal(campaigns)
		r.parameters.Set(constants.ParameterNameCampaignIDs, string(campaignsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameCampaignIDs, "null")
	}
	return r
}

// CampaignIDsString - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsRequest) CampaignIDsString(ids string) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameCampaignIDs, ids)
	return r
}

// AdIDs - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsRequest) AdIDs(ads []int) *AdsGetAdsRequest {
	if len(ads) > 0 {
		adsJSON, _ := json.Marshal(ads)
		r.parameters.Set(constants.ParameterNameAdIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameAdIDs, "null")
	}
	return r
}

// AdIDsString - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsRequest) AdIDsString(ids string) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameAdIDs, ids)
	return r
}

// Limit - method to set the optional parameter limit (maximum number of ads returned)
func (r *AdsGetAdsRequest) Limit(limit int) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameLimit, strconv.Itoa(limit))
	return r
}

// Offset - method to set the optional parameter offset (used for pagination)
func (r *AdsGetAdsRequest) Offset(offset int) *AdsGetAdsRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// AdsGetAdsLayoutRequest defines the request for ads.getAdsLayout
// Returns descriptions of ad layouts.
// If the campaign_ids and ad_ids filters are enabled at the same time, then ads are displayed according to the following principles:
// 1.If the ad ID is specified in ad_ids, then it is displayed.
// 2.If the ad belongs to an advertising campaign whose ID is specified in campaign_ids, then it is displayed.
// 3.If there is an ad with an ID specified in ad_ids and it belongs to a campaign whose ID is specified in campaign_ids, then this advertising campaign is ignored, i.e. Rule 2 doesn't work for her.
// Displays no more than 2,000 ads from no more than 2,000 campaigns.
// Doc: https://dev.vk.com/method/ads.getAdsLayout
type AdsGetAdsLayoutRequest struct {
	BaseRequest
}

// NewAdsGetAdsLayoutRequest creates a new request for ads.getAdsLayout
func NewAdsGetAdsLayoutRequest(a *api.API, actor actor.Actor) *AdsGetAdsLayoutRequest {
	return &AdsGetAdsLayoutRequest{*NewMethodBaseRequest(a, actor, "ads.getAdsLayout")}
}

// Exec executes the request and unmarshals the response into AdsGetAdsLayoutResponse
func (r *AdsGetAdsLayoutRequest) Exec(ctx context.Context) (response response.AdsGetAdsLayoutResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetAdsLayoutRequest) AccountID(id int) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (for agencies only)
func (r *AdsGetAdsLayoutRequest) ClientID(id int) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// IncludeDeleted - method to set the optional parameter include_deleted (0 for active ads, 1 for all ads)
func (r *AdsGetAdsLayoutRequest) IncludeDeleted(flag bool) *AdsGetAdsLayoutRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "0")
	}
	return r
}

// OnlyDeleted - method to set the optional parameter only_deleted (1 for only deleted ads)
func (r *AdsGetAdsLayoutRequest) OnlyDeleted(flag bool) *AdsGetAdsLayoutRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "0")
	}
	return r
}

// CampaignIDs - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsLayoutRequest) CampaignIDs(campaigns []int) *AdsGetAdsLayoutRequest {
	if len(campaigns) > 0 {
		campaignsJSON, _ := json.Marshal(campaigns)
		r.parameters.Set(constants.ParameterNameCampaignIDs, string(campaignsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameCampaignIDs, "null")
	}
	return r
}

// AdIDs - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsLayoutRequest) AdIDs(ads []int) *AdsGetAdsLayoutRequest {
	if len(ads) > 0 {
		adsJSON, _ := json.Marshal(ads)
		r.parameters.Set(constants.ParameterNameAdIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameAdIDs, "null")
	}
	return r
}

// CampaignIDsString - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsLayoutRequest) CampaignIDsString(ids string) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameCampaignIDs, ids)
	return r
}

// AdIDsString - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsLayoutRequest) AdIDsString(ids string) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameAdIDs, ids)
	return r
}

// Limit - method to set the optional parameter limit (maximum number of ads returned)
func (r *AdsGetAdsLayoutRequest) Limit(limit int) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameLimit, strconv.Itoa(limit))
	return r
}

// Offset - method to set the optional parameter offset (used for pagination)
func (r *AdsGetAdsLayoutRequest) Offset(offset int) *AdsGetAdsLayoutRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// AdsGetAdsTargetingRequest defines the request for ads.getAdsTargeting
// Returns ad targeting parameters.
// If the campaign_ids and ad_ids filters are enabled at the same time, the following ads will be displayed:
// 1.If the ad id is specified in ad_ids, then it is displayed.
// 2.If the ad belongs to an advertising campaign whose id is specified in campaign_ids, then it is displayed.
// 3.If there is an ad with an id specified in ad_ids, and it belongs to a campaign whose id is specified in campaign_ids, then this advertising campaign is ignored, i.e. Rule 2 doesn't work for her.
// Displays no more than 2,000 ads from no more than 2,000 campaigns.
// Doc: https://dev.vk.com/method/ads.getAdsTargeting
type AdsGetAdsTargetingRequest struct {
	BaseRequest
}

// NewAdsGetAdsTargetingRequest creates a new request for ads.getAdsTargeting
func NewAdsGetAdsTargetingRequest(a *api.API, actor actor.Actor) *AdsGetAdsTargetingRequest {
	return &AdsGetAdsTargetingRequest{*NewMethodBaseRequest(a, actor, "ads.getAdsTargeting")}
}

// Exec executes the request and unmarshals the response into AdsGetAdsTargetingResponse
func (r *AdsGetAdsTargetingRequest) Exec(ctx context.Context) (response response.AdsGetAdsTargetingResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetAdsTargetingRequest) AccountID(id int) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (for agencies only)
func (r *AdsGetAdsTargetingRequest) ClientID(id int) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// IncludeDeleted - method to set the optional parameter include_deleted (0 for active ads, 1 for all ads)
func (r *AdsGetAdsTargetingRequest) IncludeDeleted(flag bool) *AdsGetAdsTargetingRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "0")
	}
	return r
}

// OnlyDeleted - method to set the optional parameter only_deleted (1 for only deleted ads)
func (r *AdsGetAdsTargetingRequest) OnlyDeleted(flag bool) *AdsGetAdsTargetingRequest {
	if flag {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameOnlyDeleted, "0")
	}
	return r
}

// CampaignIDs - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsTargetingRequest) CampaignIDs(campaigns []int) *AdsGetAdsTargetingRequest {
	if len(campaigns) > 0 {
		campaignsJSON, _ := json.Marshal(campaigns)
		r.parameters.Set(constants.ParameterNameCampaignIDs, string(campaignsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameCampaignIDs, "null")
	}
	return r
}

// AdIDs - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsTargetingRequest) AdIDs(ads []int) *AdsGetAdsTargetingRequest {
	if len(ads) > 0 {
		adsJSON, _ := json.Marshal(ads)
		r.parameters.Set(constants.ParameterNameAdIDs, string(adsJSON))
	} else {
		r.parameters.Set(constants.ParameterNameAdIDs, "null")
	}
	return r
}

// CampaignIDsString - method to set the optional parameter campaign_ids (JSON array of campaign IDs)
func (r *AdsGetAdsTargetingRequest) CampaignIDsString(ids string) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameCampaignIDs, ids)
	return r
}

// AdIDsString - method to set the optional parameter ad_ids (JSON array of ad IDs)
func (r *AdsGetAdsTargetingRequest) AdIDsString(ids string) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameAdIDs, ids)
	return r
}

// Limit - method to set the optional parameter limit (maximum number of ads returned)
func (r *AdsGetAdsTargetingRequest) Limit(limit int) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameLimit, strconv.Itoa(limit))
	return r
}

// Offset - method to set the optional parameter offset (used for pagination)
func (r *AdsGetAdsTargetingRequest) Offset(offset int) *AdsGetAdsTargetingRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// AdsGetBudgetRequest defines the request for ads.getBudget
// Returns the budget of an advertising account.
// Doc: https://dev.vk.com/method/ads.getBudget
type AdsGetBudgetRequest struct {
	BaseRequest
}

// NewAdsGetBudgetRequest creates a new request for ads.getBudget
func NewAdsGetBudgetRequest(a *api.API, actor actor.Actor) *AdsGetBudgetRequest {
	return &AdsGetBudgetRequest{*NewMethodBaseRequest(a, actor, "ads.getBudget")}
}

// Exec executes the request and unmarshals the response into AdsGetBudgetResponse
func (r *AdsGetBudgetRequest) Exec(ctx context.Context) (response response.AdsGetBudgetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetBudgetRequest) AccountID(id int) *AdsGetBudgetRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// AdsGetCampaignsRequest defines the request for ads.getCampaigns
// Returns the advertising campaigns.
// Doc: https://dev.vk.com/method/ads.getCampaigns
type AdsGetCampaignsRequest struct {
	BaseRequest
}

// NewAdsGetCampaignsRequest creates a new request for ads.getCampaigns
func NewAdsGetCampaignsRequest(a *api.API, actor actor.Actor) *AdsGetCampaignsRequest {
	return &AdsGetCampaignsRequest{*NewMethodBaseRequest(a, actor, "ads.getCampaigns")}
}

// Exec executes the request and unmarshals the response into AdsGetCampaignsResponse
func (r *AdsGetCampaignsRequest) Exec(ctx context.Context) (response response.AdsGetCampaignsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetCampaignsRequest) AccountID(id int) *AdsGetCampaignsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id.
// Mandatory for advertising agencies, not used in other cases. ID of the client from whom advertising campaigns are requested.
func (r *AdsGetCampaignsRequest) ClientID(id int) *AdsGetCampaignsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// IncludeDeleted - method to set the optional parameter include_deleted
// Flag specifying whether archived advertisements should be displayed.
// 0 — display only active campaigns
// 1 — display all campaigns
func (r *AdsGetCampaignsRequest) IncludeDeleted(includeDeleted bool) *AdsGetCampaignsRequest {
	if includeDeleted {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "1")
	} else {
		r.parameters.Set(constants.ParameterNameIncludeDeleted, "0")
	}
	return r
}

// CampaignIDs - method to set the optional parameter campaign_ids.
// Filter of displayed advertising campaigns. Serialized JSON array containing campaign ids.
// Only campaigns that are present in campaign_ids and are campaigns of the specified advertising account will be displayed.
// If the parameter is equal to the string null, then all campaigns will be displayed.
func (r *AdsGetCampaignsRequest) CampaignIDs(ids []int) *AdsGetCampaignsRequest {
	idsStr, _ := json.Marshal(ids)
	r.parameters.Set(constants.ParameterNameCampaignIDs, string(idsStr))
	return r
}

// CampaignIDsString - method to set the optional parameter campaign_ids.
// Filter of displayed advertising campaigns. Serialized JSON array containing campaign ids.
// Only campaigns that are present in campaign_ids and are campaigns of the specified advertising account will be displayed.
// If the parameter is equal to the string null, then all campaigns will be displayed.
func (r *AdsGetCampaignsRequest) CampaignIDsString(ids string) *AdsGetCampaignsRequest {
	r.parameters.Set(constants.ParameterNameCampaignIDs, ids)
	return r
}

// Fields - method to set the optional parameter fields.
// Adds additional fields to the response. Supported values:
// ads_count — number of ads in the campaign.
func (r *AdsGetCampaignsRequest) Fields(fields string) *AdsGetCampaignsRequest {
	r.parameters.Set(constants.ParameterNameFields, fields)
	return r
}

// AdsGetCategoriesRequest defines the request for ads.getCategories
// Allows you to get possible advertisement topics.
// Doc: https://dev.vk.com/method/ads.getCategories
type AdsGetCategoriesRequest struct {
	BaseRequest
}

// NewAdsGetCategoriesRequest creates a new request for ads.getCategories
func NewAdsGetCategoriesRequest(a *api.API, actor actor.Actor) *AdsGetCategoriesRequest {
	return &AdsGetCategoriesRequest{*NewMethodBaseRequest(a, actor, "ads.getCategories")}
}

// Exec executes the request and unmarshals the response into AdsGetCategoriesResponse
func (r *AdsGetCategoriesRequest) Exec(ctx context.Context) (response response.AdsGetCategoriesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AdsGetCategoriesRequest) Lang(lang constants.LanguageName) *AdsGetCategoriesRequest {
	r.parameters.Set(constants.ParameterNameLang, string(lang))
	return r
}

// AdsGetClientsRequest defines the request for ads.getClients
// Method for returning a list of advertising agency clients. Available only to advertising agencies.
// Doc: https://dev.vk.com/method/ads.getClients
type AdsGetClientsRequest struct {
	BaseRequest
}

// NewAdsGetClientsRequest creates a new request for ads.getClients
func NewAdsGetClientsRequest(a *api.API, actor actor.Actor) *AdsGetClientsRequest {
	return &AdsGetClientsRequest{*NewMethodBaseRequest(a, actor, "ads.getClients")}
}

// Exec executes the request and unmarshals the response into AdsGetClientsResponse
func (r *AdsGetClientsRequest) Exec(ctx context.Context) (response response.AdsGetClientsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id
func (r *AdsGetClientsRequest) AccountID(id int) *AdsGetClientsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// AdsGetDemographicsRequest defines the request for ads.getDemographics
// Returns demographic statistics for advertisements or campaigns.
// Doc: https://dev.vk.com/method/ads.getDemographics
type AdsGetDemographicsRequest struct {
	BaseRequest
}

// NewAdsGetDemographicsRequest creates a new request for ads.getDemographics
func NewAdsGetDemographicsRequest(a *api.API, actor actor.Actor) *AdsGetDemographicsRequest {
	return &AdsGetDemographicsRequest{*NewMethodBaseRequest(a, actor, "ads.getDemographics")}
}

// Exec executes the request and unmarshals the response into AdsGetDemographicsResponse
func (r *AdsGetDemographicsRequest) Exec(ctx context.Context) (response response.AdsGetDemographicsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetDemographicsRequest) AccountID(id int) *AdsGetDemographicsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDsType - method to set the required parameter ids_type (type of objects requested: "ad" or "campaign")
func (r *AdsGetDemographicsRequest) IDsType(idsType string) *AdsGetDemographicsRequest {
	if idsType == "ad" || idsType == "campaign" {
		r.parameters.Set(constants.ParameterNameIDsType, idsType)
	}
	return r
}

// IDs - method to set the required parameter ids (comma-separated list of ad or campaign IDs)
func (r *AdsGetDemographicsRequest) IDs(ids []int) *AdsGetDemographicsRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsGetDemographicsRequest) IDsString(ids string) *AdsGetDemographicsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// Period - method to set the required parameter period (grouping method: "day", "month", or "overall")
func (r *AdsGetDemographicsRequest) Period(period string) *AdsGetDemographicsRequest {
	if period == "day" || period == "month" || period == "overall" {
		r.parameters.Set(constants.ParameterNamePeriod, period)
	}
	return r
}

// DateFrom - method to set the required parameter date_from (start date for statistics)
// day: YYYY-MM-DD, example: 2011-09-27 - September 27, 2011
// 0—day of creation
// month: YYYY-MM, example: 2011-09 - September 2011
// 0 — month of creation
// overall: 0.
func (r *AdsGetDemographicsRequest) DateFrom(date string) *AdsGetDemographicsRequest {
	r.parameters.Set(constants.ParameterNameDateFrom, date)
	return r
}

// DateTo - method to set the required parameter date_to (end date for statistics)
// The end date of the displayed statistics. Different date formats are used for different values ​​of the period parameter
// day: YYYY-MM-DD, example: 2011-09-27 - September 27, 2011;
// 0 — current day
// month: YYYY-MM, example: 2011-09 - September 2011
// 0 - current month
// overall: 0
func (r *AdsGetDemographicsRequest) DateTo(date string) *AdsGetDemographicsRequest {
	r.parameters.Set(constants.ParameterNameDateTo, date)
	return r
}

// AdsGetFloodStatsRequest defines the request for ads.getFloodStats
// Returns information about the current state of the counter - the number of remaining method runs and the time until the next counter reset in seconds.
// Doc: https://dev.vk.com/method/ads.getFloodStats
type AdsGetFloodStatsRequest struct {
	BaseRequest
}

// NewAdsGetFloodStatsRequest creates a new request for ads.getFloodStats
func NewAdsGetFloodStatsRequest(a *api.API, actor actor.Actor) *AdsGetFloodStatsRequest {
	return &AdsGetFloodStatsRequest{*NewMethodBaseRequest(a, actor, "ads.getFloodStats")}
}

// Exec executes the request and unmarshals the response into AdsGetFloodStatsResponse
func (r *AdsGetFloodStatsRequest) Exec(ctx context.Context) (response response.AdsGetFloodStatsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetFloodStatsRequest) AccountID(id int) *AdsGetFloodStatsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// AdsGetLookalikeRequestsRequest defines the request for ads.getLookalikeRequests
// Returns lookalike audience requests.
// Doc: https://dev.vk.com/method/ads.getLookalikeRequests
type AdsGetLookalikeRequestsRequest struct {
	BaseRequest
}

// NewAdsGetLookalikeRequestsRequest creates a new request for ads.getLookalikeRequests
func NewAdsGetLookalikeRequestsRequest(a *api.API, actor actor.Actor) *AdsGetLookalikeRequestsRequest {
	return &AdsGetLookalikeRequestsRequest{*NewMethodBaseRequest(a, actor, "ads.getLookalikeRequests")}
}

// Exec executes the request and unmarshals the response into AdsGetLookalikeRequestsResponse
func (r *AdsGetLookalikeRequestsRequest) Exec(ctx context.Context) (response response.AdsGetLookalikeRequestsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetLookalikeRequestsRequest) AccountID(id int) *AdsGetLookalikeRequestsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (only for agencies)
func (r *AdsGetLookalikeRequestsRequest) ClientID(id int) *AdsGetLookalikeRequestsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// RequestsIDs - method to set the optional parameter requests_ids (comma-separated list of requested request identifiers. The maximum number of identifiers in the list is 200)
// If this parameter is empty, all requests will be returned.
func (r *AdsGetLookalikeRequestsRequest) RequestsIDs(ids []int) *AdsGetLookalikeRequestsRequest {
	if len(ids) > 0 {
		strIDs := make([]string, len(ids))
		for i, id := range ids {
			strIDs[i] = strconv.Itoa(id)
		}
		r.parameters.Set(constants.ParameterNameRequestIDs, strings.Join(strIDs, ","))
	} else {
		r.parameters.Set(constants.ParameterNameRequestIDs, "")
	}
	return r
}

// Offset - method to set the optional parameter offset (used with limit for pagination)
func (r *AdsGetLookalikeRequestsRequest) Offset(offset int) *AdsGetLookalikeRequestsRequest {
	r.parameters.Set(constants.ParameterNameOffset, strconv.Itoa(offset))
	return r
}

// Limit - method to set the optional parameter limit (number of requests to return)
func (r *AdsGetLookalikeRequestsRequest) Limit(limit int) *AdsGetLookalikeRequestsRequest {
	r.parameters.Set(constants.ParameterNameLimit, strconv.Itoa(limit))
	return r
}

// SortBy - method to set the optional parameter sort_by (sort by "id" or "update_time")
// id – sort by ascending identifiers
// update_time – sort in descending order of the time of the last status update
func (r *AdsGetLookalikeRequestsRequest) SortBy(sortBy string) *AdsGetLookalikeRequestsRequest {
	if sortBy == "id" || sortBy == "update_time" {
		r.parameters.Set(constants.ParameterNameSortBy, sortBy)
	}
	return r
}

// AdsGetMusiciansRequest defines the request for ads.getMusicians
// Returns information about musicians whose listeners can be targeted.
// Doc: https://dev.vk.com/method/ads.getMusicians
type AdsGetMusiciansRequest struct {
	BaseRequest
}

// NewAdsGetMusiciansRequest creates a new request for ads.getMusicians
func NewAdsGetMusiciansRequest(a *api.API, actor actor.Actor) *AdsGetMusiciansRequest {
	return &AdsGetMusiciansRequest{*NewMethodBaseRequest(a, actor, "ads.getMusicians")}
}

// Exec executes the request and unmarshals the response into AdsGetMusiciansResponse
func (r *AdsGetMusiciansRequest) Exec(ctx context.Context) (response response.AdsGetMusiciansResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

func (r *AdsGetMusiciansRequest) ArtistName(name string) *AdsGetMusiciansRequest {
	r.parameters.Set(constants.ParameterNameArtistName, name)
	return r
}

// AdsGetMusiciansByIdsRequest defines the request for ads.getMusiciansByIds
// Returns information about musicians for listeners for whom targeting is available.
// Doc: https://dev.vk.com/method/ads.getMusiciansByIds
type AdsGetMusiciansByIdsRequest struct {
	BaseRequest
}

// NewAdsGetMusiciansByIdsRequest creates a new request for ads.getMusiciansByIds
func NewAdsGetMusiciansByIdsRequest(a *api.API, actor actor.Actor) *AdsGetMusiciansByIdsRequest {
	return &AdsGetMusiciansByIdsRequest{*NewMethodBaseRequest(a, actor, "ads.getMusiciansByIds")}
}

// Exec executes the request and unmarshals the response into AdsGetMusiciansByIdsResponse
func (r *AdsGetMusiciansByIdsRequest) Exec(ctx context.Context) (response response.AdsGetMusiciansByIdsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// IDs - method to set the required parameter ids (comma-separated list of ad or campaign IDs)
func (r *AdsGetMusiciansByIdsRequest) IDs(ids []int) *AdsGetMusiciansByIdsRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsGetMusiciansByIdsRequest) IDsString(ids string) *AdsGetMusiciansByIdsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// AdsGetOfficeUsersRequest defines the request for ads.getOfficeUsers
// Returns a list of administrators and observers of the advertising account.
// Doc: https://dev.vk.com/method/ads.getOfficeUsers
type AdsGetOfficeUsersRequest struct {
	BaseRequest
}

// NewAdsGetOfficeUsersRequest creates a new request for ads.getOfficeUsers
func NewAdsGetOfficeUsersRequest(a *api.API, actor actor.Actor) *AdsGetOfficeUsersRequest {
	return &AdsGetOfficeUsersRequest{*NewMethodBaseRequest(a, actor, "ads.getOfficeUsers")}
}

// Exec executes the request and unmarshals the response into AdsGetOfficeUsersResponse
func (r *AdsGetOfficeUsersRequest) Exec(ctx context.Context) (response response.AdsGetOfficeUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetOfficeUsersRequest) AccountID(id int) *AdsGetOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// AdsGetPostsReachRequest defines the request for ads.getPostsReach
// Returns detailed statistics on the reach of promotional posts from ads and campaigns for promoting community posts.
// Campaign statistics can only be obtained for campaigns created since December 20, 2016.
// Doc: https://dev.vk.com/method/ads.getPostsReach
type AdsGetPostsReachRequest struct {
	BaseRequest
}

// NewAdsGetPostsReachRequest creates a new request for ads.getPostsReach
func NewAdsGetPostsReachRequest(a *api.API, actor actor.Actor) *AdsGetPostsReachRequest {
	return &AdsGetPostsReachRequest{*NewMethodBaseRequest(a, actor, "ads.getPostsReach")}
}

// Exec executes the request and unmarshals the response into AdsGetPostsReachResponse
func (r *AdsGetPostsReachRequest) Exec(ctx context.Context) (response response.AdsGetPostsReachResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetPostsReachRequest) AccountID(id int) *AdsGetPostsReachRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDs - method to set the required parameter ids (comma-separated list of ad or campaign IDs)
func (r *AdsGetPostsReachRequest) IDs(ids []int) *AdsGetPostsReachRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsGetPostsReachRequest) IDsString(ids string) *AdsGetPostsReachRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// IDsType - method to set the required parameter ids_type (type of objects requested: "ad" or "campaign")
func (r *AdsGetPostsReachRequest) IDsType(idsType string) *AdsGetPostsReachRequest {
	if idsType == "ad" || idsType == "campaign" {
		r.parameters.Set(constants.ParameterNameIDsType, idsType)
	}
	return r
}

// AdsGetRejectionReasonRequest defines the request for ads.getRejectionReason
// Returns the reason why the specified ad was denied pre-moderation.
// Doc: https://dev.vk.com/method/ads.getRejectionReason
type AdsGetRejectionReasonRequest struct {
	BaseRequest
}

// NewAdsGetRejectionReasonRequest creates a new request for ads.getRejectionReason
func NewAdsGetRejectionReasonRequest(a *api.API, actor actor.Actor) *AdsGetRejectionReasonRequest {
	return &AdsGetRejectionReasonRequest{*NewMethodBaseRequest(a, actor, "ads.getRejectionReason")}
}

// Exec executes the request and unmarshals the response into AdsGetRejectionReasonResponse
func (r *AdsGetRejectionReasonRequest) Exec(ctx context.Context) (response response.AdsGetRejectionReasonResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetRejectionReasonRequest) AccountID(id int) *AdsGetRejectionReasonRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// AdID - method to set the required parameter ad_id
func (r *AdsGetRejectionReasonRequest) AdID(id int) *AdsGetRejectionReasonRequest {
	r.parameters.Set(constants.ParameterNameAdID, strconv.Itoa(id))
	return r
}

// AdsGetStatisticsRequest defines the request for ads.getStatistics
// Returns performance indicator statistics for advertisements, campaigns, clients, or the entire account.
// Doc: https://dev.vk.com/method/ads.getStatistics
type AdsGetStatisticsRequest struct {
	BaseRequest
}

// NewAdsGetStatisticsRequest creates a new request for ads.getStatistics
func NewAdsGetStatisticsRequest(a *api.API, actor actor.Actor) *AdsGetStatisticsRequest {
	return &AdsGetStatisticsRequest{*NewMethodBaseRequest(a, actor, "ads.getStatistics")}
}

// Exec executes the request and unmarshals the response into AdsGetStatisticsResponse
func (r *AdsGetStatisticsRequest) Exec(ctx context.Context) (response response.AdsGetStatisticsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetStatisticsRequest) AccountID(id int) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDsType - method to set the required parameter ids_type (type of requested objects: "ad", "campaign", "client", or "office")
func (r *AdsGetStatisticsRequest) IDsType(idsType constants.AdsStatisticIDType) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameIDsType, string(idsType))
	return r
}

// IDs - method to set the required parameter ids.
// ids of the requested ads, campaigns, clients or account, separated by commas, depending on what is specified in the ids_type parameter. Maximum 2000 objects
func (r *AdsGetStatisticsRequest) IDs(ids []int) *AdsGetStatisticsRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids
func (r *AdsGetStatisticsRequest) IDsString(ids string) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// Period - method to set the required parameter period (grouping by "day", "week", "month", "year", or "overall")
func (r *AdsGetStatisticsRequest) Period(period constants.AdsStatisticPeriod) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNamePeriod, string(period))
	return r
}

// DateFrom - method to set the required parameter date_from (start date of the statistics)
func (r *AdsGetStatisticsRequest) DateFrom(date string) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameDateFrom, date)
	return r
}

// DateTo - method to set the required parameter date_to (end date of the statistics)
func (r *AdsGetStatisticsRequest) DateTo(date string) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameDateTo, date)
	return r
}

// StatsFields - method to set the required parameter ids.
// Additional statistics: views_times - distribution of the number of impressions per user. Available for ad and campaign types that contain only advertising entries.
func (r *AdsGetStatisticsRequest) StatsFields(stats string) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameStatsFields, stats)
	return r
}

// StatsFieldsArr - method to set the required parameter ids.
// Additional statistics: views_times - distribution of the number of impressions per user. Available for ad and campaign types that contain only advertising entries.
func (r *AdsGetStatisticsRequest) StatsFieldsArr(stats []string) *AdsGetStatisticsRequest {
	r.parameters.Set(constants.ParameterNameStatsFields, strings.Join(stats, ","))
	return r
}

// AdsGetSuggestionsRequest defines the request for ads.getSuggestions
// Returns a set of tooltips for various targeting options.
// For sections: countries, positions, interest_categories, religions, user_devices, user_os, user_browsers
// Doc: https://dev.vk.com/method/ads.getSuggestions
type AdsGetSuggestionsRequest struct {
	BaseRequest
}

// NewAdsGetSuggestionsRequest creates a new request for ads.getSuggestions
func NewAdsGetSuggestionsRequest(a *api.API, actor actor.Actor) *AdsGetSuggestionsRequest {
	return &AdsGetSuggestionsRequest{*NewMethodBaseRequest(a, actor, "ads.getSuggestions")}
}

// Exec executes the request and unmarshals the response into AdsGetSuggestionsResponse
func (r *AdsGetSuggestionsRequest) Exec(ctx context.Context) (response response.AdsGetSuggestionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// Section - method to set the required parameter section.
// The section for which tips are requested. The following values are supported:
func (r *AdsGetSuggestionsRequest) Section(ids string) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// IDs - method to set the required parameter ids.
// Object IDs, separated by commas. Used to decrypt IDs returned in the ads.getAdsTargeting method.
// If this parameter is specified, then the parameters q, country, cities should not be passed, thus eliminating their requirement for a specific section. Objects are returned in the same order in which they were specified in this parameter.
func (r *AdsGetSuggestionsRequest) IDs(ids []int) *AdsGetSuggestionsRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids.
// Object IDs, separated by commas. Used to decrypt IDs returned in the ads.getAdsTargeting method.
// If this parameter is specified, then the parameters q, country, cities should not be passed, thus eliminating their requirement for a specific section. Objects are returned in the same order in which they were specified in this parameter.
func (r *AdsGetSuggestionsRequest) IDsString(ids string) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// Query - method to set the required parameter query. Query filter string.
func (r *AdsGetSuggestionsRequest) Query(q string) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameQuery, q)
	return r
}

// Country - method to set the required parameter country. Id of the country in which objects are searched.
func (r *AdsGetSuggestionsRequest) Country(country int) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameCountry, strconv.Itoa(country))
	return r
}

// Cities - method to set the required parameter cities. Ids of cities in which objects are searched, separated by commas.
func (r *AdsGetSuggestionsRequest) Cities(cities []int) *AdsGetSuggestionsRequest {
	strCities := make([]string, len(cities))
	for i, id := range cities {
		strCities[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameCities, strings.Join(strCities, ","))
	return r
}

// CitiesString - method to set the required parameter cities. Ids of cities in which objects are searched, separated by commas.
func (r *AdsGetSuggestionsRequest) CitiesString(cities string) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameCities, cities)
	return r
}

// Lang - method to set the required parameter lang
// The language of the returned string values. Supported languages:
// ru - Russian
// ua - Ukrainian
// en - English
func (r *AdsGetSuggestionsRequest) Lang(lang constants.LanguageName) *AdsGetSuggestionsRequest {
	r.parameters.Set(constants.ParameterNameLang, string(lang))
	return r
}

// AdsGetSuggestionsRegionsRequest defines the request for ads.getSuggestions
// Returns a set of tooltips for various targeting options.
// For the regions section.
// Doc: https://dev.vk.com/method/ads.getSuggestions
type AdsGetSuggestionsRegionsRequest struct {
	AdsGetSuggestionsRequest
}

// NewAdsGetSuggestionsRegionsRequest creates a new request for ads.getSuggestions
func NewAdsGetSuggestionsRegionsRequest(a *api.API, actor actor.Actor) *AdsGetSuggestionsRegionsRequest {
	return &AdsGetSuggestionsRegionsRequest{*NewAdsGetSuggestionsRequest(a, actor)}
}

// Exec executes the request and unmarshals the response into AdsGetSuggestionsResponse
func (r *AdsGetSuggestionsRegionsRequest) Exec(ctx context.Context) (response response.AdsGetSuggestionsRegionsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdsGetSuggestionsCitiesRequest defines the request for ads.getSuggestions
// Returns a set of tooltips for various targeting options.
// For the cities, districts, streets sections.
// Doc: https://dev.vk.com/method/ads.getSuggestions
type AdsGetSuggestionsCitiesRequest struct {
	AdsGetSuggestionsRequest
}

// NewAdsGetSuggestionsCitiesRequest creates a new request for ads.getSuggestions
func NewAdsGetSuggestionsCitiesRequest(a *api.API, actor actor.Actor) *AdsGetSuggestionsCitiesRequest {
	return &AdsGetSuggestionsCitiesRequest{*NewAdsGetSuggestionsRequest(a, actor)}
}

// Exec executes the request and unmarshals the response into AdsGetSuggestionsResponse
func (r *AdsGetSuggestionsCitiesRequest) Exec(ctx context.Context) (response response.AdsGetSuggestionsCitiesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdsGetSuggestionsSchoolsRequest defines the request for ads.getSuggestions
// Returns a set of tooltips for various targeting options.
// For the schools section
// Doc: https://dev.vk.com/method/ads.getSuggestions
type AdsGetSuggestionsSchoolsRequest struct {
	AdsGetSuggestionsRequest
}

// NewAdsGetSuggestionsSchoolsRequest creates a new request for ads.getSuggestions
func NewAdsGetSuggestionsSchoolsRequest(a *api.API, actor actor.Actor) *AdsGetSuggestionsSchoolsRequest {
	return &AdsGetSuggestionsSchoolsRequest{*NewAdsGetSuggestionsRequest(a, actor)}
}

// Exec executes the request and unmarshals the response into AdsGetSuggestionsResponse
func (r *AdsGetSuggestionsSchoolsRequest) Exec(ctx context.Context) (response response.AdsGetSuggestionsSchoolsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdsGetTargetGroupsRequest defines the request for ads.getTargetGroups
// Returns a list of retargeting audiences.
// Doc: https://dev.vk.com/method/ads.getTargetGroups
type AdsGetTargetGroupsRequest struct {
	BaseRequest
}

// NewAdsGetTargetGroupsRequest creates a new request for ads.getTargetGroups
func NewAdsGetTargetGroupsRequest(a *api.API, actor actor.Actor) *AdsGetTargetGroupsRequest {
	return &AdsGetTargetGroupsRequest{*NewMethodBaseRequest(a, actor, "ads.getTargetGroups")}
}

// Exec executes the request and unmarshals the response into AdsGetTargetGroupsResponse
func (r *AdsGetTargetGroupsRequest) Exec(ctx context.Context) (response response.AdsGetTargetGroupsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetTargetGroupsRequest) AccountID(id int) *AdsGetTargetGroupsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (only for agencies)
// ID of the client whose advertising account contains audiences.
func (r *AdsGetTargetGroupsRequest) ClientID(id int) *AdsGetTargetGroupsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// Extended - method to set the optional parameter extended
// If 1, the results will contain a code for posting on the site.
// Deprecated option. Used only for old retargeting groups that were replenished without the help of a pixel. For new audiences it should be omitted.
func (r *AdsGetTargetGroupsRequest) Extended(v bool) *AdsGetTargetGroupsRequest {
	if v {
		r.parameters.Set(constants.ParameterNameClientID, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameClientID)
	}
	return r
}

// AdsGetTargetPixelsRequest defines the request for ads.getTargetPixels
// Returns a list of retargeting pixels.
// Doc: https://dev.vk.com/method/ads.getTargetPixels
type AdsGetTargetPixelsRequest struct {
	BaseRequest
}

// NewAdsGetTargetPixelsRequest creates a new request for ads.getTargetPixels
func NewAdsGetTargetPixelsRequest(a *api.API, actor actor.Actor) *AdsGetTargetPixelsRequest {
	return &AdsGetTargetPixelsRequest{*NewMethodBaseRequest(a, actor, "ads.getTargetPixels")}
}

// Exec executes the request and unmarshals the response into AdsGetTargetPixelsResponse
func (r *AdsGetTargetPixelsRequest) Exec(ctx context.Context) (response response.AdsGetTargetPixelsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetTargetPixelsRequest) AccountID(id int) *AdsGetTargetPixelsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the optional parameter client_id (only for agencies)
// ID of the client in whose advertising account the pixels are located.
func (r *AdsGetTargetPixelsRequest) ClientID(id int) *AdsGetTargetPixelsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// AdsGetTargetingStatsRequest defines the request for ads.getTargetingStats
// Returns the target audience size of the targeting, as well as recommended CPC and CPM values.
// Doc: https://dev.vk.com/method/ads.getTargetingStats
type AdsGetTargetingStatsRequest struct {
	BaseRequest
}

// NewAdsGetTargetingStatsRequest creates a new request for ads.getTargetingStats
func NewAdsGetTargetingStatsRequest(a *api.API, actor actor.Actor) *AdsGetTargetingStatsRequest {
	return &AdsGetTargetingStatsRequest{*NewMethodBaseRequest(a, actor, "ads.getTargetingStats")}
}

// Exec executes the request and unmarshals the response into AdsGetTargetingStatisticResponse
func (r *AdsGetTargetingStatsRequest) Exec(ctx context.Context) (response response.AdsGetTargetingStatisticResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsGetTargetingStatsRequest) AccountID(id int) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
func (r *AdsGetTargetingStatsRequest) ClientID(id int) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// Criteria - method to set the optional parameter criteria (objects.AdsTargetStatsCriteria objects for targeting)
// Ignored if ad_id is specified
func (r *AdsGetTargetingStatsRequest) Criteria(criteria *objects.AdsTargetStatsCriteria) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameCriteria, criteria.ToJSON())
	return r
}

// CriteriaString - method to set the optional parameter criteria (objects.AdsTargetStatsCriteria objects for targeting)
// Ignored if ad_id is specified
func (r *AdsGetTargetingStatsRequest) CriteriaString(criteria string) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameCriteria, criteria)
	return r
}

// AdID - method to set the optional parameter ad_id (The ID of the ad whose targeting parameters you want to analyze. If this parameter is specified, criteria is ignored)
func (r *AdsGetTargetingStatsRequest) AdID(id int) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameAdID, strconv.Itoa(id))
	return r
}

// AdFormat - method to set the required parameter ad_format (ad format)
// 1 - image and text
// 2 - large image
// 3 - exclusive format
// 4 — promotion of communities or applications, square image
// 7 - special application format
// 8 - special format of communities
// 9 - entry in the community
// 10 — application showcase
// 11 - adaptive format.
func (r *AdsGetTargetingStatsRequest) AdFormat(format constants.AdsTargetingStatisticAdFormat) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameAdFormat, strconv.Itoa(int(format)))
	return r
}

// AdPlatform - method to set the optional parameter ad_platform
// Advertising platforms on which the ad will be shown:
// (if ad_format is 1)
//
//	0 — VK and partner sites
//	1 - VK only
//
// (if ad_format is 9)
//
//	"all" — all sites
//	"desktop" - full version of the site
//	"mobile" - mobile website and applications
//
// (if ad_format is 11)
//
//	"all" — all sites
//	"desktop" - full version of the site
//	"mobile" - mobile website and applications.
func (r *AdsGetTargetingStatsRequest) AdPlatform(platform string) *AdsGetTargetingStatsRequest {
	r.parameters.Set("ad_platform", platform)
	return r
}

// AdPlatformNoWall - method to set the optional parameter ad_platform_no_wall (exclude ad from community walls)
// For ad_format = 9 and 11. 1 - do not show the ad on community walls. Default: 0.
func (r *AdsGetTargetingStatsRequest) AdPlatformNoWall(value bool) *AdsGetTargetingStatsRequest {
	if value {
		r.parameters.Set(constants.ParameterNameAdPlatformNoWall, "1")
	} else {
		r.parameters.Set(constants.ParameterNameAdPlatformNoWall, "0")
	}
	return r
}

// AdPlatformNoAdNetwork - method to set the optional parameter ad_platform_no_ad_network
// For ad_format = 9 and 11. 1 - do not show the ad in the advertising network. Default: 0.
func (r *AdsGetTargetingStatsRequest) AdPlatformNoAdNetwork(value bool) *AdsGetTargetingStatsRequest {
	if value {
		r.parameters.Set(constants.ParameterNameAdPlatformNoAdNetwork, "1")
	} else {
		r.parameters.Set(constants.ParameterNameAdPlatformNoAdNetwork, "0")
	}
	return r
}

// LinkURL - method to set the required parameter link_url (url of the advertised objects)
func (r *AdsGetTargetingStatsRequest) LinkURL(url string) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameLinkURL, url)
	return r
}

// LinkDomain - method to set the required parameter link_domain (domain of the advertised objects)
func (r *AdsGetTargetingStatsRequest) LinkDomain(domain string) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameLinkDomain, domain)
	return r
}

// NeedPrecise - method to set the optional checkbox need_precise.
// If installed, the response will additionally contain recommended CPC and CPM values to achieve 5, 10, 15... 95% coverage.
func (r *AdsGetTargetingStatsRequest) NeedPrecise(value int) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameNeedPrecise, strconv.Itoa(value))
	return r
}

// ImpressionsLimitPeriod - method to set the optional parameter impressions_limit_period
func (r *AdsGetTargetingStatsRequest) ImpressionsLimitPeriod(limit int) *AdsGetTargetingStatsRequest {
	r.parameters.Set(constants.ParameterNameImpressionsLimitPeriod, strconv.Itoa(limit))
	return r
}

// AdsGetUploadURLRequest defines the request for ads.getUploadURL
// Returns the url to download the ad photo.
// For details on uploading images for ads, see https://dev.vk.com/ru/method/ads/upload-photo-ads
// Doc: https://dev.vk.com/method/ads.getUploadURL
type AdsGetUploadURLRequest struct {
	BaseRequest
}

// NewAdsGetUploadURLRequest creates a new request for ads.getUploadURL
func NewAdsGetUploadURLRequest(a *api.API, actor actor.Actor) *AdsGetUploadURLRequest {
	return &AdsGetUploadURLRequest{*NewMethodBaseRequest(a, actor, "ads.getUploadURL")}
}

// Exec executes the request and unmarshals the response into AdsGetUploadURLResponse
func (r *AdsGetUploadURLRequest) Exec(ctx context.Context) (response response.AdsGetUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdFormat - method to set the required parameter ad_format.
// 1 - image and text
// 2 - large image
// 3 - exclusive format
// 4 - promotion of communities or applications, square image
// 5 - application in the news feed
// 6 - mobile application
// 11 - adaptive format.
// For the “community post” ad format, this method is not used, because the photos are part of a community wall post. See the ads.createAds and wall.postAdsStealth methods.
func (r *AdsGetUploadURLRequest) AdFormat(format constants.AdsUploadURLAdFormat) *AdsGetUploadURLRequest {
	r.parameters.Set(constants.ParameterNameAdFormat, strconv.Itoa(int(format)))
	return r
}

// Icon - method to set the required parameter icon.
// 1 - Get the url to download the logo, not the main image. Only used for ad_format = 11.
func (r *AdsGetUploadURLRequest) Icon(value bool) *AdsGetUploadURLRequest {
	if value {
		r.parameters.Set(constants.ParameterNameIcon, "1")
	} else {
		r.parameters.Remove(constants.ParameterNameIcon)
	}
	return r
}

// AdsGetVideoUploadURLRequest defines the request for ads.getVideoUploadURL
// Returns the url to download the ad video.
// For details on uploading videos for ads, see https://dev.vk.com/ru/method/ads/upload-video-ads
// Doc: https://dev.vk.com/method/ads.getVideoUploadURL
type AdsGetVideoUploadURLRequest struct {
	BaseRequest
}

// NewAdsGetVideoUploadURLRequest creates a new request for ads.getVideoUploadURL
func NewAdsGetVideoUploadURLRequest(a *api.API, actor actor.Actor) *AdsGetVideoUploadURLRequest {
	return &AdsGetVideoUploadURLRequest{*NewMethodBaseRequest(a, actor, "ads.getVideoUploadURL")}
}

// Exec executes the request and unmarshals the response into AdsGetVideoUploadURLResponse
func (r *AdsGetVideoUploadURLRequest) Exec(ctx context.Context) (response response.AdsGetVideoUploadURLResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AdsImportTargetContactsRequest defines the request for ads.importTargetContacts
// Imports a list of advertiser contacts to count users registered on VKontakte in the retargeting audience.
// Doc: https://dev.vk.com/method/ads.importTargetContacts
type AdsImportTargetContactsRequest struct {
	BaseRequest
}

// NewAdsImportTargetContactsRequest creates a new request for ads.importTargetContacts
func NewAdsImportTargetContactsRequest(a *api.API, actor actor.Actor) *AdsImportTargetContactsRequest {
	return &AdsImportTargetContactsRequest{*NewMethodBaseRequest(a, actor, "ads.importTargetContacts")}
}

// Exec executes the request and unmarshals the response into AdsImportTargetContactsResponse
func (r *AdsImportTargetContactsRequest) Exec(ctx context.Context) (response response.AdsImportTargetContactsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsImportTargetContactsRequest) AccountID(id int) *AdsImportTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
func (r *AdsImportTargetContactsRequest) ClientID(id int) *AdsImportTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the required parameter target_group_id. Targeting audience ID.
func (r *AdsImportTargetContactsRequest) TargetGroupID(id int) *AdsImportTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// Contacts - method to set the required parameter contacts.
// List of phone numbers, email addresses, mobile advertising identifiers (IDFA, GAID) or user IDs, separated by commas. Their MD5 hashes or SHA256 hashes are also accepted.
func (r *AdsImportTargetContactsRequest) Contacts(contacts string) *AdsImportTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameContacts, contacts)
	return r
}

// AdsRemoveOfficeUsersRequest defines the request for ads.removeOfficeUsers
// Removes administrators and/or observers from the advertising account.
// Doc: https://dev.vk.com/method/ads.removeOfficeUsers
type AdsRemoveOfficeUsersRequest struct {
	BaseRequest
}

// NewAdsRemoveOfficeUsersRequest creates a new request for ads.removeOfficeUsers
func NewAdsRemoveOfficeUsersRequest(a *api.API, actor actor.Actor) *AdsRemoveOfficeUsersRequest {
	return &AdsRemoveOfficeUsersRequest{*NewMethodBaseRequest(a, actor, "ads.removeOfficeUsers")}
}

// Exec executes the request and unmarshals the response into AdsRemoveOfficeUsersResponse
func (r *AdsRemoveOfficeUsersRequest) Exec(ctx context.Context) (response response.AdsRemoveOfficeUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsRemoveOfficeUsersRequest) AccountID(id int) *AdsRemoveOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// IDs - method to set the required parameter ids. A serialized JSON array containing the ids of the administrators to be deleted.
func (r *AdsRemoveOfficeUsersRequest) IDs(ids []int) *AdsRemoveOfficeUsersRequest {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}
	r.parameters.Set(constants.ParameterNameIDs, strings.Join(strIDs, ","))
	return r
}

// IDsString - method to set the required parameter ids. A serialized JSON array containing the ids of the administrators to be deleted.
func (r *AdsRemoveOfficeUsersRequest) IDsString(ids string) *AdsRemoveOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameIDs, ids)
	return r
}

// AdsRemoveTargetContactsRequest defines the request for ads.removeTargetContacts
// Accepts a request to exclude advertiser contacts from the retargeting audience.
// Doc: https://dev.vk.com/method/ads.removeTargetContacts
type AdsRemoveTargetContactsRequest struct {
	BaseRequest
}

// NewAdsRemoveTargetContactsRequest creates a new request for ads.removeTargetContacts
func NewAdsRemoveTargetContactsRequest(a *api.API, actor actor.Actor) *AdsRemoveTargetContactsRequest {
	return &AdsRemoveTargetContactsRequest{*NewMethodBaseRequest(a, actor, "ads.removeTargetContacts")}
}

// Exec executes the request and unmarshals the response into AdsRemoveTargetContactsResponse
func (r *AdsRemoveTargetContactsRequest) Exec(ctx context.Context) (response response.AdsRemoveTargetContactsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsRemoveTargetContactsRequest) AccountID(id int) *AdsRemoveTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
// Client ID in the advertising account in which the audience is located.
func (r *AdsRemoveTargetContactsRequest) ClientID(id int) *AdsRemoveTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the required parameter target_group_id. Targeting audience ID.
func (r *AdsRemoveTargetContactsRequest) TargetGroupID(id int) *AdsRemoveTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// Contacts - method to set the required parameter contacts.
// List of phone numbers, email addresses, mobile advertising identifiers (IDFA, GAID) or user IDs, separated by commas. Their MD5 hashes or SHA256 hashes are also accepted.
func (r *AdsRemoveTargetContactsRequest) Contacts(contacts string) *AdsRemoveTargetContactsRequest {
	r.parameters.Set(constants.ParameterNameContacts, contacts)
	return r
}

// AdsSaveLookalikeRequestResultRequest defines the request for ads.saveLookalikeRequestResult
// Saves the result of a similar audience search.
// Doc: https://dev.vk.com/method/ads.saveLookalikeRequestResult
type AdsSaveLookalikeRequestResultRequest struct {
	BaseRequest
}

// NewAdsSaveLookalikeRequestResultRequest creates a new request for ads.saveLookalikeRequestResult
func NewAdsSaveLookalikeRequestResultRequest(a *api.API, actor actor.Actor) *AdsSaveLookalikeRequestResultRequest {
	return &AdsSaveLookalikeRequestResultRequest{*NewMethodBaseRequest(a, actor, "ads.saveLookalikeRequestResult")}
}

// Exec executes the request and unmarshals the response into AdsSaveLookalikeRequestResultResponse
func (r *AdsSaveLookalikeRequestResultRequest) Exec(ctx context.Context) (response response.AdsSaveLookalikeRequestResultResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsSaveLookalikeRequestResultRequest) AccountID(id int) *AdsSaveLookalikeRequestResultRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
// Client ID in the advertising account in which the audience is located.
func (r *AdsSaveLookalikeRequestResultRequest) ClientID(id int) *AdsSaveLookalikeRequestResultRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// RequestID - method to set the required parameter client_id (for agency accounts)
// Lookalike audience search request ID. You can get a list of all requests to search for a similar audience for a given account using ads.getLookalikeRequests.
func (r *AdsSaveLookalikeRequestResultRequest) RequestID(id int) *AdsSaveLookalikeRequestResultRequest {
	r.parameters.Set(constants.ParameterNameRequestID, strconv.Itoa(id))
	return r
}

// Level - method to set the required parameter level (for agency accounts)
// Level of specific lookalike audience size to retain. You can get a list of all available audience sizes using ads.getLookalikeRequests.
func (r *AdsSaveLookalikeRequestResultRequest) Level(id int) *AdsSaveLookalikeRequestResultRequest {
	r.parameters.Set(constants.ParameterNameLevel, strconv.Itoa(id))
	return r
}

// AdsShareTargetGroupRequest defines the request for ads.shareTargetGroup
// Shares a target group with other advertising accounts.
// Doc: https://dev.vk.com/method/ads.shareTargetGroup
type AdsShareTargetGroupRequest struct {
	BaseRequest
}

// NewAdsShareTargetGroupRequest creates a new request for ads.shareTargetGroup
func NewAdsShareTargetGroupRequest(a *api.API, actor actor.Actor) *AdsShareTargetGroupRequest {
	return &AdsShareTargetGroupRequest{*NewMethodBaseRequest(a, actor, "ads.shareTargetGroup")}
}

// Exec executes the request and unmarshals the response into AdsShareTargetGroupResponse
func (r *AdsShareTargetGroupRequest) Exec(ctx context.Context) (response response.AdsShareTargetGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsShareTargetGroupRequest) AccountID(id int) *AdsShareTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
// Client ID in the advertising account in which the audience is located.
func (r *AdsShareTargetGroupRequest) ClientID(id int) *AdsShareTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the required parameter target_group_id. Source audience ID.
func (r *AdsShareTargetGroupRequest) TargetGroupID(id int) *AdsShareTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// ShareWithClientID - method to set the parameter share_with_client_id.
// ID of the client whose advertising account needs to be granted access to the audience.
func (r *AdsShareTargetGroupRequest) ShareWithClientID(id int) *AdsShareTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameShareWithClientID, strconv.Itoa(id))
	return r
}

// AdsUpdateAdsRequest defines the request for ads.updateAds
// Edits advertisements
// Doc: https://dev.vk.com/method/ads.updateAds
type AdsUpdateAdsRequest struct {
	BaseRequest
}

// NewAdsUpdateAdsRequest creates a new request for ads.updateAds
func NewAdsUpdateAdsRequest(a *api.API, actor actor.Actor) *AdsUpdateAdsRequest {
	return &AdsUpdateAdsRequest{*NewMethodBaseRequest(a, actor, "ads.updateAds")}
}

// Exec executes the request and unmarshals the response into AdsUpdateAdsResponse
func (r *AdsUpdateAdsRequest) Exec(ctx context.Context) (response response.AdsUpdateAdsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateAdsRequest) AccountID(id int) *AdsUpdateAdsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
// A serialized JSON array of objects describing changes to declarations.
// For a description of ad_edit_specification objects, see below.
// The maximum allowed number of advertisements edited using one request is 5. If you need more, you can use execute.
func (r *AdsUpdateAdsRequest) Data(us *objects.AdsAdEditSpecifications) *AdsUpdateAdsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
// A serialized JSON array of objects describing changes to declarations.
// For a description of ad_edit_specification objects, see below.
// The maximum allowed number of advertisements edited using one request is 5. If you need more, you can use execute.
func (r *AdsUpdateAdsRequest) DataString(us string) *AdsUpdateAdsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsUpdateCampaignsRequest defines the request for ads.updateCampaigns
// Edits advertising campaigns.
// Doc: https://dev.vk.com/method/ads.updateCampaigns
type AdsUpdateCampaignsRequest struct {
	BaseRequest
}

// NewAdsUpdateCampaignsRequest creates a new request for ads.updateCampaigns
func NewAdsUpdateCampaignsRequest(a *api.API, actor actor.Actor) *AdsUpdateCampaignsRequest {
	return &AdsUpdateCampaignsRequest{*NewMethodBaseRequest(a, actor, "ads.updateCampaigns")}
}

// Exec executes the request and unmarshals the response into AdsUpdateCampaignsResponse
func (r *AdsUpdateCampaignsRequest) Exec(ctx context.Context) (response response.AdsUpdateCampaignsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateCampaignsRequest) AccountID(id int) *AdsUpdateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
// Serialized JSON array of objects describing changes in campaigns. For a description of campaign_mod_specification objects
func (r *AdsUpdateCampaignsRequest) Data(us *objects.AdsCampaignModSpecifications) *AdsUpdateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
// Serialized JSON array of objects describing changes in campaigns. For a description of campaign_mod_specification objects
func (r *AdsUpdateCampaignsRequest) DataString(us string) *AdsUpdateCampaignsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsUpdateClientsRequest defines the request for ads.updateClients
// The method edits advertising agency clients. Available only to advertising agencies.
// Doc: https://dev.vk.com/method/ads.updateClients
type AdsUpdateClientsRequest struct {
	BaseRequest
}

// NewAdsUpdateClientsRequest creates a new request for ads.updateClients
func NewAdsUpdateClientsRequest(a *api.API, actor actor.Actor) *AdsUpdateClientsRequest {
	return &AdsUpdateClientsRequest{*NewMethodBaseRequest(a, actor, "ads.updateClients")}
}

// Exec executes the request and unmarshals the response into AdsUpdateClientsResponse
func (r *AdsUpdateClientsRequest) Exec(ctx context.Context) (response response.AdsUpdateClientsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateClientsRequest) AccountID(id int) *AdsUpdateClientsRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
// Required parameter. Serialized JSON array of client_specification objects.
func (r *AdsUpdateClientsRequest) Data(us *objects.AdsClientSpecifications) *AdsUpdateClientsRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
// Required parameter. Serialized JSON array of client_specification objects.
func (r *AdsUpdateClientsRequest) DataString(us string) *AdsUpdateClientsRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsUpdateOfficeUsersRequest defines the request for ads.updateOfficeUsers
// Adds or edits administrators and/or observers to the advertising account.
// Doc: https://dev.vk.com/method/ads.updateOfficeUsers
type AdsUpdateOfficeUsersRequest struct {
	BaseRequest
}

// NewAdsUpdateOfficeUsersRequest creates a new request for ads.updateOfficeUsers
func NewAdsUpdateOfficeUsersRequest(a *api.API, actor actor.Actor) *AdsUpdateOfficeUsersRequest {
	return &AdsUpdateOfficeUsersRequest{*NewMethodBaseRequest(a, actor, "ads.updateOfficeUsers")}
}

// Exec executes the request and unmarshals the response into AdsUpdateOfficeUsersResponse
func (r *AdsUpdateOfficeUsersRequest) Exec(ctx context.Context) (response response.AdsUpdateOfficeUsersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateOfficeUsersRequest) AccountID(id int) *AdsUpdateOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// Data Required parameter. A serialized JSON array of objects describing the declarations to be created.
// A serialized JSON array of objects describing composite administrators. For a description of the user_specification objects, see below.
func (r *AdsUpdateOfficeUsersRequest) Data(us *objects.AdsUserSpecifications) *AdsUpdateOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameData, us.ToJSON())
	return r
}

// DataString Required parameter. A serialized JSON array of objects describing the declarations to be created.
// A serialized JSON array of objects describing composite administrators. For a description of the user_specification objects, see below.
func (r *AdsUpdateOfficeUsersRequest) DataString(us string) *AdsUpdateOfficeUsersRequest {
	r.parameters.Set(constants.ParameterNameData, us)
	return r
}

// AdsUpdateTargetGroupRequest defines the request for ads.updateTargetGroup
// Edits the retargeting audience.
// Doc: https://dev.vk.com/method/ads.updateTargetGroup
type AdsUpdateTargetGroupRequest struct {
	BaseRequest
}

// NewAdsUpdateTargetGroupRequest creates a new request for ads.updateTargetGroup
func NewAdsUpdateTargetGroupRequest(a *api.API, actor actor.Actor) *AdsUpdateTargetGroupRequest {
	return &AdsUpdateTargetGroupRequest{*NewMethodBaseRequest(a, actor, "ads.updateTargetGroup")}
}

// Exec executes the request and unmarshals the response into AdsUpdateTargetGroupResponse
func (r *AdsUpdateTargetGroupRequest) Exec(ctx context.Context) (response response.AdsUpdateTargetGroupResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateTargetGroupRequest) AccountID(id int) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
// Client ID in the advertising account in which the audience is located.
func (r *AdsUpdateTargetGroupRequest) ClientID(id int) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// TargetGroupID - method to set the required parameter target_group_id. Source audience ID.
func (r *AdsUpdateTargetGroupRequest) TargetGroupID(id int) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameTargetGroupID, strconv.Itoa(id))
	return r
}

// Name - method to set the required parameter name. The new retargeting audience name is a string of up to 64 characters.
func (r *AdsUpdateTargetGroupRequest) Name(name string) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameName, name)
	return r
}

// Domain - method to set the parameter domain.
// The domain of the site on which the user registration code will be placed.
// Deprecated option. Used only for old retargeting groups that were replenished without the help of a pixel. For new audiences, it should be omitted, otherwise an error will be returned.
func (r *AdsUpdateTargetGroupRequest) Domain(domain string) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameDomain, domain)
	return r
}

// LifeTime - method to set the required parameter life_time.
// The number of days after which users added to the audience will be automatically excluded from it.
func (r *AdsUpdateTargetGroupRequest) LifeTime(time int) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameLifeTime, strconv.Itoa(time))
	return r
}

// TargetPixelID - method to set the parameter target_pixel_id.
// Pass the pixel ID in this parameter if you want to collect audiences from a website.
func (r *AdsUpdateTargetGroupRequest) TargetPixelID(id int) *AdsUpdateTargetGroupRequest {
	r.parameters.Set(constants.ParameterNameTargetPixelID, strconv.Itoa(id))
	return r
}

// TargetPixelRules - method to set the parameter target_pixel_rules.
// Pass the pixel ID in this parameter if you want to collect audiences from a website.
// [
//
//	{"type": args},
//	{"type": args},
//	...
//	{"type": args}
//
// ](
//
//	{"type": args},
//	{"type": args},
//	...
//	{"type": args}
//
// )
func (r *AdsUpdateTargetGroupRequest) TargetPixelRules(rules []map[string]interface{}) *AdsUpdateTargetGroupRequest {
	rulesJSON, _ := json.Marshal(rules)
	r.parameters.Set(constants.ParameterNameTargetPixelRules, string(rulesJSON))
	return r
}

// AdsUpdateTargetPixelRequest defines the request for ads.updateTargetPixel
// Edits a retargeting pixel.
// Doc: https://dev.vk.com/method/ads.updateTargetPixel
type AdsUpdateTargetPixelRequest struct {
	BaseRequest
}

// NewAdsUpdateTargetPixelRequest creates a new request for ads.updateTargetPixel
func NewAdsUpdateTargetPixelRequest(a *api.API, actor actor.Actor) *AdsUpdateTargetPixelRequest {
	return &AdsUpdateTargetPixelRequest{*NewMethodBaseRequest(a, actor, "ads.updateTargetPixel")}
}

// Exec executes the request and unmarshals the response into AdsUpdateTargetPixelResponse
func (r *AdsUpdateTargetPixelRequest) Exec(ctx context.Context) (response response.AdsUpdateTargetPixelResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// AccountID - method to set the required parameter account_id (advertising account ID)
func (r *AdsUpdateTargetPixelRequest) AccountID(id int) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameAccountID, strconv.Itoa(id))
	return r
}

// ClientID - method to set the required parameter client_id (for agency accounts)
// Client ID in the advertising account in which the audience is located.
func (r *AdsUpdateTargetPixelRequest) ClientID(id int) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameClientID, strconv.Itoa(id))
	return r
}

// Name - method to set the required parameter name. The new retargeting audience name is a string of up to 64 characters.
func (r *AdsUpdateTargetPixelRequest) Name(name string) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameName, name)
	return r
}

// Domain - method to set the parameter domain.
// The new site domain on which the pixel code will be placed.
func (r *AdsUpdateTargetPixelRequest) Domain(domain string) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameDomain, domain)
	return r
}

// TargetPixelID - method to set the parameter target_pixel_id.
// Pass the pixel ID in this parameter if you want to collect audiences from a website.
func (r *AdsUpdateTargetPixelRequest) TargetPixelID(id int) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameTargetPixelID, strconv.Itoa(id))
	return r
}

// CategoryID - method to set the required parameter category_id.
// The new category identifier of the site on which the pixel will be placed. To get a list of possible IDs,
// you can use the ads.getSuggestions - https://dev.vk.com/ru/method/ads.getSuggestions method (interest_categories section).
func (r *AdsUpdateTargetPixelRequest) CategoryID(id int) *AdsUpdateTargetPixelRequest {
	r.parameters.Set(constants.ParameterNameCategoryID, strconv.Itoa(id))
	return r
}
