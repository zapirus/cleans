package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(echo *echo.Echo) *Server {
	return &Server{echo: echo}
}

func (c *Server) Start(ctx context.Context) error {
	fmt.Println("Starting server")
	go func() {
		if err := c.echo.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			c.echo.Logger.Fatal("shutting down the server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := c.echo.Shutdown(ctx); err != nil {
		c.echo.Logger.Fatal(err)
	}
	return nil
}

func (c *Server) Shutdown(ctx context.Context) error {
	fmt.Println("stop server")
	return nil
}
