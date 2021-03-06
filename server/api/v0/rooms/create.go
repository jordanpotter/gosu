package rooms

import (
	"errors"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/api/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
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
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	passwordHash, err := password.ComputeBcryptHash(req.Password)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	room, err := h.dbConn.CreateRoom(req.Name, passwordHash, authToken.Account.ID, req.MemberName)
	if err == db.DuplicateError {
		c.AbortWithError(409, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, sanitization.ToRoom(room))
}
