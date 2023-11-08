package runner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo"
)

type StartStopInterface interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Runner struct {
	services []StartStopInterface
}

func New(services ...StartStopInterface) *Runner {
	return &Runner{
		services: services,
	}
}
func (r *Runner) Run(e *echo.Echo) error {
	ctx := context.Background()

	for _, service := range r.services {
		if err := service.Start(ctx); err != nil {
			return fmt.Errorf("failed to start usecase")
		}
	}

	go func() {
		if err := e.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println("\tshutting down the server")

	for i := len(r.services) - 1; i >= 0; i-- {
		if err := r.services[i].Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to stop usecase")
		}
	}

	return nil
}
