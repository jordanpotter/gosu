package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/auth/password"
	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedRoom struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	PasswordHash []byte    `db:"password_hash"`
	Created      time.Time `db:"created"`
}

func (sr *storedRoom) toRoom() *db.Room {
	return &db.Room{
		ID:           sr.ID,
		Name:         sr.Name,
		PasswordHash: sr.PasswordHash,
		Created:      sr.Created,
	}
}

func (c *conn) CreateRoom(name, pwd string, adminAccountID int, adminName string) (*db.Room, error) {
	pwdHash, err := password.ComputeHash(pwd)
	if err != nil {
		return nil, err
	}

	tx, err := c.Beginx()
	if err != nil {
		return nil, err
	}

	sr := new(storedRoom)
	insertRoom := "INSERT INTO rooms (name, password_hash, created) VALUES ($1, $2, $3) RETURNING *"
	err = tx.Get(sr, insertRoom, name, pwdHash, time.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// TODO: insert member
	err = tx.Commit()
	return sr.toRoom(), err
}

func (c *conn) GetRoom(id int) (*db.Room, error) {
	sr := new(storedRoom)
	selectRoom := "SELECT * FROM rooms WHERE id=$1 LIMIT 1"
	err := c.Get(sr, selectRoom, id)
	return sr.toRoom(), err
}

func (c *conn) GetRoomByName(name string) (*db.Room, error) {
	sr := new(storedRoom)
	selectRoom := "SELECT * FROM rooms WHERE name=$1 LIMIT 1"
	err := c.Get(sr, selectRoom, name)
	return sr.toRoom(), err
}
