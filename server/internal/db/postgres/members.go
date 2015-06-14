package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedMember struct {
	ID        int       `db:"id"`
	AccountID int       `db:"account_id"`
	RoomID    int       `db:"room_id"`
	ChannelID int       `db:"channel_id"`
	Name      string    `db:"name"`
	Admin     bool      `db:"admin"`
	Banned    bool      `db:"banned"`
	Created   time.Time `db:"created"`
	LastLogin time.Time `db:"last_login"`
}

func (sm *storedMember) toMember() *db.Member {
	return &db.Member{
		ID:        sm.ID,
		AccountID: sm.AccountID,
		RoomID:    sm.RoomID,
		ChannelID: sm.ChannelID,
		Name:      sm.Name,
		Admin:     sm.Admin,
		Banned:    sm.Banned,
		Created:   sm.Created,
		LastLogin: sm.LastLogin,
	}
}

func toMembers(sms []storedMember) []db.Member {
	members := make([]db.Member, 0, len(sms))
	for _, sm := range sms {
		members = append(members, *sm.toMember())
	}
	return members
}

func (c *conn) CreateMember(accountID, roomID int, name string) (*db.Member, error) {
	sm := new(storedMember)
	insertMember := "INSERT INTO members (account_id, room_id, name, admin, banned, created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	err := c.Get(sm, insertMember, accountID, roomID, name, false, false, time.Now())
	return sm.toMember(), err
}

func (c *conn) GetMember(id int) (*db.Member, error) {
	sm := new(storedMember)
	selectMember := "SELECT * FROM members WHERE id=$1 LIMIT 1"
	err := c.Get(sm, selectMember, id)
	return sm.toMember(), err
}

func (c *conn) GetMembersByAccount(accountID int) ([]db.Member, error) {
	sms := []storedMember{}
	selectMembers := "SELECT * FROM members WHERE account_id=$1"
	err := c.Select(&sms, selectMembers, accountID)
	return toMembers(sms), err
}

func (c *conn) GetMembersByRoom(roomID int) ([]db.Member, error) {
	sms := []storedMember{}
	selectMembers := "SELECT * FROM members WHERE room_id=$1"
	err := c.Select(&sms, selectMembers, roomID)
	return toMembers(sms), err
}

func (c *conn) SetMemberAdmin(id int, admin bool) (*db.Member, error) {
	sm := new(storedMember)
	updateMember := "UPDATE members SET admin=$1 WHERE id=$2"
	err := c.Get(sm, updateMember, admin, id)
	return sm.toMember(), err
}

func (c *conn) SetMemberBanned(id int, banned bool) (*db.Member, error) {
	sm := new(storedMember)
	updateMember := "UPDATE members SET banned=$1 WHERE id=$2"
	err := c.Get(sm, updateMember, banned, id)
	return sm.toMember(), err
}

func (c *conn) DeleteMember(id int) error {
	deleteMember := "DELETE FROM members WHERE id=$1"
	_, err := c.Exec(deleteMember, id)
	return err
}
