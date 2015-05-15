package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedAccount struct {
	ID      bson.ObjectId  `bson:"_id,omitempty"`
	Email   string         `bson:"email"`
	Devices []storedDevice `bson:"devices"`
	Created time.Time      `bson:"created"`
}

type storedDevice struct {
	ID           bson.ObjectId `bson:"id"`
	Name         string        `bson:"name"`
	PasswordHash []byte        `bson:"passwordHash"`
	Created      time.Time     `bson:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	devices := make([]db.Device, len(sa.Devices))
	for _, sDevice := range sa.Devices {
		devices = append(devices, *sDevice.toDevice())
	}

	return &db.Account{
		ID:      sa.ID.Hex(),
		Email:   sa.Email,
		Devices: devices,
		Created: sa.Created,
	}
}

func (sd *storedDevice) toDevice() *db.Device {
	return &db.Device{
		ID:           sd.ID.Hex(),
		Name:         sd.Name,
		PasswordHash: sd.PasswordHash,
		Created:      sd.Created,
	}
}
