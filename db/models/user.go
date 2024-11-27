package models

type User struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"` // айди из кейклоака
	Fio    string `db:"fio"`
	Phone  string `db:"phone"`
	Role   string `db:"role"` // роль задачется в кейклоаке
}
