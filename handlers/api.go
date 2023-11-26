package handlers

import (
	"context"

	"github.com/labstack/echo"

	"clean/pkg/types"
)

const statusSuccess status = "success"
const statusError status = "error"

type Api struct {
	app Application
}

type Application interface {
	Login(ctx context.Context, login, password string) (*string, error)
	Register(ctx context.Context, user *types.User) (*string, error)
	Verify(ctx context.Context, guid, verify string) error
	Reset(ctx context.Context, login, password, retryPassword string) error
	Resend(ctx context.Context, login, password string) error
}

type response struct {
	Body    any    `json:"body,omitempty"`
	Status  status `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type status string

func NewHandler(service Application) *Api {
	return &Api{
		app: service,
	}
}

func (a *Api) Setup(s *echo.Echo) {

	v1 := s.Group("/v1")
	v1.POST("", a.login)   // change
	v1.DELETE("", a.login) // delete

	v1.POST("/register", a.register)
	v1.POST("/login", a.login)
	v1.POST("/verify", a.verify)
	v1.POST("/reset", a.reset)

	v1.POST("/resend", a.resend)
}
