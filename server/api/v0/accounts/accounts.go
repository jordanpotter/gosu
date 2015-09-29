package accounts

import (
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/accounts/devices"
	"github.com/jordanpotter/gosu/server/api/v0/accounts/memberships"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/events/pubsub"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	dbConn             db.Conn
	tf                 *token.Factory
	pub                pubsub.Publisher
	devicesHandler     *devices.Handler
	membershipsHandler *memberships.Handler
}

func New(dbConn db.Conn, tf *token.Factory, pub pubsub.Publisher) *Handler {
	return &Handler{
		dbConn:             dbConn,
		tf:                 tf,
		pub:                pub,
		devicesHandler:     devices.New(dbConn, tf, pub),
		membershipsHandler: memberships.New(dbConn, tf, pub),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", h.create)
	rg.GET("", middleware.AuthRequired(h.tf), h.get)

	h.devicesHandler.AddRoutes(rg.Group("/devices"))
	h.membershipsHandler.AddRoutes(rg.Group("/memberships"))
}
