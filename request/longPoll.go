package request

import (
	"context"
	"go-vk-sdk/api"
	"go-vk-sdk/constants"
	"go-vk-sdk/response"
	"net/http"
	"strconv"
)

// LongPollUserRequest
//
//	Long Polling is a technology that allows you to obtain data about new events using “long queries”. The server receives the request, but does not send a response to it immediately, but only when some event occurs (for example, a new message arrives) or the specified waiting time has expired.
//	Using this approach, you can instantly display important events in your application. Using the user Long Poll api, you cannot send a message; to do this, use the messages.send method.
//	For the first request within a session, the values for the server, key and ts parameters must be obtained using the messages.getLongPollServer method. In subsequent requests, use the same server and key and the new ts value that will come to you in the response from the Long Poll server.
//
//	After receiving any response, to continue communication, you need to send a request with the new ts received in the last response.
//
//	Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started
type LongPollUserRequest struct {
	BaseRequest
}

func NewLongPollUserRequest(a *api.API, url string) *LongPollUserRequest {
	r := &LongPollUserRequest{BaseRequest{
		url:        url,
		transport:  a.Client,
		header:     &http.Header{},
		parameters: &BaseParameters{},
	}}

	r.parameters.Set(constants.ParameterNameAct, "a_check")
	return r
}

// Exec executes the request and unmarshals the response into LongPollUserResponse
func (r *LongPollUserRequest) Exec(ctx context.Context) (response response.LongPollUserResponse, err error) {
	err = r.GetUnmarshal(ctx, &response)
	return
}

// Key Session secret key
func (r *LongPollUserRequest) Key(key string) *LongPollUserRequest {
	r.parameters.Set(constants.ParameterNameKey, key)
	return r
}

// Ts The number of the last event, starting from which you want to receive data
func (r *LongPollUserRequest) Ts(ts int) *LongPollUserRequest {
	r.parameters.Set(constants.ParameterNameTs, strconv.Itoa(ts))
	return r
}

// Wait
//
//	Waiting time. The maximum value is 90
//	Since some proxy servers drop the connection after 30 seconds, we recommend specifying wait= 25.
func (r *LongPollUserRequest) Wait(period int) *LongPollUserRequest {
	if period > 90 {
		period = 90
	}
	r.parameters.Set(constants.ParameterNameWait, strconv.Itoa(period))
	return r
}

// Mode
//
//	Additional response options. Sum of option codes from the list:
//	• 2 — receive investments.
//	• 8—return an extended set of events.
//	• 32 — return pts (this is required for the messages.getLongPollHistory method to work without the limit of 256 recent events).
//	• 64 — in event code 8 (a friend has become online), return additional data in the $extra field (see Event Structure).
//	• 128 — return the random_id field (random_id can be passed when sending a message using the messages.send method).
func (r *LongPollUserRequest) Mode(mode int) *LongPollUserRequest {
	r.parameters.Set(constants.ParameterNameMode, strconv.Itoa(mode))
	return r
}

// Version
//
//	Current version: 3.
//	For version 0 (default), community identifiers will be sent in the format group_id + 1000000000 to maintain backward compatibility.
//	We recommend using the current version.
func (r *LongPollUserRequest) Version(v int) *LongPollUserRequest {
	r.parameters.Set(constants.ParameterNameVersionLP, strconv.Itoa(v))
	return r
}

// LongPollGroupRequest
//
//	Bots Long Poll api allows you to work with events from your community in real time. Unlike the Callback api,
//	the queue of events is stored on the VK side - we will not send a separate notification for each event.
//	In contrast to user Long Poll, Bots Long Poll api works with community events, not user events.
//	Before connecting to the Long Poll server, you need to obtain session data (server, key, ts) using the groups.getLongPollServer method.
//	• groups.getLongPollServer - gets the address for connecting to the Long Poll server on VK;
//	• groups.getLongPollSettings - gets settings for Bots Long Poll api events;
//	• groups.setLongPollSettings - sets the Bots Long Poll api event settings.
//	Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started
type LongPollGroupRequest struct {
	BaseRequest
}

func NewLongPollGroupRequest(a *api.API, url string) *LongPollGroupRequest {
	r := &LongPollGroupRequest{BaseRequest{
		url:        url,
		transport:  a.Client,
		header:     &http.Header{},
		parameters: &BaseParameters{},
	}}

	r.parameters.Set(constants.ParameterNameAct, "a_check")
	return r
}

// Exec executes the request and unmarshals the response into LongPollGroupResponse
func (r *LongPollGroupRequest) Exec(ctx context.Context) (response response.LongPollGroupResponse, err error) {
	err = r.GetUnmarshal(ctx, &response)
	return
}

// Key Session secret key
func (r *LongPollGroupRequest) Key(key string) *LongPollGroupRequest {
	r.parameters.Set(constants.ParameterNameKey, key)
	return r
}

// Ts The number of the last event, starting from which you want to receive data
func (r *LongPollGroupRequest) Ts(ts string) *LongPollGroupRequest {
	r.parameters.Set(constants.ParameterNameTs, ts)
	return r
}

// Wait
//
//	Waiting time. The maximum value is 90
//	Since some proxy servers drop the connection after 30 seconds, we recommend specifying wait= 25.
func (r *LongPollGroupRequest) Wait(period int) *LongPollGroupRequest {
	if period > 90 {
		period = 90
	}
	r.parameters.Set(constants.ParameterNameWait, strconv.Itoa(period))
	return r
}
