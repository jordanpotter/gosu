package nanomsg

import (
	"fmt"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pub"
	"github.com/gdamore/mangos/transport/tcp"
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

func (p *publisher) Send(message []byte) error {
	fmt.Println("Sending", string(message))
	return p.sock.Send(message)
}

func (p *publisher) Close() error {
	return p.sock.Close()
}
