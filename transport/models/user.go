package models

type User struct {
	ID     int64  `json:"id" validate:"required"`
	UserID string `json:"userId" validate:"required"`
	Fio    string `json:"fio" validate:"required"`
	Role   string `json:"role" validate:"required"`
}
