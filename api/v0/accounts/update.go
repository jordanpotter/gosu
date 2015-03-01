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

func login(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	// TODO: create account in database
	id := "id"

	auth := auth.New(id)
	authEncrypted, err := auth.Encrypt()
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, CreateResponse{id, authEncrypted, auth.Expires})
}

func logout(c *gin.Context) {
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
