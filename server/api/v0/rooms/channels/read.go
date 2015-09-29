package channels

import (
	"strconv"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func (h *Handler) getAll(c *gin.Context) {
	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	channels, err := h.dbConn.GetChannelsByRoom(roomID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, sanitization.ToChannels(channels))
}
