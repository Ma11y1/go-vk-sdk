package longPollGroup

import (
	"fmt"
	internalErrors "go-vk-sdk/errors"
)

type FailedError struct {
	Code int
}

func (e *FailedError) Error() string {
	return fmt.Sprintf("%s Long poll: failed code %d", internalErrors.MessagePrefix, e.Code)
}
