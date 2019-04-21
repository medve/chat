package saver

import "github.com/chat/receiver"

func NewDbSaver(driver DbDriver) (*DbSaver, error) {
	return &DbSaver{
		driver: driver,
	}, nil
}

type DbSaver struct {
	driver DbDriver
}

func (storage *DbSaver) Save(message *receiver.Message) error {
	//err := storage.driver.SaveMessage(message)
	//return err
	return nil
}
