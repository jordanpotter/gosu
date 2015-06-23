package pubsub

type Publisher interface {
	Send(name string, data []byte) error
	Close() error
}
