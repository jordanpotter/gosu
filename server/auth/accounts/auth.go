package accounts

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/auth/token"
	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/jordanpotter/gosu/server/internal/middleware"
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
	err := c.Bind(&req)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	account, err := h.dbConn.GetAccountByEmail(req.Email)
	if err == db.NotFoundError {
		c.AbortWithError(403, err)
		return
	} else if err != nil {
		c.AbortWithError(500, err)
		return
	}

	devices, err := h.dbConn.GetDevicesByAccount(account.ID)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if !hasValidDeviceCredentials(devices, req.DeviceName, req.DevicePassword) {
		c.AbortWithError(403, errors.New("no matching device name and password"))
		return
	}

	authToken := h.tf.New()
	authToken.Account.ID = account.ID
	authTokenEncrypted, err := h.tf.Encrypt(authToken)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, AuthenticationResponse{authTokenEncrypted, authToken.Expires})
}

func hasValidDeviceCredentials(devices []db.Device, deviceName, devicePassword string) bool {
	for _, device := range devices {
		if deviceName == device.Name && password.MatchesBcryptHash(devicePassword, device.PasswordHash) {
			return true
		}
	}
	return false
}

func (h *Handler) reauthenticate(c *gin.Context) {
	t, ok := c.Get(middleware.TokenKey)
	if !ok {
		c.AbortWithError(500, errors.New("missing auth token"))
		return
	}
	authToken := t.(*token.Token)

	h.tf.Extend(authToken)
	authTokenEncrypted, err := h.tf.Encrypt(authToken)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, ReauthenticationResponse{authTokenEncrypted, authToken.Expires})
}
