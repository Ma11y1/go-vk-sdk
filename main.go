package main

import (
	internalErrors "go-vk-sdk/errors"
	"go-vk-sdk/logger"
	"go-vk-sdk/payments"
	"net/url"
	"time"
)

// TODO - Change all generated errors to errors.InternalError or call errors.Error() function
// TODO - Test all requests
// TODO - Make display of logger information using logger.Log()
// TODO - Add methods to query structures to set parameter values

func main() {
	logger.Enable()

	u, err := url.Parse("http://localhost:8080/payments")
	if err != nil {
		panic(err)
	}
	c := payments.NewCallback(u, "SECRET")

	err = c.Run()
	if err != nil {
		panic(err)
	}

	c.OnGetItem(func(e *payments.GetItemRequest) (*payments.GetItemResponse, *internalErrors.PaymentsError) {
		resp := &payments.GetItemResponse{}

		return resp, nil
	}, false)

	for c.IsRunning() {
		t := time.NewTicker(2 * time.Second)
		<-t.C
	}
}
