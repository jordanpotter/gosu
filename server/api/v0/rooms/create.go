package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type CreateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	fmt.Println("TODO: make sure auth token is provided and valid first")

	err := h.dbConn.Rooms.Create(req.Name, req.Password)
	if err == db.DuplicateError {
		c.Fail(409, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
