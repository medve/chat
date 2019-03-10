package tests

import (
	"github.com/chat/receiver/saver/db"
)

func (suite *DBSuite) TestSaveMessage() {
	tx := suite.GetConnection("TestSaveMessage")

	suite.NotNil(tx, "Error get suite")

	driver := db.NewPostgresDbDriver(tx)

	err := driver.TestFunc()

	if err != nil {
		suite.Error(err, "Error in query")
	}

}
