package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
)

type CreateRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	DeviceName     string `json:"deviceName" form:"deviceName" binding:"required"`
	DevicePassword string `json:"devicePassword" form:"devicePassword" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	account, err := h.dbConn.CreateAccount(req.Email)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	devicePasswordHash, err := password.ComputeHash(req.DevicePassword)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	_, err = h.dbConn.CreateDevice(account.ID, req.DeviceName, devicePasswordHash)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.String(200, "ok")
}
