package accounts

import (
	"time"

	"github.com/JordanPotter/gosu-server/internal/auth"
	"github.com/gin-gonic/gin"
)

type CreateResponse struct {
	Id            string    `json:"id"`
	Password      string    `json:"password"`
	AuthEncrypted string    `json:"auth"`
	AuthExpires   time.Time `json:"authExpires"`
}

func (h *Handler) create(c *gin.Context) {
	account, password, err := h.dbConn.CreateAccount()
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

	c.JSON(200, CreateResponse{account.Id, password, authEncrypted, auth.Expires})
}
