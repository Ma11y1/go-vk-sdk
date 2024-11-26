package payments

import internalErrors "go-vk-sdk/errors"

type response struct {
	Response interface{}                   `json:"response,omitempty"`
	Error    *internalErrors.PaymentsError `json:"error,omitempty"`
}

type GetItemResponse struct {
	Title    string `json:"title"`               // product name, max 48 characters. Required parameter
	PhotoURL string `json:"photo_url,omitempty"` // URL of product image on the developer's server
	Price    int    `json:"price"`               // product price, in votes. Required parameter
	ItemID   string `json:"item_id,omitempty"`   // product ID in the application

	// the size of the discount on the goods in the votes. Should be within
	// 1 to 1000 votes and less than the price of goods.
	Discount int `json:"discount,omitempty"`

	// Expiration allows product caching for {Expiration} seconds. Allowed
	// range from 600 to 604800 seconds.
	//
	// Warning! In the absence of the parameter is possible to cache the goods
	// for 3600 seconds with a large number of consecutive identical
	// responses. For cancellation of caching it is necessary to pass 0 as
	// value of parameter.
	Expiration int `json:"expiration,omitempty"`
}

type OrderStatusChangeResponse struct {
	OrderID    int `json:"order_id"`               // order ID in VK payments system
	AppOrderID int `json:"app_order_id,omitempty"` // ID of the order in the application. Shall be unique for each order
}

// GetSubscriptionResponse
//
// Using the item identifier received in the notification, developers should
// return current information about it. If there is no product, error 20
// “Subscription does not exist” should be returned in the response.
//
//	&errors.Error{
//		Code:     payments.ProductNotExist,
//		Msg:      "Subscription does not exist",
//		Critical: true,
//	}
//
// Warning! Due to the item being sent on the client-side, this parameter can be changed by users.
type GetSubscriptionResponse struct {
	ItemID        int    `json:"item_id,omitempty"`
	Title         string `json:"title"`
	PhotoURL      string `json:"photo_url,omitempty"`      // Image URL on the developer’s server for the subscription.
	Price         int    `json:"price"`                    // Subscription price shown in votes.
	Period        int    `json:"period"`                   // Subscription period duration in days. Possible values: 3, 7, 30.
	TrialDuration int    `json:"trial_duration,omitempty"` // Trial period duration in days. Possible values: 3, 7, 30.

	// Allows product caching for {expiration} seconds. Permitted range is
	// from 600 to 604,800 seconds.
	//
	// Warning! Without the parameter, caching of the product is possible for
	// 3,600 seconds if many identical consecutive responses occur. To disable
	// caching, it is necessary to pass 0 as the parameter value.
	Expiration int `json:"expiration,omitempty"`
}

type SubscriptionStatusChangeResponse struct {
	SubscriptionID int `json:"subscription_id"`        // Global subscription identifier.
	AppOrderID     int `json:"app_order_id,omitempty"` // Order identifier within the app. Must be unique for each order.
}
