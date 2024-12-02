package models

type RefResponse struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}
