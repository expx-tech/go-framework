package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ApiConfig struct {
	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"DEBUG"`
}

func newApiConfig() (*ApiConfig, error) {
	api := new(ApiConfig)
	err := envconfig.Process(apiPrefix, api)
	if err != nil {
		return nil, err
	}

	return api, nil
}
