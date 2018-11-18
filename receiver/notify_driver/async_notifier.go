package notify

import (
	"github.com/chat/receiver/message"
	"github.com/chat/receiver/queue"
)

func NewNotifier(queue queue.Queue) (*AsyncNotifierDriver, error) {
	return &AsyncNotifierDriver{
		Queue: queue,
	}, nil
}

type AsyncNotifierDriver struct {
	Queue queue.Queue
}

func (notifier *AsyncNotifierDriver) Notify(message *message.Message) error {
	err := notifier.Queue.Add(message)
	return err
}
