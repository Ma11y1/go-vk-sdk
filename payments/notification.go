package payments

type NotificationType string

const (
	// GetItem For acquiring product information.
	//
	// https://dev.vk.com/ru/api/payments/notifications/get-item
	GetItem NotificationType = "get_item"

	// OrderStatusChange For changing the order’s status.
	//
	// https://dev.vk.com/ru/payments/notifications/order-status-change
	OrderStatusChange NotificationType = "order_status_change"

	// GetSubscription For receiving subscription information.
	//
	// https://dev.vk.com/ru/api/payments/notifications/get-subscription
	GetSubscription NotificationType = "get_subscription"

	// SubscriptionStatusChange For changing a subscription’s status.
	//
	// https://dev.vk.com/ru/api/payments/notifications/subscription-status-change
	SubscriptionStatusChange NotificationType = "subscription_status_change"
)

// TestMode return NotificationType for test mode
//
//	In test mode, the suffix '_test' is added to the notification_type parameter.
//	See https://dev.vk.com/ru/api/payments/getting-started
func (t NotificationType) TestMode() NotificationType {
	return t + "_test"
}

// Notification Selection of fields depends on the type of notification.
// However all types of notifications contain the following fields.
type Notification struct {
	Type   NotificationType `schema:"notification_type,required"`
	AppID  int              `schema:"app_id,required"`
	UserID int              `schema:"user_id,required"` // ID of the user who made the order.
	Sig    string           `schema:"sig,required"`     // Notification signature.
}
