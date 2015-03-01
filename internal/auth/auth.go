package auth

import (
	"encoding/hex"
	"errors"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

const (
	duration = time.Hour
)

var (
	signatureMethod  = jwt.SigningMethodHS256
	signatureKey     []byte
	signatureKeyOnce sync.Once
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
	signatureKeyOnce.Do(loadSignatureKey)

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
	signatureKeyOnce.Do(loadSignatureKey)

	token := jwt.New(signatureMethod)
	token.Claims["id"] = a.Id
	token.Claims["expires"] = a.Expires.Unix()
	return token.SignedString(signatureKey)
}

func loadSignatureKey() {
	var err error
	signatureKeyHex := viper.GetStringMapString("auth")["signatureKey"]
	signatureKey, err = hex.DecodeString(signatureKeyHex)
	if err != nil {
		panic(err)
	} else if len(signatureKey) == 0 {
		panic("auth: invalid signature key with length 0")
	}
}
