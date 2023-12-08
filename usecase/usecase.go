package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"

	"clean/pkg/types"
)

type UseCase struct {
	repo      RepoCaseInterface
	apiClient ApiClientInterface
	mail      MailClientInterface
}

type RepoCaseInterface interface {
	Login(ctx context.Context, login, password string) (*string, error)
	Register(ctx context.Context, user *types.User) (*string, error)
	Verify(ctx context.Context, mail, verifyCode string) (string, error)
	Reset(ctx context.Context, login string) (*string, error)
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

func (u *UseCase) Register(ctx context.Context, user *types.User) (*string, error) {
	hash := sha256.New()
	hash.Write([]byte(user.Password))
	hashValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashValue)
	user.Password = hashString

	us, err := u.repo.Register(ctx, user)

	return us, wrapError(err)
}

func (u *UseCase) Login(ctx context.Context, login, password string) (*string, error) {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashValue)
	password = hashString

	us, err := u.repo.Login(ctx, login, password)

	return us, wrapError(err)
}

func (u *UseCase) Reset(ctx context.Context, login string) (*string, error) {

	repoReset, err := u.repo.Reset(ctx, login)
	if *repoReset != "" {
		val := u.GenerateCode()
		strVal := strconv.Itoa(val)
		cache[*repoReset] = strVal
		res := strconv.Itoa(val)
		u.mail.SendEmail("Код подтверждения", res, []string{*repoReset})
	}

	return repoReset, err
}

func (u *UseCase) Verify(ctx context.Context, mail, verifyCode string) (string, error) {
	for key, val := range cache {
		if key == mail && verifyCode == val {
			logPass, err := u.repo.Verify(ctx, mail, "")
			if err != nil {
				return "", err
			}
			if err = u.mail.SendEmail("Восстановление пароля", logPass, []string{mail}); err != nil {
				return "Ошибка отправки письма на почту", err

			}
			return "Логин / Пароль отправлены на почту!", nil
		}

	}
	return "", nil
}

func (u *UseCase) Resend(mail string) (string, error) {
	if cache[mail] != "" {
		u.mail.SendEmail("Код подтверждения", cache[mail], []string{mail})

	} else {
		cache[mail] = strconv.Itoa(u.GenerateCode())
		u.mail.SendEmail("Код подтверждения", cache[mail], []string{mail})
	}
	return "Код подтверждения отправлен заново", nil
}
