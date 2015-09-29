package v0

import (
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/accounts"
	"github.com/jordanpotter/gosu/server/api/v0/rooms"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
)

type Handler struct {
	accountsHandler *accounts.Handler
	roomsHandler    *rooms.Handler
}

func New(dbConn db.Conn, tf *token.Factory, pub pubsub.Publisher) *Handler {
	return &Handler{
		accountsHandler: accounts.New(dbConn, tf, pub),
		roomsHandler:    rooms.New(dbConn, tf, pub),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	h.accountsHandler.AddRoutes(rg.Group("/accounts"))
	h.roomsHandler.AddRoutes(rg.Group("/rooms"))
}
