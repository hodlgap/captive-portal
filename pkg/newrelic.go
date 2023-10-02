package pkg

import (
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
)

func NewNewrelic(c config.Config) (*newrelic.Application, error) {
	if c.App.ENV != config.EnvProd {
		return nil, nil
	}

	nr, err := newrelic.NewApplication(
		newrelic.ConfigAppName(c.App.Name),
		newrelic.ConfigLicense(c.Newrelic.LicenseKey),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return nr, nil
}
