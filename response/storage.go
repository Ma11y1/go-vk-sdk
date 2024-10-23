package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/storage

type StorageGetResponse struct {
	BaseResponse
	Response []objects.KeyValue `json:"response"`
}

func (s *StorageGetResponse) ToMap() map[string]string {
	m := make(map[string]string)
	for _, item := range s.Response {
		m[item.Key] = item.Value
	}

	return m
}

type StorageGetKeysResponse struct {
	BaseResponse
	Response []string `json:"response"`
}

type StorageSetResponse struct {
	BaseResponse
	Response int `json:"response"`
}
