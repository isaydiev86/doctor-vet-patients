package models

// PatientUpdate for update
type PatientUpdate struct {
	ID         int64    `json:"id" validate:"required"`
	Fio        string   `json:"fio"  validate:"required"`                         // ФИО пациента, обязательно
	Phone      string   `json:"phone" validate:"required"`                        // Телефон обязательно
	Name       string   `json:"name" validate:"required"`                         // Имя животного, обязательно
	Animal     *string  `json:"animal"`                                           // Животное, обязательно
	Breed      *string  `json:"breed"`                                            // Порода животного, обязательно
	Age        *float64 `json:"age"`                                              // Возраст животного в годах, >= 0
	Gender     string   `json:"gender" validate:"required,oneof=мужской женский"` // Пол животного, "мужской" или "женский"
	IsNeutered bool     `json:"isNeutered"`                                       // Информация о стерилизации
}

// Patient представляет пациента и связанную с ним информацию.
type Patient struct {
	ID         int64    `json:"id"`
	Fio        string   `json:"fio"  validate:"required"`                         // ФИО пациента, обязательно
	Phone      string   `json:"phone" validate:"required"`                        // Телефон обязательно
	Name       string   `json:"name" validate:"required"`                         // Имя животного, обязательно
	Animal     *string  `json:"animal"`                                           // Животное, обязательно
	Breed      *string  `json:"breed"`                                            // Порода животного, обязательно
	Age        *float64 `json:"age"`                                              // Возраст животного в годах, >= 0
	Gender     string   `json:"gender" validate:"required,oneof=мужской женский"` // Пол животного, "мужской" или "женский"
	IsNeutered bool     `json:"isNeutered"`                                       // Информация о стерилизации
}

type Response struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
