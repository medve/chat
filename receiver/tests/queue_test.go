package receiver

import (
	"testing"

	"github.com/chat/receiver/queue"

	"github.com/chat/receiver/message"
	nats "github.com/nats-io/go-nats"
)

func TestSendAndReceiveNats(t *testing.T) {

	nc, _ := nats.Connect(nats.DefaultURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	sendCh := make(chan *message.Message)
	ec.BindSendChan("hello", sendCh)

	recvCh := make(chan *message.Message)
	ec.BindRecvChan("hello", recvCh)

	msg := &message.Message{Text: "asd", AuthorId: 123, Timestamp: 1}

	// Send via Go channels
	sendCh <- msg

	// Receive via Go channels
	rcvd_msg := <-recvCh

	if *rcvd_msg != *msg {
		t.Errorf("Sended message not equals to received %v != %v", msg, rcvd_msg)
	}
}

func createQueue(t *testing.T, chanMod queue.ChanMode) (*queue.NatsQueue, *nats.EncodedConn) {
	ec, err := queue.CreateNatsConn(nats.DefaultURL)

	if err != nil {
		t.Error("Error should not be occured", err)
	}

	nq, err := queue.CreateNatsQueue("asd", ec, chanMod)

	if err != nil {
		t.Error("Error should not be occured", err)
	}

	if nq == nil {
		t.Error("Nats queue should not be nil")
	}

	return nq, ec
}

func TestCreateQueue(t *testing.T) {
	nq, ec := createQueue(t, queue.Write)
	defer ec.Close()

	_, err := nq.Receive()

	if err == nil {
		t.Error("Error should be occured. Write mode is setted.")
	}

	_, err = nq.GetSubscribeChan()

	if err == nil {
		t.Error("Error should be occured. Write mode is setted.")
	}

	msg := &message.Message{Text: "asd", AuthorId: 123, Timestamp: 1}
	err = nq.Add(msg)

	if err != nil {
		t.Error("Error should not be occured.", err)
	}

}

func TestAddMsg(t *testing.T) {

	// NOTE: If we will create write queue before read queue, than deadlock will be occured

	rcvNq, recvEc := createQueue(t, queue.Read)
	defer recvEc.Close()

	nq, ec := createQueue(t, queue.Write)
	defer ec.Close()

	rcvMsgChan, err := rcvNq.GetSubscribeChan()

	if err != nil {
		t.Error("Error should not be occured.", err)
	}

	msg := &message.Message{Text: "asd", AuthorId: 123, Timestamp: 1}
	err = nq.Add(msg)

	rcv_msg := <-rcvMsgChan

	if err != nil {
		t.Error("Error should not be occured.", err)
	}

	if *rcv_msg != *msg {
		t.Errorf("Sended message not equals to received %v != %v", msg, rcv_msg)
	}

}

func TestAddMsgRecv(t *testing.T) {

	rcvNq, recvEc := createQueue(t, queue.Read)
	defer recvEc.Close()

	nq, ec := createQueue(t, queue.Write)
	defer ec.Close()

	rcvMsgChan, err := rcvNq.GetSubscribeChan()

	if err != nil {
		t.Error("Error should not be occured.", err)
	}

	msg := &message.Message{Text: "asd", AuthorId: 123, Timestamp: 1}
	err = nq.Add(msg)

	rcv_msg := <-rcvMsgChan

	if err != nil {
		t.Error("Error should not be occured.", err)
	}

	if *rcv_msg != *msg {
		t.Errorf("Sended message not equals to received %v != %v", msg, rcv_msg)
	}

}
