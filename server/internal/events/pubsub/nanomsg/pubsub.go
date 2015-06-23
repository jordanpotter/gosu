package nanomsg

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/events"
)

type message struct {
	Type      events.Type `msgpack:"type"`
	EventData []byte      `msgpack:"eventData"`
	Timestamp time.Time   `msgpack:"timestamp"`
}
