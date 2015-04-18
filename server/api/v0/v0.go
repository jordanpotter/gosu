package v0

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/accounts"
	"github.com/jordanpotter/gosu/server/api/v0/rooms"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type Handler struct {
	accountsHandler *accounts.Handler
	roomsHandler    *rooms.Handler
}

func New(dbConn *db.Conn, tokenFactory *token.Factory) *Handler {
	return &Handler{
		accountsHandler: accounts.New(dbConn, tokenFactory),
		roomsHandler:    rooms.New(dbConn, tokenFactory),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	h.accountsHandler.AddRoutes(rg.Group("/accounts"))
	h.roomsHandler.AddRoutes(rg.Group("/rooms"))
}
