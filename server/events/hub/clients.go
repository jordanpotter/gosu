package hub

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/events/types"
)

type clientsManager struct {
	clients []client
	lock    sync.RWMutex
	counter int
}

type client struct {
	id int
	w  gin.ResponseWriter
}

func (cm *clientsManager) add(w gin.ResponseWriter) int {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	c := client{cm.counter, w}
	cm.counter++
	cm.clients = append(cm.clients, c)
	return c.id
}

func (cm *clientsManager) remove(id int) error {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	for i, c := range cm.clients {
		if c.id == id {
			cm.clients = append(cm.clients[:i], cm.clients[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("missing client with id %d", id)
}

func (cm *clientsManager) sendEvent(event types.Event) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()

	for _, c := range cm.clients {
		fmt.Fprint(c.w, event)
		c.w.Flush()
	}
}
