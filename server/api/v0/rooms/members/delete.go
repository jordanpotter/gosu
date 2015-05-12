package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	memberName := c.Params.ByName("memberName")

	err := h.dbConn.Rooms.RemoveMember(roomName, memberName)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
