package receiver

import (
	"testing"

	"github.com/chat/receiver"
)

func TestCreateReceiver(t *testing.T) {
	notifyMock := NotifyMock{}
	storageMock := StorageMock{}

	receiverObj, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	if receiverObj == nil {
		t.Error("Receiver shuld not be nil")
	}
}

func TestSendMessageErrorNotify(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: true}
	storageMock := StorageMock{SaveCalled: false, WithError: false}

	receiverObj, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	_, err = receiverObj.SendMessage(
		nil,
		&receiver.ChatMessage{Text: "hello", AuthorId: 1, Timestamp: 123},
	)

	if err == nil {
		t.Error("Error should be occured")
	}
}

func TestSendMessageErrorStorage(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: false}
	storageMock := StorageMock{SaveCalled: false, WithError: true}

	receiverObj, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	_, err = receiverObj.SendMessage(
		nil,
		&receiver.ChatMessage{Text: "hello", AuthorId: 1, Timestamp: 123},
	)

	if err == nil {
		t.Error("Error should be occured")
	}
}

func TestSendMessageError(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: false}
	storageMock := StorageMock{SaveCalled: false, WithError: false}

	receiverObj, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	_, err = receiverObj.SendMessage(
		nil,
		&receiver.ChatMessage{Text: "hello", AuthorId: 1, Timestamp: 123},
	)

	if err != nil {
		t.Error("Error should not be occured")
	}

	if !storageMock.SaveCalled {
		t.Error("Receiver should save message")
	}

	if !notifyMock.NotifyCalled {
		t.Error("Receiver should save message")
	}

}
