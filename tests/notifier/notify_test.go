package tests

import (
	"errors"
)

func (suite *TestSuite) TestNotifierNotify() {

	suite.Queue.On("Add", suite.TestMessage).Return(nil)

	err := suite.Notifier.Notify(suite.TestMessage)

	suite.Nil(err, "Error should be nil")

	suite.Queue.AssertExpectations(suite.T())

}

func (suite *TestSuite) TestNotifierNotifyError() {

	suite.Queue.On("Add", suite.TestMessage).Return(errors.New("some error"))

	err := suite.Notifier.Notify(suite.TestMessage)

	suite.NotNil(err, "There should be error")
	suite.Queue.AssertExpectations(suite.T())

}
