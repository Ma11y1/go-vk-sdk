package objects

type OrdersAmount struct {
	Amounts  []OrdersAmountItem `json:"amounts"`
	Currency string             `json:"currency"`
}

type OrdersAmountItem struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Votes       string `json:"votes"`
}

type OrdersOrder struct {
	Amount              int    `json:"amount"`
	AppOrderID          int    `json:"app_order_id"`
	CancelTransactionID int    `json:"cancel_transaction_id"`
	Date                int    `json:"date"`
	ID                  int    `json:"id"`
	Item                string `json:"item"`
	ReceiverID          int    `json:"receiver_id"`
	Status              string `json:"status"`
	TransactionID       int    `json:"transaction_id"`
	UserID              int    `json:"user_id"`
}

type OrdersSubscription struct {
	CancelReason    string  `json:"cancel_reason"`
	CreateTime      int     `json:"create_time"`
	ID              int     `json:"id"`
	ItemID          string  `json:"item_id"`
	NextBillTime    int     `json:"next_bill_time"`
	Period          int     `json:"period"`
	PeriodStartTime int     `json:"period_start_time"`
	Price           int     `json:"price"`
	Status          string  `json:"status"`
	PendingCancel   BoolInt `json:"pending_cancel"`
	TestMode        BoolInt `json:"test_mode"`
	TrialExpireTime int     `json:"trial_expire_time"`
	UpdateTime      int     `json:"update_time"`
}
