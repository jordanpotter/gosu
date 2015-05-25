package nanomsg

import (
	"fmt"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pub"
	"github.com/gdamore/mangos/transport/tcp"
	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/jordanpotter/gosu/server/internal/events"
)

type publisher struct {
	sock mangos.Socket
}

func NewPublisher(addr string) (events.Publisher, error) {
	sock, err := pub.NewSocket()
	if err != nil {
		return nil, err
	}

	sock.AddTransport(tcp.NewTransport())
	err = sock.Listen("tcp://" + addr)
	if err != nil {
		return nil, err
	}

	return &publisher{sock}, nil
}

func (p *publisher) Send(event interface{}) error {
	fmt.Println("Sending", event)
	fmt.Println("TODO: send timestamp in event")

	m, err := newMessage(event)
	if err != nil {
		return err
	}

	b, err := msgpack.Marshal(&m)
	if err != nil {
		return err
	}

	return p.sock.Send(b)
}

func (p *publisher) Close() error {
	return p.sock.Close()
}
