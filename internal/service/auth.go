package service

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/isaydiev86/doctor-vet-patients/internal/dto"
	"github.com/pkg/errors"
)

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*dto.RefreshTokenShort, error) {
	refresh, err := s.Keycloak.RefreshToken(ctx, refreshToken)
	if err != nil {
		s.Logger.Error("failed on keycloak refreshToken", err)
		return nil, errors.WithMessage(err, "failed on keycloak refreshToken")
	}
	return &dto.RefreshTokenShort{
		AccessToken:  refresh.AccessToken,
		RefreshToken: refresh.RefreshToken,
	}, nil
}

func (s *Service) Login(ctx context.Context, login dto.LoginRequest) (*dto.LoginResponse, error) {
	jwt, err := s.Keycloak.Login(ctx, login.Username, login.Password)
	if err != nil {
		s.Logger.Error("failed on keycloak login", err)
		return nil, errors.WithMessage(err, "failed on keycloak login")
	}

	userID, err := s.Keycloak.ExtractUserIDFromToken(jwt.AccessToken)
	if err != nil {
		s.Logger.Error("failed on keycloak get userID", err)
		return nil, err
	}

	name, err := s.Keycloak.ExtractNameFromToken(jwt.AccessToken)
	if err != nil {
		s.Logger.Error("failed on keycloak get name", err)
		return nil, err
	}

	role, err := s.Keycloak.ExtractRoleFromToken(jwt.AccessToken)
	if err != nil {
		s.Logger.Error("failed on keycloak get role", err)
		return nil, err
	}

	// проверить по userID - есть ли такой в бд, если нету создать
	exist, err := s.svc.DB.UserExists(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !exist {
		err = s.svc.DB.CreateUser(ctx, userID, name, role)
		if err != nil {
			return nil, err
		}
	}

	return mapKeycloakToDTO(jwt, userID, name, role), err
}

func mapKeycloakToDTO(jwt *gocloak.JWT, userID, name, role string) *dto.LoginResponse {
	return &dto.LoginResponse{
		Role:         role,
		Name:         name,
		UserID:       userID,
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
	}
}
