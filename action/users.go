package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Users Doc: https://dev.vk.com/ru/method/search
type Users struct {
	BaseAction
}

// Get Doc: https://dev.vk.com/method/users.get
func (u *Users) Get(actor actor.Actor) *request.UsersGetRequest {
	return request.NewUsersGetRequest(u.api, actor)
}

// GetFollowers Doc: https://dev.vk.com/method/users.getFollowers
func (u *Users) GetFollowers(actor actor.Actor) *request.UsersGetFollowersRequest {
	return request.NewUsersGetFollowersRequest(u.api, actor)
}

// GetSubscriptions Doc: https://dev.vk.com/method/users.getSubscriptions
func (u *Users) GetSubscriptions(actor actor.Actor) *request.UsersGetSubscriptionsRequest {
	return request.NewUsersGetSubscriptionsRequest(u.api, actor)
}

// Report Doc: https://dev.vk.com/method/users.report
func (u *Users) Report(actor actor.Actor) *request.UsersReportRequest {
	return request.NewUsersReportRequest(u.api, actor)
}

// Search Doc: https://dev.vk.com/method/users.search
func (u *Users) Search(actor actor.Actor) *request.UsersSearchRequest {
	return request.NewUsersSearchRequest(u.api, actor)
}
