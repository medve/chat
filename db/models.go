package db

type Message struct {
	Id        int
	Text      string
	AuthorId  int64
	Timestamp int32
}
