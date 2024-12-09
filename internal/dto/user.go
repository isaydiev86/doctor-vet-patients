package dto

type User struct {
	ID     int64
	UserID string
	Fio    string
	Role   string
}

type UserFilters struct {
	Role string
}
