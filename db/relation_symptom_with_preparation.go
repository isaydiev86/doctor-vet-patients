package db

import (
	"context"
	"fmt"
)

func (db *DB) AddRelationSymptomWithPreparation(ctx context.Context, symptomID, preparationID int64) error {
	query := `
		INSERT INTO symptom_relation_preparation (symptom_id, preparation_id)
		VALUES ($1, $2)
		ON CONFLICT (symptom_id, preparation_id) DO NOTHING;
	`
	_, err := db.Exec(ctx, query, symptomID, preparationID)

	if err != nil {
		db.logger.Error("failed to create symptom_relation_preparation", err)
		return fmt.Errorf("failed to create symptom_relation_preparation: %w", err)
	}

	return nil
}
