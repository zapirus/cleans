package handlers

import (
	"context"

	"github.com/labstack/echo"
)

type Api struct {
	app Application
}

type Application interface {
	GetUser(ctx context.Context, name string) (string, error)
	SendUser(ctx context.Context, user string) error
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
	a.app.SendUser(c.Request().Context(), "user")
	return c.JSON(200, "barra")
}
