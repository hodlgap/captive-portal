package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type App struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	GracefulTimeoutSec int    `yaml:"graceful_timeout_second"`
}

type Config struct {
	App App `yaml:"app"`
}

// Parse yaml to go struct
func Parse(p string) (*Config, error) {
	abs, err := filepath.Abs(p)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get absolute path")
	}

	f, err := os.Open(abs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}(f)

	cfg := &Config{}
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, errors.Wrap(err, "failed to decode yaml")
	}

	return cfg, nil
}
