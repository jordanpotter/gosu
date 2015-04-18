package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedAccount struct {
	Id      bson.ObjectId  `bson:"_id,omitempty"`
	Email   string         `bson:"email"`
	Devices []storedDevice `bson:"devices"`
}

type storedDevice struct {
	Name         string    `bson:"name"`
	PasswordHash []byte    `bson:"passwordHash"`
	Created      time.Time `bson:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	devices := make([]db.Device, len(sa.Devices))
	for _, sDevice := range sa.Devices {
		devices = append(devices, *sDevice.toDevice())
	}

	return &db.Account{
		Id:      sa.Id.Hex(),
		Email:   sa.Email,
		Devices: devices,
	}
}

func (sd *storedDevice) toDevice() *db.Device {
	return &db.Device{
		Name:         sd.Name,
		PasswordHash: sd.PasswordHash,
		Created:      sd.Created,
	}
}
