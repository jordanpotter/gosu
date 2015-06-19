package channels

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type CreateRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	channel, err := h.dbConn.CreateChannel(roomID, req.Name)
	if err == db.DuplicateError {
		c.AbortWithError(409, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// e := &types.RoomChannelCreated{
	// 	RoomID:      roomID,
	// 	ChannelID:   channel.ID,
	// 	ChannelName: channel.Name,
	// 	Created:     channel.Created,
	// }
	// err = h.pub.Send(e)
	// if err != nil {
	// 	fmt.Printf("Failed to send event: %v", err)
	// }

	c.JSON(200, channel)
}
