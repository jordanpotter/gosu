package rooms

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

const (
	roomNameQueryParam = "name"
)

type GetIDResponse struct {
	ID int `json:"id"`
}

func (h *Handler) get(c *gin.Context) {
	roomIDString := c.Params.ByName("roomID")
	roomID, err := strconv.Atoi(roomIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	room, err := h.dbConn.GetRoom(roomID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, room)
}

func (h *Handler) getID(c *gin.Context) {
	q := c.Request.URL.Query()
	if q[roomNameQueryParam] == nil {
		c.AbortWithError(400, errors.New("missing room name"))
		return
	}

	name := q[roomNameQueryParam][0]
	room, err := h.dbConn.GetRoomByName(name)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, GetIDResponse{room.ID})
}
