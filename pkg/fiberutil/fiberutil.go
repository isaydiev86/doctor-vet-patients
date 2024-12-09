package fiberutil

import "time"

type Config struct {
	Address          string          `yaml:"address" default:"0.0.0.0:8081"`
	StartTimeout     time.Duration   `yaml:"start_timeout" default:"1s"`
	ReadTimeout      time.Duration   `yaml:"read_timeout" default:"30s"`
	WriteTimeout     time.Duration   `yaml:"write_timeout" default:"30s"`
	BodyLimit        int             `yaml:"body_limit" default:"4194304"`
	Cache            []CacheConfig   `yaml:"cache" default:"-"`
	DisableKeepalive bool            `yaml:"disable_keepalive" default:"false"`
	Timeout          []TimeoutConfig `yaml:"timeout" default:"-"`
	ReadBufferSize   int             `yaml:"read_buffer_size" default:"4096"`
	//RateLimit        []RateLimitConfig `yaml:"rate_limit" default:"-"`
}
type RateLimitConfig struct {
	Method string        `yaml:"method"`
	Path   string        `yaml:"path"`
	Period time.Duration `yaml:"period"`
	Number int           `yaml:"number"`
	Type   string        `yaml:"type" default:"-"`
}
type TimeoutConfig struct {
	Method  string        `yaml:"method"`
	Path    string        `yaml:"path"`
	Timeout time.Duration `yaml:"timeout"`
}
type CacheConfig struct {
	IsCacheControl bool          `yaml:"is_cache_control"`
	Method         string        `yaml:"method"`
	Path           string        `yaml:"path"`
	Period         time.Duration `yaml:"period"`
}
