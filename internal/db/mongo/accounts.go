package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/JordanPotter/gosu-server/internal/auth/password"
	"github.com/JordanPotter/gosu-server/internal/db"
)

const (
	accountsCollectionName = "accounts"
)

type storedAccount struct {
	Id      bson.ObjectId  `bson:"_id,omitempty"`
	Email   string         `bson:"email"`
	Clients []storedClient `bson:"clients"`
}

type storedClient struct {
	Name         string    `bson:"name"`
	PasswordHash []byte    `bson:"passwordHash"`
	Created      time.Time `bson:"created"`
}

func (sa *storedAccount) toAccount() *db.Account {
	clients := make([]db.Client, len(sa.Clients))
	for _, sClient := range sa.Clients {
		clients = append(clients, *sClient.toClient())
	}

	return &db.Account{
		Id:      sa.Id.Hex(),
		Email:   sa.Email,
		Clients: clients,
	}
}

func (sc *storedClient) toClient() *db.Client {
	return &db.Client{
		Name:         sc.Name,
		PasswordHash: sc.PasswordHash,
		Created:      sc.Created,
	}
}

func (c *conn) CreateAccount(email, clientName, clientPassword string) error {
	cpHash, err := password.ComputeHash(clientPassword)
	if err != nil {
		return err
	}

	findBson := bson.M{"email": email}
	clientBson := bson.M{"name": clientName, "passwordHash": cpHash, "created": time.Now()}
	dataBson := bson.M{"$push": bson.M{"clients": clientBson}}
	col := c.session.DB(databaseName).C(accountsCollectionName)
	_, err = col.Upsert(findBson, dataBson)
	return err
}

func (c *conn) GetAccount(email string) (*db.Account, error) {
	var sa storedAccount
	col := c.session.DB(databaseName).C(accountsCollectionName)
	err := col.Find(bson.M{"email": email}).One(&sa)
	if err == mgo.ErrNotFound {
		return nil, db.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return sa.toAccount(), nil
}
