package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpServAddr    = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject, content string, to []string) error
}

type EmailConfig struct {
	NameMail string `env:"NAMEMAIL"`
	AddrMail string `env:"ADDRMAIL"`
	PassMail string `env:"PASSMAIL"`
}

type Mail struct {
	cfg *EmailConfig
}

func NewMail(cfg *EmailConfig) *Mail {
	return &Mail{cfg: cfg}
}

func (u *Mail) SendEmail(subject, content string, to []string) error {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", u.cfg.NameMail, u.cfg.AddrMail)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to

	return e.Send(smtpServAddr, smtp.PlainAuth("", u.cfg.AddrMail, u.cfg.PassMail, smtpAuthAddress))
}
