package db

import (
	"github.com/go-pg/pg/orm"

	"github.com/chat/receiver"
)

type PostgresDbDriver struct {
	db orm.DB
}

func NewPostgresDbDriver(db orm.DB) *PostgresDbDriver {
	return &PostgresDbDriver{
		db: db,
	}
}

func (driver *PostgresDbDriver) TestFunc() error {

	queries := []string{
		`DROP TABLE IF EXISTS tx_test`,
		`CREATE TABLE tx_test(counter int)`,
		`INSERT INTO tx_test (counter) VALUES (0)`,
	}
	for _, q := range queries {
		_, err := driver.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func (driver *PostgresDbDriver) SaveMessage(message *receiver.Message) error {
	err := driver.db.Insert(&Message{
		Text:      message.Text,
		AuthorId:  message.AuthorId,
		Timestamp: message.Timestamp,
	})
	return err
}
