package accounts

import (
	"time"

	"github.com/JordanPotter/gosu-server/internal/auth"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CreateResponse struct {
	Id            string    `json:"id"`
	AuthEncrypted string    `json:"authToken"`
	AuthExpires   time.Time `json:"authExpires"`
}

func create(c *gin.Context) {
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
