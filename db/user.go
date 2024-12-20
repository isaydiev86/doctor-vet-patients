package db

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/isaydiev86/doctor-vet-patients/transport/models"
)

func (db *DB) CreateUser(ctx context.Context, userID, name, role string) error {
	query := `
		INSERT INTO users (user_id, fio, role)
		VALUES ($1, $2, $3);
	`
	_, err := db.Exec(ctx, query, userID, name, role)

	if err != nil {
		db.logger.Error("failed to create user", err)
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (db *DB) UserExists(ctx context.Context, userID string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM users 
			WHERE user_id = $1
		);
	`
	var exists bool
	err := pgxscan.Get(ctx, db.DB, &exists, query, userID)
	if err != nil {
		db.logger.Error("failed to check if user exists", err)
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return exists, nil
}

func (db *DB) GetUsers(ctx context.Context, filter dto.UserFilters) ([]*dto.User, error) {
	query := `
		select id, user_id, fio, role
		from users
		where 
		    ($1::TEXT IS NULL OR role = $1)
	`

	var users []*models.User

	err := pgxscan.Select(ctx, db.DB, &users, query, utils.NilIfEmpty(filter.Role))
	if err != nil {
		db.logger.Error("db on GetUsers", err)
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	return mapUserDBToDTO(users), nil
}

func mapUserDBToDTO(rows []*models.User) []*dto.User {
	users := make([]*dto.User, len(rows))

	for i, row := range rows {
		item := &dto.User{
			ID:     row.ID,
			UserID: row.UserID,
			Fio:    row.Fio,
			Role:   row.Role,
		}
		users[i] = item
	}

	return users
}
