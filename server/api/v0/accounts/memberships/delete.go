package memberships

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

func (h *Handler) delete(c *gin.Context) {
	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	memberIDString := c.Params.ByName("memberID")
	memberID, err := strconv.Atoi(memberIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.dbConn.DeleteMemberForAccount(memberID, authToken.Account.ID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "ok")
}
