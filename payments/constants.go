package payments

type LanguageType string

const (
	LangRussian    LanguageType = "ru_RU"
	LangUkrainian  LanguageType = "uk_UA"
	LangBelarusian LanguageType = "be_BY"
	LangEnglish    LanguageType = "en_US"
)

type Status string

const (
	// StatusChargeable Subscription is ready for payments. An order for the user
	// needs to be processed within the application. If the response is
	// successful, the payments system credits votes to the application
	// balance. If the response is a failure, the order is canceled.
	StatusChargeable Status = "chargeable"
	StatusRefunded   Status = "refunded"  // Available as of API version 5.132. Order cancelled. It is necessary to pick up game values given to the user for payments.
	StatusActive     Status = "active"    // Subscription is active.
	StatusCancelled  Status = "cancelled" // Subscription is canceled.
)

type Reason string

const (
	ReasonUserDecision Reason = "user_decision" // Subscription canceled by the user.
	ReasonAppDecision  Reason = "app_decision"  // Subscription canceled by the application (orders.cancelSubscription).
	ReasonPaymentFail  Reason = "payment_fail"  // Subscription canceled due to a failed payments.
	ReasonUnknown      Reason = "unknown"       // Subscription canceled for a different reason.
)
