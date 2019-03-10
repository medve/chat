package receiver

func NewSimpleReceiver(saver Saver, notifier Notifier) *SimpleReceiver {
	return &SimpleReceiver{
		saver:    saver,
		notifier: notifier,
	}
}

type SimpleReceiver struct {
	saver    Saver
	notifier Notifier
}

func (receiver *SimpleReceiver) Receive(message *Message) error {

	err := receiver.saver.Save(message)

	if err != nil {
		return err
	}

	err = receiver.notifier.Notify(message)

	if err != nil {
		return err
	}

	return nil
}
