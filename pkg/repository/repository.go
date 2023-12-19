package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"clean/handlers"
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

func (p *Postgres) CreateUser(ctx context.Context, user *handlers.RequestUser) (*string, error) {
	var userID string
	query := "INSERT INTO users (guid, login, password, name, email) VALUES ($1, $2, $3, $4, $5) RETURNING guid"

	if err := p.conn.QueryRow(ctx, query, user.Guid, user.Login, user.Password, user.Name, user.Email).Scan(&userID); err != nil {
		return nil, err
	}
	return &userID, nil

}

func (p *Postgres) FindUser(ctx context.Context, val map[string]string) (*types.User, error) {
	var user types.User

	query := "SELECT login, password, name, email FROM users WHERE login = $1"

	err := p.conn.QueryRow(ctx, query, val["login"]).Scan(&user.Login, &user.Password, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
