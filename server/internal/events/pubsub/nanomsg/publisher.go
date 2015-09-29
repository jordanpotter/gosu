package nanomsg

import (
	"time"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gdamore/mangos"
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gdamore/mangos/protocol/pub"
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gdamore/mangos/transport/tcp"
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/gopkg.in/vmihailenco/msgpack.v2"

	"github.com/jordanpotter/gosu/server/internal/events"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
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

func (p *publisher) Send(event events.Event) error {
	eventData, err := msgpack.Marshal(event)
	if err != nil {
		return err
	}

	m := message{
		Type:      event.GetType(),
		EventData: eventData,
		Timestamp: time.Now(),
	}
	b, err := msgpack.Marshal(m)
	if err != nil {
		return err
	}

	return p.sock.Send(b)
}

func (p *publisher) Close() error {
	return p.sock.Close()
}
