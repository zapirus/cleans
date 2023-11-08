package main

import (
	"fmt"

	"github.com/jinzhu/configor"

	"clean/pkg/api_client"
	"clean/pkg/repository"
)

type Config struct {
	repositoryConf repository.Configs
	apiClientConf  api_client.Config
}

func ParseConf(conf *Config) error {
	if err := configor.Load(&conf.repositoryConf, ".env"); err != nil {
		fmt.Println("Ошибка при загрузке конфигурации:", err)
		return err
	}

	return nil
}
