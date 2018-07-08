package queue

import "github.com/chat/receiver/message"

type Queue interface {
	Add(message *message.Message) error
}
