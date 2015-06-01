package hub

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/pubsub"
)

const (
	subMessageBuffer = 100
)

type Hub struct {
	tf            *token.Factory
	listenChan    chan *pubsub.SubMessage
	clientsByRoom map[string]*gin.Context
}

func New(tf *token.Factory, sub pubsub.Subscriber) (*Hub, error) {
	h := &Hub{
		tf:            tf,
		listenChan:    make(chan *pubsub.SubMessage, subMessageBuffer),
		clientsByRoom: make(map[string]*gin.Context),
	}

	err := sub.Listen(h.listenChan)
	if err != nil {
		return nil, err
	}

	go h.handleMessages()
	return h, nil
}

func (h *Hub) handleMessages() {
	for m := range h.listenChan {
		fmt.Println(m.Event)
	}
}

func (h *Hub) AddRoutes(rg *gin.RouterGroup) {
	// rg.Use(middleware.AuthRequired(h.tf))
	rg.POST("/subscribe", h.subscribe)
}

func (h *Hub) subscribe(c *gin.Context) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	for {
		time.Sleep(5 * time.Second)
		fmt.Fprint(c.Writer, "hello")
		c.Writer.Flush()
		fmt.Println("wrote hello")
	}

	// c.String(200, "ok")
}
