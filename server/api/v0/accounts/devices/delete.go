package devices

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
)

func (h *Handler) delete(c *gin.Context) {
	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	deviceIDString := c.Params.ByName("deviceID")
	deviceID, err := strconv.Atoi(deviceIDString)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	accountID := authToken.Account.ID
	err = h.dbConn.DeleteDeviceForAccount(deviceID, accountID)
	if err == db.NotFoundError {
		c.AbortWithError(404, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "ok")
}
