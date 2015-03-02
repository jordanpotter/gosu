package db

import (
	"code.google.com/p/go-uuid/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword() string {
	return uuid.New()
}

func ComputePasswordHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func DoesPasswordMatchHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
