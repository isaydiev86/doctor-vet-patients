package keycloak

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/isaydiev86/doctor-vet-patients/pkg/utils"
	"github.com/pkg/errors"
)

const MAX_USERS = 10

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
func (k *Service) GetToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	}

	if token == "" {
		return "", errors.New("invalid Authorization header format")
	}

	return token, nil
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

func (k *Service) GetUsers(ctx context.Context, role string) ([]*gocloak.User, error) {

	t, err := k.client.LoginClient(ctx, k.Config.ClientID, k.Config.Secret, k.Realm)
	if err != nil {
		return nil, err
	}

	params := gocloak.GetUsersByRoleParams{
		Max: gocloak.IntP(100),
	}

	users, err := k.client.GetUsersByRoleName(ctx, t.AccessToken, k.Realm, role, params)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == nil {
			continue
		}

		realmRoles, err := k.client.GetRealmRolesByUserID(ctx, t.AccessToken, k.Realm, *user.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching realm roles for user %s: %w", *user.Username, err)
		}
		roles := make([]string, 0, len(realmRoles))
		for _, role := range realmRoles {
			roles = append(roles, gocloak.PString(role.Name))
		}
		user.RealmRoles = &roles
	}

	return users, nil
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
