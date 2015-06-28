package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedChannel struct {
	ID      int       `db:"id"`
	RoomID  int       `db:"room_id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (sc storedChannel) toChannel() db.Channel {
	return db.Channel{
		ID:      sc.ID,
		Name:    sc.Name,
		Created: sc.Created,
	}
}

func toChannels(scs []storedChannel) []db.Channel {
	channels := make([]db.Channel, 0, len(scs))
	for _, sc := range scs {
		channels = append(channels, sc.toChannel())
	}
	return channels
}

func (c *conn) CreateChannel(roomID int, name string) (db.Channel, error) {
	sc := storedChannel{}
	insertChannel := "INSERT INTO channels (room_id, name, created) VALUES ($1, $2, $3) RETURNING *"
	err := c.Get(&sc, insertChannel, roomID, name, time.Now())
	return sc.toChannel(), convertError(err)
}

func (c *conn) GetChannelsByRoom(roomID int) ([]db.Channel, error) {
	scs := []storedChannel{}
	selectChannels := "SELECT * FROM channels WHERE room_id=$1"
	err := c.Select(&scs, selectChannels, roomID)
	return toChannels(scs), convertError(err)
}

func (c *conn) DeleteChannelForRoom(id, roomID int) error {
	tx, err := c.Beginx()
	if err != nil {
		return convertError(err)
	}

	var count int
	numMembersInChannel := "SELECT COUNT(*) FROM members WHERE channel_id=$1 AND room_id=$2"
	err = tx.Get(&count, numMembersInChannel, id, roomID)
	if err != nil {
		tx.Rollback()
		return convertError(err)
	} else if count > 0 {
		return db.NotEmptyError
	}

	deleteChannel := "DELETE FROM channels WHERE id=$1 AND room_id=$2"
	_, err = tx.Exec(deleteChannel, id, roomID)
	if err != nil {
		tx.Rollback()
		return convertError(err)
	}
	return tx.Commit()
}
