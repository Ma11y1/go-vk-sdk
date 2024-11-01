package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Likes Doc: https://dev.vk.com/ru/method/likes
type Likes struct {
	BaseAction
}

// Add Doc: https://dev.vk.com/ru/method/likes.add
func (l *Likes) Add(actor actor.Actor) *request.LikesAddRequest {
	return request.NewLikesAddRequest(l.api, actor)
}

// Delete Doc: https://dev.vk.com/ru/method/likes.delete
func (l *Likes) Delete(actor actor.Actor) *request.LikesDeleteRequest {
	return request.NewLikesDeleteRequest(l.api, actor)
}

// GetList Doc: https://dev.vk.com/ru/method/likes.getList
func (l *Likes) GetList(actor actor.Actor) *request.LikesGetListRequest {
	return request.NewLikesGetListRequest(l.api, actor)
}

// GetListExtended Doc: https://dev.vk.com/ru/method/likes.getList
func (l *Likes) GetListExtended(actor actor.Actor) *request.LikesGetListExtendedRequest {
	return request.NewLikesGetListExtendedRequest(l.api, actor)
}

// IsLiked Doc: https://dev.vk.com/ru/method/likes.isLiked
func (l *Likes) IsLiked(actor actor.Actor) *request.LikesIsLikedRequest {
	return request.NewLikesIsLikedRequest(l.api, actor)
}
