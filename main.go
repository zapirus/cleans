package main

import (
	"log"

	"github.com/labstack/echo"

	"clean/handlers"
	"clean/pkg/api_client"
	"clean/pkg/repository"
	"clean/pkg/runner"
	"clean/usecase"
)

func main() {
	var cfg Config
	if err := ParseConf(&cfg); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgres(&cfg.repositoryConf)
	apiClient := api_client.New(&cfg.apiClientConf)
	services := usecase.New(repo, apiClient)

	serv := handlers.NewHandler(services)

	e := echo.New()
	serv.Setup(e)
	runServ := []runner.StartStopInterface{
		apiClient,
		repo,
		services,
	}

	r := runner.New(runServ...)
	if err := r.Run(e); err != nil {
		log.Fatal(err)
	}

}
