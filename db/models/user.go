package models

type User struct {
	ID     int64  `db:"id"`
	UserID string `db:"user_id"` // айди из keycloak
	Fio    string `db:"fio"`
	Role   string `db:"role"` // роль задается в keycloak
}
