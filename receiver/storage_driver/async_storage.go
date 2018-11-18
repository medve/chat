package storage

import (
	"github.com/chat/receiver/message"
	"github.com/chat/receiver/queue"
)

func NewStorage(queue queue.Queue) (*AsyncStorageDriver, error) {
	return &AsyncStorageDriver{
		Queue: queue,
	}, nil
}

type AsyncStorageDriver struct {
	Queue queue.Queue
}

func (storage *AsyncStorageDriver) Save(message *message.Message) error {
	err := storage.Queue.Add(message)
	return err
}
