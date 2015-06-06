package db

// import (
// 	"time"
// )
//
// type RoomsConn interface {
// 	Create(name, password, adminAccountID, adminName string) (*Room, error)
// 	Get(id string) (*Room, error)
// 	GetByName(name string) (*Room, error)
// 	Delete(id string) error
//
// 	AddChannel(id, channelName string) (*Channel, error)
// 	GetChannel(id, channelID string) (*Channel, error)
// 	RemoveChannel(id, channelID string) error
//
// 	AddMember(id, accountID, memberName string) (*Member, error)
// 	GetMember(id, memberID string) (*Member, error)
// 	GetMemberByAccount(id, accountID string) (*Member, error)
// 	SetMemberAdmin(id, memberID string, admin bool) error
// 	SetMemberBanned(id, memberID string, banned bool) error
// 	RemoveMember(id, memberID string) error
// }
//
// type Room struct {
// 	ID           int
// 	Name         string
// 	PasswordHash []byte
// 	Channels     []Channel
// 	Members      []Member
// 	Created      time.Time
// }
//
// type Channel struct {
// 	ID      int
// 	Name    string
// 	Created time.Time
// }
//
// type Member struct {
// 	ID          int
// 	AccountID   int
// 	Name        string
// 	ChannelName string
// 	Admin       bool
// 	Banned      bool
// 	Created     time.Time
// }
