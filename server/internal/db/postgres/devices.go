package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedDevice struct {
	ID           int       `db:"id"`
	AccountID    int       `db:"account_id"`
	Name         string    `db:"email"`
	PasswordHash []byte    `db:"password_hash"`
	Created      time.Time `db:"created"`
	LastLogin    time.Time `db:"last_login"`
}

func (sd *storedDevice) toDevice() *db.Device {
	return &db.Device{
		ID:           sd.ID,
		Name:         sd.Name,
		PasswordHash: sd.PasswordHash,
		Created:      sd.Created,
		LastLogin:    sd.LastLogin,
	}
}

func toDevices(sds []storedDevice) []db.Device {
	devices := make([]db.Device, 0, len(sds))
	for _, sd := range sds {
		devices = append(devices, *sd.toDevice())
	}
	return devices
}

func (c *conn) GetDevices(accountID int) ([]db.Device, error) {
	sds := []storedDevice{}
	query := "SELECT * FROM devices WHERE account_id = $1"
	err := c.Select(&sds, query, accountID)
	return toDevices(sds), err
}
