package v0

import (
	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/api/v0/accounts"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type Handler struct {
	accountsHandler *accounts.Handler
}

func New(dbConn db.Conn) *Handler {
	return &Handler{
		accountsHandler: accounts.New(dbConn),
	}
}

func (h *Handler) AddRoutes(rg *gin.RouterGroup) {
	h.accountsHandler.AddRoutes(rg.Group("/accounts"))
}
