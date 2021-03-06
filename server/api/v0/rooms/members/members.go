package members

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
	rg.GET("", h.getAll)

	rg.POST("/join", h.join)
	rg.DELETE("/leave", h.leave)

	rgWithID := rg.Group("/id/:memberID")
	rgWithID.Use(middleware.IsRoomAdmin(), middleware.IsNotSameMember("memberID"))
	rgWithID.PUT("/admin", h.setAdmin)
	rgWithID.PUT("/banned", h.setBanned)
	rgWithID.DELETE("", h.delete)
}
