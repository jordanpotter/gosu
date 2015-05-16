package token

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	accountIDKey    = "accountID"
	roomIDKey       = "roomID"
	roomMemberIDKey = "roomMemberID"
	expiresKey      = "expires"
)

type Factory struct {
	signatureKey []byte
	duration     time.Duration
}

type Token struct {
	Account Account
	Room    Room
	Expires time.Time
}

type Account struct {
	ID string
}

type Room struct {
	ID       string
	MemberID string
}

func NewFactory(signatureKey []byte, duration time.Duration) *Factory {
	return &Factory{signatureKey, duration}
}

func (f *Factory) New() *Token {
	expires := time.Now().Add(f.duration).UTC()
	return &Token{Expires: expires}
}

func (f *Factory) Encrypt(t *Token) (string, error) {
	jwt := jwt.New(jwt.SigningMethodHS256)
	jwt.Claims[accountIDKey] = t.Account.ID
	jwt.Claims[roomIDKey] = t.Room.ID
	jwt.Claims[roomMemberIDKey] = t.Room.MemberID
	jwt.Claims[expiresKey] = t.Expires.Unix()
	return jwt.SignedString(f.signatureKey)
}

func (f *Factory) Decrypt(str string) (*Token, error) {
	t, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		return f.signatureKey, nil
	})
	if err != nil {
		return nil, err
	} else if !t.Valid {
		return nil, errors.New("invalid token")
	}

	account := Account{}
	account.ID = t.Claims[accountIDKey].(string)

	room := Room{}
	room.ID, _ = t.Claims[roomIDKey].(string)
	room.MemberID, _ = t.Claims[roomMemberIDKey].(string)

	expiresUnix := int64(t.Claims[expiresKey].(float64))
	expires := time.Unix(expiresUnix, 0).UTC()
	return &Token{account, room, expires}, nil
}
