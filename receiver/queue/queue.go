package queue

import "github.com/chat/receiver/message"

type Queue interface {
	Add(message *message.Message) error
	Receive() (*message.Message, error)
	GetSubscribeChan() (chan *message.Message, error)
}
