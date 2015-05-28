package rooms

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type AuthenticationRequest struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type AuthenticationResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func (h *Handler) authenticate(c *gin.Context) {
	var req AuthenticationRequest
	if !c.Bind(&req) {
		return
	}

	t, err := c.Get(middleware.TokenKey)
	if err != nil {
		c.Fail(500, err)
		return
	}
	authToken := t.(*token.Token)

	accountID := authToken.Account.ID
	member, err := h.dbConn.Rooms.GetMemberByAccount(req.ID, accountID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Println(member)

	h.tf.Extend(authToken)
	authToken.Room.ID = req.ID
	authToken.Room.MemberID = member.ID
	authToken.Room.Admin = member.Admin
	fmt.Println(authToken.Room)
	authTokenEncrypted, err := h.tf.Encrypt(authToken)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, AuthenticationResponse{authTokenEncrypted, authToken.Expires})
}
