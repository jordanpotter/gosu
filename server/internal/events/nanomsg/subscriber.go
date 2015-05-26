package nanomsg

import (
	"errors"
	"fmt"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/sub"
	"github.com/gdamore/mangos/transport/tcp"

	"github.com/jordanpotter/gosu/server/internal/events"
)

type subscriber struct {
	sock       mangos.Socket
	dialers    map[string]mangos.Dialer
	listenChan chan<- *events.Message
}

func NewSubscriber() (events.Subscriber, error) {
	sock, err := sub.NewSocket()
	if err != nil {
		return nil, err
	}

	fmt.Println("TODO: use TLSTCP instead")
	sock.AddTransport(tcp.NewTransport())
	return &subscriber{
		sock:    sock,
		dialers: make(map[string]mangos.Dialer),
	}, nil
}

func (s *subscriber) SetAddrs(addrs []string) error {
	err := s.addMissingConnections(addrs)
	if err != nil {
		return err
	}

	err = s.removeOldConnections(addrs)
	return err
}

func (s *subscriber) addMissingConnections(newAddrs []string) error {
	for _, addr := range newAddrs {
		_, exists := s.dialers[addr]
		if exists {
			continue
		}

		err := s.connect(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *subscriber) removeOldConnections(newAddrs []string) error {
Loop:
	for addr := range s.dialers {
		for _, newAddr := range newAddrs {
			if addr == newAddr {
				continue Loop
			}
		}

		err := s.disconnect(addr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *subscriber) connect(addr string) error {
	_, exists := s.dialers[addr]
	if exists {
		return fmt.Errorf("already connected to %s", addr)
	}

	dialer, err := s.sock.NewDialer("tcp://"+addr, nil)
	if err != nil {
		return err
	}

	err = dialer.Dial()
	if err != nil {
		return err
	}

	err = s.sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		dialer.Close()
		return err
	}

	s.dialers[addr] = dialer
	return nil
}

func (s *subscriber) disconnect(addr string) error {
	fmt.Println("disconnecting", addr)

	dialer, exists := s.dialers[addr]
	if !exists {
		return fmt.Errorf("not connected to %s", addr)
	}

	err := dialer.Close()
	if err != nil {
		return err
	}

	delete(s.dialers, addr)
	return nil
}

func (s *subscriber) Listen(listener chan<- *events.Message) error {
	if s.listenChan != nil {
		return errors.New("listener already set")
	} else if listener == nil {
		return errors.New("must provide listener")
	}

	s.listenChan = listener
	go s.handleEvents()
	return nil
}

func (s *subscriber) handleEvents() {
	for {
		event, err := s.getNextEvent()
		s.listenChan <- &events.Message{Event: event, Err: err}
	}
}

func (s *subscriber) getNextEvent() (interface{}, error) {
	b, err := s.sock.Recv()
	if err != nil {
		return nil, err
	}

	var m message
	err = msgpack.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	return m.getEvent()
}

func (s *subscriber) Close() error {
	err := s.sock.Close()
	s.dialers = make(map[string]mangos.Dialer)
	close(s.listenChan)
	return err
}
