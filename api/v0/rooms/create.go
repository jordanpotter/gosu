package rooms

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

	fmt.Printf("TODO: create room %s with password %s if doesn't exist\n", req.Name, req.Password)

	c.String(200, "ok")
}
