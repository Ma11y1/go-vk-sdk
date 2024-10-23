package objects

type Gift struct {
	ID          int         `json:"id"`
	FromID      int         `json:"from_id"`
	Description string      `json:"description"`
	Date        int         `json:"date"`
	Message     string      `json:"message"`
	Gift        GiftsLayout `json:"gift"`
	GiftHash    string      `json:"gift_hash"`
	Privacy     int         `json:"privacy"`
	PaymentType string      `json:"payment_type"`
	Price       int         `json:"price"`
	PriceStr    string      `json:"price_str"`
}

type GiftsLayout struct {
	ID                int            `json:"id"`
	Thumb256          string         `json:"thumb_256"`
	Thumb48           string         `json:"thumb_48"`
	Thumb96           string         `json:"thumb_96"`
	StickersProductID int            `json:"stickers_product_id"`
	IsStickersStyle   NumberFlagBool `json:"is_stickers_style"`
}
