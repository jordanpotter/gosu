package events

type Publisher interface {
	Send(message []byte) error
	Close() error
}
