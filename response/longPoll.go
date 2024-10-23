package response

import "go-vk-sdk/events"

// LongPollUserResponse
//
//	Doc: https://dev.vk.com/ru/api/user-long-poll/getting-started
type LongPollUserResponse struct {
	Ts int `json:"ts"` // Last event number. Use it in your next query.

	// An array whose elements contain a representation of new events (each element is also an array).
	// The length of the updates array can be 0 (this means that no new events occurred during the wait).

	Updates [][]interface{} `json:"updates"`
	// "failed":1
	// The event history is out of date or partially lost, the application can receive further events using the new ts value from the response.
	// "failed":2
	// The key has expired, you need to get the key again using the messages.getLongPollServer method.
	// "failed":3
	// Information about the user has been lost, you need to request a new key and ts using the messages.getLongPollServer method.
	// "failed":4
	// An invalid version number was passed in the version parameter.
	Failed int `json:"failed"`
}

// LongPollGroupResponse
//
//	Doc: https://dev.vk.com/ru/api/bots-long-poll/getting-started
type LongPollGroupResponse struct {
	Ts      string               `json:"ts"` // Last event number. Use it in your next query.
	Updates []events.EventUpdate `json:"updates"`
	// "failed":1 - the event history is out of date or was partially lost, the application can receive further events using the new ts value from the response.
	// "failed":2 - the key has expired, you need to get the key again using the groups.getLongPollServer method.
	// "failed":3 - information is lost, you need to request a new key and ts using the groups.getLongPollServer method.
	Failed int `json:"failed"`
}
