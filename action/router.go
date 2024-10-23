package action

import (
	"errors"
	"go-vk-sdk/api"
)

type Router struct {
	api             *api.API
	Account         *Account
	Ads             *Ads
	Apps            *Apps
	AppWidgets      *AppWidgets
	Asr             *Asr
	Auth            *Auth
	Board           *Board
	Calls           *Calls
	Database        *Database
	Docs            *Docs
	Donut           *Donut
	DownloadedGames *DownloadedGames
	Fave            *Fave
	Friends         *Friends
	Gifts           *Gifts
	Groups          *Groups
	LeadForms       *LeadForms
	Likes           *Likes
	Market          *Market
	Marusia         *Marusia
	Messages        *Messages
	NewsFeed        *NewsFeed
	Notes           *Notes
	Notifications   *Notifications
	Orders          *Orders
	Pages           *Pages
	Photos          *Photos
	Podcasts        *Podcasts
	Polls           *Polls
	PrettyCards     *PrettyCards
	Search          *Search
	Secure          *Secure
	Stats           *Stats
	Status          *Status
	Storage         *Storage
	Store           *Store
	Stories         *Stories
	Streaming       *Streaming
	Translations    *Translations
	Users           *Users
	Utils           *Utils
	Video           *Video
	Wall            *Wall
	Widgets         *Widgets
	custom          map[string]Action
}

func NewRouter(api *api.API) *Router {
	return &Router{
		api:             api,
		Account:         &Account{BaseAction: BaseAction{api}},
		Auth:            &Auth{BaseAction: BaseAction{api: api}},
		Ads:             &Ads{BaseAction: BaseAction{api: api}},
		AppWidgets:      &AppWidgets{BaseAction: BaseAction{api: api}},
		Apps:            &Apps{BaseAction: BaseAction{api: api}},
		Asr:             &Asr{BaseAction: BaseAction{api: api}},
		Board:           &Board{BaseAction: BaseAction{api: api}},
		Calls:           &Calls{BaseAction: BaseAction{api: api}},
		Database:        &Database{BaseAction: BaseAction{api: api}},
		Docs:            &Docs{BaseAction: BaseAction{api: api}},
		Donut:           &Donut{BaseAction: BaseAction{api: api}},
		DownloadedGames: &DownloadedGames{BaseAction: BaseAction{api: api}},
		Fave:            &Fave{BaseAction: BaseAction{api: api}},
		Friends:         &Friends{BaseAction: BaseAction{api: api}},
		Gifts:           &Gifts{BaseAction: BaseAction{api: api}},
		Groups:          &Groups{BaseAction: BaseAction{api: api}},
		LeadForms:       &LeadForms{BaseAction: BaseAction{api: api}},
		Likes:           &Likes{BaseAction: BaseAction{api: api}},
		Market:          &Market{BaseAction: BaseAction{api: api}},
		Marusia:         &Marusia{BaseAction: BaseAction{api: api}},
		Messages:        &Messages{BaseAction: BaseAction{api: api}},
		NewsFeed:        &NewsFeed{BaseAction: BaseAction{api: api}},
		Notes:           &Notes{BaseAction: BaseAction{api: api}},
		Notifications:   &Notifications{BaseAction: BaseAction{api: api}},
		Orders:          &Orders{BaseAction: BaseAction{api: api}},
		Pages:           &Pages{BaseAction: BaseAction{api: api}},
		Photos:          &Photos{BaseAction: BaseAction{api: api}},
		Podcasts:        &Podcasts{BaseAction: BaseAction{api: api}},
		Polls:           &Polls{BaseAction: BaseAction{api: api}},
		PrettyCards:     &PrettyCards{BaseAction: BaseAction{api: api}},
		Search:          &Search{BaseAction: BaseAction{api: api}},
		Secure:          &Secure{BaseAction: BaseAction{api: api}},
		Stats:           &Stats{BaseAction: BaseAction{api: api}},
		Status:          &Status{BaseAction: BaseAction{api: api}},
		Storage:         &Storage{BaseAction: BaseAction{api: api}},
		Store:           &Store{BaseAction: BaseAction{api: api}},
		Stories:         &Stories{BaseAction: BaseAction{api: api}},
		Streaming:       &Streaming{BaseAction: BaseAction{api: api}},
		Translations:    &Translations{BaseAction: BaseAction{api: api}},
		Users:           &Users{BaseAction: BaseAction{api: api}},
		Utils:           &Utils{BaseAction: BaseAction{api: api}},
		Video:           &Video{BaseAction: BaseAction{api: api}},
		Wall:            &Wall{BaseAction: BaseAction{api: api}},
		Widgets:         &Widgets{BaseAction: BaseAction{api: api}},
		custom:          make(map[string]Action),
	}
}

func (r *Router) SetAPI(a *api.API) {
	r.api = a
	r.Account.SetAPI(a)
	r.Auth.SetAPI(a)
	r.Ads.SetAPI(a)
	r.AppWidgets.SetAPI(a)
	r.Apps.SetAPI(a)
	r.Asr.SetAPI(a)
	r.Board.SetAPI(a)
	r.Calls.SetAPI(a)
	r.Database.SetAPI(a)
	r.Docs.SetAPI(a)
	r.Donut.SetAPI(a)
	r.DownloadedGames.SetAPI(a)
	r.Fave.SetAPI(a)
	r.Friends.SetAPI(a)
	r.Gifts.SetAPI(a)
	r.Groups.SetAPI(a)
	r.LeadForms.SetAPI(a)
	r.Likes.SetAPI(a)
	r.Market.SetAPI(a)
	r.Marusia.SetAPI(a)
	r.Messages.SetAPI(a)
	r.NewsFeed.SetAPI(a)
	r.Notes.SetAPI(a)
	r.Notifications.SetAPI(a)
	r.Orders.SetAPI(a)
	r.Pages.SetAPI(a)
	r.Photos.SetAPI(a)
	r.Podcasts.SetAPI(a)
	r.Polls.SetAPI(a)
	r.PrettyCards.SetAPI(a)
	r.Search.SetAPI(a)
	r.Secure.SetAPI(a)
	r.Stats.SetAPI(a)
	r.Status.SetAPI(a)
	r.Storage.SetAPI(a)
	r.Store.SetAPI(a)
	r.Stories.SetAPI(a)
	r.Streaming.SetAPI(a)
	r.Translations.SetAPI(a)
	r.Users.SetAPI(a)
	r.Utils.SetAPI(a)
	r.Video.SetAPI(a)
	r.Wall.SetAPI(a)
	r.Widgets.SetAPI(a)

	for key := range r.custom {
		r.custom[key].SetAPI(a)
	}
}

func (r *Router) SetCustom(name string, action Action) error {
	if name == "" {
		return errors.New("Router.SetCustom(): name is empty")
	}

	if action == nil {
		return errors.New("Router.SetCustom(): action is nil")
	}

	r.custom[name] = action

	return nil
}
