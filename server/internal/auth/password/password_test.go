package password

import "testing"

func TestComputeHash(t *testing.T) {
	_, err := ComputeBcryptHash("")
	if err == nil {
		t.Error("Empty password should return an error")
	}
}

func TestMatchesHash(t *testing.T) {
	password := "P4ssW0rd"

	matches := MatchesBcryptHash(password, []byte{})
	if matches {
		t.Error("Password should not match empty hash")
	}

	hash, err := ComputeBcryptHash(password)
	if err != nil {
		t.Errorf("Unexpected hashing error: %v", err)
	}

	matches = MatchesBcryptHash(password, hash)
	if !matches {
		t.Error("Password does not match hash")
	}
}
