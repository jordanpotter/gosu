package accounts

import (
	"time"

	"github.com/JordanPotter/gosu-server/internal/auth"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Id       string `json:"id" form:"id" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Id            string    `json:"id"`
	AuthEncrypted string    `json:"auth"`
	AuthExpires   time.Time `json:"authExpires"`
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if !c.Bind(&req) {
		return
	}

	account, err := h.dbConn.GetAccount(req.Id, req.Password)
	if err != nil {
		c.Fail(500, err)
		return
	}

	auth := auth.New(account.Id)
	authEncrypted, err := auth.Encrypt()
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, LoginResponse{account.Id, authEncrypted, auth.Expires})
}
