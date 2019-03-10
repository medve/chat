package tests

import (
	"github.com/chat/receiver"
	"github.com/stretchr/testify/mock"
)

type NotifyMock struct {
	mock.Mock
}

func (notify *NotifyMock) Notify(message *receiver.Message) error {
	args := notify.Called(message)
	return args.Error(0)
}
