package config

import "fmt"

const apiPrefix = "API"

type Config struct {
	Api *ApiConfig
}

func NewConfig() (*Config, error) {
	var err error
	conf := new(Config)
	conf.Api, err = newApiConfig()
	if err != nil {
		return nil, fmt.Errorf("%v: %w", ErrApiConfigLoading, err)
	}

	return conf, nil
}
