package v0

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/events/v0/accounts"
	"github.com/jordanpotter/gosu/server/events/v0/rooms"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
)

type Handler struct {
	accountsHandler *accounts.Handler
	roomsHandler    *rooms.Handler
}

func New(tf *token.Factory, sub pubsub.Subscriber) *Handler {
	return &Handler{
		accountsHandler: accounts.New(tf, sub),
		roomsHandler:    rooms.New(tf, sub),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	h.accountsHandler.AddRoutes(rg.Group("/accounts"))
	h.roomsHandler.AddRoutes(rg.Group("/rooms"))
}
