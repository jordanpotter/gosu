package postgres

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func createTestDevice(t *testing.T, account db.Account) db.Device {
	name := "device-name"
	passwordHash := []byte("device-password")
	device, err := dbConn.CreateDevice(account.ID, name, passwordHash)
	if err != nil {
		t.Fatalf("Unexpected error during device creation: %v", err)
	} else if device.Name != name {
		t.Errorf("Mismatched device name: %s != %s", device.Name, name)
	} else if !bytes.Equal(device.PasswordHash, passwordHash) {
		t.Errorf("Mismatched password hash: %v != %v", device.PasswordHash, passwordHash)
	} else if device.Created.IsZero() {
		t.Errorf("Invalid timestamp: %v", device.Created)
	}
	return device
}

func TestDeviceCreation(t *testing.T) {
	account := createTestAccount(t)
	device := createTestDevice(t, account)
	allDevices, err := dbConn.GetDevicesByAccount(account.ID)
	if err != nil {
		t.Fatalf("Unexpected error during devices retrieval: %v", err)
	} else if len(allDevices) != 1 {
		t.Errorf("Too many devices for account: %d", len(allDevices))
	} else if !reflect.DeepEqual(device, allDevices[0]) {
		t.Errorf("Devices are not equal: %v != %v", device, allDevices[0])
	}
}

func TestDeviceDeletion(t *testing.T) {
	account := createTestAccount(t)
	device := createTestDevice(t, account)

	err := dbConn.DeleteDeviceForAccount(device.ID, account.ID)
	if err != nil {
		t.Errorf("Unexpected error during device deletion: %v", err)
	}

	// Should allow repeated deletion
	err = dbConn.DeleteDeviceForAccount(device.ID, account.ID)
	if err != nil {
		t.Errorf("Unexpected error during device deletion: %v", err)
	}

	allDevices, err := dbConn.GetDevicesByAccount(account.ID)
	if err != nil {
		t.Fatalf("Unexpected error during devices retrieval: %v", err)
	} else if len(allDevices) != 0 {
		t.Errorf("Too many devices for account: %d", len(allDevices))
	}
}
