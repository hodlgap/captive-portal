package config

import (
	"strings"

	"github.com/labstack/gommon/log"
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
