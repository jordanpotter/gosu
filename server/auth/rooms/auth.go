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

	accountID := t.(*token.Token).Account.ID
	member, err := h.dbConn.Rooms.GetMemberByAccount(req.ID, accountID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	fmt.Println("TODO: add room and admin info to auth token")
	fmt.Println(member)

	c.String(200, "ok")
}
