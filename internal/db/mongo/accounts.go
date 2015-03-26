package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/internal/auth/password"
	"github.com/JordanPotter/gosu-server/internal/db"
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

func (c *conn) ensureAccountIndices() error {
	emailIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true,
		Sparse:     false,
	}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	return col.EnsureIndex(emailIndex)
}

func (c *conn) CreateAccount(email, deviceName, devicePassword string) error {
	dpHash, err := password.ComputeHash(devicePassword)
	if err != nil {
		return err
	}

	findBson := bson.M{"email": email}
	deviceBson := bson.M{"name": deviceName, "passwordHash": dpHash, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"devices": deviceBson}}
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	_, err = col.Upsert(findBson, dataBson)
	return err
}

func (c *conn) GetAccount(email string) (*db.Account, error) {
	var sa storedAccount
	col := c.session.DB(c.config.Name).C(c.config.Collections.Accounts)
	err := col.Find(bson.M{"email": email}).One(&sa)
	if err == mgo.ErrNotFound {
		return nil, db.NotFoundError
	} else if err != nil {
		return nil, err
	}
	return sa.toAccount(), nil
}
