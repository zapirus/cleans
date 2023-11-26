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
	Host     string `env:"host"`
	Port     string `env:"port"`
	Username string `env:"username"`
	Password string `env:"password"`
	DBName   string `env:"dbname"`
	SSLMode  string `env:"sslmode"`
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

	connDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.cfg.Username, p.cfg.Password, p.cfg.Host, p.cfg.Port, p.cfg.DBName, p.cfg.SSLMode)
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
	query := "SELECT id FROM users WHERE login = $1 AND password = $2 LIMIT 1"
	var userID string
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
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	var userID string
	err := p.conn.QueryRow(ctx, query, user.Login, user.Password, user.Name, user.Email, user.Guid).Scan(&userID)
	if err != nil {
		return nil, err
	}
	return &userID, nil
}

func (p *Postgres) Verify(ctx context.Context, guid, verify string) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) Reset(ctx context.Context, login, password, retryPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) Resend(ctx context.Context, login, password string) error {
	//TODO implement me
	panic("implement me")
}
