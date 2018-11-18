package receiver

import (
	"testing"

	notify "github.com/chat/receiver/notify_driver"
)

func TestCreateNotifier(t *testing.T) {
	queueDriver := &QueueMock{}

	notifyDriver, err := notify.NewNotifier(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if notifyDriver == nil {
		t.Error("Notify driver should not be nil")
	}
}

func TestNotifierNotify(t *testing.T) {
	queueDriver := &QueueMock{IsAdded: false}

	notifierDriver, err := notify.NewNotifier(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if notifierDriver == nil {
		t.Error("Storage driver should not be nil")
	}

	err = notifierDriver.Notify(getTestMessage())

	if err != nil {
		t.Error("Fail", err)
	}

	if !queueDriver.IsAdded {
		t.Error("Storage should add message to queue")
	}

}

func TestNotifierNotifyError(t *testing.T) {
	queueDriver := &QueueMock{IsAdded: false, WithError: true}

	notifierDriver, err := notify.NewNotifier(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if notifierDriver == nil {
		t.Error("Notifier driver should not be nil")
	}

	err = notifierDriver.Notify(getTestMessage())

	if err == nil {
		t.Error("Error should be occured")
	}

}
