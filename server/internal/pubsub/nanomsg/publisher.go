package nanomsg

import (
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pub"
	"github.com/gdamore/mangos/transport/tcp"
	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/jordanpotter/gosu/server/events/types"
	"github.com/jordanpotter/gosu/server/internal/pubsub"
)

type publisher struct {
	sock mangos.Socket
}

func NewPublisher(addr string) (pubsub.Publisher, error) {
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

func (p *publisher) Send(event types.Event) error {
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
