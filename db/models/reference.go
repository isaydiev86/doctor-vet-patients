package models

import "database/sql"

type Reference struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Type sql.NullString `db:"type"`
}

type Symptoms struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
}

type Preparations struct {
	ID       int64           `db:"id"`
	Name     sql.NullString  `db:"name"`
	Dose     sql.NullFloat64 `db:"dose"`
	Course   sql.NullString  `db:"course"`
	Category sql.NullString  `db:"category"`
	Option   sql.NullString  `db:"option"`
}

type PreparationsAdd struct {
	Name     sql.NullString  `db:"name"`
	Dose     sql.NullFloat64 `db:"dose"`
	Course   sql.NullString  `db:"course"`
	Category sql.NullString  `db:"category"`
	Option   sql.NullString  `db:"option"`
}

type NameRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type PreparationsToSymptoms struct {
	ID       int64           `db:"id"`
	Name     sql.NullString  `db:"name"`
	Dose     sql.NullFloat64 `db:"dose"`
	Course   sql.NullString  `db:"course"`
	Category sql.NullString  `db:"category"`
	Option   sql.NullString  `db:"option"`

	Similar []NameRow `db:"similar"` // Список остальных препаратов категории
}
