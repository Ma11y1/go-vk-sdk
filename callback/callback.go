package callback

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-vk-sdk/actor"
	"go-vk-sdk/api"
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/events"
	"go-vk-sdk/logger"
	"go-vk-sdk/objects"
	"go-vk-sdk/request"
	"go-vk-sdk/transport"
	"net/http"
	"net/url"
	"strconv"
)

// Doc: https://dev.vk.com/ru/api/callback/getting-started

type ServerStatus string

const (
	ServerStatusOk           ServerStatus = "ok"
	ServerStatusWait         ServerStatus = "wait"
	ServerStatusFailed       ServerStatus = "failed"
	ServerStatusUnconfigured ServerStatus = "unconfigured"
)

type Callback struct {
	server           transport.CallbackServer
	api              *api.API
	actor            actor.Actor
	eventEmitter     *events.EventEmitter[events.EventType, *events.EventCallback]
	Name             string
	SecretKey        string
	confirmationKeys map[int]string
	secretKeys       map[int]string
}

func NewCallback(api *api.API, actor actor.Actor, url *url.URL) *Callback {
	callback := &Callback{
		api:              api,
		actor:            actor,
		server:           transport.NewBaseCallbackServer(url),
		eventEmitter:     events.NewEventEmitter[events.EventType, *events.EventCallback](),
		Name:             "go-vk-sdk",
		confirmationKeys: make(map[int]string),
		secretKeys:       make(map[int]string),
	}

	err := callback.SetDefaultHandler(url.Path)
	if err != nil {
		panic(err)
	}

	return callback
}

func NewCallbackServer(api *api.API, actor actor.Actor, server transport.CallbackServer) *Callback {
	return &Callback{
		api:              api,
		actor:            actor,
		server:           server,
		eventEmitter:     events.NewEventEmitter[events.EventType, *events.EventCallback](),
		Name:             "go-vk-sdk",
		confirmationKeys: make(map[int]string),
		secretKeys:       make(map[int]string),
	}
}

// UpdateSettings Allows you to update the settings of the local server and VK server
//
//	It is advisable to call for initial initialization
func (c *Callback) UpdateSettings(ctx context.Context, url string) error {
	serverID := 0

	g, err := request.NewGroupsGetByIDRequest(c.api, c.actor).Exec(ctx)
	if err != nil || g.Error.Code != 0 {
		if errors.Is(err, internalErrors.ParamCode) {
			return internalErrors.ErrorLog("Callback.UpdateSettings()", "Error request groups, need group access token")
		}
		return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
	}

	groupID := g.Response.Groups[0].ID

	servers, err := c.GetServers(groupID)
	if err != nil {
		return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
	}

	// update servers and get serverID
	for _, server := range servers {
		if server.URL == url {
			if server.Status == string(ServerStatusOk) {
				serverID = server.ID
				c.secretKeys[groupID] = server.SecretKey
				break
			}

			isSuccessful, err := c.DeleteServer(server.ID, groupID)
			if err != nil {
				logger.Log("Callback.UpdateSettings()", fmt.Sprintf("Error delete callback server %d from group %d: %s", server.ID, groupID, err.Error()))
			}
			if isSuccessful {
				logger.Log("Callback.UpdateSettings()", fmt.Sprintf("Callback server %d was deleted from group %d", server.ID, groupID))
			}
		}
	}

	// create new server
	if serverID == 0 {
		secretKey, err := GenerateRandomString(24)
		if err != nil {
			return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
		}

		c.secretKeys[groupID] = secretKey

		confirmationKey, err := c.GetConfirmationKey(groupID)
		if err != nil {
			return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
		}

		c.confirmationKeys[groupID] = confirmationKey

		serverID, err = c.AddServer(groupID, c.Name, c.server.GetURL().String(), secretKey)
		if err != nil {
			return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
		}
	}

	_, err = c.SetSettings(groupID, serverID)
	if err != nil {
		return internalErrors.ErrorLog("Callback.UpdateSettings()", err.Error())
	}

	return nil
}

func (c *Callback) Run() error {
	err := c.server.Run()
	if err != nil {
		return internalErrors.ErrorLog("Callback.Run()", "Failed to start callback server: "+err.Error())
	}
	logger.Log("Callback.Run()", "Payments server is running at url: "+c.server.GetURL().String())
	return nil
}

func (c *Callback) Stop(ctx context.Context) error {
	err := c.server.Stop(ctx)
	logger.Log("Callback.Stop()", "Payments server is stopped at url: "+c.server.GetURL().String())
	if err != nil {
		return internalErrors.ErrorLog("Callback.Stop()", "Failed to stop callback server: "+err.Error())
	}
	return nil
}

