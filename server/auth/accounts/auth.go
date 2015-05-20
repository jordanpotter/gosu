package accounts

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type AuthenticationRequest struct {
	Email          string `json:"email" form:"email" binding:"required"`
	DeviceName     string `json:"deviceName" form:"deviceName" binding:"required"`
	DevicePassword string `json:"devicePassword" form:"devicePassword" binding:"required"`
}

type AuthenticationResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

type ReauthenticationResponse struct {
	AuthToken   string    `json:"authToken"`
	AuthExpires time.Time `json:"authExpires"`
}

func (h *Handler) authenticate(c *gin.Context) {
	var req AuthenticationRequest
	if !c.Bind(&req) {
		return
	}

	account, err := h.dbConn.Accounts.GetByEmail(req.Email)
	if err == db.NotFoundError {
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

	authToken := h.tf.New()
	authToken.Account.ID = account.ID
	authTokenEncrypted, err := h.tf.Encrypt(authToken)
	if err != nil {
		c.Fail(500, err)
		return
	}

	c.JSON(200, AuthenticationResponse{authTokenEncrypted, authToken.Expires})
}

func hasValidDeviceCredentials(deviceName, devicePassword string, devices []db.Device) bool {
	for _, device := range devices {
		if deviceName == device.Name && password.MatchesHash(devicePassword, device.PasswordHash) {
			return true
		}
	}
	return false
}

func (h *Handler) reauthenticate(c *gin.Context) {
	fmt.Println("TODO: reauthenticate user")
	c.String(200, "ok")
}
