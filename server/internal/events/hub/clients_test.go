package hub

import "testing"

func TestAdd(t *testing.T) {
	cm := new(clientsManager)
	w := &testResponseWriter{
		writeChan:     make(chan bool, 1),
		closeNotifier: make(chan bool, 1),
	}
	_ = cm.add(w)
	if len(cm.clients) != 1 {
		t.Errorf("Unexpected number of clients %d", len(cm.clients))
	}
}

func TestRemove(t *testing.T) {
	cm := new(clientsManager)
	w := &testResponseWriter{
		writeChan:     make(chan bool, 1),
		closeNotifier: make(chan bool, 1),
	}
	id := cm.add(w)

	err := cm.remove(-1)
	if err == nil {
		t.Error("Expected error due to missing client")
	}

	err = cm.remove(id)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestSend(t *testing.T) {
	cm := new(clientsManager)
	w := &testResponseWriter{
		writeChan:     make(chan bool, 1),
		closeNotifier: make(chan bool, 1),
	}
	_ = cm.add(w)
	m := []byte("message")
	cm.send(m)
	<-w.writeChan
}
