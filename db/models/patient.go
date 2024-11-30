package models

import "database/sql"

// Patient Основная инфа пациента
type Patient struct {
	ID         int64           `db:"id"`
	Age        sql.NullFloat64 `db:"age"`
	Fio        sql.NullString  `db:"fio"`
	Phone      sql.NullString  `db:"phone"`
	Address    sql.NullString  `db:"address"`
	Animal     sql.NullString  `db:"animal"`
	Name       sql.NullString  `db:"name"`
	Breed      sql.NullString  `db:"breed"`
	Gender     sql.NullString  `db:"gender"`
	IsNeutered bool            `db:"is_neutered"`
}
