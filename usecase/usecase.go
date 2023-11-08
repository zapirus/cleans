package usecase

import (
	"context"
	"fmt"
)

type UseCase struct {
	repo      RepositoryInterface
	apiClient ApiClientInterface
}

type RepositoryInterface interface {
	GetUserRepo(ctx context.Context, name string) (string, error)
	SendUserRepo(ctx context.Context, user string) error
}

type ApiClientInterface interface {
	SendUserClient(ctx context.Context, user string) error
}

func New(repo RepositoryInterface, apiClient ApiClientInterface) *UseCase {
	return &UseCase{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (u *UseCase) Start(ctx context.Context) error {
	fmt.Println("start usecase")

	return nil
}

func (u *UseCase) Shutdown(ctx context.Context) error {
	fmt.Println("stop usecase")
	return nil
}

func (u *UseCase) GetUserUseCase(ctx context.Context, name string) (string, error) {
	user, err := u.repo.GetUserRepo(ctx, name)
	if err != nil {
		return "", err
	}

	if user != "verify" {
		return "not verify", fmt.Errorf("no verify")
	}

	return user, nil
}

func (u *UseCase) SendUserUseCase(ctx context.Context, user string) error {
	return u.repo.SendUserRepo(ctx, user)
}
