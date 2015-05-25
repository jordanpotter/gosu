package events

type Publisher interface {
	Send(event interface{}) error
	Close() error
}
