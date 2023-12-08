package mail_test

import (
	"testing"

	"clean/pkg/mail"
)

func TestMail_SendEmail(t *testing.T) {
	cfg := &mail.EmailConfig{
		NameMail: "zapirus",
		AddrMail: "dsfdfgdfgghfh@gmail.com",
		PassMail: "nfvxurijwokcwljk",
	}

	emailCfg := mail.NewMail(cfg)

	subject := "Сброс пароля"
	content := "<h1>логин, пароль!</h1>"
	to := []string{"ibragimov-009@mail.ru"}

	err := emailCfg.SendEmail(subject, content, to)
	if err != nil {
		t.Fatalf("Не удалось отправить письмо: %v", err)
	}
}
