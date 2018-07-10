package receiver

import (
	"github.com/chat/receiver/message"
	notify "github.com/chat/receiver/notify_driver"
	storage "github.com/chat/receiver/storage_driver"
	context "golang.org/x/net/context"
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

func createMessage(msg *ChatMessage) *message.Message {
	return &message.Message{
		Text:      msg.Text,
		AuthorId:  msg.AuthorId,
		Timestamp: msg.Timestamp,
	}
}

func (receiver *GrpcReceiver) SendMessage(ctx context.Context, msg *ChatMessage) (*ReceiverReply, error) {

	message := createMessage(msg)

	err := receiver.Save(message)

	if err != nil {
		return nil, err
	}

	err = receiver.Notify(message)

	if err != nil {
		return nil, err
	}

	return &ReceiverReply{Status: "ok"}, nil
}
