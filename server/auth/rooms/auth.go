package rooms

import (
	"errors"
	"time"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type AuthenticationRequest struct {
	ID int `json:"id" form:"id" binding:"required"`
}

type AuthenticationResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func (h *Handler) authenticate(c *gin.Context) {
	var req AuthenticationRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	member, err := h.dbConn.GetMemberByAccountAndRoom(authToken.Account.ID, req.ID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	} else if member.Banned {
		c.AbortWithError(403, errors.New("banned"))
		return
	}

	h.tf.Extend(authToken)
	authToken.Room.ID = req.ID
	authToken.Room.MemberID = member.ID
	authToken.Room.Admin = member.Admin
	authTokenEncrypted, err := h.tf.Encrypt(authToken)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, AuthenticationResponse{authTokenEncrypted, authToken.Expires})
}
