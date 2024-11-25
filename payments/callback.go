package payments

import (
	"context"
	"encoding/json"
	"github.com/gorilla/schema"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"go-vk-sdk/transport"
	"net/http"
	"net/url"
)

// Callback
//
//	Doc: https://dev.vk.com/ru/api/payments/getting-started
type Callback struct {
	server                       transport.CallbackServer
	secret                       string
	getItem                      func(e *GetItemRequest) (*GetItemResponse, *internalErrors.PaymentsError)
	getItemTest                  func(e *GetItemRequest) (*GetItemResponse, *internalErrors.PaymentsError)
	orderStatusChange            func(e *OrderStatusChangeRequest) (*OrderStatusChangeResponse, *internalErrors.PaymentsError)
	orderStatusChangeTest        func(e *OrderStatusChangeRequest) (*OrderStatusChangeResponse, *internalErrors.PaymentsError)
	getSubscription              func(e *GetSubscriptionRequest) (*GetSubscriptionResponse, *internalErrors.PaymentsError)
	getSubscriptionTest          func(e *GetSubscriptionRequest) (*GetSubscriptionResponse, *internalErrors.PaymentsError)
	subscriptionStatusChange     func(e *SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *internalErrors.PaymentsError)
	subscriptionStatusChangeTest func(e *SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *internalErrors.PaymentsError)
}

func NewCallback(url *url.URL, secret string) *Callback {
	return &Callback{
		server: transport.NewBaseCallbackServer(url),
		secret: secret,
	}
}

func NewCallbackServer(server transport.CallbackServer, secret string) *Callback {
	return &Callback{
		server: server,
		secret: secret,
	}
}

func (c *Callback) handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// NOTE: what about net error?
		c.sendResponse(w, &response{
			Error: &internalErrors.PaymentsError{
				Code:       internalErrors.PaymentsBadRequest,
				Msg:        err.Error(),
				IsCritical: true,
			},
		})

		return
	}

	// If signatures do not match, give the error 10 in response
	if r.PostForm.Get("sig") != Sign(r.PostForm, c.secret) {
		c.sendResponse(w, &response{
			Error: &internalErrors.PaymentsError{
				Code:       internalErrors.PaymentsBadSignatures,
				Msg:        "The calculated and sent signatures do not match",
				IsCritical: true,
			},
		})

		return
	}

	resp, err := c.handleNotification(r.PostForm)
	if err != nil {
		c.sendResponse(w, &response{
			Error: &internalErrors.PaymentsError{
				Code:       internalErrors.PaymentsBadRequest,
				Msg:        err.Error(),
				IsCritical: true,
			},
		})

		return
	}

	c.sendResponse(w, resp)
}

func (c *Callback) handleNotification(u url.Values) (*response, error) {
	r := &response{}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	t := NotificationType(u.Get("notification_type"))

	switch {
	case t == GetItem && c.getItem != nil:
		var event GetItemRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.getItem(&event)
	case t == GetItem.TestMode() && c.getItemTest != nil:
		var event GetItemRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.getItemTest(&event)
	case t == GetSubscription && c.getSubscription != nil:
		var event GetSubscriptionRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.getSubscription(&event)
	case t == GetSubscription.TestMode() && c.getSubscriptionTest != nil:
		var event GetSubscriptionRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.getSubscriptionTest(&event)
	case t == OrderStatusChange && c.orderStatusChange != nil:
		var event OrderStatusChangeRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.orderStatusChange(&event)
	case t == OrderStatusChange.TestMode() && c.orderStatusChangeTest != nil:
		var event OrderStatusChangeRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.orderStatusChangeTest(&event)
	case t == SubscriptionStatusChange && c.subscriptionStatusChange != nil:
		var event SubscriptionStatusChangeRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.subscriptionStatusChange(&event)
	case t == SubscriptionStatusChange.TestMode() && c.subscriptionStatusChangeTest != nil:
		var event SubscriptionStatusChangeRequest
		if err := decoder.Decode(&event, u); err != nil {
			return nil, internalErrors.ErrorLog("Payments.Callback.handleNotification()", err.Error())
		}

		r.Response, r.Error = c.subscriptionStatusChangeTest(&event)
	default:
		r.Error = &internalErrors.PaymentsError{
			Code:       internalErrors.PaymentsCommonError,
			Msg:        string(t) + " not processed",
			IsCritical: true,
		}
	}

	return r, nil
}

