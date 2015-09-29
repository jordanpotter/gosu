package token

import (
	"errors"
	"time"

	jwt "github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/dgrijalva/jwt-go"
)

const (
	accountIDKey    = "accountID"
	roomIDKey       = "roomID"
	roomMemberIDKey = "roomMemberID"
	roomAdminKey    = "roomAdmin"
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
	ID int
}

type Room struct {
	ID       int
	MemberID int
	Admin    bool
}

func NewFactory(signatureKey []byte, duration time.Duration) *Factory {
	return &Factory{signatureKey, duration}
}

func (f *Factory) New() *Token {
	return &Token{Expires: f.getExpirationTime()}
}

func (f *Factory) Extend(t *Token) {
	t.Expires = f.getExpirationTime()
}

func (f *Factory) getExpirationTime() time.Time {
	return time.Now().Add(f.duration).UTC()
}

func (f *Factory) Encrypt(t *Token) (string, error) {
	jwt := jwt.New(jwt.SigningMethodHS256)
	jwt.Claims[accountIDKey] = t.Account.ID
	jwt.Claims[roomIDKey] = t.Room.ID
	jwt.Claims[roomMemberIDKey] = t.Room.MemberID
	jwt.Claims[roomAdminKey] = t.Room.Admin
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

	return &Token{
		Account: getAccountFromClaims(t.Claims),
		Room:    getRoomFromClaims(t.Claims),
		Expires: getExpiresFromClaims(t.Claims),
	}, nil
}

func getAccountFromClaims(claims map[string]interface{}) Account {
	accountIDFloat := claims[accountIDKey].(float64)
	return Account{
		ID: int(accountIDFloat),
	}
}

func getRoomFromClaims(claims map[string]interface{}) Room {
	roomIDFloat, _ := claims[roomIDKey].(float64)
	roomMemberIDFloat, _ := claims[roomMemberIDKey].(float64)
	roomAdmin, _ := claims[roomAdminKey].(bool)
	return Room{
		ID:       int(roomIDFloat),
		MemberID: int(roomMemberIDFloat),
		Admin:    roomAdmin,
	}
}

func getExpiresFromClaims(claims map[string]interface{}) time.Time {
	expiresUnixFloat := claims[expiresKey].(float64)
	expiresUnix := int64(expiresUnixFloat)
	return time.Unix(expiresUnix, 0).UTC()
}
