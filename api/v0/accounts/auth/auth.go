package auth

import "time"

import jwt "github.com/dgrijalva/jwt-go"

const (
	authTokenDuration = time.Hour
)

var (
	authTokenSigningMethod = jwt.SigningMethodHS256
	authTokenKey           = []byte{}
)

type Auth struct {
	Token   string `json:"token"`
	Expires int64  `json:"expires"`
}

func New(id string) (Auth, error) {
	expires := time.Now().Add(authTokenDuration).Unix()

	token := jwt.New(authTokenSigningMethod)
	token.Claims["id"] = id
	token.Claims["expires"] = expires
	tokenString, err := token.SignedString(authTokenKey)
	return Auth{tokenString, expires}, err
}
