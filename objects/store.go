package objects

type Product struct {
	ID                 int         `json:"id"`
	Type               string      `json:"type"`
	Active             BoolInt     `json:"active"` // Information whether the product is active (1 - yes, 0 - no)
	Promoted           BoolInt     `json:"promoted"`
	Purchased          BoolInt     `json:"purchased"`
	PurchaseDate       int         `json:"purchase_date"`
	Title              string      `json:"title"`
	Subtitle           string      `json:"subtitle"`
	TitleLangKey       string      `json:"title_lang_key"`
	DescriptionLangKey string      `json:"description_lang_key"`
	BaseURL            string      `json:"base_url"`
	URL                string      `json:"url"`
	Copyright          string      `json:"copyright"`
	Icon               ProductIcon `json:"icon"`          // Array of icon images or icon set object of the product (for stickers product type)
	BaseID             int         `json:"base_id"`       // Id of the base pack (for sticker pack styles)
	HasAnimation       bool        `json:"has_animation"` // Information whether the product is an animated sticker pack (for stickers product type)
	IsNew              bool        `json:"is_new"`        // Information whether sticker product wasn't used after being purchased
	IsPopup            bool        `json:"is_popup"`      // Information whether the product is a sticker pack with popup stickers (for stickers product type)
	IsVmoji            bool        `json:"is_vmoji"`      // Information whether sticker pack is a vmoji pack
	PaymentRegion      string      `json:"payment_region"`
	Previews           []Image     `json:"previews"`
	Stickers           []Sticker   `json:"stickers"`
	StyleIDs           []int       `json:"style_ids"`         // Array of style ids available for the sticker pack
	StyleStickerIDs    []int       `json:"style_sticker_ids"` // Array of style sticker ids (for sticker pack styles)
	StickerIDs         []int       `json:"sticker_ids"`
}

type ProductIcon struct {
	Version int     `json:"version"`
	BaseURL string  `json:"base_url"`
	Images  []Image `json:"images"`
}

type StickersKeywords struct {
	BaseURL    string            `json:"base_url"`
	Count      int               `json:"count"`
	ChunkCount int               `json:"chunk_count"` // Total count of chunks to load
	ChunkHash  string            `json:"chunk_hash"`  // Chunks version hash
	Dictionary []StickersKeyword `json:"dictionary"`
}

type StickersKeyword struct {
	PromotedStickers []Sticker         `json:"promoted_stickers"`
	UserStickers     []Sticker         `json:"user_stickers"`
	Words            []string          `json:"words"`
	Stickers         []StickersKeyword `json:"stickers"`
}

type StickerKeyword struct {
	PackID    int `json:"pack_id"`
	StickerID int `json:"sticker_id"`
}
