package postgres

import (
	"context"

	"doctor-vet-patients/internal/dto"
)

func (r *repo) CreatePatient(ctx context.Context, patient dto.Patient) error {
	/// TODO в одной транзакции создать основные данные пациента и первоначальные данные для лечения
	/// TODO это делать из бизнес логики или сразу тут ?
	return nil
}
