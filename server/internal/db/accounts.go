package db

// import (
// 	"time"
// )
//
// type AccountsConn interface {
// 	Create(email, deviceName, devicePassword string) (*Account, error)
// 	Get(id string) (*Account, error)
// 	GetByEmail(email string) (*Account, error)
// 	Delete(id string) error
// }
//
// type Account struct {
// 	ID      int
// 	Email   string
// 	Devices []Device
// 	Created time.Time
// }
//
// type Device struct {
// 	ID           int
// 	Name         string
// 	PasswordHash []byte
// 	Created      time.Time
// }
