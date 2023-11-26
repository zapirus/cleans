package usecase

import (
	"context"

	"clean/pkg/types"
)

type UseCase struct {
	repo      RepoCaseInterface
	apiClient ApiClientInterface
}

type RepoCaseInterface interface {
	Login(ctx context.Context, login, password string) (*string, error)
	Register(ctx context.Context, user *types.User) (*string, error)
	Verify(ctx context.Context, guid, verify string) error
	Reset(ctx context.Context, login, password, retryPassword string) error
	Resend(ctx context.Context, login, password string) error
}

type ApiClientInterface interface {
	SendUserClient(ctx context.Context, user string) error
}

func New(repo RepoCaseInterface, apiClient ApiClientInterface) *UseCase {
	return &UseCase{
		repo:      repo,
		apiClient: apiClient,
	}
}

func (u *UseCase) Login(ctx context.Context, login, password string) (*string, error) {
	us, err := u.repo.Login(ctx, login, password)
	return us, err
}

func (u *UseCase) Register(ctx context.Context, user *types.User) (*string, error) {
	us, err := u.repo.Register(ctx, user)
	return us, err
}

func (u *UseCase) Verify(ctx context.Context, guid, verify string) error {
	err := u.repo.Verify(ctx, guid, verify)
	return err
}

func (u *UseCase) Reset(ctx context.Context, login, password, retryPassword string) error {
	err := u.repo.Reset(ctx, login, password, retryPassword)
	return err
}
func (u *UseCase) Resend(ctx context.Context, login, password string) error {
	u.repo.Resend(ctx, login, password)
	return nil
}
