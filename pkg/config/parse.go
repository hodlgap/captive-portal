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

type Config struct {
	App      `yaml:"app"`
	Newrelic `yaml:"newrelic"`
	Openwrt  `yaml:"openwrt"`
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