// handle base handler at a given path
func (c *Callback) handle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var updateEvent events.EventUpdate
	err := decoder.Decode(&updateEvent)
	if err != nil {
		logger.Log("Callback.handle()", "error JSON decode result body: "+err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	secretKey, ok := c.secretKeys[updateEvent.GroupID]
	if !ok {
		secretKey = c.SecretKey
	}

	if secretKey != "" && updateEvent.Secret != secretKey {
		logger.Log("Callback.handle()", fmt.Sprintf("Bad secret key %s for grouod id %d", secretKey, updateEvent.GroupID))
		http.Error(w, "Bad secret", http.StatusForbidden)

		return
	}

	if updateEvent.Type == events.EventTypeConfirmation {
		key, exists := c.confirmationKeys[updateEvent.GroupID]
		if exists && key != "" {
			_, err = w.Write([]byte(key))
			if err != nil {
				logger.Log("Callback.handle()", "failed to write confirmation key to response writer: "+err.Error())
			}
		} else {
			c.eventEmitter.Emit(events.EventTypeConfirmation, &events.EventCallback{Event: &events.EventConfirmation{
				Response: w,
				GroupID:  updateEvent.GroupID,
			}})
		}
		return
	}

	event, err := events.NewEvent(&updateEvent)
	if err != nil {
		logger.Log("Callback.handle()", "failed to create event: "+err.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	callbackEvent := &events.EventCallback{Event: event}

	callbackEvent.RetryCounter, _ = strconv.Atoi(r.Header.Get("X-Retry-Counter"))

	c.eventEmitter.Emit(updateEvent.Type, callbackEvent)

	if callbackEvent.Error != nil {
		logger.Log("Callback.handle()", "failed to handle callback event: "+callbackEvent.Error.Error())
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if callbackEvent.IsRemove {
		_, _ = w.Write([]byte("remove"))
		logger.Log("Callback.handle()", fmt.Sprintf("group %d server was deleted", updateEvent.GroupID))
		return
	}

	if callbackEvent.Code != 0 {
		w.Header().Set("Retry-After", callbackEvent.Date.Format(http.TimeFormat))
		http.Error(w, http.StatusText(callbackEvent.Code), callbackEvent.Code)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (c *Callback) AddEventListener(event events.EventType, listener *events.EventListener[*events.EventCallback]) {
	if event == "" || listener == nil {
		logger.Log("Callback.AddEventListener()", "attempted to add nil event or listener")
		return
	}
	c.eventEmitter.On(event, listener)
}

func (c *Callback) RemoveEventListener(event events.EventType, listener *events.EventListener[*events.EventCallback]) {
	if event == "" || listener == nil {
		logger.Log("Callback.RemoveEventListener()", "attempted to remove nil event or listener")
		return
	}
	c.eventEmitter.Off(event, listener)
}

func (c *Callback) ClearEventListeners(event events.EventType) {
	c.eventEmitter.Clear(event)
}

func (c *Callback) SetDefaultHandler(path string) error {
	if path == "" {
		return internalErrors.ErrorLog("Callback.SetDefaultHandler()", "Invalid value handle path. Path is empty")
	}
	return c.SetHandleFunc(path, c.handle)
}

func (c *Callback) SetHandler(path string, handler http.Handler) error {
	if path == "" {
		return internalErrors.ErrorLog("Callback.SetHandler()", "Invalid value handle path. Path is empty")
	}
	return c.server.SetHandler(path, handler)
}

func (c *Callback) SetHandleFunc(path string, fn func(http.ResponseWriter, *http.Request)) error {
	if path == "" {
		return internalErrors.ErrorLog("Callback.SetHandleFunc()", "Invalid value handle path. Path is empty")
	}
	return c.server.SetHandleFunc(path, fn)
}

// GetServers Retrieves information about servers for the Callback API in the community
func (c *Callback) GetServers(groupID int) ([]objects.GroupCallbackServer, error) {
	res, err := request.NewGroupsGetCallbackServersRequest(c.api, c.actor).
		GroupID(groupID).
		Exec(context.Background())
	if err != nil {
		return nil, internalErrors.ErrorLog("Callback.GetCallbackServers()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.Items, nil
}

// AddServer Adds a server for the Callback API to the community
//
//	After successful execution, returns the identifier of the added server in the server_id (integer) field
func (c *Callback) AddServer(groupID int, title, url, secret string) (int, error) {
	res, err := request.NewGroupsAddCallbackServerRequest(c.api, c.actor).
		GroupID(groupID).
		Title(title).
		URL(url).
		SecretKey(secret).
		Exec(context.Background())

	if err != nil {
		return -1, internalErrors.ErrorLog("Callback.AddServer()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.ServerID, nil
}

// DeleteServer Removes the server for the Callback API from the community
func (c *Callback) DeleteServer(serverID, groupID int) (bool, error) {
	res, err := request.NewGroupsDeleteCallbackServerRequest(c.api, c.actor).
		GroupID(groupID).
		ServerID(serverID).
		Exec(context.Background())

	if err != nil {
		return false, internalErrors.ErrorLog("Callback.DeleteServer()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	if res.Response == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

// SetSettings Allows you to set event notification settings in the Callback API
func (c *Callback) SetSettings(groupID, serverID int) (bool, error) {
	return c.SetSettingsEvents(groupID, serverID, c.eventEmitter.Keys())
}

// SetSettingsEvents Allows you to set event notification settings in the Callback API
func (c *Callback) SetSettingsEvents(groupID, serverID int, e []events.EventType) (bool, error) {
	req := request.NewGroupsSetCallbackSettingsRequest(c.api, c.actor).
		GroupID(groupID).
		ServerID(serverID).
		APIVersion(c.api.Version)

	for _, event := range e {
		req.SetEvent(string(event), true)
	}

	res, err := req.Exec(context.Background())
	if err != nil {
		return false, internalErrors.ErrorLog("Callback.SetSettings()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return true, nil
}

// GetConfirmationKey Allows you to get the string required to confirm the server address in the Callback API
func (c *Callback) GetConfirmationKey(groupID int) (string, error) {
	res, err := request.NewGroupsGetCallbackConfirmationCodeRequest(c.api, c.actor).
		GroupID(groupID).
		Exec(context.Background())

	if err != nil {
		return "", internalErrors.ErrorLog("Callback.GetConfirmationKey()", "Request error: "+err.Error()+"\nResponse error: "+res.Error.Error())
	}

	return res.Response.Code, nil
}

func (c *Callback) IsRunning() bool {
	return c.server.IsRunning()
}
