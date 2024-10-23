package response

import (
	"go-vk-sdk/errors"
)

type BaseResponse struct {
	Response interface{}     `json:"response"` // override
	Error    errors.APIError `json:"error"`
}
