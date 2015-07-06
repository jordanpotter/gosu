package postgres

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net"
	"time"

	"github.com/jordanpotter/gosu/server/internal/config"
	"github.com/jordanpotter/gosu/server/internal/db"
)

const (
	testPostgresIP         = "127.0.0.1"
	testPostgresPort       = 5432
	testPostgresConfigPath = "../../../../conf/db/postgres.json"
	testMigrationsPath     = "../../../../conf/db/migrations"
)

var dbConn db.Conn

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	postgresAddr := config.PostgresNode{
		IP:     net.ParseIP(testPostgresIP),
		DBPort: testPostgresPort,
	}
	postgresAddrs := []config.PostgresNode{postgresAddr}

	postgresConfigBytes, err := ioutil.ReadFile(testPostgresConfigPath)
	if err != nil {
		panic(err)
	}

	var postgresConfig config.Postgres
	err = json.Unmarshal(postgresConfigBytes, &postgresConfig)
	if err != nil {
		panic(err)
	}

	dbConn, err = New(postgresAddrs, &postgresConfig, testMigrationsPath)
	if err != nil {
		panic(err)
	}
}
