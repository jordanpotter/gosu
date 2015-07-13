package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jordanpotter/gosu/server/internal/db"
	"github.com/lib/pq"
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
