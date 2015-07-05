package password

import "testing"

func TestComputeHash(t *testing.T) {
	_, err := ComputeHash("")
	if err == nil {
		t.Error("Empty password should return an error")
	}
}

func TestMatchesHash(t *testing.T) {
	password := "P4ssW0rd"

	matches := MatchesHash(password, []byte{})
	if matches {
		t.Error("Password should not match empty hash")
	}

	hash, err := ComputeHash(password)
	if err != nil {
		t.Errorf("Unexpected hashing error: %v", err)
	}

	matches = MatchesHash(password, hash)
	if !matches {
		t.Error("Password does not match hash")
	}
}
