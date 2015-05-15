package members

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) leave(c *gin.Context) {
	roomID := c.Params.ByName("roomID")
	fmt.Printf("TODO: leave room %s", roomID)
	c.String(200, "ok")
}

func (h *Handler) delete(c *gin.Context) {
	fmt.Println("TODO: check account is admin for room")
	fmt.Println("TODO: check not revoking admin for self")

	roomID := c.Params.ByName("roomID")
	memberID := c.Params.ByName("memberID")
	err := h.dbConn.Rooms.RemoveMember(roomID, memberID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
