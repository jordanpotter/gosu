package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const duration = time.Hour

var (
	signatureMethod = jwt.SigningMethodHS256
	signatureKey    = []byte{78, 217, 118, 44, 83, 152, 15, 40, 165, 52, 191, 235, 169, 203, 8, 138}
)

type Auth struct {
	Id      string
	Expires time.Time
}

func New(id string) *Auth {
	expires := time.Now().Add(duration)
	return &Auth{id, expires}
}

func Decrypt(str string) (*Auth, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return signatureKey, nil
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
	token := jwt.New(signatureMethod)
	token.Claims["id"] = a.Id
	token.Claims["expires"] = a.Expires.Unix()
	return token.SignedString(signatureKey)
}
