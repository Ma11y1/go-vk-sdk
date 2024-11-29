package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Market Doc: https://dev.vk.com/ru/method/market
//
// Doc2: https://dev.vk.com/ru/api/market/getting-started
type Market struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/ru/method/market.add
func (m *Market) Add(actor actor.Actor) *request.MarketAddRequest {
	return request.NewMarketAddRequest(m.api, actor)
}

// AddAlbum Doc: https://dev.vk.com/ru/method/market.addAlbum
func (m *Market) AddAlbum(actor actor.Actor) *request.MarketAddAlbumRequest {
	return request.NewMarketAddAlbumRequest(m.api, actor)
}

// AddProperty Doc: https://dev.vk.com/ru/method/market.addProperty
func (m *Market) AddProperty(actor actor.Actor) *request.MarketAddPropertyRequest {
	return request.NewMarketAddPropertyRequest(m.api, actor)
}

// AddPropertyVariant Doc: https://dev.vk.com/ru/method/market.addPropertyVariant
func (m *Market) AddPropertyVariant(actor actor.Actor) *request.MarketAddPropertyVariantRequest {
	return request.NewMarketAddPropertyVariantRequest(m.api, actor)
}

// AddToAlbum Doc: https://dev.vk.com/ru/method/market.addToAlbum
func (m *Market) AddToAlbum(actor actor.Actor) *request.MarketAddToAlbumRequest {
	return request.NewMarketAddToAlbumRequest(m.api, actor)
}

