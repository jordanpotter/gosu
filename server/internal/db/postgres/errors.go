package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jordanpotter/gosu/Godeps/_workspace/src/github.com/lib/pq"
	"github.com/jordanpotter/gosu/server/internal/db"
)

func convertError(err error) error {
	if err == sql.ErrNoRows {
		return db.NotFoundError
	}

	pgErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	switch pgErr.Code {
	case "23505":
		return db.DuplicateError
	}

	fmt.Println("TODO: special case pq error:", pgErr.Code)
	return err
}
