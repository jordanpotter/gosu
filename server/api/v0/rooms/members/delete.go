package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

func (h *Handler) leave(c *gin.Context) {
	t, err := c.Get(middleware.TokenKey)
	if err != nil {
		c.Fail(500, err)
		return
	}
	authToken := t.(*token.Token)

	roomID := c.Params.ByName("roomID")
	accountID := authToken.Account.ID
	member, err := h.dbConn.Rooms.GetMemberByAccount(roomID, accountID)
	if err != nil {
		c.Fail(500, err)
		return
	}

	err = h.dbConn.Rooms.RemoveMember(roomID, member.AccountID)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}

func (h *Handler) delete(c *gin.Context) {
	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not revoking admin for self")

	roomID := c.Params.ByName("roomID")
	memberID := c.Params.ByName("memberID")
	err := h.dbConn.Rooms.RemoveMember(roomID, memberID)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
