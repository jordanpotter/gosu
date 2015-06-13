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
	LastLogin    time.Time
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
	Cretaed time.Time
}

type Member struct {
	ID        int
	RoomID    int
	AccountID int
	ChannelID int
	Name      string
	Admin     bool
	Banned    bool
	Created   time.Time
	LastLogin time.Time
}
