package rooms

import (
	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type Handler struct {
	dbConn db.Conn
	tf     *token.Factory
}

func New(dbConn db.Conn, tf *token.Factory) *Handler {
	return &Handler{dbConn, tf}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/authenticate", middleware.AuthRequired(h.tf), h.authenticate)
}
