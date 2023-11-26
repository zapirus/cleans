package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"

	"clean/pkg/api_client"
	"clean/pkg/repository"
)

type Config struct {
	repositoryConf repository.Configs
	apiClientConf  api_client.Config
}

func ParseConf(conf *Config) error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("ошибка при загрузке .env файла: %s", err)
	}

	if err := configor.Load(&conf.repositoryConf, ""); err != nil {
		fmt.Println("Ошибка при загрузке конфигурации:", err)
		return err

	}
	conf.repositoryConf.Host = os.Getenv("host")
	conf.repositoryConf.Port = os.Getenv("port")
	conf.repositoryConf.Username = os.Getenv("username")
	conf.repositoryConf.Password = os.Getenv("password")
	conf.repositoryConf.DBName = os.Getenv("dbname")
	conf.repositoryConf.SSLMode = os.Getenv("sslmode")
	//
	//_ = godotenv.Load(".env")
	//
	//configorLoader := configor.New(&configor.Config{
	//	Silent:               true,
	//	ErrorOnUnmatchedKeys: true,
	//	Environment:          "",
	//	ENVPrefix:            "-",
	//	Debug:                false,
	//	Verbose:              false,
	//	AutoReload:           false,
	//	AutoReloadInterval:   0,
	//	AutoReloadCallback:   nil,
	//})
	//
	//if err := configorLoader.Load(conf.repositoryConf); err != nil {
	//	return fmt.Errorf("loading environment: %w", err)
	//}

	return nil
}
