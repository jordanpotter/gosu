package members

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
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

	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.dbConn.DeleteMemberForAccountAndRoom(authToken.Account.ID, roomID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// e := &types.RoomMemberDeleted{
	// 	RoomID:   roomID,
	// 	MemberID: member.ID,
	// }
	// err = h.pub.Send(e)
	// if err != nil {
	// 	fmt.Printf("Failed to send event: %v", err)
	// }

	c.String(200, "ok")
}

// func (h *Handler) delete(c *gin.Context) {
// 	fmt.Println("TODO: check not revoking admin for self")
//
// 	roomID := c.Params.ByName("roomID")
// 	memberID := c.Params.ByName("memberID")
// 	err := h.dbConn.Rooms.RemoveMember(roomID, memberID)
// 	if err != nil {
// 		c.AbortWithError(500, err)
// 		return
// 	}
//
// 	e := &types.RoomMemberDeleted{
// 		RoomID:   roomID,
// 		MemberID: memberID,
// 	}
// 	err = h.pub.Send(e)
// 	if err != nil {
// 		fmt.Printf("Failed to send event: %v", err)
// 	}
//
// 	c.String(200, "ok")
// }
