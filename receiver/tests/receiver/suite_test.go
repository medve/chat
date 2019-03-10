package tests

import (
	"github.com/chat/receiver"
	"github.com/chat/receiver/tests/notifier"
	tests2 "github.com/chat/receiver/tests/saver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	Notifier    *tests.NotifyMock
	Saver       *tests2.SaverMock
	Receiver    receiver.Receiver
	TestMessage *receiver.Message
}

func (suite *TestSuite) SetupTest() {

	suite.Notifier = new(tests.NotifyMock)
	suite.Saver = new(tests2.SaverMock)

	receiverObj := receiver.NewSimpleReceiver(
		suite.Saver,
		suite.Notifier,
	)

	assert.NotNil(suite.T(), receiverObj, "Receiver should not be nil")

	suite.Receiver = receiverObj

	suite.TestMessage = &receiver.Message{
		Text:      "asd",
		AuthorId:  1,
		Timestamp: 123,
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
