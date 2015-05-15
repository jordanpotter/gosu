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

	fmt.Println("TODO: make sure admin for room")

	roomID := c.Params.ByName("roomID")
	err := h.dbConn.Rooms.AddChannel(roomID, req.Name)
	if err == db.DuplicateError {
		c.Fail(409, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
