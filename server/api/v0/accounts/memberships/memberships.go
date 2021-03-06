package memberships

import (
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	dbConn db.Conn
	tf     *token.Factory
	pub    pubsub.Publisher
}

func New(dbConn db.Conn, tf *token.Factory, pub pubsub.Publisher) *Handler {
	return &Handler{dbConn, tf, pub}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthRequired(h.tf))
	rg.GET("", h.getAll)
	rg.DELETE("/id/:memberID", h.delete)
}
