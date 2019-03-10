package db

import (
	"github.com/go-pg/pg/orm"
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
