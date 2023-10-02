package config

import (
	"github.com/labstack/gommon/log"
	"strings"
)

type ENV string

const (
	EnvUnknown ENV = ""
	EnvDev     ENV = "dev"
	EnvProd    ENV = "prod"
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

type App struct {
	Name               string `yaml:"name"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	ENV                `yaml:"env"`
	GracefulTimeoutSec int `yaml:"graceful_timeout_second"`
	LogLevel           `yaml:"log_level"`
}
