package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Openwrt struct {
	EncryptionKey string `yaml:"encryption_key"`
}

type Newrelic struct {
	LicenseKey string `yaml:"license_key"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type App struct {
	Name               string `yaml:"name"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	ENV                `yaml:"env"`
	GracefulTimeoutSec int `yaml:"graceful_timeout_second"`
	LogLevel           `yaml:"log_level"`
}

type Config struct {
	App      `yaml:"app"`
	Newrelic `yaml:"newrelic"`
	Openwrt  `yaml:"openwrt"`
	Redis    `yaml:"redis"`
}

// Parse yaml to go struct
func Parse(p string) (Config, error) {
	cfg := Config{}

	abs, err := filepath.Abs(p)
	if err != nil {
		return cfg, errors.Wrap(err, "failed to get absolute path")
	}

	f, err := os.Open(abs)
	if err != nil {
		return cfg, errors.Wrap(err, "failed to open file")
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}(f)

	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to decode yaml")
	}

	return cfg, nil
}
