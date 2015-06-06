package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func New(addrs []config.PostgresNode, config *config.Postgres) (*db.Conn, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", config.Username, config.Password, config.Name, config.SSLMode)
	postgresConn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	conn := &db.Conn{
		Closer: postgresConn,
	}
	return conn, nil

	// session, err := createSession(addrs, config)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// accountsConn, err := accounts.New(session, config)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// roomsConn, err := rooms.New(session, config)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// conn := &db.Conn{
	// 	Accounts: accountsConn,
	// 	Rooms:    roomsConn,
	// 	Closer:   session,
	// }
	// return conn, nil
}

//
// func createSession(addrs []config.MongoNode, config *config.Mongo) (*mgo.Session, error) {
// 	dialInfo := mgo.DialInfo{
// 		Addrs:     getAddrsAsStrings(addrs),
// 		Database:  config.Name,
// 		Username:  config.Username,
// 		Password:  config.Password,
// 		Mechanism: "SCRAM-SHA-1",
// 		Direct:    false,
// 		Timeout:   10 * time.Second,
// 	}
//
// 	session, err := mgo.DialWithInfo(&dialInfo)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	session.SetMode(mgo.Strong, false)
// 	session.SetSafe(&mgo.Safe{
// 		WMode:    config.WriteParams.Mode,
// 		WTimeout: int(config.WriteParams.Timeout.Seconds()),
// 		J:        config.WriteParams.Journaling})
// 	return session, nil
// }
//
// func getAddrsAsStrings(addrs []config.MongoNode) []string {
// 	addrsStr := make([]string, 0, len(addrs))
// 	for _, addr := range addrs {
// 		addr := fmt.Sprintf("%s:%d", addr.IP.String(), addr.DBPort)
// 		addrsStr = append(addrsStr, addr)
// 	}
// 	return addrsStr
// }
