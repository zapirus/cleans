package runner

import (
	"context"
	"fmt"
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
func (r *Runner) Run() error {
	ctx := context.Background()

	for _, service := range r.services {
		if err := service.Start(ctx); err != nil {
			return fmt.Errorf("failed to start: %s", err)
		}
	}

	fmt.Println("\tshutting down the services")

	for i := len(r.services) - 1; i >= 0; i-- {
		if err := r.services[i].Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to stop usecase")
		}
	}

	return nil
}
