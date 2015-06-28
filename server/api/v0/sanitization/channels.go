package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Channel struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func ToChannel(dbChannel db.Channel) Channel {
	return Channel{
		ID:      dbChannel.ID,
		Name:    dbChannel.Name,
		Created: dbChannel.Created,
	}
}

func ToChannels(dbChannels []db.Channel) []Channel {
	channels := make([]Channel, 0, len(dbChannels))
	for _, dbChannel := range dbChannels {
		channels = append(channels, ToChannel(dbChannel))
	}
	return channels
}
