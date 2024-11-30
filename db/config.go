package db

import "time"

type Config struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Schema          string        `yaml:"schema"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	Name            string        `yaml:"name"`
	ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time" default:"1h"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" default:"30m"`
	MaxOpenConns    int           `yaml:"max_open_conns" default:"4"`
	SSL             bool          `yaml:"ssl" default:"false"`
}
