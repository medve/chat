package tests

import (
	"github.com/chat/receiver"
	"github.com/stretchr/testify/mock"
)

type QueueMock struct {
	mock.Mock
}

func (queue *QueueMock) Add(message *receiver.Message) error {

	args := queue.Called(message)
	return args.Error(0)
}

func (queue *QueueMock) Receive() (*receiver.Message, error) {

	args := queue.Called()
	return args.Get(0).(*receiver.Message), args.Error(1)
}

func (queue *QueueMock) GetSubscribeChan() (chan *receiver.Message, error) {

	args := queue.Called()
	return args.Get(0).(chan *receiver.Message), args.Error(1)
}
