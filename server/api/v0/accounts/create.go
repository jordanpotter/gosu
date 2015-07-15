package accounts

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/api/v0/sanitization"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/events"
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

	devicePasswordHash, err := password.ComputeBcryptHash(req.DevicePassword)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	device, err := h.dbConn.CreateDevice(account.ID, req.DeviceName, devicePasswordHash)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = h.pub.Send(events.AccountDeviceCreated{
		AccountID:  account.ID,
		DeviceID:   device.ID,
		DeviceName: device.Name,
		Created:    device.Created,
	})
	if err != nil {
		fmt.Printf("Failed to send event: %v", err)
	}

	c.JSON(200, sanitization.ToAccount(account))
}
