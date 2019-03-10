package tests

import (
	"github.com/chat/receiver"
	"github.com/chat/receiver/notifications"
	"github.com/chat/receiver/tests/queue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	Notifier    receiver.Notifier
	TestMessage *receiver.Message
	Queue       *tests.QueueMock
}

func (suite *TestSuite) SetupTest() {
	suite.Queue = &tests.QueueMock{}

	notifier, err := notify.NewNotifier(suite.Queue)

	assert.Nil(suite.T(), err, "There is an error in process of test suite setup")

	assert.NotNil(suite.T(), notifier, "Notifier should not be nil")

	suite.Notifier = notifier

	suite.TestMessage = &receiver.Message{
		Text:      "asd",
		AuthorId:  1,
		Timestamp: 123,
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
