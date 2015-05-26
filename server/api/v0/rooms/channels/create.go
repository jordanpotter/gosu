package channels

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type CreateRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	roomID := c.Params.ByName("roomID")
	channel, err := h.dbConn.Rooms.AddChannel(roomID, req.Name)
	if err == db.DuplicateError {
		c.Fail(409, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}
	fmt.Println(channel)

	// e := events.RoomChannelCreated{
	// 	RoomID: roomID,
	// 	ChannelID: "channel-id"
	// 	ChannelName: "channel-name"
	// }
	// err := h.pub.Send(e)
	// if err != nil {
	// 	fmt.Println("Failed to send event: %v", err)
	// }

	c.String(200, "ok")
}
