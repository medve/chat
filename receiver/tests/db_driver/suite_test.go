package tests

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/suite"
)

type DBSuite struct {
	suite.Suite
	transactions map[string]*pg.Tx
	db           *pg.DB
}

func (suite *DBSuite) SetupSuite() {
	// drop and create db if need
	suite.db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5432",
		Database: "postgres",
	})

	suite.NotNil(suite.db, "Error connect to db")

	suite.transactions = make(map[string]*pg.Tx)

	// TODO drop/create tables, migrations

}

func (suite *DBSuite) BeforeTest(suiteName, testName string) {

	tx, err := suite.db.Begin()

	suite.NotNil(tx, "Error begin transaction")

	if err != nil {
		suite.Error(err, "Error in begin of transaction")
	}

	suite.transactions[testName] = tx
}

func (suite *DBSuite) AfterTest(suiteName, testName string) {
	tx := suite.transactions[testName]

	err := tx.Rollback()
	if err != nil {
		suite.Error(err, "Error in rollback of transaction")
	}
}

func (suite *DBSuite) TearDownSuite() {
	err := suite.db.Close()

	if err != nil {
		suite.Error(err, "Error while close")
	}
}

func (suite *DBSuite) GetConnection(testName string) *pg.Tx {
	return suite.transactions[testName]
}

func TestDBSuite(t *testing.T) {
	suite.Run(t, new(DBSuite))
}
