package models

import "database/sql"

type Reference struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Type sql.NullString `db:"type"`
}
