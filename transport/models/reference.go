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

type NameAdd struct {
	Name string `json:"name" validate:"required"`
}

type NameResponse struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RelationSymptomWithPreparation struct {
	SymptomID     int64 `json:"symptomId" validate:"required"`
	PreparationID int64 `json:"preparationId" validate:"required"`
}

type Preparations struct {
	ID       int64   `json:"id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Dose     float64 `json:"dose" validate:"required"`
	Course   string  `json:"course" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Option   string  `json:"option" validate:"required"`
}

type PreparationsAdd struct {
	Name     string  `json:"name" validate:"required"`
	Dose     float64 `json:"dose" validate:"required"`
	Course   string  `json:"course" validate:"required"`
	Category string  `json:"category" validate:"required"`
	Option   string  `json:"option" validate:"required"`
}

type PreparationsToSymptoms struct {
	ID       int64          `json:"id" validate:"required"`
	Name     string         `json:"name" validate:"required"`
	Dose     float64        `json:"dose" validate:"required"`
	Course   string         `json:"course" validate:"required"`
	Category string         `json:"category" validate:"required"`
	Option   string         `json:"option" validate:"required"`
	Similar  []NameResponse `json:"similar" validate:"required"`
}
