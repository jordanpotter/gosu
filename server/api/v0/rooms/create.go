package rooms

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

type CreateRequest struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Password   string `json:"password" form:"password" binding:"required"`
	MemberName string `json:"memberName" form:"memberName" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	if !c.Bind(&req) {
		return
	}

	t, err := c.Get(middleware.TokenKey)
	if err != nil {
		c.Fail(500, err)
		return
	}

	accountID := t.(*token.Token).Account.ID
	err = h.dbConn.Rooms.Create(req.Name, req.Password, accountID, req.MemberName)
	if err == db.DuplicateError {
		c.Fail(409, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	c.String(200, "ok")
}
