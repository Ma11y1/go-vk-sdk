package request

import (
	"context"
	"encoding/json"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
)

// Doc: https://dev.vk.com/ru/method/execute

// ExecuteRequest The method executes the VKScript code passed to it, in which other api methods can be called, with intermediate results being saved and processed.
//
//	Doc: https://dev.vk.com/ru/method/execute
type ExecuteRequest struct {
	BaseRequest
}

func NewExecuteRequest(a *api.API, actor actor.Actor) *ExecuteRequest {
	return &ExecuteRequest{*NewMethodBaseRequest(a, actor, "execute")}
}

func (r *ExecuteRequest) Exec(ctx context.Context, target interface{}) error {
	resp := &response.ExecuteResponse{}
	err := r.PostUnmarshal(ctx, resp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Response, &target)
	if err != nil {
		return err
	}

	if resp.ExecuteErrors != nil {
		return &resp.ExecuteErrors
	}

	return err
}

func (r *ExecuteRequest) Code(code string) *ExecuteRequest {
	if code != "" {
		r.parameters.Set(constants.ParameterNameCode, code)
	}
	return r
}
