package receiver

import (
	"testing"

	"github.com/chat/receiver"
	"github.com/chat/receiver/message"
)

func getTestMessage() *message.Message {
	return &message.Message{
		Text:      "asd",
		AuthorId:  1,
		Timestamp: 123,
	}
}

func TestSave(t *testing.T) {
	notifyMock := NotifyMock{}
	storageMock := StorageMock{SaveCalled: false, WithError: false}

	receiver, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	receiver.Save(getTestMessage())

	if err != nil {
		t.Error("Fail", err)
	}

	if !storageMock.SaveCalled {
		t.Error("Receiver should save message")
	}
}

func TestSaveError(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: false}
	storageMock := StorageMock{SaveCalled: false, WithError: true}

	receiver, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	err = receiver.Save(getTestMessage())

	if !storageMock.SaveCalled {
		t.Error("Receiver should save message")
	}

	if err == nil {
		t.Error("Error should be occured")
	}
}

func TestNotify(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: false}
	storageMock := StorageMock{SaveCalled: false, WithError: false}

	receiver, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	receiver.Notify(getTestMessage())

	if err != nil {
		t.Error("Fail", err)
	}

	if !notifyMock.NotifyCalled {
		t.Error("Receiver should notify message")
	}
}

func TestNotifyError(t *testing.T) {
	notifyMock := NotifyMock{NotifyCalled: false, WithError: true}
	storageMock := StorageMock{SaveCalled: false, WithError: false}

	receiver, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	err = receiver.Notify(getTestMessage())

	if !notifyMock.NotifyCalled {
		t.Error("Receiver should save message")
	}

	if err == nil {
		t.Error("Error should be occured")
	}
}
