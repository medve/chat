package receiver

import (
	"testing"

	"github.com/chat/receiver"
)

func TestCreateReceiver(t *testing.T) {
	notifyMock := NotifyMock{}
	storageMock := StorageMock{}

	receiver, err := receiver.NewGrpcReceiver(
		&notifyMock,
		&storageMock,
	)

	if err != nil {
		t.Error("Fail", err)
	}

	if receiver == nil {
		t.Error("Receiver shuld not be nil")
	}
}
