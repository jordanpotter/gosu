package accounts

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events"
)

type Handler struct {
	dbConn *db.Conn
	tf     *token.Factory
	pub    events.Publisher
}

func New(dbConn *db.Conn, tf *token.Factory, pub events.Publisher) *Handler {
	return &Handler{dbConn, tf, pub}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", h.create)
}
