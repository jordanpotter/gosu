package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/events"
)

const (
	AccountDeviceCreatedName = "accountDeviceCreated"
	AccountDeviceDeletedName = "accountDeviceDeleted"
)

type AccountDeviceCreated struct {
	EventName  string    `json:"eventName"`
	AccountID  int       `json:"accountId"`
	DeviceID   int       `json:"deviceId"`
	DeviceName string    `json:"deviceName"`
	Created    time.Time `json:"created"`
	Timestamp  time.Time `json:"timestamp"`
}

func ToAccountDeviceCreated(adc events.AccountDeviceCreated, timestamp time.Time) AccountDeviceCreated {
	return AccountDeviceCreated{
		EventName:  AccountDeviceCreatedName,
		AccountID:  adc.AccountID,
		DeviceID:   adc.DeviceID,
		DeviceName: adc.DeviceName,
		Created:    adc.Created,
		Timestamp:  timestamp,
	}
}

type AccountDeviceDeleted struct {
	EventName string    `json:"eventName"`
	AccountID int       `json:"accountId"`
	DeviceID  int       `json:"deviceId"`
	Timestamp time.Time `json:"timestamp"`
}

func ToAccountDeviceDeleted(add events.AccountDeviceDeleted, timestamp time.Time) AccountDeviceDeleted {
	return AccountDeviceDeleted{
		EventName: AccountDeviceDeletedName,
		AccountID: add.AccountID,
		DeviceID:  add.DeviceID,
		Timestamp: timestamp,
	}
}
