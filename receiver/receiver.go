package receiver

type Notifier interface {
	Notify(message *Message) error
}

type Queue interface {
	Add(message *Message) error
	Receive() (*Message, error)
	GetSubscribeChan() (chan *Message, error)
}

type Saver interface {
	Save(message *Message) error
}

type Receiver interface {
	Receive(message *Message) error
}

type Message struct {
	Text      string
	AuthorId  int64
	Timestamp int32
}
