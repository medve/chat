package receiver

import (
	"github.com/chat/receiver/message"
)

type Receiver interface {
	Save(message *message.Message) error
	Notify(message *message.Message) error
}
