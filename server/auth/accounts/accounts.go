package accounts

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type Handler struct {
	dbConn db.Conn
	tf     *token.Factory
}

func New(dbConn db.Conn, tf *token.Factory) *Handler {
	return &Handler{dbConn, tf}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/authenticate", h.authenticate)
	// rg.POST("/reauthenticate", middleware.AuthRequired(h.tf), h.reauthenticate)
}
