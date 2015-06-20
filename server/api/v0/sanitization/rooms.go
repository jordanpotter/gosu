package sanitization

import (
	"time"

	"github.com/jordanpotter/gosu/server/internal/db"
)

type Room struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func SanitizeRoom(dbRoom *db.Room) *Room {
	return &Room{
		ID:      dbRoom.ID,
		Name:    dbRoom.Name,
		Created: dbRoom.Created,
	}
}

func SanitizeRooms(dbRooms []db.Room) []Room {
	rooms := make([]Room, 0, len(dbRooms))
	for _, dbRoom := range dbRooms {
		rooms = append(rooms, *SanitizeRoom(&dbRoom))
	}
	return rooms
}
