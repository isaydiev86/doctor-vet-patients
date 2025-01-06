package dto

type Reference struct {
	ID   int64
	Name string
	Type string
}

type Symptoms struct {
	ID   int64
	Name string
}

type NameResponse struct {
	ID   int64
	Name string
}

type Preparations struct {
	ID       int64
	Name     string
	Dose     float64
	Course   string
	Category string
	Option   string
}

type PreparationsWithSimilar struct {
	ID       int64
	Name     string
	Dose     float64
	Course   string
	Category string
	Option   string
	Similar  []NameResponse
}
