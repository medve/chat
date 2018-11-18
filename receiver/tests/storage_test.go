package receiver

import (
	"testing"

	storage "github.com/chat/receiver/storage_driver"
)

func TestCreateStorage(t *testing.T) {
	queueDriver := &QueueMock{}

	storageDriver, err := storage.NewStorage(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if storageDriver == nil {
		t.Error("Storage driver should not be nil")
	}
}

func TestStorageSave(t *testing.T) {
	queueDriver := &QueueMock{IsAdded: false}

	storageDriver, err := storage.NewStorage(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if storageDriver == nil {
		t.Error("Storage driver should not be nil")
	}

	err = storageDriver.Save(getTestMessage())

	if err != nil {
		t.Error("Fail", err)
	}

	if !queueDriver.IsAdded {
		t.Error("Storage should add message to queue")
	}

}

func TestStorageSaveError(t *testing.T) {
	queueDriver := &QueueMock{IsAdded: false, WithError: true}

	storageDriver, err := storage.NewStorage(queueDriver)

	if err != nil {
		t.Error("Fail", err)
	}

	if storageDriver == nil {
		t.Error("Storage driver should not be nil")
	}

	err = storageDriver.Save(getTestMessage())

	if err == nil {
		t.Error("Error should be occured")
	}

}
