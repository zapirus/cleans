package service

import (
	"context"
	"fmt"
)

type Service struct {
	repo      ServInterface
	apiClient ApiClientInterface
}

type ServInterface interface {
	GetUser(ctx context.Context, name string) (string, error)
	SendUser(ctx context.Context, user string) error
}

type ApiClientInterface interface {
	SendUser(ctx context.Context, user string) error
}

func New(repo ServInterface, apiClient ApiClientInterface) *Service {
	return &Service{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (u *Service) GetUser(ctx context.Context, guid string) (string, error) {
	user, err := u.repo.GetUser(ctx, guid)
	if err != nil {
		return "", err
	}

	if user != "verify" {
		return "not verify", fmt.Errorf("no verify")
	}

	return user, nil
}

func (u *Service) SendUser(ctx context.Context, user string) error {
	return u.repo.SendUser(ctx, user)
}

func (u *Service) Start(ctx context.Context) error {
	fmt.Println("start service")

	return nil
}

func (u *Service) Shutdown(ctx context.Context) error {
	fmt.Println("stop service")
	return nil
}
