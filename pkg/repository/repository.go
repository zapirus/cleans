package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"clean/pkg/types"
)

type Configs struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USERS"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DBNAME"`
	SSLMode  string `env:"SSLMODE"`
}

type Postgres struct {
	cfg  *Configs
	conn *pgx.Conn
}

func NewPostgres(cfg *Configs) *Postgres {

	return &Postgres{cfg: cfg}
}

func (p *Postgres) Start(ctx context.Context) error {
	fmt.Println("start postgres")

	connDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", p.cfg.User, p.cfg.Password, p.cfg.Host, p.cfg.Port, p.cfg.DBName, p.cfg.SSLMode)

	conn, err := pgx.Connect(ctx, connDB)
	if err != nil {
		return err
	}
	p.conn = conn
	return nil
}

func (p *Postgres) Shutdown(ctx context.Context) error {
	err := p.conn.Close(ctx)
	if err != nil {
		return err
	}
	fmt.Println("stop postgres")
	return nil
}

func (p *Postgres) Login(ctx context.Context, login, password string) (*string, error) {
	var userID string
	query := "SELECT login FROM users WHERE login = $1 AND password = $2 LIMIT 1"
	err := p.conn.QueryRow(ctx, query, login, password).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &userID, nil
}

func (p *Postgres) Register(ctx context.Context, user *types.User) (*string, error) {
	var userID string
	val := p.checkLogin(ctx, user.Login)

	if val {
		return &userID, fmt.Errorf("пользователь с логином - %s уже существует. Попробуйте другой логин", user.Login)
	}

	query := "INSERT INTO users (login, password, name, email) VALUES ($1, $2, $3, $4) RETURNING login"

	err := p.conn.QueryRow(ctx, query, user.Login, user.Password, user.Name, user.Email).Scan(&userID)
	if err != nil {
		return nil, err
	}

	return &userID, nil
}

func (p *Postgres) Reset(ctx context.Context, login string) (*string, error) {
	var userID string
	query := "SELECT email FROM users WHERE login = $1  LIMIT 1"
	err := p.conn.QueryRow(ctx, query, login).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &userID, nil
}

func (p *Postgres) Verify(ctx context.Context, mail, verifyCode string) (string, error) {
	login := ""
	password := ""

	query := "SELECT login, password FROM users WHERE email = $1"
	err := p.conn.QueryRow(ctx, query, mail).Scan(&login, &password)

	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("user not found")
	}
	return "Ваш логин: " + login + " " + "Ваш пароль: " + password, err
}
