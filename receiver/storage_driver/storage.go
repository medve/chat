package storage

import "github.com/chat/receiver/message"

type StorageDriver interface {
	Save(message *message.Message) error
}
