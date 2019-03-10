package grpc

import (
	"github.com/chat/receiver"
	"golang.org/x/net/context"
)

func NewGrpcReceiver(
	receiver receiver.Receiver,
) (*GrpcReceiver, error) {
	return &GrpcReceiver{
		receiver: receiver,
	}, nil
}

type GrpcReceiver struct {
	receiver receiver.Receiver
}

func (client *GrpcReceiver) SendMessage(ctx context.Context, msg *ChatMessage) (*ReceiverReply, error) {

	message := &receiver.Message{
		Text:      msg.Text,
		AuthorId:  msg.AuthorId,
		Timestamp: msg.Timestamp,
	}

	err := client.receiver.Receive(message)

	if err != nil {
		return &ReceiverReply{Status: "fail"}, err
	}

	return &ReceiverReply{Status: "ok"}, nil
}
