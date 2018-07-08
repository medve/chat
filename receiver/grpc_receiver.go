package receiver

import (
	"github.com/chat/receiver/message"
	notify "github.com/chat/receiver/notify_driver"
	storage "github.com/chat/receiver/storage_driver"
)

func NewGrpcReceiver(
	notifier notify.NotifyDriver,
	storage storage.StorageDriver,
) (*GrpcReceiver, error) {
	return &GrpcReceiver{
		StorageDriver: storage,
		NotifyDriver:  notifier,
	}, nil
}

type GrpcReceiver struct {
	StorageDriver storage.StorageDriver
	NotifyDriver  notify.NotifyDriver
}

func (receiver *GrpcReceiver) Save(message *message.Message) error {

	err := receiver.StorageDriver.Save(message)

	return err
}

func (receiver *GrpcReceiver) Notify(message *message.Message) error {

	err := receiver.NotifyDriver.Notify(message)

	return err
}
