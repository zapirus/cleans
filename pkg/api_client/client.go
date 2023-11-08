package api_client

import (
	"context"
	"fmt"
)

type Client struct {
	cfg *Config
}

type Config struct {
	BaseUrl string `env:"BASE_URL"`
}

func New(cfg *Config) *Client {
	return &Client{cfg: cfg}
}

func (c *Client) Start(ctx context.Context) error {
	fmt.Println("start client")
	return nil
}

func (c *Client) Shutdown(ctx context.Context) error {
	fmt.Println("stop client")
	return nil
}

func (c *Client) SendUser(ctx context.Context, user string) error {
	return nil
}
