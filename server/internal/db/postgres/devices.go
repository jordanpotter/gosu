package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedDevice struct {
	ID           int       `db:"id"`
	AccountID    int       `db:"account_id"`
	Name         string    `db:"name"`
	PasswordHash []byte    `db:"password_hash"`
	Created      time.Time `db:"created"`
}

func (sd *storedDevice) toDevice() *db.Device {
	device := &db.Device{
		ID:           sd.ID,
		Name:         sd.Name,
		PasswordHash: sd.PasswordHash,
		Created:      sd.Created,
	}
	return device
}

func toDevices(sds []storedDevice) []db.Device {
	devices := make([]db.Device, 0, len(sds))
	for _, sd := range sds {
		devices = append(devices, *sd.toDevice())
	}
	return devices
}

func (c *conn) CreateDevice(accountID int, deviceName string, devicePasswordHash []byte) (*db.Device, error) {
	sd := new(storedDevice)
	insertDevice := "INSERT INTO devices (account_id, name, password_hash, created) VALUES ($1, $2, $3, $4) RETURNING *"
	err := c.Get(sd, insertDevice, accountID, deviceName, devicePasswordHash, time.Now())
	return sd.toDevice(), convertError(err)
}

func (c *conn) GetDevicesByAccount(accountID int) ([]db.Device, error) {
	sds := []storedDevice{}
	selectDevices := "SELECT * FROM devices WHERE account_id=$1"
	err := c.Select(&sds, selectDevices, accountID)
	return toDevices(sds), convertError(err)
}

func (c *conn) DeleteDeviceForAccount(id, accountID int) error {
	deleteDevice := "DELETE FROM devices WHERE id=$1 AND account_id=$2"
	_, err := c.Exec(deleteDevice, id, accountID)
	return convertError(err)
}
