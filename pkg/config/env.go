package config

import (
	"strings"

	"github.com/caarlos0/env/v9"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type AppEnv string

const (
	AppEnvUnknown AppEnv = ""
	AppEnvDEV     AppEnv = "dev"
	AppEnvPROD    AppEnv = "prod"
)

type LogLevel string

const (
	LogLevelDEBUG = "DEBUG"
	LogLevelINFO  = "INFO"
)

func (l LogLevel) ToLevel() log.Lvl {
	switch strings.ToUpper(string(l)) {
	case LogLevelDEBUG:
		return log.DEBUG
	case LogLevelINFO:
		return log.INFO
	default:
		return log.INFO
	}
}

type Openwrt struct {
	EncryptionKey string `env:"ENCRYPTION_KEY"`
}

type Newrelic struct {
	LicenseKey string `env:"LICENSE_KEY"`
}

type Redis struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB"`
}

type DB struct {
	Host     string `env:"HOST"`
	Port     int    `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
	SslMode  string `env:"SSL_MODE"`
}

type App struct {
	Name               string   `env:"NAME"`
	Host               string   `env:"HOST"`
	Port               int      `env:"PORT"`
	Env                AppEnv   `env:"ENV"`
	GracefulTimeoutSec int      `env:"GRACEFUL_TIMEOUT_SECOND"`
	LogLevel           LogLevel `env:"LOG_LEVEL"`
}

type Config struct {
	App      `envPrefix:"APP_"`
	Newrelic `envPrefix:"NEWRELIC_"`
	Openwrt  `envPrefix:"OPENWRT_"`
	Redis    `envPrefix:"REDIS_"`
	DB       `envPrefix:"DB_"`
}

// FromEnv parses env to Config
func FromEnv() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to parse env")
	}

	return cfg, nil
}
