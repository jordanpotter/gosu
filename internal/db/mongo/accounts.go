package mongo

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/internal/db"
)

const (
	accountsCollectionName = "accounts"
)

type storedAccount struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	PasswordHash []byte        `bson:"passwordHash"`
	Created      time.Time     `bson:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	return &db.Account{
		Id:      sa.Id.Hex(),
		Created: sa.Created,
	}
}

func (c *conn) CreateAccount() (*db.Account, string, error) {
	p := db.GeneratePassword()
	pHash, err := db.ComputePasswordHash(p)
	if err != nil {
		return nil, "", err
	}

	sa := storedAccount{
		Id:           bson.NewObjectId(),
		PasswordHash: pHash,
		Created:      bson.Now(),
	}

	col := c.session.DB(databaseName).C(accountsCollectionName)
	err = col.Insert(&sa)
	if err != nil {
		return nil, "", err
	}

	return sa.toAccount(), p, nil
}

func (c *conn) GetAccount(id, password string) (*db.Account, error) {
	var sa storedAccount
	col := c.session.DB(databaseName).C(accountsCollectionName)
	err := col.FindId(bson.ObjectIdHex(id)).One(&sa)
	if err != nil {
		return nil, err
	}

	if !db.DoesPasswordMatchHash(password, sa.PasswordHash) {
		return nil, errors.New("mongo: password does not match hash")
	}

	return sa.toAccount(), nil
}