func (c *Callback) sendResponse(w http.ResponseWriter, data *response) {
	w.Header().Add("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(data)
}

func (c *Callback) Run() error {
	err := c.server.Run()
	if err != nil {
		return internalErrors.ErrorLog("Payments.Callback.Run()", "Failed to start callback server: "+err.Error())
	}
	logger.Log("Payments.Callback.Run()", "Payments server is running at url: "+c.server.GetURL().String())
	return nil
}

func (c *Callback) Stop(ctx context.Context) error {
	err := c.server.Stop(ctx)
	logger.Log("Payments.Callback.Stop()", "Payments server is stopped at url: "+c.server.GetURL().String())
	if err != nil {
		return internalErrors.ErrorLog("Payments.Callback.Stop()", "Failed to stop callback server: "+err.Error())
	}
	return nil
}

func (c *Callback) IsRunning() bool {
	return c.server.IsRunning()
}

func (c *Callback) SetDefaultHandler(path string) error {
	if path == "" {
		return internalErrors.ErrorLog("Payments.Callback.SetDefaultHandler()", "Invalid value handle path. Path is empty")
	}
	return c.SetHandleFunc(path, c.handle)
}

func (c *Callback) SetHandler(path string, handler http.Handler) error {
	if path == "" {
		return internalErrors.ErrorLog("Payments.Callback.SetHandler()", "Invalid value handle path. Path is empty")
	}
	return c.server.SetHandler(path, handler)
}

func (c *Callback) SetHandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) error {
	if path == "" {
		return internalErrors.ErrorLog("Payments.Callback.SetHandleFunc()", "Invalid value handle path. Path is empty")
	}
	return c.server.SetHandleFunc(path, fn)
}

// OnGetItem notification is sent when the product purchase dialog box is
// opened in the application. Then, obtained information about the product
// is shown in the purchase dialog box.
//
// With item specifying the product name received in the notification, the
// developer shall return actual information about such product. When there
// is no product available, reply with Error 20 "Product does not exist".
//
//	&errors.PaymentsError{
//		Code:     payments.ProductNotExist,
//		Msg:      "Product does not exist",
//		Critical: true,
//	}
//
// Please note! As item is passed at the client side, user can change this
// parameters.
//
// Please note! We recommend you to use product caching for all products,
// this will decrease the number of calls to the application server and users
// will not wait when information about them is retrieved.
//
// In case information about the product with item ID was cached for
// expiration period, following requests for such product will not be run
// within the given time.
//
// Doc: https://dev.vk.com/ru/api/payments/notifications/get-item
func (c *Callback) OnGetItem(f func(e *GetItemRequest) (*GetItemResponse, *internalErrors.PaymentsError), isTest bool) {
	if isTest {
		c.getItemTest = f
	} else {
		c.getItem = f
	}
}

// OnOrderStatusChange
//
// Please note! In case of repeated notification of Changing order status type
// (with the same order_id), the reply shall be the exact copy of the reply
// for initial notification.
//
// For example, if such notification was sent and positive reply was received
// but it was not received by VK, or for any temporary reason the blocked
// funds could not be charged to the application account immediately after
// reply, such notification will be resent. And there is no need to place a
// new order, you just send the same reply as before (by keeping order_id and
// checking whether such notification was received).
//
// Doc: https://dev.vk.com/ru/payments/notifications/order-status-change
func (c *Callback) OnOrderStatusChange(f func(e *OrderStatusChangeRequest) (*OrderStatusChangeResponse, *internalErrors.PaymentsError), isTest bool) {
	if isTest {
		c.orderStatusChangeTest = f
	} else {
		c.orderStatusChange = f
	}
}

// OnGetSubscription is sent when a subscription dialog window is opened via application.
//
// Doc: https://dev.vk.com/ru/api/payments/notifications/get-subscription
func (c *Callback) OnGetSubscription(f func(e *GetSubscriptionRequest) (*GetSubscriptionResponse, *internalErrors.PaymentsError), isTest bool) {
	if isTest {
		c.getSubscriptionTest = f
	} else {
		c.getSubscription = f
	}
}

// OnSubscriptionStatusChange is sent the moment the subscription status
// changes. Please note that the subscription status doesn’t change after
// renewal (withdrawal of a new payments from a user). You won’t receive a
// payments notification about subscription renewal. It’s sent only if a user
// decides to not renew and thus cancels the subscription.
// Warning! If the notification of Change in subscription status type (with
// the same subscription_id) occurs repeatedly, the response should be
// identical to the response of the initial notification.
//
// For example, if this notification was sent and a positive response was
// received, but it didn’t reach VK servers or for some temporary reasons
// transferring frozen funds to the application’s account failed right after
// receiving the response, the same notification will be sent again. There is
// no need to issue a new order. It is necessary to send the same response as
// the previous time (having saved the subscription_id and verified through it
// as if the same notification had already been received).
//
// Note that if the subscription was canceled due to lack of funds on a user’s
// account, it can be renewed within five days (provided that the user
// replenishes their account balance within these five days). In this case, the
// existing subscription is the one renewed, and there is no need to create a
// new one.
//
// Doc: https://dev.vk.com/ru/api/payments/notifications/subscription-status-change
func (c *Callback) OnSubscriptionStatusChange(f func(e *SubscriptionStatusChangeRequest) (*SubscriptionStatusChangeResponse, *internalErrors.PaymentsError), isTest bool) {
	if isTest {
		c.subscriptionStatusChangeTest = f
	} else {
		c.subscriptionStatusChange = f
	}
}
