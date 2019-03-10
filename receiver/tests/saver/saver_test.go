package tests

import (
	"errors"
)

func (suite *TestSuite) TestStorageSave() {

	suite.Queue.On("Add", suite.TestMessage).Return(nil)

	err := suite.Saver.Save(suite.TestMessage)

	suite.Nil(err, "Error should be nil")

	suite.Queue.AssertExpectations(suite.T())

}

func (suite *TestSuite) TestStorageSaveError() {

	suite.Queue.On("Add", suite.TestMessage).Return(errors.New("some error"))

	err := suite.Saver.Save(suite.TestMessage)

	suite.NotNil(err, "There should be error")
	suite.Queue.AssertExpectations(suite.T())

}
