package accounts

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type Handler struct {
	dbConn       *db.Conn
	tokenFactory *token.Factory
}

func New(dbConn *db.Conn, tokenFactory *token.Factory) *Handler {
	return &Handler{dbConn, tokenFactory}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", h.create)
}
