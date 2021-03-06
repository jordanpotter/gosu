package password

import (
	"errors"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/golang.org/x/crypto/bcrypt"
)

func ComputeBcryptHash(password string) ([]byte, error) {
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func MatchesBcryptHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
