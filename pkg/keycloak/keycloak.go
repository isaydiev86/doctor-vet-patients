package keycloak

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/pkg/errors"
)

type Config struct {
	URL      string `yaml:"url"`
	Realm    string `yaml:"realm"`
	ClientID string `yaml:"client_id"`
	Secret   string `yaml:"secret"`
	Rsa256   string `yaml:"rsa256"`
}

type Service struct {
	Config Config
	Realm  string
	client *gocloak.GoCloak
}

func New(config Config) *Service {
	client := gocloak.NewClient(config.URL)

	return &Service{
		Config: config,
		Realm:  config.Realm,
		client: client,
	}
}

func (k *Service) ValidateToken(ctx context.Context, token string) (bool, error) {
	_, err := k.client.RetrospectToken(ctx, token, k.Config.ClientID, k.Config.Secret, k.Realm)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (k *Service) GetUserRoles(token string) ([]string, error) {
	_, claims, err := k.client.DecodeAccessToken(context.Background(), token, k.Realm)
	if err != nil {
		return nil, err
	}

	return parseRealmRoles(utils.FromPtr(claims)), nil
}

func (k *Service) Login(ctx context.Context, username, password string) (*gocloak.JWT, error) {
	jw, err := k.client.Login(ctx, k.Config.ClientID, k.Config.Secret, k.Realm, username, password)
	if err != nil {
		return nil, err
	}

	return jw, nil
}

// RefreshToken refreshes the given token.
func (k *Service) RefreshToken(ctx context.Context, refreshToken string) (*gocloak.JWT, error) {
	jw, err := k.client.RefreshToken(ctx, k.Config.ClientID, k.Config.Secret, k.Realm, refreshToken)
	if err != nil {
		return nil, err
	}

	return jw, nil
}

func (k *Service) GetUserByID(ctx context.Context, accessToken, userID string) (*gocloak.User, error) {
	user, err := k.client.GetUserByID(ctx, k.Realm, accessToken, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (k *Service) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {
	rptResult, err := k.client.RetrospectToken(ctx, accessToken, k.Config.ClientID, k.Config.Secret, k.Realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}

	return rptResult, nil
}

// ExtractUserIDFromToken Функция для извлечения userID из access_token
func (k *Service) ExtractUserIDFromToken(accessToken string) (string, error) {
	claims, err := parseTokenClaims(accessToken)
	if err != nil {
		return "", err
	}

	userID, _ := claims["sub"].(string)
	if userID == "" {
		return "", fmt.Errorf("userID не найден в токене")
	}
	return userID, nil
}

// ExtractNameFromToken Функция для извлечения name из access_token
func (k *Service) ExtractNameFromToken(accessToken string) (string, error) {
	claims, err := parseTokenClaims(accessToken)
	if err != nil {
		return "", err
	}

	name, _ := claims["name"].(string)
	if name == "" {
		return "", fmt.Errorf("имя не найдено в токене")
	}
	return name, nil
}

// ExtractRoleFromToken Функция для извлечения роли из access_token
func (k *Service) ExtractRoleFromToken(accessToken string) (string, error) {
	claims, err := parseTokenClaims(accessToken)
	if err != nil {
		return "", err
	}

	if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
		if roles, ok := realmAccess["roles"].([]interface{}); ok && len(roles) > 0 {
			hasAdmin := false
			hasDoctor := false

			for _, r := range roles {
				role, _ := r.(string)
				if role == "admin" {
					hasAdmin = true
				}
				if role == "doctor" {
					hasDoctor = true
				}
			}

			if hasAdmin {
				return "admin", nil
			}
			if hasDoctor {
				return "doctor", nil
			}
		}
	}

	return "", fmt.Errorf("ни одна из приоритетных ролей не найдена в токене")
}

// Общая вспомогательная функция для парсинга токена и получения claims
func parseTokenClaims(accessToken string) (jwt.MapClaims, error) {
	parts := strings.Split(accessToken, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("некорректный токен")
	}

	claims := jwt.MapClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(accessToken, claims)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить токен: %w", err)
	}
	return claims, nil
}
