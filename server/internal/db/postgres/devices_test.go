package postgres

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestDeviceCreation(t *testing.T) {
	email := fmt.Sprintf("test-%d@email.com", rand.Uint32())
	account, err := dbConn.CreateAccount(email)
	if err != nil {
		t.Errorf("Unexpected error during account creation: %v", err)
	}

	deviceName := "device-name"
	devicePassword := []byte("device-password")
	device, err := dbConn.CreateDevice(account.ID, deviceName, devicePassword)
	if err != nil {
		t.Errorf("Unexpected error during device creation: %v", err)
	} else if device.Name != deviceName {
		t.Errorf("Mismatched device name, %s != %s", device.Name, deviceName)
	} else if device.Created.IsZero() {
		t.Errorf("Invalid timestamp, %v", device.Created)
	}

	allDevices, err := dbConn.GetDevicesByAccount(account.ID)
	if err != nil {
		t.Errorf("Unexpected error during devices retrieval: %v", err)
	} else if len(allDevices) != 1 {
		t.Errorf("Too many devices for account: %d", len(allDevices))
	} else if !reflect.DeepEqual(device, allDevices[0]) {
		t.Errorf("Devices are not equal, %v != %v", device, allDevices[0])
	}
}

func TestDeviceDeletion(t *testing.T) {
	email := fmt.Sprintf("test-%d@email.com", rand.Uint32())
	account, err := dbConn.CreateAccount(email)
	if err != nil {
		t.Errorf("Unexpected error during account creation: %v", err)
	}

	deviceName := "device-name"
	devicePassword := []byte("device-password")
	device, err := dbConn.CreateDevice(account.ID, deviceName, devicePassword)
	if err != nil {
		t.Errorf("Unexpected error during device creation: %v", err)
	}

	err = dbConn.DeleteDeviceForAccount(device.ID, account.ID)
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
		t.Errorf("Unexpected error during devices retrieval: %v", err)
	} else if len(allDevices) != 0 {
		t.Errorf("Too many devices for account: %d", len(allDevices))
	}
}
