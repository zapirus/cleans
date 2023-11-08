package handlers

import (
	"context"

	"github.com/labstack/echo"
)

type Api struct {
	app Application
}

type Application interface {
	GetUserUseCase(ctx context.Context, name string) (string, error)
	SendUserUseCase(ctx context.Context, user string) error
}

func NewHandler(service Application) *Api {
	return &Api{
		app: service,
	}
}

func (a *Api) Setup(s *echo.Echo) {
	s.GET("/user", a.HandlerTake)
}

func (a *Api) HandlerTake(c echo.Context) error {
	as, _ := a.app.GetUserUseCase(c.Request().Context(), "user")
	return c.JSON(200, as)
}
