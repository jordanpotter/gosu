package db

import (
	"golang.org/x/crypto/bcrypt"
)

func ComputePasswordHash(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func DoesPasswordMatchHash(p string, h []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, []byte(p))
	return err == nil
}
