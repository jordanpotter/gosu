package token

import (
	"testing"
	"time"
)

const (
	testSignatureKey = "secret signature key"
	testDuration     = time.Hour
	testAccountID    = 17
	testRoomID       = 19
	testRoomMemberID = 103
	testRoomAdmin    = true
)

func TestExpiration(t *testing.T) {
	factory := NewFactory([]byte(testSignatureKey), testDuration)
	token := factory.New()
	expectedExpiration := time.Now().Add(testDuration)
	if token.Expires.After(expectedExpiration) {
		t.Errorf("Token expiration %v should be before %v", token.Expires, expectedExpiration)
	}
}

func TestExtend(t *testing.T) {
	duration := time.Microsecond
	factory := NewFactory([]byte(testSignatureKey), duration)
	token := factory.New()
	firstExpiration := token.Expires
	factory.Extend(token)
	secondExpiration := token.Expires
	if firstExpiration.Equal(secondExpiration) || firstExpiration.After(secondExpiration) {
		t.Errorf("Expiration not extended, %v vs %v", firstExpiration, secondExpiration)
	}
}

func TestEncryption(t *testing.T) {
	factory := NewFactory([]byte(testSignatureKey), testDuration)
	token := factory.New()
	token.Account.ID = testAccountID
	token.Room.ID = testRoomID
	token.Room.MemberID = testRoomMemberID
	token.Room.Admin = testRoomAdmin

	encrypted, err := factory.Encrypt(token)
	if err != nil {
		t.Errorf("Unexpected error during encryption: %v", err)
	}

	decrypted, err := factory.Decrypt(encrypted)
	if err != nil {
		t.Errorf("Unexpected error during decryption: %v", err)
	}

	if decrypted.Account.ID != testAccountID {
		t.Errorf("Mismatch account id, %d != %d", decrypted.Account.ID, testAccountID)
	} else if decrypted.Room.ID != testRoomID {
		t.Errorf("Mismatch room id, %d != %d", decrypted.Room.ID, testRoomID)
	} else if decrypted.Room.MemberID != testRoomMemberID {
		t.Errorf("Mismatch room member id, %d != %d", decrypted.Room.MemberID, testRoomMemberID)
	} else if decrypted.Room.Admin != testRoomAdmin {
		t.Errorf("Mismatch room admin, %t != %t", decrypted.Room.Admin, testRoomAdmin)
	}
}

func TestMissingFields(t *testing.T) {
	factory := NewFactory([]byte(testSignatureKey), testDuration)
	token := factory.New()
	token.Account.ID = testAccountID
	token.Room.ID = testRoomID

	encrypted, err := factory.Encrypt(token)
	if err != nil {
		t.Errorf("Unexpected error during encryption: %v", err)
	}

	decrypted, err := factory.Decrypt(encrypted)
	if err != nil {
		t.Errorf("Unexpected error during decryption: %v", err)
	}

	if decrypted.Account.ID != testAccountID {
		t.Errorf("Mismatch account id, %d != %d", decrypted.Account.ID, testAccountID)
	} else if decrypted.Room.ID != testRoomID {
		t.Errorf("Mismatch room id, %d != %d", decrypted.Room.ID, testRoomID)
	} else if decrypted.Room.MemberID != 0 {
		t.Errorf("Expected empty room member id, received %d", decrypted.Room.MemberID)
	} else if decrypted.Room.Admin {
		t.Errorf("Expected falsy room admin, received %t", decrypted.Room.Admin)
	}
}
