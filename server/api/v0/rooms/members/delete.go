package members

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/events/types"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

func (h *Handler) leave(c *gin.Context) {
	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	roomID := c.Params.ByName("roomID")
	accountID := authToken.Account.ID
	member, err := h.dbConn.Rooms.GetMemberByAccount(roomID, accountID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.dbConn.Rooms.RemoveMember(roomID, member.ID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	e := &types.RoomMemberDeleted{
		RoomID:   roomID,
		MemberID: member.ID,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.String(200, "ok")
}

func (h *Handler) delete(c *gin.Context) {
	fmt.Println("TODO: check not revoking admin for self")

	roomID := c.Params.ByName("roomID")
	memberID := c.Params.ByName("memberID")
	err := h.dbConn.Rooms.RemoveMember(roomID, memberID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	e := &types.RoomMemberDeleted{
		RoomID:   roomID,
		MemberID: memberID,
	}
	err = h.pub.Send(e)
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.String(200, "ok")
}
