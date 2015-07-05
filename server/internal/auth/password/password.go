package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ComputeHash(password string) ([]byte, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func MatchesHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
