package payments

import (
	"strconv"
	"strings"
)

type GetItemRequest struct {
	Notification
	ReceiverID int    `schema:"receiver_id,required"`
	OrderID    int    `schema:"order_id,required"`
	Lang       string `schema:"lang,required"` // User language as language_country.
	Item       string `schema:"item,required"` // Product name passed to the purchase dialog box.
}

type OrderStatusChangeRequest struct {
	Notification
	ReceiverID int `schema:"receiver_id,required"` // The recipient’s ID
	OrderID    int `schema:"order_id,required"`    // Order ID in VK payments system
	Date       int `schema:"date,required"`        // when the order was created (as "unix timestamp")

	// Status New status of the order
	//
	// chargeable - order ready for payments. User shall complete the order
	// in the application. In case of successful reply the payments system will
	// charge votes to the application account. In case of error message the
	// order will be canceled.
	//
	// refunded - Available as of API version 5.132. Order cancelled. It is necessary
	// to pick up game values given to the user for payments.
	Status Status `schema:"status,required"`

	// Item Product name passed to the purchase dialog box.
	//
	// For special offer (call of the payments dialog box with type=offers), Item
	// parameter will include "offer_{offer_id}", item_price parameter is the
	// number of votes charged for such special offer.
	Item               string `schema:"item,required"`
	ItemID             string `schema:"item_id"`
	ItemTitle          string `schema:"item_title,required"`
	ItemPhotoURL       string `schema:"item_photo_url"`
	ItemPrice          string `schema:"item_price,required"`
	ItemCurrencyAmount string `schema:"item_currency_amount"` // Cost in application currency. Doc: https://dev.vk.com/ru/api/payments/getting-started
	ItemDiscount       string `schema:"item_discount"`
}

// OfferID returns offer_id.
func (r OrderStatusChangeRequest) OfferID() int {
	offerID, _ := strconv.Atoi(strings.TrimPrefix(r.Item, "offer_"))
	return offerID
}

type GetSubscriptionRequest struct {
	Notification
	ReceiverID int          `schema:"receiver_id,required"`
	OrderID    int          `schema:"order_id,required"`
	Lang       LanguageType `schema:"lang,required"`
	Item       string       `schema:"item,required"` // Product name passed to the subscription dialog window.
}

type SubscriptionStatusChangeRequest struct {
	Notification
	CancelReason Reason `schema:"cancel_reason"`
	ItemID       string `schema:"item_id,required"` // Product identifier in the application.
	ItemPrice    string `schema:"item_price,required"`

	// Status New subscription status. Possible values:
	//
	// "chargeable" — subscription is ready for payments. An order for the user
	// needs to be processed within the application. If the response is
	// successful, the payments system credits votes to the application
	// balance. If the response is a failure, the order is canceled.
	//
	// "active" — subscription is active;
	//
	// "cancelled" — subscription is canceled.
	Status         Status `schema:"status,required"`
	PendingCancel  int    `schema:"pending_cancel"` // Subscription is active until the end of the paid period (status = active). integer, [1]
	SubscriptionID int    `schema:"subscription_id,required"`
	NextBillTime   int    `schema:"next_bill_time"` // Subscription next bill time
}
