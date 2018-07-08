package receiver

import (
	"errors"

	"github.com/chat/receiver/message"
)

type StorageMock struct {
	SaveCalled bool
	WithError  bool
}

func (storage *StorageMock) Save(message *message.Message) error {
	storage.SaveCalled = true

	if storage.WithError {
		return errors.New("Error occured")
	}
	return nil
}
