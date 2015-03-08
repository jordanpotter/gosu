package accounts

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/auth/password"
	"github.com/JordanPotter/gosu-server/internal/auth/token"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type LoginRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	ClientName     string `json:"clientName" form:"clientName" binding:"required"`
	ClientPassword string `json:"clientPassword" form:"clientPassword" binding:"required"`
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

	account, err := h.dbConn.GetAccount(req.Email)
	if err == db.ErrNotFound {
		c.Fail(403, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	if !hasValidClientCredentials(req.ClientName, req.ClientPassword, account.Clients) {
		c.Fail(403, errors.New("no matching client name and password"))
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

func hasValidClientCredentials(clientName, clientPassword string, clients []db.Client) bool {
	for _, client := range clients {
		if clientName == client.Name && password.MatchesHash(clientPassword, client.PasswordHash) {
			return true
		}
	}
	return false
}
