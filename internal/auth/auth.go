package auth

import (
	"encoding/hex"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	duration = time.Hour
)

var (
	signingMethod = jwt.SigningMethodHS256
	signingString []byte
)

type Auth struct {
	Id      string
	Expires time.Time
}

func init() {
	var err error
	signingString, err = hex.DecodeString("DEADBEEF") // TODO: read from environment variable
	if err != nil {
		panic(err)
	}
}

func New(id string) *Auth {
	expires := time.Now().Add(duration)
	return &Auth{id, expires}
}

func Decrypt(val string) (*Auth, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return signingString, nil
	})
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("auth: invalid token")
	}

	id := token.Claims["id"].(string)
	expiresUnix := int64(token.Claims["expires"].(float64))
	expires := time.Unix(expiresUnix, 0)
	return &Auth{id, expires}, nil
}

func (a *Auth) Encrypt() (string, error) {
	token := jwt.New(signingMethod)
	token.Claims["id"] = a.Id
	token.Claims["expires"] = a.Expires.Unix()
	return token.SignedString(signingString)
}
