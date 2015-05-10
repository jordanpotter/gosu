package accounts

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedAccount struct {
	Id          bson.ObjectId      `bson:"_id,omitempty"`
	Email       string             `bson:"email"`
	Devices     []storedDevice     `bson:"devices"`
	Memberships []storedMembership `bson:"memberships"`
	Created     time.Time          `bson:"created"`
}

type storedDevice struct {
	Name         string    `bson:"name"`
	PasswordHash []byte    `bson:"passwordHash"`
	Created      time.Time `bson:"created"`
}

type storedMembership struct {
	RoomId   bson.ObjectId `bson:"roomId"`
	PeerName string        `bson:"peerName"`
	Admin    bool          `bson:"admin"`
	Banned   bool          `bson:"banned"`
	Created  time.Time     `bson:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	devices := make([]db.Device, len(sa.Devices))
	for _, sDevice := range sa.Devices {
		devices = append(devices, *sDevice.toDevice())
	}

	memberships := make([]db.Membership, len(sa.Memberships))
	for _, sMembership := range sa.Memberships {
		memberships = append(memberships, *sMembership.toMembership())
	}

	return &db.Account{
		Id:          sa.Id.Hex(),
		Email:       sa.Email,
		Devices:     devices,
		Memberships: memberships,
		Created:     sa.Created,
	}
}

func (sd *storedDevice) toDevice() *db.Device {
	return &db.Device{
		Name:         sd.Name,
		PasswordHash: sd.PasswordHash,
		Created:      sd.Created,
	}
}

func (sm *storedMembership) toMembership() *db.Membership {
	return &db.Membership{
		RoomId:   sm.RoomId.Hex(),
		PeerName: sm.PeerName,
		Admin:    sm.Admin,
		Banned:   sm.Banned,
		Created:  sm.Created,
	}
}
