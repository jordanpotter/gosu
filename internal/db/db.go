package db

type Conn interface {
	Close()
	CreateAccount(name string, password string) (*Account, error)
	GetAccount(name string) (*Account, error)
}

type Account struct {
	Name     string   `json:"name", bson:"name"`
	Password Password `json:"name", bson:"password"`
}

type Password struct {
	Hash []byte `json:"hash", bson:"hash"`
	Salt []byte `json:"salt", bson:"salt"`
}
