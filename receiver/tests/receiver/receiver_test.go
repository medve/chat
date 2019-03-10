package tests

import (
	"errors"
)

func (suite *TestSuite) TestSendMessageErrorNotify() {
	suite.Notifier.On("Notify", suite.TestMessage).Return(errors.New("Some error"))
	suite.Saver.On("Save", suite.TestMessage).Return(nil)

	err := suite.Receiver.Receive(
		suite.TestMessage,
	)

	suite.NotNil(err, "Error should be got")
}

func (suite *TestSuite) TestSendMessageErrorSave() {
	suite.Notifier.On("Notify", suite.TestMessage).Return(nil)
	suite.Saver.On("Save", suite.TestMessage).Return(errors.New("Some error"))

	err := suite.Receiver.Receive(
		suite.TestMessage,
	)

	suite.NotNil(err, "Error should be got")
}

func (suite *TestSuite) TestSendMessageOk() {
	suite.Notifier.On("Notify", suite.TestMessage).Return(nil)
	suite.Saver.On("Save", suite.TestMessage).Return(nil)

	err := suite.Receiver.Receive(
		suite.TestMessage,
	)

	suite.Nil(err, "There is should not be any errors")

}
