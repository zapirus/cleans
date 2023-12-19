package usecase

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"clean/handlers"
	"clean/pkg/types"
)

type UseCase struct {
	repo      RepoCaseInterface
	apiClient ApiClientInterface
	mail      MailClientInterface
}

type RepoCaseInterface interface {
	CreateUser(ctx context.Context, user *handlers.RequestUser) (*string, error)
	FindUser(ctx context.Context, val map[string]string) (*types.User, error)
}

type ApiClientInterface interface {
	SendUserClient(ctx context.Context, user string) error
}

type MailClientInterface interface {
	SendEmail(subject, content string, to []string) error
}

func New(repo RepoCaseInterface, apiClient ApiClientInterface, mail MailClientInterface) *UseCase {
	return &UseCase{
		repo:      repo,
		apiClient: apiClient,
		mail:      mail,
	}
}

var cache = make(map[string]string)
var mp = make(map[string]string)

func (u *UseCase) Register(ctx context.Context, user handlers.RequestUser) (*string, error) {
	hashString := u.generateHashPass(user.Password)
	user.Password = hashString

	guid := u.generateGuid()

	user.Guid = guid

	us, err := u.repo.CreateUser(ctx, &user)

	return us, wrapError(err)
}

func (u *UseCase) Login(ctx context.Context, login, password string) (*string, error) {
	hashString := u.generateHashPass(password)
	password = hashString

	mp["login"] = login
	mp["password"] = password

	us, err := u.repo.FindUser(ctx, mp)

	return &us.Login, wrapError(err)
}

func (u *UseCase) Reset(ctx context.Context, login string) (*string, error) {
	mp["login"] = login

	user, err := u.repo.FindUser(ctx, mp)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	if user.Login != "" {
		val := u.generateCode()
		strVal := strconv.Itoa(val)
		cache[login] = strVal
		res := strconv.Itoa(val)
		u.mail.SendEmail("Код подтверждения", res, []string{user.Email})
	}

	return &user.Email, err
}

func (u *UseCase) Verify(ctx context.Context, login, verifyCode string) (string, error) {
	mp["login"] = login
	for key, val := range cache {

		if key == login && val == verifyCode {
			user, err := u.repo.FindUser(ctx, mp)
			if err != nil {
				return "", err
			}
			err = u.mail.SendEmail("Восстановление пароля", user.Login+" "+user.Password, []string{user.Email})
			if err != nil {
				return "Ошибка отправки письма на почту", err

			}
			return "Логин / Пароль отправлены на почту!", nil
		}

	}
	return "", nil
}

func (u *UseCase) Resend(ctx context.Context, login string) (string, error) {

	user, err := u.repo.FindUser(ctx, map[string]string{"login": login})
	if err != nil {
		return "", err
	}
	if cache[login] != "" {
		u.mail.SendEmail("Код подтверждения", cache[login], []string{user.Email})

	} else {
		cache[login] = strconv.Itoa(u.generateCode())
		u.mail.SendEmail("Код подтверждения", cache[login], []string{user.Email})
	}
	return "Код подтверждения отправлен заново", nil
}
