package tests

import (
	"github.com/chat/receiver"
	"github.com/stretchr/testify/mock"
)

type SaverMock struct {
	mock.Mock
}

func (saver *SaverMock) Save(message *receiver.Message) error {
	args := saver.Called(message)
	return args.Error(0)
}
