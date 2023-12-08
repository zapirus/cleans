package main

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"

	"clean/pkg/api_client"
	"clean/pkg/mail"
	"clean/pkg/repository"
)

type Config struct {
	RepositoryConf repository.Configs
	ApiClientConf  api_client.Config
	MailClientConf mail.EmailConfig
}

func ParseConf(conf *Config) error {
	_ = godotenv.Load()

	configorLoader := configor.New(&configor.Config{
		Silent:               true,
		ErrorOnUnmatchedKeys: true,
		Environment:          "",
		ENVPrefix:            "-",
		Debug:                false,
		Verbose:              false,
		AutoReload:           false,
		AutoReloadInterval:   0,
		AutoReloadCallback:   nil,
	})

	if err := configorLoader.Load(conf); err != nil {
		return fmt.Errorf("loading environment: %w", err)
	}

	return nil
}
