package accounts

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/db"
)

type Handler struct {
	dbConn db.Conn
}

func New(dbConn db.Conn) *Handler {
	return &Handler{dbConn}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", h.create)
	rg.POST("/login", h.login)
	rg.POST("/logout", h.logout)
}
