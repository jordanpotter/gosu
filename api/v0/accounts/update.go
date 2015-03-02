package accounts

import (
	"errors"
	"time"

	"github.com/JordanPotter/gosu-server/internal/auth"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResponse struct {
	Id            string    `json:"id"`
	AuthEncrypted string    `json:"authToken"`
	AuthExpires   time.Time `json:"authExpires"`
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if !c.Bind(&req) {
		return
	}

	account, err := h.dbConn.GetAccount(req.Name, req.Password)
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

	c.JSON(200, CreateResponse{account.Id, authEncrypted, auth.Expires})
}

func (h *Handler) logout(c *gin.Context) {
	authEncrypted := c.Request.Header.Get("authorization")
	auth, err := auth.Decrypt(authEncrypted)
	if err != nil {
		c.Fail(500, err)
		return
	}

	if auth.Expires.Before(time.Now()) {
		c.Fail(403, errors.New("accounts: token has expired"))
		return
	}

	// TODO: update database

	c.String(200, "success")
}
