package db

import "time"

type Account struct {
	ID      int       `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

type Device struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"`
	Created      time.Time `json:"created"`
}

type Room struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	PasswordHash []byte    `json:"-"`
	Created      time.Time `json:"created"`
}

type Channel struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type Member struct {
	ID        int       `json:"id"`
	AccountID int       `json:"-"`
	RoomID    int       `json:"roomId"`
	ChannelID int       `json:"channelId"`
	Name      string    `json:"name"`
	Admin     bool      `json:"admin"`
	Banned    bool      `json:"banned"`
	Created   time.Time `json:"created"`
}
