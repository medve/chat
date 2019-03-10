package notify

import (
	"github.com/chat/receiver"
)

func NewNotifier(queue receiver.Queue) (*AsyncNotifierDriver, error) {
	return &AsyncNotifierDriver{
		Queue: queue,
	}, nil
}

type AsyncNotifierDriver struct {
	Queue receiver.Queue
}

func (notifier *AsyncNotifierDriver) Notify(message *receiver.Message) error {
	err := notifier.Queue.Add(message)
	return err
}
