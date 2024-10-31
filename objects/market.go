package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-vk-sdk/constants"
)

type MarketCurrency struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

type MarketAlbum struct {
	Count       int     `json:"count"`
	ID          int     `json:"id"`
	OwnerID     int     `json:"owner_id"`
	Photo       Photo   `json:"photo"`
	Title       string  `json:"title"`
	UpdatedTime int     `json:"updated_time"`
	IsMain      BoolInt `json:"is_main"`
	IsHidden    BoolInt `json:"is_hidden"`
}

func (m *MarketAlbum) ToAttachment() string {
	return fmt.Sprintf("market_album%d_%d", m.OwnerID, m.ID)
}

type MarketCategory struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Parent *MarketCategory `json:"parent"`
}

type MarketItem struct {
	ID                 int                  `json:"id"`
	Title              string               `json:"title"`
	Description        string               `json:"description"`
	AccessKey          string               `json:"access_key"`
	Date               int                  `json:"date,omitempty"`
	Availability       int                  `json:"availability"`
	Category           MarketCategory       `json:"category"`
	OwnerID            int                  `json:"owner_id"`
	Price              MarketPrice          `json:"price"`
	ThumbPhoto         string               `json:"thumb_photo"`
	CanComment         BoolInt              `json:"can_comment"`
	CanRepost          BoolInt              `json:"can_repost"`
	IsFavorite         BoolInt              `json:"is_favorite"`
	IsMainVariant      BoolInt              `json:"is_main_variant"`
	AlbumsIDs          []int                `json:"albums_ids"`
	Photos             []Photo              `json:"photos"`
	Likes              LikesInfo            `json:"likes"`
	Reposts            RepostsInfo          `json:"reposts"`
	ViewsCount         int                  `json:"views_count,omitempty"`
	URL                string               `json:"url"`
	ButtonTitle        string               `json:"button_title"`
	ExternalID         string               `json:"external_id"`
	Dimensions         MarketDimensions     `json:"dimensions"`
	Weight             int                  `json:"weight"`
	VariantsGroupingID int                  `json:"variants_grouping_id"`
	PropertyValues     []MarketItemProperty `json:"property_values"`
	CartQuantity       int                  `json:"cart_quantity"`
	SKU                string               `json:"sku"`
}

func (m *MarketItem) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("false")) {
		return nil
	}

	type renamedMarketMarketItem MarketItem

	var r renamedMarketMarketItem

	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("objects.MarketItem: %w", err)
	}

	*m = MarketItem(r)

	return nil
}

type MarketItemProperty struct {
	VariantID    int    `json:"variant_id"`
	VariantName  string `json:"variant_name"`
	PropertyName string `json:"property_name"`
}

type MarketDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"length"`
}

func (m *MarketItem) ToAttachment() string {
	return fmt.Sprintf("m%d_%d", m.OwnerID, m.ID)
}

type MarketPrice struct {
	Amount        string         `json:"amount"`
	Currency      MarketCurrency `json:"currency"`
	DiscountRate  int            `json:"discount_rate"`
	OldAmount     string         `json:"old_amount"`
	Text          string         `json:"text"`
	OldAmountText string         `json:"old_amount_text"`
}

func (m *MarketPrice) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}

	type renamedMarketPrice MarketPrice

	var r renamedMarketPrice

	err := json.Unmarshal(data, &r)
	if err != nil {
		return fmt.Errorf("objects.MarketPrice: %w", err)
	}

	*m = MarketPrice(r)

	return nil
}

type MarketSection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MarketOrder struct {
	ID                int                         `json:"id"`
	GroupID           int                         `json:"group_id"`
	UserID            int                         `json:"user_id"`
	Date              int                         `json:"date"`
	Status            constants.MarketOrderStatus `json:"status"`
	ItemsCount        int                         `json:"items_count"`
	TotalPrice        MarketPrice                 `json:"total_price"`
	DisplayOrderID    string                      `json:"display_order_id"`
	Comment           string                      `json:"comment"`
	PreviewOrderItems []MarketOrderItem           `json:"preview_order_items"`
	PriceDetails      []MarketPriceDetail         `json:"price_details"`
	Delivery          MarketOrderDelivery         `json:"delivery"`
	Recipient         MarketOrderRecipient        `json:"recipient"`
}

type MarketOrderDelivery struct {
	TrackNumber   string              `json:"track_number"`
	TrackLink     string              `json:"track_link"`
	Address       string              `json:"address"`
	Type          string              `json:"type"`
	DeliveryPoint MarketDeliveryPoint `json:"delivery_point,omitempty"`
}

type MarketDeliveryPoint struct {
	ID           int                        `json:"id"`
	ExternalID   string                     `json:"external_id"`
	OutpostOnly  BoolInt                    `json:"outpost_only"`
	CashOnly     BoolInt                    `json:"cash_only"`
	Address      MarketDeliveryPointAddress `json:"address"`
	DisplayTitle string                     `json:"display_title"`
	ServiceID    int                        `json:"service_id"`
}

type MarketDeliveryPointAddress struct {
	ID             int     `json:"id"`
	Address        string  `json:"address"`
	CityID         int     `json:"city_id"`
	CountryID      int     `json:"country_id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Phone          string  `json:"phone"`
	Title          string  `json:"title"`
	WorkInfoStatus string  `json:"work_info_status"`
}

type MarketOrderRecipient struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DisplayText string `json:"display_text"`
}

type MarketOrderItem struct {
	OwnerID  int         `json:"owner_id"`
	ItemID   int         `json:"item_id"`
	Price    MarketPrice `json:"price"`
	Quantity int         `json:"quantity"`
	Item     MarketItem  `json:"item"`
	Title    string      `json:"title"`
	Photo    Photo       `json:"photo"`
	Variants []string    `json:"variants"`
}

type MarketPriceDetail struct {
	Title    string      `json:"title"`
	Price    MarketPrice `json:"price"`
	IsAccent BoolInt     `json:"is_accent,omitempty"`
}

type MarketCategoryTree struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Icon     []Image                `json:"icon"`
	Children []MarketCategoryTree   `json:"children"`
	View     MarketCategoryTreeView `json:"view"`
	URL      string                 `json:"url"`
}

type MarketCategoryTreeView struct {
	Type     string   `json:"type"`
	Selected bool     `json:"selected"`
	RootPath []string `json:"root_path"`
}
