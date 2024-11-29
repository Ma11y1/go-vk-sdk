package request

import (
	"context"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/market
//
// Doc2: https://dev.vk.com/ru/api/market/getting-started

// MarketAddRequest defines the request for market.add
//
// Adds a new product.
// Doc: https://dev.vk.com/method/market.add
type MarketAddRequest struct {
	BaseRequest
}

// NewMarketAddRequest creates a new request for market.add
func NewMarketAddRequest(a *api.API, actor actor.Actor) *MarketAddRequest {
	return &MarketAddRequest{*NewMethodBaseRequest(a, actor, "market.add")}
}

// Exec executes the request and unmarshals the response into MarketAddResponse
func (r *MarketAddRequest) Exec(ctx context.Context) (response response.MarketAddResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketAddAlbumRequest defines the request for market.addAlbum
//
// Adds a new album with products to the community.
// Doc: https://dev.vk.com/method/market.addAlbum
type MarketAddAlbumRequest struct {
	BaseRequest
}

// NewMarketAddAlbumRequest creates a new request for market.addAlbum
func NewMarketAddAlbumRequest(a *api.API, actor actor.Actor) *MarketAddAlbumRequest {
	return &MarketAddAlbumRequest{*NewMethodBaseRequest(a, actor, "market.addAlbum")}
}

// Exec executes the request and unmarshals the response into MarketAddAlbumResponse
func (r *MarketAddAlbumRequest) Exec(ctx context.Context) (response response.MarketAddAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketAddPropertyRequest defines the request for market.addProperty
//
// Adds a new property (e.g., "color", "size") for community products.
// Doc: https://dev.vk.com/method/market.addProperty
type MarketAddPropertyRequest struct {
	BaseRequest
}

// NewMarketAddPropertyRequest creates a new request for market.addProperty
func NewMarketAddPropertyRequest(a *api.API, actor actor.Actor) *MarketAddPropertyRequest {
	return &MarketAddPropertyRequest{*NewMethodBaseRequest(a, actor, "market.addProperty")}
}

// Exec executes the request and unmarshals the response into MarketAddPropertyResponse
func (r *MarketAddPropertyRequest) Exec(ctx context.Context) (response response.MarketAddPropertyResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketAddPropertyVariantRequest defines the request for market.addPropertyVariant
//
// Adds a variant for a property. Each property can have up to 50 variants.
// Doc: https://dev.vk.com/method/market.addPropertyVariant
type MarketAddPropertyVariantRequest struct {
	BaseRequest
}

// NewMarketAddPropertyVariantRequest creates a new request for market.addPropertyVariant
func NewMarketAddPropertyVariantRequest(a *api.API, actor actor.Actor) *MarketAddPropertyVariantRequest {
	return &MarketAddPropertyVariantRequest{*NewMethodBaseRequest(a, actor, "market.addPropertyVariant")}
}

// Exec executes the request and unmarshals the response into MarketAddPropertyVariantResponse
func (r *MarketAddPropertyVariantRequest) Exec(ctx context.Context) (response response.MarketAddPropertyVariantResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketAddToAlbumRequest defines the request for market.addToAlbum
//
// Adds a product to one or more selected albums.
// Doc: https://dev.vk.com/method/market.addToAlbum
type MarketAddToAlbumRequest struct {
	BaseRequest
}

// NewMarketAddToAlbumRequest creates a new request for market.addToAlbum
func NewMarketAddToAlbumRequest(a *api.API, actor actor.Actor) *MarketAddToAlbumRequest {
	return &MarketAddToAlbumRequest{*NewMethodBaseRequest(a, actor, "market.addToAlbum")}
}

// Exec executes the request and unmarshals the response into MarketAddToAlbumResponse
func (r *MarketAddToAlbumRequest) Exec(ctx context.Context) (response response.MarketAddToAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketCreateCommentRequest defines the request for market.createComment
//
// Creates a new comment for a product.
// Doc: https://dev.vk.com/method/market.createComment
type MarketCreateCommentRequest struct {
	BaseRequest
}

// NewMarketCreateCommentRequest creates a new request for market.createComment
func NewMarketCreateCommentRequest(a *api.API, actor actor.Actor) *MarketCreateCommentRequest {
	return &MarketCreateCommentRequest{*NewMethodBaseRequest(a, actor, "market.createComment")}
}

// Exec executes the request and unmarshals the response into MarketCreateCommentResponse
func (r *MarketCreateCommentRequest) Exec(ctx context.Context) (response response.MarketCreateCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketDeleteRequest defines the request for market.delete
//
// Deletes a product.
// Doc: https://dev.vk.com/method/market.delete
type MarketDeleteRequest struct {
	BaseRequest
}

// NewMarketDeleteRequest creates a new request for market.delete
func NewMarketDeleteRequest(a *api.API, actor actor.Actor) *MarketDeleteRequest {
	return &MarketDeleteRequest{*NewMethodBaseRequest(a, actor, "market.delete")}
}

// Exec executes the request and unmarshals the response into MarketDeleteResponse
func (r *MarketDeleteRequest) Exec(ctx context.Context) (response response.MarketDeleteResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketDeleteAlbumRequest defines the request for market.deleteAlbum
//
// Deletes an album with products.
// Doc: https://dev.vk.com/method/market.deleteAlbum
type MarketDeleteAlbumRequest struct {
	BaseRequest
}

// NewMarketDeleteAlbumRequest creates a new request for market.deleteAlbum
func NewMarketDeleteAlbumRequest(a *api.API, actor actor.Actor) *MarketDeleteAlbumRequest {
	return &MarketDeleteAlbumRequest{*NewMethodBaseRequest(a, actor, "market.deleteAlbum")}
}

// Exec executes the request and unmarshals the response into MarketDeleteAlbumResponse
func (r *MarketDeleteAlbumRequest) Exec(ctx context.Context) (response response.MarketDeleteAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketDeleteCommentRequest defines the request for market.deleteComment
//
// Deletes a comment for a product.
// Doc: https://dev.vk.com/method/market.deleteComment
type MarketDeleteCommentRequest struct {
	BaseRequest
}

// NewMarketDeleteCommentRequest creates a new request for market.deleteComment
func NewMarketDeleteCommentRequest(a *api.API, actor actor.Actor) *MarketDeleteCommentRequest {
	return &MarketDeleteCommentRequest{*NewMethodBaseRequest(a, actor, "market.deleteComment")}
}

// Exec executes the request and unmarshals the response into MarketDeleteCommentResponse
func (r *MarketDeleteCommentRequest) Exec(ctx context.Context) (response response.MarketDeleteCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketDeletePropertyRequest defines the request for market.deleteProperty
//
// Deletes a property for a product.
// Doc: https://dev.vk.com/method/market.deleteProperty
type MarketDeletePropertyRequest struct {
	BaseRequest
}

// NewMarketDeletePropertyRequest creates a new request for market.deleteProperty
func NewMarketDeletePropertyRequest(a *api.API, actor actor.Actor) *MarketDeletePropertyRequest {
	return &MarketDeletePropertyRequest{*NewMethodBaseRequest(a, actor, "market.deleteProperty")}
}

// Exec executes the request and unmarshals the response into MarketDeletePropertyResponse
func (r *MarketDeletePropertyRequest) Exec(ctx context.Context) (response response.MarketDeletePropertyResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketDeletePropertyVariantRequest defines the request for market.deletePropertyVariant
//
// Deletes a variant for a property of a product.
// Doc: https://dev.vk.com/method/market.deletePropertyVariant
type MarketDeletePropertyVariantRequest struct {
	BaseRequest
}

// NewMarketDeletePropertyVariantRequest creates a new request for market.deletePropertyVariant
func NewMarketDeletePropertyVariantRequest(a *api.API, actor actor.Actor) *MarketDeletePropertyVariantRequest {
	return &MarketDeletePropertyVariantRequest{*NewMethodBaseRequest(a, actor, "market.deletePropertyVariant")}
}

// Exec executes the request and unmarshals the response into MarketDeletePropertyVariantResponse
func (r *MarketDeletePropertyVariantRequest) Exec(ctx context.Context) (response response.MarketDeletePropertyVariantResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditRequest defines the request for market.edit
//
// Edits information about a product.
// Doc: https://dev.vk.com/method/market.edit
type MarketEditRequest struct {
	BaseRequest
}

// NewMarketEditRequest creates a new request for market.edit
func NewMarketEditRequest(a *api.API, actor actor.Actor) *MarketEditRequest {
	return &MarketEditRequest{*NewMethodBaseRequest(a, actor, "market.edit")}
}

// Exec executes the request and unmarshals the response into MarketEditResponse
func (r *MarketEditRequest) Exec(ctx context.Context) (response response.MarketEditResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditAlbumRequest defines the request for market.editAlbum
//
// Edits an album with products in the community.
// Doc: https://dev.vk.com/method/market.editAlbum
type MarketEditAlbumRequest struct {
	BaseRequest
}

// NewMarketEditAlbumRequest creates a new request for market.editAlbum
func NewMarketEditAlbumRequest(a *api.API, actor actor.Actor) *MarketEditAlbumRequest {
	return &MarketEditAlbumRequest{*NewMethodBaseRequest(a, actor, "market.editAlbum")}
}

// Exec executes the request and unmarshals the response into MarketEditAlbumResponse
func (r *MarketEditAlbumRequest) Exec(ctx context.Context) (response response.MarketEditAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditCommentRequest defines the request for market.editComment
//
// Edits the text of a comment on a product.
// Doc: https://dev.vk.com/method/market.editComment
type MarketEditCommentRequest struct {
	BaseRequest
}

// NewMarketEditCommentRequest creates a new request for market.editComment
func NewMarketEditCommentRequest(a *api.API, actor actor.Actor) *MarketEditCommentRequest {
	return &MarketEditCommentRequest{*NewMethodBaseRequest(a, actor, "market.editComment")}
}

// Exec executes the request and unmarshals the response into MarketEditCommentResponse
func (r *MarketEditCommentRequest) Exec(ctx context.Context) (response response.MarketEditCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditOrderRequest defines the request for market.editOrder
//
// Edits an order.
// Doc: https://dev.vk.com/method/market.editOrder
type MarketEditOrderRequest struct {
	BaseRequest
}

// NewMarketEditOrderRequest creates a new request for market.editOrder
func NewMarketEditOrderRequest(a *api.API, actor actor.Actor) *MarketEditOrderRequest {
	return &MarketEditOrderRequest{*NewMethodBaseRequest(a, actor, "market.editOrder")}
}

// Exec executes the request and unmarshals the response into MarketEditOrderResponse
func (r *MarketEditOrderRequest) Exec(ctx context.Context) (response response.MarketEditOrderResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditPropertyRequest defines the request for market.editProperty
//
// Edits a product property.
// Doc: https://dev.vk.com/method/market.editProperty
type MarketEditPropertyRequest struct {
	BaseRequest
}

// NewMarketEditPropertyRequest creates a new request for market.editProperty
func NewMarketEditPropertyRequest(a *api.API, actor actor.Actor) *MarketEditPropertyRequest {
	return &MarketEditPropertyRequest{*NewMethodBaseRequest(a, actor, "market.editProperty")}
}

// Exec executes the request and unmarshals the response into MarketEditPropertyResponse
func (r *MarketEditPropertyRequest) Exec(ctx context.Context) (response response.MarketEditPropertyResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketEditPropertyVariantRequest defines the request for market.editPropertyVariant
//
// Edits a variant of a product property.
// Doc: https://dev.vk.com/method/market.editPropertyVariant
type MarketEditPropertyVariantRequest struct {
	BaseRequest
}

// NewMarketEditPropertyVariantRequest creates a new request for market.editPropertyVariant
func NewMarketEditPropertyVariantRequest(a *api.API, actor actor.Actor) *MarketEditPropertyVariantRequest {
	return &MarketEditPropertyVariantRequest{*NewMethodBaseRequest(a, actor, "market.editPropertyVariant")}
}

// Exec executes the request and unmarshals the response into MarketEditPropertyVariantResponse
func (r *MarketEditPropertyVariantRequest) Exec(ctx context.Context) (response response.MarketEditPropertyVariantResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketFilterCategoriesRequest defines the request for market.filterCategories
//
// Filters and retrieves product categories.
// Doc: https://dev.vk.com/method/market.filterCategories
type MarketFilterCategoriesRequest struct {
	BaseRequest
}

// NewMarketFilterCategoriesRequest creates a new request for market.filterCategories
func NewMarketFilterCategoriesRequest(a *api.API, actor actor.Actor) *MarketFilterCategoriesRequest {
	return &MarketFilterCategoriesRequest{*NewMethodBaseRequest(a, actor, "market.filterCategories")}
}

// Exec executes the request and unmarshals the response into MarketFilterCategoriesResponse
func (r *MarketFilterCategoriesRequest) Exec(ctx context.Context) (response response.MarketFilterCategoriesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetRequest defines the request for market.get
//
// Retrieves a list of community products.
// Doc: https://dev.vk.com/method/market.get
type MarketGetRequest struct {
	BaseRequest
}

// NewMarketGetRequest creates a new request for market.get
func NewMarketGetRequest(a *api.API, actor actor.Actor) *MarketGetRequest {
	return &MarketGetRequest{*NewMethodBaseRequest(a, actor, "market.get")}
}

// Exec executes the request and unmarshals the response into MarketGetResponse
func (r *MarketGetRequest) Exec(ctx context.Context) (response response.MarketGetResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetAlbumByIDRequest defines the request for market.getAlbumById
//
// Retrieves data about a product album.
// Doc: https://dev.vk.com/method/market.getAlbumById
type MarketGetAlbumByIDRequest struct {
	BaseRequest
}

// NewMarketGetAlbumByIDRequest creates a new request for market.getAlbumById
func NewMarketGetAlbumByIDRequest(a *api.API, actor actor.Actor) *MarketGetAlbumByIDRequest {
	return &MarketGetAlbumByIDRequest{*NewMethodBaseRequest(a, actor, "market.getAlbumById")}
}

// Exec executes the request and unmarshals the response into MarketGetAlbumByIDResponse
func (r *MarketGetAlbumByIDRequest) Exec(ctx context.Context) (response response.MarketGetAlbumByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetAlbumsRequest defines the request for market.getAlbums
//
// Retrieves a list of product albums.
// Doc: https://dev.vk.com/method/market.getAlbums
type MarketGetAlbumsRequest struct {
	BaseRequest
}

// NewMarketGetAlbumsRequest creates a new request for market.getAlbums
func NewMarketGetAlbumsRequest(a *api.API, actor actor.Actor) *MarketGetAlbumsRequest {
	return &MarketGetAlbumsRequest{*NewMethodBaseRequest(a, actor, "market.getAlbums")}
}

// Exec executes the request and unmarshals the response into MarketGetAlbumsResponse
func (r *MarketGetAlbumsRequest) Exec(ctx context.Context) (response response.MarketGetAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetByIDRequest defines the request for market.getById
//
// Retrieves information about products by IDs.
// Doc: https://dev.vk.com/method/market.getById
type MarketGetByIDRequest struct {
	BaseRequest
}

// NewMarketGetByIDRequest creates a new request for market.getById
func NewMarketGetByIDRequest(a *api.API, actor actor.Actor) *MarketGetByIDRequest {
	return &MarketGetByIDRequest{*NewMethodBaseRequest(a, actor, "market.getById")}
}

// Exec executes the request and unmarshals the response into MarketGetByIDResponse
func (r *MarketGetByIDRequest) Exec(ctx context.Context) (response response.MarketGetByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetCategoriesRequest defines the request for market.getCategories
//
// Retrieves a list of product categories.
// Doc: https://dev.vk.com/method/market.getCategories
type MarketGetCategoriesRequest struct {
	BaseRequest
}

// NewMarketGetCategoriesRequest creates a new request for market.getCategories
func NewMarketGetCategoriesRequest(a *api.API, actor actor.Actor) *MarketGetCategoriesRequest {
	return &MarketGetCategoriesRequest{*NewMethodBaseRequest(a, actor, "market.getCategories")}
}

// Exec executes the request and unmarshals the response into MarketGetCategoriesResponse
func (r *MarketGetCategoriesRequest) Exec(ctx context.Context) (response response.MarketGetCategoriesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetCommentsRequest defines the request for market.getComments
//
// Retrieves a list of comments on a product.
// Doc: https://dev.vk.com/method/market.getComments
type MarketGetCommentsRequest struct {
	BaseRequest
}

// NewMarketGetCommentsRequest creates a new request for market.getComments
func NewMarketGetCommentsRequest(a *api.API, actor actor.Actor) *MarketGetCommentsRequest {
	return &MarketGetCommentsRequest{*NewMethodBaseRequest(a, actor, "market.getComments")}
}

// Exec executes the request and unmarshals the response into MarketGetCommentsResponse
func (r *MarketGetCommentsRequest) Exec(ctx context.Context) (response response.MarketGetCommentsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetCommentsExtendedRequest defines the request for market.getComments
//
// Retrieves a list of comments on a product.
// Doc: https://dev.vk.com/method/market.getComments
type MarketGetCommentsExtendedRequest struct {
	BaseRequest
}

// NewMarketGetCommentsExtendedRequest creates a new request for market.getComments
func NewMarketGetCommentsExtendedRequest(a *api.API, actor actor.Actor) *MarketGetCommentsExtendedRequest {
	r := &MarketGetCommentsExtendedRequest{*NewMethodBaseRequest(a, actor, "market.getComments")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into MarketGetCommentsExtendedResponse
func (r *MarketGetCommentsExtendedRequest) Exec(ctx context.Context) (response response.MarketGetCommentsExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetGroupOrdersRequest defines the request for market.getGroupOrders
//
// Retrieves a list of group orders.
// Doc: https://dev.vk.com/method/market.getGroupOrders
type MarketGetGroupOrdersRequest struct {
	BaseRequest
}

// NewMarketGetGroupOrdersRequest creates a new request for market.getGroupOrders
func NewMarketGetGroupOrdersRequest(a *api.API, actor actor.Actor) *MarketGetGroupOrdersRequest {
	return &MarketGetGroupOrdersRequest{*NewMethodBaseRequest(a, actor, "market.getGroupOrders")}
}

// Exec executes the request and unmarshals the response into MarketGetGroupOrdersResponse
func (r *MarketGetGroupOrdersRequest) Exec(ctx context.Context) (response response.MarketGetGroupOrdersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetOrderByIDRequest defines the request for market.getOrderById
//
// Retrieves an order by ID.
// Doc: https://dev.vk.com/method/market.getOrderById
type MarketGetOrderByIDRequest struct {
	BaseRequest
}

// NewMarketGetOrderByIDRequest creates a new request for market.getOrderById
func NewMarketGetOrderByIDRequest(a *api.API, actor actor.Actor) *MarketGetOrderByIDRequest {
	return &MarketGetOrderByIDRequest{*NewMethodBaseRequest(a, actor, "market.getOrderById")}
}

// Exec executes the request and unmarshals the response into MarketGetOrderByIDResponse
func (r *MarketGetOrderByIDRequest) Exec(ctx context.Context) (response response.MarketGetOrderByIDResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetOrderItemsRequest defines the request for market.getOrderItems
//
// Retrieves items within an order.
// Doc: https://dev.vk.com/method/market.getOrderItems
type MarketGetOrderItemsRequest struct {
	BaseRequest
}

// NewMarketGetOrderItemsRequest creates a new request for market.getOrderItems
func NewMarketGetOrderItemsRequest(a *api.API, actor actor.Actor) *MarketGetOrderItemsRequest {
	return &MarketGetOrderItemsRequest{*NewMethodBaseRequest(a, actor, "market.getOrderItems")}
}

// Exec executes the request and unmarshals the response into MarketGetOrderItemsResponse
func (r *MarketGetOrderItemsRequest) Exec(ctx context.Context) (response response.MarketGetOrderItemsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetOrdersRequest defines the request for market.getOrders
//
// Retrieves a list of orders.
// Doc: https://dev.vk.com/method/market.getOrders
type MarketGetOrdersRequest struct {
	BaseRequest
}

// NewMarketGetOrdersRequest creates a new request for market.getOrders
func NewMarketGetOrdersRequest(a *api.API, actor actor.Actor) *MarketGetOrdersRequest {
	return &MarketGetOrdersRequest{*NewMethodBaseRequest(a, actor, "market.getOrders")}
}

// Exec executes the request and unmarshals the response into MarketGetOrdersResponse
func (r *MarketGetOrdersRequest) Exec(ctx context.Context) (response response.MarketGetOrdersResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetOrdersExtendedRequest defines the request for market.getOrders
//
// Retrieves a list of orders.
// Doc: https://dev.vk.com/method/market.getOrders
type MarketGetOrdersExtendedRequest struct {
	BaseRequest
}

// NewMarketGetOrdersExtendedRequest creates a new request for market.getOrders
func NewMarketGetOrdersExtendedRequest(a *api.API, actor actor.Actor) *MarketGetOrdersExtendedRequest {
	r := &MarketGetOrdersExtendedRequest{*NewMethodBaseRequest(a, actor, "market.getOrders")}
	r.parameters.Set(constants.ParameterNameExtended, "1")
	return r
}

// Exec executes the request and unmarshals the response into MarketGetOrdersExtendedResponse
func (r *MarketGetOrdersExtendedRequest) Exec(ctx context.Context) (response response.MarketGetOrdersExtendedResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetProductPhotoUploadServerRequest defines the request for market.getProductPhotoUploadServer
//
// Retrieves the URL for uploading product photos.
// Doc: https://dev.vk.com/method/market.getProductPhotoUploadServer
type MarketGetProductPhotoUploadServerRequest struct {
	BaseRequest
}

// NewMarketGetProductPhotoUploadServerRequest creates a new request for market.getProductPhotoUploadServer
func NewMarketGetProductPhotoUploadServerRequest(a *api.API, actor actor.Actor) *MarketGetProductPhotoUploadServerRequest {
	return &MarketGetProductPhotoUploadServerRequest{*NewMethodBaseRequest(a, actor, "market.getProductPhotoUploadServer")}
}

// Exec executes the request and unmarshals the response into MarketGetProductPhotoUploadServerResponse
func (r *MarketGetProductPhotoUploadServerRequest) Exec(ctx context.Context) (response response.MarketGetProductPhotoUploadServerResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGetPropertiesRequest defines the request for market.getProperties
//
// Retrieves a list of properties for the specified community.
// Doc: https://dev.vk.com/method/market.getProperties
type MarketGetPropertiesRequest struct {
	BaseRequest
}

// NewMarketGetPropertiesRequest creates a new request for market.getProperties
func NewMarketGetPropertiesRequest(a *api.API, actor actor.Actor) *MarketGetPropertiesRequest {
	return &MarketGetPropertiesRequest{*NewMethodBaseRequest(a, actor, "market.getProperties")}
}

// Exec executes the request and unmarshals the response into MarketGetPropertiesResponse
func (r *MarketGetPropertiesRequest) Exec(ctx context.Context) (response response.MarketGetPropertiesResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketGroupItemsRequest defines the request for market.groupItems
//
// Groups items into a single product group.
// Doc: https://dev.vk.com/method/market.groupItems
type MarketGroupItemsRequest struct {
	BaseRequest
}

// NewMarketGroupItemsRequest creates a new request for market.groupItems
func NewMarketGroupItemsRequest(a *api.API, actor actor.Actor) *MarketGroupItemsRequest {
	return &MarketGroupItemsRequest{*NewMethodBaseRequest(a, actor, "market.groupItems")}
}

// Exec executes the request and unmarshals the response into MarketGroupItemsResponse
func (r *MarketGroupItemsRequest) Exec(ctx context.Context) (response response.MarketGroupItemsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketRemoveFromAlbumRequest defines the request for market.removeFromAlbum
//
// Removes a product from one or more selected albums.
// Doc: https://dev.vk.com/method/market.removeFromAlbum
type MarketRemoveFromAlbumRequest struct {
	BaseRequest
}

// NewMarketRemoveFromAlbumRequest creates a new request for market.removeFromAlbum
func NewMarketRemoveFromAlbumRequest(a *api.API, actor actor.Actor) *MarketRemoveFromAlbumRequest {
	return &MarketRemoveFromAlbumRequest{*NewMethodBaseRequest(a, actor, "market.removeFromAlbum")}
}

// Exec executes the request and unmarshals the response into MarketRemoveFromAlbumResponse
func (r *MarketRemoveFromAlbumRequest) Exec(ctx context.Context) (response response.MarketRemoveFromAlbumResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketReorderAlbumsRequest defines the request for market.reorderAlbums
//
// Changes the position of an album in the list.
// Doc: https://dev.vk.com/method/market.reorderAlbums
type MarketReorderAlbumsRequest struct {
	BaseRequest
}

// NewMarketReorderAlbumsRequest creates a new request for market.reorderAlbums
func NewMarketReorderAlbumsRequest(a *api.API, actor actor.Actor) *MarketReorderAlbumsRequest {
	return &MarketReorderAlbumsRequest{*NewMethodBaseRequest(a, actor, "market.reorderAlbums")}
}

// Exec executes the request and unmarshals the response into MarketReorderAlbumsResponse
func (r *MarketReorderAlbumsRequest) Exec(ctx context.Context) (response response.MarketReorderAlbumsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketReorderItemsRequest defines the request for market.reorderItems
//
// Changes the position of an item in an album.
// Doc: https://dev.vk.com/method/market.reorderItems
type MarketReorderItemsRequest struct {
	BaseRequest
}

// NewMarketReorderItemsRequest creates a new request for market.reorderItems
func NewMarketReorderItemsRequest(a *api.API, actor actor.Actor) *MarketReorderItemsRequest {
	return &MarketReorderItemsRequest{*NewMethodBaseRequest(a, actor, "market.reorderItems")}
}

// Exec executes the request and unmarshals the response into MarketReorderItemsResponse
func (r *MarketReorderItemsRequest) Exec(ctx context.Context) (response response.MarketReorderItemsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketReportRequest defines the request for market.report
//
// Submits a complaint about a product.
// Doc: https://dev.vk.com/method/market.report
type MarketReportRequest struct {
	BaseRequest
}

// NewMarketReportRequest creates a new request for market.report
func NewMarketReportRequest(a *api.API, actor actor.Actor) *MarketReportRequest {
	return &MarketReportRequest{*NewMethodBaseRequest(a, actor, "market.report")}
}

// Exec executes the request and unmarshals the response into MarketReportResponse
func (r *MarketReportRequest) Exec(ctx context.Context) (response response.MarketReportResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketReportCommentRequest defines the request for market.reportComment
//
// Submits a complaint about a comment on a product.
// Doc: https://dev.vk.com/method/market.reportComment
type MarketReportCommentRequest struct {
	BaseRequest
}

// NewMarketReportCommentRequest creates a new request for market.reportComment
func NewMarketReportCommentRequest(a *api.API, actor actor.Actor) *MarketReportCommentRequest {
	return &MarketReportCommentRequest{*NewMethodBaseRequest(a, actor, "market.reportComment")}
}

// Exec executes the request and unmarshals the response into MarketReportCommentResponse
func (r *MarketReportCommentRequest) Exec(ctx context.Context) (response response.MarketReportCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketRestoreRequest defines the request for market.restore
//
// Restores a deleted product.
// Doc: https://dev.vk.com/method/market.restore
type MarketRestoreRequest struct {
	BaseRequest
}

// NewMarketRestoreRequest creates a new request for market.restore
func NewMarketRestoreRequest(a *api.API, actor actor.Actor) *MarketRestoreRequest {
	return &MarketRestoreRequest{*NewMethodBaseRequest(a, actor, "market.restore")}
}

// Exec executes the request and unmarshals the response into MarketRestoreResponse
func (r *MarketRestoreRequest) Exec(ctx context.Context) (response response.MarketRestoreResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketRestoreCommentRequest defines the request for market.restoreComment
//
// Restores a deleted comment on a product.
// Doc: https://dev.vk.com/method/market.restoreComment
type MarketRestoreCommentRequest struct {
	BaseRequest
}

// NewMarketRestoreCommentRequest creates a new request for market.restoreComment
func NewMarketRestoreCommentRequest(a *api.API, actor actor.Actor) *MarketRestoreCommentRequest {
	return &MarketRestoreCommentRequest{*NewMethodBaseRequest(a, actor, "market.restoreComment")}
}

// Exec executes the request and unmarshals the response into MarketRestoreCommentResponse
func (r *MarketRestoreCommentRequest) Exec(ctx context.Context) (response response.MarketRestoreCommentResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketSaveProductPhotoRequest defines the request for market.saveProductPhoto
//
// Prepares an uploaded photo for adding to a community product.
// Doc: https://dev.vk.com/method/market.saveProductPhoto
type MarketSaveProductPhotoRequest struct {
	BaseRequest
}

// NewMarketSaveProductPhotoRequest creates a new request for market.saveProductPhoto
func NewMarketSaveProductPhotoRequest(a *api.API, actor actor.Actor) *MarketSaveProductPhotoRequest {
	return &MarketSaveProductPhotoRequest{*NewMethodBaseRequest(a, actor, "market.saveProductPhoto")}
}

// Exec executes the request and unmarshals the response into MarketSaveProductPhotoResponse
func (r *MarketSaveProductPhotoRequest) Exec(ctx context.Context) (response response.MarketSaveProductPhotoResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketSearchRequest defines the request for market.search
//
// Retrieves items from the community catalog.
// Doc: https://dev.vk.com/method/market.search
type MarketSearchRequest struct {
	BaseRequest
}

// NewMarketSearchRequest creates a new request for market.search
func NewMarketSearchRequest(a *api.API, actor actor.Actor) *MarketSearchRequest {
	return &MarketSearchRequest{*NewMethodBaseRequest(a, actor, "market.search")}
}

// Exec executes the request and unmarshals the response into MarketSearchResponse
func (r *MarketSearchRequest) Exec(ctx context.Context) (response response.MarketSearchResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketSearchItemsRequest defines the request for market.searchItems
//
// Searches for items in the community catalog.
// Doc: https://dev.vk.com/method/market.searchItems
type MarketSearchItemsRequest struct {
	BaseRequest
}

// NewMarketSearchItemsRequest creates a new request for market.searchItems
func NewMarketSearchItemsRequest(a *api.API, actor actor.Actor) *MarketSearchItemsRequest {
	return &MarketSearchItemsRequest{*NewMethodBaseRequest(a, actor, "market.searchItems")}
}

// Exec executes the request and unmarshals the response into MarketSearchItemsResponse
func (r *MarketSearchItemsRequest) Exec(ctx context.Context) (response response.MarketSearchItemsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketSearchItemsBasicRequest defines the request for market.searchItemsBasic
//
// Performs a basic search for items in the community catalog.
// Doc: https://dev.vk.com/method/market.searchItemsBasic
type MarketSearchItemsBasicRequest struct {
	BaseRequest
}

// NewMarketSearchItemsBasicRequest creates a new request for market.searchItemsBasic
func NewMarketSearchItemsBasicRequest(a *api.API, actor actor.Actor) *MarketSearchItemsBasicRequest {
	return &MarketSearchItemsBasicRequest{*NewMethodBaseRequest(a, actor, "market.searchItemsBasic")}
}

// Exec executes the request and unmarshals the response into MarketSearchItemsBasicResponse
func (r *MarketSearchItemsBasicRequest) Exec(ctx context.Context) (response response.MarketSearchItemsBasicResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}

// MarketUngroupItemsRequest defines the request for market.ungroupItems
//
// Splits a product group into multiple individual items.
// Doc: https://dev.vk.com/method/market.ungroupItems
type MarketUngroupItemsRequest struct {
	BaseRequest
}

// NewMarketUngroupItemsRequest creates a new request for market.ungroupItems
func NewMarketUngroupItemsRequest(a *api.API, actor actor.Actor) *MarketUngroupItemsRequest {
	return &MarketUngroupItemsRequest{*NewMethodBaseRequest(a, actor, "market.ungroupItems")}
}

// Exec executes the request and unmarshals the response into MarketUngroupItemsResponse
func (r *MarketUngroupItemsRequest) Exec(ctx context.Context) (response response.MarketUngroupItemsResponse, err error) {
	err = r.PostUnmarshal(ctx, &response)
	return
}
