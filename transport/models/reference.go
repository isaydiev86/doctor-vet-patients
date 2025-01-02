package models

type RefResponse struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}

type Symptoms struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Preparations struct {
	ID       int64   `json:"id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Dose     float64 `json:"dose" validate:"required"`
	Course   string  `json:"course" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Option   string  `json:"option" validate:"required"`
}