// CreateComment Doc: https://dev.vk.com/ru/method/market.createComment
func (m *Market) CreateComment(actor actor.Actor) *request.MarketCreateCommentRequest {
	return request.NewMarketCreateCommentRequest(m.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/market.delete
func (m *Market) Delete(actor actor.Actor) *request.MarketDeleteRequest {
	return request.NewMarketDeleteRequest(m.api, actor)
}

// DeleteAlbum Doc: https://dev.vk.com/ru/method/market.deleteAlbum
func (m *Market) DeleteAlbum(actor actor.Actor) *request.MarketDeleteAlbumRequest {
	return request.NewMarketDeleteAlbumRequest(m.api, actor)
}

// DeleteComment Doc: https://dev.vk.com/ru/method/market.deleteComment
func (m *Market) DeleteComment(actor actor.Actor) *request.MarketDeleteCommentRequest {
	return request.NewMarketDeleteCommentRequest(m.api, actor)
}

// DeleteProperty Doc: https://dev.vk.com/ru/method/market.deleteProperty
func (m *Market) DeleteProperty(actor actor.Actor) *request.MarketDeletePropertyRequest {
	return request.NewMarketDeletePropertyRequest(m.api, actor)
}

// DeletePropertyVariant Doc: https://dev.vk.com/ru/method/market.deletePropertyVariant
func (m *Market) DeletePropertyVariant(actor actor.Actor) *request.MarketDeletePropertyVariantRequest {
	return request.NewMarketDeletePropertyVariantRequest(m.api, actor)
}

// Edit Doc: https://dev.vk.com/ru/method/market.edit
func (m *Market) Edit(actor actor.Actor) *request.MarketEditRequest {
	return request.NewMarketEditRequest(m.api, actor)
}

// EditAlbum Doc: https://dev.vk.com/ru/method/market.editAlbum
func (m *Market) EditAlbum(actor actor.Actor) *request.MarketEditAlbumRequest {
	return request.NewMarketEditAlbumRequest(m.api, actor)
}

// EditComment Doc: https://dev.vk.com/ru/method/market.editComment
func (m *Market) EditComment(actor actor.Actor) *request.MarketEditCommentRequest {
	return request.NewMarketEditCommentRequest(m.api, actor)
}

// EditOrder Doc: https://dev.vk.com/ru/method/market.editOrder
func (m *Market) EditOrder(actor actor.Actor) *request.MarketEditOrderRequest {
	return request.NewMarketEditOrderRequest(m.api, actor)
}

// EditProperty Doc: https://dev.vk.com/ru/method/market.editProperty
func (m *Market) EditProperty(actor actor.Actor) *request.MarketEditPropertyRequest {
	return request.NewMarketEditPropertyRequest(m.api, actor)
}

// EditPropertyVariant Doc: https://dev.vk.com/ru/method/market.editPropertyVariant
func (m *Market) EditPropertyVariant(actor actor.Actor) *request.MarketEditPropertyVariantRequest {
	return request.NewMarketEditPropertyVariantRequest(m.api, actor)
}

// FilterCategories Doc: https://dev.vk.com/ru/method/market.filterCategories
func (m *Market) FilterCategories(actor actor.Actor) *request.MarketFilterCategoriesRequest {
	return request.NewMarketFilterCategoriesRequest(m.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/market.get
func (m *Market) Get(actor actor.Actor) *request.MarketGetRequest {
	return request.NewMarketGetRequest(m.api, actor)
}

// GetAlbumById Doc: https://dev.vk.com/ru/method/market.getAlbumById
func (m *Market) GetAlbumById(actor actor.Actor) *request.MarketGetAlbumByIDRequest {
	return request.NewMarketGetAlbumByIDRequest(m.api, actor)
}

// GetAlbums Doc: https://dev.vk.com/ru/method/market.getAlbums
func (m *Market) GetAlbums(actor actor.Actor) *request.MarketGetAlbumsRequest {
	return request.NewMarketGetAlbumsRequest(m.api, actor)
}

// GetById Doc: https://dev.vk.com/ru/method/market.getById
func (m *Market) GetById(actor actor.Actor) *request.MarketGetByIDRequest {
	return request.NewMarketGetByIDRequest(m.api, actor)
}

// GetCategories Doc: https://dev.vk.com/ru/method/market.getCategories
func (m *Market) GetCategories(actor actor.Actor) *request.MarketGetCategoriesRequest {
	return request.NewMarketGetCategoriesRequest(m.api, actor)
}

// GetComments Doc: https://dev.vk.com/ru/method/market.getComments
func (m *Market) GetComments(actor actor.Actor) *request.MarketGetCommentsRequest {
	return request.NewMarketGetCommentsRequest(m.api, actor)
}

// GetCommentsExtended Doc: https://dev.vk.com/ru/method/market.getComments
func (m *Market) GetCommentsExtended(actor actor.Actor) *request.MarketGetCommentsExtendedRequest {
	return request.NewMarketGetCommentsExtendedRequest(m.api, actor)
}

// GetGroupOrders Doc: https://dev.vk.com/ru/method/market.getGroupOrders
func (m *Market) GetGroupOrders(actor actor.Actor) *request.MarketGetGroupOrdersRequest {
	return request.NewMarketGetGroupOrdersRequest(m.api, actor)
}

// GetOrderById Doc: https://dev.vk.com/ru/method/market.getOrderById
func (m *Market) GetOrderById(actor actor.Actor) *request.MarketGetOrderByIDRequest {
	return request.NewMarketGetOrderByIDRequest(m.api, actor)
}

// GetOrderItems Doc: https://dev.vk.com/ru/method/market.getOrderItems
func (m *Market) GetOrderItems(actor actor.Actor) *request.MarketGetOrderItemsRequest {
	return request.NewMarketGetOrderItemsRequest(m.api, actor)
}

// GetOrders Doc: https://dev.vk.com/ru/method/market.getOrders
func (m *Market) GetOrders(actor actor.Actor) *request.MarketGetOrdersRequest {
	return request.NewMarketGetOrdersRequest(m.api, actor)
}

// GetOrdersExtended Doc: https://dev.vk.com/ru/method/market.getOrders
func (m *Market) GetOrdersExtended(actor actor.Actor) *request.MarketGetOrdersExtendedRequest {
	return request.NewMarketGetOrdersExtendedRequest(m.api, actor)
}

// GetProductPhotoUploadServer Doc: https://dev.vk.com/ru/method/market.getProductPhotoUploadServer
func (m *Market) GetProductPhotoUploadServer(actor actor.Actor) *request.MarketGetProductPhotoUploadServerRequest {
	return request.NewMarketGetProductPhotoUploadServerRequest(m.api, actor)
}

// GetProperties Doc: https://dev.vk.com/ru/method/market.getProperties
func (m *Market) GetProperties(actor actor.Actor) *request.MarketGetPropertiesRequest {
	return request.NewMarketGetPropertiesRequest(m.api, actor)
}

// GroupItems Doc: https://dev.vk.com/ru/method/market.groupItems
func (m *Market) GroupItems(actor actor.Actor) *request.MarketGroupItemsRequest {
	return request.NewMarketGroupItemsRequest(m.api, actor)
}

// RemoveFromAlbum Doc: https://dev.vk.com/ru/method/market.removeFromAlbum
func (m *Market) RemoveFromAlbum(actor actor.Actor) *request.MarketRemoveFromAlbumRequest {
	return request.NewMarketRemoveFromAlbumRequest(m.api, actor)
}

// ReorderAlbums Doc: https://dev.vk.com/ru/method/market.reorderAlbums
func (m *Market) ReorderAlbums(actor actor.Actor) *request.MarketReorderAlbumsRequest {
	return request.NewMarketReorderAlbumsRequest(m.api, actor)
}

// ReorderItems Doc: https://dev.vk.com/ru/method/market.reorderItems
func (m *Market) ReorderItems(actor actor.Actor) *request.MarketReorderItemsRequest {
	return request.NewMarketReorderItemsRequest(m.api, actor)
}

// Report Doc: https://dev.vk.com/ru/method/market.report
func (m *Market) Report(actor actor.Actor) *request.MarketReportRequest {
	return request.NewMarketReportRequest(m.api, actor)
}

// ReportComment Doc: https://dev.vk.com/ru/method/market.reportComment
func (m *Market) ReportComment(actor actor.Actor) *request.MarketReportCommentRequest {
	return request.NewMarketReportCommentRequest(m.api, actor)
}

// Restore Doc: https://dev.vk.com/ru/method/market.restore
func (m *Market) Restore(actor actor.Actor) *request.MarketRestoreRequest {
	return request.NewMarketRestoreRequest(m.api, actor)
}

// RestoreComment Doc: https://dev.vk.com/ru/method/market.restoreComment
func (m *Market) RestoreComment(actor actor.Actor) *request.MarketRestoreCommentRequest {
	return request.NewMarketRestoreCommentRequest(m.api, actor)
}

// SaveProductPhoto Doc: https://dev.vk.com/ru/method/market.saveProductPhoto
func (m *Market) SaveProductPhoto(actor actor.Actor) *request.MarketSaveProductPhotoRequest {
	return request.NewMarketSaveProductPhotoRequest(m.api, actor)
}

// Search Doc: https://dev.vk.com/ru/method/market.search
func (m *Market) Search(actor actor.Actor) *request.MarketSearchRequest {
	return request.NewMarketSearchRequest(m.api, actor)
}

// SearchItems Doc: https://dev.vk.com/ru/method/market.searchItems
func (m *Market) SearchItems(actor actor.Actor) *request.MarketSearchItemsRequest {
	return request.NewMarketSearchItemsRequest(m.api, actor)
}

// UngroupItems Doc: https://dev.vk.com/ru/method/market.ungroupItems
func (m *Market) UngroupItems(actor actor.Actor) *request.MarketUngroupItemsRequest {
	return request.NewMarketUngroupItemsRequest(m.api, actor)
}
