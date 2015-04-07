package token

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Factory struct {
	signatureKey []byte
	duration     time.Duration
}

type Token struct {
	Id      string
	Expires time.Time
}

func NewFactory(signatureKey []byte, duration time.Duration) *Factory {
	return &Factory{signatureKey, duration}
}

func (f *Factory) New(id string) *Token {
	expires := time.Now().Add(f.duration)
	return &Token{id, expires}
}

func (f *Factory) Encrypt(token *Token) (string, error) {
	jwt := jwt.New(jwt.SigningMethodHS256)
	jwt.Claims["id"] = token.Id
	jwt.Claims["expires"] = token.Expires.Unix()
	return jwt.SignedString(f.signatureKey)
}

func (f *Factory) Decrypt(str string) (*Token, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return f.signatureKey, nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("invalid token")
	}

	id := token.Claims["id"].(string)
	expiresUnix := int64(token.Claims["expires"].(float64))
	expires := time.Unix(expiresUnix, 0)
	return &Token{id, expires}, nil
}
