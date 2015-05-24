package nanomsg

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/sub"
	"github.com/gdamore/mangos/transport/tcp"

	"github.com/jordanpotter/gosu/server/internal/events"
)

type subscriber struct {
	sock mangos.Socket
}

func NewSubscriber(addr string) (events.Subscriber, error) {
	sock, err := sub.NewSocket()
	if err != nil {
		return nil, err
	}

	sock.AddTransport(tcp.NewTransport())
	err = sock.Dial("tcp://" + addr)
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		return nil, err
	}

	return &subscriber{sock}, nil
}

func (s *subscriber) Listen() <-chan events.Message {
	c := make(chan events.Message)
	go s.handleMessages(c)
	return c
}

func (s *subscriber) handleMessages(c chan<- events.Message) {
	for {
		data, err := s.sock.Recv()
		c <- events.Message{Data: data, Err: err}
	}
}

func (s *subscriber) Close() error {
	return s.sock.Close()
}
