package nanomsg

import "time"

type message struct {
	Name      string    `msgpack:"name"`
	Data      []byte    `msgpack:"data"`
	Timestamp time.Time `msgpack:"timestamp"`
}
