package main

import (
	"context"
	"go-vk-sdk/callback"
	"go-vk-sdk/events"
	"go-vk-sdk/logger"
	"time"
)

// TODO - Change all generated errors to errors.InternalError or call errors.Error() function
// TODO - Test all requests
// TODO - Make display of logger information using logger.Log()
// TODO - Add methods to query structures to set parameter values

func main() {
	logger.Enable()
	c := callback.NewCallback("name1", "localhost:8080", "/qq")
	err := c.Run()
	if err != nil {
		panic(err)
	}

	c.AddEventListener(events.EventTypeConfirmation, events.NewEventListener(ConfirmationHandler))

	for c.IsRunning() {
		t := time.NewTicker(2 * time.Second)
		<-t.C
		c.Stop(context.TODO())
	}
}

func ConfirmationHandler(event events.Event) {
	conf := event.(*events.EventConfirmation)
	err := conf.Confirm("code_confirm")
	if err != nil {
		panic(err)
	}
}
