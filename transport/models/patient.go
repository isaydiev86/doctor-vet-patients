package models

// Patient представляет пациента и связанную с ним информацию.
type Patient struct {
	ID         int64   `json:"id" example:"1" validate:"required"`
	DoctorID   string  `json:"doctorId"  validate:"required" format:"uuid"`      // идентификатор доктора из кейклоак UUID
	Fio        string  `json:"fio"  validate:"required"`                         // ФИО пациента, обязательно
	Phone      string  `json:"phone" validate:"required,e164"`                   // Телефон в формате E.164, обязательно
	Address    string  `json:"address" validate:"required"`                      // Адрес проживания, обязательно
	Animal     string  `json:"animal" validate:"required"`                       // Животное, обязательно
	Name       string  `json:"name" validate:"required"`                         // Имя животного, обязательно
	Breed      string  `json:"breed" validate:"required"`                        // Порода животного, обязательно
	Age        float64 `json:"age" validate:"gte=0"`                             // Возраст животного в годах, >= 0
	Gender     string  `json:"gender" validate:"required,oneof=мужской женский"` // Пол животного, "male" или "female"
	Status     string  `json:"status" validate:"required"`                       // Статус пациента, обязательно
	IsNeutered bool    `json:"isNeutered"`                                       // Информация о стерилизации
}

type Response struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
