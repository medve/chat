package receiver

import (
	"errors"

	"github.com/chat/receiver/message"
)

type QueueMock struct {
	IsAdded   bool
	WithError bool
}

func (queue *QueueMock) Add(message *message.Message) error {
	queue.IsAdded = true
	if queue.WithError {
		return errors.New("Error")
	}
	return nil
}

func (queue *QueueMock) Receive() (*message.Message, error) {

	return nil, nil
}

func (queue *QueueMock) GetSubscribeChan() (chan *message.Message, error) {

	return nil, nil
}
