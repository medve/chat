package saver

import (
	"github.com/chat/receiver"
)

func NewAsyncSaver(queue receiver.Queue) (*AsyncSaver, error) {
	return &AsyncSaver{
		Queue: queue,
	}, nil
}

type AsyncSaver struct {
	Queue receiver.Queue
}

func (storage *AsyncSaver) Save(message *receiver.Message) error {
	err := storage.Queue.Add(message)
	return err
}
