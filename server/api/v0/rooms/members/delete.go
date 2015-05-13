package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) leave(c *gin.Context) {
	roomName := c.Params.ByName("roomName")
	fmt.Printf("TODO: leave room %s", roomName)
	c.String(200, "ok")
}

func (h *Handler) delete(c *gin.Context) {
	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not revoking admin for self")

	roomName := c.Params.ByName("roomName")
	accountID := c.Params.ByName("memberAccountID")
	err := h.dbConn.Rooms.RemoveMember(roomName, accountID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
