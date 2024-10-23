package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/market

type MarketAddResponse struct {
	BaseResponse
	Response struct {
		MarketItemID int `json:"market_item_id"` // Item ID
	} `json:"response"`
}

type MarketAddAlbumResponse struct {
	BaseResponse
	Response struct {
		MarketAlbumID int `json:"market_album_id"` // Album ID
		AlbumsCount   int `json:"albums_count"`
	} `json:"response"`
}

type MarketAddToAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketCreateCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketDeleteResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketDeleteAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketDeleteCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketEditResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketEditAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketEditCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketEditOrderResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketGetResponse struct {
	BaseResponse
	Response struct {
		Count int                  `json:"count"`
		Items []objects.MarketItem `json:"items"`
	} `json:"response"`
}

type MarketGetAlbumByIDResponse struct {
	BaseResponse
	Response struct {
		Count int                   `json:"count"`
		Items []objects.MarketAlbum `json:"items"`
	} `json:"response"`
}

type MarketGetAlbumsResponse struct {
	BaseResponse
	Response struct {
		Count int                   `json:"count"`
		Items []objects.MarketAlbum `json:"items"`
	} `json:"response"`
}

type MarketGetByIDResponse struct {
	BaseResponse
	Response struct {
		Count int                  `json:"count"`
		Items []objects.MarketItem `json:"items"`
	} `json:"response"`
}

type MarketGetCategoriesResponse struct {
	BaseResponse
	Response struct {
		Items []objects.MarketCategoryTree `json:"items"`
	} `json:"response"`
}

type MarketGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.WallWallComment `json:"items"`
	} `json:"response"`
}

type MarketGetCommentsExtendedResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.WallWallComment `json:"items"`
		objects.UsersAndGroups
	} `json:"response"`
}

type MarketGetGroupOrdersResponse struct {
	BaseResponse
	Response struct {
		Count int                   `json:"count"`
		Items []objects.MarketOrder `json:"items"`
	} `json:"response"`
}

type MarketGetOrderByIDResponse struct {
	BaseResponse
	Response struct {
		Order objects.MarketOrder `json:"order"`
	} `json:"response"`
}

type MarketGetOrderItemsResponse struct {
	BaseResponse
	Response struct {
		Count int                       `json:"count"`
		Items []objects.MarketOrderItem `json:"items"`
	} `json:"response"`
}

type MarketRemoveFromAlbumResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketReorderAlbumsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketReorderItemsResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketReportResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketReportCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketRestoreResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketRestoreCommentResponse struct {
	BaseResponse
	Response int `json:"response"`
}

type MarketSearchResponse struct {
	BaseResponse
	Response struct {
		Count    int                  `json:"count"`
		Items    []objects.MarketItem `json:"items"`
		ViewType int                  `json:"view_type"`
	} `json:"response"`
}

type MarketSearchItemsResponse struct {
	BaseResponse
	Response struct {
		Count    int                  `json:"count"`
		ViewType int                  `json:"view_type"`
		Items    []objects.MarketItem `json:"items"`
		Groups   []objects.GroupFull  `json:"groups,omitempty"`
	} `json:"response"`
}
