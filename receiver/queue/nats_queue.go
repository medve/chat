package queue

import (
	"errors"

	"github.com/chat/receiver/message"
	"github.com/nats-io/go-nats"
)

const BufferSize = 0

type ChanMode int

const (
	Read ChanMode = iota
	Write
	ReadAndWrite
)

type NatsQueue struct {
	sendChan chan *message.Message
	recvChan chan *message.Message
}

func CreateNatsConn(url string) (*nats.EncodedConn, error) {
	nc, err := nats.Connect(url)

	if err != nil {
		return nil, err
	}

	return nats.NewEncodedConn(nc, nats.JSON_ENCODER)
}

func CreateNatsQueue(chanName string, natsConn *nats.EncodedConn, chanMode ChanMode) (*NatsQueue, error) {

	var (
		recvCh chan *message.Message
		sendCh chan *message.Message
		err    error
	)

	if chanMode == Read || chanMode == ReadAndWrite {
		recvCh = make(chan *message.Message, BufferSize)
		_, err = natsConn.BindRecvChan(chanName, recvCh)

		if err != nil {
			return nil, err
		}
	}

	if chanMode == Write || chanMode == ReadAndWrite {
		sendCh = make(chan *message.Message, BufferSize)
		err = natsConn.BindSendChan(chanName, sendCh)

		if err != nil {
			return nil, err
		}
	}

	return &NatsQueue{
		sendChan: sendCh,
		recvChan: recvCh,
	}, nil
}

func (queue *NatsQueue) Add(message *message.Message) error {

	if queue.sendChan == nil {
		return errors.New("Send channel did not created")
	}

	queue.sendChan <- message

	return nil
}

func (queue *NatsQueue) Receive() (*message.Message, error) {

	if queue.recvChan == nil {
		return nil, errors.New("Recv channel did not created")
	}

	result := <-queue.recvChan

	return result, nil
}

func (queue *NatsQueue) GetSubscribeChan() (chan *message.Message, error) {
	if queue.recvChan == nil {
		return nil, errors.New("Recv channel did not created")
	}

	return queue.recvChan, nil
}
