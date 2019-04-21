package tests

import (
	"github.com/chat/db"
	"github.com/chat/receiver"
)

func (suite *DBSuite) TestTxSuite() {
	tx := suite.GetConnection("TestTxSuite")

	suite.NotNil(tx, "Error get suite")

	driver := db.NewPostgresDbDriver(tx)

	err := driver.TestFunc()

	if err != nil {
		suite.Error(err, "Error in query")
	}

}

func (suite *DBSuite) TestSaveMessage() {
	tx := suite.GetConnection("TestSaveMessage")

	suite.NotNil(tx, "Error get suite")

	driver := db.NewPostgresDbDriver(tx)

	sourceMessage := &receiver.Message{
		Text:      "test",
		AuthorId:  1,
		Timestamp: 123,
	}

	err := driver.SaveMessage(sourceMessage)

	if err != nil {
		suite.Error(err, "Error in query")
	}

	var messages []db.Message
	err = suite.db.Model(&messages).Select()
	if err != nil {
		suite.Error(err, "Error while select messages")
	}

	suite.Len(messages, 1)
	messageFromDb := messages[0]
	suite.Equal(messageFromDb.AuthorId, sourceMessage.AuthorId)
	suite.Equal(messageFromDb.Text, sourceMessage.Text)
	suite.Equal(messageFromDb.Timestamp, sourceMessage.Timestamp)

}
