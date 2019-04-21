package saver

import "github.com/chat/receiver"

type DbDriver interface {
	SaveMessage(message *receiver.Message) error
}
