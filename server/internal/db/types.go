package db

import "time"

type Account struct {
	ID      int
	Email   string
	Created time.Time
}

type Device struct {
	ID           int
	Name         string
	PasswordHash []byte
	Created      time.Time
}

type Room struct {
	ID           int
	Name         string
	PasswordHash []byte
	Created      time.Time
}

type Channel struct {
	ID      int
	Name    string
	Created time.Time
}

type Member struct {
	ID        int
	AccountID int
	RoomID    int
	ChannelID int
	Name      string
	Admin     bool
	Banned    bool
	Created   time.Time
}
