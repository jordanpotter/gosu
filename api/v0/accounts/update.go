package accounts

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/JordanPotter/gosu-server/internal/auth/password"
	"github.com/JordanPotter/gosu-server/internal/db"
)

type LoginRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	DeviceName     string `json:"deviceName" form:"deviceName" binding:"required"`
	DevicePassword string `json:"devicePassword" form:"devicePassword" binding:"required"`
}

type LoginResponse struct {
	Id          string    `json:"id"`
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if !c.Bind(&req) {
		return
	}

	account, err := h.dbConn.GetAccount(req.Email)
	if err == db.ErrNotFound {
		c.Fail(403, err)
		return
	} else if err != nil {
		c.Fail(500, err)
		return
	}

	if !hasValidDeviceCredentials(req.DeviceName, req.DevicePassword, account.Devices) {
		c.Fail(403, errors.New("no matching device name and password"))
		return
	}

	authToken := h.tokenFactory.New(account.Id)
	authTokenEncrypted, err := h.tokenFactory.Encrypt(authToken)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, LoginResponse{account.Id, authTokenEncrypted, authToken.Expires})
}

func hasValidDeviceCredentials(deviceName, devicePassword string, devices []db.Device) bool {
	for _, device := range devices {
		if deviceName == device.Name && password.MatchesHash(devicePassword, device.PasswordHash) {
			return true
		}
	}
	return false
}
