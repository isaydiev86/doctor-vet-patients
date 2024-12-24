package service

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
)

const (
	RoleAdmin   = "admin"
	RoleDoctor  = "doctor"
	RoleUnknown = "unknown"
)

func (s *Service) GetUsers(ctx context.Context, role string) ([]*dto.User, error) {
	usersData, err := s.Keycloak.GetUsers(ctx, role)
	if err != nil {
		return nil, err
	}

	return mapUsersDataToDTO(usersData), nil
}

func mapUsersDataToDTO(usersData []*gocloak.User) []*dto.User {
	users := make([]*dto.User, 0, len(usersData))

	for i, u := range usersData {
		if u == nil {
			continue
		}

		user := &dto.User{
			ID:     int64(i),
			UserID: utils.FromPtr(u.ID),
			Fio:    utils.FromPtr(u.FirstName) + " " + utils.FromPtr(u.LastName),
			Role:   getRole(utils.FromPtr(u.RealmRoles)),
		}

		users = append(users, user)
	}

	return users
}

func getRole(roles []string) string {
	hasAdmin := false
	hasDoctor := false

	for _, r := range roles {
		if r == RoleAdmin {
			hasAdmin = true
		}
		if r == RoleDoctor {
			hasDoctor = true
		}
	}

	if hasAdmin {
		return RoleAdmin
	}
	if hasDoctor {
		return RoleDoctor
	}
	return RoleUnknown
}
