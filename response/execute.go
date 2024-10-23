package response

import (
	"go-vk-sdk/errors"
)

type ExecuteResponse struct {
	BaseResponse
	ExecuteErrors errors.ExecuteAPIErrors `json:"execute_errors"`
	Response      []byte                  `json:"response"`
}

func (e *ExecuteResponse) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}

	return append([]byte{}, e.Response...), nil
}

func (e *ExecuteResponse) UnmarshalJSON(data []byte) error {
	e.Response = append(e.Response[0:0], data...)
	return nil
}
