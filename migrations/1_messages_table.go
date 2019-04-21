package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

type Message struct {
	Id        int
	Text      string
	AuthorId  int32
	Timestamp int32
}

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table message...")
		_, err := db.Exec(`CREATE TABLE message(id integer PRIMARY KEY, text varchar(1000), author_id integer, timestamp integer)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table message...")
		_, err := db.Exec(`DROP TABLE message`)
		return err
	})
}
