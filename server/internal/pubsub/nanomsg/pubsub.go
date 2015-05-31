package nanomsg

import (
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/jordanpotter/gosu/server/events/types"
)

type message struct {
	Type       types.Type `msgpack:"type"`
	EventBytes []byte     `msgpack:"eventBytes"`
	Timestamp  time.Time  `msgpack:"timestamp"`
}

func newMessage(event types.Event) (*message, error) {
	b, err := msgpack.Marshal(event)
	return &message{event.GetType(), b, time.Now()}, err
}
