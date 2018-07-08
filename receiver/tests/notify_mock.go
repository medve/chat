package receiver

import (
	"errors"

	"github.com/chat/receiver/message"
)

type NotifyMock struct {
	NotifyCalled bool
	WithError    bool
}

func (notify *NotifyMock) Notify(message *message.Message) error {
	notify.NotifyCalled = true
	if notify.WithError {
		return errors.New("Error occured")
	}
	return nil
}
