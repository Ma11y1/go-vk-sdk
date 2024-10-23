package objects

import "go-vk-sdk/constants"

type AppWidgetsImage struct {
	ID     string                        `json:"id"`
	Type   constants.AppWidgetsImageType `json:"type"`
	Images []Image                       `json:"images"`
}
