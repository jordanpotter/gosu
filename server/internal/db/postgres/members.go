package postgres

import (
	"database/sql"
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedMember struct {
	ID        int           `db:"id"`
	AccountID int           `db:"account_id"`
	RoomID    int           `db:"room_id"`
	ChannelID sql.NullInt64 `db:"channel_id"`
	Name      string        `db:"name"`
	Admin     bool          `db:"admin"`
	Banned    bool          `db:"banned"`
	Created   time.Time     `db:"created"`
	LastLogin *time.Time    `db:"last_login"`
}

func (sm *storedMember) toMember() *db.Member {
	member := &db.Member{
		ID:        sm.ID,
		AccountID: sm.AccountID,
		RoomID:    sm.RoomID,
		Name:      sm.Name,
		Admin:     sm.Admin,
		Banned:    sm.Banned,
		Created:   sm.Created,
	}
	if sm.ChannelID.Valid {
		member.ChannelID = int(sm.ChannelID.Int64)
	}
	if sm.LastLogin != nil {
		member.LastLogin = *sm.LastLogin
	}
	return member
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
	return sm.toMember(), convertError(err)
}

func (c *conn) GetMemberByAccountAndRoom(accountID, roomID int) (*db.Member, error) {
	sm := new(storedMember)
	selectMember := "SELECT * FROM members WHERE account_id=$1 AND room_id=$2 LIMIT 1"
	err := c.Get(sm, selectMember, accountID, roomID)
	return sm.toMember(), convertError(err)
}

func (c *conn) GetMembersByAccount(accountID int) ([]db.Member, error) {
	sms := []storedMember{}
	selectMembers := "SELECT * FROM members WHERE account_id=$1"
	err := c.Select(&sms, selectMembers, accountID)
	return toMembers(sms), convertError(err)
}

func (c *conn) GetMembersByRoom(roomID int) ([]db.Member, error) {
	sms := []storedMember{}
	selectMembers := "SELECT * FROM members WHERE room_id=$1"
	err := c.Select(&sms, selectMembers, roomID)
	return toMembers(sms), convertError(err)
}

func (c *conn) SetMemberAdminForRoom(id, roomID int, admin bool) (*db.Member, error) {
	sm := new(storedMember)
	updateMember := "UPDATE members SET admin=$1 WHERE id=$2 AND room_id=$3 RETURNING *"
	err := c.Get(sm, updateMember, admin, id, roomID)
	return sm.toMember(), convertError(err)
}

func (c *conn) SetMemberBannedForRoom(id, roomID int, banned bool) (*db.Member, error) {
	sm := new(storedMember)
	updateMember := "UPDATE members SET banned=$1 WHERE id=$2 AND room_id=$3 RETURNING *"
	err := c.Get(sm, updateMember, banned, id, roomID)
	return sm.toMember(), convertError(err)
}

func (c *conn) DeleteMemberForAccount(id, accountID int) error {
	deleteMember := "DELETE FROM members WHERE id=$1 AND account_id=$2"
	_, err := c.Exec(deleteMember, id, accountID)
	return convertError(err)
}

func (c *conn) DeleteMemberForRoom(id, roomID int) error {
	deleteMember := "DELETE FROM members WHERE id=$1 AND room_id=$2"
	_, err := c.Exec(deleteMember, id, roomID)
	return convertError(err)
}

func (c *conn) DeleteMemberForAccountAndRoom(accountID, roomID int) error {
	deleteMember := "DELETE FROM members WHERE account_id=$1 AND room_id=$2"
	_, err := c.Exec(deleteMember, accountID, roomID)
	return convertError(err)
}
