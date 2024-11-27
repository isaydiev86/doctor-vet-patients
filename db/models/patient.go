package models

// Patient Основная инфа пациента
type Patient struct {
	ID         int64  `db:"id"`
	Fio        string `db:"fio"`
	Phone      string `db:"phone"`
	Address    string `db:"address"`
	Animal     string `db:"animal"`
	Name       string `db:"name"`
	Breed      string `db:"breed"`
	Gender     string `db:"gender"`
	IsNeutered bool   `db:"is_neutered"`
}
