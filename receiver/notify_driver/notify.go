package notify

import "github.com/chat/receiver/message"

type NotifyDriver interface {
	Notify(message *message.Message) error
}
