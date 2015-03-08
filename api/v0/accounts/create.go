package accounts

import (
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	ClientName     string `json:"clientName" form:"clientName" binding:"required"`
	ClientPassword string `json:"clientPassword" form:"clientPassword" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	err := h.dbConn.CreateAccount(req.Email, req.ClientName, req.ClientPassword)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
