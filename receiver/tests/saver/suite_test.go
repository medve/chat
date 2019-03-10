package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/chat/receiver"
	"github.com/chat/receiver/saver"
	"github.com/chat/receiver/tests/queue"
)

type TestSuite struct {
	suite.Suite
	Saver       receiver.Saver
	TestMessage *receiver.Message
	Queue       *tests.QueueMock
}

func (suite *TestSuite) SetupTest() {
	suite.Queue = &tests.QueueMock{}

	saver, err := saver.NewAsyncSaver(suite.Queue)

	assert.Nil(suite.T(), err, "There is an error in process of test suite setup")

	assert.NotNil(suite.T(), saver, "Saver should not be nil")

	suite.Saver = saver

	suite.TestMessage = &receiver.Message{
		Text:      "asd",
		AuthorId:  1,
		Timestamp: 123,
	}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
