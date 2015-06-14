package postgres

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type storedChannel struct {
	ID      int       `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (sc *storedChannel) toChannel() *db.Channel {
	return &db.Channel{
		ID:      sc.ID,
		Name:    sc.Name,
		Created: sc.Created,
	}
}

func toChannels(scs []storedChannel) []db.Channel {
	channels := make([]db.Channel, 0, len(scs))
	for _, sc := range scs {
		channels = append(channels, *sc.toChannel())
	}
	return channels
}

func (c *conn) CreateChannel(roomID int, name string) (*db.Channel, error) {
	sc := new(storedChannel)
	insertChannel := "INSERT INTO channels (room_id, name, created) VALUES ($1, $2, $3) RETURNING *"
	err := c.Get(sc, insertChannel, roomID, name, time.Now())
	return sc.toChannel(), err
}

func (c *conn) GetChannel(id int) (*db.Channel, error) {
	sc := new(storedChannel)
	selectChannel := "SELECT * FROM channels WHERE id=$1 LIMIT 1"
	err := c.Get(sc, selectChannel, id)
	return sc.toChannel(), err
}

func (c *conn) GetChannelsByRoom(roomID int) ([]db.Channel, error) {
	scs := []storedChannel{}
	selectChannels := "SELECT * FROM channels WHERE room_id=$1"
	err := c.Select(&scs, selectChannels, roomID)
	return toChannels(scs), err
}

func (c *conn) DeleteChannel(id int) error {
	tx, err := c.Beginx()
	if err != nil {
		return err
	}

	var count int
	numMembersInChannel := "SELECT COUNT(*) FROM members WHERE channel_id=$1"
	err = tx.Get(&count, numMembersInChannel, id)
	if err != nil {
		tx.Rollback()
		return err
	} else if count > 0 {
		return db.NotEmptyError
	}

	deleteChannel := "DELETE FROM channels WHERE id=$1"
	_, err = tx.Exec(deleteChannel, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
