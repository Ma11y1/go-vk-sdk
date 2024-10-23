package objects

type PrettyCardsPrettyCard struct {
	Button     string  `json:"button"`      // Button key
	ButtonText string  `json:"button_text"` // Button text in current language
	CardID     string  `json:"card_id"`     // Card ID (long int returned as string)
	Images     []Image `json:"images"`
	LinkURL    string  `json:"link_url"`  // Link url
	Photo      string  `json:"photo"`     // Photo ID (format "<owner_id>_<media_id>")
	Price      string  `json:"price"`     // Price if set (decimal number returned as string)
	PriceOld   string  `json:"price_old"` // Old price if set (decimal number returned as string)
	Title      string  `json:"title"`     // Title
}
