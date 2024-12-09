package config

import (
	"fmt"
	"os"

	"github.com/isaydiev86/doctor-vet-patients/pkg/dbutil"
	"github.com/isaydiev86/doctor-vet-patients/pkg/keycloak"
	"github.com/isaydiev86/doctor-vet-patients/transport"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB       *dbutil.Config    `yaml:"db"`
	Srv      *transport.Config `yaml:"server"`
	Keycloak *keycloak.Config  `yaml:"keycloak"`
}

func New() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	return &cfg, nil
}
