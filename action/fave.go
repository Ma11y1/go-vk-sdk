package action

import (
	"go-vk-sdk/actor"
	"go-vk-sdk/request"
)

// Fave Doc: https://dev.vk.com/ru/method/fave
type Fave struct {
	BaseAction
}

// AddArticle Doc: https://dev.vk.com/ru/method/fave.addArticle
func (f *Fave) AddArticle(actor actor.Actor) *request.FaveAddArticleRequest {
	return request.NewFaveAddArticleRequest(f.api, actor)
}

// AddLink Doc: https://dev.vk.com/ru/method/fave.addLink
func (f *Fave) AddLink(actor actor.Actor) *request.FaveAddLinkRequest {
	return request.NewFaveAddLinkRequest(f.api, actor)
}

// AddPage Doc: https://dev.vk.com/ru/method/fave.addPage
func (f *Fave) AddPage(actor actor.Actor) *request.FaveAddPageRequest {
	return request.NewFaveAddPageRequest(f.api, actor)
}

// AddPost Doc: https://dev.vk.com/ru/method/fave.addPost
func (f *Fave) AddPost(actor actor.Actor) *request.FaveAddPostRequest {
	return request.NewFaveAddPostRequest(f.api, actor)
}

// AddProduct Doc: https://dev.vk.com/ru/method/fave.addProduct
func (f *Fave) AddProduct(actor actor.Actor) *request.FaveAddProductRequest {
	return request.NewFaveAddProductRequest(f.api, actor)
}

// AddTag Doc: https://dev.vk.com/ru/method/fave.addTag
func (f *Fave) AddTag(actor actor.Actor) *request.FaveAddTagRequest {
	return request.NewFaveAddTagRequest(f.api, actor)
}

// AddVideo Doc: https://dev.vk.com/ru/method/fave.addVideo
func (f *Fave) AddVideo(actor actor.Actor) *request.FaveAddVideoRequest {
	return request.NewFaveAddVideoRequest(f.api, actor)
}

// EditTag Doc: https://dev.vk.com/ru/method/fave.editTag
func (f *Fave) EditTag(actor actor.Actor) *request.FaveEditTagRequest {
	return request.NewFaveEditTagRequest(f.api, actor)
}

// Get Doc: https://dev.vk.com/ru/method/fave.get
func (f *Fave) Get(actor actor.Actor) *request.FaveGetRequest {
	return request.NewFaveGetRequest(f.api, actor)
}

// GetExtended Doc: https://dev.vk.com/ru/method/fave.get
func (f *Fave) GetExtended(actor actor.Actor) *request.FaveGetExtendedRequest {
	return request.NewFaveGetExtendedRequest(f.api, actor)
}

// GetPages Doc: https://dev.vk.com/ru/method/fave.getPages
func (f *Fave) GetPages(actor actor.Actor) *request.FaveGetPagesRequest {
	return request.NewFaveGetPagesRequest(f.api, actor)
}

// GetTags Doc: https://dev.vk.com/ru/method/fave.getTags
func (f *Fave) GetTags(actor actor.Actor) *request.FaveGetTagsRequest {
	return request.NewFaveGetTagsRequest(f.api, actor)
}

// MarkSeen Doc: https://dev.vk.com/ru/method/fave.markSeen
func (f *Fave) MarkSeen(actor actor.Actor) *request.FaveMarkSeenRequest {
	return request.NewFaveMarkSeenRequest(f.api, actor)
}

// RemoveArticle Doc: https://dev.vk.com/ru/method/fave.removeArticle
func (f *Fave) RemoveArticle(actor actor.Actor) *request.FaveRemoveArticleRequest {
	return request.NewFaveRemoveArticleRequest(f.api, actor)
}

// RemoveLink Doc: https://dev.vk.com/ru/method/fave.removeLink
func (f *Fave) RemoveLink(actor actor.Actor) *request.FaveRemoveLinkRequest {
	return request.NewFaveRemoveLinkRequest(f.api, actor)
}

// RemovePage Doc: https://dev.vk.com/ru/method/fave.removePage
func (f *Fave) RemovePage(actor actor.Actor) *request.FaveRemovePageRequest {
	return request.NewFaveRemovePageRequest(f.api, actor)
}

// RemovePost Doc: https://dev.vk.com/ru/method/fave.removePost
func (f *Fave) RemovePost(actor actor.Actor) *request.FaveRemovePostRequest {
	return request.NewFaveRemovePostRequest(f.api, actor)
}

// RemoveProduct Doc: https://dev.vk.com/ru/method/fave.removeProduct
func (f *Fave) RemoveProduct(actor actor.Actor) *request.FaveRemoveProductRequest {
	return request.NewFaveRemoveProductRequest(f.api, actor)
}

// RemoveTag Doc: https://dev.vk.com/ru/method/fave.removeTag
func (f *Fave) RemoveTag(actor actor.Actor) *request.FaveRemoveTagRequest {
	return request.NewFaveRemoveTagRequest(f.api, actor)
}

// RemoveVideo Doc: https://dev.vk.com/ru/method/fave.removeVideo
func (f *Fave) RemoveVideo(actor actor.Actor) *request.FaveRemoveVideoRequest {
	return request.NewFaveRemoveVideoRequest(f.api, actor)
}

// ReorderTags Doc: https://dev.vk.com/ru/method/fave.reorderTags
func (f *Fave) ReorderTags(actor actor.Actor) *request.FaveReorderTagsRequest {
	return request.NewFaveReorderTagsRequest(f.api, actor)
}

// SetPageTags Doc: https://dev.vk.com/ru/method/fave.setPageTags
func (f *Fave) SetPageTags(actor actor.Actor) *request.FaveSetPageTagsRequest {
	return request.NewFaveSetPageTagsRequest(f.api, actor)
}

// SetTags Doc: https://dev.vk.com/ru/method/fave.setTags
func (f *Fave) SetTags(actor actor.Actor) *request.FaveSetTagsRequest {
	return request.NewFaveSetTagsRequest(f.api, actor)
}

// TrackPageInteraction Doc: https://dev.vk.com/ru/method/fave.trackPageInteraction
func (f *Fave) TrackPageInteraction(actor actor.Actor) *request.FaveTrackPageInteractionRequest {
	return request.NewFaveTrackPageInteractionRequest(f.api, actor)
}
