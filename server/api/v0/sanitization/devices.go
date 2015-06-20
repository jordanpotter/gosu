package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Device struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func SanitizeDevice(dbDevice *db.Device) *Device {
	return &Device{
		ID:      dbDevice.ID,
		Name:    dbDevice.Name,
		Created: dbDevice.Created,
	}
}

func SanitizeDevices(dbDevices []db.Device) []Device {
	devices := make([]Device, 0, len(dbDevices))
	for _, dbDevice := range dbDevices {
		devices = append(devices, *SanitizeDevice(&dbDevice))
	}
	return devices
}
