package longPollGroup

import (
	"context"
	"errors"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	"go-vk-sdk/events"
	"go-vk-sdk/request"
	"sync/atomic"
)

// Doc: https://dev.vk.com/ru/api/bots-long-poll/getting-started

type Server struct {
	URL string `json:"url"`
	Key string `json:"key"`
	Ts  string `json:"ts"`
}

type EventUpdate struct {
	Type  events.EventType
	Event interface{}
}

type LongPoll struct {
	api           *api.API
	user          actor.Actor
	groupID       int
	url           string
	key           string
	ts            string
	wait          int
	trackedEvents map[events.EventType]int
	chanUpdate    chan *EventUpdate
	isRunning     int32

	req             *request.LongPollGroupRequest             // cache
	reqSetSettings  *request.GroupsSetLongPollSettingsRequest // cache
	reqUpdateServer *request.GroupsGetLongPollServerRequest   // cache
}

func NewLongPoll(a *api.API, user actor.Actor, groupID int) *LongPoll {
	return &LongPoll{
		api:             a,
		user:            user,
		groupID:         groupID,
		url:             "",
		key:             "",
		ts:              "",
		wait:            25,
		trackedEvents:   map[events.EventType]int{},
		chanUpdate:      make(chan *EventUpdate, 2),
		isRunning:       0,
		req:             request.NewLongPollGroupRequest(a, "").Wait(25),
		reqSetSettings:  request.NewGroupsSetLongPollSettingsRequest(a, user).GroupID(groupID).Enabled(true).APIVersion(a.Version),
		reqUpdateServer: request.NewGroupsGetLongPollServerRequest(a, user).GroupID(groupID),
	}
}

func NewLongPollServer(a *api.API, user actor.Actor, groupID int, server *Server) *LongPoll {
	lp := NewLongPoll(a, user, groupID)
	lp.url = server.URL
	lp.key = server.Key
	lp.ts = server.Ts
	return lp
}

func (l *LongPoll) UpdateServer(isUpdateTs bool) error {
	serverSettings, err := l.reqUpdateServer.Exec(context.Background())
	if err != nil {
		return err
	}

	if serverSettings.Response.Key == "" {
		return errors.New("LongPoll.UpdateServer(): response server settings is empty")
	}

	l.key = serverSettings.Response.Key
	l.url = serverSettings.Response.Server

	if isUpdateTs {
		l.ts = serverSettings.Response.Ts
	}

	l.req.
		Key(l.key).
		Ts(l.ts).
		SetURL(l.url)

	return nil
}

func (l *LongPoll) autoSetSettings() error {
	l.reqSetSettings.ResetEvents()

	for key, value := range l.trackedEvents {
		if value > 0 {
			l.reqSetSettings.SetEvent(string(key), true)
		}
	}

	_, err := l.reqSetSettings.Exec(context.Background())
	if err != nil {
		return fmt.Errorf("LongPoll.autoSetSettings(): APIError long poll group auto set settings: %v\n", err)
	}

	return nil
}

func (l *LongPoll) Run(ctx context.Context) error {
	if l.url == "" || l.key == "" {
		return errors.New("LongPoll.Run(): server is undefined")
	}

	if !atomic.CompareAndSwapInt32(&l.isRunning, 0, 1) {
		return fmt.Errorf("LongPoll.Run(): long poll group is already running")
	}

	defer atomic.StoreInt32(&l.isRunning, 0)

	err := l.autoSetSettings()
	if err != nil {
		return err
	}

	for atomic.LoadInt32(&l.isRunning) == 1 {
		select {
		case _, ok := <-ctx.Done():
			if !ok {
				return ctx.Err()
			}
		default:
			resp, err := l.req.Exec(ctx)
			if err != nil {
				return err // TODO complete error handling so as not to complete verification cycle
			}

			switch FailedType(resp.Failed) {
			case 0, FailedTypeOutdatedStory:
				l.ts = resp.Ts
			case FailedTypeExpiredKey:
				err = l.UpdateServer(false)
			case FailedTypeOutdatedUserInfo:
				err = l.UpdateServer(true)
			default:
				err = &FailedError{Code: resp.Failed}
			}
			if err != nil {
				return err // TODO complete error handling so as not to complete verification cycle
			}

			for _, update := range resp.Updates {
				event, err := events.NewEvent(&update)
				if err != nil {
					return err // TODO complete error handling so as not to complete verification cycle
				}

				l.chanUpdate <- &EventUpdate{
					Type:  update.Type,
					Event: event,
				}
			}
		}
	}

	return nil
}

func (l *LongPoll) Stop() error {
	atomic.StoreInt32(&l.isRunning, 0)
	return nil
}

func (l *LongPoll) TrackEvent(event events.EventType) {
	l.trackedEvents[event] = 1
}

func (l *LongPoll) UntrackEvent(event events.EventType) {
	delete(l.trackedEvents, event)
}

func (l *LongPoll) Updates() chan *EventUpdate {
	return l.chanUpdate
}

func (l *LongPoll) SetServer(server *Server) error {
	if server == nil || server.URL == "" || server.Key == "" {
		return fmt.Errorf("LongPoll.SetServer(): invalid server configuration %+v\n", server)
	}

	l.url = server.URL
	l.key = server.Key
	l.ts = server.Ts
	l.req.
		Key(l.key).
		Ts(l.ts).
		SetURL(l.url)

	return nil
}

// SetWait wait > 0 and < 90
func (l *LongPoll) SetWait(wait int) error {
	if wait <= 0 || wait > 90 {
		return fmt.Errorf("LongPoll.SetWait(): invalid wait time: %d, must be between 1 and 90", wait)
	}

	l.wait = wait
	l.req.Wait(wait)

	return nil
}

func (l *LongPoll) IsRunning() bool {
	return atomic.LoadInt32(&l.isRunning) == 1
}
