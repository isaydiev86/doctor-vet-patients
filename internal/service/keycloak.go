package service

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/pkg/errors"
)

type KeycloakConfig struct {
	URL      string `yaml:"url"`
	Realm    string `yaml:"realm"`
	ClientID string `yaml:"client_id"`
	Secret   string `yaml:"secret"`
	Rsa256   string `yaml:"rsa256"`
}

type KeycloakService struct {
	Config KeycloakConfig
	Realm  string
}

func NewKeycloakService(config KeycloakConfig) *KeycloakService {
	return &KeycloakService{
		Config: config,
		Realm:  config.Realm,
	}
}

func (k *KeycloakService) loginRestApiClient(ctx context.Context) (*gocloak.JWT, error) {
	client := gocloak.NewClient(k.Config.URL)

	token, err := client.LoginClient(ctx, k.Config.ClientID, k.Config.Secret, k.Realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to login the rest client")
	}
	return token, nil
}

func (k *KeycloakService) Login(ctx context.Context, username, password string) (string, error) {
	//token, err := k.loginRestApiClient(ctx)
	//if err != nil {
	//	return "", err
	//}

	client := gocloak.NewClient(k.Config.URL)
	jwt, err := client.Login(ctx, k.Config.ClientID, k.Config.Secret, k.Realm, username, password)
	if err != nil {
		return "", err
	}
	return jwt.AccessToken, nil
}

func (k *KeycloakService) RetrospectToken(ctx context.Context, accessToken string) (*gocloak.IntroSpectTokenResult, error) {

	client := gocloak.NewClient(k.Config.URL)

	rptResult, err := client.RetrospectToken(ctx, accessToken, k.Config.ClientID, k.Config.Secret, k.Realm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrospect token")
	}
	return rptResult, nil
}
