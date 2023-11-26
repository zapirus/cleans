package main

import (
	"log"

	"github.com/labstack/echo"

	"clean/handlers"
	"clean/pkg/api_client"
	"clean/pkg/repository"
	"clean/pkg/runner"
	"clean/pkg/server"
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

	echoServer := echo.New()
	serv.Setup(echoServer)
	newEchoServer := server.NewServer(echoServer)

	runServices := []runner.StartStopInterface{
		apiClient,
		repo,
		newEchoServer,
	}

	r := runner.New(runServices...)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

}
