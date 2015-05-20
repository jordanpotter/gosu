package rooms

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/db"
)

const (
	roomNameQueryParam = "name"
)

func (h *Handler) get(c *gin.Context) {
	roomID := c.Params.ByName("roomID")
	room, err := h.dbConn.Rooms.Get(roomID)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, room)
}

func (h *Handler) getID(c *gin.Context) {
	q := c.Request.URL.Query()
	if q[roomNameQueryParam] == nil {
		c.Fail(400, errors.New("missing room name"))
		return
	}

	name := q[roomNameQueryParam][0]
	room, err := h.dbConn.Rooms.GetByName(name)
	if err == db.NotFoundError {
		c.Fail(404, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, room.ID)
}
