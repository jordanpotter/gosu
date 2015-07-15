package hub

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

type testResponseWriter struct {
	numWrites     int
	writeChan     chan bool
	closeNotifier chan bool
}

func (trw *testResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (trw *testResponseWriter) Write(b []byte) (int, error) {
	trw.numWrites++
	trw.writeChan <- true
	return len(b), nil
}

func (trw *testResponseWriter) WriteHeader(code int) {
}

func (trw *testResponseWriter) Flush() {
}

func (trw *testResponseWriter) CloseNotify() <-chan bool {
	return trw.closeNotifier
}

func createSubscription(hub *Hub, key string) *testResponseWriter {
	w := &testResponseWriter{
		writeChan:     make(chan bool, 1000),
		closeNotifier: make(chan bool, 1),
	}
	go hub.SubscribeAndBlock(key, w)
	return w
}

func TestSubscription(t *testing.T) {
	hub := New()
	key := "some-key"
	w := createSubscription(hub, key)
	time.Sleep(10 * time.Millisecond)

	if hub.clientManagers[key] == nil {
		t.Fatalf("Client manager for key %s not added", key)
	} else if len(hub.clientManagers[key].clients) != 1 {
		t.Errorf("Unexpected number of client managers for key %s: %d", key, len(hub.clientManagers[key].clients))
	}

	w.closeNotifier <- true
	time.Sleep(10 * time.Millisecond)

	if len(hub.clientManagers[key].clients) != 0 {
		t.Errorf("Expected no remaining clients, received: %d", len(hub.clientManagers[key].clients))
	}
}

func TestSending(t *testing.T) {
	hub := New()
	key := "some-key"
	w := createSubscription(hub, key)
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 100; i++ {
		m := make([]byte, i+1000)
		hub.Send("some-key", m)
		<-w.writeChan
	}
}

func TestHeavyLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	createKeys := func(numKeys int) []string {
		keys := make([]string, numKeys)
		for i := 0; i < len(keys); i++ {
			keys[i] = fmt.Sprintf("key-%d", i)
		}
		return keys
	}

	createWriters := func(hub *Hub, keys []string, numWritersPerKey int) map[string][]*testResponseWriter {
		writers := make(map[string][]*testResponseWriter)
		for _, key := range keys {
			writers[key] = make([]*testResponseWriter, 0)
			for i := 0; i < numWritersPerKey; i++ {
				writers[key] = append(writers[key], createSubscription(hub, key))
			}
		}
		time.Sleep(10 * time.Millisecond)
		return writers
	}

	writeMessages := func(hub *Hub, keys []string, writers map[string][]*testResponseWriter, numMessages int) {
		completionChan := make(chan bool, numMessages)
		numWrites := 0

		for i := 0; i < numMessages; i++ {
			key := keys[rand.Intn(len(keys))]
			m := []byte("hello")
			hub.Send(key, m)

			for _, w := range writers[key] {
				numWrites++
				go func(w *testResponseWriter) {
					completionChan <- <-w.writeChan
				}(w)
			}
		}

		for i := 0; i < numWrites; i++ {
			<-completionChan
		}
	}

	hub := New()
	keys := createKeys(1000)
	writers := createWriters(hub, keys, 20)
	writeMessages(hub, keys, writers, 100000)
}
